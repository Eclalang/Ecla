package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddStringFloat(t *testing.T) {
	t1 := eclaType.String("Hello")
	t2 := eclaType.Float(1.1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("Hello1.1") {
		t.Error("Expected Hello1.1, got ", result)
	}
}
