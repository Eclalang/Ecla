package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"testing"
)

// Bool interacts with Bool

func TestNewBool(t *testing.T) {
	_, err := NewBool("true")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	_, err = NewBool("not a bool")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestBoolGetSize(t *testing.T) {
	t1 := Bool(true)
	expected := utils.Sizeof(t1)

	result := t1.GetSize()
	if result != utils.Sizeof(t1) {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

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

func TestEqBoolsFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(false)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
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

func TestNotEqBoolsFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(true)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
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

func TestAndBoolsFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolsFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolsFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBools(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolsFalseLeft(t *testing.T) {
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

func TestOrBoolsFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolsFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBools(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolsFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolsFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolsFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
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

func TestEqBoolFloatFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(0.0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(0.0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolFloatFalse(t *testing.T) {
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

func TestAndBoolFloatFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(1.0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolFloatFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(0.0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolFloatFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(0.0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(1.0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolFloatFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(1.0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolFloatFalseRight(t *testing.T) {
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

func TestOrBoolFloatFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(0.0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolFloat(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(1.0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolFloatFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(1.0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolFloatFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Float(0.0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolFloatFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Float(0.0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
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

func TestEqBoolIntFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolIntFalse(t *testing.T) {
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

func TestAndBoolIntFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolIntFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolIntFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolIntFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolIntFalseRight(t *testing.T) {
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

func TestOrBoolIntFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolInt(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolIntFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolIntFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolIntFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Bool interacts with Char

func TestEqBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqBoolCharFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolCharFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := Char('A')

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
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndBoolCharFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolCharFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolCharFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolCharFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolCharFalseRight(t *testing.T) {
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

func TestOrBoolCharFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolChar(t *testing.T) {
	t1 := Bool(true)
	t2 := Char('A')

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolCharFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := Char('A')

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolCharFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolCharFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
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

func TestBoolString(t *testing.T) {
	t1, _ := NewBool("true")
	t2 := "true"
	if t1.String() != t2 {
		t.Errorf("Expected \"true\", got %s", t1.String())
	}
}

func TestBoolGetString(t *testing.T) {
	t1, _ := NewBool("false")
	t2 := String("false")
	if t1.GetString() != t2 {
		t.Errorf("Expected \"false\", got %s", t1.GetString())
	}
}

// Bool interacts with Var

func TestEqBoolVar(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqBoolVarFalse(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqBoolVar(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolVarFalse(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolVar(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndBoolVarFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolVarFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolVarFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBoolVar(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolVarFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolVarFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolVarFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolVar(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolVarFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(true))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolVarFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolVarFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2, _ := NewVar("test", "bool", Bool(false))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Bool interacts with Any

func TestEqBoolAny(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(true))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqBoolAnyFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(false))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqBoolAny(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(false))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolAnyFalse(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(true))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolAny(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(true))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndBoolAnyFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(true))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolAnyFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(false))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolAnyFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(false))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrBoolAny(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(true))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolAnyFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(true))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolAnyFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(false))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolAnyFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(false))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolAny(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(true))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorBoolAnyFalseLeft(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(true))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolAnyFalseRight(t *testing.T) {
	t1 := Bool(true)
	t2 := NewAny(Bool(false))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorBoolAnyFalseBoth(t *testing.T) {
	t1 := Bool(false)
	t2 := NewAny(Bool(false))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// test bool errors

func TestBoolSetValueErr(t *testing.T) {
	t1 := Bool(false)

	err := t1.SetValue(1)
	if err == nil {
		t.Error("Expected error when using SetValue of Bool")
	}
}

func TestBoolGetIndexErr(t *testing.T) {
	t1 := Bool(false)

	_, err := t1.GetIndex(Bool(true))
	if err == nil {
		t.Error("Expected error when using SetValue of Bool")
	}
}

func TestBoolAddErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding a bool to a bool")
	}
}

func TestBoolSubErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting a bool to a bool")
	}
}

func TestBoolModErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when getting the modulo of a bool")
	}
}

func TestBoolMulErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying a bool")
	}
}

func TestBoolDivErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing a bool")
	}
}

func TestBoolDivEcErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing ec a bool")
	}
}

func TestBoolEqErr(t *testing.T) {
	t1 := Bool(false)
	t2 := String("false")

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when comparing a string to a bool")
	}
}

func TestBoolNotEqErr(t *testing.T) {
	t1 := Bool(false)
	t2 := String("false")

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when comparing a string to a bool")
	}
}

func TestBoolGtErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when comparing size of a bool")
	}
}

func TestBoolGtEqErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when comparing size of a bool")
	}
}

func TestBoolLwErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when comparing size of a bool")
	}
}

func TestBoolLwEqErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when comparing size of a bool")
	}
}

func TestBoolAndErr(t *testing.T) {
	t1 := Bool(false)
	t2 := String("false")

	_, err := t1.And(t2)
	if err == nil {
		t.Error("Expected error when comparing a string to a bool")
	}
}

func TestBoolOrErr(t *testing.T) {
	t1 := Bool(false)
	t2 := String("false")

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("Expected error when comparing a string to a bool")
	}
}

func TestBoolXorErr(t *testing.T) {
	t1 := Bool(false)
	t2 := String("false")

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when comparing a string to a bool")
	}
}

func TestBoolAppendErr(t *testing.T) {
	t1 := Bool(false)
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending a bool")
	}
}

func TestBoolLenErr(t *testing.T) {
	t1 := Bool(false)

	_, err := t1.Len()
	if err == nil {
		t.Error("Expected error when getting length of a bool")
	}
}
