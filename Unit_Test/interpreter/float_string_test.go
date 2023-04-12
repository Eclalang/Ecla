package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"strconv"
	"testing"
)

func TestAddFloatString(t *testing.T) {
	t1 := eclaType.Float(3.14)
	t2 := eclaType.String("hello")

	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}
	expect1, _ := strconv.ParseFloat("3.14", 32)
	Newexpect1 := fmt.Sprintf("%6f", expect1)
	expect2 := "hello"
	expect := Newexpect1 + expect2
	expected := eclaType.String(expect)
	if result.GetValue() != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
