package interpreter

import (
	"fmt"
	"strconv"

	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/interpreter/eclaKeyWord"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
)

// RunImportStmt executes a parser.ImportStmt.
func RunImportStmt(stmt parser.ImportStmt, env *Env) {
	env.Import(stmt)
}

// RunTypeStmt executes a parser.TypeStmt.
func RunTypeStmt(tree parser.TypeStmt, env *Env) {
	t := RunTree(tree.Expression, env)[0].GetVal()
	var typ string
	typ = t.GetType()
	fmt.Println(typ)
	//return eclaType.NewString(RunTree(tree.Expression, env).GetType())
}

// AssignementTypeChecking checks if the type of the variable is the same as the type of the expression.
// If the type of the variable is any, it returns true else it returns false.
func AssignementTypeChecking(tree parser.VariableAssignStmt, type1 string, type2 string, env *Env) bool {
	if type1[:3] == parser.Any {
		return true
	}
	if type1 != type2 {
		env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Can't assign %s to %s", type2, type1), errorHandler.LevelFatal)
	}
	return false
}

// TODO : Remove this function @mkarten
func HandleError(tree parser.VariableAssignStmt, err error, env *Env) {
	if err != nil {
		env.ErrorHandle.HandleError(0, tree.StartPos(), err.Error(), errorHandler.LevelFatal)
	}
}

func getPointerToSelectorExpr(tree parser.SelectorExpr, env *Env, parent *eclaType.Type) *eclaType.Type {
	var temp *eclaType.Type
	if parent != nil {
		temp = parent
	} else {
		switch tree.Expr.(type) {
		case parser.Literal:
			if tree.Expr.(parser.Literal).Type == "VAR" {
				variable, ok := env.GetVar(tree.Expr.(parser.Literal).Value)
				if !ok {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "variable not found", errorHandler.LevelFatal)
				}
				temp = &variable.Value
			} else {
				env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Cant run assignement on type %s", tree.Expr.(parser.Literal).Type), errorHandler.LevelFatal)
			}
		case parser.IndexableAccessExpr:
			temp = IndexableAssignmentChecks(tree.Expr.(parser.IndexableAccessExpr), env)
		default:
			env.ErrorHandle.HandleError(0, tree.StartPos(), "Cant run assignement on type", errorHandler.LevelFatal)
		}
	}

	if parent != nil {
		switch tree.Expr.(type) {
		case parser.Literal:
			if tree.Expr.(parser.Literal).Type == "VAR" {
				switch (*temp).(type) {
				case *eclaType.Struct:
					temp = (*temp).(*eclaType.Struct).GetField(tree.Expr.(parser.Literal).Value)
					if temp == nil {
						env.ErrorHandle.HandleError(0, tree.StartPos(), "field not found", errorHandler.LevelFatal)
					}
				}
			} else {
				env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Cant run assignement on type %s", tree.Expr.(parser.Literal).Type), errorHandler.LevelFatal)
			}
		case parser.IndexableAccessExpr:
			switch (*temp).(type) {
			case *eclaType.Struct:
				temp = (*temp).(*eclaType.Struct).GetField(tree.Expr.(parser.IndexableAccessExpr).VariableName)
				if temp == nil {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "field not found", errorHandler.LevelFatal)
				}
				for i := range tree.Expr.(parser.IndexableAccessExpr).Indexes {
					busCollection := RunTree(tree.Expr.(parser.IndexableAccessExpr).Indexes[i], env)
					if IsMultipleBus(busCollection) {
						env.ErrorHandle.HandleError(0, tree.StartPos(), "MULTIPLE BUS IN getPointerToSelectorExpr", errorHandler.LevelFatal)
					}
					elem := busCollection[0].GetVal()
					var err error
					temp, err = (*temp).GetIndex(elem)
					if err != nil {
						env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: "+err.Error(), errorHandler.LevelFatal)
					}
				}
			}
		}
	}

	switch tree.Sel.(type) {
	case parser.Literal:
		if tree.Sel.(parser.Literal).Type == "VAR" {
			switch (*temp).(type) {
			case *eclaType.Struct:
				temp = (*temp).(*eclaType.Struct).GetField(tree.Sel.(parser.Literal).Value)
				if temp == nil {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "field not found", errorHandler.LevelFatal)
				}
			}
		} else {
			env.ErrorHandle.HandleError(0, tree.StartPos(), fmt.Sprintf("Cant run assignement on type %s", tree.Sel.(parser.Literal).Type), errorHandler.LevelFatal)
		}
	case parser.IndexableAccessExpr:
		switch (*temp).(type) {
		case *eclaType.Struct:
			temp = (*temp).(*eclaType.Struct).GetField(tree.Sel.(parser.IndexableAccessExpr).VariableName)
			if temp == nil {
				env.ErrorHandle.HandleError(0, tree.StartPos(), "field not found", errorHandler.LevelFatal)
			}
			for i := range tree.Sel.(parser.IndexableAccessExpr).Indexes {
				busCollection := RunTree(tree.Sel.(parser.IndexableAccessExpr).Indexes[i], env)
				if IsMultipleBus(busCollection) {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "MULTIPLE BUS IN getPointerToSelectorExpr", errorHandler.LevelFatal)
				}
				elem := busCollection[0].GetVal()
				var err error
				temp, err = (*temp).GetIndex(elem)
				if err != nil {
					env.ErrorHandle.HandleError(0, tree.StartPos(), "indexable variable assign: "+err.Error(), errorHandler.LevelFatal)
				}
			}
		}
	case parser.SelectorExpr:
		temp = getPointerToSelectorExpr(tree.Sel.(parser.SelectorExpr), env, temp)
	default:
		env.ErrorHandle.HandleError(0, tree.StartPos(), "Cant run assignement on type", errorHandler.LevelFatal)
	}

	return temp
}

