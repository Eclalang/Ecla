package parser

import (
	"github.com/tot0p/Ecla/lexer"
)

func TokenPrecedence(tok lexer.Token) int {
	switch tok.TokenType {
	case "temp1":
		return 1
	case "temp2":
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
