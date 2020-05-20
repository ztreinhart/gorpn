package main

import (
	"fmt"
	"math/big"
)

//Arithmetic Operators
func (r *RPEngine) plus() {
	raw1 := r.stack.Pop()
	raw2 := r.stack.Pop()
	switch op1 := raw1.(type) {
	case *big.Int:
		switch op2 := raw2.(type) {
		case *big.Int:
			op1 = op1.Add(op1, op2)
			r.stack.Push(op1)
		case *big.Float:
			flop1 := new(big.Float).SetInt(op1)
			flop1 = flop1.Add(flop1, op2)
			r.stack.Push(flop1)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch op2 := raw2.(type) {
		case *big.Int:
			flop2 := new(big.Float).SetInt(op2)
			flop2 = flop2.Add(flop2, op1)
			r.stack.Push(flop2)
		case *big.Float:
			op1 = op1.Add(op1, op2)
			r.stack.Push(op1)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch op2 := raw2.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			r.stack.Push(op1 || op2)
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) minus() {
	raw1 := r.stack.Pop()
	raw2 := r.stack.Pop()
	switch op1 := raw1.(type) {
	case *big.Int:
		switch op2 := raw2.(type) {
		case *big.Int:
			op2 = op2.Sub(op2, op1)
			r.stack.Push(op2)
		case *big.Float:
			flop1 := new(big.Float).SetInt(op1)
			op2 = op2.Sub(op2, flop1)
			r.stack.Push(op2)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch op2 := raw2.(type) {
		case *big.Int:
			flop2 := new(big.Float).SetInt(op2)
			flop2 = flop2.Sub(flop2, op1)
			r.stack.Push(flop2)
		case *big.Float:
			op2 = op2.Sub(op2, op1)
			r.stack.Push(op2)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch raw2.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			fmt.Println("Operation undefined between booleans")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) multiply() {
	raw1 := r.stack.Pop()
	raw2 := r.stack.Pop()
	switch op1 := raw1.(type) {
	case *big.Int:
		switch op2 := raw2.(type) {
		case *big.Int:
			op1 = op1.Mul(op1, op2)
			r.stack.Push(op1)
		case *big.Float:
			flop1 := new(big.Float).SetInt(op1)
			flop1 = flop1.Mul(flop1, op2)
			r.stack.Push(flop1)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch op2 := raw2.(type) {
		case *big.Int:
			flop2 := new(big.Float).SetInt(op2)
			flop2 = flop2.Mul(flop2, op1)
			r.stack.Push(flop2)
		case *big.Float:
			op1 = op1.Mul(op1, op2)
			r.stack.Push(op1)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch op2 := raw2.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			r.stack.Push(op1 && op2)
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) divide() {
	raw1 := r.stack.Pop()
	raw2 := r.stack.Pop()
	switch op1 := raw1.(type) {
	case *big.Int:
		switch op2 := raw2.(type) {
		case *big.Int:
			op2 = op2.Div(op2, op1)
			r.stack.Push(op2)
		case *big.Float:
			flop1 := new(big.Float).SetInt(op1)
			op2 = op2.Quo(op2, flop1)
			r.stack.Push(op2)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch op2 := raw2.(type) {
		case *big.Int:
			flop2 := new(big.Float).SetInt(op2)
			flop2 = flop2.Quo(flop2, op1)
			r.stack.Push(flop2)
		case *big.Float:
			op2 = op2.Quo(op2, op1)
			r.stack.Push(op2)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch raw2.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			fmt.Println("Operation undefined between booleans")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) cla() {
	r.stack.Init()
	r.vars = make(map[string]interface{})
}

func (r *RPEngine) clr() {
	r.stack.Init()
}

func (r *RPEngine) clv() {
	r.vars = make(map[string]interface{})

}

func (r *RPEngine) mod() {
	raw1 := r.stack.Pop() //y
	raw2 := r.stack.Pop() //x
	switch op1 := raw1.(type) {
	case *big.Int:
		switch op2 := raw2.(type) {
		case *big.Int:
			if op1 == big.NewInt(0) {
				fmt.Println("Undefined: Division by zero!")
				return
			}
			op2 = op2.Mod(op2, op1)
		case *big.Float:
			//TODO
			intop2, _ := op2.Int(nil)

		case bool:
			//TODO
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch op2 := raw2.(type) {
		case *big.Int:
			//TODO
		case *big.Float:
			//TODO
		case bool:
			//TODO
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch op2 := raw2.(type) {
		case *big.Int:
			//TODO
		case *big.Float:
			//TODO
		case bool:
			//TODO
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) increment() {
	inc := big.NewInt(1)
	r.stack.Push(inc)
	r.plus()
}

func (r *RPEngine) decrement() {
	dec := big.NewInt(1)
	r.stack.Push(dec)
	r.minus()
}

//Bitwise Operators
func (r *RPEngine) bitwiseAND() {
	//TODO: Bitwise AND

}

func (r *RPEngine) bitwiseOR() {
	//TODO: Bitwise OR

}

func (r *RPEngine) bitwiseXOR() {
	//TODO: Bitwise XOR

}

func (r *RPEngine) bitwiseNOT() {
	//TODO: Bitwise NOT

}

func (r *RPEngine) bitwiseLeftShift() {
	//TODO: Bitwise shift left

}

func (r *RPEngine) bitwiseRightShift() {
	//TODO: Bitwise shift right

}

//Boolean Operators
func (r *RPEngine) boolNOT() {
	//TODO: Boolean NOT
	arg1 := r.stack.Pop()
	arg2 := r.stack.Pop()
	r.stack.Push(float64(arg1 != arg2))
}

func (r *RPEngine) boolAND() {
	//TODO: Boolean AND
}

func (r *RPEngine) boolOR() {
	//TODO: Boolean OR
}

func (r *RPEngine) boolXOR() {
	//TODO: Boolean XOR
}

//Comparison Operators
func (r *RPEngine) notEqual() {
	//TODO: Not equal to

}

func (r *RPEngine) lessThan() {
	//TODO: Less than

}

func (r *RPEngine) lessThanEqualTo() {
	//TODO: Less than or equal to

}

func (r *RPEngine) equalTo() {
	//TODO: Equal to

}

func (r *RPEngine) greaterThan() {
	//TODO: Greater than

}

func (r *RPEngine) greaterThanEqualTo() {
	//TODO: Greater than or equal to

}

//Trig Functions
func (r *RPEngine) acos() {
	//TODO: Arc Cosine

}

func (r *RPEngine) asin() {
	//TODO: Arc Sine

}

func (r *RPEngine) atan() {
	//TODO: Arc Tangent

}

func (r *RPEngine) cos() {
	//TODO: Cosine

}

func (r *RPEngine) cosh() {
	//TODO: Hyberbolic Cosine

}

func (r *RPEngine) sin() {
	//TODO: Sine

}

func (r *RPEngine) sinh() {
	//TODO: Hyberbolic Cosine

}

func (r *RPEngine) tanh() {
	//TODO: Hyperbolic Tangent

}

//Numeric Utilities
func (r *RPEngine) ceiling() {
	//TODO: Ceiling

}

func (r *RPEngine) floor() {
	//TODO: Floor

}

func (r *RPEngine) round() {
	//TODO: Round

}

func (r *RPEngine) integerPart() {
	//TODO: Integer part

}

func (r *RPEngine) floatingPart() {
	//TODO: Floating part

}

func (r *RPEngine) sign() {
	//TODO: Push -1, 0, 1 depending on the sign of the operand

}

func (r *RPEngine) abs() {
	//TODO: Absolute value

}

func (r *RPEngine) max() {
	//TODO: Max

}

func (r *RPEngine) min() {
	//TODO: Min

}

//Display Modes
func (r *RPEngine) stackDisplay() {
	//TODO: Toggles the stack display from horizontal to vertical

}

func (r *RPEngine) hexDisplay() {
	//TODO: Switch display mode to hex

}

func (r *RPEngine) decDisplay() {
	//TODO: Switch display mode to decimal (default)

}

func (r *RPEngine) binDisplay() {
	//TODO: Switch display mode to binary

}

func (r *RPEngine) octDisplay() {
	//TODO: Switch display mode to octal

}

//Constants
func (r *RPEngine) constE() {
	//TODO: Push e onto the stack

}

func (r *RPEngine) constPi() {
	//TODO: Push pi onto the stack

}

func (r *RPEngine) random() {
	//TODO: Push a random integer onto the stack

}

//Mathematic functions
func (r *RPEngine) exp() {
	//TODO: Exponentiation

}

func (r *RPEngine) factorial() {
	//TODO: Factorial

}

func (r *RPEngine) sqrt() {
	//TODO: Square root

}

func (r *RPEngine) naturalLog() {
	//TODO: Natural log

}

func (r *RPEngine) log() {
	//TODO: Logarithm (base 2 or base 10?)

}

func (r *RPEngine) pow() {
	//TODO: Raise a number to a power

}

//Networking
func (r *RPEngine) hnl() {
	//TODO: Host to network long

}

func (r *RPEngine) hns() {
	//TODO: Host to network short

}

func (r *RPEngine) nhl() {
	//TODO: Network to host long

}

func (r *RPEngine) nhs() {
	//TODO: Network to host short

}

//Stack Manipulation
func (r *RPEngine) pick() {
	n := int(r.stack.Pop())
	r.stack.Push(r.stack.Pick(n))
}

func (r *RPEngine) depth() {
	r.stack.Push(float64(r.stack.Depth()))
}

func (r *RPEngine) drop() {
	r.stack.Drop()
}

func (r *RPEngine) dropn() {
	n := int(r.stack.Pop())
	r.stack.Dropn(n)
}

func (r *RPEngine) dup() {
	r.stack.Dup()
}

func (r *RPEngine) dupn() {
	n := int(r.stack.Pop())
	r.stack.Dupn(n)
}

func (r *RPEngine) roll() {
	n := int(r.stack.Pop())
	r.stack.Roll(n)
}

func (r *RPEngine) rolld() {
	n := int(r.stack.Pop())
	r.stack.Rolld(n)
}

func (r *RPEngine) swap() {
	r.stack.Swap()
}

//Macros and variables
func (r *RPEngine) repeat(op string) {
	n := int(r.stack.Pop())
	for i := 0; i < n; i++ {
		r.Eval([]string{op})
	}
}

func (r *RPEngine) macro(ops []string) {

}
