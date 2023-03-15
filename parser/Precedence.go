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
	case "temp1":
		return 1
	case "temp2":
		return 2
	case lexer.EQUAL, lexer.NEQ, lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ:
		return 3
	case lexer.ADD, lexer.SUB, lexer.OR:
		return 4
	case lexer.MULT, lexer.DIV, lexer.MOD, lexer.QOT, lexer.AND:
		return 5
	}
	return LowestPrecedence
}