// RunVariableAssignStmt Run assigns a variable.
func RunVariableAssignStmt(tree parser.VariableAssignStmt, env *Env) {
	var exprs []eclaType.Type
	var exprsTypes []string
	var vars []*eclaType.Type
	var varsTypes []string
	// Calculate values of expressions behind the variable
	if tree.Values[0] != nil {
		for _, v := range tree.Values {
			busses := RunTree(v, env)
			for _, bus := range busses {
				busVal := bus.GetVal()
				switch busVal.(type) {
				case *eclaType.List:
					busVal = busVal.GetValue().(eclaType.Type)
				case *eclaType.Map:
					busVal = busVal.GetValue().(eclaType.Type)
				case *eclaType.Var:
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
			temp := IndexableAssignmentChecks(v.(parser.IndexableAccessExpr), env)
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
		case parser.SelectorExpr:
			variable := getPointerToSelectorExpr(v.(parser.SelectorExpr), env, nil)
			vars = append(vars, variable)
			varsTypes = append(varsTypes, (*variable).GetType())
		}
	}

	PreExecLen := len(exprs)
	NamesLen := len(tree.Names)
	opp := tree.Operator

	if PreExecLen == NamesLen {
		switch opp {
		case parser.ASSIGN:
			for i := 0; i < NamesLen; i++ {
				switch exprs[i].(type) {
				case *eclaType.Function:
					fnTemp := exprs[i].(*eclaType.Function)
					switch (*vars[i]).(type) {
					case *eclaType.Function:
						fn := (*vars[i]).(*eclaType.Function)
						err := fn.Override(fnTemp.Args[0], fnTemp.GetBody(), fnTemp.GetReturn())
						if err != nil {
							env.ErrorHandle.HandleError(0, 0, err.Error(), errorHandler.LevelFatal)
						}
					case *eclaType.Any:
						tmp := (*vars[i]).(*eclaType.Any)
						switch tmp.Value.(type) {
						case *eclaType.Function:
							fn := tmp.Value.(*eclaType.Function)
							err := fn.Override(fnTemp.Args[0], fnTemp.GetBody(), fnTemp.GetReturn())
							if err != nil {
								env.ErrorHandle.HandleError(0, 0, err.Error(), errorHandler.LevelFatal)
							}
						default:
							err := tmp.SetAny(fnTemp)
							if err != nil {
								env.ErrorHandle.HandleError(0, 0, err.Error(), errorHandler.LevelFatal)
							}
						}
					default:
						fmt.Printf("%T\n", *vars[i])
						env.ErrorHandle.HandleError(0, 0, "cannot assign function to none function", errorHandler.LevelFatal)
					}
				default:
					isAny := AssignementTypeChecking(tree, varsTypes[i], exprsTypes[i], env)
					if isAny {
						*vars[i] = eclaType.NewAny(exprs[i])
					} else {
						*vars[i] = exprs[i]
					}
				}
			}
		case parser.ADDASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).Add(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
			}
		case parser.SUBASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).Sub(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
			}
		case parser.DIVASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).Div(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
			}
		case parser.MODASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).Mod(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
			}
		case parser.QOTASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).DivEc(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
			}

		case parser.MULTASSIGN:
			for i := 0; i < NamesLen; i++ {
				temp, err := (*vars[i]).Mul(exprs[i])
				HandleError(tree, err, env)
				isAny := AssignementTypeChecking(tree, varsTypes[i], temp.GetType(), env)
				if isAny {
					*vars[i] = eclaType.NewAny(temp)
				} else {
					*vars[i] = temp
				}
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
func IndexableAssignmentChecks(index parser.IndexableAccessExpr, env *Env) *eclaType.Type {
	v, ok := env.GetVar(index.VariableName)
	if !ok {
		env.ErrorHandle.HandleError(0, index.StartPos(), "indexable variable assign: variable not found", errorHandler.LevelFatal)
	}
	t := eclaType.Type(v)
	var temp = &t
	for i := range index.Indexes {
		busCollection := RunTree(index.Indexes[i], env)
		if IsMultipleBus(busCollection) {
			env.ErrorHandle.HandleError(0, index.StartPos(), "indexable variable assign: MULTIPLE BUS IN INDEXABLEASSIGMENTCHECK", errorHandler.LevelFatal)
			return nil
		}
		elem := busCollection[0].GetVal()
		var err error
		//fmt.Printf("%T\n", result.GetValue())
		temp, err = (*temp).GetIndex(elem)
		if err != nil {
			env.ErrorHandle.HandleError(0, index.StartPos(), "indexable variable assign: "+err.Error(), errorHandler.LevelFatal)
		}
	}
	return temp
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
		env.NewScope(SCOPE_LOOP)
		for _, stmt := range while.Body {
			BusCollection2 := RunTree(stmt, env)
			if IsMultipleBus(BusCollection2) {
				env.ErrorHandle.HandleError(0, tree.StartPos(), "while: MULTIPLE BUS IN RunWhileStmt", errorHandler.LevelFatal)
			}
			temp := BusCollection2[0]

			// TODO: add break and continue
			// TODO add multiple bus
			if temp.IsReturn() {
				env.EndScope()
				return temp
			}
		}
		env.EndScope()
		BusCollection = RunTree(while.Condition, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, tree.StartPos(), "while: MULTIPLE BUS IN RunWhileStmt", errorHandler.LevelFatal)
		}
	}
	return NewNoneBus()
}

