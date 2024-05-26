package eclaType

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
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
