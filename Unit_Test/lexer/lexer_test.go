package lexer

import (
	"github.com/tot0p/Ecla/lexer"
	"testing"
)

func TestLexer(t *testing.T) {
	code := `(1+1);`
	l := lexer.Lexer(code)
	if l == nil {
		t.Error("Expected a lexer, got nil")
	} else if len(l) != 7 {
		t.Error("Expected 6 tokens, got ", len(l))
	} else if l[0].Value != "(" {
		t.Error("Expected (, got ", l[0].Value)
	} else if l[1].Value != "1" {
		t.Error("Expected 1, got ", l[1].Value)
	} else if l[2].Value != "+" {
		t.Error("Expected +, got ", l[2].Value)
	} else if l[3].Value != "1" {
		t.Error("Expected 1, got ", l[3].Value)
	} else if l[4].Value != ")" {
		t.Error("Expected ), got ", l[4].Value)
	} else if l[5].Value != ";" {
		t.Error("Expected ;, got ", l[5].Value)
	} else if l[6].Value != "" {
		t.Error("Expected empty string, got ", l[6].Value)
	}
}
