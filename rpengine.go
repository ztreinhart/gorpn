package main

import (
	"runtime"
	"strconv"
	"strings"
)

type RPEngine struct {
	stack  RPNStack
	vars   map[string]float64
	macros map[string]string
}

func (r *RPEngine) Init() {
	r.stack.Init()
	r.vars = make(map[string]float64)
	r.macros = make(map[string]string)
}

func (r *RPEngine) EvalString(input string) float64 {
	//TODO
	input = strings.ToLower(strings.TrimSpace(input))
	tokens := strings.Split(input, " ")
	return r.Eval(tokens)
}

func (r *RPEngine) Eval(tokens []string) float64 {
	for _, token := range tokens {
		//fmt.Println(token)
		val, err := strconv.ParseFloat(token, 64)
		switch {
		case err == nil:
			r.stack.Push(val)

		case token == "+":
			op1 := r.stack.Pop()
			op2 := r.stack.Pop()
			r.stack.Push(op1 + op2)

		case token == "-":
			op1 := r.stack.Pop()
			op2 := r.stack.Pop()
			r.stack.Push(op2 - op1)

		case token == "*":
			op1 := r.stack.Pop()
			op2 := r.stack.Pop()
			r.stack.Push(op1 * op2)

		case token == "/":
			op1 := r.stack.Pop()
			op2 := r.stack.Pop()
			r.stack.Push(op2 / op1)

		case token == "cla":
			r.stack.Init()
			r.vars = make(map[string]float64)

		case token == "clr":
			r.stack.Init()

		case token == "clv":
			r.vars = make(map[string]float64)

		case token == "!":
			//TODO

		case token == "!=":
			//TODO

		case token == "%":
			//TODO: Doublecheck this...
			op1 := int64(r.stack.Pop())
			op2 := int64(r.stack.Pop())
			r.stack.Push(float64(op2 % op1))

		case token == "++":
			r.stack.Push(r.stack.Pop() + 1)

		case token == "--":
			r.stack.Push(r.stack.Pop() - 1)

		case token == "exit":
			runtime.Goexit()
		}
	}
	return r.stack.Peek()
}

func (r *RPEngine) PromptExecutor(in string) {
	r.EvalString(in)

	LivePrefixState.LivePrefix = r.stack.AsHorizString() + "> "
	LivePrefixState.IsEnable = true
}
