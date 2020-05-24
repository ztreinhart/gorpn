package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//const rcfile = "~/.rpnrc"
const rcfile = "rcfile"

func main() {
	//To help us make a clean exit
	defer os.Exit(0)

	//Initialize the RPEngine
	var engine RPEngine
	engine.Init()

	if _, err := os.Stat(rcfile); err == nil {
		if rawfile, err := ioutil.ReadFile(rcfile); err == nil {
			lines := strings.Split(string(rawfile), "\n")
			for _, line := range lines {
				_ = engine.EvalString(line)
			}
		}
	}

	//If there are command line args, process them
	if len(os.Args[1:]) > 0 {
		argStr := strings.Join(os.Args[1:], " ")
		fmt.Print(engine.EvalString(argStr))

	} else { //Otherwise run the REPL
		engine.InitREPL()
		engine.RunREPL()
	}

}
