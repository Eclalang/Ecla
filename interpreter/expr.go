package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/errorHandler"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

// RunTree executes a parser.Tree.
func RunTree(tree parser.Node, env *Env) *Bus {
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
		RunPrintStmt(tree.(parser.PrintStmt), env)
	case parser.TypeStmt:
		RunTypeStmt(tree.(parser.TypeStmt), env)
	case parser.VariableDecl:
		RunVariableDecl(tree.(parser.VariableDecl), env)
	case parser.VariableAssignStmt:
		RunVariableAssignStmt(tree.(parser.VariableAssignStmt), env)
	case parser.WhileStmt:
		return RunWhileStmt(tree.(parser.WhileStmt), env)
	case parser.ForStmt:
		return RunForStmt(tree.(parser.ForStmt), env)
	case parser.IfStmt:
		return RunIfStmt(tree.(parser.IfStmt), env)
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
	case parser.MapLiteral:
		return RunMapLiteral(tree.(parser.MapLiteral), env)
	case parser.ReturnStmt:
		r := RunReturnStmt(tree.(parser.ReturnStmt), env)
		fn := env.GetFunctionExecuted()
		ok := fn.CheckReturn([]eclaType.Type{r})
		if !ok {
			panic("Return type of function" + fn.Name + "is incorrect")
		}
		return NewReturnBus(r)
	}
	return NewNoneBus()
}

// RunMethodCallExpr executes a parser.MethodCallExpr.
func RunMethodCallExpr(expr parser.MethodCallExpr, env *Env) *Bus {
	var args []eclaType.Type
	for _, v := range expr.FunctionCall.Args {
		temp := RunTree(v, env).GetVal()
		switch temp.(type) {
		case *eclaType.Var:
			temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		args = append(args, temp)
	}
	call := env.Libs[expr.ObjectName].Call(expr.FunctionCall.Name, args)
	if call == nil {
		panic(fmt.Sprintf("Method %s not found in module %s", expr.FunctionCall.Name, expr.ObjectName))
	}
	return NewMainBus(call)
}

// RunBinaryExpr executes a parser.BinaryExpr.
func RunBinaryExpr(tree parser.BinaryExpr, env *Env) *Bus {
	left := RunTree(tree.LeftExpr, env).GetVal()
	right := RunTree(tree.RightExpr, env).GetVal()
	var t eclaType.Type
	var err error
	switch tree.Operator.TokenType {
	case lexer.ADD:
		t, err = left.Add(right)
	case lexer.SUB:
		t, err = left.Sub(right)
	case lexer.MULT:
		t, err = left.Mul(right)
	case lexer.DIV:
		t, err = left.Div(right)
	case lexer.MOD:
		t, err = left.Mod(right)
	case lexer.QOT:
		t, err = left.DivEc(right)
	case lexer.EQUAL:
		t, err = left.Eq(right)
	case lexer.LSS:
		t, err = left.Lw(right)
	case lexer.LEQ:
		t, err = left.LwEq(right)
	case lexer.GTR:
		t, err = left.Gt(right)
	case lexer.GEQ:
		t, err = left.GtEq(right)
	case lexer.NEQ:
		t, err = left.NotEq(right)
	case lexer.AND:
		t, err = left.And(right)
	case lexer.OR:
		t, err = left.Or(right)
	default:
		return NewNoneBus()
	}
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(t)
}

// RunUnaryExpr executes a parser.UnaryExpr.
func RunUnaryExpr(tree parser.UnaryExpr, env *Env) *Bus {
	switch tree.Operator.TokenType {
	case lexer.SUB:
		t, err := eclaType.Int(0).Sub(RunTree(tree.RightExpr, env).GetVal()) // TODO: Fix this
		if err != nil {
			panic(err)
		}
		return NewMainBus(t)
	case lexer.ADD:
		return RunTree(tree.RightExpr, env)
	case lexer.NOT:
		t, err := RunTree(tree.RightExpr, env).GetVal().Not()
		if err != nil {
			panic(err)
		}
		return NewMainBus(t)
	}
	return NewNoneBus()
}

func RunFunctionCallExpr(tree parser.FunctionCallExpr, env *Env) *Bus {
	env.NewScope(SCOPE_FUNCTION)
	defer env.EndScope()
	var args []eclaType.Type
	for _, v := range tree.Args {
		temp := RunTree(v, env).GetVal()
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
	env.AddFunctionExecuted(fn)
	defer env.RemoveFunctionExecuted()
	r, err := RunBodyFunction(fn, env)
	if err != nil {
		panic(err)
	}
	return NewMainBus(r)
}

func RunBodyFunction(fn *eclaType.Function, env *Env) (eclaType.Type, error) {
	for _, v := range fn.Body {
		temp := RunTree(v, env)
		if temp.IsReturn() {
			return temp.GetVal().GetValue().(eclaType.Type), nil
		}
	}
	return nil, nil
}

func RunIndexableAccessExpr(tree parser.IndexableAccessExpr, env *Env) *Bus {
	v, ok := env.GetVar(tree.VariableName)
	if !ok {
		panic(fmt.Sprintf("Variable %s not found", tree.VariableName))
	}
	var result eclaType.Type = v
	for i := range tree.Indexes {
		elem := RunTree(tree.Indexes[i], env).GetVal()
		//fmt.Printf("%s\n", elem.GetValue())
		//fmt.Printf("%T\n", result.GetValue())
		temp, err := result.GetIndex(elem)

		result = *temp

		if err != nil {
			panic(err)
		}
	}
	return NewMainBus(result)
}
