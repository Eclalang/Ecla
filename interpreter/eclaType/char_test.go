package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"testing"
)

// Char interacts with Char

func TestNewChar(t *testing.T) {
	t1, _ := NewChar("c")

	if t1 != 'c' {
		t.Error("Error when creating a Char")
	}
}

func TestNewEmptyChar(t *testing.T) {
	t1, _ := NewChar("")

	if t1 != Char(0) {
		t.Error("Error when creating a Char")
	}
}

func TestCharGetValue(t *testing.T) {
	t1 := Char('c')

	result := t1.GetValue()
	if result != Char('c') {
		t.Error("Expected 'c', got ", result)
	}
}

func TestCharGetType(t *testing.T) {
	t1 := Char('c')

	result := t1.GetType()
	if result != "char" {
		t.Error("Expected \"char\", got ", result)
	}
}

func TestCharIsNull(t *testing.T) {
	t1 := Char('c')

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
	t1 := Char('B')
	t2 := Char('!')

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('!') {
		t.Error("Expected '!', got ", result)
	}
}

func TestNegSubChar(t *testing.T) {
	t1 := Char('c')
	t2 := Char(-3)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('f') {
		t.Error("Expected 'f', got ", result)
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
		t.Error("Expected 32, got ", result)
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
		t.Error("Expected 'B', got ", result)
	}
}

