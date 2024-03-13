package eclaType

import (
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

func TestListSetValue(t *testing.T) {
	t1, _ := NewList(parser.Int)
	t2 := t1.GetValue()

	if t1 != t2 {
		t.Error("error when getting value of list")
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
