package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/utils"
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

func TestMapSetValue(t *testing.T) {
	t1 := NewMap()
	t2 := &Map{[]Type{Int(1)}, []Type{String("One")}, "test", parser.Int, parser.String}
	err := t1.SetValue(t2)
	if err != nil {
		t.Error(err)
	}
	if t1.Typ != t2.Typ || t1.TypVal != t2.TypVal || t1.TypKey != t2.TypKey {
		t.Error("The types do not match")
	}
	if len(t1.Keys) != len(t2.Keys) {
		t.Error("There is a different number of keys")
	}
	for i, elem := range t1.Keys {
		if elem != t2.Keys[i] {
			t.Error("The keys do not match")
		}
	}
	if len(t1.Values) != len(t2.Values) {
		t.Error("There is a different number of values")
	}
	for i, elem := range t1.Values {
		if elem != t2.Values[i] {
			t.Error("The values do not match")
		}
	}
}

func TestMapSetValueAny(t *testing.T) {
	t1 := NewMap()
	t2 := &Map{[]Type{Int(1)}, []Type{String("One")}, "test", parser.Int, parser.String}
	a := &Any{t2, t2.Typ}
	err := t1.SetValue(a)
	if err != nil {
		t.Error(err)
	}
	if t1.Typ != t2.Typ || t1.TypVal != t2.TypVal || t1.TypKey != t2.TypKey {
		t.Error("The types do not match")
	}
	if len(t1.Keys) != len(t2.Keys) {
		t.Error("There is a different number of keys")
	}
	for i, elem := range t1.Keys {
		if elem != t2.Keys[i] {
			t.Error("The keys do not match")
		}
	}
	if len(t1.Values) != len(t2.Values) {
		t.Error("There is a different number of values")
	}
	for i, elem := range t1.Values {
		if elem != t2.Values[i] {
			t.Error("The values do not match")
		}
	}
}

func TestMapString(t *testing.T) {
	t1 := &Map{[]Type{Bool(true), Bool(false)}, []Type{Int(1), Int(0)}, "", "", ""}
	expected := "{true: 1, false: 0}"
	result := t1.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapGetString(t *testing.T) {
	t1 := &Map{[]Type{Bool(true), Bool(false)}, []Type{Int(1), Int(0)}, "", "", ""}
	expected := String("{true: 1, false: 0}")
	result := t1.GetString()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapGetType(t *testing.T) {
	t1 := NewMap()
	expected := "test"
	t1.Typ = expected
	result := t1.GetType()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapGetTypeOfKeyAndValue(t *testing.T) {
	str := "map[int]string"
	resultKey, resultVal := GetTypeOfKeyAndValue(str)
	if resultKey != "int" {
		t.Errorf("expected key: %s, got %s", "int", resultKey)
	}
	if resultVal != "string" {
		t.Errorf("expected val: %s, got %s", "string", resultVal)
	}
}

func TestMapEqTrue(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(true) {
		t.Error("Expected true, got false")
	}
}

func TestMapEqTrueAny(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Any{&Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}, "test"}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(true) {
		t.Error("Expected true, got false")
	}
}

func TestMapEqTrueVar(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Var{"var", &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(true) {
		t.Error("Expected true, got false")
	}
}

func TestMapEqFalseLength(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Int(2)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(false) {
		t.Error("Expected false, got true")
	}
}

func TestMapEqFalseKey(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(false) {
		t.Error("Expected false, got true")
	}
}

func TestMapEqFalseValue(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("fail")}, "type", "keys", "values"}
	b, err := t1.Eq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(false) {
		t.Error("Expected false, got true")
	}
}

func TestMapNotEqTrue(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Int(2)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	b, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(true) {
		t.Error("Expected true, got false")
	}
}

func TestMapNotEqFalse(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	t2 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	b, err := t1.NotEq(t2)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(false) {
		t.Error("Expected false, got true")
	}
}

func TestMapIsNull(t *testing.T) {
	t1 := NewMap()
	if t1.IsNull() {
		t.Error("Expected true, got false")
	}
}

func TestIsMapTrue(t *testing.T) {
	str := "map[int]string"
	if !IsMap(str) {
		t.Error("Expected true, got false")
	}
}

func TestIsMapFalseLength(t *testing.T) {
	str := "map"
	if IsMap(str) {
		t.Error("Expected false, got true")
	}
}

func TestIsMapFalseType(t *testing.T) {
	str := "123456"
	if IsMap(str) {
		t.Error("Expected false, got true")
	}
}

func TestMapGetSize(t *testing.T) {
	t1 := &Map{[]Type{Int(1)}, []Type{String("test")}, "type", "keys", "values"}
	expected := utils.Sizeof(t1)
	result := t1.GetSize()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapLen(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Int(2), Int(3)}, []Type{}, "", "", ""}
	expected := 3
	result, err := t1.Len()
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapSetType(t *testing.T) {
	str := "map[int]string"
	t1 := NewMap()
	t1.SetType(str)
	if t1.Typ != str {
		t.Error("Error in type")
	}
	if t1.TypKey != "int" {
		t.Error("Error in key type")
	}
	if t1.TypVal != "string" {
		t.Error("Error in value type")
	}
}

func TestMapSetNewKey(t *testing.T) {
	t1 := &Map{[]Type{}, []Type{}, "map[int]string", parser.Int, parser.String}
	t1.Set(Int(0), String("0"))
	if len(t1.Values) == 0 || len(t1.Keys) == 0 {
		t.Error("Value was not added")
	}
}

func TestMapSetReplaceKey(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("zero")}, "map[int]string", parser.Int, parser.String}
	t1.Set(Int(0), String("0"))
	if len(t1.Values) == 0 || len(t1.Keys) == 0 {
		t.Error("Value was not added")
	}
}

