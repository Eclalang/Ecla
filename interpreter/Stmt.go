package interpreter

import (
	"errors"
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaKeyWord"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
	"strconv"
)

// RunPrintStmt executes a parser.PrintStmt.
func RunPrintStmt(tree parser.PrintStmt, env *Env) eclaType.Type {
	fmt.Print(RunTree(tree.Expression, env).GetString())
	return nil
}

func RunImportStmt(stmt parser.ImportStmt, env *Env) {
	env.Import(stmt.ModulePath)
}

// RunTypeStmt executes a parser.TypeStmt.
func RunTypeStmt(tree parser.TypeStmt, env *Env) eclaType.Type {
	t := RunTree(tree.Expression, env)
	var typ string
	typ = t.GetType()
	fmt.Println(typ)
	return nil
	//return eclaType.NewString(RunTree(tree.Expression, env).GetType())
}

// RunVariableDecrementStmt Run decrements a variable.
func RunVariableDecrementStmt(tree parser.VariableDecrementStmt, env *Env) {
	v, ok := env.GetVar(tree.Name)
	if !ok {
		panic(errors.New("variable not found"))
	}
	v.Decrement()
}

// RunVariableIncrementStmt Run increments a variable.
func RunVariableIncrementStmt(tree parser.VariableIncrementStmt, env *Env) {
	v, ok := env.GetVar(tree.Name)
	if !ok {
		panic(errors.New("variable not found"))
	}
	v.Increment()
}

// RunVariableAssignStmt Run assigns a variable.
func RunVariableAssignStmt(tree parser.VariableAssignStmt, env *Env) {
	if len(tree.Name) == 1 {
		v, ok := env.GetVar(tree.Name[0])
		if !ok {
			panic(errors.New("variable not found"))
		}
		err := v.SetVar(RunTree(tree.Value[0], env))
		if err != nil {
			panic(err)
		}
	}
}

// RunWhileStmt
func RunWhileStmt(tree parser.WhileStmt, env *Env) {
	env.NewScope()
	defer env.EndScope()
	while := eclaKeyWord.NewWhile(tree.Cond, tree.Body)
	for RunTree(while.Condition, env).GetString() == "true" { //TODO add error
		for _, stmt := range while.Body {
			RunTree(stmt, env)
		}
	}
}

func RunForStmt(For parser.ForStmt, env *Env) {
	env.NewScope()
	defer env.EndScope()
	tokenEmpty := lexer.Token{}
	if For.RangeToken != tokenEmpty {
		f := eclaKeyWord.NewForRange([]eclaType.Type{}, For.RangeExpr, For.KeyToken, For.ValueToken, For.Body)
		k, err := eclaType.NewVar(f.KeyToken.Value, "int", eclaType.NewInt("0"))
		if err != nil {
			panic(err)
		}
		list := RunTree(f.RangeExpr, env)
		var typ string
		var l int //...
		//fmt.Printf("%T", list)
		switch list.(type) {
		case *eclaType.List:
			typ = list.(*eclaType.List).GetType()[2:]
			l = list.(*eclaType.List).Len()
		case eclaType.String:
			typ = list.GetType()
			l = list.(eclaType.String).Len()
		case *eclaType.Var:
			temp := list.(*eclaType.Var).GetValue()
			//fmt.Printf("%T", temp)
			switch temp.(type) {
			case *eclaType.List:
				typ = temp.(*eclaType.List).GetType()[2:]
				l = temp.(*eclaType.List).Len()
			case eclaType.String:
				typ = temp.(eclaType.String).GetType()
				l = temp.(eclaType.String).Len()
			default:
				panic(errors.New("for range: type " + list.GetType() + " not supported"))
			}
		default:
			panic(errors.New("type " + list.GetType() + " not supported"))
		}

		env.SetVar(f.KeyToken.Value, k)
		v, err := eclaType.NewVarEmpty(f.ValueToken.Value, typ)
		if err != nil {
			panic(err)
		}
		env.SetVar(f.ValueToken.Value, v)
		for i := 0; i < l; i++ {
			k.SetVar(eclaType.NewInt(strconv.Itoa(i)))
			val, err := list.GetIndex(eclaType.Int(i))
			if err != nil {
				panic(err)
			}
			err = v.SetVar(val)
			if err != nil {
				panic(err)
			}
			for _, stmt := range f.Body {
				RunTree(stmt, env)
			}
		}
	} else {
		f := eclaKeyWord.NewForI([]eclaType.Type{}, For.Body, For.CondExpr, For.PostAssignStmt)
		RunTree(For.InitDecl, env)
		for RunTree(f.Condition, env).GetString() == "true" { //TODO add error
			for _, stmt := range f.Body {
				RunTree(stmt, env)
			}
			RunTree(f.Post, env)
		}
	}
}

// RunIfStmt
func RunIfStmt(tree parser.IfStmt, env *Env) {
	if RunTree(tree.Cond, env).GetString() == "true" { //TODO add error
		env.NewScope()
		defer env.EndScope()
		for _, stmt := range tree.Body {
			RunTree(stmt, env)
		}
	} else if tree.ElseStmt != nil {
		if tree.ElseStmt.IfStmt != nil {
			RunIfStmt(*tree.ElseStmt.IfStmt, env)
		} else {
			env.NewScope()
			defer env.EndScope()
			for _, stmt := range tree.ElseStmt.Body {
				RunTree(stmt, env)
			}
		}
	}
}

func RunReturnStmt(tree parser.ReturnStmt, env *Env) eclaType.Type {
	l := []eclaType.Type{}
	for _, expr := range tree.ReturnValues {
		l = append(l, RunTree(expr, env))
	}
	return l[0]
}

func RunIndexableVariableAssignStmt(tree parser.IndexableVariableAssignStmt, env *Env) eclaType.Type {
	var index parser.IndexableAccessExpr

	switch tree.IndexableAccess[0].(type) {
	case parser.IndexableAccessExpr:
		index = tree.IndexableAccess[0].(parser.IndexableAccessExpr)
	default:
		panic(errors.New("indexable variable assign: indexable access not found"))
	}

	v, ok := env.GetVar(index.VariableName)
	if !ok {
		panic(errors.New("indexable variable assign: variable not found"))
	}

	var temp *eclaType.Type

	switch v.Value.(type) {
	case *eclaType.List:
		temp = &v.Value
	default:
		panic(fmt.Sprintf("Variable %s is not indexable", index.VariableName))
	}

	for i := range index.Indexes {
		elem := RunTree(index.Indexes[i], env)
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		if elem.GetType() != "int" {
			panic(fmt.Sprintf("Index must be an integer"))
		}

		switch (*temp).(type) {
		case *eclaType.List:
			temp = &((*temp).(*eclaType.List).Value[elem.(eclaType.Int)])
		default:
			panic(fmt.Sprintf("Variable %s is not indexable", index.VariableName))
		}
	}
	*temp = RunTree(tree.Value[0], env)
	return nil
}