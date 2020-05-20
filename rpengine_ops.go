package main

import (
	"fmt"
	"math/big"
)

//Arithmetic Operators
func (r *RPEngine) plus() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			y = y.Add(y, x)
			r.stack.Push(y)
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			fly = fly.Add(fly, x)
			r.stack.Push(fly)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Add(flx, y)
			r.stack.Push(flx)
		case *big.Float:
			y = y.Add(y, x)
			r.stack.Push(y)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch x := rawX.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			r.stack.Push(y || x)
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) minus() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			x = x.Sub(x, y)
			r.stack.Push(x)
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			x = x.Sub(x, fly)
			r.stack.Push(x)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Sub(flx, y)
			r.stack.Push(flx)
		case *big.Float:
			x = x.Sub(x, y)
			r.stack.Push(x)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch rawX.(type) {
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
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			y = y.Mul(y, x)
			r.stack.Push(y)
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			fly = fly.Mul(fly, x)
			r.stack.Push(fly)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Mul(flx, y)
			r.stack.Push(flx)
		case *big.Float:
			y = y.Mul(y, x)
			r.stack.Push(y)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch x := rawX.(type) {
		case *big.Int:
			fmt.Println("Operation undefined between boolean and integer")
		case *big.Float:
			fmt.Println("Operation undefined between boolean and float")
		case bool:
			r.stack.Push(y && x)
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case nil:
		fmt.Println("Operation undefined on nils")
	}
}

func (r *RPEngine) divide() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			x = x.Div(x, y)
			r.stack.Push(x)
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			x = x.Quo(x, fly)
			r.stack.Push(x)
		case bool:
			fmt.Println("Operation undefined between integer and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Quo(flx, y)
			r.stack.Push(flx)
		case *big.Float:
			x = x.Quo(x, y)
			r.stack.Push(x)
		case bool:
			fmt.Println("Operation undefined between float and boolean.")
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case bool:
		switch rawX.(type) {
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
	rawY := r.stack.Pop() //y
	rawX := r.stack.Pop() //x
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			if y == big.NewInt(0) {
				fmt.Println("Undefined: Division by zero!")
				return
			}
			x = x.Mod(x, y)
		case *big.Float:
			//TODO
			intx, _ := x.Int(nil)

		case bool:
			//TODO
		case nil:
			fmt.Println("Operation undefined on nils")
		}
	case *big.Float:
		switch x := rawX.(type) {
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
		switch x := rawX.(type) {
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
