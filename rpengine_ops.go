package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
)

//Arithmetic Operators
func (r *RPEngine) plus() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int: //Y is an integer
		switch x := rawX.(type) {
		case *big.Int: //Y is an integer, X is an integer
			y = y.Add(y, x)
			r.stack.Push(y)
			return
		case *big.Float: //Y is an integer, X is a float
			fly := new(big.Float).SetInt(y)
			fly = fly.Add(fly, x)
			r.stack.Push(fly)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float: //Y is a float
		switch x := rawX.(type) {
		case *big.Int: //Y is a float, X is an integer
			flx := new(big.Float).SetInt(x)
			flx = flx.Add(flx, y)
			r.stack.Push(flx)
			return
		case *big.Float: //Y is a float, X is a float
			y = y.Add(y, x)
			r.stack.Push(y)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case bool:
		switch x := rawX.(type) {
		case bool:
			r.stack.Push(y || x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
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
			return
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			x = x.Sub(x, fly)
			r.stack.Push(x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Sub(flx, y)
			r.stack.Push(flx)
			return
		case *big.Float:
			x = x.Sub(x, y)
			r.stack.Push(x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
}

func (r *RPEngine) multiply() { //TODO: Add checks for division by zero
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case *big.Int:
		switch x := rawX.(type) {
		case *big.Int:
			y = y.Mul(y, x)
			r.stack.Push(y)
			return
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			fly = fly.Mul(fly, x)
			r.stack.Push(fly)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Mul(flx, y)
			r.stack.Push(flx)
			return
		case *big.Float:
			y = y.Mul(y, x)
			r.stack.Push(y)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)

		}
	case bool:
		switch x := rawX.(type) {
		case bool:
			r.stack.Push(y && x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
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
			return
		case *big.Float:
			fly := new(big.Float).SetInt(y)
			x = x.Quo(x, fly)
			r.stack.Push(x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float:
		switch x := rawX.(type) {
		case *big.Int:
			flx := new(big.Float).SetInt(x)
			flx = flx.Quo(flx, y)
			r.stack.Push(flx)
			return
		case *big.Float:
			x = x.Quo(x, y)
			r.stack.Push(x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
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
	case *big.Int: //Y is an integer
		if y == big.NewInt(0) {
			fmt.Println("Undefined: Division by zero!")
			break
		}

		switch x := rawX.(type) {
		case *big.Int: //Y is an integer and X is an integer
			x = x.Mod(x, y)
			r.stack.Push(x)
			return
		case *big.Float: //Y is an integer and X is a float
			fmt.Println("Warning: implicit conversion from float to integer")
			intx, _ := x.Int(nil)
			intx = intx.Mod(intx, y)
			r.stack.Push(intx)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float: //Y is a float
		fmt.Println("Warning: implicit conversion from float to integer")
		inty, _ := y.Int(nil)

		if inty == big.NewInt(0) {
			fmt.Println("Undefined: Division by zero!")
			break
		}

		switch x := rawX.(type) {
		case *big.Int: //Y is a float and X is an integer
			x = x.Mod(x, inty)
			r.stack.Push(x)
			return
		case *big.Float: //Y is a float and X is a float
			intx, _ := x.Int(nil)
			intx = intx.Mod(intx, inty)
			r.stack.Push(intx)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
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
	rawX := r.stack.Pop()
	switch x := rawX.(type) {
	case bool:
		r.stack.Push(!x)
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
		r.stack.Push(rawX)
	}
}

func (r *RPEngine) boolAND() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case bool:
		switch x := rawX.(type) {
		case bool:
			r.stack.Push(x && y)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
}

func (r *RPEngine) boolOR() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case bool:
		switch x := rawX.(type) {
		case bool:
			r.stack.Push(x || y)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
}

func (r *RPEngine) boolXOR() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()
	switch y := rawY.(type) {
	case bool:
		switch x := rawX.(type) {
		case bool:
			r.stack.Push(x != y)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	default:
		fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
	}
	r.stack.Push(rawX)
	r.stack.Push(rawY)
}

//Comparison Operators
func compareHelper(rawX, rawY interface{}) (int, error) {
	switch y := rawY.(type) {
	case *big.Int: //Y is an integer
		switch x := rawX.(type) {
		case *big.Int: //Y is an integer, X is an integer
			sign := x.Cmp(y)
			return sign, nil
		case *big.Float: //Y is an integer, X is a float
			fly := new(big.Float).SetInt(y)
			sign := x.Cmp(fly)
			return sign, nil
		default:
			errstr := fmt.Sprintf("Operation undefined between %T and %T.", rawX, rawY)
			return 0, errors.New(errstr)
		}
	case *big.Float: //Y is a float
		switch x := rawX.(type) {
		case *big.Int: //Y is a float, X is an integer
			flx := new(big.Float).SetInt(x)
			sign := flx.Cmp(y)
			return sign, nil
		case *big.Float: //Y is a float, X is a float
			sign := x.Cmp(y)
			return sign, nil
		default:
			errstr := fmt.Sprintf("Operation undefined between %T and %T.", rawX, rawY)
			return 0, errors.New(errstr)
		}
	default:
		errstr := fmt.Sprintf("Operation undefined between %T and %T.", rawX, rawY)
		return 0, errors.New(errstr)
	}
}

func (r *RPEngine) notEqual() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		x, ok1 := rawX.(bool)
		y, ok2 := rawY.(bool)

		if ok1 && ok2 {
			r.stack.Push(x != y)
		} else {
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
			r.stack.Push(rawX)
			r.stack.Push(rawY)
		}
		return
	}

	r.stack.Push(sign != 0)
}

func (r *RPEngine) lessThan() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(sign == -1)

}

func (r *RPEngine) lessThanEqualTo() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(sign == -1 || sign == 0)
}

func (r *RPEngine) equalTo() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		x, ok1 := rawX.(bool)
		y, ok2 := rawY.(bool)

		if ok1 && ok2 {
			r.stack.Push(x == y)
		} else {
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
			r.stack.Push(rawX)
			r.stack.Push(rawY)
		}
		return
	}

	r.stack.Push(sign == 0)
}

func (r *RPEngine) greaterThan() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(sign == 1)

}

func (r *RPEngine) greaterThanEqualTo() {
	rawY := r.stack.Pop()
	rawX := r.stack.Pop()

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(sign == 0 || sign == 1)

}

//Trig Functions
func (r *RPEngine) acos() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for arc cosine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Acos(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) asin() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for arc sine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Asin(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) atan() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for arc tangent calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Atan(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) cos() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for cosine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Cos(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) cosh() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for hyberbolic cosine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Cosh(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) sin() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for sine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Sin(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) sinh() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for hyperbolic sine calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Sinh(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) tanh() {
	rawX := r.stack.Pop()
	var flx float64
	switch x := rawX.(type) {
	case *big.Int:
		var err error
		flx, err = strconv.ParseFloat(x.String(), 64)
		if err != nil {
			fmt.Println("Error parsing integer into float for hyperbolic tangent calculation.")
		}
	case *big.Float:
		flx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
	}
	flx = math.Tanh(flx)
	r.stack.Push(big.NewFloat(flx))
}

//Numeric Utilities
func (r *RPEngine) popFloatHelper() (float64, error) {
	rawX := r.stack.Pop()
	var floatx float64
	switch x := rawX.(type) {
	case *big.Int:
		floatx, _ = strconv.ParseFloat(x.String(), 64)
	case *big.Float:
		floatx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined with operand of type %T\n", rawX)
		return math.NaN(), errors.New("Unsupported type")
	}
	return floatx, nil
}

func (r *RPEngine) pushFloatOrInt(z *big.Float) {
	if z.IsInt() {
		zint, _ := z.Int(nil)
		r.stack.Push(zint)
	} else {
		r.stack.Push(z)
	}
}

//float64 precision
func (r *RPEngine) ceiling() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Ceil(x))
		r.pushFloatOrInt(z)
	}

}

//float64 precision
func (r *RPEngine) floor() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Floor(x))
		r.pushFloatOrInt(z)
	}
}

