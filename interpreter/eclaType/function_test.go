package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
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
		t.Errorf("Expected exactly 1 set of return types, got %d", len(f.Return))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Body))
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
		t.Errorf("Expected exactly 1 set of return types, got %d", len(f.Return))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Body))
	}
	for i, b := range f.Body["int"] {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestAddOverload(t *testing.T) {
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

	f := NewFunction("test", nil, nil, nil)
	if f == nil {
		t.Errorf("Error when creating function")
	}

	f.AddOverload(args, body, ret)

	//test args
	if len(f.Args) != 2 {
		t.Errorf("Expected exactly 2 set of arguments, got %d", len(f.Args))
	}
	for i, arg := range f.Args[1] {
		if arg != args[i] {
			t.Errorf("Expected %v, got %v", args[i], arg)
		}
	}

	//test return
	if len(f.Return) != 2 {
		t.Errorf("Expected exactly 2 set of return types, got %d", len(f.Return))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 2 {
		t.Errorf("Expected exactly 2 bodies, got %d", len(f.Body))
	}
	for i, b := range f.Body["int"] {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestGetTypeEmptyFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetType() != "function()" {
		t.Errorf("Expected function(), got %v", f.GetType())
	}
}

func TestGetTypeWithArgsFunction(t *testing.T) {
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

func TestGetTypeWithReturnsFunction(t *testing.T) {
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

func TestGetValueFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetValue() != f {
		t.Errorf("Expected %v, got %v", f, f.GetValue())
	}
}

func TestStringFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.String() != "function" {
		t.Errorf("Expected function, got %v", f.String())
	}
}

func TestGetStringFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetString() != String("function") {
		t.Errorf("Expected function, got %v", f.GetString())
	}
}

func TestIsNullFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.IsNull() != false {
		t.Errorf("Expected false, got %v", f.IsNull())
	}
}

func TestGetSizeFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)

	expected := utils.Sizeof(f)
	result := f.GetSize()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetBodySimple(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
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

	f := NewFunction("test", args, body, nil)

	if len(f.GetBody()) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Body))
	}
	for i, b := range f.GetBody() {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestGetBodyOverload(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
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

	f := NewFunction("test", args, body, nil)
	f.AddOverload(nil, nil, nil)

	if len(f.GetBody()) != 1 {
		t.Errorf("Expected exactly 1 bodies, got %d", len(f.GetBody()))
	}
	for i, b := range f.GetBody() {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestGetReturnSimple(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	var ret []string
	ret = append(ret, "string")

	f := NewFunction("test", args, nil, ret)

	if len(f.GetReturn()) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Return))
	}
	for i, r := range f.GetReturn() {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}
}

func TestGetReturnOverload(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	var ret []string
	ret = append(ret, "string")

	f := NewFunction("test", args, nil, ret)
	f.AddOverload(nil, nil, nil)

	if len(f.GetReturn()) != 1 {
		t.Errorf("Expected exactly 1 bodies, got %d", len(f.GetReturn()))
	}
	for i, r := range f.GetReturn() {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}
}

func TestOverride(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	f := NewFunction("test", args, nil, nil)

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

	err := f.Override(args, body, ret)
	if err != nil {
		t.Error(err)
	}

	//test return
	if len(f.Return) != 1 {
		t.Errorf("Expected exactly 1 set of return types, got %d", len(f.Return))
	}
	for i, r := range f.Return["int"] {
		if r != ret[i] {
			t.Errorf("Expected %v, got %v", ret[i], r)
		}
	}

	//test body
	if len(f.Body) != 1 {
		t.Errorf("Expected exactly 1 body, got %d", len(f.Body))
	}
	for i, b := range f.Body["int"] {
		if b != body[i] {
			t.Errorf("Expected %v, got %v", body[i], b)
		}
	}
}

func TestGetIndexOfArgsFalse(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "char"})
	var types []Type
	types = append(types, Int(0))

	f := NewAnonymousFunction(args, nil, nil)

	result := f.GetIndexOfArgs(types)
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}
}

func TestGetIndexOfArgsWithSeveralArgsFalse(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "char"})
	var types []Type
	types = append(types, Int(0))
	types = append(types, Int(0))

	f := NewAnonymousFunction(args, nil, nil)

	result := f.GetIndexOfArgs(types)
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}
}

