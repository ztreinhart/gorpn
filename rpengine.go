package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"math/big"
	"regexp"
	"runtime"
	"strings"
)

type RPEngine struct {
	stack      RPNStack
	vars       map[string]interface{}
	macros     map[string]string
	replPrefix string
	replPrompt *prompt.Prompt
	regex      map[string]*regexp.Regexp
	//precision   uint
	helpCalled  bool
	displayBase int
	stackDisp   string
}

func (r *RPEngine) Init() {
	r.stack.Init()
	r.vars = make(map[string]interface{})
	r.macros = make(map[string]string)

	r.regex = make(map[string]*regexp.Regexp)
	r.regex["setvar"] = regexp.MustCompile("[a-z]=")
	//r.precision = 128
	r.helpCalled = false
	r.displayBase = 10
	r.stackDisp = "horiz"
}

func (r *RPEngine) InitREPL() {
	r.replPrefix = "[]> "
	r.replPrompt = prompt.New(
		r.replExecutor,
		r.replCompleter,
		//prompt.OptionPrefix("[]> "),
		prompt.OptionLivePrefix(r.changeReplPrefix),
		prompt.OptionTitle("gorpn"),
	)
}

func (r *RPEngine) RunREPL() {
	r.replPrompt.Run()
}

func (r *RPEngine) EvalString(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	tokens := strings.Split(input, " ")
	r.Eval(tokens)

	if r.helpCalled {
		return ""
	}
	return r.valString(r.stack.Peek())
}

func (r *RPEngine) Eval(tokens []string) {
	r.helpCalled = false

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		var literal interface{}
		literalFound := false

		//Handle boolean literal tokens
		switch token {
		case "true", "t":
			literalFound = true
			literal = true
		case "false", "f":
			literalFound = true
			literal = false
		}

		//Handle integer literal tokens
		if !literalFound {
			literal, literalFound = new(big.Int).SetString(token, 0)
		}

		//Handle float literal tokens
		if !literalFound {
			literal, literalFound = new(big.Float).SetString(token)
			//literal.(*big.Float).SetPrec(r.precision)
		}

		//Main parsing tree
		switch {
		//Found a literal value
		case literalFound:
			r.stack.Push(literal)

		//Arithmetic Operators
		case token == "+":
			r.plus()
		case token == "-":
			r.minus()
		case token == "*":
			r.multiply()
		case token == "/":
			r.divide()
		case token == "cla":
			r.cla()
		case token == "clr":
			r.clr()
		case token == "clv":
			r.clv()
		case token == "%":
			r.mod()
		case token == "++":
			r.increment()
		case token == "--":
			r.decrement()

		//Bitwise Operators
		case token == "&":
			r.bitwiseAND()
		case token == "|":
			r.bitwiseOR()
		case token == "^":
			r.bitwiseXOR()
		case token == "~":
			r.bitwiseNOT()
		case token == "<<":
			r.bitwiseLeftShift()
		case token == ">>":

		//Boolean Operations
		case token == "!":
			r.boolNOT()
		case token == "&&":
			r.boolAND()
		case token == "||":
			r.boolOR()
		case token == "^^":
			r.boolXOR()

		//Comparison Operators
		case token == "!=":
			r.notEqual()
		case token == "<":
			r.lessThan()
		case token == "<=":
			r.lessThanEqualTo()
		case token == "==":
			r.equalTo()
		case token == ">":
			r.greaterThan()
		case token == ">=":
			r.greaterThanEqualTo()

		//Trigonometric Functions
		case token == "acos":
			r.acos()
		case token == "asin":
			r.asin()
		case token == "atan":
			r.atan()
		case token == "cos":
			r.cos()
		case token == "cosh":
			r.cosh()
		case token == "sin":
			r.sin()
		case token == "sinh":
			r.sinh()
		case token == "tan":
			r.tan()
		case token == "tanh":
			r.tanh()

		//Numeric Utilities
		case token == "ceil":
			r.ceiling()
		case token == "floor":
			r.floor()
		case token == "round":
			r.round()
		case token == "ip":
			r.integerPart()
		case token == "fp":
			r.floatingPart()
		case token == "sign":
			r.sign()
		case token == "abs":
			r.abs()
		case token == "max":
			r.max()
		case token == "min":
			r.min()

		//Calc params and display modes
		case token == "hex":
			r.hexDisplay()
		case token == "dec":
			r.decDisplay()
		case token == "bin":
			r.binDisplay()
		case token == "oct":
			r.octDisplay()
		case token == "stack":
			r.stackDisplay()

		//Constants
		case token == "e":
			r.constE()
		case token == "pi":
			r.constPi()
		case token == "rand":
			r.random()

		//Mathematic functions
		case token == "exp":
			r.exp()
		case token == "fact":
			r.factorial()
		case token == "sqrt":
			r.sqrt()
		case token == "ln":
			r.naturalLog()
		case token == "log":
			r.log()
		case token == "pow":
			r.pow()

		//Networking
		case token == "hnl":
			r.hnl()
		case token == "hns":
			r.hns()
		case token == "nhl":
			r.nhl()
		case token == "nhs":
			r.nhs()

		//Stack Manipulation
		case token == "pick":
			r.pick()
		case token == "depth":
			r.depth()
		case token == "drop":
			r.drop()
		case token == "dropn":
			r.dropn()
		case token == "dup":
			r.dup()
		case token == "dupn":
			r.dupn()
		case token == "roll":
			r.roll()
		case token == "rolld":
			r.rolld()
		case token == "swap":
			r.swap()

		//Macros and Variables:
		case token == "repeat":
			i++
			r.repeat(tokens[i])
		case r.regex["setvar"].MatchString(token):
			r.setVar(token)
		case token == "macro":
			r.setMacro(tokens[i:])
			i = len(tokens)
		case token == "lsmacros":
			r.listMacros()
		case token == "rmmacro":
			i++
			r.rmMacro(tokens[i])
		case token == "clm":
			r.clearMacros()
		case r.varFound(token):
			r.getVar(token)
		case r.macroFound(token):
			r.runMacro(token)

		//Other operations
		case token == "help":
			fmt.Println("Help!")
			r.helpCalled = true
		case token == "exit":
			runtime.Goexit()
		case token == "quit":
			runtime.Goexit()
		case token == "type":
			fmt.Printf("Type: %T\n", r.stack.Peek())
		default:
			fmt.Println("Unrecognized input.")
		}
	}
}

