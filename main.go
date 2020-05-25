package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//const rcfile = "~/.gorpnrc"
const rcfile = "rcfile"

func main() {
	//To help us make a clean exit
	defer os.Exit(0)

	//Initialize the RPEngine
	var engine RPEngine
	engine.Init()

	//If it exists, run rcfile upon startup.
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
		if strings.ToLower(os.Args[1]) == "help" {
			cliHelp()
		} else {
			argStr := strings.Join(os.Args[1:], " ")
			fmt.Print(engine.EvalString(argStr))
		}
	} else { //Otherwise run the REPL
		engine.InitREPL()
		engine.RunREPL()
	}

}

func cliHelp() {
	helpStr := `
goRPN
=====
goRPN is a console based RPN (reverse polish notation) calculator. Traditional
calculators (other than those from Hewlett Packard) use infix notation, where 
the mathematical operator comes between its two operands (3 + 2). In contrast,
RPN calculators use postfix notation, where the operator follows its two
operands. For example, to add 3 and 2, the sequence of commands looks like this:
3 2 +

Features
========
goRPN has all of the features you would expect of a scientific/programmer's 
calculator (trig functions, multiple memories, multiple display modes, macros,
etc). goRPN also implements arbitrary precision arithmetic (where feasible) for 
both floating point and integer values, which means you can do tricks like 
calculate the factorial of 10,000 (try it with: ./gorpn 10000 fact).

Operating modes
===============
goRPN will operate in two modes, CLI and interactive. In CLI mode it can be used
like a standard unix command line tool, suitable for scripting or general use. 
For example:
$ ./gorpn 3 2 + 
will result in \"5\" being printed to the console.

If goRPN is run with no commands:
$ ./gorpn
it will enter interactive mode. In this mode, it provides tab completion of 
commands, command history, online help, and live display of variables
and stack contents, among other features.

RC file
=======
Upon startup in either mode, goRPN will process the contents of the file:
~/.gorpnrc
If it exists, each line will be run as a separate command sequence. This is
useful for saving frequently used macros.

`
	fmt.Println(helpStr)
}
