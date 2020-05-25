package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

const REPLHelpStr = `
 goRPN Command Reference
 ===============================
 goRPN is a full featured console RPN calculator. It provides arbitrary
 precision integer and floating point arithmetic where feasible. The following
 commands are supported, with max precision noted in parentheses where
 applicable.

 Arithmetic Operators
 ===============================

    +          Add        (arbitrary)
    -          Subtract   (arbitrary)
    *          Multiply   (arbitrary)
    /          Divide     (arbitrary)
    cla        Clear the stack and variables
    clr        Clear the stack
    clv        Clear the variables
    %          Modulus    (arbitrary)
    ++         Increment  (arbitrary)
    --         Decrement  (arbitrary)

 Bitwise Operators
 ===============================

    &          Bitwise AND
    |          Bitwise OR
    ^          Bitwise XOR
    ~          Bitwise NOT
    <<         Bitwise shift left
    >>         Bitwise shift right

 Boolean Operators
 ===============================

    !          Boolean NOT
    &&         Boolean AND
    ||         Boolean OR
    ^^         Boolean XOR

 Comparison Operators
 ===============================

    !=         Not equal to
    <          Less than
    <=         Less than or equal to
    ==         Equal to
    >          Greater than
    >=         Greater than or equal to

 Trigonometric Functions
 ===============================

    acos       Arc Cosine         (64-bit floating point)
    asin       Arc Sine           (64-bit floating point)
    atan       Arc Tangent        (64-bit floating point)
    cos        Cosine             (64-bit floating point)
    cosh       Hyperbolic Cosine  (64-bit floating point)
    sin        Sine               (64-bit floating point)
    sinh       Hyperbolic Sine    (64-bit floating point)
    tanh       Hyperbolic tangent (64-bit floating point)

 Numeric Utilities
 ===============================

    ceil       Ceiling            (64-bit floating point)
    floor      Floor              (64-bit floating point)
    round      Round              (64-bit floating point)
    ip         Integer part       (64-bit floating point)
    fp         Floating part      (64-bit floating point)
    sign       Push -1, 0, or 0 depending on the sign
    abs        Absolute value     (arbitrary)
    max        Max
    min        Min
 
 Display Modes
 ===============================

    hex        Switch display mode to hexadecimal
    dec        Switch display mode to decimal (default)
    bin        Switch display mode to binary
    oct        Switch display mode to octal

 Constants
 ===============================

    e          Push e                     (64-bit floating point)
    pi         Push Pi                    (64-bit floating point)
    rand       Generate a random number   (integer)

 Mathematic Functions
 ===============================

    exp        Exponentiation             (64-bit floating point)
    fact       Factorial                  (arbitrary)
    sqrt       Square Root                (arbitrary)
    ln         Natural Logarithm          (64-bit floating point)
    log        Logarithm                  (64-bit floating point)
    pow        Raise a number to a power  (64-bit floating point)

 Networking
 ===============================

    hnl        Host to network long       (32-bit unsigned integer)
    hns        Host to network short      (16-bit unsigned integer)
    nhl        Network to host long       (32-bit unsigned integer)
    nhs        Network to host short      (16-bit unsigned integer)

 Stack Manipulation
 ===============================

    pick       Pick the -n'th item from the stack
    repeat     Repeat an operation n times, e.g. '3 repeat +'
    depth      Push the current stack depth
    drop       Drops the top item from the stack
    dropn      Drops n items from the stack
    dup        Duplicates the top stack item
    dupn       Duplicates the top n stack items in order
    roll       Roll the stack upwards by n
    rolld      Roll the stack downwards by n
    stack      Toggles stack display from horizontal to vertical
    swap       Swap the top 2 stack items

 Macros and Variables
 ===============================

    macro      Defines a macro, e.g. 'macro kib 1024 *'
    lsmacros   List currently defined macros
    clm        Clear macros
    x=         Assigns a variable, e.g. '1024 x='

 Other
 ===============================

    help       Print the help message
	type	   Gets the data type of the element at top of stack
    exit       Exit the calculator
    quit       Exit the calculator
`

func (r *RPEngine) replCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "+", Description: "Add"},
		{Text: "-", Description: "Subtract"},
		{Text: "*", Description: "Multiply"},
		{Text: "/", Description: "Divide"},
		{Text: "cla", Description: "Clear the stack and variables"},
		{Text: "clr", Description: "Clear the stack"},
		{Text: "clv", Description: "Clear the variables"},
		{Text: "%", Description: "Modulus"},
		{Text: "++", Description: "Increment"},
		{Text: "--", Description: "Decrement"},
		{Text: "&", Description: "Bitwise AND"},
		{Text: "|", Description: "Bitwise OR"},
		{Text: "^", Description: "Bitwise XOR"},
		{Text: "~", Description: "Bitwise NOT"},
		{Text: "<<", Description: "Bitwise shift left"},
		{Text: ">>", Description: "Bitwise shift right"},
		{Text: "!", Description: "Boolean NOT"},
		{Text: "&&", Description: "Boolean AND"},
		{Text: "||", Description: "Boolean OR"},
		{Text: "^^", Description: "Boolean XOR"},
		{Text: "!=", Description: "Not equal to"},
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
		{Text: "rmmacro", Description: "Remove a defined macro, e.g. 'rmmacro kib'"},
		{Text: "clm", Description: "Clears all defined macros"},
		{Text: "x=", Description: "Assigns a variable, e.g. '1024 x='"},
		{Text: "type", Description: "Prints type of top variable in the stack. (useful for debugging)"},
		{Text: "help", Description: "Print the help message"},
		{Text: "exit", Description: "Exit the calculator"},
		{Text: "quit", Description: "Exit the calculator"},
	}
	//return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
}

func (r *RPEngine) replHelp() {
	fmt.Print(REPLHelpStr)
}
