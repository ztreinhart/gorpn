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

//Helpers
//returns -1 if x is less than y, +1 if x is greater than y, 0 if they're equal.
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
			errstr := fmt.Sprintf("operation undefined between %T and %T.", rawX, rawY)
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
			errstr := fmt.Sprintf("operation undefined between %T and %T", rawX, rawY)
			return 0, errors.New(errstr)
		}
	default:
		errstr := fmt.Sprintf("operation undefined between %T and %T", rawX, rawY)
		return 0, errors.New(errstr)
	}
}

func (r *RPEngine) popInt64Helper() (int64, error) {
	rawN, err := r.stack.Pop()
	if err != nil {
		return 0, err
	}
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
		return 0, errors.New("wrong type")
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
	rawX, err := r.stack.Pop()
	if err != nil {
		return math.NaN(), err
	}
	var floatx float64
	switch x := rawX.(type) {
	case *big.Int:
		floatx, _ = strconv.ParseFloat(x.String(), 64)
	case *big.Float:
		floatx, _ = x.Float64()
	default:
		fmt.Printf("Operation undefined with operand of type %T\n", rawX)
		r.stack.Push(rawX)
		return math.NaN(), errors.New("unsupported type")
	}
	return floatx, nil
}

func (r *RPEngine) popBigIntHelper() (*big.Int, error) {
	rawX, err := r.stack.Pop()
	if err != nil {
		return nil, err
	}
	switch x := rawX.(type) {
	case *big.Int:
		return x, nil
	case *big.Float:
		fmt.Println("Implicit conversion to integer.")
		intx, _ := x.Int(nil)
		return intx, nil
	default:
		fmt.Printf("Operation undefined with operand of type %T\n", rawX)
		r.stack.Push(rawX)
		return nil, errors.New("unsupported type")
	}
}

