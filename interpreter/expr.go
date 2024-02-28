package interpreter

import (
	"fmt"
	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
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
		ok := fn.CheckReturn(r, env.TypeDecl)
		if !ok {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Return type of function "+fn.Name+" is incorrect", errorHandler.LevelFatal)
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
		return RunSelectorExpr(tree.(parser.SelectorExpr), env, nil)
	case parser.StructInstantiationExpr:
		return RunStructInstantiationExpr(tree.(parser.StructInstantiationExpr), env)
	default:
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), fmt.Sprintf("Not implemented : %T\n", tree), errorHandler.LevelFatal)
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
	case parser.StructDecl:
		RunStructDecl(tree.(parser.StructDecl), env)
	}
	return []*Bus{NewNoneBus()}
}

// RunBinaryExpr executes a parser.BinaryExpr.
func RunBinaryExpr(tree parser.BinaryExpr, env *Env) *Bus {
	BusCollection := RunTree(tree.LeftExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(tree.LeftExpr.StartLine(), tree.LeftExpr.StartPos(), "MULTIPLE BUS IN RunBinaryExpr.\nPlease open issue", errorHandler.LevelFatal)
	}
	left := BusCollection[0].GetVal()
	BusCollection = RunTree(tree.RightExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(tree.RightExpr.StartLine(), tree.RightExpr.StartPos(), "MULTIPLE BUS IN RunBinaryExpr.\nPlease open issue", errorHandler.LevelFatal)
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
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return NewMainBus(t)
}

// RunUnaryExpr executes a parser.UnaryExpr.
func RunUnaryExpr(tree parser.UnaryExpr, env *Env) *Bus {
	BusCollection := RunTree(tree.RightExpr, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(tree.RightExpr.StartLine(), tree.RightExpr.StartPos(), "MULTIPLE BUS IN RunUnaryExpr.\nPlease open issue", errorHandler.LevelFatal)
	}
	switch tree.Operator.TokenType {
	case lexer.SUB:
		t, err := eclaType.Int(0).Sub(BusCollection[0].GetVal()) // TODO: Fix this
		if err != nil {
			env.ErrorHandle.HandleError(tree.RightExpr.StartLine(), tree.RightExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		return NewMainBus(t)
	case lexer.ADD:
		return BusCollection[0]
	case lexer.NOT:
		t, err := BusCollection[0].GetVal().Not()
		if err != nil {
			env.ErrorHandle.HandleError(tree.RightExpr.StartLine(), tree.RightExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
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
	v, ok := env.GetVar(tree.Name)
	if !ok {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), fmt.Sprintf("Function %s not found", tree.Name), errorHandler.LevelFatal)
	}
	var fn *eclaType.Function
	if v.IsFunction() {
		fn = v.GetFunction()
	}
	var r []eclaType.Type
	var err error
	if fn != nil {
		r, err = RunFunctionCallExprWithArgs(tree.Name, env, fn, args)
		if err != nil {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
	} else {
		switch v.Value.(type) {
		case *eclaType.FunctionBuiltIn:
			r, err = v.Value.(*eclaType.FunctionBuiltIn).Call(args)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
		default:
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), fmt.Sprintf("Function %s not found", tree.Name), errorHandler.LevelFatal)
		}
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
	ok, argsList := fn.TypeAndNumberOfArgsIsCorrect(args, env.TypeDecl)
	if !ok {
		return nil, fmt.Errorf("function %s called with incorrect arguments", Name)
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
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), fmt.Sprintf("Variable %s not found", tree.VariableName), errorHandler.LevelFatal)
	}
	var result eclaType.Type = v
	for i := range tree.Indexes {
		BusCollection := RunTree(tree.Indexes[i], env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(tree.Indexes[i].StartLine(), tree.Indexes[i].StartPos(), "MULTIPLE BUS IN RunIndexableAccessExpr.\nPlease open issue", errorHandler.LevelFatal)
		}
		elem := BusCollection[0].GetVal()
		temp, err := result.GetIndex(elem)

		if err != nil {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		result = *temp

	}
	return NewMainBus(result)
}

func RunAnonymousFunctionCallExpr(tree parser.AnonymousFunctionCallExpr, env *Env) []*Bus {
	fn := RunTree(tree.AnonymousFunction, env)
	if IsMultipleBus(fn) {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunAnonymousFunctionCallExpr.\nPlease open issue", errorHandler.LevelFatal)
	}
	var f *eclaType.Function
	switch fn[0].GetVal().(type) {
	case *eclaType.Function:
		f = fn[0].GetVal().(*eclaType.Function)
	default:
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "Cannot call a non-function", errorHandler.LevelFatal)
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
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
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

func RunSelectorExpr(expr parser.SelectorExpr, env *Env, Struct eclaType.Type) []*Bus {
	prev := Struct
	if Struct == nil {
		expr1 := RunTree(expr.Expr, env)
		if IsMultipleBus(expr1) {
			env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "MULTIPLE BUS IN RunSelectorExpr.\nPlease open issue", errorHandler.LevelFatal)
		}

		switch expr1[0].GetVal().(type) {
		case *eclaType.Var:
			prev = expr1[0].GetVal().(*eclaType.Var).Value
		default:
			prev = expr1[0].GetVal()
		}
	}

	switch prev.(type) {
	case *eclaType.Lib:
		lib := env.Libs[prev.(*eclaType.Lib).Name]
		lastLib := env.Libs
		defer func() { env.Libs = lastLib }()
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
			switch lib.(type) {
			case *envLib:
				env.SetScope(lib.(*envLib).Var)
				env.Libs = lib.(*envLib).Libs
			}
			result, err := lib.Call(expr.Sel.(parser.FunctionCallExpr).Name, args)
			if err != nil {
				env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			for _, elem := range result {
				returnBuses = append(returnBuses, NewMainBus(elem))
			}
			switch lib.(type) {
			case *envLib:
				env.EndScope()
			}
			return returnBuses
		case parser.Literal:
			sel := expr.Sel.(parser.Literal)
			if sel.Type == "VAR" { //TODO don't hard code "VAR"
				v, ok := lib.(*envLib).GetVar(sel.Value)
				if !ok {
					env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "variable "+sel.Value+" does not exist", errorHandler.LevelFatal)
				}
				return []*Bus{NewMainBus(v)}
			}
		case parser.SelectorExpr:
			sel := expr.Sel.(parser.SelectorExpr)
			//check if sel is a struct
			expr := RunTree(sel.Expr, env)
			if IsMultipleBus(expr) {
				env.ErrorHandle.HandleError(sel.StartLine(), sel.StartPos(), "MULTIPLE BUS IN RunSelectorExpr.\nPlease open issue", errorHandler.LevelFatal)
			}
			switch expr[0].GetVal().(type) {
			case *eclaType.Struct:
				//s := expr[0].GetVal().(*eclaType.Struct)
				//return RunSelectorExpr(sel, env, nil)
			}
		default:
			env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "cannot use "+prev.String()+" here", errorHandler.LevelFatal)
		}
	case *eclaType.Struct:
		switch expr.Sel.(type) {
		case parser.Literal:
			sel := expr.Sel.(parser.Literal)
			if sel.Type == "VAR" { //TODO don't hard code "VAR"
				s := prev.(*eclaType.Struct)
				result, ok := s.Fields[sel.Value]
				if !ok {
					env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "field "+sel.Value+" does not exist", errorHandler.LevelFatal)
				}
				return []*Bus{NewMainBus(*result)}
			}
		case parser.FunctionCallExpr:
			tree := expr.Sel.(parser.FunctionCallExpr)
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

			fn, ok := prev.(*eclaType.Struct).Fields[tree.Name]
			if !ok {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "field "+tree.Name+" does not exist", errorHandler.LevelFatal)
			}
			var foo *eclaType.Function
			switch (*fn).(type) {
			case *eclaType.Function:
				foo = (*fn).(*eclaType.Function)
			}
			r, err := RunFunctionCallExprWithArgs(tree.Name, env, foo, args)
			if err != nil {
				env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			var retValues []*Bus
			for _, v := range r {
				retValues = append(retValues, NewMainBus(v))
			}
			return retValues
		case parser.SelectorExpr:
			sel := expr.Sel.(parser.SelectorExpr)
			switch sel.Expr.(type) {

			case parser.Literal:
				sel := sel.Expr.(parser.Literal)
				if sel.Type == "VAR" { //TODO don't hard code "VAR"
					s := prev.(*eclaType.Struct)
					result, ok := s.Fields[sel.Value]
					if !ok {
						env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "field "+sel.Value+" does not exist", errorHandler.LevelFatal)
					}
					prev = *result
				}
			case parser.FunctionCallExpr:
				tree := sel.Expr.(parser.FunctionCallExpr)
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

				fn, ok := prev.(*eclaType.Struct).Fields[tree.Name]
				if !ok {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "field "+tree.Name+" does not exist", errorHandler.LevelFatal)
				}
				var foo *eclaType.Function
				switch (*fn).(type) {
				case *eclaType.Function:
					foo = (*fn).(*eclaType.Function)
				}
				r, err := RunFunctionCallExprWithArgs(tree.Name, env, foo, args)
				if err != nil {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
				}
				var retValues []*Bus
				for _, v := range r {
					retValues = append(retValues, NewMainBus(v))
				}
				if len(retValues) == 1 {
					prev = retValues[0].GetVal()
				} else {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN RunSelectorExpr.\nPlease open issue", errorHandler.LevelFatal)
				}
			case parser.IndexableAccessExpr:
				tree := sel.Expr.(parser.IndexableAccessExpr)
				s := prev.(*eclaType.Struct)
				result, ok := s.Fields[tree.VariableName]
				if !ok {
					env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "field "+tree.VariableName+" does not exist", errorHandler.LevelFatal)
				}
				for i := range tree.Indexes {
					BusCollection := RunTree(tree.Indexes[i], env)
					if IsMultipleBus(BusCollection) {
						env.ErrorHandle.HandleError(tree.Indexes[i].StartLine(), tree.Indexes[i].StartPos(), "MULTIPLE BUS IN RunIndexableAccessExpr\nPlease open issue", errorHandler.LevelFatal)
					}
					elem := BusCollection[0].GetVal()
					temp, err := (*result).GetIndex(elem)

					result = temp

					if err != nil {
						env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
					}
				}
				prev = *result
			default:
				fmt.Printf("%T\n", expr.Sel)
			}
			return RunSelectorExpr(sel, env, prev)
		case parser.IndexableAccessExpr:
			tree := expr.Sel.(parser.IndexableAccessExpr)
			s := prev.(*eclaType.Struct)
			result, ok := s.Fields[tree.VariableName]
			if !ok {
				env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "field "+tree.VariableName+" does not exist", errorHandler.LevelFatal)
			}
			for i := range tree.Indexes {
				BusCollection := RunTree(tree.Indexes[i], env)
				if IsMultipleBus(BusCollection) {
					env.ErrorHandle.HandleError(tree.Indexes[i].StartLine(), tree.Indexes[i].StartPos(), "MULTIPLE BUS IN RunIndexableAccessExpr\nPlease open issue", errorHandler.LevelFatal)
				}
				elem := BusCollection[0].GetVal()
				temp, err := (*result).GetIndex(elem)

				result = temp

				if err != nil {
					env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
				}
			}
			return []*Bus{NewMainBus(*result)}
		default:
			env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "struct cannot have filed of type "+prev.GetType(), errorHandler.LevelFatal)
		}
	default:
		env.ErrorHandle.HandleError(expr.StartLine(), expr.StartPos(), "type "+prev.GetType()+" has no fields", errorHandler.LevelFatal)
	}
	return []*Bus{NewNoneBus()}
}

func RunStructInstantiationExpr(tree parser.StructInstantiationExpr, env *Env) []*Bus {
	decl, ok := env.GetTypeDecl(tree.Name)
	if !ok {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "unknown type: "+tree.Name, errorHandler.LevelFatal)
	}
	s := eclaType.NewStruct(decl.(*eclaDecl.StructDecl))
	s.SetType(tree.Name)
	for i, arg := range tree.Args {
		val := RunTree(arg, env)
		if IsMultipleBus(val) {
			env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), "MULTIPLE BUS IN StructInstantiationExpr.\nPlease open issue", errorHandler.LevelFatal)
		}
		temp := val[0].GetVal()
		switch temp.(type) {
		case *eclaType.Var:
			temp = temp.(*eclaType.Var).GetValue().(eclaType.Type)
		}
		s.AddField(i, temp)
	}
	err := s.Verify()
	if err != nil {
		env.ErrorHandle.HandleError(tree.StartLine(), tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	return []*Bus{NewMainBus(s)}
}
