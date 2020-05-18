package main

import "fmt"

type RPNStack struct {
	stack []float64
}

func (s *RPNStack) Init() {
	s.stack = make([]float64, 0)
}

func (s *RPNStack) Push(val float64) {
	s.stack = append(s.stack, val)
}

func (s *RPNStack) Pop() float64 {
	end := len(s.stack) - 1
	val := s.stack[end]
	s.stack = s.stack[:end]
	return val
}

func (s *RPNStack) Peek() float64 {
	end := len(s.stack) - 1
	if end >= 0 {
		return s.stack[end]
	} else {
		return 0
	}
}

func (s *RPNStack) AsHorizString() string {
	return fmt.Sprintf("%v", s.stack)
}
