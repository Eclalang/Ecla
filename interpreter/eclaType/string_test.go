package eclaType

import (
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// String interacts with Bool

func TestAddStringBool(t *testing.T) {
	t1 := String("Hello")
	t2 := Bool(true)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("Hellotrue") {
		t.Error("Expected Hellotrue, got ", result)
	}
}

// String interacts with Int

func TestAddStringInt(t *testing.T) {
	t1 := String("Hello")
	t2 := Int(1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulStringInt(t *testing.T) {
	t1 := String("Hello")
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestGetIndexStringInt(t *testing.T) {
	t1 := String("123")
	t2 := Int(0)

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// String interacts with Float

func TestAddStringFloat(t *testing.T) {
	t1 := String("Hello")
	t2 := Float(1.1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	Newexpect1 := fmt.Sprintf("%g", 1.1)
	expect2 := "Hello"
	expect := expect2 + Newexpect1
	expected := String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// String interacts with Any

func TestAddStringAny(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(Int(1))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulStringAny(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestEqStringAny(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(String("Hello"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStringAnyFalse(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(String(" world"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqStringAny(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(String(" world"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqStringAnyFalse(t *testing.T) {
	t1 := String("Hello")
	t2 := NewAny(String("Hello"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtStringAnyFalse(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtStringAnyEq(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("123"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("12"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqStringAnyFalse(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("1234"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqStringAnyEq(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("123"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStringAnyFalse(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwStringAnyEq(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("123"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("1234"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqStringAnyFalse(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("12"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqStringAnyEq(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("123"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAppendStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(String("4"))

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result != String("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

func TestGetIndexStringAny(t *testing.T) {
	t1 := String("123")
	t2 := NewAny(Int(0))

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// String interacts with Var

func TestAddStringVar(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.Int, Int(1))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulStringVar(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.Int, Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestEqStringVar(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.String, String("Hello"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStringVarFalse(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.String, String(" world"))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqStringVar(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.String, String(" world"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqStringVarFalse(t *testing.T) {
	t1 := String("Hello")
	t2, _ := NewVar("test", parser.String, String("Hello"))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtStringVar(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtStringVarFalse(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtStringVarEq(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("123"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqStringVar(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("12"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqStringVarFalse(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("1234"))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqStringVarEq(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("123"))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStringVar(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStringVarFalse(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwStringVarEq(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("123"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqStringVar(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("1234"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqStringVarFalse(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("12"))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqStringVarEq(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.String, String("123"))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGetIndexStringVar(t *testing.T) {
	t1 := String("123")
	t2, _ := NewVar("test", parser.Int, Int(0))

	result, err := t1.GetIndex(t2)
	if err != nil {
		t.Error(err)
	}
	if (*result).GetValue() != Char('1') {
		t.Error("Expected \"1\", got ", result)
	}
}

// String interacts with String

func TestStringGetSize(t *testing.T) {
	t1 := String("test")
	expected := utils.Sizeof(t1)

	result := t1.GetSize()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestStrinLen(t *testing.T) {
	t1 := String("test")
	expected := 4

	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestIsNullString(t *testing.T) {
	t1 := String("test")
	result := t1.IsNull()

	if result == true {
		t.Error("expected false, got ", result)
	}
}

func TestStringString(t *testing.T) {
	t1 := String("test")
	t2 := "test"

	result := t1.String()

	if result != t2 {
		t.Errorf("expected %s, got %s", t2, result)
	}
}

func TestAddStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String(" world")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("hello world") {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String("hello")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStringsFalse(t *testing.T) {
	t1 := String("hello")
	t2 := String(" world")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String("world")

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtStrings(t *testing.T) {
	t1 := String("hello!")
	t2 := String("hello")

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqStrings(t *testing.T) {
	t1 := String("hello!")
	t2 := String("hello")

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String("hello!")

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String("hello")

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAppendStrings(t *testing.T) {
	t1 := String("hello")
	t2 := String(" world")

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("hello world") {
		t.Error("Expected hello world, got ", result)
	}
}

// String interacts with Char

func TestAddStringChar(t *testing.T) {
	t1 := String("Hello")
	t2 := Char('A')

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("HelloA") {
		t.Error("Expected HelloA, got ", result)
	}
}

func TestMulStringChar(t *testing.T) {
	t1 := String("Hello")
	t2 := Char(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}

func TestAppendStringChar(t *testing.T) {
	t1 := String("123")
	t2 := Char('4')

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result != String("1234") {
		t.Error("Expected \"1234\", got ", result)
	}
}

// test String errors

func TestStringSetValue(t *testing.T) {
	t1, _ := NewString("test")
	result := t1.SetValue("err")

	if result == nil {
		t.Error("expected error when setting value of string")
	}
}

func TestStringNewStringEscapeErr(t *testing.T) {
	_, err := NewString("\n")

	if err == nil {
		t.Error("expected error when creating string with only escape char")
	}
}

func TestStringNewStringErr(t *testing.T) {
	_, err := NewString("'\"\"'")

	if err == nil {
		t.Error("expected error when creating invalid strings")
	}
}

func TestGetIndexStringOutOfRangeErr(t *testing.T) {
	t1 := String("123")
	t2 := Int(3)

	_, err := t1.GetIndex(t2)
	if err == nil {
		t.Error("Expected error when indexing out of range")
	}
}

func TestGetIndexStringTypeErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.GetIndex(t2)
	if err == nil {
		t.Error("Expected error when indexing string with bool")
	}
}

func TestSubStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting from string")
	}
}

func TestModStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when getting remainder of string")
	}
}

func TestMulStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying string with bool")
	}
}

func TestDivStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing string")
	}
}

func TestDivEcStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing string")
	}
}

func TestEqStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestNotEqStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestGtStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestGtEqStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestLwStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestLwEqStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestAndStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.And(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestOrStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestXorStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Xor(t2)
	if err == nil {
		t.Error("Expected error when comparing string and bool")
	}
}

func TestNotStringErr(t *testing.T) {
	t1 := String("123")

	_, err := t1.Not()
	if err == nil {
		t.Error("Expected error when getting \"not\" of string")
	}
}

func TestAppendStringErr(t *testing.T) {
	t1 := String("123")
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending bool to string")
	}
}
