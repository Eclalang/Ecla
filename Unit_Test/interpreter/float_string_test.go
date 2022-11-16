package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddFloatString(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.String("3.14hello") {
		t.Error("Expected 3.14hello, got ", result)
	}
}