//float64 precision
func (r *RPEngine) round() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Round(x))
		r.pushFloatOrInt(z)
	}
}

//float64 precision
func (r *RPEngine) integerPart() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		zint, _ := math.Modf(x)
		r.stack.Push(big.NewInt(int64(zint)))
	}
}

//float64 precision
func (r *RPEngine) floatingPart() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		_, zfrac := math.Modf(x)
		r.stack.Push(big.NewFloat(zfrac))
	}
}

//Native math/big.
func (r *RPEngine) sign() {
	//TODO: Push -1, 0, 1 depending on the sign of the operand

}

//Native math/big.
func (r *RPEngine) abs() {
	//TODO: Absolute value

}

//float64 precision
func (r *RPEngine) max() {
	y, erry := r.popFloatHelper()
	x, errx := r.popFloatHelper()
	if errx == nil && erry == nil {
		z := big.NewFloat(math.Max(x, y))
		r.pushFloatOrInt(z)
	}
}

//float64 precision
func (r *RPEngine) min() {
	y, erry := r.popFloatHelper()
	x, errx := r.popFloatHelper()
	if errx == nil && erry == nil {
		z := big.NewFloat(math.Min(x, y))
		r.pushFloatOrInt(z)
	}
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
	r.stack.Push(big.NewFloat(math.E))
}

