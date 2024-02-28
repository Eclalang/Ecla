package eclaType

import (
	"testing"
)

// Char interacts with Char

func TestNewChar(t *testing.T) {
	t1 := Char('A')

	if t1 != 'A' {
		t.Error("Error when creating a char")
	}
}

func TestCharGetValue(t *testing.T) {
	t1 := Char('A')

	result := t1.GetValue()
	if result != Char('A') {
		t.Error("Expected 'A', got ", result)
	}
}

func TestCharGetType(t *testing.T) {
	t1 := Char('A')

	result := t1.GetType()
	if result != "char" {
		t.Error("Expected char, got ", result)
	}
}

func TestCharIsNull(t *testing.T) {
	t1 := Char('A')

	result := t1.IsNull()
	if result != false {
		t.Error("Expected false, got", result)
	}
}

func TestAddChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('!')

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('b') {
		t.Error("Expected b, got ", result)
	}
}

func TestSubChar(t *testing.T) {
	t1 := Char('b')
	t2 := Char('!')

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('A') {
		t.Error("Expected A, got ", result)
	}
}

func TestModChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('!')

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(32) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulCharChar(t *testing.T) {
	t1 := Char('!')
	t2 := Char(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected B, got ", result)
	}
}

func TestDivChar(t *testing.T) {
	t1 := Char('B')
	t2 := Char('!')

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected !, got ", result)
	}
}

func TestDivEcChar(t *testing.T) {
	t1 := Char('B')
	t2 := Char('!')

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected !, got ", result)
	}
}

func TestEqChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqCharFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqChar(t *testing.T) {
	t1 := Char('!')
	t2 := Char('"')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqCharFalse(t *testing.T) {
	t1 := Char('!')
	t2 := Char('!')

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharTrue(t *testing.T) {
	t1 := Char('A')
	t2 := Char('!')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtCharFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharEq(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharEq(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwCharFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Char('!')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharEq(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqCharFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Char('!')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharEq(t *testing.T) {
	t1 := Char('A')
	t2 := Char('A')

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Char('B')

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Char(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Char('B')

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Char(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotCharFalse(t *testing.T) {
	t1 := Char('A')

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got", result)
	}
}

func TestNotCharTrue(t *testing.T) {
	t1 := Char(0)

	result, err := t1.Not()
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got", result)
	}
}

// Char interacts with Int

func TestAddCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('C') {
		t.Error("Expected C, got ", result)
	}
}

func TestAddCharIntNeg(t *testing.T) {
	t1 := Char('B')
	t2 := Int(-1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('A') {
		t.Error("Expected A, got ", result)
	}
}

func TestSubCharInt(t *testing.T) {
	t1 := Char('B')
	t2 := Int(1)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('A') {
		t.Error("Expected A, got ", result)
	}
}

func TestSubCharIntNeg(t *testing.T) {
	t1 := Char('A')
	t2 := Int(-1)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected B, got ", result)
	}
}

func TestModCharInt(t *testing.T) {
	t1 := Char('!')
	t2 := Int(2)

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulCharCharInt(t *testing.T) {
	t1 := Char('!')
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected B, got ", result)
	}
}

func TestDivCharInt(t *testing.T) {
	t1 := Char('B')
	t2 := Int(2)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('!') {
		t.Error("Expected !, got ", result)
	}
}

func TestDivEcCharInt(t *testing.T) {
	t1 := Char('C')
	t2 := Int(2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('!') {
		t.Error("Expected !, got ", result)
	}
}

func TestEqCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(65)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqCharIntFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Int(66)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqCharInt(t *testing.T) {
	t1 := Char('!')
	t2 := Int(32)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqCharCharIntFalse(t *testing.T) {
	t1 := Char('!')
	t2 := Int(33)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharIntTrue(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtCharIntFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Int(66)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharIntEq(t *testing.T) {
	t1 := Char('A')
	t2 := Int(65)

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharIntEq(t *testing.T) {
	t1 := Char('A')
	t2 := Int(65)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharIntFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Int(66)

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(66)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwCharIntFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Int(64)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharIntEq(t *testing.T) {
	t1 := Char('A')
	t2 := Int(65)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(66)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqCharIntFalse(t *testing.T) {
	t1 := Char('A')
	t2 := Int(64)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharIntEq(t *testing.T) {
	t1 := Char('A')
	t2 := Int(65)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharIntFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharIntFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Int(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharIntFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Int(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharIntFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharIntFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Int(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharIntFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Int(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

// Char interacts with String

func TestAddCharString(t *testing.T) {
	t1 := Char('!')
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("!hello") {
		t.Error("Expected 1hello, got ", result)
	}
}

func TestCharString(t *testing.T) {
	t1 := Char('!')
	t2 := "!"

	result := t1.String() == (t2)
	if result != true {
		t.Error("Expected true, got ", result)
	}
}
func TestCharGetString(t *testing.T) {
	t1 := Char('!')
	t2 := String("!")

	result, err := t1.GetString().Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestCharAppendChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCharAppendString(t *testing.T) {
	t1 := Char('A')
	t2 := String("BC")

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error")
	}
}

// test errors Char

func TestCharSetValueErr(t *testing.T) {
	t1 := Char('A')

	err := t1.SetValue('B')
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCharGetIndexErr(t *testing.T) {
	t1 := Char('A')

	_, err := t1.GetIndex(Int(0))
	if err == nil {
		t.Error("Expected error when indexing")
	}
}

func TestAddCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding a bool to a char")
	}
}

func TestSubCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting a bool to a char")
	}
}

func TestMulCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying a char by a bool")
	}
}

func TestDivCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by a bool")
	}
}

func TestDivBy0CharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by 0")
	}
}

func TestDivBy0CharIntErr(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by 0")
	}
}

func TestModCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when doing a modulo of a char by a bool")
	}
}

func TestModBy0CharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a char by 0")
	}
}

func TestModBy0CharIntErr(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a char by 0")
	}
}

func TestDivEcCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by a bool")
	}
}

func TestDivEcBy0CharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by 0")
	}
}

func TestDivEcBy0CharIntErr(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a char by 0")
	}
}

func TestEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when testing equality between char and bool")
	}
}

func TestNotEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when testing inequality between char and bool")
	}
}

func TestGtCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when testing if a char is greater than a bool")
	}
}

func TestGtEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a char is greater or equal to a bool")
	}
}

func TestLwCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when testing if a char is lower than a bool")
	}
}

func TestLwEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a char is lower or equal to a bool")
	}
}

func TestCharAppendErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending a bool to a char")
	}
}

// Char interacts with float
func TestAndCharFloat(t *testing.T) {
	t1 := Char('A')
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFloatFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFloatFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Float(1)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharFloatFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Float(0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFloat(t *testing.T) {
	t1 := Char('A')
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFloatFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFloatFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Float(1)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharFloatFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Float(0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}
