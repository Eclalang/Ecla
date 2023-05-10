package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestAddBoolString(t *testing.T) {
	t1 := eclaType.NewBool("true")
	t2 := eclaType.NewString("Hello")
	actual, err := t1.Add(t2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	expected := eclaType.NewString("trueHello")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
