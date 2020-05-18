package main

import (
	"os"
)

func main() {
	//To help us make a clean exit
	defer os.Exit(0)

	//Initialize the RPEngine
	var engine RPEngine
	engine.Init()

	//If we're going into interactive mode
	engine.InitREPL()
	engine.RunREPL()

}
