package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddIntString(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("1hello") {
		t.Error("Expected 1hello, got ", result)
	}
}

func TestMulIntString(t *testing.T) {
	t1 := eclaType.Int(2)
	t2 := eclaType.String("hello")

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("hellohello") {
		t.Error("Expected hellohello, got ", result)
	}
}
