package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"reflect"
	"testing"
)

func TestMulListInt(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.Int(2)

	result, err := t1.Mul(t2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result.GetValue(), eclaType.List{eclaType.Int(1), eclaType.Int(2), eclaType.Int(1), eclaType.Int(2)}) {
		t.Error("Expected true, got ", result)
	}
}
