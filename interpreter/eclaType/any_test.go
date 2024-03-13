package eclaType

import "testing"

func TestAnyString(t *testing.T) {
	t1 := NewAny(Bool(true))
	expected := "true"
	result := t1.String()

	if result != expected {
		t.Error("expected \"true\", got ", result)
	}
}

func TestAnyGetString(t *testing.T) {
	t1 := NewAny(Bool(true))
	expected, _ := NewString("true")
	result := t1.GetString()

	if result != expected {
		t.Error("expected \"true\", got ", result)
	}
}

func TestAnyGetValue(t *testing.T) {
	t1 := NewAny(Bool(true))
	result := t1.GetValue()

	if result != t1 {
		t.Error("expected ", t1, ", got ", result)
	}
}
