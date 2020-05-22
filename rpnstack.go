package main

import (
	"fmt"
)

//TODO: Consider switching all 'n' arguments to uint.

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

func (s *RPNStack) Pop() interface{} { //Returns nil if Stack empty
	if len(s.Stack) < 1 {
		return nil
	}
	end := len(s.Stack) - 1
	val := s.Stack[end]
	s.Stack = s.Stack[:end]
	return val
}

func (s *RPNStack) Peek() interface{} { //Returns nil if Stack empty
	end := len(s.Stack) - 1
	if end >= 0 {
		return s.Stack[end]
	} else {
		return nil
	}
}

func (s *RPNStack) PushBottom(val interface{}) {
	s.Stack = append([]interface{}{val}, s.Stack...)
}

func (s *RPNStack) PopBottom() interface{} { //Returns nil if Stack empty
	if len(s.Stack) < 1 {
		return nil
	}
	val := s.Stack[0]
	s.Stack = s.Stack[1:]
	return val
}

func (s *RPNStack) PeekBottom() interface{} { //Returns nil if Stack empty
	if len(s.Stack) < 1 {
		return nil
	}
	return s.Stack[0]
}

func (s *RPNStack) Pick(n int) interface{} { //Returns NaN if out of bounds
	pos := len(s.Stack) - 1 - n
	if pos >= 0 {
		val := s.Stack[pos]
		s.Stack = append(s.Stack[:n], s.Stack[n+1:]...)
		return val
	} else {
		return nil
	}
}

func (s *RPNStack) Depth() int {
	return len(s.Stack)
}
func (s *RPNStack) Drop() {
	if len(s.Stack) > 0 {
		_ = s.Pop()
	}
}

func (s *RPNStack) DropBottom() {
	if len(s.Stack) > 0 {
		_ = s.PopBottom()
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
	if len(s.Stack) > 0 {
		s.Push(s.Peek())
	}
}

func (s *RPNStack) Dupn(n int) { //If n >= Stack length, duplicates whole Stack
	if n <= 0 {
		return
	}
	pos := len(s.Stack) - 1 - n
	if pos < 0 {
		pos = 0
	}
	s.Stack = append(s.Stack, s.Stack[pos:]...)
}

func (s *RPNStack) Roll(n int) {
	if n < 0 {
		s.Rolld(-n)
		return
	}
	for i := 0; i < n; i++ {
		s.PushBottom(s.Pop())
	}
}

func (s *RPNStack) Rolld(n int) {
	if n < 0 {
		s.Roll(-n)
		return
	}
	for i := 0; i < n; i++ {
		s.Push(s.PopBottom())
	}
}

func (s *RPNStack) Swap() {
	if len(s.Stack) > 1 {
		first := s.Pop()
		second := s.Pop()
		s.Push(first)
		s.Push(second)
	}
}

func (s *RPNStack) AsHorizString() string {
	return fmt.Sprintf("%v", s.Stack)
}
