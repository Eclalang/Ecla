package eclaType

import (
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"testing"
)

// Float interacts with Float

func TestNewFloat(t *testing.T) {
	t1 := NewFloat("0.3")

	if t1 != 0.3 {
		t.Error("Error when creating a Float")
	}
}

func TestFloatGetValue(t *testing.T) {
	t1 := Float(2.5)

	result := t1.GetValue()
	if result != Float(2.5) {
		t.Error("Expected 2.5, got ", result)
	}
}

func TestFloatGetType(t *testing.T) {
	t1 := Float(0)

	result := t1.GetType()
	if result != "float" {
		t.Error("Expected float, got ", result)
	}
}

func TestFloatIsNull(t *testing.T) {
	t1 := Float(0)

	result := t1.IsNull()
	if result != false {
		t.Error("Expected false, got", result)
	}
}

func TestAddFloat(t *testing.T) {
	t1 := Float(1.2)
	t2 := Float(2.3)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.5) {
		t.Error("Expected 3.5, got ", result)
	}
}

func TestAddNegFloat(t *testing.T) {
	t1 := Float(1.5)
	t2 := Float(-2.3)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-0.8) {
		t.Error("Expected -0.8, got ", result)
	}
}

func TestSubFloat(t *testing.T) {
	t1 := Float(4.6)
	t2 := Float(3.1)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.5) {
		t.Error("Expected 2.5, got ", result)
	}
}

func TestNegSubFloat(t *testing.T) {
	t1 := Float(4.1)
	t2 := Float(-3.5)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7.6) {
		t.Error("Expected 7.6, got ", result)
	}
}

func TestMulFloat(t *testing.T) {
	t1 := Float(1.2)
	t2 := Float(2.6)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.12) {
		t.Error("Expected 3.12, got ", result)
	}
}

func TestMulNegFloat(t *testing.T) {
	t1 := Float(1.2)
	t2 := Float(-2.6)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-3.12) {
		t.Error("Expected -3.12, got ", result)
	}
}

func TestDivFloat(t *testing.T) {
	t1 := Float(3.12)
	t2 := Float(1.2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.6) {
		t.Error("Expected 2.6, got ", result)
	}
}

func TestDivNegFloat(t *testing.T) {
	t1 := Float(4.5)
	t2 := Float(-2.0)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.25) {
		t.Error("Expected -2.25, got ", result)
	}
}

