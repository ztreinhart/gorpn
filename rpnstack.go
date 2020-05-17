package main

type RPNStack struct {
	stack []float64
}

func (s *RPNStack) Init() {
	s.stack = make([]float64, 10)
}

func (s *RPNStack) push(val float64) {
	s.stack = append(s.stack, val)
}

func (s *RPNStack) pop() float64 {
	end := len(s.stack) - 1
	val := s.stack[end]
	s.stack = s.stack[:end]
	return val
}

func (s *RPNStack) peek() float64 {
	end := len(s.stack) - 1
	return s.stack[end]
}
