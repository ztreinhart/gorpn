package main

import (
	"fmt"
	"strings"
)

func main() {
	//Initialize an RPNStack
	var engine RPEngine
	engine.Init()

	//Get some sample input
	input := "3 2 + 2 * 5 - 5 /"

	//Sanitize the input
	input = strings.ToLower(strings.TrimSpace(input))
	tokens := strings.Split(input, " ")
	fmt.Printf("Result: %f\n", engine.Eval(tokens))

	//fmt.Println(tokens)

}
