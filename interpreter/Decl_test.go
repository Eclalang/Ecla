package interpreter

import (
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewINT(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.INT, Value: "0"}, &Env{})
	if bus.GetVal().GetType() != "int" {
		t.Error("Expected INT, got", bus.GetVal().GetType())
	}
}

func TestNewSTRING(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.STRING, Value: "test"}, &Env{})
	if bus.GetVal().GetType() != "string" {
		t.Error("Expected STRING, got", bus.GetVal().GetType())
	}

	//bus = New(parser.Literal{Type: lexer.STRING, Value: "\n"}, &Env{})

}
