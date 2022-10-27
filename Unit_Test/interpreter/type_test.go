package interpreter

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestAddTwoIntegers(t *testing.T) {
	t1 := eclaType.Int(1)
	t2 := eclaType.Int(2)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	if result.GetValue() != eclaType.Int(3) {
		t.Error("Expected 3, got ", result)
	}
}

//func TestAddTwoFloats(t *testing.T) {
//	t1 := eclaType.NewFloat(1.1)
//	t2 := eclaType.NewFloat(2.2)
//
//	result, err := t1.ADD(t2)
//	if err != nil {
//		t.Error(err)
//	}
//	if result.GetValue() != 3.3 {
//		t.Error("Expected 3.3, got ", result.GetValue())
//	}
//}
