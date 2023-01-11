package lexer

import (
	"strconv"
	"testing"

	"github.com/tot0p/Ecla/lexer"
)

func TestLexer(t *testing.T) {
	code := `"\"a"`
	expected := []lexer.Token{
		{
			TokenType: lexer.DQUOTE,
			Value:     "\"",
			Position:  0,
			Line:      0,
		},
		{
			TokenType: lexer.STRING,
			Value:     `\"a`,
			Position:  1,
			Line:      0,
		},
		{
			TokenType: lexer.DQUOTE,
			Value:     "\"",
			Position:  4,
			Line:      0,
		},
		{
			TokenType: lexer.EOF,
			Value:     "",
			Position:  len(code),
			Line:      0,
		},
	}

	expected_lenth := len(expected)
	diff := 0
	l := lexer.Lexer(code)
	if l == nil {
		t.Error("Expected a lexer, got nil")
	} else if len(l) != expected_lenth {
		t.Error("Expected "+strconv.Itoa(expected_lenth)+" tokens, got ", len(l))
		diff++
	}
	for Position, expct := range expected {
		if Position < len(l) {
			if expct != l[Position] {
				diff++
				t.Error("Diff ", diff, " Expected ", expct, " for the token n°", Position+1, " , got ", l[Position])
			}
		}
	}
	t.Log("\n\t\t---GLOBAL RESULT---\n")
	t.Log("entry \t <", code, ">")
	t.Log("Expected \t", expected)
	t.Log("Got \t", l)
	t.Log("diff total ", diff)
}
