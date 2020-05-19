package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"runtime"
	"strconv"
	"strings"
)

type RPEngine struct {
	stack      RPNStack
	vars       map[string]float64
	macros     map[string]string
	replPrefix string
	replPrompt *prompt.Prompt
}

func (r *RPEngine) Init() {
	r.stack.Init()
	r.vars = make(map[string]float64)
	r.macros = make(map[string]string)
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
	return r.Eval(tokens)
}

func (r *RPEngine) Eval(tokens []string) float64 {
	//for _, token := range tokens {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

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

		case token == "macro":
			r.macro(tokens[i:])
			i = len(tokens)

		case token == "x=":
			//TODO: Assigns a variable, e.g. '1024 x='

		//Other operations
		case token == "help":
			fmt.Println("Help!")

		case token == "exit":
			runtime.Goexit()
		}
	}
	return r.stack.Peek()
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

//Arithmetic Operators
func (r *RPEngine) plus() {
	op1 := r.stack.Pop()
	op2 := r.stack.Pop()
	r.stack.Push(op1 + op2)
}

func (r *RPEngine) minus() {
	op1 := r.stack.Pop()
	op2 := r.stack.Pop()
	r.stack.Push(op2 - op1)
}

func (r *RPEngine) multiply() {
	op1 := r.stack.Pop()
	op2 := r.stack.Pop()
	r.stack.Push(op1 * op2)
}

func (r *RPEngine) divide() {
	op1 := r.stack.Pop()
	op2 := r.stack.Pop()
	r.stack.Push(op2 / op1)
}

func (r *RPEngine) cla() {
	r.stack.Init()
	r.vars = make(map[string]float64)
}

func (r *RPEngine) clr() {
	r.stack.Init()
}

func (r *RPEngine) clv() {
	r.vars = make(map[string]float64)

}

func (r *RPEngine) mod() {
	op1 := int64(r.stack.Pop())
	op2 := int64(r.stack.Pop())
	r.stack.Push(float64(op2 % op1))
}

func (r *RPEngine) increment() {
	r.stack.Push(r.stack.Pop() + 1)
}

func (r *RPEngine) decrement() {
	r.stack.Push(r.stack.Pop() - 1)
}

//Bitwise Operators
func (r *RPEngine) bitwiseAND() {
	//TODO: Bitwise AND

}

func (r *RPEngine) bitwiseOR() {
	//TODO: Bitwise OR

}

func (r *RPEngine) bitwiseXOR() {
	//TODO: Bitwise XOR

}

func (r *RPEngine) bitwiseNOT() {
	//TODO: Bitwise NOT

}

func (r *RPEngine) bitwiseLeftShift() {
	//TODO: Bitwise shift left

}

func (r *RPEngine) bitwiseRightShift() {
	//TODO: Bitwise shift right

}

//Boolean Operators
func (r *RPEngine) boolNOT() {
	//TODO: Boolean NOT
	arg1 := r.stack.Pop()
	arg2 := r.stack.Pop()
	r.stack.Push(float64(arg1 != arg2))
}

func (r *RPEngine) boolAND() {
	//TODO: Boolean AND
}

func (r *RPEngine) boolOR() {
	//TODO: Boolean OR
}

func (r *RPEngine) boolXOR() {
	//TODO: Boolean XOR
}

//Comparison Operators
func (r *RPEngine) notEqual() {
	//TODO: Not equal to

}

func (r *RPEngine) lessThan() {
	//TODO: Less than

}

func (r *RPEngine) lessThanEqualTo() {
	//TODO: Less than or equal to

}

func (r *RPEngine) equalTo() {
	//TODO: Equal to

}

func (r *RPEngine) greaterThan() {
	//TODO: Greater than

}

func (r *RPEngine) greaterThanEqualTo() {
	//TODO: Greater than or equal to

}

//Trig Functions
func (r *RPEngine) acos() {
	//TODO: Arc Cosine

}

func (r *RPEngine) asin() {
	//TODO: Arc Sine

}

