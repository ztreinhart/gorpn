package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"unsafe"
)

//TODO: DO NOT REUSE big.Int AND big.Float pointers!

//Helpers
func copyHelper(rawX interface{}) interface{} {
	switch x := rawX.(type) {
	case *big.Int:
		return new(big.Int).Set(x)
	case *big.Float:
		return new(big.Float).Set(x)
	default:
		return rawX
	}
}

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

//TODO: func (r *RPEngine) popFloat64Helper() (float64, error) {}

func (r *RPEngine) popInt64Helper() (int64, error) {
	rawN := r.stack.Pop()
	switch n := rawN.(type) {
	case *big.Int:
		return n.Int64(), nil
	case *big.Float:
		int64n, _ := n.Int64()
		fmt.Println("Warning: implicit conversion of float to integer")
		return int64n, nil
	default:
		fmt.Printf("Operands of type %T cannot be converted to integers.\n", rawN)
		r.stack.Push(rawN)
		return 0, errors.New("Wrong type.")
	}
}

func (r *RPEngine) popIntHelper() (int, error) {
	i, err := r.popInt64Helper()
	return int(i), err
}

func (r *RPEngine) popUintHelper() (uint, error) {
	i, err := r.popInt64Helper()
	return uint(i), err
}

func (r *RPEngine) popUint16Helper() (uint16, error) {
	i, err := r.popInt64Helper()
	return uint16(i), err
}

func (r *RPEngine) popUint32Helper() (uint32, error) {
	i, err := r.popInt64Helper()
	return uint32(i), err
}

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

func (r *RPEngine) popBigIntHelper() (*big.Int, error) {
	rawX := r.stack.Pop()
	switch x := rawX.(type) {
	case *big.Int:
		return x, nil
	case *big.Float:
		fmt.Println("Implicit conversion to integer.")
		intx, _ := x.Int(nil)
		return intx, nil
	default:
		fmt.Printf("Operation undefined with operand of type %T\n", rawX)
		return nil, errors.New("Unsupported type")
	}
}

func (r *RPEngine) pushFloatOrInt(z *big.Float) {
	if z.IsInt() {
		zint, _ := z.Int(nil)
		r.stack.Push(zint)
	} else {
		r.stack.Push(z)
	}
}

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