func TestMulNegChar(t *testing.T) {
	t1 := Char(4)
	t2 := Char(-2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-8) {
		t.Error("Expected -8, got ", result)
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
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcChar(t *testing.T) {
	t1 := Char('B')
	t2 := Char('!')

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegChar(t *testing.T) {
	t1 := Char(5)
	t2 := Char(-2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-2) {
		t.Error("Expected -2, got ", result)
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
	t1 := Char('B')
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
		t.Error("Expected false, got ", result)
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
		t.Error("Expected false, got ", result)
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
		t.Error("Expected false, got ", result)
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

func TestXorChar(t *testing.T) {
	t1 := Char('A')
	t2 := Char('B')

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Char('A')

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Char(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotCharFalse(t *testing.T) {
	t1 := Char(1)

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

func TestCharGetSize(t *testing.T) {
	t1 := Char(0)
	expected := utils.Sizeof(t1)

	result := t1.GetSize()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestCharGetValueAsInt(t *testing.T) {
	t1 := Char('A')

	result := t1.GetValueAsInt()
	if result != Int(65) {
		t.Error("Expected 65, got ", result)
	}
}

// Char interacts with Int

func TestAddCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected 'B', got ", result)
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
		t.Error("Expected 'A', got ", result)
	}
}

func TestSubCharInt(t *testing.T) {
	t1 := Char('B')
	t2 := Int(33)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('!') {
		t.Error("Expected '!', got ", result)
	}
}

func TestSubCharIntNeg(t *testing.T) {
	t1 := Char('!')
	t2 := Int(-33)

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected 'B', got ", result)
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

func TestMulCharInt(t *testing.T) {
	t1 := Char('!')
	t2 := Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char('B') {
		t.Error("Expected 'B', got ", result)
	}
}

func TestDivCharInt(t *testing.T) {
	t1 := Char('B')
	t2 := Int(33)

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcCharInt(t *testing.T) {
	t1 := Char('B')
	t2 := Int(33)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
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
	t2 := Int(25)

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

func TestNotEqCharIntFalse(t *testing.T) {
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
	t1 := Char('B')
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

func TestGtCharCharEq(t *testing.T) {
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
	t2 := Char(64)

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
	t1 := Char('B')
	t2 := Int(1)

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

func TestXorCharInt(t *testing.T) {
	t1 := Char('A')
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharIntFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Int(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharIntFalseRight(t *testing.T) {
	t1 := Char('A')
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharIntFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Int(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Char interacts with String

func TestAddCharString(t *testing.T) {
	t1 := Char('a')
	t2 := String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != String("ahello") {
		t.Error("Expected ahello, got ", result)
	}
}

func TestCharString(t *testing.T) {
	t1 := Char('0')
	t2 := "0"

	result := t1.String()
	if result != t2 {
		t.Error("Expected \"0\", got ", result)
	}
}

func TestCharGetString(t *testing.T) {
	t1 := Char('0')
	t2 := String("0")

	result, err := t1.GetString().Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected \"0\", got ", result)
	}
}

// test errors Char

func TestCharAndErr(t *testing.T) {
	t1 := Char(5)
	t2 := String("test")

	_, err := t1.And(t2)
	if err == nil {
		t.Error("Expected error when checking Char and string")
	}
}

func TestCharOrErr(t *testing.T) {
	t1 := Char(5)
	t2 := String("test")

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("Expected error when checking Char or string")
	}
}

func TestCharXorErr(t *testing.T) {
	t1 := Char(5)
	t2 := String("test")

	_, err := t1.Xor(t2)
	if err == nil {
		t.Error("Expected error when checking Char xor string")
	}
}

func TestCharLenErr(t *testing.T) {
	t1 := Char(5)

	_, err := t1.Len()
	if err == nil {
		t.Error("Expected error when checking len of Char")
	}
}

func TestDivEcCharFloatErr(t *testing.T) {
	t1 := Char(5)
	t2 := Float(2)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing ec by float")
	}
}

func TestNewCharWithString(t *testing.T) {
	_, err := NewChar("test")

	if err == nil {
		t.Error("Expected error when creating a char with a string")
	}
}

func TestCharAppendErr(t *testing.T) {
	t1 := Char(0)
	t2 := Char(1)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCharSetValueErr(t *testing.T) {
	t1 := Char(0)

	err := t1.SetValue(1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCharGetIndexErr(t *testing.T) {
	t1 := Char(0)

	_, err := t1.GetIndex(Char(0))
	if err == nil {
		t.Error("Expected error when indexing")
	}
}

func TestAddCharErr(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding a bool to an Char")
	}
}

func TestSubCharErr(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when subtracting a bool to an Char")
	}
}

func TestMulCharErr(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("Expected error when multiplying an Char by a bool")
	}
}

func TestDivCharErr(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an Char by a bool")
	}
}

func TestDivBy0CharErr(t *testing.T) {
	t1 := Char(0)
	t2 := Char(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an Char by 0")
	}
}

func TestDivBy0CharIntErr(t *testing.T) {
	t1 := Char(5)
	t2 := Int(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("Expected error when dividing an Char by 0")
	}
}

func TestModCharErr(t *testing.T) {
	t1 := Char(1)
	t2 := Bool(true)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when doing a modulo of an Char by a bool")
	}
}

func TestModBy0CharErr(t *testing.T) {
	t1 := Char(1)
	t2 := Char(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a Char by 0")
	}
}

func TestModBy0CharIntErr(t *testing.T) {
	t1 := Char(1)
	t2 := Int(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("Expected error when modding a Char by 0")
	}
}

func TestDivEcCharErr(t *testing.T) {
	t1 := Char(1)
	t2 := Bool(true)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Char by a bool")
	}
}

func TestDivEcBy0CharErr(t *testing.T) {
	t1 := Char(1)
	t2 := Char(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Char by 0")
	}
}

func TestDivEcBy0CharIntErr(t *testing.T) {
	t1 := Char(1)
	t2 := Int(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("Expected error when dividing a Char by 0")
	}
}

func TestEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Eq(t2)
	if err == nil {
		t.Error("Expected error when testing equality between Char and bool")
	}
}

func TestNotEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.NotEq(t2)
	if err == nil {
		t.Error("Expected error when testing inequality between Char and bool")
	}
}

func TestGtCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("Expected error when testing if a Char is greater than a bool")
	}
}

func TestGtEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Char is greater or equal to a bool")
	}
}

func TestLwCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("Expected error when testing if a Char is lower than a bool")
	}
}

func TestLwEqCharErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("Expected error when testing if a Char is lower or equal to a bool")
	}
}

func TestCharAppendTypeErr(t *testing.T) {
	t1 := Char('A')
	t2 := Bool(true)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("Expected error when appending a bool to a Char")
	}
}

// Char Chareracts with float

