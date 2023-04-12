package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"strconv"
	"testing"
)

func TestAddStringFloat(t *testing.T) {
	t1 := eclaType.String("Hello")
	t2 := eclaType.Float(1.1)

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	expect1, _ := strconv.ParseFloat("1.1", 32)
	Newexpect1 := fmt.Sprintf("%6f", expect1)
	expect2 := "Hello"
	expect := expect2 + Newexpect1
	expected := eclaType.String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
