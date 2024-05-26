package eclaType

import (
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewMap(t *testing.T) {
	t1 := NewMap()
	if len(t1.Keys) != 0 || t1.Keys == nil {
		t.Error("Error when initializing Keys")
	}
	if len(t1.Values) != 0 || t1.Values == nil {
		t.Error("Error when initializing Values")
	}
	if t1.Typ != "" {
		t.Error("Error when initialising Typ")
	}
	if t1.TypKey != "" {
		t.Error("Error when initialising TypKey")
	}
	if t1.TypVal != "" {
		t.Error("Error when initialising TypVal")
	}
}

func TestMapGetValue(t *testing.T) {
	t1 := NewMap()
	t2 := t1.GetValue()
	if t1 != t2 {
		t.Errorf("Expected %s, got %s", t1, t2)
	}
}

func TestMapSetAutoTypeEmpty(t *testing.T) {
	t1 := NewMap()
	expected := "empty"
	err := t1.SetAutoType()
	if err != nil {
		t.Error(err)
	}
	if t1.Typ != expected {
		t.Errorf("Expected %s, got %s", expected, t1.Typ)
	}
}

func TestMapSetAutoTypeIntString(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("One")}, "", parser.Int, parser.String}
	err := t1.SetAutoType()
	if err != nil {
		t.Error(err)
	}
	expected := "map[int]string"

	if t1.Typ != expected {
		t.Errorf("Expected %s, got %s", expected, t1.Typ)
	}
}

// Test Map errors

func TestMapSetAutoTypeErrorKeys(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Bool(false)}, []Type{Int(1), Int(2)}, "", "", ""}
	err := t1.SetAutoType()
	if err == nil {
		t.Error("Expected error when creating map with wrong type of keys")
	}
}

func TestMapSetAutoTypeErrorValues(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Int(2)}, []Type{Int(1), Bool(false)}, "", "", ""}
	err := t1.SetAutoType()
	if err == nil {
		t.Error("Expected error when creating map with wrong type of keys")
	}
}

func TestMapSetValueNonMap(t *testing.T) {
	t1 := NewMap()
	err := t1.SetValue(Int(0))
	if err == nil {
		t.Error("Expected error when setting value with Int")
	}
}

func TestMapSetErrKeys(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic when adding a bool as key in a map[test]int")
		}
	}()

	t1 := &Map{[]Type{}, []Type{}, "", "test", parser.Int}
	t1.Set(Bool(false), Int(0))
}

func TestMapSetErrValues(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic when adding a bool as value in a map[int]test")
		}
	}()

	t1 := &Map{[]Type{}, []Type{}, "", parser.Int, "test"}
	t1.Set(Int(0), Bool(false))
}

func TestAddKeyError(t *testing.T) {
	t1 := &Map{[]Type{}, []Type{}, "", "test", parser.Int}
	err := t1.AddKey(Int(0))
	if err == nil {
		t.Error("Expected error when adding int as key of map[test]int")
	}
}

func TestGetIndexErrorType(t *testing.T) {
	t1 := &Map{[]Type{}, []Type{}, "", parser.Int, parser.Int}
	_, err := t1.GetIndex(Bool(false))
	if err == nil {
		t.Error("Expected error when getting index of bool in map[int]int")
	}
}

func TestGetIndexNotFound(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{Int(1)}, "", parser.Int, parser.Int}
	_, err := t1.GetIndex(Int(-1))
	if err == nil {
		t.Error("Expected error when getting index of non existent key")
	}
}

func TestMapAddNonMap(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{Int(1)}, "", parser.Int, parser.Int}
	_, err := t1.Add(Int(0))
	if err == nil {
		t.Error("Expected error when adding int to map")
	}
}

func TestMapAddMapWrongType(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{Int(1)}, "test", parser.Int, parser.Int}
	t2 := &Map{[]Type{Bool(true)}, []Type{Int(1)}, "different type", parser.Bool, parser.Int}
	_, err := t1.Add(t2)
	if err == nil {
		t.Error("Expected error when adding int to map")
	}
}

func TestMapSubMapWrongType(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{Int(1)}, "test", parser.Int, parser.Int}
	t2 := &Map{[]Type{Bool(true)}, []Type{Int(1)}, "different type", parser.Bool, parser.Int}
	_, err := t1.Sub(t2)
	if err == nil {
		t.Error("Expected error when adding int to map")
	}
}

func TestMapMulErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Mul(Int(2))
	if err == nil {
		t.Error("Expected error when multiplying a map")
	}
}

func TestMapDivErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Div(Int(2))
	if err == nil {
		t.Error("Expected error when dividing a map")
	}
}

func TestMapModErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Mod(Int(2))
	if err == nil {
		t.Error("Expected error when getting remainder of a map")
	}
}

func TestMapDivEcErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.DivEc(Int(2))
	if err == nil {
		t.Error("Expected error when getting quotient of a map")
	}
}

func TestMapEqNonMap(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Eq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map  to int")
	}
}

func TestMapAndErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.And(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapOrErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Or(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapXorErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Xor(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapGtErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Gt(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapGtEqErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.GtEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapLwErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Lw(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapLwEqErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.LwEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapAppendErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Append(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map")
	}
}

func TestMapNotErr(t *testing.T) {
	t1 := NewMap()
	_, err := t1.Not()
	if err == nil {
		t.Error("Expected error when getting \"not\" of map")
	}
}
