package main

import "github.com/c-bata/go-prompt"

func (r *RPEngine) replCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "+", Description: "Add"},
		{Text: "-", Description: "Subtract"},
		{Text: "*", Description: "Multiply"},
		{Text: "/", Description: "Divide"},
		{Text: "cla", Description: "Clear the stack and variables"},
		{Text: "clr", Description: "Clear the stack"},
		{Text: "clv", Description: "Clear the variables"},
		{Text: "!", Description: "Boolean NOT"},
		{Text: "!=", Description: "Not equal to"},
		{Text: "%", Description: "Modulus"},
		{Text: "++", Description: "Increment"},
		{Text: "--", Description: "Decrement"},
		{Text: "&", Description: "Bitwise AND"},
		{Text: "|", Description: "Bitwise OR"},
		{Text: "^", Description: "Bitwise XOR"},
		{Text: "~", Description: "Bitwise NOT"},
		{Text: "<<", Description: "Bitwise shift left"},
		{Text: ">>", Description: "Bitwise shift right"},
		{Text: "&&", Description: "Boolean AND"},
		{Text: "||", Description: "Boolean OR"},
		{Text: "^^", Description: "Boolean XOR"},
		{Text: "<", Description: "Less than"},
		{Text: "<=", Description: "Less than or equal to"},
		{Text: "==", Description: "Equal to"},
		{Text: ">", Description: "Greater than"},
		{Text: ">=", Description: "Greater than or equal to"},
		{Text: "acos", Description: "Arc Cosine"},
		{Text: "asin", Description: "Arc Sine"},
		{Text: "atan", Description: "Arc Tangent"},
		{Text: "cos", Description: "Cosine"},
		{Text: "cosh", Description: "Hyperbolic Cosine"},
		{Text: "sin", Description: "Sine"},
		{Text: "sinh", Description: "Hyperbolic Sine"},
		{Text: "tanh", Description: "Hyperbolic tangent"},
		{Text: "ceil", Description: "Ceiling"},
		{Text: "floor", Description: "Floor"},
		{Text: "round", Description: "Round"},
		{Text: "ip", Description: "Integer part"},
		{Text: "fp", Description: "Floating part"},
		{Text: "sign", Description: "Push -1, 0, or 0 depending on the sign"},
		{Text: "abs", Description: "Absolute value"},
		{Text: "max", Description: "Max"},
		{Text: "min", Description: "Min"},
		{Text: "hex", Description: "Switch display mode to hexadecimal"},
		{Text: "dec", Description: "Switch display mode to decimal (default)"},
		{Text: "bin", Description: "Switch display mode to binary"},
		{Text: "oct", Description: "Switch display mode to octal"},
		{Text: "e", Description: "Push e"},
		{Text: "pi", Description: "Push Pi"},
		{Text: "rand", Description: "Generate a random number"},
		{Text: "exp", Description: "Exponentiation"},
		{Text: "fact", Description: "Factorial"},
		{Text: "sqrt", Description: "Square Root"},
		{Text: "ln", Description: "Natural Logarithm"},
		{Text: "log", Description: "Logarithm"},
		{Text: "pow", Description: "Raise a number to a power"},
		{Text: "hnl", Description: "Host to network long"},
		{Text: "hns", Description: "Host to network short"},
		{Text: "nhl", Description: "Network to host long"},
		{Text: "nhs", Description: "Network to host short"},
		{Text: "pick", Description: "Pick the -n'th item from the stack"},
		{Text: "repeat", Description: "Repeat an operation n times, e.g. '3 repeat +'"},
		{Text: "depth", Description: "Push the current stack depth"},
		{Text: "drop", Description: "Drops the top item from the stack"},
		{Text: "dropn", Description: "Drops n items from the stack"},
		{Text: "dup", Description: "Duplicates the top stack item"},
		{Text: "dupn", Description: "Duplicates the top n stack items in order"},
		{Text: "roll", Description: "Roll the stack upwards by n"},
		{Text: "rolld", Description: "Roll the stack downwards by n"},
		{Text: "stack", Description: "Toggles stack display from horizontal to vertical"},
		{Text: "swap", Description: "Swap the top 2 stack items"},
		{Text: "macro", Description: "Defines a macro, e.g. 'macro kib 1024 *'"},
		{Text: "lsmacros", Description: "Lists currently defined macros"},
		{Text: "x=", Description: "Assigns a variable, e.g. '1024 x='"},
		{Text: "type", Description: "Prints type of top variable in the stack. (useful for debugging)"},
		{Text: "help", Description: "Print the help message"},
		{Text: "exit", Description: "Exit the calculator"},
	}
	//return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
}
