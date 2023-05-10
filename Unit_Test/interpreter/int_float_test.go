package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(3.2) {
		t.Error("Expected 3.2, got ", result)
	}
}

func TestSubIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(-1.2) {
		t.Error("Expected -1.2, got ", result)
	}
}

func TestMulIntFloat(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.Float(2.2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(4.4) {
		t.Error("Expected 4.4, got ", result)
	}
}

func TestDivIntFloat(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.Float(2.2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(0.9090909) {
		t.Error("Expected 0.9090909, got ", result)
	}
}

func TestEqIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(1.0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(1.1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(1.1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqIntFloat(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.Float(1.1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntFloat(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Float(1.1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntFloat(t *testing.T) {
	t1 := eclaType.Int(5)
	t2 := eclaType.Float(1.1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
