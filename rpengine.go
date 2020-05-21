package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type RPEngine struct {
	stack      RPNStack
	vars       map[string]interface{}
	macros     map[string]string
	replPrefix string
	replPrompt *prompt.Prompt
	regex      map[string]*regexp.Regexp
}

func (r *RPEngine) Init() {
	r.stack.Init()
	r.vars = make(map[string]interface{})
	r.macros = make(map[string]string)

	r.regex = make(map[string]*regexp.Regexp)
	r.regex["setvar"] = regexp.MustCompile("[a-z]=")
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

func (r *RPEngine) EvalString(input string) float64 {
	input = strings.ToLower(strings.TrimSpace(input))
	tokens := strings.Split(input, " ")
	r.Eval(tokens)
	//TODO
	return 0
}

func (r *RPEngine) Eval(tokens []string) {
	//for _, token := range tokens {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		//Handle boolean literal tokens
		switch token {
		case "true", "t":
			r.stack.Push("true")
			break
		case "false", "f":
			r.stack.Push("false")
			break
		}

		//Handle integer literal tokens
		//TODO

		//Handle float literal tokens
		//TODO
		val, err := strconv.ParseFloat(token, 64)

		switch {
		case err == nil:
			r.stack.Push(val)

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
		case token == "tanh":
			r.tanh()

		//Numeric Utilities
		case token == "ceil":
			r.ceiling()
		case token == "floor":
			r.floor()
		case token == "round":
			r.round()
		case token == "integerPart":
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

		//Display Modes
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
		case r.varFound(token):
			r.getVar(token)
		case r.macroFound(token):
			r.runMacro(token)

		//Other operations
		case token == "help":
			fmt.Println("Help!")
		case token == "exit":
			runtime.Goexit()
		}
	}
}

func (r *RPEngine) replExecutor(in string) {
	r.EvalString(in)

	r.replPrefix = r.stack.AsHorizString() + "> "
	//LivePrefixState.IsEnable = true
}

func (r *RPEngine) replCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "+", Description: "Addition"},
		{Text: "-", Description: "Subtraction"},
		{Text: "*", Description: "Multiplication"},
		{Text: "/", Description: "Division"},
		{Text: "cla", Description: "Clears stack and variables"},
		{Text: "exit", Description: "Exits from the calculator"},
	}
	//return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
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
