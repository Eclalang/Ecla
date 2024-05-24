package parser

import (
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

func TestTokenPrecedence(t *testing.T) {
	OR := lexer.Token{TokenType: lexer.OR, Value: "OR"}
	XOR := lexer.Token{TokenType: lexer.XOR, Value: "XOR"}
	AND := lexer.Token{TokenType: lexer.AND, Value: "AND"}
	EQUAL := lexer.Token{TokenType: lexer.EQUAL, Value: "EQUAL"}
	NEQ := lexer.Token{TokenType: lexer.NEQ, Value: "NEQ"}
	LSS := lexer.Token{TokenType: lexer.LSS, Value: "LSS"}
	LEQ := lexer.Token{TokenType: lexer.LEQ, Value: "LEQ"}
	GTR := lexer.Token{TokenType: lexer.GTR, Value: "GTR"}
	GEQ := lexer.Token{TokenType: lexer.GEQ, Value: "GEQ"}
	ADD := lexer.Token{TokenType: lexer.ADD, Value: "ADD"}
	SUB := lexer.Token{TokenType: lexer.SUB, Value: "SUB"}
	MULT := lexer.Token{TokenType: lexer.MULT, Value: "MULT"}
	DIV := lexer.Token{TokenType: lexer.DIV, Value: "DIV"}
	MOD := lexer.Token{TokenType: lexer.MOD, Value: "MOD"}
	QOT := lexer.Token{TokenType: lexer.QOT, Value: "QOT"}
	RANDOM := lexer.Token{TokenType: lexer.TEXT, Value: "RANDOM"}

	if TokenPrecedence(OR) != 1 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(XOR) != 2 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(AND) != 3 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(EQUAL) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(NEQ) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(LSS) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(LEQ) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(GTR) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(GEQ) != 4 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(ADD) != 5 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(SUB) != 5 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(MULT) != 6 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(DIV) != 6 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(MOD) != 6 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(QOT) != 6 {
		t.Error("TokenPrecedence failed to return the correct value")
	}
	if TokenPrecedence(RANDOM) != LowestPrecedence {
		t.Error("TokenPrecedence failed to return the correct value")
	}
}
