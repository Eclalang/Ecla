package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// function interacts with function

func TestGenerateArgsString(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	args = append(args, parser.FunctionParams{"arg1", "string"})

	result := generateArgsString(args)
	expected := "intstring"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestNewFunction(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	var ret []string
	ret = append(ret, "string")
	var body []parser.Node
	body = append(
		body, parser.Literal{
			lexer.Token{
				"type",
				"val",
				0,
				0},
			"type",
			"val"})

	f := NewFunction("test", args, body, ret)

	//test allocation
	if f == nil {
		t.Errorf("Expected function, got %v", f)
	}

	//test name
	if f.Name != "test" {
		t.Error("Expected test, got " + f.Name)
	}

	//test args
	if len(f.Args) != 1 {
		t.Errorf("Expected exactly 1 set of arguments, got %d", len(f.Args))
	}
	for i, arg := range f.Args[0] {
		if arg != args[i] {
			t.Errorf("Expected %v, got %v", args[i], arg)
		}
	}

	//test return
	if len(f.Return) != 1 {
		t.Errorf("Expected exactly 1 set of return types, got %d", len(f.Args))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Args))
	}
	for i, b := range f.Body["int"] {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestNewAnonymousFunction(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	var ret []string
	ret = append(ret, "string")
	var body []parser.Node
	body = append(
		body, parser.Literal{
			lexer.Token{
				"type",
				"val",
				0,
				0},
			"type",
			"val"})

	f := NewAnonymousFunction(args, body, ret)

	//test allocation
	if f == nil {
		t.Errorf("Expected function, got %v", f)
	}

	//test name
	if f.Name != "" {
		t.Error("Expected \"\", got " + f.Name)
	}

	//test args
	if len(f.Args) != 1 {
		t.Errorf("Expected exactly 1 set of arguments, got %d", len(f.Args))
	}
	for i, arg := range f.Args[0] {
		if arg != args[i] {
			t.Errorf("Expected %v, got %v", args[i], arg)
		}
	}

	//test return
	if len(f.Return) != 1 {
		t.Errorf("Expected exactly 1 set of return types, got %d", len(f.Args))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Args))
	}
	for i, b := range f.Body["int"] {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestGetType(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetType() != "function()" {
		t.Errorf("Expected function(), got %v", f.GetType())
	}
}

// Useless tests for code coverage purpose
func TestGetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetValue() == nil {
		t.Error("Expected not nil, got nil")
	}
}

func TestSetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if err := f.SetValue(nil); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func TestString(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.String() != "function" {
		t.Errorf("Expected function, got %v", f.String())
	}
}

func TestGetString(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetString() != "function" {
		t.Errorf("Expected function, got %v", f.GetString().String())
	}
}

func TestGetIndex(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	expect, err := f.GetIndex(nil)
	if expect != nil || err == nil {
		t.Errorf("Expected nil & error, got %v & %v", expect, err)
	}
}

func TestIsNull(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.IsNull() {
		t.Errorf("Expected false, got %v", f.IsNull())
	}
}

func TestTypeAndNumberOfArgsIsCorrect(t *testing.T) {
	var structDecl []eclaDecl.TypeDecl
	f := NewFunction("test", nil, nil, nil)
	var args []Type
	test, expect := f.TypeAndNumberOfArgsIsCorrect(args, structDecl)
	if !test || expect == nil {
		t.Errorf("Expected true, got %v", test)
	}
}

func TestCheckReturn(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	var structDecl []eclaDecl.TypeDecl
	var args []Type
	expect := f.CheckReturn(args, structDecl)
	if !expect {
		t.Errorf("Expected true, got %v", expect)
	}
}
