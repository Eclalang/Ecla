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
func RunPrintStmt(tree parser.PrintStmt, env *Env) {
	fmt.Print(RunTree(tree.Expression, env)[0].GetVal().GetString())
}

// RunImportStmt executes a parser.ImportStmt.
func RunImportStmt(stmt parser.ImportStmt, env *Env) {
	env.Import(stmt.ModulePath)
}

// RunTypeStmt executes a parser.TypeStmt.
func RunTypeStmt(tree parser.TypeStmt, env *Env) {
	t := RunTree(tree.Expression, env)[0].GetVal()
	var typ string
	typ = t.GetType()
	fmt.Println(typ)
	//return eclaType.NewString(RunTree(tree.Expression, env).GetType())
}

func AssignementTypeChecking(tree parser.VariableAssignStmt, type1 string, type2 string, env *Env) {
	if type1 != type2 {
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Can't assign %s to %s", type2, type1), errorHandler.LevelFatal)
	}
}

func HandleError(tree parser.VariableAssignStmt, err error, env *Env) {
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunVariableAssignStmt Run assigns a variable.
func RunVariableAssignStmt(tree parser.VariableAssignStmt, env *Env) {
	var exprs []eclaType.Type
	var exprsTypes []string
	var vars []*eclaType.Type
	var varsTypes []string
	if tree.Values[0] != nil {
		for _, v := range tree.Values {
			busses := RunTree(v, env)
			for _, bus := range busses {
				busVal := bus.GetVal()
				switch busVal.(type) {
				case *eclaType.List:
					continue
				case *eclaType.Map:
					continue
				case *eclaType.Var:
					busVal.GetType()
					busVal = busVal.GetValue().(eclaType.Type)
				}
				exprs = append(exprs, busVal)
				exprsTypes = append(exprsTypes, busVal.GetType())
			}
		}
	}
	for _, v := range tree.Names {
		switch v.(type) {
		case parser.IndexableAccessExpr:
			temp := IndexableAssignmentChecks(tree, v.(parser.IndexableAccessExpr), env)
			vars = append(vars, temp)
			varsTypes = append(varsTypes, (*temp).GetType())
		case parser.Literal:
			if v.(parser.Literal).Type == "VAR" {
				variable, ok := env.GetVar(v.(parser.Literal).Value)
				if !ok {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: variable not found", errorHandler.LevelFatal)
				}
				vars = append(vars, &(variable.Value))
				varsTypes = append(varsTypes, variable.Value.GetType())
			} else {
				env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Cant run assignement on type %s", tree.Names[0].(parser.Literal).Type), errorHandler.LevelFatal)
			}

		}
	}

	PreExecLen := len(exprs)
	NamesLen := len(tree.Names)
	opp := tree.Operator

	if PreExecLen == NamesLen {
		switch opp {
		case parser.ASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				*vars[i] = exprs[i]
			}
		case parser.ADDASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).Add(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.SUBASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).Sub(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.DIVASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).Div(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.MODASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).Mod(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.QOTASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).DivEc(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}

		case parser.MULTASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
				temp, err := (*vars[i]).Mul(exprs[i])
				HandleError(tree, err, env)
				*vars[i] = temp
			}

		default:
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("%s is not a valid assignement operator", tree.Operator), errorHandler.LevelFatal)
		}

	} else if PreExecLen == 1 && NamesLen > 1 {
		switch opp {
		case parser.ASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				*vars[i] = exprs[0]
			}
		case parser.ADDASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).Add(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.SUBASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).Sub(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.DIVASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).Div(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.MODASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).Mod(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.QOTASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).DivEc(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}

		case parser.MULTASSIGN:
			for i := 0; i < NamesLen; i++ {
				AssignementTypeChecking(tree, varsTypes[i], exprsTypes[0], env)
				temp, err := (*vars[i]).Mul(exprs[0])
				HandleError(tree, err, env)
				*vars[i] = temp
			}

		default:
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("%s is not a valid assignement operator", tree.Operator), errorHandler.LevelFatal)
		}
	} else if PreExecLen == 0 && NamesLen >= 1 {
		switch opp {
		case parser.INCREMENT:
			for i := 0; i < NamesLen; i++ {
				addOne := eclaType.NewInt("1")
				AssignementTypeChecking(tree, varsTypes[i], addOne.GetType(), env)
				temp, err := (*vars[i]).Add(addOne)
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		case parser.DECREMENT:
			for i := 0; i < NamesLen; i++ {
				subOne := eclaType.NewInt("1")
				AssignementTypeChecking(tree, varsTypes[i], subOne.GetType(), env)
				temp, err := (*vars[i]).Sub(subOne)
				HandleError(tree, err, env)
				*vars[i] = temp
			}
		default:
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("%s is not a valid assignement operator", tree.Operator), errorHandler.LevelFatal)
		}

	} else {
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Invalid assignment: %d rValues to %d lValues", PreExecLen, NamesLen), errorHandler.LevelFatal)
	}
}