func (r *RPEngine) popBoolHelper() (bool, error) {
	rawX, err := r.stack.Pop()
	if err != nil {
		return false, err
	}
	x, ok := rawX.(bool)
	if !ok {
		fmt.Println("Operation undefined with operand of type %T\n")
		r.stack.Push(rawX)
		return false, errors.New("unsupported type")
	}
	return x, nil
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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

	switch y := rawY.(type) {
	case *big.Int: //Y is an int
		if y == big.NewInt(0) {
			fmt.Println("Undefined: Division by zero!")
			break
		}

		switch x := rawX.(type) {
		case *big.Int: //X is int, y is int
			fly := new(big.Float).SetInt(y)
			flx := new(big.Float).SetInt(x)
			flx = flx.Quo(flx, fly)
			r.stack.Push(flx)
			return
		case *big.Float: //X is float, y is int
			fly := new(big.Float).SetInt(y)
			x = x.Quo(x, fly)
			r.stack.Push(x)
			return
		default:
			fmt.Printf("Operation undefined between %T and %T.\n", rawX, rawY)
		}
	case *big.Float: //Y is a float
		if y == big.NewFloat(0) {
			fmt.Println("Undefined: Division by zero!")
			break
		}

		switch x := rawX.(type) {
		case *big.Int: //Y is float, X is int
			flx := new(big.Float).SetInt(x)
			flx = flx.Quo(flx, y)
			r.stack.Push(flx)
			return
		case *big.Float: //X and Y are both floats.
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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, erry := r.stack.Peek()
	y, erry := r.popBigIntHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBigIntHelper()
	if errx != nil {
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(new(big.Int).And(x, y))
}

func (r *RPEngine) bitwiseOR() {
	rawY, erry := r.stack.Peek()
	y, erry := r.popBigIntHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBigIntHelper()
	if errx != nil {
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(new(big.Int).Or(x, y))
}

func (r *RPEngine) bitwiseXOR() {
	rawY, erry := r.stack.Peek()
	y, erry := r.popBigIntHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBigIntHelper()
	if errx != nil {
		r.stack.Push(rawY)
		return
	}

	r.stack.Push(new(big.Int).Xor(x, y))
}

func (r *RPEngine) bitwiseNOT() {
	x, errx := r.popBigIntHelper()
	if errx != nil {
		return
	}

	r.stack.Push(new(big.Int).Not(x))
}

func (r *RPEngine) bitwiseLeftShift() {
	rawN, errn := r.stack.Peek()
	n, errn := r.popIntHelper()
	if errn != nil {
		return
	}

	x, errx := r.popBigIntHelper()
	if errx != nil {
		r.stack.Push(rawN)
		return
	}

	r.stack.Push(new(big.Int).Lsh(x, uint(n)))
}

func (r *RPEngine) bitwiseRightShift() {
	rawN, errn := r.stack.Peek()
	n, errn := r.popIntHelper()
	if errn != nil {
		return
	}

	x, errx := r.popBigIntHelper()
	if errx != nil {
		r.stack.Push(rawN)
		return
	}

	r.stack.Push(new(big.Int).Rsh(x, uint(n)))
}

//Boolean Operators
func (r *RPEngine) boolNOT() {
	x, err := r.popBoolHelper()
	if err != nil {
		return
	}

	r.stack.Push(!x)
}

func (r *RPEngine) boolAND() {
	y, erry := r.popBoolHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBoolHelper()
	if errx != nil {
		r.stack.Push(y)
		return
	}

	r.stack.Push(x && y)
}

func (r *RPEngine) boolOR() {
	y, erry := r.popBoolHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBoolHelper()
	if errx != nil {
		r.stack.Push(y)
		return
	}

	r.stack.Push(x || y)
}

func (r *RPEngine) boolXOR() {
	y, erry := r.popBoolHelper()
	if erry != nil {
		return
	}

	x, errx := r.popBoolHelper()
	if errx != nil {
		r.stack.Push(y)
		return
	}

	r.stack.Push(x != y)
}

//Comparison Operators
func (r *RPEngine) notEqual() {
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

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
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Acos(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) asin() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Asin(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) atan() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Atan(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) cos() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Cos(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) cosh() {

	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Cosh(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) sin() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Sin(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) sinh() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Sinh(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) tan() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Tan(flx)
	r.stack.Push(big.NewFloat(flx))
}

func (r *RPEngine) tanh() {
	flx, err := r.popFloatHelper()
	if err != nil {
		return
	}

	flx = math.Tanh(flx)
	r.stack.Push(big.NewFloat(flx))
}

//Numeric Utilities

//float64 precision
func (r *RPEngine) ceiling() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Ceil(x))
	r.pushFloatOrInt(z)
}

//float64 precision
func (r *RPEngine) floor() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Floor(x))
	r.pushFloatOrInt(z)
}

//float64 precision
func (r *RPEngine) round() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Round(x))
	r.pushFloatOrInt(z)
}

//float64 precision
func (r *RPEngine) integerPart() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	zint, _ := math.Modf(x)
	r.stack.Push(big.NewInt(int64(zint)))
}

//float64 precision
func (r *RPEngine) floatingPart() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	_, zfrac := math.Modf(x)
	r.stack.Push(big.NewFloat(zfrac))
}

//Native math/big.
func (r *RPEngine) sign() {
	rawX, err := r.stack.Pop()
	if err != nil {
		return
	}

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
	rawX, err := r.stack.Pop()
	if err != nil {
		return
	}

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
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	if sign >= 0 {
		r.stack.Push(rawX)
	} else {
		r.stack.Push(rawY)
	}
}

//float64 precision
func (r *RPEngine) min() {
	rawY, err1 := r.stack.Pop()
	if err1 != nil {
		return
	}
	rawX, err2 := r.stack.Pop()
	if err2 != nil {
		r.stack.Push(rawY)
		return
	}

	sign, err := compareHelper(rawX, rawY)
	if err != nil {
		fmt.Println(err.Error())
		r.stack.Push(rawX)
		r.stack.Push(rawY)
		return
	}

	if sign >= 0 {
		r.stack.Push(rawY)
	} else {
		r.stack.Push(rawX)
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
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Exp(x))
	r.pushFloatOrInt(z)
}

//Native implementation in math/big. Arbitrary precision
func (r *RPEngine) factorial() {
	x, err := r.popInt64Helper()
	if err != nil {
		return
	}

	r.stack.Push(new(big.Int).MulRange(1, x))
}

