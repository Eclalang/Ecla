package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaKeyWord"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

// RunTree executes a parser.Tree.
func RunTree(tree parser.Node, env *Env) eclaType.Type {
	//fmt.Printf("%T\n", tree)
	switch tree.(type) {
	case parser.Literal:
		return New(tree.(parser.Literal), env)
	case parser.BinaryExpr:
		return RunBinaryExpr(tree.(parser.BinaryExpr), env)
	case parser.UnaryExpr:
		return RunUnaryExpr(tree.(parser.UnaryExpr), env)
	case parser.ParenExpr:
		return RunTree(tree.(parser.ParenExpr).Expression, env)
	case parser.PrintStmt:
		return RunPrintStmt(tree.(parser.PrintStmt), env)
	case parser.TypeStmt:
		return RunTypeStmt(tree.(parser.TypeStmt), env)
	case parser.VariableDecl:
		return RunVariableDecl(tree.(parser.VariableDecl), env)
	case parser.VariableDecrementStmt:
		RunVariableDecrementStmt(tree.(parser.VariableDecrementStmt), env)
	case parser.VariableIncrementStmt:
		RunVariableIncrementStmt(tree.(parser.VariableIncrementStmt), env)
	case parser.VariableAssignStmt:
		RunVariableAssignStmt(tree.(parser.VariableAssignStmt), env)
	case parser.WhileStmt:
		RunWhileStmt(tree.(parser.WhileStmt), env)
	case parser.ForStmt:
		RunForStmt(tree.(parser.ForStmt), env)
	case parser.IfStmt:
		RunIfStmt(tree.(parser.IfStmt), env)
	case parser.ArrayLiteral:
		return RunArrayLiteral(tree.(parser.ArrayLiteral), env)
	case parser.ImportStmt:
		RunImportStmt(tree.(parser.ImportStmt), env)
	case parser.MethodCallExpr:
		return RunMethodCallExpr(tree.(parser.MethodCallExpr), env)
	case parser.FunctionDecl:
		RunFunctionDecl(tree.(parser.FunctionDecl), env)
	case parser.FunctionCallExpr:
		return RunFunctionCallExpr(tree.(parser.FunctionCallExpr), env)
	case parser.IndexableAccessExpr:
		return RunIndexableAccessExpr(tree.(parser.IndexableAccessExpr), env)
	case parser.IndexableVariableAssignStmt:
		RunIndexableVariableAssignStmt(tree.(parser.IndexableVariableAssignStmt), env)
	}
	return nil
}

func RunMethodCallExpr(expr parser.MethodCallExpr, env *Env) eclaType.Type {
	var args []eclaType.Type
	for _, v := range expr.FunctionCall.Args {
		temp := RunTree(v, env)
		switch temp.(type) {
		case *eclaType.Var:
			temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		args = append(args, temp)
	}
	return env.Libs[expr.ObjectName].Call(expr.FunctionCall.Name, args)
}

// RunBinaryExpr executes a parser.BinaryExpr.
func RunBinaryExpr(tree parser.BinaryExpr, env *Env) eclaType.Type {
	//fmt.Printf("%T\n", tree)
	left := RunTree(tree.LeftExpr, env)
	right := RunTree(tree.RightExpr, env)
	switch tree.Operator.TokenType {
	case lexer.ADD:
		t, err := left.Add(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.SUB:
		t, err := left.Sub(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.MULT:
		t, err := left.Mul(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.DIV:
		t, err := left.Div(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.MOD:
		t, err := left.Mod(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.QOT:
		t, err := left.DivEc(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.EQUAL:
		t, err := left.Eq(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.LSS:
		t, err := left.Lw(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.LEQ:
		t, err := left.LwEq(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.GTR:
		t, err := left.Gt(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.GEQ:
		t, err := left.GtEq(right)
		if err != nil {
			panic(err)
		}
		return t
	case lexer.NEQ:
		t, err := left.NotEq(right)
		if err != nil {
			panic(err)
		}
		return t
	}
	return nil
}

// RUnUnaryExpr executes a parser.UnaryExpr.
func RunUnaryExpr(tree parser.UnaryExpr, env *Env) eclaType.Type {
	switch tree.Operator.TokenType {
	case lexer.SUB:
		t, err := eclaType.Int(0).Sub(RunTree(tree.RightExpr, env)) // TODO: Fix this
		if err != nil {
			panic(err)
		}
		return t
	case lexer.ADD:
		return RunTree(tree.RightExpr, env)
	case lexer.NOT:
		t, err := RunTree(tree.RightExpr, env).Not()
		if err != nil {
			panic(err)
		}
		return t
	}
	return nil
}

func RunFunctionCallExpr(tree parser.FunctionCallExpr, env *Env) eclaType.Type {
	env.NewScope()
	defer env.EndScope()
	var args []eclaType.Type
	for _, v := range tree.Args {
		temp := RunTree(v, env)
		switch temp.(type) {
		case *eclaType.Var:
			temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		args = append(args, temp)
	}
	fn, ok := env.GetFunction(tree.Name)
	if !ok {
		panic(fmt.Sprintf("Function %s not found", tree.Name))
	}
	ok, argsList := fn.TypeAndNumberOfArgsIsCorrect(args)
	if !ok {
		panic(fmt.Sprintf("Function %s called with incorrect arguments", tree.Name))
	}

	for i, v := range argsList {
		env.SetVar(i, v)
	}

	r, err := RunBodyFunction(fn, env)
	if err != nil {
		panic(err)
	}
	return r
}

func RunBodyFunction(fn *eclaKeyWord.Function, env *Env) (eclaType.Type, error) {
	for _, v := range fn.Body {
		switch v.(type) {
		case parser.ReturnStmt:
			r := RunReturnStmt(v.(parser.ReturnStmt), env)
			ok := fn.CheckReturn([]eclaType.Type{r})
			if !ok {
				return nil, fmt.Errorf("Return type of function %s is incorrect", fn.Name)
			}
			return r, nil
		}
		RunTree(v, env)
	}
	return nil, nil
}

func RunIndexableAccessExpr(tree parser.IndexableAccessExpr, env *Env) eclaType.Type {
	v, ok := env.GetVar(tree.VariableName)
	if !ok {
		panic(fmt.Sprintf("Variable %s not found", tree.VariableName))
	}
	var result eclaType.Type
	switch v.Value.(type) {
	case *eclaType.List:
		result = v.Value
	default:
		panic(fmt.Sprintf("Variable %s is not indexable", tree.VariableName))
	}

	for i := range tree.Indexes {
		elem := RunTree(tree.Indexes[i], env)
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		if elem.GetType() != "int" {
			panic(fmt.Sprintf("Index must be an integer"))
		}

		switch result.(type) {
		case *eclaType.List:
			result = result.(*eclaType.List).Value[elem.(eclaType.Int)]
		default:
			panic(fmt.Sprintf("Variable %s is not indexable", tree.VariableName))
		}
	}
	return result
}
