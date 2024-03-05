package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// New

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
	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.STRING, Value: "\n"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewBool(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.BOOL, Value: "true"}, &Env{})
	if bus.GetVal().GetType() != "bool" {
		t.Error("Expected BOOL, got", bus.GetVal().GetType())
	}

	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.BOOL, Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewFLOAT(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.FLOAT, Value: "0.0"}, &Env{})
	if bus.GetVal().GetType() != "float" {
		t.Error("Expected FLOAT, got", bus.GetVal().GetType())
	}
}

func TestNewCHAR(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.CHAR, Value: "a"}, &Env{})
	if bus.GetVal().GetType() != "char" {
		t.Error("Expected CHAR, got", bus.GetVal().GetType())
	}

	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.CHAR, Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewVAR(t *testing.T) {
	s, err := eclaType.NewString("test")
	if err != nil {
		t.Error("Error creating string: ", err.Error())
	}
	env := NewEnv()
	v, err := eclaType.NewVar("test", "", s)
	if err != nil {
		t.Error("Error creating var: ", err.Error())
	}
	env.SetVar("test", v)

	bus := New(parser.Literal{Type: "VAR", Value: "test"}, env)
	if bus.GetVal().GetType() != "string" {
		t.Error("Expected STRING, got", bus.GetVal().GetType())
	}

	var errCheck = false
	env.ErrorHandle.HookExit(
		func(i int) {
			errCheck = true
		})

	bus = New(parser.Literal{Type: "VAR", Value: "test2"}, env)
	if !errCheck {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewNull(t *testing.T) {
	bus := New(parser.Literal{Type: "NULL", Value: "null"}, &Env{})
	if bus.GetVal().GetType() != "" {
		t.Error("Expected NULL, got", bus.GetVal().GetType())
	}
}

func TestNewIfNodeNotExist(t *testing.T) {
	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus := New(parser.Literal{Type: "NODE", Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}
