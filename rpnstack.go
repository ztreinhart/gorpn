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
func (s *RPNStack) Pick(n int) float64 {
	pos := len(s.stack) - 1 - n
	if pos >= 0 {
		val := s.stack[pos]
		s.stack = append(s.stack[:n], s.stack[n+1]) //TODO: double check this...
		return val
	} else {
		return math.NaN()
	}
}

func (s *RPNStack) Depth() int {
	return len(s.stack)
}

func (s *RPNStack) Dupn(n int) {
	pos := len(s.stack) - 1 - n
	s.stack = append(s.stack, s.stack[pos:]...)
}

func (s *RPNStack) AsHorizString() string {
	return fmt.Sprintf("%v", s.stack)
}
