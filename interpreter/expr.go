package interpreter

import (
	"fmt"

	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
)

// RunTree executes a parser.Node
func RunTree(tree parser.Node, env *Env) []*Bus {
	switch tree.(type) {
	case parser.Literal:
		return []*Bus{New(tree.(parser.Literal), env)}
	case parser.BinaryExpr:
		return []*Bus{RunBinaryExpr(tree.(parser.BinaryExpr), env)}
	case parser.UnaryExpr:
		return []*Bus{RunUnaryExpr(tree.(parser.UnaryExpr), env)}
	case parser.ParenExpr:
		return RunTree(tree.(parser.ParenExpr).Expression, env)
	case parser.TypeStmt:
		RunTypeStmt(tree.(parser.TypeStmt), env)
	case parser.VariableDecl:
		RunVariableDecl(tree.(parser.VariableDecl), env)
	case parser.VariableAssignStmt:
		RunVariableAssignStmt(tree.(parser.VariableAssignStmt), env)
	case parser.WhileStmt:
		return []*Bus{RunWhileStmt(tree.(parser.WhileStmt), env)}
	case parser.ForStmt:
		return []*Bus{RunForStmt(tree.(parser.ForStmt), env)}
	case parser.IfStmt:
		return []*Bus{RunIfStmt(tree.(parser.IfStmt), env)}
	case parser.ArrayLiteral:
		return []*Bus{RunArrayLiteral(tree.(parser.ArrayLiteral), env)}
	case parser.ImportStmt:
		RunImportStmt(tree.(parser.ImportStmt), env)
	case parser.MethodCallExpr:
		return RunMethodCallExpr(tree.(parser.MethodCallExpr), env)
	case parser.FunctionDecl:
		RunFunctionDecl(tree.(parser.FunctionDecl), env)
	case parser.FunctionCallExpr:
		return RunFunctionCallExpr(tree.(parser.FunctionCallExpr), env)
	case parser.IndexableAccessExpr:
		return []*Bus{RunIndexableAccessExpr(tree.(parser.IndexableAccessExpr), env)}
	case parser.MapLiteral:
		return []*Bus{RunMapLiteral(tree.(parser.MapLiteral), env)}
	case parser.ReturnStmt:
		r := RunReturnStmt(tree.(parser.ReturnStmt), env)
		fn := env.GetFunctionExecuted()
		ok := fn.CheckReturn(r)
		if !ok {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "Return type of function "+fn.Name+" is incorrect", errorHandler.LevelFatal)
		}
		var temp []*Bus
		for _, v := range r {
			temp = append(temp, NewReturnBus(v))
		}
		return temp
	case parser.MurlocStmt:
		RunMurlocStmt(tree.(parser.MurlocStmt), env)
	case parser.AnonymousFunctionExpr:
		return RunAnonymousFunctionExpr(tree.(parser.AnonymousFunctionExpr), env)
	case parser.BlockScopeStmt:
		return RunBlockScopeStmt(tree.(parser.BlockScopeStmt), env)
	case parser.AnonymousFunctionCallExpr:
		return RunAnonymousFunctionCallExpr(tree.(parser.AnonymousFunctionCallExpr), env)
	case parser.StructDecl:
		RunStructDecl(tree.(parser.StructDecl), env)
	case parser.SelectorExpr:
		return RunSelectorExpr(tree.(parser.SelectorExpr), env)
	default:
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Not implemented : %T\n", tree), errorHandler.LevelFatal)
	}

	return []*Bus{NewNoneBus()}
}

func RunAnonymousFunctionExpr(AnonymousFunc parser.AnonymousFunctionExpr, env *Env) []*Bus {
	fn := eclaType.NewAnonymousFunction(AnonymousFunc.Prototype.Parameters, AnonymousFunc.Body, AnonymousFunc.Prototype.ReturnTypes)
	returnBus := []*Bus{NewMainBus(fn)}
	return returnBus
}

// RunTreeLoad is special version of RunTree that is used to load the environment (function, variable, import)
func RunTreeLoad(tree parser.Node, env *Env) []*Bus {
	switch tree.(type) {
	case parser.VariableDecl:
		RunVariableDecl(tree.(parser.VariableDecl), env)
	case parser.FunctionDecl:
		RunFunctionDecl(tree.(parser.FunctionDecl), env)
	case parser.ImportStmt:
		RunImportStmt(tree.(parser.ImportStmt), env)
	}
	return []*Bus{NewNoneBus()}
}

