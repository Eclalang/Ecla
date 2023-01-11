package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddFloats(t *testing.T) {
	t1 := eclaType.NewFloat("1.1")
	t2 := eclaType.NewFloat("0.2")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(1.3) {
		t.Error("Expected 1.3, got ", result)
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
