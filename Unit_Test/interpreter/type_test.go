package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddTwoIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

//func TestAddTwoFloats(t *testing.T) {
//	t1 := eclaType.NewFloat(1.1)
//	t2 := eclaType.NewFloat(2.2)
//
//	result, err := t1.ADD(t2)
//	if err != nil {
//		t.Error(err)
//	}
//	if result.GetValue() != 3.3 {
//		t.Error("Expected 3.3, got ", result.GetValue())
//	}
//}

func TestDivEcTwoIntegers(t *testing.T) {
	t1 := eclaType.Int(6)
	t2 := eclaType.Int(2)

	result, err := t1.DivEc(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

func TestEqTwoIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqStrings(t *testing.T) {
	t1 := eclaType.String("hello")
	t2 := eclaType.String("hello")

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3)}

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2), eclaType.Int(3)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3)}

	result, err := t1.Gt(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestGtEqStrings(t *testing.T) {
	t1 := eclaType.String("hello!")
	t2 := eclaType.String("hello")

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(1)

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