func (r *RPEngine) replExecutor(in string) {
	r.EvalString(in)
	//TODO: Updated display logic.
	//r.replPrefix = r.stack.AsHorizString() + "> "
	r.replPrefix = r.buildREPLPrefix()
	//LivePrefixState.IsEnable = true
}

func (r *RPEngine) changeReplPrefix() (string, bool) {
	return r.replPrefix, true
}

func (r *RPEngine) varFound(key string) bool {
	_, ok := r.vars[key]
	return ok
}

func (r *RPEngine) macroFound(key string) bool {
	_, ok := r.macros[key]
	return ok
}

func (r *RPEngine) valString(val interface{}) string {
	switch x := val.(type) {
	case bool:
		return fmt.Sprintf("%t", x)
	case *big.Int:
		return x.Text(r.displayBase)
	case *big.Float:
		switch r.displayBase {
		case 16:
			return x.Text('x', -1)
		default:
			return x.Text('g', -1)
		}
	default:
		return fmt.Sprintf("%v", x)
	}
}

func (r *RPEngine) stackString() string {
	var b strings.Builder
	if r.stackDisp == "vert" {
		for i := len(r.stack.Stack) - 1; i >= 0; i-- {
			_, _ = fmt.Fprint(&b, r.valString(r.stack.Stack[i]), "\n")
		}
	} else {
		for idx, val := range r.stack.Stack {
			_, _ = fmt.Fprint(&b, r.valString(val), "")
			if idx < len(r.stack.Stack)-1 {
				_, _ = fmt.Fprint(&b, "  ")
			}
		}
	}

	return b.String()
}

func (r *RPEngine) varsString() string {
	if len(r.vars) > 0 {
		var b strings.Builder
		_, _ = fmt.Fprint(&b, "[variables: ")
		for key, value := range r.vars {
			_, _ = fmt.Fprint(&b, key, "=", r.valString(value), " ")
		}
		_, _ = fmt.Fprint(&b, "]")
		return b.String()
	}
	return ""
}

func (r *RPEngine) buildREPLPrefix() string {
	var b strings.Builder
	if r.stackDisp == "vert" {
		_, _ = fmt.Fprint(&b, r.stackString())
		_, _ = fmt.Fprint(&b, r.varsString())
		_, _ = fmt.Fprint(&b, " >")
	} else {
		_, _ = fmt.Fprint(&b, r.varsString())
		_, _ = fmt.Fprint(&b, "[ ", r.stackString(), " ] > ")
	}
	return b.String()
}
