package interpreter

import (
	"errors"
	"fmt"
	"github.com/tot0p/Ecla/errorHandler"
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
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Can't assign %d rValues to 1 lValue", len(tree.Values)), errorHandler.LevelFatal)
		}
		switch tree.Names[0].(type) {
		case parser.IndexableAccessExpr:
			switch tree.Operator {
			case parser.ASSIGN:
				RunIndexableVariableAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.INCREMENT:
				RunIndexableVariableIncrementStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.DECREMENT:
				RunIndexableVariableDecrementStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.ADDASSIGN:
				RunIndexableVariableAddAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.SUBASSIGN:
				RunIndexableVariableSubAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.DIVASSIGN:
				RunIndexableVariableDivAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.MODASSIGN:
				RunIndexableVariableModAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.QOTASSIGN:
				RunIndexableVariableQotAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			case parser.MULTASSIGN:
				RunIndexableVariableMultAssignStmt(tree, tree.Names[0].(parser.IndexableAccessExpr), env)
			default:
				env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("%s is not a valid assignement operator", tree.Operator), errorHandler.LevelFatal)
			}
		case parser.Literal:
			if tree.Names[0].(parser.Literal).Type == "VAR" {
				switch tree.Operator {
				case parser.ASSIGN:
					RunVariableNonIndexableAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.INCREMENT:
					RunVariableIncrementStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.DECREMENT:
					RunVariableDecrementStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.ADDASSIGN: //TODO Change by method
					RunVariableAddAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.SUBASSIGN:
					RunVariableSubAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.DIVASSIGN:
					RunVariableDivAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.MODASSIGN:
					RunVariableModAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.QOTASSIGN:
					RunVariableQotAssignStmt(tree, tree.Names[0].(parser.Literal), env)
				case parser.MULTASSIGN:
					RunVariableMultAssignStmt(tree, tree.Names[0].(parser.Literal), env)

				default:
					env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("%s is not a valid assignement operator", tree.Operator), errorHandler.LevelFatal)
				}
			} else {
				env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Cant run assignement on type %s", tree.Names[0].(parser.Literal).Type), errorHandler.LevelFatal)
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
					default:
						env.ErrorHandle.HandleError(0, tree.StartPos(), "invalid assignment", errorHandler.LevelFatal)
					}
				} else {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "invalid assignment", errorHandler.LevelFatal)
				}
			}
		}
	} else {
		//TODO
		fmt.Println("multiple values, multiple names")
	}
}

func IndexableAssignementChecks(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) *eclaType.Type {
	v, ok := env.GetVar(index.VariableName)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: variable not found", errorHandler.LevelFatal)
	}
	var temp *eclaType.Type
	switch v.Value.(type) {
	case *eclaType.List:
		temp = &v.Value
	default:
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Variable %s is not indexable", index.VariableName), errorHandler.LevelFatal)
	}
	for i := range index.Indexes {
		elem := RunTree(index.Indexes[i], env)
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		if elem.GetType() != "int" {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "Index must be an integer", errorHandler.LevelFatal)
		}
		switch (*temp).(type) {
		case *eclaType.List:
			temp = &((*temp).(*eclaType.List).Value[elem.(eclaType.Int)])
		default:
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Variable %s is not indexable", index.VariableName), errorHandler.LevelFatal)
		}
	}
	return temp
}

func RunIndexableVariableAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) eclaType.Type {
	temp := IndexableAssignementChecks(tree, index, env)
	*temp = RunTree(tree.Values[0], env)
	return nil
}

func RunVariableNonIndexableAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	temp := RunTree(tree.Values[0], env)
	err := v.SetVar(temp)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)

	}
}

func RunIndexableVariableIncrementStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Add(eclaType.Int(1))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}
func RunVariableIncrementStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	v.Increment()
}

func RunIndexableVariableDecrementStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Sub(eclaType.Int(1))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableDecrementStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	v.Decrement()
}

func RunIndexableVariableAddAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Add(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableAddAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.Add(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func RunIndexableVariableSubAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Sub(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableSubAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.Sub(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func RunIndexableVariableDivAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Div(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableDivAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.Div(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func RunIndexableVariableModAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Mod(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableModAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.Mod(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func RunIndexableVariableQotAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).DivEc(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableQotAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.DivEc(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func RunIndexableVariableMultAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignementChecks(tree, index, env)
	res, err := (*temp).Mul(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

func RunVariableMultAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	t, err := v.Mul(RunTree(tree.Values[0], env))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
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
func RunIfStmt(tree parser.IfStmt, env *Env) eclaType.Type {
	if RunTree(tree.Cond, env).GetString() == "true" { //TODO add error
		env.NewScope(SCOPE_CONDITION)
		defer env.EndScope()
		for _, stmt := range tree.Body {
			temp := RunTree(stmt, env)
			if temp != nil {
				return temp
			}
		}
	} else if tree.ElseStmt != nil {
		if tree.ElseStmt.IfStmt != nil {
			return RunIfStmt(*tree.ElseStmt.IfStmt, env)
		} else {
			env.NewScope(SCOPE_CONDITION)
			defer env.EndScope()
			for _, stmt := range tree.ElseStmt.Body {
				temp := RunTree(stmt, env)
				if temp != nil {
					return temp
				}
			}
		}
	}
	return nil
}

func RunReturnStmt(tree parser.ReturnStmt, env *Env) eclaType.Type {
	l := []eclaType.Type{}
	for _, expr := range tree.ReturnValues {
		l = append(l, RunTree(expr, env))
	}
	return l[0]
}
