package main

import (
	"os"
)

//var LivePrefixState struct {
//	LivePrefix string
//	IsEnable   bool
//}
//
//func executor(in string) {
//	fmt.Println("Your input: " + in)
//	if in == "" {
//		LivePrefixState.IsEnable = false
//		LivePrefixState.LivePrefix = in
//		return
//	}
//	LivePrefixState.LivePrefix = in + "> "
//	LivePrefixState.IsEnable = true
//}
//
//func completer(d prompt.Document) []prompt.Suggest {
//	s := []prompt.Suggest{
//		{Text: "+", Description: "Addition"},
//		{Text: "-", Description: "Subtraction"},
//		{Text: "*", Description: "Multiplication"},
//		{Text: "/", Description: "Division"},
//		{Text: "cla", Description: "Clears stack and variables"},
//		{Text: "exit", Description: "Exits from the calculator"},
//	}
//	//return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
//	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
//}
//
//func changeLivePrefix() (string, bool) {
//	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
//}

func main() {
	//To help us make a clean exit
	defer os.Exit(0)

	//Initialize the RPEngine
	var engine RPEngine
	engine.Init()
	engine.InitREPL()
	engine.RunREPL()
	//
	////Get some sample input
	//input := "3 2 + 2 * 5 - 5 /"
	//
	////Sanitize the input
	//input = strings.ToLower(strings.TrimSpace(input))
	//tokens := strings.Split(input, " ")
	//fmt.Printf("Result: %f\n", engine.Eval(tokens))

	//fmt.Println(tokens)

	//p := prompt.New(
	//	engine.replExecutor,
	//	completer,
	//	prompt.OptionPrefix(">>> "),
	//	prompt.OptionLivePrefix(changeReplPrefix),
	//	prompt.OptionTitle("live-prefix-example"),
	//)
	//p.Run()
}
