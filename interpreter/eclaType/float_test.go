package eclaType

import (
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

/*
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

func TestModFloatChar(t *testing.T) {
	t1 := Float(60)
	t2 := Char('!')

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(27) {
		t.Error("Expected 27, got ", result)
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

func TestDivEcFloatChar(t *testing.T) {
	t1 := Float(67)
	t2 := Char('!')

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2) {
		t.Error("Expected 2, got ", result)
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

func TestLwEqFloatFloatFalse(t *testing.T) {
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

func TestLwEqFloatFloatEq(t *testing.T) {
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t2 := Char(1)

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

// Int interacts with String

func TestAddFloatString(t *testing.T) {
	t1 := Float(0)
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	Newexpect1 := fmt.Sprintf("%g", 3.14)
	expect2 := "hello"
	expect := Newexpect1 + expect2
	expected := String(expect)
	fmt.Println(expected)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDivEcFloatErr(t *testing.T) {
	t1 := Float(1)
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Float by a bool")
	}
}

func TestDivEcBy0FloatErr(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Float by 0")
	}
}

func TestDivEcBy0FloatCharErr(t *testing.T) {
	t1 := Float(1)
	t2 := Char(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Float by 0")
	}
}

func TestEqFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when testing equality between Float and bool")
	}
}

func TestNotEqFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when testing inequality between Float and bool")
	}
}

func TestGtFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when testing if a Float is greater than a bool")
	}
}

func TestGtEqFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Float is greater or equal to a bool")
	}
}

func TestLwFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when testing if a Float is lower than a bool")
	}
}

func TestLwEqFloatErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Float is lower or equal to a bool")
	}
}

func TestFloatAppendTypeErr(t *testing.T) {
	t1 := Float('A')
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending a bool to a Float")
	}
}

// Int interacts with float

func TestAddFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.0) {
		t.Error("Expected 3.0, got ", result)
	}
}

func TestAddNegFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(-2.0)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.0) {
		t.Error("Expected -1.0, got ", result)
	}
}

func TestSubFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(3.0)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.0) {
		t.Error("Expected 1.0, got ", result)
	}
}

func TestNegSubFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(-3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7.0) {
		t.Error("Expected 7.0, got ", result)
	}
}

func TestMulFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(2.0)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8.0) {
		t.Error("Expected 8.0, got ", result)
	}
}

func TestMulNegFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(-2.0)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8.0) {
		t.Error("Expected -8.0, got ", result)
	}
}

func TestDivFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(2.0)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivNegFloatFloat(t *testing.T) {
	t1 := Float(4)
	t2 := Float(-2.0)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.0) {
		t.Error("Expected -2.0, got ", result)
	}
}

func TestAndFloatFloat(t *testing.T) {
	t1 := Float(2)
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatFloatFalseRight(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatFloatFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFloatFalseRight(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatFloatFalseBoth(t *testing.T) {
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

func TestXorFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatFloatFalseRight(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatFloatFalseLeft(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatFloatFalseBoth(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFloat(t *testing.T) {
	t1 := Float(0)
	t2 := Float(0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFloatFalse(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatFloat(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatFloatFalse(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0.5)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatFloatEq(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatFloatFalse(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(0.5)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatFloatEq(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatFloatFalse(t *testing.T) {
	t1 := Float(0)
	t2 := Float(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatFloatEq(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatFloatFalse(t *testing.T) {
	t1 := Float(2)
	t2 := Float(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatFloat(t *testing.T) {
	t1 := Float(1)
	t2 := Float(2.5)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatFloatEq(t *testing.T) {
	t1 := Float(1)
	t2 := Float(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected tue, got ", result)
	}
}

func TestLwEqFloatFloatFalse(t *testing.T) {
	t1 := Float(2)
	t2 := Float(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Int interacts with Bool

func TestAndFloatBool(t *testing.T) {
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
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
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestAddNegFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(-2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestSubFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestNegSubFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(-3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7) {
		t.Error("Expected 7, got ", result)
	}
}

func TestModFloatVar(t *testing.T) {
	t1 := Float(3)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8) {
		t.Error("Expected 8, got ", result)
	}
}

func TestMulNegFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(-2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8) {
		t.Error("Expected -8, got ", result)
	}
}

func TestDivFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivNegFloatVar(t *testing.T) {
	t1 := Float(4)
	t2, _ := NewVar("testVar", "Float", Float(-2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.0) {
		t.Error("Expected -2.0, got ", result)
	}
}

func TestDivEcFloatVar(t *testing.T) {
	t1 := Float(5)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegFloatVar(t *testing.T) {
	t1 := Float(5)
	t2, _ := NewVar("testVar", "Float", Float(-2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestEqFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFalseVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatVarFalse(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatVarTrue(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatVarFalse(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatVarEq(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(0))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatVarEq(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatVarFalse(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatVarFalse(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatVarEq(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatVar(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(3))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatVarFalse(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(1))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatVarEq(t *testing.T) {
	t1 := Float(2)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatVarFalseRight(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(0))

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
	t2, _ := NewVar("testVar", "Float", Float(2))

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
	t2, _ := NewVar("testVar", "Float", Float(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatVarFalseRight(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(0))

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
	t2, _ := NewVar("testVar", "Float", Float(2))

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
	t2, _ := NewVar("testVar", "Float", Float(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorFloatVar(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(2))

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
	t2, _ := NewVar("testVar", "Float", Float(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatVarFalseRight(t *testing.T) {
	t1 := Float(1)
	t2, _ := NewVar("testVar", "Float", Float(0))

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
	t2, _ := NewVar("testVar", "Float", Float(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Int interacts with Any

func TestAddFloatAnt(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestAddNegFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(-2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestSubFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestNegSubFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(-3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7) {
		t.Error("Expected 7, got ", result)
	}
}

func TestModFloatAny(t *testing.T) {
	t1 := Float(3)
	t2 := NewAny(Float(2))

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8) {
		t.Error("Expected 8, got ", result)
	}
}

func TestMulNegFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(-2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8) {
		t.Error("Expected -8, got ", result)
	}
}

func TestDivFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivNegFloatAny(t *testing.T) {
	t1 := Float(4)
	t2 := NewAny(Float(-2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.0) {
		t.Error("Expected -2.0, got ", result)
	}
}

func TestDivEcFloatAny(t *testing.T) {
	t1 := Float(5)
	t2 := NewAny(Float(2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegFloatAny(t *testing.T) {
	t1 := Float(5)
	t2 := NewAny(Float(-2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestEqFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqFloatFalseAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqFloatAnyFalse(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(1))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatAnyTrue(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtFloatAnyFalse(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtFloatAnyEq(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(0))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatAnyEq(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqFloatAnyFalse(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloatAnyFalse(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(1))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwFloatAnyEq(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatAny(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(3))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqFloatAnyFalse(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(1))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqFloatAnyEq(t *testing.T) {
	t1 := Float(2)
	t2 := NewAny(Float(2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatAny(t *testing.T) {
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1)
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
	t2 := NewAny(Float(2))

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
	t1 := Float(1)
	t2 := NewAny(Float(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1)
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
	t2 := NewAny(Float(2))

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
	t1 := Float(1)
	t2 := NewAny(Float(2))

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
	t2 := NewAny(Float(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorFloatAnyFalseRight(t *testing.T) {
	t1 := Float(1)
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
*/
