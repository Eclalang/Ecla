package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"reflect"
	"testing"
)

func TestAddListString(t *testing.T) {
	t1 := eclaType.List{eclaType.Int(1), eclaType.Int(2)}
	t2 := eclaType.String("Hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result.GetValue(), eclaType.String("[1 2]Hello")) {
		t.Error("Expected [1 2]Hello, got ", result)
	}
}
