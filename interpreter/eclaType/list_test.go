package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// List interacts with List

func TestListNewList(t *testing.T) {
	t1, _ := NewList(parser.Int)
	t2 := &List{[]Type{}, parser.Int}

	switch t1.(type) {
	case *List:
		if t1.(*List).Typ != t2.Typ {
			t.Error("cannot compare list of " + t1.(*List).Typ + " with list of " + t2.Typ)
		}
		if len(t1.(*List).Value) != len(t2.Value) {
			t.Error("error when creating a new list")
		}
		for i, v := range t1.(*List).Value {
			if v != t2.Value[i] {
				t.Error("error when creating a new list")
			}
		}
	}
}

func TestListGetValue(t *testing.T) {
	t1, _ := NewList(parser.Int)
	t2 := t1.GetValue()

	if t1 != t2 {
		t.Error("error when getting value of list")
	}
}

func TestListSetValueWithList(t *testing.T) {
	t1, _ := NewList(parser.Int)
	var array []Type
	array = append(array, Int(0))
	t2 := &List{array, parser.Int}
	err := t1.SetValue(t2)

	if err != nil {
		t.Error(err)
	}
	switch t1.(type) {
	case *List:
		if t1.(*List).Typ != t2.Typ {
			t.Errorf("Expected %s, got %s", t1.(*List).Typ, t2.Typ)
		}
		if len(t1.(*List).Value) != len(t2.Value) {
			t.Errorf("Expected list of length %d, got list of length %d", len(t1.(*List).Value), len(t2.Value))
		}
		for i, elem := range t1.(*List).Value {
			if elem != t2.Value[i] {
				t.Error("The lists contain different elements")
			}
		}
	default:
		t.Errorf("Expected %T, got %T", t2, t1)
	}
}

func TestListSetValueWithSliceOfTypes(t *testing.T) {
	t1 := &List{[]Type{}, "[]int"}
	var array []Type
	array = append(array, Int(0))
	err := t1.SetValue(array)

	if err != nil {
		t.Error(err)
	}

	if t1.Typ[2:] != array[0].GetType() {
		t.Errorf("Expected %s, got %s", t1.Typ[2:], parser.Int)
	}
	if len(t1.Value) != len(array) {
		t.Errorf("Expected list of length %d, got list of length %d", len(t1.Value), len(array))
	}
	for i, elem := range t1.Value {
		if elem != array[i] {
			t.Error("The lists contain different elements")
		}
	}
}

