package eclaType

import (
	"testing"
)

// Int interacts with Int

func TestNewInt(t *testing.T) {
	t1 := NewInt("0")

	if t1 != 0 {
		t.Error("Error when creating a Int")
	}
}

func TestIntGetValue(t *testing.T) {
	t1 := Int(0)

	result := t1.GetValue()
	if result != Int(0) {
		t.Error("Expected 0, got ", result)
	}
}

func TestIntGetType(t *testing.T) {
	t1 := Int(0)

	result := t1.GetType()
	if result != "int" {
		t.Error("Expected int, got ", result)
	}
}

func TestIntIsNull(t *testing.T) {
	t1 := Int(0)

	result := t1.IsNull()
	if result != false {
		t.Error("Expected false, got", result)
	}
}

func TestAddInt(t *testing.T) {
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

func TestAddNegInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(-2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestSubInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestNegSubInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(-3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(7) {
		t.Error("Expected 7, got ", result)
	}
}

func TestModInt(t *testing.T) {
	t1 := Int(3)
	t2 := Int(2)

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(8) {
		t.Error("Expected B, got ", result)
	}
}

func TestMulNegInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(-2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(-8) {
		t.Error("Expected -8, got ", result)
	}
}

func TestDivInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivNegInt(t *testing.T) {
	t1 := Int(4)
	t2 := Int(-2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.0) {
		t.Error("Expected -2.0, got ", result)
	}
}

func TestDivEcInt(t *testing.T) {
	t1 := Int(5)
	t2 := Int(2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegInt(t *testing.T) {
	t1 := Int(5)
	t2 := Int(-2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestEqInt(t *testing.T) {
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

func TestEqIntFalse(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntFalse(t *testing.T) {
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

func TestGtIntTrue(t *testing.T) {
	t1 := Int(2)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtIntFalse(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntEq(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntEq(t *testing.T) {
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

func TestGtEqIntFalse(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntFalse(t *testing.T) {
	t1 := Int(2)
	t2 := Int(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwIntEq(t *testing.T) {
	t1 := Int(2)
	t2 := Int(2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqInt(t *testing.T) {
	t1 := Int(2)
	t2 := Int(3)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntFalse(t *testing.T) {
	t1 := Int(2)
	t2 := Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqIntEq(t *testing.T) {
	t1 := Int(2)
	t2 := Int(2)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(2)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorInt(t *testing.T) {
	t1 := Int(1)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotIntFalse(t *testing.T) {
	t1 := Int(1)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Int(0) {
		t.Error("Expected false, got", result)
	}
}

func TestNotIntTrue(t *testing.T) {
	t1 := Int(0)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Int(1) {
		t.Error("Expected true, got", result)
	}
}

// Int interacts with Char

func TestAddIntChar(t *testing.T) {
	t1 := Int(1)
	t2 := Char('A')

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(66) {
		t.Error("Expected 66, got ", result)
	}
}

func TestSubIntChar(t *testing.T) {
	t1 := Int(90)
	t2 := Char('A')

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(25) {
		t.Error("Expected 25, got ", result)
	}
}

func TestModIntChar(t *testing.T) {
	t1 := Int(60)
	t2 := Int('!')

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(27) {
		t.Error("Expected 27, got ", result)
	}
}

func TestMulIntChar(t *testing.T) {
	t1 := Int(2)
	t2 := Char('!')

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(66) {
		t.Error("Expected 66, got ", result)
	}
}

func TestDivIntChar(t *testing.T) {
	t1 := Int(66)
	t2 := Char('!')

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivEcIntChar(t *testing.T) {
	t1 := Int(67)
	t2 := Int('!')

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Int(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestEqIntChar(t *testing.T) {
	t1 := Int(65)
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqIntCharFalse(t *testing.T) {
	t1 := Int(66)
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqIntChar(t *testing.T) {
	t1 := Int(32)
	t2 := Char('!')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntCharFalse(t *testing.T) {
	t1 := Int(33)
	t2 := Char('!')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntCharTrue(t *testing.T) {
	t1 := Int(67)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtIntCharFalse(t *testing.T) {
	t1 := Int(64)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntCharEq(t *testing.T) {
	t1 := Int(65)
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqIntChar(t *testing.T) {
	t1 := Int(66)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntCharEq(t *testing.T) {
	t1 := Int(66)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntCharFalse(t *testing.T) {
	t1 := Int(64)
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwIntChar(t *testing.T) {
	t1 := Int(64)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntCharFalse(t *testing.T) {
	t1 := Int(67)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwIntCharEq(t *testing.T) {
	t1 := Int(65)
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqIntChar(t *testing.T) {
	t1 := Int(64)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntIntFalse(t *testing.T) {
	t1 := Int(66)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqIntIntEq(t *testing.T) {
	t1 := Int(65)
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntChar(t *testing.T) {
	t1 := Int(1)
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntCharFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntCharFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Char('A')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntCharFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntChar(t *testing.T) {
	t1 := Int(1)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntCharFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntCharFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Char('A')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntCharFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorIntChar(t *testing.T) {
	t1 := Int(1)
	t2 := Char(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntCharFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Char(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntCharFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntCharFalseBoth(t *testing.T) {
	t1 := Int(0)
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

func TestAddIntString(t *testing.T) {
	t1 := Int(0)
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("0hello") {
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

func TestIntString(t *testing.T) {
	t1 := Int(0)
	t2 := "0"

	result := t1.String() == (t2)
	if result != true {
		t.Error("Expected true, got ", result)
	}
}

func TestIntGetString(t *testing.T) {
	t1 := Int(0)
	t2 := String("0")

	result, err := t1.GetString().Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

// test errors Int

func TestIntLenErr(t *testing.T) {
	t1 := Int(5)

	_, err := t1.Len()
	if err == nil {
		t.Error("Expected error when checking len of int")
	}
}

func TestDivEcIntFloatErr(t *testing.T) {
	t1 := Int(5)
	t2 := Float(2)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing ec by float")
	}
}

func TestIntAppendErr(t *testing.T) {
	t1 := Int(0)
	t2 := Int(1)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestIntSetValueErr(t *testing.T) {
	t1 := Int(0)

	err := t1.SetValue(1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestIntGetIndexErr(t *testing.T) {
	t1 := Int(0)

	_, err := t1.GetIndex(Int(0))
	if err == nil {
		t.Error("Expected error when indexing")
	}
}

func TestAddIntErr(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding a bool to an int")
	}
}

func TestSubIntErr(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting a bool to an int")
	}
}

func TestMulIntErr(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying an int by a bool")
	}
}

func TestDivIntErr(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an int by a bool")
	}
}

func TestDivBy0IntErr(t *testing.T) {
	t1 := Int(0)
	t2 := Int(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an int by 0")
	}
}

func TestDivEcBy0IntCharErr(t *testing.T) {
	t1 := Int(5)
	t2 := Char(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an int by 0")
	}
}

func TestModIntErr(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when doing a modulo of an int by a bool")
	}
}

func TestModBy0IntErr(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a Int by 0")
	}
}

func TestModBy0IntCharErr(t *testing.T) {
	t1 := Int(1)
	t2 := Char(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a Int by 0")
	}
}

func TestDivEcIntErr(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Int by a bool")
	}
}

func TestDivEcBy0IntErr(t *testing.T) {
	t1 := Int(1)
	t2 := Int(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Int by 0")
	}
}

func TestDivEcBy0IntIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Int(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Int by 0")
	}
}

func TestEqIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when testing equality between Int and bool")
	}
}

func TestNotEqIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when testing inequality between Int and bool")
	}
}

func TestGtIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when testing if a Int is greater than a bool")
	}
}

func TestGtEqIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Int is greater or equal to a bool")
	}
}

func TestLwIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when testing if a Int is lower than a bool")
	}
}

func TestLwEqIntErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Int is lower or equal to a bool")
	}
}

func TestIntAppendTypeErr(t *testing.T) {
	t1 := Int('A')
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending a bool to a Int")
	}
}

// Int interacts with float

func TestAddIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(3.0) {
		t.Error("Expected 3.0, got ", result)
	}
}

func TestAddNegIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(-2.0)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-1.0) {
		t.Error("Expected -1.0, got ", result)
	}
}

func TestSubIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(3.0)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(1.0) {
		t.Error("Expected 1.0, got ", result)
	}
}

func TestNegSubIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(-3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(7.0) {
		t.Error("Expected 7.0, got ", result)
	}
}

func TestMulIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(2.0)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(8.0) {
		t.Error("Expected 8.0, got ", result)
	}
}

func TestMulNegIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(-2.0)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-8.0) {
		t.Error("Expected -8.0, got ", result)
	}
}

func TestDivIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(2.0)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(2.0) {
		t.Error("Expected 2.0, got ", result)
	}
}

func TestDivNegIntFloat(t *testing.T) {
	t1 := Int(4)
	t2 := Float(-2.0)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Float(-2.0) {
		t.Error("Expected -2.0, got ", result)
	}
}

func TestAndIntFloat(t *testing.T) {
	t1 := Int(2)
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFloatFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFloatFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntFloatFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFloatFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFloatFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntFloatFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqIntFloat(t *testing.T) {
	t1 := Int(0)
	t2 := Float(0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqIntFloatFalse(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqIntFloat(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqIntFloatFalse(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(0.5)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtIntFloatEq(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtIntFloatFalse(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(0.5)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntFloatEq(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqIntFloatFalse(t *testing.T) {
	t1 := Int(0)
	t2 := Float(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwIntFloatEq(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwIntFloatFalse(t *testing.T) {
	t1 := Int(2)
	t2 := Float(1)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqIntFloat(t *testing.T) {
	t1 := Int(1)
	t2 := Float(2.5)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntFloatEq(t *testing.T) {
	t1 := Int(1)
	t2 := Float(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected tue, got ", result)
	}
}

func TestLwEqIntFloatFalse(t *testing.T) {
	t1 := Int(2)
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

func TestAndIntBool(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndIntBoolFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndIntBoolFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndIntBoolFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrIntBool(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrIntBoolFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrIntBoolFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrIntBoolFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntBool(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntBoolFalseLeft(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntBoolFalseRight(t *testing.T) {
	t1 := Int(1)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorIntBoolFalseBoth(t *testing.T) {
	t1 := Int(0)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
