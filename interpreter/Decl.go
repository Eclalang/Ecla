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
			env.ErrorHandle.HandleError(0, t.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(str)
	case lexer.BOOL:
		return NewMainBus(eclaType.NewBool(t.Value))
	case lexer.FLOAT:
		return NewMainBus(eclaType.NewFloat(t.Value))
	case lexer.CHAR:
		c, err := eclaType.NewChar(t.Value)
		if err != nil {
			env.ErrorHandle.HandleError(0, t.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(c)
	case "VAR":
		v, ok := env.GetVar(t.Value)
		if !ok {
			env.ErrorHandle.HandleError(0, t.StartPos(), "variable not found", errorHandler.LevelFatal)
		}
		return NewMainBus(v)
	case "NULL":
		return NewMainBus(eclaType.NewNull())
	default:
		env.ErrorHandle.HandleError(0, t.StartPos(), "Unknown type "+t.Type, errorHandler.LevelFatal)
		return NewNoneBus()
	}
}

// RunVariableDecl executes a parser.VariableDecl.
func RunVariableDecl(tree parser.VariableDecl, env *Env) {
	if tree.Value == nil {
		if env.CheckIfVarExistsInCurrentScope(tree.Name) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable already exists", errorHandler.LevelFatal)
			return
		}
		switch tree.Type {
		case parser.Int:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewInt("0"))
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.String:
			str, err := eclaType.NewString("")
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, str)
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Bool:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewBool("false"))
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Float:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewFloat("0.0"))
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		case parser.Any:
			val := eclaType.NewNull()
			v, err := eclaType.NewVar(tree.Name, tree.Type, val)
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		}
		if eclaType.IsList(tree.Type) {
			l, err := eclaType.NewList(tree.Type)
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, l)
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		} else if eclaType.IsMap(tree.Type) {
			m := eclaType.NewMap()
			m.SetType(tree.Type)
			v, err := eclaType.NewVar(tree.Name, tree.Type, m)
			if err != nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.SetVar(tree.Name, v)
		} else if _, ok := parser.Keywords[tree.Type]; ok {
			decl, ok := env.GetTypeDecl(tree.Type)
			if !ok {
				env.ErrorHandle.HandleError(0, tree.StartPos(), "unknown type: "+tree.Type, errorHandler.LevelFatal)
			}
			switch decl.(type) {
			case *eclaDecl.StructDecl:
				m := eclaType.NewStruct(decl.(*eclaDecl.StructDecl))
				m.SetType(tree.Type)
				v, err := eclaType.NewVar(tree.Name, tree.Type, m)
				if err != nil {
					env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
				}
				env.SetVar(tree.Name, v)
			}
		}
	} else {
		busCollection := RunTree(tree.Value, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable decl : MULTIPLE BUS IN RunVariableDecl", errorHandler.LevelFatal)
		}
		v, err := eclaType.NewVar(tree.Name, tree.Type, busCollection[0].GetVal())
		if err != nil {
			env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		if v.IsFunction() {
			if fn, ok := env.GetVar(tree.Name); ok {
				if fn.IsFunction() {
					fn2 := v.GetFunction()
					if slices.Contains(fn.GetFunction().GetTypes(), fn2.GetType()) {
						env.ErrorHandle.HandleError(0, 0, "Cannot overwrite this function", errorHandler.LevelFatal)
					}
					fn.GetFunction().AddOverload(fn2.Args[0], fn2.GetBody(), fn2.GetReturn())
				} else {
					env.ErrorHandle.HandleError(0, 0, "Cannot overload a non-function variable", errorHandler.LevelFatal)
				}
			} else {
				if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
					env.SetVar(tree.Name, v)
				} else {
					env.ErrorHandle.HandleError(0, 0, "Cannot reassign a variable", errorHandler.LevelFatal)
				}
			}
		} else {
			if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
				env.SetVar(tree.Name, v)
			} else {
				env.ErrorHandle.HandleError(0, 0, "Cannot reassign a variable", errorHandler.LevelFatal)
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
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable decl : MULTIPLE BUS IN RunVariableDecl", errorHandler.LevelFatal)
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
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = l.SetValue(values)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(l)
}

// RunFunctionDecl executes a parser.FunctionDecl.
func RunFunctionDecl(tree parser.FunctionDecl, env *Env) {
	declared, _ := env.Vars.Get(tree.Name)
	if !env.CheckIfVarExistsInCurrentScope(tree.Name) {
		fn := eclaType.NewFunction(tree.Name, tree.Prototype.Parameters, tree.Body, tree.Prototype.ReturnTypes)
		env.SetFunction(tree.Name, fn)
	} else {
		if !declared.IsFunction() {
			env.ErrorHandle.HandleError(tree.StartPos(),
				0,
				"var "+tree.Name+" already exists.",
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
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable decl : MULTIPLE BUS IN RunVariableDecl", errorHandler.LevelFatal)
		}
		values = append(values, busCollection[0].GetVal())
	}
	for _, k := range tree.Keys {
		busCollection := RunTree(k, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable decl : MULTIPLE BUS IN RunVariableDecl", errorHandler.LevelFatal)
		}
		keys = append(keys, busCollection[0].GetVal())
	}
	m := eclaType.NewMap()
	m.Keys = keys
	m.Values = values
	err := m.SetAutoType()
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(m)
}

func RunStructDecl(tree parser.StructDecl, env *Env) {
	strdecl := eclaDecl.NewStructDecl(tree)
	env.AddTypeDecl(strdecl)
}