func (r *RPEngine) atan() {
	//TODO: Arc Tangent

}

func (r *RPEngine) cos() {
	//TODO: Cosine

}

func (r *RPEngine) cosh() {
	//TODO: Hyberbolic Cosine

}

func (r *RPEngine) sin() {
	//TODO: Sine

}

func (r *RPEngine) sinh() {
	//TODO: Hyberbolic Cosine

}

func (r *RPEngine) tanh() {
	//TODO: Hyperbolic Tangent

}

//Numeric Utilities
func (r *RPEngine) ceiling() {
	//TODO: Ceiling

}

func (r *RPEngine) floor() {
	//TODO: Floor

}

func (r *RPEngine) round() {
	//TODO: Round

}

func (r *RPEngine) integerPart() {
	//TODO: Integer part

}

func (r *RPEngine) floatingPart() {
	//TODO: Floating part

}

func (r *RPEngine) sign() {
	//TODO: Push -1, 0, 1 depending on the sign of the operand

}

func (r *RPEngine) abs() {
	//TODO: Absolute value

}

func (r *RPEngine) max() {
	//TODO: Max

}

func (r *RPEngine) min() {
	//TODO: Min

}

//Display Modes
func (r *RPEngine) stackDisplay() {
	//TODO: Toggles the stack display from horizontal to vertical

}

func (r *RPEngine) hexDisplay() {
	//TODO: Switch display mode to hex

}

func (r *RPEngine) decDisplay() {
	//TODO: Switch display mode to decimal (default)

}

func (r *RPEngine) binDisplay() {
	//TODO: Switch display mode to binary

}

func (r *RPEngine) octDisplay() {
	//TODO: Switch display mode to octal

}

//Constants
func (r *RPEngine) constE() {
	//TODO: Push e onto the stack

}

func (r *RPEngine) constPi() {
	//TODO: Push pi onto the stack

}

func (r *RPEngine) random() {
	//TODO: Push a random integer onto the stack

}

//Mathematic functions
func (r *RPEngine) exp() {
	//TODO: Exponentiation

}

func (r *RPEngine) factorial() {
	//TODO: Factorial

}

func (r *RPEngine) sqrt() {
	//TODO: Square root

}

func (r *RPEngine) naturalLog() {
	//TODO: Natural log

}

func (r *RPEngine) log() {
	//TODO: Logarithm (base 2 or base 10?)

}

func (r *RPEngine) pow() {
	//TODO: Raise a number to a power

}

//Networking
func (r *RPEngine) hnl() {
	//TODO: Host to network long

}

func (r *RPEngine) hns() {
	//TODO: Host to network short

}

func (r *RPEngine) nhl() {
	//TODO: Network to host long

}

func (r *RPEngine) nhs() {
	//TODO: Network to host short

}

//Stack Manipulation
func (r *RPEngine) pick() {
	n := int(r.stack.Pop())
	r.stack.Push(r.stack.Pick(n))
}

func (r *RPEngine) depth() {
	r.stack.Push(float64(r.stack.Depth()))
}

func (r *RPEngine) drop() {
	r.stack.Drop()
}

func (r *RPEngine) dropn() {
	n := int(r.stack.Pop())
	r.stack.Dropn(n)
}

func (r *RPEngine) dup() {
	r.stack.Dup()
}

func (r *RPEngine) dupn() {
	n := int(r.stack.Pop())
	r.stack.Dupn(n)
}

func (r *RPEngine) roll() {
	n := int(r.stack.Pop())
	r.stack.Roll(n)
}

func (r *RPEngine) rolld() {
	n := int(r.stack.Pop())
	r.stack.Rolld(n)
}

func (r *RPEngine) swap() {
	r.stack.Swap()
}

//Macros and variables
func (r *RPEngine) repeat(op string) {
	n := int(r.stack.Pop())
	for i := 0; i < n; i++ {
		r.Eval([]string{op})
	}
}

func (r *RPEngine) macro(ops []string) {

}
