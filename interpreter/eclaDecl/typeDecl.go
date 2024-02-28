package eclaDecl

type TypeDecl interface {
	GetName() string
	GetFieldsInOrder() []Field
}

type Field struct {
	Name string
	Type string
}
