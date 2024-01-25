package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestTransformTo(t *testing.T) {
	t1 := NewMainBus(nil)
	t1.TransformTo(BUS_RETURN)
	if !t1.IsReturn() {
		t.Errorf("Expected %s, got %s", "BUS_RETURN", "BUS_MAIN")
	}
}

func TestIsReturn(t *testing.T) {
	t1 := NewReturnBus(nil)
	if !t1.IsReturn() {
		t.Errorf("Expected %s, got %s", "BUS_RETURN", "BUS_MAIN")
	}
}

func TestIsMain(t *testing.T) {
	t1 := NewMainBus(nil)
	if !t1.IsMain() {
		t.Errorf("Expected %s, got %s", "BUS_MAIN", "BUS_RETURN")
	}
}

func TestIsNone(t *testing.T) {
	t1 := NewNoneBus()
	if !t1.IsNone() {
		t.Errorf("Expected %s, got %s", "BUS_NONE", "BUS_MAIN")
	}
}

func TestGetVal(t *testing.T) {
	value, err := eclaType.NewString("hello world")
	if err != nil {
		t.Errorf("Error creating string: %s", err.Error())
	}
	t1 := NewMainBus(value)
	if t1.GetVal() != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
}

func TestNewMainBus(t *testing.T) {
	value, err := eclaType.NewString("hello world")
	if err != nil {
		t.Errorf("Error creating string: %s", err.Error())
	}
	t1 := NewMainBus(value)
	if t1.Val != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
	if t1.Type != BUS_MAIN {
		t.Errorf("Expected %s, got %d", "BUS_MAIN", t1.Type)
	}
}

func TestNewReturnBus(t *testing.T) {
	value, err := eclaType.NewString("hello world")
	if err != nil {
		t.Errorf("Error creating string: %s", err.Error())
	}
	t1 := NewReturnBus(value)
	if t1.Val != value {
		t.Errorf("Expected %s, got %s", value, t1.GetVal())
	}
	if t1.Type != BUS_RETURN {
		t.Errorf("Expected %s, got %d", "BUS_RETURN", t1.Type)
	}
}

func TestNewNoneBus(t *testing.T) {
	t1 := NewNoneBus()
	if t1.Val != nil {
		t.Errorf("Expected %s, got %s", "nil", t1.GetVal())
	}
	if t1.Type != BUS_NONE {
		t.Errorf("Expected %s, got %d", "BUS_NONE", t1.Type)
	}
}
