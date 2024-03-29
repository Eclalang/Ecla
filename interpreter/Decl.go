package interpreter

import (
	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"slices"
)

// New returns a new eclaType.Type from a parser.Literal.
func New(t parser.Literal, env *Env) *Bus {
	switch t.Type {
	case lexer.INT:
		return NewMainBus(eclaType.NewInt(t.Value))
	case lexer.STRING:
		str, err := eclaType.NewString(t.Value)
		if err != nil {
			env.ErrorHandle.HandleError(t.StartLine(), t.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(str)
	case lexer.BOOL:
		b, err := eclaType.NewBool(t.Value)
		if err != nil {
			env.ErrorHandle.HandleError(t.StartLine(), t.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(b)
	case lexer.FLOAT:
		return NewMainBus(eclaType.NewFloat(t.Value))
	case lexer.CHAR:
		c, err := eclaType.NewChar(t.Value)
		if err != nil {
			env.ErrorHandle.HandleError(t.StartLine(), t.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(c)
	case "VAR":
		v, ok := env.GetVar(t.Value)
		if !ok {
			env.ErrorHandle.HandleError(t.StartLine(), t.StartPos(), "variable "+t.Value+" not found", errorHandler.LevelFatal)
		}
		return NewMainBus(v)
	case "NULL":
		return NewMainBus(eclaType.NewNull())
	default:
		env.ErrorHandle.HandleError(t.StartLine(), t.StartPos(), "Unknown type "+t.Type, errorHandler.LevelFatal)
		return NewNoneBus()
	}
}

// RunVariableDecl executes a parser.VariableDecl.
func RunVariableDecl(tree parser.VariableDecl, env *Env) {
	if tree.Value == nil {
		if env.CheckIfVarExistsInCurrentScope(tree.Name) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "variable "+tree.Name+" already exists", errorHandler.LevelFatal)
			return
		}
		switch tree.Type {
		case parser.Int:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewInt("0"))
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.String:
			str, err := eclaType.NewString("")
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, str)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Bool:
			b, err := eclaType.NewBool("false")
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, b)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Float:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewFloat("0.0"))
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Any:
			val, e := eclaType.NewAnyEmpty()
			if e != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), e.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, val)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Char:
			c, err := eclaType.NewChar("")
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, c)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		}
		if eclaType.IsList(tree.Type) {
			l, err := eclaType.NewList(tree.Type)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, l)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		} else if eclaType.IsMap(tree.Type) {
			m := eclaType.NewMap()
			m.SetType(tree.Type)
			v, err := eclaType.NewVar(tree.Name, tree.Type, m)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		} else if decl, ok := env.GetTypeDecl(tree.Type); ok {
			switch decl.(type) {
			case *eclaDecl.StructDecl:
				m := eclaType.NewStruct(decl.(*eclaDecl.StructDecl))
				m.SetType(tree.Type)
				v, err := eclaType.NewVar(tree.Name, tree.Type, m)
				if err != nil {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
				}
				env.SetVar(tree.Name, v)
			}
		}
	} else {
		busCollection := RunTree(tree.Value, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunVariableDecl.\nPlease open issue", errorHandler.LevelFatal)
		}
		v, err := eclaType.NewVar(tree.Name, tree.Type, busCollection[0].GetVal())
		if err != nil {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		if v.IsFunction() {
			if fn, ok := env.GetVar(tree.Name); ok {
				if fn.IsFunction() {
					fn2 := v.GetFunction()
					if slices.Contains(fn.GetFunction().GetTypes(), fn2.GetType()) {
						env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Cannot overwrite this function "+tree.Name, errorHandler.LevelFatal)
					}
					fn.GetFunction().AddOverload(fn2.Args[0], fn2.GetBody(), fn2.GetReturn())
				} else {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Cannot overload a non-function variable "+tree.Name, errorHandler.LevelFatal)
				}
			} else {
				if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
					env.SetVar(tree.Name, v)
				} else {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Cannot reassign a variable "+tree.Name, errorHandler.LevelFatal)
				}
			}
		} else {
			if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
				env.SetVar(tree.Name, v)
			} else {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Cannot reassign a variable "+tree.Name, errorHandler.LevelFatal)
			}
		}
	}
}

// RunArrayLiteral executes a parser.ArrayLiteral.
func RunArrayLiteral(tree parser.ArrayLiteral, env *Env) *Bus {
	var values []eclaType.Type
	for _, v := range tree.Values {
		busCollection := RunTree(v, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunArrayLiteral.\nPlease open issue", errorHandler.LevelFatal)
		}
		values = append(values, busCollection[0].GetVal())
	}
	// Construct type of list
	var typ string
	if len(values) == 0 {
		typ = "empty"
	} else {
		typ = "[]" + values[0].GetType()
	}
	l, err := eclaType.NewList(typ)
	if err != nil {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = l.SetValue(values)
	if err != nil {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(l)
}

// RunFunctionDecl executes a parser.FunctionDecl.
func RunFunctionDecl(tree parser.FunctionDecl, env *Env) {
	declared, _ := env.Vars.Get(tree.Name)
	if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
		fn := eclaType.NewFunction(tree.Name, tree.Prototype.Parameters, tree.Body, tree.Prototype.ReturnTypes)
		err := env.SetFunction(tree.Name, fn)
		if err != nil {
			env.ErrorHandle.HandleError(tree.StartPos(), tree.StartLine(), err.Error(), errorHandler.LevelFatal)
		}
	} else {
		if !declared.IsFunction() {
			env.ErrorHandle.HandleError(tree.StartPos(),
				0,
				"Variable "+tree.Name+" already exists.",
				errorHandler.LevelFatal)
		} else {
			declared.GetFunction().AddOverload(tree.Prototype.Parameters, tree.Body, tree.Prototype.ReturnTypes)
		}
	}
}

// RunMapLiteral executes a parser.MapLiteral.
func RunMapLiteral(tree parser.MapLiteral, env *Env) *Bus {
	var keys []eclaType.Type
	var values []eclaType.Type
	for _, v := range tree.Values {
		busCollection := RunTree(v, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunMapLiteral.\nPlease open issue", errorHandler.LevelFatal)
		}
		values = append(values, busCollection[0].GetVal())
	}
	for _, k := range tree.Keys {
		busCollection := RunTree(k, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunMapLiteral.\nPlease open issue", errorHandler.LevelFatal)
		}
		keys = append(keys, busCollection[0].GetVal())
	}
	m := eclaType.NewMap()
	m.Keys = keys
	m.Values = values
	err := m.SetAutoType()
	if err != nil {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(m)
}

func RunStructDecl(tree parser.StructDecl, env *Env) {
	strdecl := eclaDecl.NewStructDecl(tree)
	env.AddTypeDecl(strdecl)
}
