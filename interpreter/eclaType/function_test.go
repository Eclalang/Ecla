package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"testing"
)

// function interacts with function

func TestGetType(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetType() != "function()" {
		t.Errorf("Expected function(), got %v", f.GetType())
	}
}

// Useless tests for code coverage purpose
func TestGetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetValue() == nil {
		t.Error("Expected not nil, got nil")
	}
}

func TestSetValue(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if err := f.SetValue(nil); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func TestString(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.String() != "function" {
		t.Errorf("Expected function, got %v", f.String())
	}
}

func TestGetString(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.GetString() != "function" {
		t.Errorf("Expected function, got %v", f.GetString().String())
	}
}

func TestGetIndex(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	expect, err := f.GetIndex(nil)
	if expect != nil || err == nil {
		t.Errorf("Expected nil & error, got %v & %v", expect, err)
	}
}

func TestIsNull(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f.IsNull() {
		t.Errorf("Expected false, got %v", f.IsNull())
	}
}

func TestNewFunction(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	if f == nil {
		t.Errorf("Expected not nil, got %v", f)
	}
}

func TestTypeAndNumberOfArgsIsCorrect(t *testing.T) {
	var structDecl []eclaDecl.TypeDecl
	f := NewFunction("test", nil, nil, nil)
	var args []Type
	test, expect := f.TypeAndNumberOfArgsIsCorrect(args, structDecl)
	if !test || expect == nil {
		t.Errorf("Expected true, got %v", test)
	}
}

func TestCheckReturn(t *testing.T) {
	f := NewFunction("test", nil, nil, nil)
	var structDecl []eclaDecl.TypeDecl
	var args []Type
	expect := f.CheckReturn(args, structDecl)
	if !expect {
		t.Errorf("Expected true, got %v", expect)
	}
}
