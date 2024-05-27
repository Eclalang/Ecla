package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
)

type Function struct {
	Name            string
	Args            [][]parser.FunctionParams
	Body            map[string][]parser.Node
	Return          map[string][]string
	lastIndexOfArgs int
}

// Function Method for interface Type

func (f *Function) GetValue() any {
	return f
}

func (f *Function) SetValue(value any) error {
	return errors.New("cannot set value of function")
}

func (f *Function) String() string {
	return "function"
}

func (f *Function) GetString() String {
	return "function"
}

func (f *Function) GetType() string {
	typ := "function("
	length := len(f.Args[0])
	for i := 0; i < length-1; i++ {
		typ += f.Args[0][i].Type
		typ += ","
	}
	if length > 0 {
		typ += f.Args[0][length-1].Type
	}
	typ += ")"
	key := generateArgsString(f.Args[0])
	length = len(f.Return[key])
	if length > 0 {
		typ += "("
		for i := 0; i < length-1; i++ {
			typ += f.Return[key][i]
			typ += ","
		}
		typ += f.Return[key][length-1] + ")"
	}

	return typ
}

func (f *Function) GetIndex(number Type) (*Type, error) {
	return nil, errors.New("cannot get index of function")
}

func (f *Function) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add " + f.String() + " and " + other.String())
}

func (f *Function) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract " + other.String() + " from " + f.String())
}

func (f *Function) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply " + f.String() + " by " + other.String())
}

func (f *Function) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide " + f.String() + " by " + other.String())
}

func (f *Function) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot get remainder of " + f.String() + " by " + other.String())
}

func (f *Function) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of " + f.String() + " by " + other.String())
}

func (f *Function) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) And(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) Or(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) Not() (Type, error) {
	return nil, errors.New("cannot \"not\" a function")
}

func (f *Function) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
}

func (f *Function) Append(other Type) (Type, error) {
	return nil, errors.New("cannot add " + other.String() + " to " + f.String())
}

func (f *Function) IsNull() bool {
	return false
}

// End of Function Method for interface Type

func generateArgsString(args []parser.FunctionParams) string {
	result := ""
	for _, arg := range args {
		result += arg.Type
	}
	return result
}

func NewFunction(Name string, args []parser.FunctionParams, body []parser.Node, ret []string) *Function {
	var argsList [][]parser.FunctionParams
	argsList = append(argsList, args)
	argsString := generateArgsString(args)
	var returnMap = make(map[string][]string)
	returnMap[argsString] = ret
	var bodyMap = make(map[string][]parser.Node)
	bodyMap[argsString] = body
	return &Function{
		Name:            Name,
		Args:            argsList,
		Body:            bodyMap,
		Return:          returnMap,
		lastIndexOfArgs: 0,
	}
}

func NewAnonymousFunction(args []parser.FunctionParams, body []parser.Node, ret []string) *Function {
	var argsList [][]parser.FunctionParams
	argsList = append(argsList, args)
	argsString := generateArgsString(args)
	var returnMap = make(map[string][]string)
	returnMap[argsString] = ret
	var bodyMap = make(map[string][]parser.Node)
	bodyMap[argsString] = body
	return &Function{
		Name:   "",
		Args:   argsList,
		Body:   bodyMap,
		Return: returnMap,
	}
}

// Method for function

func (f *Function) AddOverload(args []parser.FunctionParams, body []parser.Node, ret []string) {
	f.Args = append(f.Args, args)
	key := generateArgsString(args)
	f.Body[key] = body
	f.Return[key] = ret
}