func TestAndCharFloat(t *testing.T) {
	t1 := Char('2')
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
	t1 := Char('1')
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
	t1 := Char('1')
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
	t1 := Char('1')
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

func TestXorCharFloat(t *testing.T) {
	t1 := Char('1')
	t2 := Float(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharFloatFalseRight(t *testing.T) {
	t1 := Char('1')
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorCharFloatFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Float(1)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorCharFloatFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Float(0)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

// Char interacts with Bool

func TestAndCharBool(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharBoolFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharBoolFalseRight(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharBoolFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(false)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrCharBool(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharBoolFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrCharBoolFalseRight(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrCharBoolFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharBool(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharBoolFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(true)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharBoolFalseRight(t *testing.T) {
	t1 := Char('1')
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharBoolFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := Bool(false)

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Char intreracts with Var

func TestAddCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestAddNegCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(-2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestSubCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestNegSubCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(-3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(7) {
		t.Error("Expected 7, got ", result)
	}
}

func TestModCharVar(t *testing.T) {
	t1 := Char(3)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(8) {
		t.Error("Expected 8, got ", result)
	}
}

func TestMulNegCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(-2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-8) {
		t.Error("Expected -8, got ", result)
	}
}

func TestDivCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivNegCharVar(t *testing.T) {
	t1 := Char(4)
	t2, _ := NewVar("testVar", "char", Char(-2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestDivEcCharVar(t *testing.T) {
	t1 := Char(5)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegCharVar(t *testing.T) {
	t1 := Char(5)
	t2, _ := NewVar("testVar", "char", Char(-2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestEqCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqCharFalseVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqCharVarFalse(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharVarTrue(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtCharVarFalse(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharVarEq(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharVarEq(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharVarFalse(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwCharVarFalse(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharVarEq(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharVar(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(3))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqCharVarFalse(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(1))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharVarEq(t *testing.T) {
	t1 := Char(2)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharVarFalseRight(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharVarFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharVarFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharVarFalseRight(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharVarFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharVarFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorCharVar(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharVarFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharVarFalseRight(t *testing.T) {
	t1 := Char(1)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharVarFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2, _ := NewVar("testVar", "char", Char(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

// Char Chareracts with Any

func TestAddCharAnt(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestAddNegCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(-2))

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-1) {
		t.Error("Expected -1, got ", result)
	}
}

func TestSubCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestNegSubCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(-3))

	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(7) {
		t.Error("Expected 7, got ", result)
	}
}

func TestModCharAny(t *testing.T) {
	t1 := Char(3)
	t2 := NewAny(Char(2))

	result, err := t1.Mod(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(1) {
		t.Error("Expected 1, got ", result)
	}
}

func TestMulCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(8) {
		t.Error("Expected 8, got ", result)
	}
}

func TestMulNegCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(-2))

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-8) {
		t.Error("Expected -8, got ", result)
	}
}

func TestDivCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivNegCharAny(t *testing.T) {
	t1 := Char(4)
	t2 := NewAny(Char(-2))

	result, err := t1.Div(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestDivEcCharAny(t *testing.T) {
	t1 := Char(5)
	t2 := NewAny(Char(2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(2) {
		t.Error("Expected 2, got ", result)
	}
}

func TestDivEcNegCharAny(t *testing.T) {
	t1 := Char(5)
	t2 := NewAny(Char(-2))

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Char(-2) {
		t.Error("Expected -2, got ", result)
	}
}

func TestEqCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(1))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqCharFalseAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestNotEqCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqCharAnyFalse(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(1))

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharAnyTrue(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtCharAnyFalse(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtCharAnyEq(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(1))

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestGtEqCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(0))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharAnyEq(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(1))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqCharAnyFalse(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwCharAnyFalse(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(1))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwCharAnyEq(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(2))

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharAny(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(3))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqCharAnyFalse(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(1))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestLwEqCharAnyEq(t *testing.T) {
	t1 := Char(2)
	t2 := NewAny(Char(2))

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestAndCharAnyFalseRight(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharAnyFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(2))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndCharAnyFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(0))

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestOrCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharAnyFalseRight(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharAnyFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(2))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrCharAnyFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(0))

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected true, got ", result)
	}
}

func TestXorCharAny(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharAnyFalseLeft(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(2))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharAnyFalseRight(t *testing.T) {
	t1 := Char(1)
	t2 := NewAny(Char(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(true) {
		t.Error("Expected false, got ", result)
	}
}

func TestXorCharAnyFalseBoth(t *testing.T) {
	t1 := Char(0)
	t2 := NewAny(Char(0))

	result, err := t1.Xor(t2)
	if err != nil {
		t.Error(err)
	}
	if result != Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
