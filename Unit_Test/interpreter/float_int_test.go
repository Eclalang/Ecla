package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

/* func TestAddFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(5.14) {
		t.Error("Expected 5.14, got ", result)
	}
} */

/* func TestSubFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(2)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(1.14) {
		t.Error("Expected 1.14, got ", result)
	}
} */

func TestMulFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(6.28) {
		t.Error("Expected 6.28, got ", result)
	}
}

func TestDivFloatInt(t *testing.T) {
	t1 := eclaType.Float(4.5)
	t2 := eclaType.Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(2.25) {
		t.Error("Expected 2.25, got ", result)
	}
}

func TestEqFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.000)
	t2 := eclaType.Int(3)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(3)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(3)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(3)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(3)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatInt(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.Int(3)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
