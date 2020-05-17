package main

import (
	"fmt"
	"strconv"
)

type RPEngine struct {
	stack RPNStack
	vars map[string]float64
	macros map[string]string
}

func (r *RPEngine) Init() {
	r.stack.Init()
	r.vars = make(map[string]float64)
	r.macros = make(map[string]string)
}

func (r *RPEngine) EvalString(input string) float64 {
	//TODO
	return 0
}

func (r *RPEngine) Eval(tokens []string) float64 {
	//TODO
	for _, token := range tokens {
		fmt.Println(token)
		val, err := strconv.ParseFloat(token, 64)
		switch {
		case err == nil:
			fmt.Printf("Numeric value: %f\n", val)
			r.stack.push(val)
		case token == "+":
			op1 := r.stack.pop()
			op2 := r.stack.pop()
			res := op1 + op2
			r.stack.push(res)
			fmt.Printf("Addition operator. Result: %f\n", res)
		case token == "-":
			op1 := r.stack.pop()
			op2 := r.stack.pop()
			res := op2 - op1
			r.stack.push(res)
			fmt.Printf("Subtraction operator. Result: %f\n", res)
		case token == "*":
			op1 := r.stack.pop()
			op2 := r.stack.pop()
			res := op1 * op2
			r.stack.push(res)
			fmt.Printf("Multiplication operator. Result: %f\n", res)
		case token == "/":
			op1 := r.stack.pop()
			op2 := r.stack.pop()
			res := op2 / op1
			r.stack.push(res)
			fmt.Printf("Division operator. Result: %f\n", res)
		}
	}
	return r.stack.peek()
}