func (r *RPEngine) multiply() {
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

func (r *RPEngine) divide() { //TODO: Add checks for division by zero
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
			_, rem := new(big.Int).QuoRem(x, y, new(big.Int))
			r.stack.Push(rem)
			return
		case *big.Float: //Y is an integer and X is a float
			fmt.Println("Warning: implicit conversion from float to integer")
			intx, _ := x.Int(nil)
			_, rem := new(big.Int).QuoRem(intx, y, new(big.Int))
			r.stack.Push(rem)
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
			_, rem := new(big.Int).QuoRem(x, inty, new(big.Int))
			r.stack.Push(rem)
			return
		case *big.Float: //Y is a float and X is a float
			intx, _ := x.Int(nil)
			//intx = intx.Mod(intx, inty)
			//r.stack.Push(intx)
			_, rem := new(big.Int).QuoRem(intx, inty, new(big.Int))
			r.stack.Push(rem)
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
	y, erry := r.popBigIntHelper()
	x, errx := r.popBigIntHelper()
	if errx == nil && erry == nil {
		r.stack.Push(new(big.Int).And(x, y))
	}

}

func (r *RPEngine) bitwiseOR() {
	y, erry := r.popBigIntHelper()
	x, errx := r.popBigIntHelper()
	if errx == nil && erry == nil {
		r.stack.Push(new(big.Int).Or(x, y))
	}
}

func (r *RPEngine) bitwiseXOR() {
	y, erry := r.popBigIntHelper()
	x, errx := r.popBigIntHelper()
	if errx == nil && erry == nil {
		r.stack.Push(new(big.Int).Xor(x, y))
	}
}

func (r *RPEngine) bitwiseNOT() {
	x, errx := r.popBigIntHelper()
	if errx == nil {
		r.stack.Push(new(big.Int).Not(x))
	}
}

func (r *RPEngine) bitwiseLeftShift() {
	n, errn := r.popIntHelper()
	x, errx := r.popBigIntHelper()
	if errn == nil && errx == nil {
		r.stack.Push(new(big.Int).Lsh(x, uint(n)))
	}
}

func (r *RPEngine) bitwiseRightShift() {
	n, errn := r.popIntHelper()
	x, errx := r.popBigIntHelper()
	if errn == nil && errx == nil {
		r.stack.Push(new(big.Int).Rsh(x, uint(n)))
	}
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

func (r *RPEngine) tan() {
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
	flx = math.Tan(flx)
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
	rawX := r.stack.Pop()
	switch x := rawX.(type) {
	case *big.Int:
		r.stack.Push(big.NewInt(int64(x.Sign())))
	case *big.Float:
		r.stack.Push(big.NewInt(int64(x.Sign())))
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
		r.stack.Push(rawX)
	}
}

//Native math/big.
func (r *RPEngine) abs() {
	rawX := r.stack.Pop()
	switch x := rawX.(type) {
	case *big.Int:
		r.stack.Push(x.Abs(x))
	case *big.Float:
		r.stack.Push(x.Abs(x))
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
		r.stack.Push(rawX)
	}
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
func (r *RPEngine) hexDisplay() {
	r.displayBase = 16

}

func (r *RPEngine) decDisplay() {
	r.displayBase = 10

}

func (r *RPEngine) binDisplay() {
	r.displayBase = 2

}

func (r *RPEngine) octDisplay() {
	r.displayBase = 8

}

func (r *RPEngine) stackDisplay() {
	switch r.stackDisp {
	case "vert":
		r.stackDisp = "horiz"
	case "horiz":
		r.stackDisp = "vert"
	}

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
	x, err := r.popInt64Helper()
	if err == nil {
		r.stack.Push(new(big.Int).MulRange(1, x))
	}
}

//Native implementation in math/big. Arbitrary precision
func (r *RPEngine) sqrt() {
	rawX := r.stack.Pop()
	switch x := rawX.(type) {
	case *big.Int:
		r.stack.Push(x.Sqrt(x))
	case *big.Float:
		r.stack.Push(x.Sqrt(x))
	default:
		fmt.Printf("Operation undefined on type: %T\n ", rawX)
		r.stack.Push(rawX)
	}
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
//Takes value from top of Stack as little-endian uint32,
//pushes a big-endian uint32 back on the Stack
func (r *RPEngine) hnl() {
	rawX := r.stack.Peek()
	x, err := r.popUint32Helper()
	if err == nil {
		buf := (*[4]byte)(unsafe.Pointer(&x))[:]
		nl := binary.BigEndian.Uint32(buf)
		r.stack.Push(big.NewInt(int64(nl)))
		return
	}
	r.stack.Push(rawX)
}

//Takes value from top of Stack as little-endian uint16,
//pushes a big-endian uint16 back on the Stack
func (r *RPEngine) hns() {
	rawX := r.stack.Peek()
	x, err := r.popUint16Helper()
	if err == nil {
		buf := (*[2]byte)(unsafe.Pointer(&x))[:]
		ns := binary.BigEndian.Uint16(buf)
		r.stack.Push(big.NewInt(int64(ns)))
		return
	}
	r.stack.Push(rawX)
}

//Takes value from top of Stack as big-endian uint32,
//pushes a little-endian uint32 back on the Stack
func (r *RPEngine) nhl() {
	rawX := r.stack.Peek()
	x, err := r.popUint32Helper()
	if err == nil {
		buf := (*[4]byte)(unsafe.Pointer(&x))[:]
		hl := binary.LittleEndian.Uint32(buf)
		r.stack.Push(big.NewInt(int64(hl)))
		return
	}
	r.stack.Push(rawX)
}

//Takes value from top of Stack as big-endian uint16,
//pushes a little-endian uint16 back on the Stack
func (r *RPEngine) nhs() {
	rawX := r.stack.Peek()
	x, err := r.popUint16Helper()
	if err == nil {
		buf := (*[2]byte)(unsafe.Pointer(&x))[:]
		hs := binary.BigEndian.Uint16(buf)
		r.stack.Push(big.NewInt(int64(hs)))
		return
	}
	r.stack.Push(rawX)
}

//Stack Manipulation
//TODO: Doesn't work
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
	//TODO: doesn't work
	n, err := r.popIntHelper()
	if err == nil {
		r.stack.Dropn(n)
	}
}

func (r *RPEngine) dup() {
	//r.stack.Dup()
	rawX := r.stack.Peek()

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

func (r *RPEngine) setMacro(ops []string) {
	r.macros[ops[1]] = strings.Join(ops[2:], " ")

}

func (r *RPEngine) listMacros() {
	fmt.Println("Currently defined macros: ")
	for key, value := range r.macros {
		fmt.Println(key, ":\t", value)
	}
}

func (r *RPEngine) rmMacro(macro string) {
	delete(r.macros, macro)
}

func (r *RPEngine) clearMacros() {
	r.macros = make(map[string]string)
}

func (r *RPEngine) runMacro(key string) {
	r.Eval(strings.Split(r.macros[key], " "))
}

func (r *RPEngine) setVar(token string) {
	varName := string([]rune(token)[0])
	r.vars[varName] = r.stack.Pop()
}

func (r *RPEngine) getVar(key string) {
	r.stack.Push(r.vars[key])
}