func (r *RPEngine) constPi() {
	r.stack.Push(big.NewFloat(math.Pi))
}

func (r *RPEngine) random() {
	r.stack.Push(big.NewInt(int64(rand.Int())))

}

//Mathematic functions
//Currently only float64 precision due to use of math.Exp()
func (r *RPEngine) exp() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Exp(x))
		r.pushFloatOrInt(z)
	}
}

//Native implementation in math/big. Arbitrary precision
func (r *RPEngine) factorial() {
	//TODO: Factorial

}

//Native implementation in math/big. Arbitrary precision
func (r *RPEngine) sqrt() {
	//TODO: Square root

}

//Currently only float64 precision due to use of math.Log()
func (r *RPEngine) naturalLog() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Log(x))
		r.pushFloatOrInt(z)
	}
}

//Currently only float64 precision due to use of math.Log10()
func (r *RPEngine) log() {
	x, errx := r.popFloatHelper()
	if errx == nil {
		z := big.NewFloat(math.Log10(x))
		r.pushFloatOrInt(z)
	}
}

//Currently only float64 precision due to use of math.Pow()
func (r *RPEngine) pow() {
	y, erry := r.popFloatHelper()
	x, errx := r.popFloatHelper()
	if errx == nil && erry == nil {
		z := big.NewFloat(math.Pow(x, y))
		r.pushFloatOrInt(z)
	}
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
func (r *RPEngine) popIntHelper() (int, error) {
	rawN := r.stack.Pop()
	switch n := rawN.(type) {
	case *big.Int:
		return int(n.Int64()), nil
	case *big.Float:
		int64n, _ := n.Int64()
		fmt.Println("Warning: using float value as integer index or counter.")
		return int(int64n), nil
	default:
		fmt.Printf("Operands of type %T cannot be used as integer stack indices or counters.\n", rawN)
		r.stack.Push(rawN)
		return -1, errors.New("Wrong type.")
	}
}

func (r *RPEngine) pick() {
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Push(r.stack.Pick(n))
	}
}

func (r *RPEngine) depth() {
	r.stack.Push(float64(r.stack.Depth()))
}

func (r *RPEngine) drop() {
	r.stack.Drop()
}

func (r *RPEngine) dropn() {
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Dropn(n)
	}
}

func (r *RPEngine) dup() {
	r.stack.Dup()
}

func (r *RPEngine) dupn() {
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Dupn(n)
	}
}

func (r *RPEngine) roll() {
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Roll(n)
	}
}

func (r *RPEngine) rolld() {
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Rolld(n)
	}
}

func (r *RPEngine) swap() {
	r.stack.Swap()
}

//Macros and variables
func (r *RPEngine) repeat(op string) {
	n, err := r.popIntHelper()
	if err == nil {
		for i := 0; i < n; i++ {
			r.Eval([]string{op})
		}
	}
}

func (r *RPEngine) macro(ops []string) {

}
