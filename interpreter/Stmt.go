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

// RunVariableAssignStmt Run assigns a variable.
func RunVariableAssignStmt(tree parser.VariableAssignStmt, env *Env) {
	lValueEqualOne := len(tree.Values) == 1
	if len(tree.Names) == 1 {
		if !lValueEqualOne {
			panic(errors.New("invalid assignment"))
		}
		switch tree.Names[0].(type) {
		case parser.IndexableAccessExpr:
			switch tree.Operator {
			case parser.ASSIGN:
				RunIndexableVariableAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.INCREMENT:
				RunIndexableVariableIncrementStmt(tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.DECREMENT:
				RunIndexableVariableDecrementStmt(tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.ADDASSIGN:
				RunIndexableVariableAddAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			}
		case parser.Literal:
			if tree.Names[0].(parser.Literal).Type == "VAR" {
				switch tree.Operator {
				case parser.ASSIGN:
					RunVariableNonIndexableAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.INCREMENT:
					RunVariableIncrementStmt(tree.Names[0].(parser.Literal), env)
				case parser.DECREMENT:
					RunVariableDecrementStmt(tree.Names[0].(parser.Literal), env)
				case parser.ADDASSIGN: //TODO Change by method
					RunVariableAddAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				default:
					fmt.Println(tree.Operator == "")
					panic(errors.New("invalid assignment"))
				}
			} else {
				panic(errors.New("invalid assignment"))
			}
		default:
			fmt.Printf("%T", tree.Names[0])
		}
	} else if lValueEqualOne {
		for Name := range tree.Names {
			switch tree.Names[Name].(type) {
			case parser.IndexableAccessExpr:
				switch tree.Operator {
				case parser.ASSIGN:
					RunIndexableVariableAssignStmt(tree, tree.Names[Name].(parser.IndexableAccessExpr), env)
				}
			case parser.Literal:
				if tree.Names[Name].(parser.Literal).Type == "VAR" {
					switch tree.Operator {
					case parser.ASSIGN:
						RunVariableNonIndexableAssignStmt(tree, tree.Names[Name].(parser.Literal), env)
					}
				} else {
					panic(errors.New("invalid assignment"))
				}
			}
		}
	} else {
		//TODO
		fmt.Println("multiple values, multiple names")
	}
}

func RunIndexableVariableAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) eclaType.Type {
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
	*temp = RunTree(tree.Values[0], env)
	return nil
}

func RunVariableNonIndexableAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		panic(errors.New("variable not found"))
	}
	temp := RunTree(tree.Values[0], env)
	err := v.SetVar(temp)
	if err != nil {
		panic(err)
	}
}

func RunIndexableVariableIncrementStmt(index parser.IndexableAccessExpr, env *Env) {
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
	switch (*temp).(type) {
	case eclaType.Int:
		*temp = (*temp).(eclaType.Int) + 1
	case eclaType.Float:
		*temp = (*temp).(eclaType.Float) + 1
	default:

		panic(fmt.Sprintf("Variable %s is not incrementable", index.VariableName))
	}
}

func RunVariableIncrementStmt(variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		panic(errors.New("variable not found"))
	}
	v.Increment()
}

func RunIndexableVariableDecrementStmt(index parser.IndexableAccessExpr, env *Env) {
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
	switch (*temp).(type) {
	case eclaType.Int:
		*temp = (*temp).(eclaType.Int) - 1
	case eclaType.Float:
		*temp = (*temp).(eclaType.Float) - 1
	default:
		panic(fmt.Sprintf("Variable %s is not decrementable", index.VariableName))
	}
}

func RunVariableDecrementStmt(variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		panic(errors.New("variable not found"))
	}
	v.Decrement()
}

func RunIndexableVariableAddAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
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
	switch (*temp).(type) {
	case *eclaType.Int:
		*(*temp).(*eclaType.Int) += RunTree(tree.Values[0], env).(eclaType.Int)
	case *eclaType.Float:
		*(*temp).(*eclaType.Float) += RunTree(tree.Values[0], env).(eclaType.Float)
	default:
		panic(fmt.Sprintf("Variable %s is not incrementable", index.VariableName))
	}
}

func RunVariableAddAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		panic(errors.New("variable not found"))
	}
	t, err := v.Add(RunTree(tree.Values[0], env))
	if err != nil {
		panic(err)
	}
	v.SetValue(t)
}

// RunWhileStmt
func RunWhileStmt(tree parser.WhileStmt, env *Env) {
	env.NewScope(SCOPE_LOOP)
	defer env.EndScope()
	while := eclaKeyWord.NewWhile(tree.Cond, tree.Body)
	for RunTree(while.Condition, env).GetString() == "true" { //TODO add error
		for _, stmt := range while.Body {
			RunTree(stmt, env)
		}
	}
}

func RunForStmt(For parser.ForStmt, env *Env) {
	env.NewScope(SCOPE_LOOP)
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
		env.NewScope(SCOPE_CONDITION)
		defer env.EndScope()
		for _, stmt := range tree.Body {
			RunTree(stmt, env)
		}
	} else if tree.ElseStmt != nil {
		if tree.ElseStmt.IfStmt != nil {
			RunIfStmt(*tree.ElseStmt.IfStmt, env)
		} else {
			env.NewScope(SCOPE_CONDITION)
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
