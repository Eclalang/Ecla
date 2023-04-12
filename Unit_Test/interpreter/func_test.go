package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

// Useless tests for code coverage purpose
func TestGetValue(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f.GetValue() != nil {
		t.Errorf("Expected nil, got %v", f.GetValue())
	}
}

func TestSetValue(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if err := f.SetValue(nil); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestString(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f.String() != "function" {
		t.Errorf("Expected function, got %v", f.String())
	}
}

func TestGetString(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f.GetString() != "function" {
		t.Errorf("Expected function, got %v", f.GetString().String())
	}
}

func TestGetType(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f.GetType() != "function" {
		t.Errorf("Expected function, got %v", f.GetType())
	}
}

func TestGetIndex(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	expect, err := f.GetIndex(nil)
	if expect != nil || err != nil {
		t.Errorf("Expected nil & nil, got %v & %v", expect, err)
	}
}

func TestIsNull(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f.IsNull() {
		t.Errorf("Expected false, got %v", f.IsNull())
	}
}

func TestNewFunction(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	if f == nil {
		t.Errorf("Expected not nil, got %v", f)
	}
}

func TestTypeAndNumberOfArgsIsCorrect(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	var args []eclaType.Type
	test, expect := f.TypeAndNumberOfArgsIsCorrect(args)
	if !test || expect == nil {
		t.Errorf("Expected true, got %v", test)
	}
}

func TestCheckReturn(t *testing.T) {
	f := eclaType.NewFunction("test", nil, nil, nil)
	var args []eclaType.Type
	expect := f.CheckReturn(args)
	if !expect {
		t.Errorf("Expected true, got %v", expect)
	}
}
