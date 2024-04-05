package eclaType

import (
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

func TestGetTypeEmpty(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetType() != "function()" {
		t.Errorf("Expected function(), got %v", f.GetType())
	}
}

func TestGetTypeWithArgs(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	args = append(args, parser.FunctionParams{"arg1", "string"})

	f := NewFunction("test", args, nil, nil)
	result := f.GetType()
	expected := "function(int,string)"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetTypeWithReturns(t *testing.T) {
	var ret []string
	ret = append(ret, "string")
	ret = append(ret, "int")

	f := NewFunction("test", nil, nil, ret)
	result := f.GetType()
	expected := "function()(string, int)"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetValue() != f {
		t.Errorf("Expected %v, got %v", f, f.GetValue())
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
	if f.GetString() != String("function") {
		t.Errorf("Expected function, got %v", f.GetString())
	}
}

func TestIsNull(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.IsNull() != false {
		t.Errorf("Expected false, got %v", f.IsNull())
	}
}

// Test errors in function

func TestSetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if err := f.SetValue(nil); err == nil {
		t.Errorf("Expected error when setting value of function")
	}
}

func TestGetIndex(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.GetIndex(Int(0)); err == nil {
		t.Errorf("Expected error when getting index of function")
	}
}

func TestAdd(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Add(Int(0)); err == nil {
		t.Errorf("Expected error when adding an int to a function")
	}
}

func TestSub(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Sub(Int(0)); err == nil {
		t.Errorf("Expected error when subtracting an int from a function")
	}
}

func TestMul(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Mul(Int(0)); err == nil {
		t.Errorf("Expected error when multiplying a function by an int")
	}
}

func TestDiv(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Div(Int(0)); err == nil {
		t.Errorf("Expected error when dividing a function by an int")
	}
}

func TestDivEc(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.DivEc(Int(0)); err == nil {
		t.Errorf("Expected error when getting quotient of a function by an int")
	}
}

func TestMod(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Mod(Int(0)); err == nil {
		t.Errorf("Expected error when getting remainder of a function by an int")
	}
}

func TestOr(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Or(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestXor(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Xor(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestAnd(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.And(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestNot(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Not(); err == nil {
		t.Errorf("Expected error when getting \"not\" of a function")
	}
}

func TestEq(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Eq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestNotEq(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.NotEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestGt(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Gt(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestGtEq(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.GtEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestLw(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Lw(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestLwEq(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.LwEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestAppend(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Append(Int(0)); err == nil {
		t.Errorf("Expected error when appending to a function")
	}
}

/*
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
*/
