package main

import "math/big"

type ArgType int

const (
	integer ArgType = iota
	floating
	boolean
)

func (a ArgType) String() string {
	return [...]string{"integer", "floating", "boolean"}[a]
}

type RPNArg struct {
	typeof   ArgType
	integer  *big.Int
	floating *big.Float
	boolean  *bool
}