// RunMethodCallExpr executes a parser.MethodCallExpr.
func RunMethodCallExpr(expr parser.MethodCallExpr, env *Env) []*Bus {
	var args []eclaType.Type
	for _, v := range expr.FunctionCall.Args {
		BusCollection := RunTree(v, env)
		for _, bus := range BusCollection {
			temp := bus.GetVal()
			switch temp.(type) {
			case *eclaType.Var:
				temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
			}
			args = append(args, temp)
		}
	}
	var returnBuses []*Bus
	call, callErr := env.Libs[expr.ObjectName].Call(expr.FunctionCall.Name, args)
	if callErr != nil {
		env.ErrorHandle.HandleError(0, expr.StartPos(), callErr.Error(), errorHandler.LevelFatal)
	}
	for _, elem := range call {
		returnBuses = append(returnBuses, NewMainBus(elem))
	}
	return returnBuses
}

// RunBinaryExpr executes a parser.BinaryExpr.
func RunBinaryExpr(tree parser.BinaryExpr, env *Env) *Bus {
	BusCollection := RunTree(tree.LeftExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.LeftExpr.StartPos(), "MULTIPLE BUS IN RunBinaryExpr", errorHandler.LevelFatal)
	}
	left := BusCollection[0].GetVal()
	BusCollection = RunTree(tree.RightExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.RightExpr.StartPos(), "MULTIPLE BUS IN RunBinaryExpr", errorHandler.LevelFatal)
	}
	right := BusCollection[0].GetVal()
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
	case lexer.XOR:
		t, err = left.Xor(right)
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
	BusCollection := RunTree(tree.RightExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.RightExpr.StartPos(), "MULTIPLE BUS IN RunUnaryExpr", errorHandler.LevelFatal)
	}
	switch tree.Operator.TokenType {
	case lexer.SUB:
		t, err := eclaType.Int(0).Sub(BusCollection[0].GetVal()) // TODO: Fix this
		if err != nil {
			env.ErrorHandle.HandleError(0, tree.RightExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(t)
	case lexer.ADD:
		return BusCollection[0]
	case lexer.NOT:
		t, err := BusCollection[0].GetVal().Not()
		if err != nil {
			env.ErrorHandle.HandleError(0, tree.RightExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(t)
	}
	return NewNoneBus()
}

// RunFunctionCallExpr executes a parser.FunctionCallExpr.
func RunFunctionCallExpr(tree parser.FunctionCallExpr, env *Env) []*Bus {
	var args []eclaType.Type
	for _, v := range tree.Args {
		BusCollection := RunTree(v, env)
		for _, bus := range BusCollection {
			temp := bus.GetVal()
			switch temp.(type) {
			case *eclaType.Var:
				temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
			}
			args = append(args, temp)
		}

	}
	fn, ok := env.GetFunction(tree.Name)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Function %s not found", tree.Name), errorHandler.LevelFatal)
	}
	r, err := RunFunctionCallExprWithArgs(tree.Name, env, fn, args)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	var retValues []*Bus
	for _, v := range r {
		retValues = append(retValues, NewMainBus(v))
	}
	return retValues
}

// RunFunctionCallExprWithArgs executes a parser.FunctionCallExpr with the given arguments.
func RunFunctionCallExprWithArgs(Name string, env *Env, fn *eclaType.Function, args []eclaType.Type) ([]eclaType.Type, error) {
	env.NewScope(SCOPE_FUNCTION)
	defer env.EndScope()
	ok, argsList := fn.TypeAndNumberOfArgsIsCorrect(args)
	if !ok {
		env.ErrorHandle.HandleError(0, 0, fmt.Sprintf("Function %s called with incorrect arguments", Name), errorHandler.LevelFatal)
	}

	for i, v := range argsList {
		env.SetVar(i, v)
	}
	env.AddFunctionExecuted(fn)
	defer env.RemoveFunctionExecuted()

	return RunBodyFunction(fn, env)
}

// RunBodyFunction executes the code associated with the function.
func RunBodyFunction(fn *eclaType.Function, env *Env) ([]eclaType.Type, error) {
	for _, v := range fn.GetBody() {
		BusCollection := RunTree(v, env)
		if IsMultipleBus(BusCollection) {
			ret := true
			var retVal []eclaType.Type
			for _, bus := range BusCollection {
				if !bus.IsReturn() {
					ret = false
				}
				retVal = append(retVal, bus.GetVal().GetValue().(eclaType.Type))
			}
			if ret {
				return retVal, nil
			}
		} else if len(BusCollection) == 1 {
			if BusCollection[0].IsReturn() {
				return []eclaType.Type{BusCollection[0].GetVal().GetValue().(eclaType.Type)}, nil
			}
		}
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}

// RunIndexableAccessExpr executes a parser.IndexableAccessExpr.
func RunIndexableAccessExpr(tree parser.IndexableAccessExpr, env *Env) *Bus {
	v, ok := env.GetVar(tree.VariableName)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Variable %s not found", tree.VariableName), errorHandler.LevelFatal)
	}
	var result eclaType.Type = v
	for i := range tree.Indexes {
		BusCollection := RunTree(tree.Indexes[i], env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, tree.Indexes[i].StartPos(), "MULTIPLE BUS IN RunIndexableAccessExpr", errorHandler.LevelFatal)
		}
		elem := BusCollection[0].GetVal()
		temp, err := result.GetIndex(elem)

		result = *temp

		if err != nil {
			env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
	}
	return NewMainBus(result)
}

func RunAnonymousFunctionCallExpr(tree parser.AnonymousFunctionCallExpr, env *Env) []*Bus {
	fn := RunTree(tree.AnonymousFunction, env)
	if IsMultipleBus(fn) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "MULTIPLE BUS IN RunAnonymousFunctionCallExpr", errorHandler.LevelFatal)
	}
	var f *eclaType.Function
	switch fn[0].GetVal().(type) {
	case *eclaType.Function:
		f = fn[0].GetVal().(*eclaType.Function)
	default:
		env.ErrorHandle.HandleError(0, tree.StartPos(), "Cannot call a non-function", errorHandler.LevelFatal)
	}
	var args []eclaType.Type
	for _, v := range tree.Args {
		BusCollection := RunTree(v, env)
		for _, bus := range BusCollection {
			temp := bus.GetVal()
			switch temp.(type) {
			case *eclaType.Var:
				temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
			}
			args = append(args, temp)
		}
	}
	r, err := RunFunctionCallExprWithArgs("anonymous function", env, f, args)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	var retValues []*Bus
	for _, v := range r {

		retValues = append(retValues, NewReturnBus(v))
	}
	return retValues
}

