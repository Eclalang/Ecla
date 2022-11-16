package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestSubIntegers(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.Int(1)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestModIntegers(t *testing.T) {
	t1 := eclaType.Int(5)
	t2 := eclaType.Int(2)

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulIntegers(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.Int(3)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(6) {
		t.Error("Expected 6, got ", result)
	}
}

func TestDivIntegers(t *testing.T) {
	t1 := eclaType.Int(6)
	t2 := eclaType.Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestDivEcIntegers(t *testing.T) {
	t1 := eclaType.Int(6)
	t2 := eclaType.Int(2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestEqIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntegers(t *testing.T) {
	t1 := eclaType.Int(3)
	t2 := eclaType.Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(3)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
