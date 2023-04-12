package interpreter

import (
	"github.com/tot0p/Ecla/interpreter"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestTransformTo(t *testing.T) {
	t1 := interpreter.NewMainBus(nil)
	t1.TransformTo(interpreter.BUS_RETURN)
	if !t1.IsReturn() {
		t.Errorf("Expected %s, got %s", "BUS_RETURN", "BUS_MAIN")
	}
}

func TestIsReturn(t *testing.T) {
	t1 := interpreter.NewReturnBus(nil)
	if !t1.IsReturn() {
		t.Errorf("Expected %s, got %s", "BUS_RETURN", "BUS_MAIN")
	}
}

func TestIsMain(t *testing.T) {
	t1 := interpreter.NewMainBus(nil)
	if !t1.IsMain() {
		t.Errorf("Expected %s, got %s", "BUS_MAIN", "BUS_RETURN")
	}
}

func TestIsNone(t *testing.T) {
	t1 := interpreter.NewNoneBus()
	if !t1.IsNone() {
		t.Errorf("Expected %s, got %s", "BUS_NONE", "BUS_MAIN")
	}
}

func TestGetVal(t *testing.T) {
	value := eclaType.NewString("hello world")
	t1 := interpreter.NewMainBus(value)
	if t1.GetVal() != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
}

func TestNewMainBus(t *testing.T) {
	value := eclaType.NewString("hello world")
	t1 := interpreter.NewMainBus(value)
	if t1.Val != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
	if t1.Type != interpreter.BUS_MAIN {
		t.Errorf("Expected %s, got %d", "BUS_MAIN", t1.Type)
	}
}

func TestNewReturnBus(t *testing.T) {
	value := eclaType.NewString("hello world")
	t1 := interpreter.NewReturnBus(value)
	if t1.Val != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
	if t1.Type != interpreter.BUS_RETURN {
		t.Errorf("Expected %s, got %d", "BUS_RETURN", t1.Type)
	}
}

func TestNewNoneBus(t *testing.T) {
	t1 := interpreter.NewNoneBus()
	if t1.Val != nil {
		t.Errorf("Expected %s, got %s", "nil", t1.GetVal())
	}
	if t1.Type != interpreter.BUS_NONE {
		t.Errorf("Expected %s, got %d", "BUS_NONE", t1.Type)
	}
}
