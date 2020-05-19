package main

import (
	"fmt"
	"math"
)

type RPNStack struct {
	stack []float64
}

func (s *RPNStack) Init() {
	s.stack = make([]float64, 0)
}

func (s *RPNStack) Clear() {
	s.Init()
}

func (s *RPNStack) Push(val float64) {
	s.stack = append(s.stack, val)
}

func (s *RPNStack) Pop() float64 { //Returns NaN if stack empty
	if len(s.stack) < 1 {
		return math.NaN()
	}
	end := len(s.stack) - 1
	val := s.stack[end]
	s.stack = s.stack[:end]
	return val
}

func (s *RPNStack) Peek() float64 { //Returns NaN if stack empty
	end := len(s.stack) - 1
	if end >= 0 {
		return s.stack[end]
	} else {
		return math.NaN()
	}
}

func (s *RPNStack) PushBottom(val float64) {
	s.stack = append([]float64{val}, s.stack...)
}

func (s *RPNStack) PopBottom() float64 { //Returns NaN if stack empty
	if len(s.stack) < 1 {
		return math.NaN()
	}
	val := s.stack[0]
	s.stack = s.stack[1:]
	return val
}

func (s *RPNStack) PeekBottom() float64 { //Returns NaN if stack empty
	if len(s.stack) < 1 {
		return math.NaN()
	}
	return s.stack[0]
}

func (s *RPNStack) Pick(n int) float64 { //Returns NaN if out of bounds
	pos := len(s.stack) - 1 - n
	if pos >= 0 {
		val := s.stack[pos]
		s.stack = append(s.stack[:n], s.stack[n+1:]...)
		return val
	} else {
		return math.NaN()
	}
}

func (s *RPNStack) Depth() int {
	return len(s.stack)
}
func (s *RPNStack) Drop() {
	if len(s.stack) > 0 {
		_ = s.Pop()
	}
}

func (s *RPNStack) DropBottom() {
	if len(s.stack) > 0 {
		_ = s.PopBottom()
	}
}

func (s *RPNStack) Dropn(n int) {
	if n >= len(s.stack) {
		s.Clear()
		return
	}
	for i := 0; i < n; i++ {
		s.Drop()
	}
}

func (s *RPNStack) DropBottomn(n int) {
	if n >= len(s.stack) {
		s.Clear()
		return
	}
	for i := 0; i < n; i++ {
		s.DropBottom()
	}
}

func (s *RPNStack) Dup() {
	if len(s.stack) > 0 {
		s.Push(s.Peek())
	}
}

func (s *RPNStack) Dupn(n int) { //If n >= stack length, duplicates whole stack
	if n <= 0 {
		return
	}
	pos := len(s.stack) - 1 - n
	if pos < 0 {
		pos = 0
	}
	s.stack = append(s.stack, s.stack[pos:]...)
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
	if len(s.stack) > 1 {
		first := s.Pop()
		second := s.Pop()
		s.Push(first)
		s.Push(second)
	}
}

func (s *RPNStack) AsHorizString() string {
	return fmt.Sprintf("%v", s.stack)
}