//Native implementation in math/big. Arbitrary precision
func (r *RPEngine) sqrt() {
	rawX, err := r.stack.Pop()
	if err != nil {
		return
	}

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
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Log(x))
	r.pushFloatOrInt(z)
}

//Currently only float64 precision due to use of math.Log10()
func (r *RPEngine) log() {
	x, errx := r.popFloatHelper()
	if errx != nil {
		return
	}

	z := big.NewFloat(math.Log10(x))
	r.pushFloatOrInt(z)
}

//Currently only float64 precision due to use of math.Pow()
func (r *RPEngine) pow() {
	rawY, erry := r.stack.Peek()
	y, erry := r.popFloatHelper()
	if erry != nil {
		return
	}

	x, errx := r.popFloatHelper()
	if errx != nil {
		r.stack.Push(rawY)
		return
	}

	z := big.NewFloat(math.Pow(x, y))
	r.pushFloatOrInt(z)
}

//Networking
//Takes value from top of Stack as little-endian uint32,
//pushes a big-endian uint32 back on the Stack
func (r *RPEngine) hnl() {
	x, err := r.popUint32Helper()
	if err != nil {
		return
	}

	buf := (*[4]byte)(unsafe.Pointer(&x))[:]
	nl := binary.BigEndian.Uint32(buf)
	r.stack.Push(big.NewInt(int64(nl)))
}

//Takes value from top of Stack as little-endian uint16,
//pushes a big-endian uint16 back on the Stack
func (r *RPEngine) hns() {
	x, err := r.popUint32Helper()
	if err != nil {
		return
	}

	buf := (*[2]byte)(unsafe.Pointer(&x))[:]
	ns := binary.BigEndian.Uint16(buf)
	r.stack.Push(big.NewInt(int64(ns)))
}

//Takes value from top of Stack as big-endian uint32,
//pushes a little-endian uint32 back on the Stack
func (r *RPEngine) nhl() {
	x, err := r.popUint32Helper()
	if err != nil {
		return
	}

	buf := (*[4]byte)(unsafe.Pointer(&x))[:]
	hl := binary.LittleEndian.Uint32(buf)
	r.stack.Push(big.NewInt(int64(hl)))
}

//Takes value from top of Stack as big-endian uint16,
//pushes a little-endian uint16 back on the Stack
func (r *RPEngine) nhs() {
	x, err := r.popUint16Helper()
	if err != nil {
		return
	}

	buf := (*[2]byte)(unsafe.Pointer(&x))[:]
	hs := binary.BigEndian.Uint16(buf)
	r.stack.Push(big.NewInt(int64(hs)))
}

//Stack Manipulation
func (r *RPEngine) pick() {
	rawN, errn := r.stack.Peek()
	n, errn := r.popIntHelper()
	if errn != nil {
		return
	}

	rawX, errx := r.stack.Pick(n)
	if errx != nil {
		r.stack.Push(rawN)
		return
	}

	r.stack.Push(copyHelper(rawX))
}

func (r *RPEngine) depth() {
	r.stack.Push(float64(r.stack.Depth()))
}

func (r *RPEngine) drop() {
	r.stack.Drop()
}

func (r *RPEngine) dropn() {
	n, err := r.popIntHelper()
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		r.stack.Drop()
	}
}

func (r *RPEngine) dup() {
	r.stack.Dup()
}

func (r *RPEngine) dupn() {
	n, err := r.popIntHelper()
	if err != nil {
		return
	}

	r.stack.Dupn(n)
}

func (r *RPEngine) roll() {
	n, err := r.popIntHelper()
	if err != nil {
		return
	}

	r.stack.Roll(n)
}

func (r *RPEngine) rolld() {
	n, err := r.popIntHelper()
	if err != nil {
		return
	}

	r.stack.Rolld(n)
}

func (r *RPEngine) swap() {
	r.stack.Swap()
}

//Macros and variables
func (r *RPEngine) repeat(op string) {
	n, err := r.popIntHelper()
	if err != nil {

	}

	for i := 0; i < n; i++ {
		r.Eval([]string{op})
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
	value, err := r.stack.Pop()
	if err != nil {
		return
	}
	r.vars[varName] = value
}

func (r *RPEngine) getVar(key string) {
	value := r.vars[key]
	if value == nil {
		return
	}

	r.stack.Push(value)

}
