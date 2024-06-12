package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"testing"
)

func TestNewStruct(t *testing.T) {
	m := make(map[string]string)
	fName := "field1"
	vName := "value1"
	m[fName] = vName
	sName := "struct"

	decl := &eclaDecl.StructDecl{m, []string{fName}, sName}
	t1 := NewStruct(decl)
	if t1.Typ != sName {
		t.Errorf("Expected %s, got %s", sName, t1.Typ)
		return
	}
	if t1.Definition != decl {
		t.Errorf("Expected %s, got %s", decl, t1.Definition)
		return
	}
	if t1.Fields == nil {
		t.Error("Expected map, got nil")
		return
	}
	if len(t1.Fields) != 0 {
		t.Error("The maps do not match")
		return
	}
}

func TestStructAddField(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{}, []string{"field0"}, "testStruct"}
	s := &Struct{map[string]*Type{}, "testType", decl}

	i := Int(42)
	s.AddField(0, i)

	if len(s.Fields) == 0 {
		t.Error("Field was not added")
	}
	expected := *(s.Fields[decl.Order[0]])
	if expected != i {
		t.Errorf("Expected %d, got %d", i, expected)
	}
}

func TestStructVerify(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{}, []string{"field0"}, "testStruct"}
	var v Type = Int(42)
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	err := s.Verify()
	if err != nil {
		t.Error(err)
	}
}

func TestStructGetValue(t *testing.T) {
	s := &Struct{nil, "", nil}
	result := s.GetValue()
	if result != s {
		t.Error("The two structs do not match")
	}
}

func TestStructGetType(t *testing.T) {
	expected := "testing get type"
	s := &Struct{nil, expected, nil}
	result := s.GetType()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStructString(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	expected := "{42, test}"
	result := s.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStructGetString(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	expected := String("{42, test}")
	result := s.GetString()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStructSetType(t *testing.T) {
	oldType := "old type"
	s := &Struct{nil, oldType, nil}
	newType := "new type"
	s.SetType(newType)
	if (s.Typ == oldType) || (s.Typ != newType) {
		t.Errorf("Expected %s, got %s", newType, oldType)
	}
}

func TestStructGet(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	result, err := s.Get("field0")
	if err != nil {
		t.Error(err)
	}
	if result != i {
		t.Errorf("Expected %d, got %s", i, result)
	}
}

func TestStructGetIndex(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	result, err := s.GetIndex(String("field0"))
	if err != nil {
		t.Error(err)
	}
	if *result != i {
		t.Errorf("Expected %d, got %s", i, *result)
	}
}

func TestStructIsNull(t *testing.T) {
	s := &Struct{nil, "", nil}
	result := s.IsNull()
	if result {
		t.Error("Expected false, got true")
	}
}

func TestStructGetField(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	result := s.GetField("field0")
	if *result != i {
		t.Errorf("Expected %d, got %s", i, *result)
	}
}

func TestStructGetFieldFalse(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{}, ""}
	s := &Struct{map[string]*Type{}, "", decl}
	result := s.GetField("field0")
	if result != nil {
		t.Errorf("Expected nil, got %s", *result)
	}
}

func TestStructLen(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	result, err := s.Len()
	if err != nil {
		t.Error(err)
	}
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

func TestStructGetSize(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	s := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	expected := utils.Sizeof(s)
	result := s.GetSize()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestIsStruct(t *testing.T) {
	str := "struct test"
	result := IsStruct(str)
	if !result {
		t.Errorf("%s should be a valid struct type", str)
	}
}

func TestIsStructFalseLen(t *testing.T) {
	str := "s"
	result := IsStruct(str)
	if result {
		t.Errorf("%s should not be a valid struct type", str)
	}
}

func TestIsStructFalse(t *testing.T) {
	str := "not a valid type"
	result := IsStruct(str)
	if result {
		t.Errorf("%s should not be a valid struct type", str)
	}
}

func TestSetValueStruct(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{"field0", "field1"}, ""}
	var i Type = Int(42)
	var str Type = String("test")
	expected := &Struct{map[string]*Type{"field0": &i, "field1": &str}, "", decl}
	s := &Struct{nil, "", nil}
	err := s.SetValue(expected)
	if err != nil {
		t.Error(err)
		return
	}

	if s.Typ != expected.Typ || s.Definition != expected.Definition || len(s.Fields) != len(expected.Fields) {
		t.Errorf("Expected %s, got %s", expected, s)
		return
	}
	for k, arg := range expected.Fields {
		if s.Fields[k] != arg {
			t.Error("The fields are not the same")
			return
		}
	}
}

// Tests struct errors

func TestStructVerifyError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{}, []string{"field0"}, "testStruct"}
	s := &Struct{map[string]*Type{}, "testStruct", decl}
	err := s.Verify()
	if err == nil {
		t.Error("Expected error when the number of fields did not match")
	}
}

func TestStructGetError(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{}, ""}
	s := &Struct{map[string]*Type{}, "", decl}
	_, err := s.Get("field0")
	if err == nil {
		t.Error("Expected error when getting non existent field")
	}
}