func TestGetIndexOfArgsSimple(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})

	f := NewAnonymousFunction(args, nil, nil)

	var types []Type
	types = append(types, Int(0))

	result := f.GetIndexOfArgs(types)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestGetIndexOfArgsWithAny(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", parser.Any})

	f := NewAnonymousFunction(args, nil, nil)

	var types []Type
	types = append(types, Int(0))

	result := f.GetIndexOfArgs(types)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// Test errors in function

func TestSetValueFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if err := f.SetValue(nil); err == nil {
		t.Errorf("Expected error when setting value of function")
	}
}

func TestGetIndexFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.GetIndex(Int(0)); err == nil {
		t.Errorf("Expected error when getting index of function")
	}
}

func TestAddFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Add(Int(0)); err == nil {
		t.Errorf("Expected error when adding an int to a function")
	}
}

func TestSubFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Sub(Int(0)); err == nil {
		t.Errorf("Expected error when subtracting an int from a function")
	}
}

func TestMulFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Mul(Int(0)); err == nil {
		t.Errorf("Expected error when multiplying a function by an int")
	}
}

func TestDivFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Div(Int(0)); err == nil {
		t.Errorf("Expected error when dividing a function by an int")
	}
}

func TestDivEcFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.DivEc(Int(0)); err == nil {
		t.Errorf("Expected error when getting quotient of a function by an int")
	}
}

func TestModFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Mod(Int(0)); err == nil {
		t.Errorf("Expected error when getting remainder of a function by an int")
	}
}

func TestOrFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Or(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestXorFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Xor(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestAndFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.And(Bool(true)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestNotFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Not(); err == nil {
		t.Errorf("Expected error when getting \"not\" of a function")
	}
}

func TestEqFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Eq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestNotEqFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.NotEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestGtFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Gt(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestGtEqFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.GtEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestLwFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Lw(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestLwEqFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.LwEq(Int(0)); err == nil {
		t.Errorf("Expected error when comparing a function")
	}
}

func TestAppendFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Append(Int(0)); err == nil {
		t.Errorf("Expected error when appending to a function")
	}
}

func TestLenFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if _, err := f.Len(); err == nil {
		t.Errorf("Expected error when getting length of a function")
	}
}

func TestOverrideError(t *testing.T) {
	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	f := NewFunction("test", args, nil, nil)
	args = append(args, parser.FunctionParams{"arg1", "string"})

	err := f.Override(args, nil, nil)
	if err == nil {
		t.Errorf("Expected error when overriding non-existing prototype")
	}
}

func TestGetTypeFunctionEmpty(t *testing.T) {
	expected := "function()"
	foo := NewFunction("test", nil, nil, nil)
	result := foo.GetTypes()

	if len(result) != 1 {
		t.Error("Expected exactly 1 types, got ", len(result))
	}
	if expected != result[0] {
		t.Errorf("Expected %s, got %s", expected, result[0])
	}
}

func TestGetTypeFunctionWithArgsAndReturn(t *testing.T) {
	expected := "function(int, string)(string, int)"

	var args []parser.FunctionParams
	args = append(args, parser.FunctionParams{"arg0", "int"})
	args = append(args, parser.FunctionParams{"arg1", "string"})
	var ret []string
	ret = append(ret, "string")
	ret = append(ret, "int")

	foo := NewFunction("test", args, nil, ret)
	result := foo.GetTypes()

	if len(result) != 1 {
		t.Error("Expected exactly 1 types, got ", len(result))
	}
	if expected != result[0] {
		t.Errorf("Expected %s, got %s", expected, result[0])
	}
}

func TestGetTypeFunctionsWithArgsAndReturn(t *testing.T) {
	var expected []string
	expected = append(expected, "function(int)(string)")
	expected = append(expected, "function(double)(char)")

	var args1 []parser.FunctionParams
	args1 = append(args1, parser.FunctionParams{"arg0", "int"})
	var ret1 []string
	ret1 = append(ret1, "string")

	foo := NewFunction("test", args1, nil, ret1)

	var args2 []parser.FunctionParams
	args2 = append(args2, parser.FunctionParams{"arg0", "double"})
	var ret2 []string
	ret2 = append(ret2, "char")

	foo.AddOverload(args2, nil, ret2)

	result := foo.GetTypes()

	if len(result) != 2 {
		t.Error("Expected exactly 2 types, got ", len(result))
	}
	for i := 0; i < 2; i++ {
		if expected[i] != result[i] {
			t.Errorf("Expected \"%s\", got \"%s\"", expected[i], result[i])
		}
	}
}

/*
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
