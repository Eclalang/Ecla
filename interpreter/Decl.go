package interpreter

import (
	"errors"
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

// New returns a new eclaType.Type from a parser.Literal.
func New(t parser.Literal, env *Env) eclaType.Type {
	switch t.Type {
	case lexer.INT:
		return eclaType.NewInt(t.Value)
	case lexer.STRING:
		return eclaType.NewString(t.Value)
	case lexer.BOOL:
		return eclaType.NewBool(t.Value)
	case lexer.FLOAT:
		return eclaType.NewFloat(t.Value)
	case "VAR":
		v, ok := env.GetVar(t.Value)
		if !ok {
			panic(errors.New("variable not found"))
		}
		return v
	case "NULL":
		return eclaType.NewNull()
	default:
		panic("Unknown type")
	}
}

// RunVariableDecl executes a parser.VariableDecl.
func RunVariableDecl(tree parser.VariableDecl, env *Env) eclaType.Type {
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
		v, err := eclaType.NewVar(tree.Name, tree.Type, RunTree(tree.Value, env))
		if err != nil {
			panic(err)
		}
		env.SetVar(tree.Name, v)
	}
	return nil
}

func RunArrayLiteral(tree parser.ArrayLiteral, env *Env) eclaType.Type {
	var values []eclaType.Type
	for _, v := range tree.Values {
		values = append(values, RunTree(v, env))
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
	return l
}

func RunFunctionDecl(tree parser.FunctionDecl, env *Env) {
	//fn := eclaKeyWord.NewFunction(tree.Name, tree.Parameters, tree.Body, tree.ReturnTypes)

	fn := eclaType.NewFunction(tree.Name, tree.Parameters, tree.Body, tree.ReturnTypes)

	env.SetFunction(tree.Name, fn)
}

func RunMapLiteral(tree parser.MapLiteral, env *Env) eclaType.Type {
	var keys []eclaType.Type
	var values []eclaType.Type
	for _, v := range tree.Values {
		values = append(values, RunTree(v, env))
	}
	for _, k := range tree.Keys {
		keys = append(keys, RunTree(k, env))
	}
	m := eclaType.NewMap()
	m.Keys = keys
	m.Values = values
	m.SetAutoType()
	return m
}
