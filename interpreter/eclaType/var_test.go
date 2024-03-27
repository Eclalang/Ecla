package eclaType

import (
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewVarEmpty(t *testing.T) {
	t1, err := NewVarEmpty("test", "int")
	if err != nil {
		t.Error(err)
	}

	expected := &Var{Name: "test", Value: NewNullType("int")}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarString(t *testing.T) {
	t1, err := NewVar("test", parser.String, String("value"))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", String("value")}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarFloat(t *testing.T) {
	t1, err := NewVar("test", parser.Float, Int(1))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", Float(1.0)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarAny(t *testing.T) {
	tmp := NewAny(Int(0))
	t1, err := NewVar("test", parser.Any, Int(0))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", tmp}

	if t1.Value == expected.Value {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarNull(t *testing.T) {
	t1, err := NewVar("test", parser.Int, NewNullType(parser.Int))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", NewNullType(parser.Int)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarEmptyType(t *testing.T) {
	t1, err := NewVar("test", "", Int(0))
	if err != nil {
		t.Error(err)
	}

	expected := &Var{"test", Int(0)}

	if *t1 != *expected {
		t.Error("expected ", expected, ", got ", t1)
	}
}

func TestNewVarErr(t *testing.T) {
	_, err := NewVar("test", parser.Bool, Int(1))

	if err == nil {
		t.Error("Expected error when creating a var with the wrong type")
	}
}

func TestNewVarVar(t *testing.T) {
	t1, err := NewVar("test", parser.Int, Int(0))
	if err != nil {
		t.Error(err)
	}

	t2, e := NewVar("test", parser.Int, t1)
	if e != nil {
		t.Error(e)
	}

	if *t1 != *t2 {
		t.Error("Expected ", t1, ", got ", t2)
	}
}
