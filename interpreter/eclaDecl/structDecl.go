package eclaDecl

import "github.com/Eclalang/Ecla/parser"

// StructDecl is the struct declaration.
type StructDecl struct {
	Fields map[string]string
	Order  []string
	Name   string
}

func NewStructDecl(tree parser.StructDecl) *StructDecl {
	var strdecl = StructDecl{
		Fields: make(map[string]string),
		Order:  make([]string, 0),
		Name:   tree.Name,
	}

	for _, field := range tree.Fields {
		strdecl.Fields[field.Name] = field.Type
		strdecl.Order = append(strdecl.Order, field.Name)
	}
	return &strdecl
}

func (s *StructDecl) GetFieldsInOrder() []Field {
	var fields []Field
	for _, field := range s.Order {
		fields = append(fields, Field{Name: field, Type: s.Fields[field]})
	}
	return fields
}

func (s *StructDecl) GetName() string {
	return s.Name
}