func (f *Function) Override(args []parser.FunctionParams, body []parser.Node, ret []string) error {
	index := -1
	var isSameArgs bool
	for i, arg := range f.Args {
		for j, argTyp := range arg {
			isSameArgs = true
			if len(arg) != len(args) || argTyp.Type != args[j].Type {
				isSameArgs = false
				break
			}
		}
		if isSameArgs {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("cannot override a prototype that was not implemented")
	}
	key := generateArgsString(args)
	f.Args[index] = args
	f.Body[key] = body
	f.Return[key] = ret
	return nil
}

func (f *Function) GetBody() []parser.Node {
	key := generateArgsString(f.Args[f.lastIndexOfArgs])
	return f.Body[key]
}

func (f *Function) GetReturn() []string {
	key := generateArgsString(f.Args[f.lastIndexOfArgs])
	return f.Return[key]
}

func (f *Function) GetIndexOfArgs(args []Type) int {
	l := len(args)
	cursor := -1
	maxNbAny := -1
	for i, arg := range f.Args {
		if l != len(arg) {
			continue
		}
		var nbAny int
		var isGoodArgs = true
		for j, typ := range args {
			if arg[j].Type == parser.Any {
				nbAny++
			} else if typ.GetType() != arg[j].Type {
				isGoodArgs = false
				break
			}
		}
		if (maxNbAny == -1 && nbAny > 0) || nbAny < maxNbAny {
			cursor = i
			maxNbAny = nbAny
		} else if isGoodArgs {
			return i
		}
	}
	return cursor
}

func (f *Function) TypeAndNumberOfArgsIsCorrect(args []Type, StructDecl []eclaDecl.TypeDecl) (bool, map[string]*Var) {
	indexOfArgs := f.GetIndexOfArgs(args)
	f.lastIndexOfArgs = indexOfArgs
	if indexOfArgs == -1 {
		return false, nil
	}
	var i int = 0
	var argsType = make(map[string]*Var)
	for _, arg := range f.Args[indexOfArgs] {
		paramName := arg.Name
		paramType := arg.Type
		elem := args[i]
		switch elem.(type) {
		case *Var:
			elem = elem.(*Var).Value
		}
		tp := elem.GetType()
		if paramType == parser.Any {
			tp = parser.Any
		}
		if tp != paramType {
			isImplemented := false
			for _, decl := range StructDecl {
				if decl.GetName() == paramType {
					isImplemented = true
					break
				}
			}
			if !isImplemented {
				return false, nil
			}
		}
		v, err := NewVar(paramName, tp, elem)
		if err != nil {
			panic(err)
		}
		argsType[paramName] = v
		i++
	}
	return true, argsType
}

func (f *Function) CheckReturn(ret []Type, StructDecl []eclaDecl.TypeDecl) bool {

	key := generateArgsString(f.Args[f.lastIndexOfArgs])
	if len(f.Return[key]) != len(ret) {
		return false
	}
	var i int = 0
	for _, r := range f.Return[key] {
		elem := ret[i]
		switch elem.(type) {
		case *Var:
			elem = elem.(*Var).Value
		}
		i++
		if r == parser.Any {
			continue
		}
		tp := elem.GetType()
		if tp != r {
			isImplemented := false
			for _, decl := range StructDecl {
				if decl.GetName() == r {
					isImplemented = true
					break
				}
			}
			if !isImplemented {
				return false
			}
		}
	}
	return true
}

func (f *Function) GetTypes() []string {
	var types []string
	for i := 0; i < len(f.Args); i++ {
		typ := "function("
		length := len(f.Args[i])
		for j := 0; j < length-1; j++ {
			typ += f.Args[i][j].Type
			typ += ","
		}
		if length > 0 {
			typ += f.Args[i][length-1].Type
		}
		typ += ")"
		key := generateArgsString(f.Args[i])
		length = len(f.Return[key])
		if length > 0 {
			typ += "("
			for j := 0; j < length-1; j++ {
				typ += f.Return[key][j]
				typ += ","
			}
			typ += f.Return[key][length-1] + ")"
		}
		types = append(types, typ)
	}
	return types
}

func (f *Function) GetSize() int {
	return utils.Sizeof(f)
}

func (f *Function) Len() (int, error) {
	return -1, errors.New("cannot get length of function")
}
