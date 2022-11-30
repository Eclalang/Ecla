package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"reflect"
	"testing"
)

func TestAddLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3)}

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result.GetValue(), eclaType.List{eclaType.Int(1), eclaType.Int(2), eclaType.Int(1), eclaType.Int(3)}) {
		t.Error("Expected true, got ", result)
	}
}

func TestEqLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
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

func TestGtEqLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2), eclaType.Int(3)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3)}

	result, err := t1.GtEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3), eclaType.String("a")}

	result, err := t1.Lw(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestLwEqLists(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.List{eclaType.Int(1), eclaType.Int(3)}

	result, err := t1.LwEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