func TestListString(t *testing.T) {
	t1 := &List{[]Type{Int(0), Int(1), Int(2)}, parser.Int}
	expected := "[0, 1, 2]"
	result := t1.String()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestListGetString(t *testing.T) {
	t1 := &List{[]Type{Int(0), Int(1), Int(2)}, parser.Int}
	expected := String("[0, 1, 2]")
	result := t1.GetString()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestListGetType(t *testing.T) {
	expected := "test type"
	t1, err := NewList(expected)
	if err != nil {
		t.Error(err)
	}

	result := t1.GetType()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestListSetType(t *testing.T) {
	expected := "test type"
	t1 := &List{[]Type{}, "wrong type"}

	t1.SetType(expected)
	if t1.Typ != expected {
		t.Errorf("Expected %s, got %s", expected, t1.Typ)
	}
}

func TestListGetIndex(t *testing.T) {
	expected := Int(5)
	t1 := &List{[]Type{Int(3), expected}, parser.Int}
	result, err := t1.GetIndex(Int(1))

	if err != nil {
		t.Error(err)
	}

	if *result != expected {
		t.Errorf("Expected %d, got %d", expected, *result)
	}
}

func TestListIsNull(t *testing.T) {
	t1 := &List{[]Type{}, parser.Int}

	if t1.IsNull() {
		t.Error("Expected false, got true")
	}
}

func TestListGetValueType(t *testing.T) {
	expected := "test"
	t1 := &List{[]Type{}, "[]" + expected}
	result := t1.GetValueType()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestListCheckTypeOfListTrue(t *testing.T) {
	t1 := &List{[]Type{Int(1), Int(0)}, parser.Int}
	if !CheckTypeOfList(t1, parser.Int) {
		t.Error("Expected true, got false")
	}
}

func TestListCheckTypeOfListFalse(t *testing.T) {
	t1 := &List{[]Type{Int(1), Char('c')}, parser.Int}
	if CheckTypeOfList(t1, parser.Int) {
		t.Error("Expected false, got true")
	}
}

func TestIsListTrue(t *testing.T) {
	if !IsList("[]int") {
		t.Error("Expected true, got false")
	}
}

func TestIsListFalse(t *testing.T) {
	if IsList("int") {
		t.Error("Expected false, got true")
	}
}

func TestListLen(t *testing.T) {
	t1 := &List{[]Type{Int(1), Int(2), Int(3)}, parser.Int}
	expected := 3
	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestListGetSize(t *testing.T) {
	t1 := &List{[]Type{Int(1), Int(2)}, parser.Int}
	expected := utils.Sizeof(t1)
	result := t1.GetSize()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test List errors

func TestListSetValueWithNonList(t *testing.T) {
	t1, _ := NewList(parser.Int)
	err := t1.SetValue(Int(0))

	if err == nil {
		t.Error("Expected error when setting list with non list value")
	}
}

func TestListSetValueWithListOfWrongType(t *testing.T) {
	t1 := &List{[]Type{}, "[]char"}
	t2 := &List{[]Type{Int(0)}, "[]int"}
	err := t1.SetValue(t2)

	if err == nil {
		t.Error("Expected error when setting value of list with another list of a different type")
	}
}

func TestListSetValueWithSliceOfWrongTypes(t *testing.T) {
	t1 := &List{[]Type{}, "[]char"}
	var array []Type
	array = append(array, Int(0))
	err := t1.SetValue(array)

	if err == nil {
		t.Error("Expected error when setting value of list of char with slice of int")
	}
}

func TestListGetIndexOutOfRange(t *testing.T) {
	expected := Int(5)
	t1 := &List{[]Type{Int(3), expected}, parser.Int}
	_, err := t1.GetIndex(Int(42))

	if err == nil {
		t.Error("Expected error when getting index out of range")
	}
}

func TestListGetIndexWrongType(t *testing.T) {
	expected := Int(5)
	t1 := &List{[]Type{Int(3), expected}, parser.Int}
	_, err := t1.GetIndex(String("this won't work"))

	if err == nil {
		t.Error("Expected error when getting index with non int")
	}
}

func TestListSub(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Sub(Int(0))
	if result == nil {
		t.Error("Expected error when subtracting from list")
	}
}

func TestListMod(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Mod(Int(0))
	if result == nil {
		t.Error("Expected error when getting remainder of a list")
	}
}

func TestListDiv(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Div(Int(0))
	if result == nil {
		t.Error("Expected error when dividing list")
	}
}

func TestListDivEc(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.DivEc(Int(0))
	if result == nil {
		t.Error("Expected error when getting quotient of a list")
	}
}

func TestListEqWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Eq(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListEqWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Eq(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListNotEqWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.NotEq(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListNotEqWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.NotEq(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListGtWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Gt(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListGtWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Gt(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListGtEqWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.GtEq(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListGtEqWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.GtEq(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListLwWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Lw(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListLwWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Lw(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListLwEqWithNonList(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.LwEq(Int(0))
	if result == nil {
		t.Error("Expected error when comparing list to int")
	}
}

func TestListLwEqWithWrongType(t *testing.T) {
	t1, err := NewList(parser.Char)
	if err != nil {
		t.Error(err)
	}

	_, result := t1.LwEq(&List{[]Type{Int(1)}, parser.Int})
	if result == nil {
		t.Error("Expected error when comparing list of char to list of int")
	}
}

func TestListAnd(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.And(Int(0))
	if result == nil {
		t.Error("Expected error when comparing a list")
	}
}

func TestListOr(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Or(Int(0))
	if result == nil {
		t.Error("Expected error when comparing a list")
	}
}

func TestListXor(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Xor(Int(0))
	if result == nil {
		t.Error("Expected error when comparing a list")
	}
}

func TestListNot(t *testing.T) {
	t1, err := NewList("test")
	if err != nil {
		t.Error(err)
	}

	_, result := t1.Not()
	if result == nil {
		t.Error("Expected error when comparing a list")
	}
}

func TestListAddNonList(t *testing.T) {
	t1 := &List{[]Type{Int(0)}, parser.Int}
	_, err := t1.Add(Int(1))

	if err == nil {
		t.Error("Expected error when adding an int to a list")
	}
}

func TestListAddWrongType(t *testing.T) {
	t1 := &List{[]Type{Int(0)}, parser.Int}
	t2 := &List{[]Type{Char('c')}, parser.Char}
	_, err := t1.Add(t2)

	if err == nil {
		t.Error("Expected error when adding a list of char to a list of int")
	}
}

func TestListMulWrongType(t *testing.T) {
	t1 := &List{[]Type{Int(0)}, parser.Int}
	_, err := t1.Mul(String("test"))

	if err == nil {
		t.Error("Expected error when multiplying a list by a string")
	}
}

func TestListAppendNonList(t *testing.T) {
	t1 := &List{[]Type{Int(0)}, parser.Int}
	_, err := t1.Append(Int(1))

	if err == nil {
		t.Error("Expected error when appending an int to a list")
	}
}