func RunBlockScopeStmt(tree parser.BlockScopeStmt, env *Env) []*Bus {
	env.NewScope(SCOPE_MAIN)
	defer env.EndScope()
	for _, v := range tree.Body {
		RunTree(v, env)
	}
	return []*Bus{NewNoneBus()}
}

func RunSelectorExpr(expr parser.SelectorExpr, env *Env) []*Bus {
	expr1 := RunTree(expr.Expr, env)
	if IsMultipleBus(expr1) {
		env.ErrorHandle.HandleError(0, expr.StartPos(), "MULTIPLE BUS IN RunSelectorExpr", errorHandler.LevelFatal)
	}
	var prev eclaType.Type
	switch expr1[0].GetVal().(type) {
	case *eclaType.Var:
		prev = expr1[0].GetVal().(*eclaType.Var).Value
	}

	switch prev.(type) {
	case *eclaType.Lib:
		lib := env.Libs[prev.(*eclaType.Lib).Name]
		switch expr.Sel.(type) {
		case parser.FunctionCallExpr:
			var args []eclaType.Type
			for _, v := range expr.Sel.(parser.FunctionCallExpr).Args {
				BusCollection := RunTree(v, env)
				for _, bus := range BusCollection {
					temp := bus.GetVal()
					switch temp.(type) {
					case *eclaType.Var:
						temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
					}
					args = append(args, temp)
				}
			}
			var returnBuses []*Bus
			result, err := lib.Call(expr.Sel.(parser.FunctionCallExpr).Name, args)
			if err != nil {
				env.ErrorHandle.HandleError(0, expr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			for _, elem := range result {
				returnBuses = append(returnBuses, NewMainBus(elem))
			}
			return returnBuses
		default:
			env.ErrorHandle.HandleError(0, expr.StartPos(), "SelectorExpr not implemented", errorHandler.LevelFatal)
		}
	}

	// TODO : implements for struct

	//expr2 := RunTree(expr.Sel, env)
	//if IsMultipleBus(expr2) {
	//	env.ErrorHandle.HandleError(0, expr.StartPos(), "MULTIPLE BUS IN RunSelectorExpr", errorHandler.LevelFatal)
	//}

	return nil
}
