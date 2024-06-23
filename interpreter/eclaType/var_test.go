package eclaType

import (
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewVarEmpty(t *testing.T) {
	t1, err := NewVarEmpty("test", "int")
	if err != nil {
		t.Error(err)
	}

	expected := &Var{Name: "test", Value: NewNullType("int")}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarString(t *testing.T) {
	t1, err := NewVar("test", parser.String, String("value"))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", String("value")}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarFloat(t *testing.T) {
	t1, err := NewVar("test", parser.Float, Int(1))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", Float(1.0)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarAny(t *testing.T) {
	tmp := NewAny(Int(0))
	t1, err := NewVar("test", parser.Any, Int(0))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", tmp}

	if t1.Value == expected.Value {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarNull(t *testing.T) {
	t1, err := NewVar("test", parser.Int, NewNullType(parser.Int))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", NewNullType(parser.Int)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarEmptyType(t *testing.T) {
	t1, err := NewVar("test", "", Int(0))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", Int(0)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarErr(t *testing.T) {
	_, err := NewVar("test", parser.Bool, Int(1))

	if err == nil {
		t.Error("Expected error when creating a var with the wrong type")
	}
}

func TestNewVarVar(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	t2, e := NewVar("test", parser.Int, t1)
	if e != nil {
		t.Error(e)
	}

	if *t1 != *t2 {
		t.Error("Expected ", t1, ", got ", t2)
	}
}

func TestVarString(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}
	expected := "test = 0"

	result := t1.String()

	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

func TestVarGetString(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}
	expected := String("0")

	result := t1.GetString()

	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

func TestVarGetValue(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.GetValue()

	if result != Int(0) {
		t.Error("Expected ", Int(0), ", got ", result)
	}
}

func TestVarSetValue(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.SetValue(Int(0))

	switch result.(type) {
	case error:
		return
	}
	t.Error("Expected error, got ", result)
}

func TestVarGetType(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.GetType()

	if result != parser.Int {
		t.Error("Expected ", parser.Int, ", got ", result)
	}
}

func TestGetIndexVarInt(t *testing.T) {
	t1, err := NewVar("test", parser.String, String("123"))
	if err != nil {
		t.Error(err)
	}
	t2 := Int(0)

	result, e := t1.GetIndex(t2)
	if e != nil {
		t.Error(e)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

func TestSetVarInt(t *testing.T) {
	t1, err := NewVarEmpty("test", parser.Int)
	if err != nil {
		t.Error(err)
	}
	t2 := Int(0)

	e := t1.SetVar(t2)
	if e != nil {
		t.Error(e)
	}
	if t1.Value != t2 {
		t.Error("Expected ", t2, ", got ", t1.Value)
	}
}

func TestSetVarVarInt(t *testing.T) {
	t1, err := NewVarEmpty("test", parser.Int)
	if err != nil {
		t.Error(err)
	}
	t2, err2 := NewVar("i", parser.Int, Int(0))
	if err2 != nil {
		t.Error(err2)
	}

	e := t1.SetVar(t2)
	if e != nil {
		t.Error(e)
	}
	if t1.Value != t2.Value {
		t.Error("Expected ", t2, ", got ", t1.Value)
	}
}

func TestSetVarNull(t *testing.T) {
	t1, err := NewVarEmpty("test", parser.Int)
	if err != nil {
		t.Error(err)
	}
	t2 := NewNullType(parser.Int)

	e := t1.SetVar(t2)
	if e != nil {
		t.Error(e)
	}
	if t1.Value != t2 {
		t.Error("Expected ", t2, ", got ", t1.Value)
	}
}

func TestSetVarErr(t *testing.T) {
	t1, err := NewVarEmpty("test", parser.Int)
	if err != nil {
		t.Error(err)
	}
	t2 := Bool(true)

	e := t1.SetVar(t2)
	if e == nil {
		t.Error("expected error when assigning boolean to int")
	}
}

func TestAddVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestSubVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestMulVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(2))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(3))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(6) {
		t.Error("Expected 6, got ", result)
	}
}

func TestDivVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.5) {
		t.Error("Expected 0.5, got ", result)
	}
}

func TestModVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(23))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestDivEcVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(9))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(4) {
		t.Error("Expected 4, got ", result)
	}
}

func TestEqVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(9))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqVarAsInt(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Int, Int(2))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Bool, Bool(true))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Bool, Bool(false))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Bool, Bool(true))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Bool, Bool(false))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Bool, Bool(true))
	if err1 != nil {
		t.Error(err1)
	}
	t2, err2 := NewVar("b", parser.Bool, Bool(false))
	if err2 != nil {
		t.Error(err2)
	}

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Bool, Bool(false))
	if err1 != nil {
		t.Error(err1)
	}

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestIncrementVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}

	t1.Increment()
	if t1.GetValue() != Int(2) {
		t.Error("Expected 2, got ", t1)
	}
}

func TestDecrementVar(t *testing.T) {
	t1, err1 := NewVar("a", parser.Int, Int(1))
	if err1 != nil {
		t.Error(err1)
	}

	t1.Decrement()
	if t1.GetValue() != Int(0) {
		t.Error("Expected 0, got ", t1)
	}
}

func TestAppendVar(t *testing.T) {
	t1, err1 := NewVar("test", parser.String, String("123"))
	if err1 != nil {
		t.Error(err1)
	}
	t2 := String("4")

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

func TestIsNullVar(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.IsNull()
	if result {
		t.Error("Expected false, got ", result)
	}
}

func TestIsFunctionVarFalse(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.IsFunction()
	if result {
		t.Error("Expected false, got ", result)
	}
}

func TestIsFunctionVarFunction(t *testing.T) {
	foo := NewFunction("foo", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
	t1, err := NewVar("test", foo.GetType(), foo)
	if err != nil {
		t.Error(err)
	}

	result := t1.IsFunction()
	if !result {
		t.Error("Expected true, got ", result)
	}
}

func TestIsFunctionVarAnyFunction(t *testing.T) {
	foo := NewFunction("foo", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
	a := NewAny(foo)
	t1, err := NewVar("test", a.GetType(), a)
	if err != nil {
		t.Error(err)
	}

	result := t1.IsFunction()
	if !result {
		t.Error("Expected true, got ", result)
	}
}

func TestIsAnyTrue(t *testing.T) {
	a := NewAny(Int(0))
	t1, err := NewVar("test", a.GetType(), a)
	if err != nil {
		t.Error(err)
	}

	result := t1.IsAny()
	if !result {
		t.Error("expected true, got ", result)
	}
}

func TestIsAnyFalse(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	result := t1.IsAny()
	if result {
		t.Error("expected false, got ", result)
	}
}

func TestGetFunctionVar(t *testing.T) {
	expected := NewFunction("test", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
	t1, err := NewVar("foo", expected.GetType(), expected)
	if err != nil {
		t.Error(err)
	}

	result := t1.GetFunction()
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

func TestGetFunctionVarAny(t *testing.T) {
	expected := NewFunction("foo", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
	a := NewAny(expected)
	t1, err := NewVar("test", a.GetType(), a)
	if err != nil {
		t.Error(err)
	}

	result := t1.GetFunction()
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

//TODO @Sanegv investigate if you put a safety to prevent any of any
//func TestGetFunctionVarAnyAny(t *testing.T) {
//	expected := NewFunction("foo", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
//	a := NewAny(expected)
//	a2 := NewAny(a)
//	t1, err := NewVar("test", a2.GetType(), a2)
//	if err != nil {
//		t.Error(err)
//	}
//
//	result := t1.GetFunction()
//	if result != expected {
//		t.Error("Expected ", expected, ", got ", result)
//	}
//}

func TestGetFunctionVarNil(t *testing.T) {
	t1, _ := NewVarEmpty("test", parser.Int)

	result := t1.GetFunction()
	if result != nil {
		t.Error("Expected \"nil\", got ", result)
	}
}

func TestGetSizeVar(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	expected := Int(0).GetSize()
	result := t1.GetSize()
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

func TestVarLen(t *testing.T) {
	t1, err := NewVar("test", parser.String, String("1234"))
	if err != nil {
		t.Error(err)
	}
	expected := 4

	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