// IndexableAssignmentChecks checks if the indexable variable is valid
func IndexableAssignmentChecks(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) *eclaType.Type {
	v, ok := env.GetVar(index.VariableName)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: variable not found", errorHandler.LevelFatal)
	}
	t := eclaType.Type(v)
	var temp = &t
	for i := range index.Indexes {
		busCollection := RunTree(index.Indexes[i], env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN INDEXABLEASSIGMENTCHECK", errorHandler.LevelFatal)
			return nil
		}
		elem := busCollection[0].GetVal()
		var err error
		//fmt.Printf("%T\n", result.GetValue())
		temp, err = (*temp).GetIndex(elem)
		if err != nil {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: "+err.Error(), errorHandler.LevelFatal)
		}
	}
	return temp
}

// RunIndexableVariableAssignStmt runs the indexable variable assignment
func RunIndexableVariableAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) eclaType.Type {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableAssignStmt", errorHandler.LevelFatal)
		return nil
	}
	*temp = busCollection[0].GetVal()
	return nil
}

// RunVariableNonIndexableAssignStmt runs the variable assignment
func RunVariableNonIndexableAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN VariableNonIndexableAssignStmt", errorHandler.LevelFatal)
		return
	}
	temp := BusCollection[0].GetVal()
	err := v.SetVar(temp)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)

	}
}

// RunIndexableVariableIncrementStmt runs the indexable variable increment
func RunIndexableVariableIncrementStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	res, err := (*temp).Add(eclaType.Int(1))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableIncrementStmt runs the variable increment
func RunVariableIncrementStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	v.Increment()
}

// RunIndexableVariableDecrementStmt runs the indexable variable decrement
func RunIndexableVariableDecrementStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	res, err := (*temp).Sub(eclaType.Int(1))
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableDecrementStmt runs the variable decrement
func RunVariableDecrementStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	v.Decrement()
}

