package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
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
