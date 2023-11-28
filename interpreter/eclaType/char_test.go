package eclaType

import (
	"testing"
)

// Char interacts with Char

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
