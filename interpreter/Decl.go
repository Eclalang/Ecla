package interpreter

import (
	"errors"
	"fmt"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"github.com/Eclalang/Ecla/errorHandler"
)

// New returns a new eclaType.Type from a parser.Literal.
func New(t parser.Literal, env *Env) *Bus {
	switch t.Type {
	case lexer.INT:
		return NewMainBus(eclaType.NewInt(t.Value))
	case lexer.STRING:
		return NewMainBus(eclaType.NewString(t.Value))
	case lexer.BOOL:
		return NewMainBus(eclaType.NewBool(t.Value))
	case lexer.FLOAT:
		return NewMainBus(eclaType.NewFloat(t.Value))
	case "VAR":
		v, ok := env.GetVar(t.Value)
		if !ok {
			panic(errors.New("variable not found"))
		}
		return NewMainBus(v)
	case "NULL":
		return NewMainBus(eclaType.NewNull())
	default:
		panic("Unknown type")
		return NewNoneBus()
	}
}

// RunVariableDecl executes a parser.VariableDecl.
func RunVariableDecl(tree parser.VariableDecl, env *Env) {
	if tree.Value == nil {
		switch tree.Type {
		case parser.Int:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewInt("0"))
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		case parser.String:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewString(""))
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		case parser.Bool:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewBool("false"))
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		case parser.Float:
			v, err := eclaType.NewVar(tree.Name, tree.Type, eclaType.NewFloat("0.0"))
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		}
		if eclaType.IsList(tree.Type) {
			l, err := eclaType.NewList(tree.Type)
			if err != nil {
				panic(err)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, l)
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		} else if eclaType.IsMap(tree.Type) {
			m := eclaType.NewMap()
			m.SetType(tree.Type)
			v, err := eclaType.NewVar(tree.Name, tree.Type, m)
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		}
	} else {
		busCollection := RunTree(tree.Value, env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "variable decl : MULTIPLE BUS IN RunVariableDecl", errorHandler.LevelFatal)
		}
		v, err := eclaType.NewVar(tree.Name, tree.Type, busCollection[0].GetVal())
		if err != nil {
			panic(err)
		}
		env.SetVar(tree.Name, v)
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
	//Modif typ par tree.$type
	/*
		[14, 15] -> x1 [ donc [], 14 -> int donc []int
	*/
	var typ string
	if len(values) == 0 {
		typ = "empty"
	} else {
		typ = "[]" + values[0].GetType()
	}
	l, err := eclaType.NewList(typ)
	if err != nil {
		panic(err)
	}
	err = l.SetValue(values)
	if err != nil {
		fmt.Println("non")
		panic(err)
	}
	return NewMainBus(l)
}

// RunFunctionDecl executes a parser.FunctionDecl.
func RunFunctionDecl(tree parser.FunctionDecl, env *Env) {
	fn := eclaType.NewFunction(tree.Name, tree.Parameters, tree.Body, tree.ReturnTypes)
	env.SetFunction(tree.Name, fn)
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
	m.SetAutoType()
	return NewMainBus(m)
}
