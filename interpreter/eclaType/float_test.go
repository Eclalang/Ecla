package eclaType

import (
	"fmt"
	"strconv"
	"testing"
)

// Float interacts with Float

func TestAddFloats(t *testing.T) {
	t1 := NewFloat("1.1")
	t2 := NewFloat("0.2")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	var expect1, _ = strconv.ParseFloat("1.1", 32)
	var expect2, _ = strconv.ParseFloat("0.2", 32)
	expect := expect1 + expect2
	expected := Float(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSubFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.1) {
		t.Error("Expected -1.1, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestMulFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.42) {
		t.Error("Expected 2.42, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestDivFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.5) {
		t.Error("Expected 0.5, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}

func TestEqFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(1.1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloats(t *testing.T) {
	t1 := Float(3.1)
	t2 := Float(2.2)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloats(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloats(t *testing.T) {
	t1 := Float(2.2)
	t2 := Float(2.2)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// Float interacts with Int

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
	t1 := Float(3.14)
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(6.28) {
		t.Error("Expected 6.28, got ", result)
	}
}

func TestDivFloatInt(t *testing.T) {
	t1 := Float(4.5)
	t2 := Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.25) {
		t.Error("Expected 2.25, got ", result)
	}
}

func TestEqFloatInt(t *testing.T) {
	t1 := Float(3.000)
	t2 := Int(3)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatInt(t *testing.T) {
	t1 := Float(3.14)
	t2 := Int(3)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatInt(t *testing.T) {
	t1 := Float(3.14)
	t2 := Int(3)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatInt(t *testing.T) {
	t1 := Float(3.14)
	t2 := Int(3)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatInt(t *testing.T) {
	t1 := Float(3.14)
	t2 := Int(3)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatInt(t *testing.T) {
	t1 := Float(3.14)
	t2 := Int(3)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Float interacts with String

func TestAddFloatString(t *testing.T) {
	t1 := Float(3.14)
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	expect1, _ := strconv.ParseFloat("3.14", 32)
	Newexpect1 := fmt.Sprintf("%6f", expect1)
	expect2 := "hello"
	expect := Newexpect1 + expect2
	expected := String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
