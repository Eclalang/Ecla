package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestEqBoolInt(t *testing.T) {
	t1 := eclaType.Bool(true)
	t2 := eclaType.Int(1)

	result, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", result)
	}
}
