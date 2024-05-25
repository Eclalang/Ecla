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
	t1 := &List{[]Type{}, expected}
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

/*
func TestListGetSize(t *testing.T) {
	t1 := List("test")
	expected := utils.Sizeof(t1)

	result := t1.GetSize()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestStrinLen(t *testing.T) {
	t1 := List("test")
	expected := 4

	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestIsNullList(t *testing.T) {
	t1 := List("test")
	result := t1.IsNull()

	if result == true {
		t.Error("expected false, got ", result)
	}
}

func TestStringList(t *testing.T) {
	t1 := List("test")
	t2 := "test"

	result := t1.String()

	if result != t2 {
		t.Errorf("expected %s, got %s", t2, result)
	}
}

func TestAddLists(t *testing.T) {
	t1 := List("hello")
	t2 := List(" world")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("hello world") {
		t.Error("Expected true, got ", result)
	}
}

func TestEqLists(t *testing.T) {
	t1 := List("hello")
	t2 := List("hello")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqListsFalse(t *testing.T) {
	t1 := List("hello")
	t2 := List(" world")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqLists(t *testing.T) {
	t1 := List("hello")
	t2 := List("world")

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtLists(t *testing.T) {
	t1 := List("hello!")
	t2 := List("hello")

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqLists(t *testing.T) {
	t1 := List("hello!")
	t2 := List("hello")

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwLists(t *testing.T) {
	t1 := List("hello")
	t2 := List("hello!")

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqLists(t *testing.T) {
	t1 := List("hello")
	t2 := List("hello")

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAppendLists(t *testing.T) {
	t1 := List("hello")
	t2 := List(" world")

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("hello world") {
		t.Error("Expected hello world, got ", result)
	}
}

/*
// List interacts with Bool

func TestAddListBool(t *testing.T) {
	t1 := List("Hello")
	t2 := Bool(true)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("Hellotrue") {
		t.Error("Expected Hellotrue, got ", result)
	}
}

// List interacts with Int

func TestAddListInt(t *testing.T) {
	t1 := List("Hello")
	t2 := Int(1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulListInt(t *testing.T) {
	t1 := List("Hello")
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestGetIndexListInt(t *testing.T) {
	t1 := List("123")
	t2 := Int(0)

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// List interacts with Float

func TestAddListFloat(t *testing.T) {
	t1 := List("Hello")
	t2 := Float(1.1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	Newexpect1 := fmt.Sprintf("%g", 1.1)
	expect2 := "Hello"
	expect := expect2 + Newexpect1
	expected := List(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// List interacts with Any

func TestAddListAny(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(Int(1))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulListAny(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestEqListAny(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(List("Hello"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqListAnyFalse(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(List(" world"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqListAny(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(List(" world"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqListAnyFalse(t *testing.T) {
	t1 := List("Hello")
	t2 := NewAny(List("Hello"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtListAnyFalse(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtListAnyEq(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("123"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("12"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqListAnyFalse(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("1234"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqListAnyEq(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("123"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwListAnyFalse(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwListAnyEq(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("123"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("1234"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqListAnyFalse(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("12"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqListAnyEq(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("123"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAppendListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(List("4"))

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result != List("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

func TestGetIndexListAny(t *testing.T) {
	t1 := List("123")
	t2 := NewAny(Int(0))

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// List interacts with Var

func TestAddListVar(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.Int, Int(1))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulListVar(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.Int, Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestEqListVar(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.List, List("Hello"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqListVarFalse(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.List, List(" world"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqListVar(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.List, List(" world"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqListVarFalse(t *testing.T) {
	t1 := List("Hello")
	t2, _ := NewVar("test", parser.List, List("Hello"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtListVar(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtListVarFalse(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtListVarEq(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("123"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqListVar(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqListVarFalse(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqListVarEq(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("123"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwListVar(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwListVarFalse(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwListVarEq(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("123"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqListVar(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqListVarFalse(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqListVarEq(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.List, List("123"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGetIndexListVar(t *testing.T) {
	t1 := List("123")
	t2, _ := NewVar("test", parser.Int, Int(0))

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// List interacts with Char

func TestAddListChar(t *testing.T) {
	t1 := List("Hello")
	t2 := Char('A')

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("HelloA") {
		t.Error("Expected HelloA, got ", result)
	}
}

func TestMulListChar(t *testing.T) {
	t1 := List("Hello")
	t2 := Char(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != List("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestAppendListChar(t *testing.T) {
	t1 := List("123")
	t2 := Char('4')

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result != List("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

// test List errors

func TestListSetValue(t *testing.T) {
	t1, _ := NewList("test")
	result := t1.SetValue("err")

	if result == nil {
		t.Error("expected error when setting value of List")
	}
}

func TestListNewListEscapeErr(t *testing.T) {
	_, err := NewList("\n")

	if err == nil {
		t.Error("expected error when creating List with only escape char")
	}
}

func TestListNewListErr(t *testing.T) {
	_, err := NewList("'\"\"'")

	if err == nil {
		t.Error("expected error when creating invalid Lists")
	}
}

func TestGetIndexListOutOfRangeErr(t *testing.T) {
	t1 := List("123")
	t2 := Int(3)

	_, err := t1.GetIndex(t2)
	if err == nil {
		t.Error("Expected error when indexing out of range")
	}
}

func TestGetIndexListTypeErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.GetIndex(t2)
	if err == nil {
		t.Error("Expected error when indexing List with bool")
	}
}

func TestSubListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting from List")
	}
}

func TestModListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when getting remainder of List")
	}
}

func TestMulListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying List with bool")
	}
}

func TestDivListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing List")
	}
}

func TestDivEcListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing List")
	}
}

func TestEqListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestNotEqListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestGtListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestGtEqListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestLwListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestLwEqListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestAndListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.And(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestOrListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestXorListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Xor(t2)
	if err == nil {
		t.Error("Expected error when comparing List and bool")
	}
}

func TestNotListErr(t *testing.T) {
	t1 := List("123")

	_, err := t1.Not()
	if err == nil {
		t.Error("Expected error when getting \"not\" of List")
	}
}

func TestAppendListErr(t *testing.T) {
	t1 := List("123")
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending bool to List")
	}
}
*/