func generateForRangeKeys(max int) []eclaType.Type {
	var keys []eclaType.Type
	for i := 0; i < max; i++ {
		keys = append(keys, eclaType.NewInt(strconv.Itoa(i)))
	}
	return keys
}

// RunForStmt runs the for statement
func RunForStmt(For parser.ForStmt, env *Env) *Bus {
	env.NewScope(SCOPE_LOOP)
	defer env.EndScope()
	tokenEmpty := lexer.Token{}
	if For.RangeToken != tokenEmpty {
		f := eclaKeyWord.NewForRange([]eclaType.Type{}, For.RangeExpr, For.KeyToken, For.ValueToken, For.Body)
		BusCollection := RunTree(f.RangeExpr, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
		}
		list := BusCollection[0].GetVal()
		var typ string
		var l int //...
		//fmt.Printf("%T", list)
		switch list.(type) {
		case *eclaType.Var:
			list = list.(*eclaType.Var).Value
		}

		var k *eclaType.Var
		var err error

		var keys []eclaType.Type

		switch list.(type) {
		case *eclaType.List:
			typ = list.(*eclaType.List).GetType()[2:]
			k, err = eclaType.NewVar(f.KeyToken.Value, parser.Int, eclaType.NewInt("0"))
			if err != nil {
				env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			l = list.(*eclaType.List).Len()
			keys = generateForRangeKeys(l)
		case eclaType.String:
			typ = parser.Char
			k, err = eclaType.NewVar(f.KeyToken.Value, parser.Int, eclaType.NewInt("0"))
			if err != nil {
				env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			l = list.(eclaType.String).Len()
			keys = generateForRangeKeys(l)
		case *eclaType.Map:
			k, err = eclaType.NewVar(f.KeyToken.Value, list.(*eclaType.Map).TypKey, list.(*eclaType.Map).Keys[0])
			if err != nil {
				env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			typ = list.(*eclaType.Map).TypVal
			l = list.(*eclaType.Map).Len()
			keys = list.(*eclaType.Map).Keys
		default:
			env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), "type "+list.GetType()+" not supported", errorHandler.LevelFatal)
		}

		env.SetVar(f.KeyToken.Value, k)
		v, err := eclaType.NewVarEmpty(f.ValueToken.Value, typ)
		if err != nil {
			env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
		}
		env.SetVar(f.ValueToken.Value, v)
		for i := 0; i < l; i++ {
			err := k.SetVar(keys[i])
			if err != nil {
				return nil
			}
			val, err := list.GetIndex(keys[i])
			if err != nil {
				env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			err = v.SetVar(*val)
			if err != nil {
				env.ErrorHandle.HandleError(0, f.RangeExpr.StartPos(), err.Error(), errorHandler.LevelFatal)
			}
			env.NewScope(SCOPE_LOOP)
			for _, stmt := range f.Body {
				BusCollection2 := RunTree(stmt, env)
				if IsMultipleBus(BusCollection2) {
					env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
				}
				temp := BusCollection2[0]
				if temp.IsReturn() {
					env.EndScope()
					return temp
				}
			}
			env.EndScope()
		}
	} else {
		f := eclaKeyWord.NewForI([]eclaType.Type{}, For.Body, For.CondExpr, For.PostAssignStmt)
		RunTree(For.InitDecl, env)
		BusCollection := RunTree(f.Condition, env)
		if IsMultipleBus(BusCollection) {
			env.ErrorHandle.HandleError(0, f.Condition.StartPos(), "for: MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
		}
		for BusCollection[0].GetVal().GetString() == "true" {
			env.NewScope(SCOPE_LOOP)
			for _, stmt := range f.Body {
				BusCollection2 := RunTree(stmt, env)
				if IsMultipleBus(BusCollection2) {
					env.ErrorHandle.HandleError(0, stmt.StartPos(), "MULTIPLE BUS IN RunForStmt", errorHandler.LevelFatal)
				}
				temp := BusCollection2[0]
				if temp.IsReturn() {
					env.EndScope()
					return temp
				}
			}
			env.EndScope()
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

// RunMurlocStmt executes a parser.MurlocStmt.
func RunMurlocStmt(stmt parser.MurlocStmt, env *Env) {
	env.ErrorHandle.HandleError(0, stmt.StartPos(), "Mrgle, Mmmm Uuua !", errorHandler.LevelFatal)
}
