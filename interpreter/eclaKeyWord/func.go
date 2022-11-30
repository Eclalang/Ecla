package eclaKeyWord

import "github.com/tot0p/Ecla/interpreter/eclaType"

type Function struct {
	Name  string
	Args  []eclaType.Type
	Scope []eclaType.Type
	Code  any
}

func (f *Function) GetType() eclaType.String {
	return "Function"
}

func (f *Function) Call() {
	//TODO EXECUTE CODE
}
