package eclaType

import (
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewAny(t *testing.T) {
	t1 := NewAny(Bool(true))
	expected := &Any{Bool(true), parser.Bool}

	if t1.GetValue() != expected.GetValue() {
		t.Error("expected \"true\", got ", t1)
	}
}

func TestNewEmptyAny(t *testing.T) {
	t1, _ := NewAnyEmpty()
	expected := &Any{Value: NewNull()}

	if t1.GetValue() != expected.GetValue() {
		t.Error("expected \"true\", got ", t1)
	}
}

func TestAnyString(t *testing.T) {
	t1 := NewAny(Bool(true))
	expected := "true"
	result := t1.String()

	if result != expected {
		t.Error("expected \"true\", got ", result)
	}
}

func TestAnyGetString(t *testing.T) {
	t1 := NewAny(Bool(true))
	expected, _ := NewString("true")
	result := t1.GetString()

	if result != expected {
		t.Error("expected \"true\", got ", result)
	}
}

func TestAnyGetValue(t *testing.T) {
	result := Bool(true)
	t1 := NewAny(result)

	if result != t1.GetValue() {
		t.Error("expected ", t1, ", got ", result)
	}
}

func TestAnySetValue(t *testing.T) {
	t2, _ := NewList("[]int")
	t1 := NewAny(t2)
	var values []Type
	values = append(values, Int(0))
	err := t1.SetValue(values)

	if err != nil {
		t.Error(err)
	}
}

func TestAnySetAny(t *testing.T) {
	t1 := NewAny(Bool(true))
	t2 := NewAny(Int(0))

	err := t1.SetAny(Int(0))
	if err != nil {
		t.Error(err)
	}

	if t1.GetValue() != t2.GetValue() {
		t.Error("expected ", t2, ", got ", t1)
	}
}

func TestGetIndexAnyInt(t *testing.T) {
	t1 := NewAny(String("123"))
	t2 := Int(0)

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

func TestAddAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestSubAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestMulAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(3))
	t2 := NewAny(Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(6) {
		t.Error("Expected 6, got ", result)
	}
}

func TestDivAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.5) {
		t.Error("Expected 0.5, got ", result)
	}
}

func TestModAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(51))
	t2 := NewAny(Int(2))

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestDivEcAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(9))
	t2 := NewAny(Int(2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(4) {
		t.Error("Expected 4, got ", result)
	}
}

func TestEqAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqAnyAsInt(t *testing.T) {
	t1 := NewAny(Int(1))
	t2 := NewAny(Int(2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndAny(t *testing.T) {
	t1 := NewAny(Bool(true))
	t2 := NewAny(Bool(false))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrAny(t *testing.T) {
	t1 := NewAny(Bool(true))
	t2 := NewAny(Bool(false))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorAny(t *testing.T) {
	t1 := NewAny(Bool(true))
	t2 := NewAny(Bool(false))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotAny(t *testing.T) {
	t1 := NewAny(Bool(false))

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestIncrementAny(t *testing.T) {
	t1 := NewAny(Int(1))

	t1.Increment()
	if t1.GetValue() != Int(2) {
		t.Error("Expected 2, got ", t1)
	}
}

func TestDecrementAny(t *testing.T) {
	t1 := NewAny(Int(1))

	t1.Decrement()
	if t1.GetValue() != Int(0) {
		t.Error("Expected 0, got ", t1)
	}
}

func TestAppendAny(t *testing.T) {
	t1 := NewAny(String("123"))
	t2 := NewAny(String("4"))

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

func TestIsNullAny(t *testing.T) {
	t1, _ := NewAnyEmpty()

	result := t1.IsNull()
	if !result {
		t.Error("Expected \"true\", got ", result)
	}
}

func TestGetFunctionAny(t *testing.T) {
	expected := NewFunction("test", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
	t1 := NewAny(expected)

	result := t1.GetFunction()
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

//TODO @Sanegv investigate if you put a safety to prevent any of any
//func TestGetFunctionAnyAny(t *testing.T) {
//	expected := NewFunction("test", []parser.FunctionParams{}, []parser.Node{}, []string{"int"})
//	a := NewAny(expected)
//	t1 := NewAny(a)
//
//	result := t1.GetFunction()
//	if result != expected {
//		t.Error("Expected ", expected, ", got ", result)
//	}
//}

func TestGetFunctionAnyNil(t *testing.T) {
	t1, _ := NewAnyEmpty()

	result := t1.GetFunction()
	if result != nil {
		t.Error("Expected \"nil\", got ", result)
	}
}

func TestGetSizeAny(t *testing.T) {
	t1 := NewAny(Int(0))
	expected := Int(0).GetSize()
	result := t1.GetSize()
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

func TestAnyLen(t *testing.T) {
	t1 := NewAny(String("test"))
	expected := 4

	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
