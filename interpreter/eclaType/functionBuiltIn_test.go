package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewTypeOf(t *testing.T) {
	typeof := NewTypeOf()

	// test assign
	if typeof == nil {
		t.Error("Expected functionBuiltIn, got ", typeof)
	}

	// test name
	if typeof.Name != "typeOf" || typeof.f == nil {
		t.Error("typeOf assignment went wrong")
	}

	// test func with int
	var arr []Type
	arr = append(arr, Int(0))
	result1, err1 := typeof.f(arr)
	if err1 != nil {
		t.Error(err1)
	}

	if len(result1) != 1 {
		t.Error("typeOf should return exactly one argument")
	}
	if result1[0].String() != parser.Int {
		t.Errorf("Expected %s, got %s", parser.Int, result1[0].String())
	}

	//test func case err
	arr = append(arr, Int(0))
	_, err2 := typeof.f(arr)
	if err2 == nil {
		t.Error("Expected error when getting type of several args")
	}
}

func TestNewSizeOf(t *testing.T) {
	sizeof := NewSizeOf()

	// test assign
	if sizeof == nil {
		t.Error("Expected functionBuiltIn, got ", sizeof)
	}

	// test name
	if sizeof.Name != "sizeOf" || sizeof.f == nil {
		t.Error("sizeOf assignment went wrong")
	}

	// test func with int
	var arr []Type
	arr = append(arr, Int(0))
	result1, err1 := sizeof.f(arr)
	if err1 != nil {
		t.Error(err1)
	}

	if len(result1) != 1 {
		t.Error("sizeOf should return exactly one argument")
	}
	switch result1[0].(type) {
	case Int:
		if result1[0].(Int) != Int(Int(0).GetSize()) {
			t.Errorf("Expected %d, got %d", Int(Int(0).GetSize()), result1[0].(Int))
		}
	default:
		t.Errorf("Expected %T, got %T", Int(0), result1[0])
	}

	//test func case err
	arr = append(arr, Int(0))
	_, err2 := sizeof.f(arr)
	if err2 == nil {
		t.Error("Expected error when getting size of several args")
	}
}

func TestNewLen(t *testing.T) {
	foo := NewLen()

	// test allocation
	if foo == nil {
		t.Error("Expected functionBuiltIn, got nil")
	}

	// test name & f
	if foo.Name != "len" || foo.f == nil {
		t.Error("Error when creating Len")
	}

	// test err
	var arr []Type
	_, err := foo.f(arr)
	if err == nil {
		t.Error("Expected error when calling len with no args")
	}

	// test len
	var list, err1 = NewList(parser.Int)
	if err1 != nil {
		t.Error(err1)
	}
	arr = append(arr, list)
	result, err2 := foo.f(arr)
	if err2 != nil {
		t.Error(err2)
	}

	if len(result) != 1 {
		t.Error("len should return exactly one argument")
	}
	switch result[0].(type) {
	case Int:
		if result[0].(Int) != Int(0) {
			t.Errorf("Expected %d, got %d", Int(0), result[0].(Int))
		}
	default:
		t.Errorf("Expected %T, got %T", Int(0), result[0])
	}
}

func TestAppend(t *testing.T) {
	foo := NewAppend()

	//test allocation
	if foo == nil {
		t.Error("Expected FunctionBuiltIn, got nil")
	}

	//test name & f
	if foo.Name != "append" || foo.f == nil {
		t.Error("Error when creating append")
	}

	//test number of args error
	var arr []Type
	_, err := foo.f(arr)
	if err == nil {
		t.Error("Expected error when appending without args")
	}

	//test append
	list, errList := NewList("[]int")
	if errList != nil {
		t.Error(errList)
	}
	arr = append(arr, list)

	listToAppend := &List{[]Type{Int(0)}, "[]" + parser.Int}
	arr = append(arr, listToAppend)

	result, errResult := foo.f(arr)
	if errResult != nil {
		t.Error(errResult)
		return
	}

	if len(result) != 1 {
		t.Error("Expected 1, got ", len(result))
	}
	switch result[0].(type) {
	case *List:
		if len(result[0].(*List).Value) != 1 {
			t.Error("Expected 1, got ", len(result[0].(*List).Value))
		}
	}

	// test type error
	var args []Type
	args = append(args, Int(0))
	args = append(args, Int(0))
	_, err = foo.f(args)
	if err == nil {
		t.Error("Expected error when appending int to int")
	}
}

