package eclaKeyWord

import "github.com/tot0p/Ecla/interpreter/eclaType"

type For struct {
	Condition any
	Body      any
	Scope     []eclaType.Type
}

func NewFor(condition any, body any, scope []eclaType.Type) *For {
	return &For{Condition: condition, Body: body, Scope: scope}
}

func (f *For) Execute() error {
	return nil
}
