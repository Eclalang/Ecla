package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// Null interacts with Null

func TestNewNull(t *testing.T) {
	t1 := NewNull()
	t2 := Null{""}

	if t1 != t2 {
		t.Error("error when creating null")
	}
}

func TestNewNullType(t *testing.T) {
	t1 := NewNullType(parser.Int)
	t2 := Null{parser.Int}

	if t1 != t2 {
		t.Error("error when creating null type")
	}
}

func TestNullString(t *testing.T) {
	t1 := NewNull()
	t2 := "null"

	result := t1.String()

	if result != t2 {
		t.Error("expected \"null\", got " + result)
	}
}

func TestNullGetString(t *testing.T) {
	t1 := NewNull()
	t2 := String("null")

	result := t1.GetString()

	if result != t2 {
		t.Error("expected \"null\", got " + result)
	}
}

func TestNullGetType(t *testing.T) {
	t1 := NewNullType(parser.Int)
	t2 := Int(0)

	result := t1.GetType()

	if result != t2.GetType() {
		t.Error("expected \"int\", got " + result)
	}
}

func TestNullEq(t *testing.T) {
	t1 := Null{parser.Int}
	t2 := Null{parser.String}

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}

	if result != Bool(true) {
		t.Error("expected true, got" + result.String())
	}
}

func TestNullEqFalse(t *testing.T) {
	t1 := Null{parser.Int}
	t2 := Int(0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}

	if result != Bool(false) {
		t.Error("expected false, got" + result.String())
	}
}

func TestNullNotEq(t *testing.T) {
	t1 := Null{parser.Int}
	t2 := Int(0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}

	if result != Bool(true) {
		t.Error("expected true, got" + result.String())
	}
}

func TestNullNotEqFalse(t *testing.T) {
	t1 := Null{parser.Int}
	t2 := Null{parser.String}

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}

	if result != Bool(false) {
		t.Error("expected false, got" + result.String())
	}
}

func TestNullIsNull(t *testing.T) {
	t1 := Null{""}

	result := t1.IsNull()
	if result != true {
		t.Error("expected \"true\", got false")
	}
}

func TestNullGetSize(t *testing.T) {
	t1 := Null{""}

	result := t1.GetSize()
	if result != utils.Sizeof(t1) {
		t.Errorf("expected %d, got %d", utils.Sizeof(t1), result)
	}
}

// Null test errors

func TestNullGetValue(t *testing.T) {
	t1 := NewNull()
	result := t1.GetValue()

	switch result.(type) {
	case error:
		return
	default:
		t.Errorf("expected error, got %T", result)
	}
}

func TestNullSetValue(t *testing.T) {
	t1 := NewNull()
	err := t1.SetValue(0)

	if err == nil {
		t.Error("expected error when setting value of null")
	}
}

func TestNullGetIndex(t *testing.T) {
	t1 := NewNull()
	_, err := t1.GetIndex(Int(0))

	if err == nil {
		t.Error("expected error when getting index of null")
	}
}

func TestNullAdd(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Add(t2)
	if err == nil {
		t.Error("expected error when adding null with non-string")
	}
}

func TestNullSub(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("expected error when subtracting from null")
	}
}

func TestNullMul(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Mul(t2)
	if err == nil {
		t.Error("expected error when multiplying with null")
	}
}

func TestNullDiv(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Div(t2)
	if err == nil {
		t.Error("expected error when dividing null")
	}
}

func TestNullMod(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Mod(t2)
	if err == nil {
		t.Error("expected error when getting remainder of null")
	}
}

func TestNullDivEc(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.DivEc(t2)
	if err == nil {
		t.Error("expected error when getting quotient of null")
	}
}

func TestNullAnd(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.And(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullOr(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Or(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}
func TestNullXor(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Xor(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullNot(t *testing.T) {
	t1 := Null{""}

	_, err := t1.Not()
	if err == nil {
		t.Error("expected error when getting \"not\" of null")
	}
}

func TestNullGt(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Gt(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullGtEq(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.GtEq(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullLw(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Lw(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullLwEq(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.LwEq(t2)
	if err == nil {
		t.Error("expected error when comparing to null")
	}
}

func TestNullAppend(t *testing.T) {
	t1 := Null{""}
	t2 := Int(0)

	_, err := t1.Append(t2)
	if err == nil {
		t.Error("expected error when appending to null")
	}
}

func TestNullGetLength(t *testing.T) {
	t1 := Null{""}

	_, err := t1.Len()
	if err == nil {
		t.Error("expected error when getting length of null")
	}
}

// Null interacts with String

func TestNullStringAdd(t *testing.T) {
	t1 := Null{""}
	t2 := String("test")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}

	if result != String("nulltest") {
		t.Error("expected \"nulltest\", got " + result.String())
	}
}
