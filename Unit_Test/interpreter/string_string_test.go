package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String(" world")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("hello world") {
		t.Error("Expected true, got ", result)
	}
}

func TestMulStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.Int(3)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("hellohellohello") {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String("hello")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String("world")

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtStrings(t *testing.T) {
	t1 := eclaType.String("hello!")
	t2 := eclaType.String("hello")

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqStrings(t *testing.T) {
	t1 := eclaType.String("hello!")
	t2 := eclaType.String("hello")

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String("hello!")

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String("hello")

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAppendStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String(" world")

	result, err := t1.Append(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("hello world") {
		t.Error("Expected hello world, got ", result)
	}
}