func TestStructGetIndexErrorType(t *testing.T) {
	s := &Struct{nil, "", nil}
	_, err := s.GetIndex(Int(0))
	if err == nil {
		t.Error("Expected error when getting index with non string arg")
	}
}

func TestStructGetIndexGetError(t *testing.T) {
	decl := &eclaDecl.StructDecl{nil, []string{}, ""}
	s := &Struct{map[string]*Type{}, "", decl}
	_, err := s.GetIndex(String("field0"))
	if err == nil {
		t.Error("Expected error when getting non existent field")
	}
}

func TestStructAddError(t *testing.T) {
	s := &Struct{nil, "", nil}
	_, err := s.Add(Int(0))
	if err == nil {
		t.Error("Expected error when adding int to struct")
	}
}

func TestStructSubError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Sub(Int(0))
	if err == nil {
		t.Error("Expected error when subtracting int from struct")
	}
}

func TestStructMulError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Mul(Int(0))
	if err == nil {
		t.Error("Expected error when multiplying int with struct")
	}
}

func TestStructDivError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Div(Int(0))
	if err == nil {
		t.Error("Expected error when dividing struct by int")
	}
}

func TestStructDivEcError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.DivEc(Int(0))
	if err == nil {
		t.Error("Expected error when getting quotient of struct")
	}
}

func TestStructModError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Mod(Int(0))
	if err == nil {
		t.Error("Expected error when getting remainder of struct")
	}
}

func TestStructEqError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Eq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructNotEqError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.NotEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructAndError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.And(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructOrError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Or(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructXorError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Xor(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructGtError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Gt(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructGtEqError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.GtEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructLwError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Lw(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructLwEqError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.LwEq(Int(0))
	if err == nil {
		t.Error("Expected error when comparing struct to int")
	}
}

func TestStructNotError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Not()
	if err == nil {
		t.Error("Expected error when getting \"not\" of struct")
	}
}

func TestStructAppendError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	_, err := s.Append(Int(0))
	if err == nil {
		t.Error("Expected error when appending to struct")
	}
}

func TestStructSetValueError(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	err := s.SetValue(Int(0))
	if err == nil {
		t.Error("Expected error when setting value of struct with int")
	}
}

func TestStructSetErrorNotExist(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	err := s.Set("wrong", Int(0))
	if err == nil {
		t.Error("Expected error when appending to struct")
	}
}

func TestStructSetErrorType(t *testing.T) {
	decl := &eclaDecl.StructDecl{map[string]string{"field0": "string"}, []string{"field0"}, "testStruct"}
	var v Type = String("test")
	s := &Struct{map[string]*Type{"field0": &v}, "testStruct", decl}
	err := s.Set("field0", Int(0))
	if err == nil {
		t.Error("Expected error when appending to struct")
	}
}
