package eclaType

import (
	"testing"
)

// Int interacts with Float

func TestAddIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(2.2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.2) {
		t.Error("Expected 3.2, got ", result)
	}
}

func TestSubIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(2.2)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.2) {
		t.Error("Expected -1.2, got ", result)
	}
}

func TestMulIntFloat(t *testing.T) {
	t1 := Int(2)
	t2 := Float(2.2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(4.4) {
		t.Error("Expected 4.4, got ", result)
	}
}

func TestDivIntFloat(t *testing.T) {
	t1 := Int(2)
	t2 := Float(2.2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.9090909) {
		t.Error("Expected 0.9090909, got ", result)
	}
}

func TestEqIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1.0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1.1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1.1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqIntFloat(t *testing.T) {
	t1 := Int(2)
	t2 := Float(1.1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1.1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntFloat(t *testing.T) {
	t1 := Int(5)
	t2 := Float(1.1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Int interacts with Int

func TestAddIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestSubIntegers(t *testing.T) {
	t1 := Int(2)
	t2 := Int(1)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestModIntegers(t *testing.T) {
	t1 := Int(5)
	t2 := Int(2)

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulIntegers(t *testing.T) {
	t1 := Int(2)
	t2 := Int(3)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(6) {
		t.Error("Expected 6, got ", result)
	}
}

func TestDivIntegers(t *testing.T) {
	t1 := Int(6)
	t2 := Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestDivEcIntegers(t *testing.T) {
	t1 := Int(6)
	t2 := Int(2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestEqIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntegers(t *testing.T) {
	t1 := Int(3)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(3)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntegers(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// Int interacts with String

func TestAddIntString(t *testing.T) {
	t1 := Int(1)
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("1hello") {
		t.Error("Expected 1hello, got ", result)
	}
}

func TestMulIntString(t *testing.T) {
	t1 := Int(2)
	t2 := String("hello")

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("hellohello") {
		t.Error("Expected hellohello, got ", result)
	}
}
