package eclaType

import (
	"testing"
)

// Bool interacts with Bool

func TestEqBools(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(true)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBools(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(false)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndBools(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBools(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotBools(t *testing.T) {
	t1 := Bool(true)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Bool interacts with Float

func TestEqBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(1.0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(1.0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(1.0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(0.0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// Bool interacts with Int

func TestEqBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// Bool interacts with Char

func TestEqBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// Bool interacts with String

func TestAddBoolString(t *testing.T) {
	t1, err := NewBool("true")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t2, _ := NewString("Hello")
	actual, err := t1.Add(t2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	expected, _ := NewString("trueHello")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
