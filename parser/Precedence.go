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
	case "temp3":
		return 3
	case lexer.ADD, lexer.SUB:
		return 4
	case lexer.MULT, lexer.DIV, lexer.MOD:
		return 5
	}
	return LowestPrecedence
}