// RunIndexableVariableAddAssignStmt runs the indexable variable add assign
func RunIndexableVariableAddAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableAddAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).Add(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableAddAssignStmt runs the variable add assign
func RunVariableAddAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableAddAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.Add(BusCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunIndexableVariableSubAssignStmt runs the indexable variable sub assign
func RunIndexableVariableSubAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableSubAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).Sub(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableSubAssignStmt runs the variable sub assign
func RunVariableSubAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableSubAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.Sub(BusCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunIndexableVariableDivAssignStmt runs the indexable variable div assign
func RunIndexableVariableDivAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableDivAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).Div(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableDivAssignStmt runs the variable div assign
func RunVariableDivAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableDivAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.Div(BusCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunIndexableVariableModAssignStmt runs the indexable variable mod assign
func RunIndexableVariableModAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableModAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).Mod(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableModAssignStmt runs the variable mod assign
func RunVariableModAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableModAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.Mod(BusCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunIndexableVariableQotAssignStmt runs the indexable variable qot assign
func RunIndexableVariableQotAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableQotAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).DivEc(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableQotAssignStmt runs the variable qot assign
func RunVariableQotAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	BusCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableQotAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.DivEc(BusCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunIndexableVariableMultAssignStmt runs the indexable variable mult assign
func RunIndexableVariableMultAssignStmt(tree parser.VariableAssignStmt, index parser.IndexableAccessExpr, env *Env) {
	temp := IndexableAssignmentChecks(tree, index, env)
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: MULTIPLE BUS IN IndexableVariableMultAssignStmt", errorHandler.LevelFatal)
	}
	res, err := (*temp).Mul(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	*temp = res
}

// RunVariableMultAssignStmt runs the variable mult assign
func RunVariableMultAssignStmt(tree parser.VariableAssignStmt, variable parser.Literal, env *Env) {
	v, ok := env.GetVar(variable.Value)
	if !ok {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
	}
	busCollection := RunTree(tree.Values[0], env)
	if IsMultipleBus(busCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "variable assign: MULTIPLE BUS IN RunVariableMultAssignStmt", errorHandler.LevelFatal)
	}
	t, err := v.Mul(busCollection[0].GetVal())
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
	err = v.SetVar(t)
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

// RunWhileStmt runs the while statement
func RunWhileStmt(tree parser.WhileStmt, env *Env) *Bus {
	env.NewScope(SCOPE_LOOP)
	defer env.EndScope()
	while := eclaKeyWord.NewWhile(tree.Cond, tree.Body)
	BusCollection := RunTree(while.Condition, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "while: MULTIPLE BUS IN RunWhileStmt", errorHandler.LevelFatal)
	}
	for BusCollection[0].GetVal().GetString() == "true" { //TODO add error
		for _, stmt := range while.Body {
			BusCollection2 := RunTree(stmt, env)
			if IsMultipleBus(BusCollection2) {
				env.ErrorHandle.HandleError(0, tree.StartPos(), "while: MULTIPLE BUS IN RunWhileStmt", errorHandler.LevelFatal)
			}
			temp := BusCollection2[0]
			// TODO: add break and continue
			// TODO add multiple bus
			if temp.IsReturn() {
				return temp
			}
		}
		BusCollection = RunTree(while.Condition, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "while: MULTIPLE BUS IN RunWhileStmt", errorHandler.LevelFatal)
		}
	}
	return NewNoneBus()
}

// RunForStmt runs the for statement
func RunForStmt(For parser.ForStmt, env *Env) *Bus {
	env.NewScope(SCOPE_LOOP)
	defer env.EndScope()
	tokenEmpty := lexer.Token{}
	if For.RangeToken != tokenEmpty {
		f := eclaKeyWord.NewForRange([]eclaType.Type{}, For.RangeExpr, For.KeyToken, For.ValueToken, For.Body)
		k, err := eclaType.NewVar(f.KeyToken.Value, "int", eclaType.NewInt("0"))
		if err != nil {
			panic(err)
		}
		BusCollection := RunTree(f.RangeExpr, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
		}
		list := BusCollection[0].GetVal()
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
			err := k.SetVar(eclaType.NewInt(strconv.Itoa(i)))
			if err != nil {
				return nil
			}
			val, err := list.GetIndex(eclaType.Int(i))
			if err != nil {
				panic(err)
			}
			err = v.SetVar(*val)
			if err != nil {
				panic(err)
			}
			for _, stmt := range f.Body {
				BusCollection2 := RunTree(stmt, env)
				if IsMultipleBus(BusCollection2) {
					env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
				}
				temp := BusCollection2[0]
				if temp.IsReturn() {
					return temp
				}
			}
		}
	} else {
		f := eclaKeyWord.NewForI([]eclaType.Type{}, For.Body, For.CondExpr, For.PostAssignStmt)
		RunTree(For.InitDecl, env)
		BusCollection := RunTree(f.Condition, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, f.Condition.StartPos(), "for: MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
		}
		fmt.Println(BusCollection[0].GetVal())
		for BusCollection[0].GetVal().GetString() == "true" {
			for _, stmt := range f.Body {
				BusCollection2 := RunTree(stmt, env)
				if IsMultipleBus(BusCollection2) {
					env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
				}
				temp := BusCollection2[0]
				if temp.IsReturn() {
					return temp
				}
			}
			RunTree(f.Post, env)
			BusCollection = RunTree(f.Condition, env)
			if IsMultipleBus(BusCollection) {
				env.ErrorHandle.HandleError(0, f.Condition.StartPos(), "for: MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
			}
		}
	}
	return NewNoneBus()
}

// RunIfStmt runs the if statement
func RunIfStmt(tree parser.IfStmt, env *Env) *Bus {
	BusCollection := RunTree(tree.Cond, env)
	if IsMultipleBus(BusCollection) {
		env.ErrorHandle.HandleError(0, tree.StartPos(), "if: MULTIPLE BUS IN RunIfStmt", errorHandler.LevelFatal)
	}
	if BusCollection[0].GetVal().GetString() == "true" { //TODO add error
		env.NewScope(SCOPE_CONDITION)
		defer env.EndScope()
		for _, stmt := range tree.Body {
			BusCollection := RunTree(stmt, env)
			if IsMultipleBus(BusCollection) {
				env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunIfStmt", errorHandler.LevelFatal)
			}
			temp := BusCollection[0]
			if temp.IsReturn() {
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
				BusCollection := RunTree(stmt, env)
				if IsMultipleBus(BusCollection) {
					env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunIfStmt", errorHandler.LevelFatal)
				}
				temp := BusCollection[0]
				if temp.IsReturn() {
					return temp
				}
			}
		}
	}
	return NewNoneBus()
}

// RunReturnStmt runs the return statement
func RunReturnStmt(tree parser.ReturnStmt, env *Env) []eclaType.Type {
	var l []eclaType.Type
	for _, expr := range tree.ReturnValues {
		BusCollection := RunTree(expr, env)
		for _, bus := range BusCollection {
			l = append(l, bus.GetVal())
		}
	}
	return l
}
