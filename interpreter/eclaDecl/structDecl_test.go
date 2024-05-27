package eclaDecl

import (
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestNewStructDecl(t *testing.T) {
	tree := parser.StructDecl{
		Name: "test",
		Fields: []parser.StructField{
			{
				Name: "field1",
				Type: "int",
			},
			{
				Name: "field2",
				Type: "string",
			},
		},
	}
	strdecl := NewStructDecl(tree)
	if strdecl.Name != "test" {
		t.Error("Expected test, got ", strdecl.Name)
	}
	if strdecl.Fields["field1"] != "int" {
		t.Error("Expected int, got ", strdecl.Fields["field1"])
	}
	if strdecl.Fields["field2"] != "string" {
		t.Error("Expected string, got ", strdecl.Fields["field2"])
	}
	if strdecl.Order[0] != "field1" {
		t.Error("Expected field1, got ", strdecl.Order[0])
	}
	if strdecl.Order[1] != "field2" {
		t.Error("Expected field2, got ", strdecl.Order[1])
	}
}

func TestStructDecl_GetFieldsInOrder(t *testing.T) {

	strdecl := StructDecl{
		Fields: map[string]string{
			"field1": "int",
			"field2": "string",
		},
		Order: []string{"field1", "field2"},
	}
	fields := strdecl.GetFieldsInOrder()
	if fields[0].Name != "field1" {
		t.Error("Expected field1, got ", fields[0].Name)
	}
	if fields[0].Type != "int" {
		t.Error("Expected int, got ", fields[0].Type)
	}
	if fields[1].Name != "field2" {
		t.Error("Expected field2, got ", fields[1].Name)
	}
	if fields[1].Type != "string" {
		t.Error("Expected string, got ", fields[1].Type)
	}

}

func TestStructDecl_GetName(t *testing.T) {
	strdecl := StructDecl{
		Name: "test",
	}
	if strdecl.GetName() != "test" {
		t.Error("Expected test, got ", strdecl.GetName())
	}
}
