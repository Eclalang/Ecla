package parser

import (
	"github.com/tot0p/Ecla/lexer"
)

const (
	LowestPrecedence  = 0
	HighestPrecedence = 6
)

func TokenPrecedence(tok lexer.Token) int {
	switch tok.TokenType {
	case lexer.OR:
		return 1
	case lexer.AND:
		return 2
	case lexer.EQUAL, lexer.NEQ, lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ:
		return 3
	case lexer.ADD, lexer.SUB:
		return 4
	case lexer.MULT, lexer.DIV, lexer.MOD, lexer.QOT:
		return 5
	}
	return LowestPrecedence
}
