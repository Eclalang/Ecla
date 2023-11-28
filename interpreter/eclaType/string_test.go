package eclaType

import (
	"fmt"
	"strconv"
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

// String interacts with Float

func TestAddStringFloat(t *testing.T) {
	t1 := String("Hello")
	t2 := Float(1.1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	expect1, _ := strconv.ParseFloat("1.1", 32)
	Newexpect1 := fmt.Sprintf("%6f", expect1)
	expect2 := "Hello"
	expect := expect2 + Newexpect1
	expected := String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// String interacts with String

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

func TestMulStrings(t *testing.T) {
	t1 := String("hello")
	t2 := Int(3)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("hellohellohello") {
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
