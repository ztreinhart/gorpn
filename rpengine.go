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
		case token == "!":
			r.boolNot()
		case token == "!=":
			r.notEqual()
		case token == "%":
			r.mod()
		case token == "++":
			r.increment()
		case token == "--":
			r.decrement()

		//Bitwise Operators
		case token == "&":
			//TODO: Bitwise AND

		case token == "|":
			//TODO: Bitwise OR

		case token == "^":
			//TODO: Bitwise XOR

		case token == "~":
			//TODO: Bitwise NOT

		case token == "<<":
			//TODO: Bitwise shift left

		case token == ">>":
			//TODO: Bitwise shift right

		//Boolean Operations
		case token == "&&":
			//TODO: Boolean AND

		case token == "||":
			//TODO: Boolean OR

		case token == "^^":
			//TODO: Boolean XOR

		//Comparison Operators
		case token == "<":
			//TODO: Less than

		case token == "<=":
			//TODO: Less than or equal to

		case token == "==":
			//TODO: Equal to

		case token == ">":
			//TODO: Greater than

		case token == ">=":
			//TODO: Greater than or equal to

		//Trigonometric Functions
		case token == "acos":
			//TODO: Arc Cosine

		case token == "asin":
			//TODO: Arc Sine

		case token == "atan":
			//TODO: Arc Tangent

		case token == "cos":
			//TODO: Cosine

		case token == "cosh":
			//TODO: Hyberbolic Cosine

		case token == "sin":
			//TODO: Sine

		case token == "sinh":
			//TODO: Hyberbolic Cosine

		case token == "tanh":
			//TODO: Hyperbolic Tangent

		//Numeric Utilities
		case token == "ceil":
			//TODO: Ceiling

		case token == "floor":
			//TODO: Floor

		case token == "round":
			//TODO: Round

		case token == "ip":
			//TODO: Integer part

		case token == "fp":
			//TODO: Floating part

		case token == "sign":
			//TODO: Push -1, 0, 1 depending on the sign of the operand

		case token == "abs":
			//TODO: Absolute value

		case token == "max":
			//TODO: Max

		case token == "min":
			//TODO: Min

		//Display Modes
		case token == "hex":
			//TODO: Switch display mode to hex

		case token == "dec":
			//TODO: Switch display mode to decimal (default)

		case token == "bin":
			//TODO: Switch display mode to binary

		case token == "oct":
			//TODO: Switch display mode to octal

		//Constants
		case token == "e":
			//TODO: Push e onto the stack

		case token == "pi":
			//TODO: Push pi onto the stack

		case token == "rand":
			//TODO: Push a random integer onto the stack

		//Mathematic functions
		case token == "exp":
			//TODO: Exponentiation

		case token == "fact":
			//TODO: Factorial

		case token == "sqrt":
			//TODO: Square root

		case token == "ln":
			//TODO: Natural log

		case token == "log":
			//TODO: Logarithm (base 2 or base 10?)

		case token == "pow":
			//TODO: Raise a number to a power

		//Networking
		case token == "hnl":
			//TODO: Host to network long

		case token == "hns":
			//TODO: Host to network short

		case token == "nhl":
			//TODO: Network to host long

		case token == "nhs":
			//TODO: Network to host short

		//Stack Manipulation
		case token == "pick":
			r.pick()
		case token == "repeat":
			//TODO: Repeat an operation n times, e.g. 3 repeat +
			i++
			r.repeat(tokens[i])
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
		case token == "stack":
			r.stackDisplay()
		case token == "swap":
			r.swap()

		//Macros and Variables:
		case token == "macro":
			//TODO: Defines a macro, e.g. 'macro kib 1024 *'

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

func (r *RPEngine) boolNot() {
	//TODO: Boolean NOT
}

func (r *RPEngine) notEqual() {
	//TODO: Not equal to

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

//Boolean Operators

//Comparison Operators

//Trig Functions

//Numeric Utilities

//Display Modes

//Constants

//Mathematic functions

//Networking

//Stack Manipulation
func (r *RPEngine) pick() {
	n := int(r.stack.Pop())
	r.stack.Push(r.stack.Pick(n))
}

func (r *RPEngine) repeat(op string) {
	//TODO: Repeat an operation n times, e.g. 3 repeat +
	n := int(r.stack.Pop())
	for i := 0; i < n; i++ {
		r.Eval([]string{op})
	}
}

func (r *RPEngine) depth() {
	r.stack.Push(float64(r.stack.Depth()))
}

func (r *RPEngine) drop() {
	_ = r.stack.Pop()
}

func (r *RPEngine) dropn() {
	//TODO: Drop n items from the stack
	r.repeat("drop")
}

func (r *RPEngine) dup() {
	r.stack.Push(r.stack.Peek())
}

func (r *RPEngine) dupn() {
	//TODO: Duplicates the top n stack items in order
	n := int(r.stack.Pop())
	r.stack.Dupn(n)
}

func (r *RPEngine) roll() {
	//TODO: Roll the stack upwards by n

}

func (r *RPEngine) rolld() {
	//TODO: Roll the stack downwards by n

}

func (r *RPEngine) stackDisplay() {
	//TODO: Toggles the stack display from horizontal to vertical

}

func (r *RPEngine) swap() {
	//TODO: Swap the top 2 stack items
	first := r.stack.Pop()
	second := r.stack.Pop()
	r.stack.Push(first)
	r.stack.Push(second)

}

//Macros and variables