func TestEqFloat(t *testing.T) {
	t1 := Float(1.3)
	t2 := Float(1.3)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFalse(t *testing.T) {
	t1 := Float(1.4)
	t2 := Float(2.6)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloat(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(2.3)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatFalse(t *testing.T) {
	t1 := Float(1.1)
	t2 := Float(1.1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatTrue(t *testing.T) {
	t1 := Float(2.4)
	t2 := Float(2.3)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatFalse(t *testing.T) {
	t1 := Float(2.1)
	t2 := Float(2.3)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatEq(t *testing.T) {
	t1 := Float(4.5)
	t2 := Float(4.5)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloat(t *testing.T) {
	t1 := Float(2.3)
	t2 := Float(0.6)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatEq(t *testing.T) {
	t1 := Float(0.2)
	t2 := Float(0.2)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatFalse(t *testing.T) {
	t1 := Float(1.2)
	t2 := Float(4.2)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloat(t *testing.T) {
	t1 := Float(1.2)
	t2 := Float(4.4)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatFalse(t *testing.T) {
	t1 := Float(3.0)
	t2 := Float(0.3)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatEq(t *testing.T) {
	t1 := Float(5.1)
	t2 := Float(5.1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloat(t *testing.T) {
	t1 := Float(2.1)
	t2 := Float(3.2)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatFalse(t *testing.T) {
	t1 := Float(2.4)
	t2 := Float(1.3)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatEq(t *testing.T) {
	t1 := Float(2.0)
	t2 := Float(2.0)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloat(t *testing.T) {
	t1 := Float(1.4)
	t2 := Float(2.3)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatFalseRight(t *testing.T) {
	t1 := Float(1.7)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1.8)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloat(t *testing.T) {
	t1 := Float(1.8)
	t2 := Float(2.6)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFalseRight(t *testing.T) {
	t1 := Float(1.3)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1.4)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloat(t *testing.T) {
	t1 := Float(0.1)
	t2 := Float(1.1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0.9)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatFalseRight(t *testing.T) {
	t1 := Float(1.5)
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotFloatFalse(t *testing.T) {
	t1 := Float(1.4)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Float(0) {
		t.Error("Expected false, got", result)
	}
}

func TestNotFloatTrue(t *testing.T) {
	t1 := Float(0)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Float(1) {
		t.Error("Expected true, got", result)
	}
}

func TestFloatGetSize(t *testing.T) {
	t1 := Float(0)
	expected := utils.Sizeof(t1)

	result := t1.GetSize()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Float interacts with Char

func TestAddFloatChar(t *testing.T) {
	t1 := Float(1)
	t2 := Char('A')

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(66) {
		t.Error("Expected 66, got ", result)
	}
}

func TestSubFloatChar(t *testing.T) {
	t1 := Float(90)
	t2 := Char('A')

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(25) {
		t.Error("Expected 25, got ", result)
	}
}

func TestMulFloatChar(t *testing.T) {
	t1 := Float(2)
	t2 := Char('!')

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(66) {
		t.Error("Expected 66, got ", result)
	}
}

func TestDivFloatChar(t *testing.T) {
	t1 := Float(66)
	t2 := Char('!')

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestEqFloatChar(t *testing.T) {
	t1 := Float(65)
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatCharFalse(t *testing.T) {
	t1 := Float(66)
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatChar(t *testing.T) {
	t1 := Float(32)
	t2 := Char('!')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatCharFalse(t *testing.T) {
	t1 := Float(33)
	t2 := Char('!')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatCharTrue(t *testing.T) {
	t1 := Float(67)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatCharFalse(t *testing.T) {
	t1 := Float(64)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatCharEq(t *testing.T) {
	t1 := Float(65)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatChar(t *testing.T) {
	t1 := Float(66)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatCharEq(t *testing.T) {
	t1 := Float(66)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatCharFalse(t *testing.T) {
	t1 := Float(64)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatChar(t *testing.T) {
	t1 := Float(64)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatCharFalse(t *testing.T) {
	t1 := Float(67)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatCharEq(t *testing.T) {
	t1 := Float(65)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatChar(t *testing.T) {
	t1 := Float(64)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatCharFalse(t *testing.T) {
	t1 := Float(66)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatCharEq(t *testing.T) {
	t1 := Float(65)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatChar(t *testing.T) {
	t1 := Float(2.0)
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatCharFalseRight(t *testing.T) {
	t1 := Float(3.1)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatCharFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatCharFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatChar(t *testing.T) {
	t1 := Float(-0.4)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatCharFalseRight(t *testing.T) {
	t1 := Float(8.6)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatCharFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatCharFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatChar(t *testing.T) {
	t1 := Float(2.4)
	t2 := Char(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatCharFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Char('A')

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatCharFalseRight(t *testing.T) {
	t1 := Float(1)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatCharFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
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

	Newexpect1 := fmt.Sprintf("%g", 3.14)
	expect2 := "hello"
	expect := Newexpect1 + expect2
	expected := String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Float errors

func TestDivEcFloatErr(t *testing.T) {
	t1 := Float(4.12)
	t2 := Float(2)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when getting quotient of a Float")
	}
}

func TestNewFloatErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic when creating a new float with wrong argument, got nil")
		}
	}()

	NewFloat("not a float")
}

func TestFloatSetValueErr(t *testing.T) {
	t1 := Float(3.14)

	err := t1.SetValue(1.2)
	if err == nil {
		t.Error("Expected error when setting value of float, got nil")
	}
}

func TestFloatGetIndexErr(t *testing.T) {
	t1 := Float(3.14)

	_, err := t1.GetIndex(Float(1.2))
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestFloatAddErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding a bool and a float, got nil")
	}
}

func TestFloatSubErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting a bool from a float, got nil")
	}
}

func TestModFloatErr(t *testing.T) {
	t1 := Float(4.12)
	t2 := Float(2)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when getting remainder of a Float")
	}
}

func TestFloatMulErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying a bool and a float, got nil")
	}
}

func TestFloatDivErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing a float by a bool, got nil")
	}
}

func TestFloatEqErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatNotEqErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatGtErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatGtEqErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatLwErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatLwEqErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when comparing a bool and a float, got nil")
	}
}

func TestFloatAppendErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := Float(1.0)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending to a float, got nil")
	}
}

func TestFloatLengthErr(t *testing.T) {
	t1 := Float(3.14)

	_, err := t1.Len()
	if err == nil {
		t.Error("Expected error when getting length of float, got nil")
	}
}

func TestFloatAndErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := String("test")

	_, err := t1.And(t2)
	if err == nil {
		t.Error("Expected error when comparing a string and a float, got nil")
	}
}

func TestFloatOrErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := String("test")

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("Expected error when comparing a string and a float, got nil")
	}
}

func TestFloatXorErr(t *testing.T) {
	t1 := Float(3.14)
	t2 := String("test")

	_, err := t1.Xor(t2)
	if err == nil {
		t.Error("Expected error when comparing a string and a float, got nil")
	}
}

// Float interacts with Int

func TestAddFloatInt(t *testing.T) {
	t1 := Float(1.3)
	t2 := Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.3) {
		t.Error("Expected 3.3, got ", result)
	}
}

func TestAddNegFloatInt(t *testing.T) {
	t1 := Float(1.5)
	t2 := Int(-2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-0.5) {
		t.Error("Expected -0.5, got ", result)
	}
}

func TestSubFloatInt(t *testing.T) {
	t1 := Float(4.2)
	t2 := Int(3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.2) {
		t.Error("Expected 1.2, got ", result)
	}
}

func TestNegSubFloatInt(t *testing.T) {
	t1 := Float(4.1)
	t2 := Int(-3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7.1) {
		t.Error("Expected 7.1, got ", result)
	}
}

func TestMulFloatInt(t *testing.T) {
	t1 := Float(4.4)
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8.8) {
		t.Error("Expected 8.8, got ", result)
	}
}

func TestMulNegFloatInt(t *testing.T) {
	t1 := Float(4.4)
	t2 := Int(-2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8.8) {
		t.Error("Expected -8.8, got ", result)
	}
}

func TestDivFloatInt(t *testing.T) {
	t1 := Float(4.2)
	t2 := Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.1) {
		t.Error("Expected 2.1, got ", result)
	}
}

func TestDivNegFloatInt(t *testing.T) {
	t1 := Float(4.2)
	t2 := Int(-2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.1) {
		t.Error("Expected -2.1, got ", result)
	}
}

func TestAndFloatInt(t *testing.T) {
	t1 := Float(2.4)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatIntFalseRight(t *testing.T) {
	t1 := Float(1.6)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatIntFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatIntFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatInt(t *testing.T) {
	t1 := Float(1.4)
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatIntFalseRight(t *testing.T) {
	t1 := Float(1.8)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatIntFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Int(4)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatIntFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatInt(t *testing.T) {
	t1 := Float(2.9)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatIntFalseRight(t *testing.T) {
	t1 := Float(1.7)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatIntFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Int(7)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatIntFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatInt(t *testing.T) {
	t1 := Float(4.0)
	t2 := Int(4)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatIntFalse(t *testing.T) {
	t1 := Float(1.1)
	t2 := Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatInt(t *testing.T) {
	t1 := Float(0.9)
	t2 := Int(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatIntFalse(t *testing.T) {
	t1 := Float(1.0)
	t2 := Int(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatInt(t *testing.T) {
	t1 := Float(1.1)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatIntEq(t *testing.T) {
	t1 := Float(1.0)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatIntFalse(t *testing.T) {
	t1 := Float(0.4)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatInt(t *testing.T) {
	t1 := Float(1.2)
	t2 := Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatIntEq(t *testing.T) {
	t1 := Float(1.0)
	t2 := Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatIntFalse(t *testing.T) {
	t1 := Float(0.3)
	t2 := Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatInt(t *testing.T) {
	t1 := Float(1.4)
	t2 := Int(2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatIntEq(t *testing.T) {
	t1 := Float(1)
	t2 := Int(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatIntFalse(t *testing.T) {
	t1 := Float(2.4)
	t2 := Int(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatInt(t *testing.T) {
	t1 := Float(1)
	t2 := Int(2)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatIntEq(t *testing.T) {
	t1 := Float(1)
	t2 := Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected tue, got ", result)
	}
}

func TestLwEqFloatIntFalse(t *testing.T) {
	t1 := Float(1.7)
	t2 := Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Float interacts with Bool

func TestAndFloatBool(t *testing.T) {
	t1 := Float(4.1)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatBoolFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatBoolFalseRight(t *testing.T) {
	t1 := Float(3.1)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatBoolFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatBool(t *testing.T) {
	t1 := Float(5.7)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatBoolFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatBoolFalseRight(t *testing.T) {
	t1 := Float(4.5)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatBoolFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatBool(t *testing.T) {
	t1 := Float(2.8)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatBoolFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatBoolFalseRight(t *testing.T) {
	t1 := Float(7.1)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatBoolFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Int interacts with Var

func TestAddFloatVar(t *testing.T) {
	t1 := Float(1.2)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.3) {
		t.Error("Expected 3.3, got ", result)
	}
}

func TestAddNegFloatVar(t *testing.T) {
	t1 := Float(4.7)
	t2, _ := NewVar("testVar", "float", Float(-2.2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.5) {
		t.Error("Expected -2.5, got ", result)
	}
}

func TestSubFloatVar(t *testing.T) {
	t1 := Float(1.1)
	t2, _ := NewVar("testVar", "float", Float(0.2))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.9) {
		t.Error("Expected 0.9, got ", result)
	}
}

func TestNegSubFloatVar(t *testing.T) {
	t1 := Float(4.2)
	t2, _ := NewVar("testVar", "float", Float(-0.9))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(5.1) {
		t.Error("Expected 5.1, got ", result)
	}
}

func TestMulFloatVar(t *testing.T) {
	t1 := Float(5.0)
	t2, _ := NewVar("testVar", "float", Float(0.1))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(0.5) {
		t.Error("Expected 0.5, got ", result)
	}
}

func TestMulNegFloatVar(t *testing.T) {
	t1 := Float(5.0)
	t2, _ := NewVar("testVar", "float", Float(-0.1))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-0.5) {
		t.Error("Expected -0.5, got ", result)
	}
}

func TestDivFloatVar(t *testing.T) {
	t1 := Float(3.6)
	t2, _ := NewVar("testVar", "int", Int(3))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.2) {
		t.Error("Expected 1.2, got ", result)
	}
}

func TestDivNegFloatVar(t *testing.T) {
	t1 := Float(3.6)
	t2, _ := NewVar("testVar", "int", Int(-3))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.2) {
		t.Error("Expected -1.2, got ", result)
	}
}

func TestEqFloatVar(t *testing.T) {
	t1 := Float(1.2)
	t2, _ := NewVar("testVar", "float", Float(1.2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFalseVar(t *testing.T) {
	t1 := Float(1.2)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatVar(t *testing.T) {
	t1 := Float(1.4)
	t2, _ := NewVar("testVar", "float", Float(4.2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatVarFalse(t *testing.T) {
	t1 := Float(1.4)
	t2, _ := NewVar("testVar", "float", Float(1.4))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatVarTrue(t *testing.T) {
	t1 := Float(2.1)
	t2, _ := NewVar("testVar", "float", Float(2.01))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatVarFalse(t *testing.T) {
	t1 := Float(1.1)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatVarEq(t *testing.T) {
	t1 := Float(1.4)
	t2, _ := NewVar("testVar", "float", Float(1.4))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatVar(t *testing.T) {
	t1 := Float(1.1)
	t2, _ := NewVar("testVar", "float", Float(0.2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatVarEq(t *testing.T) {
	t1 := Float(1.2)
	t2, _ := NewVar("testVar", "float", Float(1.2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatVarFalse(t *testing.T) {
	t1 := Float(1.2)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatVar(t *testing.T) {
	t1 := Float(1.5)
	t2, _ := NewVar("testVar", "float", Float(2.5))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatVarFalse(t *testing.T) {
	t1 := Float(2.4)
	t2, _ := NewVar("testVar", "float", Float(1.4))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatVarEq(t *testing.T) {
	t1 := Float(2.7)
	t2, _ := NewVar("testVar", "float", Float(2.7))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatVar(t *testing.T) {
	t1 := Float(2.2)
	t2, _ := NewVar("testVar", "float", Float(3.2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatVarFalse(t *testing.T) {
	t1 := Float(2.4)
	t2, _ := NewVar("testVar", "float", Float(1.4))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatVarEq(t *testing.T) {
	t1 := Float(2.1)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatVar(t *testing.T) {
	t1 := Float(1.7)
	t2, _ := NewVar("testVar", "float", Float(7.2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatVarFalseRight(t *testing.T) {
	t1 := Float(4.1)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatVarFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(2.7))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatVarFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatVar(t *testing.T) {
	t1 := Float(1.1)
	t2, _ := NewVar("testVar", "float", Float(2.1))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatVarFalseRight(t *testing.T) {
	t1 := Float(1.4)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatVarFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(2.7))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatVarFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatVar(t *testing.T) {
	t1 := Float(1.5)
	t2, _ := NewVar("testVar", "float", Float(2.5))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatVarFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(2.7))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatVarFalseRight(t *testing.T) {
	t1 := Float(1.3)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatVarFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2, _ := NewVar("testVar", "float", Float(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Float interacts with Any

func TestAddFloatAny(t *testing.T) {
	t1 := Float(1.4)
	t2 := NewAny(Float(2.2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.6) {
		t.Error("Expected 3.6, got ", result)
	}
}

func TestAddNegFloatAny(t *testing.T) {
	t1 := Float(1.1)
	t2 := NewAny(Float(-2.2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.1) {
		t.Error("Expected -1.1, got ", result)
	}
}

func TestSubFloatAny(t *testing.T) {
	t1 := Float(4.1)
	t2 := NewAny(Float(3.0))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.1) {
		t.Error("Expected 1.1, got ", result)
	}
}

func TestNegSubFloatAny(t *testing.T) {
	t1 := Float(4.1)
	t2 := NewAny(Float(-3.0))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7.1) {
		t.Error("Expected 7.1, got ", result)
	}
}

func TestMulFloatAny(t *testing.T) {
	t1 := Float(4.1)
	t2 := NewAny(Int(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8.2) {
		t.Error("Expected 8.2, got ", result)
	}
}

func TestMulNegFloatAny(t *testing.T) {
	t1 := Float(4.1)
	t2 := NewAny(Int(-2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8.2) {
		t.Error("Expected -8.2, got ", result)
	}
}

func TestDivFloatAny(t *testing.T) {
	t1 := Float(4.2)
	t2 := NewAny(Int(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.1) {
		t.Error("Expected 2.1, got ", result)
	}
}

func TestDivNegFloatAny(t *testing.T) {
	t1 := Float(4.2)
	t2 := NewAny(Int(-2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.1) {
		t.Error("Expected -2.1, got ", result)
	}
}

func TestEqFloatAny(t *testing.T) {
	t1 := Float(1.1)
	t2 := NewAny(Float(1.1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFalseAny(t *testing.T) {
	t1 := Float(1.2)
	t2 := NewAny(Float(2.2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatAny(t *testing.T) {
	t1 := Float(1.3)
	t2 := NewAny(Float(2.3))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatAnyFalse(t *testing.T) {
	t1 := Float(1.4)
	t2 := NewAny(Float(1.4))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatAnyTrue(t *testing.T) {
	t1 := Float(2.5)
	t2 := NewAny(Float(1.5))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatAnyFalse(t *testing.T) {
	t1 := Float(1.6)
	t2 := NewAny(Float(2.6))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatAnyEq(t *testing.T) {
	t1 := Float(1.7)
	t2 := NewAny(Float(1.7))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatAny(t *testing.T) {
	t1 := Float(1.8)
	t2 := NewAny(Float(0.8))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatAnyEq(t *testing.T) {
	t1 := Float(1.9)
	t2 := NewAny(Float(1.9))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatAnyFalse(t *testing.T) {
	t1 := Float(1.1)
	t2 := NewAny(Float(2.1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatAny(t *testing.T) {
	t1 := Float(1.2)
	t2 := NewAny(Float(2.2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatAnyFalse(t *testing.T) {
	t1 := Float(2.3)
	t2 := NewAny(Float(1.3))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatAnyEq(t *testing.T) {
	t1 := Float(2.4)
	t2 := NewAny(Float(2.4))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatAny(t *testing.T) {
	t1 := Float(2.5)
	t2 := NewAny(Float(3.5))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatAnyFalse(t *testing.T) {
	t1 := Float(2.6)
	t2 := NewAny(Float(1.6))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatAnyEq(t *testing.T) {
	t1 := Float(2.7)
	t2 := NewAny(Float(2.7))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatAny(t *testing.T) {
	t1 := Float(1.8)
	t2 := NewAny(Float(2.8))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1.9)
	t2 := NewAny(Float(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatAnyFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(2.1))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndFloatAnyFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatAny(t *testing.T) {
	t1 := Float(1.2)
	t2 := NewAny(Float(2.2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1.3)
	t2 := NewAny(Float(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatAnyFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(2.4))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatAnyFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatAny(t *testing.T) {
	t1 := Float(1.5)
	t2 := NewAny(Float(2.5))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatAnyFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(2.6))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1.7)
	t2 := NewAny(Float(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatAnyFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := NewAny(Float(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
