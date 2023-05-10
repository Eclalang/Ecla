package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddStringBool(t *testing.T) {
	t1 := eclaType.String("Hello")
	t2 := eclaType.Bool(true)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("Hellotrue") {
		t.Error("Expected Hellotrue, got ", result)
	}
}