func TestMapAddKey(t *testing.T) {
	t1 := &Map{[]Type{}, []Type{}, "map[int]string", parser.Int, parser.String}
	err := t1.AddKey(Int(0))
	if err != nil {
		t.Error(err)
	}
	if len(t1.Keys) != 1 || len(t1.Values) != 1 {
		t.Error("The key was not added properly")
	}
	if t1.Keys[0] != Int(0) {
		t.Errorf("Expected %s, got %s", Int(0), t1.Keys)
	}
	if !t1.Values[0].IsNull() {
		t.Errorf("Expected null, got %s", t1.Values[0])
	}
}

func TestMapGetTrue(t *testing.T) {
	key := Int(0)
	expected := String("zero")
	t1 := &Map{[]Type{key}, []Type{expected}, "map[int]string", parser.Int, parser.String}
	result, b := t1.Get(key)
	if b != true {
		t.Error("Should have found the value, but didn't")
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapGetFalse(t *testing.T) {
	key := Int(0)
	expected := String("zero")
	t1 := &Map{[]Type{key}, []Type{expected}, "map[int]string", parser.Int, parser.String}
	_, b := t1.Get(Int(1))
	if b != false {
		t.Error("Should not have found the value, but did")
	}
}

func TestMapGetIndex(t *testing.T) {
	key := Int(0)
	expected := String("zero")
	t1 := &Map{[]Type{key}, []Type{expected}, "map[int]string", parser.Int, parser.String}
	result, err := t1.GetIndex(key)
	if err != nil {
		t.Error("Should have found the value, but didn't")
	}
	if *result != expected {
		t.Errorf("Expected %s, got %s", expected, *result)
	}
}

func TestMapGetKey(t *testing.T) {
	expected := Int(0)
	value := String("zero")
	t1 := &Map{[]Type{expected}, []Type{value}, "map[int]string", parser.Int, parser.String}
	result, b := t1.GetKey(value)
	if !b {
		t.Error("Should have found the value, but didn't")
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapGetKeyFalse(t *testing.T) {
	key := Int(0)
	value := String("zero")
	t1 := &Map{[]Type{key}, []Type{value}, "map[int]string", parser.Int, parser.String}
	_, b := t1.GetKey(Int(1))
	if b {
		t.Error("Should not have found the value, but did")
	}
}

func TestMapAddMap(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	t2 := &Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}
	expected := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}

	b, e := result.Eq(expected)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapAddVar(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	t2 := &Var{"test", &Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}}
	expected := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}

	b, e := result.Eq(expected)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapAddAny(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	t2 := &Any{&Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}, "test"}
	expected := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}

	b, e := result.Eq(expected)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapAddString(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	t2 := String("test")
	expected := String("{0: 0}test")
	result, err := t1.Add(t2)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapDelete(t *testing.T) {
	t1 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	expected := NewMap()
	t1.Delete(Int(0))
	b, err := t1.Eq(expected)
	if err != nil {
		t.Error(err)
	}
	if b != Bool(true) {
		t.Error("The values were not removed")
	}
}

func TestMapSub(t *testing.T) {
	t1 := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	t2 := &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}
	expected := &Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	b, e := result.Eq(t2)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapSubVar(t *testing.T) {
	t1 := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	t2 := &Var{"test", &Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}}
	expected := &Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	b, e := result.Eq(t2)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapSubAny(t *testing.T) {
	t1 := &Map{[]Type{Int(0), Int(1)}, []Type{String("0"), String("1")}, "map[int]string", parser.Int, parser.String}
	t2 := &Any{&Map{[]Type{Int(0)}, []Type{String("0")}, "map[int]string", parser.Int, parser.String}, "test"}
	expected := &Map{[]Type{Int(1)}, []Type{String("1")}, "map[int]string", parser.Int, parser.String}
	result, err := t1.Sub(t2)
	if err != nil {
		t.Error(err)
	}
	b, e := result.Eq(t2)
	if e != nil {
		t.Error(e)
	}
	if b != Bool(true) {
		t.Errorf("Expected %s, got %s", expected, result)
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

func TestMapNotEqErr(t *testing.T) {
	t1 := &Map{[]Type{Int(1), Int(2)}, []Type{String("test")}, "type", "keys", "values"}
	_, err := t1.NotEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing map to int")
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
