package parser

import (
	"github.com/Eclalang/Ecla/lexer"
)

const (
	LowestPrecedence  = 0
	HighestPrecedence = 7
)

func TokenPrecedence(tok lexer.Token) int {
	switch tok.TokenType {
	case lexer.OR:
		return 1
	case lexer.XOR:
		return 2
	case lexer.AND:
		return 3
	case lexer.EQUAL, lexer.NEQ, lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ:
		return 4
	case lexer.ADD, lexer.SUB:
		return 5
	case lexer.MULT, lexer.DIV, lexer.MOD, lexer.QOT:
		return 6
	}
	return LowestPrecedence
}
