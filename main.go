package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const rcfile = "~/.gorpnrc"

//const rcfile = "rcfile"

const helpStr = `
goRPN
===============
goRPN is a console based RPN (reverse polish notation) calculator.

USAGE
===============
	$ gorpn						interactive mode
	$ gorpn [rpn expression]	evaluate a single expression

EXAMPLES
===============
	$ gorpn 3 3 + 					=>	6
	$ gorpn '30 400 * 15 60 * +' 	=>  12900 
			(quotes prevent shell from expanding * into file list)

	$ gorpn e ln 					=>  1
	$ gorpn 						=>  interactive mode

RC file
===============
Upon startup in either mode, goRPN will process the contents of the file:
~/.gorpnrc
If it exists, each line will be run as a separate command sequence. This is
useful for saving frequently used macros.

MORE HELP
===============
An in-depth command reference is available by executing the 'help' command
in interactive mode.
`

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

	//Check on the standard input
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe != 0 { //If we have a pipe input
		reader := bufio.NewReader(os.Stdin)
		var contents []rune
		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			contents = append(contents, input)
		}

		contentStr := string(contents)
		fmt.Println(engine.EvalString(contentStr))

	} else if len(os.Args[1:]) > 0 { //If we aren't on a pipe, but we have command line arguments
		if strings.ToLower(os.Args[1]) == "help" {
			cliHelp()
		} else {
			argStr := strings.Join(os.Args[1:], " ")
			fmt.Println(engine.EvalString(argStr))
		}

	} else { //Otherwise run the REPL
		engine.InitREPL()
		engine.RunREPL()
	}

}

func cliHelp() {
	fmt.Println(helpStr)
}
