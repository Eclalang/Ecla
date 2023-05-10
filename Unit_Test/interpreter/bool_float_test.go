package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestEqBoolFloat(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Float(1.0)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBoolFloat(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Float(1.0)

	result, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}

func TestAndBoolFloat(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Float(1.0)

	result, err := t1.And(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestOrBoolFloat(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Float(0.0)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
