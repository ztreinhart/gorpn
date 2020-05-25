package main

import (
	"errors"
	"math/big"
)

type RPNStack struct {
	Stack []interface{}
}

func (s *RPNStack) Init() {
	s.Stack = make([]interface{}, 0)
}

func (s *RPNStack) Clear() {
	s.Init()
}

func (s *RPNStack) Push(val interface{}) {
	s.Stack = append(s.Stack, val)
}

func (s *RPNStack) Pop() (interface{}, error) { //Returns nil if stack empty
	if len(s.Stack) < 1 {
		return nil, errors.New("stack empty")
	}
	end := len(s.Stack) - 1
	val := s.Stack[end]
	s.Stack = s.Stack[:end]
	return val, nil
}

func (s *RPNStack) Peek() (interface{}, error) { //Returns nil if stack empty
	end := len(s.Stack) - 1
	if end < 0 {
		return nil, errors.New("stack empty")
	}
	return s.Stack[end], nil
}

func (s *RPNStack) PushBottom(val interface{}) {
	s.Stack = append([]interface{}{val}, s.Stack...)
}

func (s *RPNStack) PopBottom() (interface{}, error) { //Returns nil if stack empty
	if len(s.Stack) < 1 {
		return nil, errors.New("stack empty")
	}
	val := s.Stack[0]
	s.Stack = s.Stack[1:]
	return val, nil
}

func (s *RPNStack) PeekBottom() (interface{}, error) { //Returns nil if stack empty
	if len(s.Stack) < 1 {
		return nil, errors.New("stack empty")
	}
	return s.Stack[0], nil
}

func (s *RPNStack) Pick(n int) (interface{}, error) {
	end := len(s.Stack) - 1
	pos := end - n
	if pos >= 0 && pos <= end {
		val := s.Stack[pos]
		s.Stack = append(s.Stack[:pos], s.Stack[pos+1:]...)
		return val, nil
	}
	return nil, errors.New("pick index out of bounds")
}

func (s *RPNStack) Depth() int {
	return len(s.Stack)
}

func (s *RPNStack) Drop() {
	if len(s.Stack) > 0 {
		_, _ = s.Pop()
	}
}

func (s *RPNStack) DropBottom() {
	if len(s.Stack) > 0 {
		_, _ = s.PopBottom()
	}
}

func (s *RPNStack) Dropn(n int) {
	if n >= len(s.Stack) {
		s.Clear()
		return
	}
	for i := 0; i < n; i++ {
		s.Drop()
	}
}

func (s *RPNStack) DropBottomn(n int) {
	if n >= len(s.Stack) {
		s.Clear()
		return
	}
	for i := 0; i < n; i++ {
		s.DropBottom()
	}
}

func (s *RPNStack) Dup() {
	item, err := s.Peek()
	if err == nil {
		s.Push(copyHelper(item))
	}
}

func (s *RPNStack) Dupn(n int) { //If n >= Stack length, duplicates whole Stack
	if n <= 0 {
		return
	}
	buf := make([]interface{}, 0)
	for i := len(s.Stack) - n; i < len(s.Stack); i++ {
		buf = append(buf, copyHelper(s.Stack[i]))
	}
	s.Stack = append(s.Stack, buf...)
}

func (s *RPNStack) Roll(n int) {
	if n < 0 {
		s.Rolld(-n)
		return
	}
	for i := 0; i < n; i++ {
		item, err := s.Pop()
		if err == nil {
			s.PushBottom(item)
		}
	}
}

func (s *RPNStack) Rolld(n int) {
	if n < 0 {
		s.Roll(-n)
		return
	}
	for i := 0; i < n; i++ {
		item, err := s.PopBottom()
		if err == nil {
			s.Push(item)
		}
	}
}

func (s *RPNStack) Swap() {
	if len(s.Stack) > 1 {
		first, err1 := s.Pop()
		if err1 != nil {
			return
		}

		second, err2 := s.Pop()
		if err2 != nil {
			s.Push(first)
			return
		}

		s.Push(first)
		s.Push(second)
	}
}

//func (s *RPNStack) AsHorizString() string {
//	return fmt.Sprintf("%v", s.Stack)
//}

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