func TestCall(t *testing.T) {
	typeof := NewTypeOf()

	// test assign
	if typeof == nil {
		t.Error("Expected functionBuiltIn, got ", typeof)
	}

	// test name
	if typeof.Name != "typeOf" || typeof.f == nil {
		t.Error("typeOf assignment went wrong")
	}

	// test func with int
	var arr []Type
	arr = append(arr, Int(0))
	result1, err1 := typeof.Call(arr)
	if err1 != nil {
		t.Error(err1)
	}

	if len(result1) != 1 {
		t.Error("typeOf should return exactly one argument")
	}
	if result1[0].String() != parser.Int {
		t.Errorf("Expected %s, got %s", parser.Int, result1[0].String())
	}

	//test func case err
	arr = append(arr, Int(0))
	_, err2 := typeof.f(arr)
	if err2 == nil {
		t.Error("Expected error when getting type of several args")
	}
}

func TestGetTypeFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	result := foo.GetType()
	expected := "function()"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetStringFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	result := foo.GetString()
	expected := String("function")

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStringFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	result := foo.String()
	expected := "function"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestIsNullFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	result := foo.IsNull()
	if result {
		t.Error("expected false, got true")
	}
}

func TestGetSizeFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	result := foo.GetSize()
	expected := utils.Sizeof(foo)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetValueFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()

	f := foo.GetValue()
	if f != foo {
		t.Errorf("expected %v, got %v", foo, f)
	}
}

// test err

func TestSetValueFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	err := foo.SetValue(nil)
	if err == nil {
		t.Error("Expected error when setting value of BuiltInFunction")
	}
}

func TestGetIndexFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.GetIndex(Int(0))
	if err == nil {
		t.Error("Expected error when getting index of BuiltInFunction")
	}
}

func TestAddFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Add(Int(0))
	if err == nil {
		t.Error("Expected error when adding to BuiltInFunction")
	}
}

func TestSubFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Sub(Int(0))
	if err == nil {
		t.Error("Expected error when subtracting from BuiltInFunction")
	}
}

func TestMulFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Mul(Int(0))
	if err == nil {
		t.Error("Expected error when multiplying BuiltInFunction")
	}
}

func TestDivFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Div(Int(0))
	if err == nil {
		t.Error("Expected error when dividing BuiltInFunction")
	}
}

func TestDivEcFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.DivEc(Int(0))
	if err == nil {
		t.Error("Expected error when getting quotient BuiltInFunction")
	}
}

func TestModFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Mod(Int(0))
	if err == nil {
		t.Error("Expected error when getting remainder BuiltInFunction")
	}
}

func TestEqFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Eq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestNotEqFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.NotEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestGtFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Gt(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestGtEqFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.GtEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestLwFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Lw(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestLwEqFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.LwEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestAndFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.And(Bool(true))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestOrFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Or(Bool(true))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestXorFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Xor(Bool(true))
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestNotFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Not()
	if err == nil {
		t.Error("Expected error when comparing BuiltInFunction")
	}
}

func TestAppendFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Append(Int(0))
	if err == nil {
		t.Error("Expected error when appending to BuiltInFunction")
	}
}

func TestLenFunctionBuiltIn(t *testing.T) {
	foo := NewTypeOf()
	_, err := foo.Len()
	if err == nil {
		t.Error("Expected error when getting length BuiltInFunction")
	}
}
