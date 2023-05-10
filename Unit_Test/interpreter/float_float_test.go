package interpreter

import (
	"strconv"
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddFloats(t *testing.T) {
	t1 := eclaType.NewFloat("1.1")
	t2 := eclaType.NewFloat("0.2")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	var expect1, _ = strconv.ParseFloat("1.1", 32)
	var expect2, _ = strconv.ParseFloat("0.2", 32)
	expect := expect1 + expect2
	expected := eclaType.Float(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSubFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(-1.1) {
		t.Error("Expected -1.1, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestMulFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(2.42) {
		t.Error("Expected 2.42, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestDivFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(0.5) {
		t.Error("Expected 0.5, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestEqFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(1.1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloats(t *testing.T) {
	t1 := eclaType.Float(3.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloats(t *testing.T) {
	t1 := eclaType.Float(2.2)
	t2 := eclaType.Float(2.2)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
