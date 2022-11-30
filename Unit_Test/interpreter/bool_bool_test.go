package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestEqBools(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Bool(true)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotEqBools(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Bool(false)

	result, err := t1.NotEq(t2)
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
	t1 := eclaType.Bool(false)
	t2 := eclaType.Bool(true)

	result, err := t1.Or(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}

func TestNotBools(t *testing.T) {
	t1 := eclaType.Bool(true)

	result, err := t1.Not(t1)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(false) {
		t.Error("Expected false, got ", result)
	}
}
