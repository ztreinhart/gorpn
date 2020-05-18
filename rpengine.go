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
	for _, token := range tokens {

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
			op1 := int64(r.stack.Pop())
			op2 := int64(r.stack.Pop())
			r.stack.Push(float64(op2 % op1))

		case token == "++":
			r.stack.Push(r.stack.Pop() + 1)

		case token == "--":
			r.stack.Push(r.stack.Pop() - 1)

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
