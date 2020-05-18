package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//To help us make a clean exit
	defer os.Exit(0)

	//Initialize the RPEngine
	var engine RPEngine
	engine.Init()

	//If there are command line args, process them
	if len(os.Args[1:]) > 0 {
		argStr := strings.Join(os.Args[1:], " ")
		fmt.Print(engine.EvalString(argStr))

	} else { //Otherwise run the REPL
		engine.InitREPL()
		engine.RunREPL()
	}

}
