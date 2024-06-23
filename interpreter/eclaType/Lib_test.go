package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"testing"
)

func TestNewLib(t *testing.T) {
	t1 := NewLib("console")
	t2 := &Lib{"console"}

	if *t1 != *t2 {
		t.Error("error when creating Lib")
	}
}

func TestLibGetValue(t *testing.T) {
	t1 := NewLib("console")
	result := t1.GetValue()

	if t1 != result {
		t.Error("error when getting value of Lib")
	}
}

func TestLibSetValue(t *testing.T) {
	t1 := NewLib("console")
	result := t1.SetValue("test")

	if result == nil {
		t.Error("expected error when setting value of Lib")
	}
}

func TestLibString(t *testing.T) {
	t1 := NewLib("console")
	result := t1.String()

	if result != "console" {
		t.Error("expected \"console\", got " + result)
	}
}

func TestLibGetString(t *testing.T) {
	t1 := NewLib("console")
	t2 := String("console")
	result := t1.GetString()

	if result != t2 {
		t.Error("expected " + t2 + ", got " + result)
	}
}

func TestLibGetType(t *testing.T) {
	t1 := NewLib("console")
	t2 := "lib"
	result := t1.GetType()

	if result != t2 {
		t.Error("expected \"lib\", got " + result)
	}
}

func TestLibGetIndex(t *testing.T) {
	t1 := NewLib("console")
	_, err := t1.GetIndex(Int(0))

	if err == nil {
		t.Error("expected error when getting index of Lib")
	}
}

func TestLibAdd(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Add(t2)

	if err == nil {
		t.Error("expected error when adding to a lib")
	}
}

func TestLibSub(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Sub(t2)

	if err == nil {
		t.Error("expected error when subtracting from a lib")
	}
}

func TestLibMul(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Mul(t2)

	if err == nil {
		t.Error("expected error when multiplying with a lib")
	}
}

func TestLibDiv(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Div(t2)

	if err == nil {
		t.Error("expected error when dividing a lib")
	}
}

func TestLibMod(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Mod(t2)

	if err == nil {
		t.Error("expected error when getting remainder of a lib")
	}
}

func TestLibDivEc(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.DivEc(t2)

	if err == nil {
		t.Error("expected error when getting quotient of a lib")
	}
}

func TestLibEq(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Eq(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibNotEq(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.NotEq(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibGt(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Gt(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibGtEq(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.GtEq(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibLw(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Lw(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibLwEq(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.LwEq(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibAnd(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.And(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibOr(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Or(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibNot(t *testing.T) {
	t1 := NewLib("console")
	_, err := t1.Not()

	if err == nil {
		t.Error("expected error when getting \"not\" of a lib")
	}
}

func TestLibXor(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Xor(t2)

	if err == nil {
		t.Error("expected error when comparing a lib")
	}
}

func TestLibIsNull(t *testing.T) {
	t1 := NewLib("console")
	result := t1.IsNull()

	if result {
		t.Error("expected false, got true")
	}
}

func TestLibAppend(t *testing.T) {
	t1 := NewLib("console")
	t2 := NewLib("test")
	_, err := t1.Append(t2)

	if err == nil {
		t.Error("expected error when appending to a lib")
	}
}

func TestLibGetSize(t *testing.T) {
	t1 := NewLib("console")
	t2 := utils.Sizeof(t1)
	result := t1.GetSize()

	if result != t2 {
		t.Errorf("expected %d, got %d", t2, result)
	}
}

func TestLibLen(t *testing.T) {
	t1 := NewLib("console")
	_, err := t1.Len()

	if err == nil {
		t.Error("expected error when getting length of a lib")
	}
}
