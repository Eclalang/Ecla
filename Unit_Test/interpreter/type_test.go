package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

/*func TestAddTwoFloats(t *testing.T) {
	t1 := eclaType.Float(1.1)
	t2 := eclaType.Float(2.2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Float(3.3) {
		t.Error("Expected 3.3, got ", result, " t1 = ", t1, " t2 = ", t2)
	}
}*/

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

func TestAndBools(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Bool(true)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBools(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Bool(false)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotBool(t *testing.T) {
	t1 := eclaType.Bool(true)

	result, err := t1.Not(t1)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
