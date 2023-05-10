package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddStringInt(t *testing.T) {
	t1 := eclaType.String("Hello")
	t2 := eclaType.Int(1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("Hello1") {
		t.Error("Expected Hello1, got ", result)
	}
}

func TestMulStringInt(t *testing.T) {
	t1 := eclaType.String("Hello")
	t2 := eclaType.Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("HelloHello") {
		t.Error("Expected HelloHello, got ", result)
	}
}
