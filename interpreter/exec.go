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

// Run executes the environment.
func Run(env *Env) {
	for _, v := range env.SyntaxTree.ParseTree.Operations {
		//txt, _ := json.MarshalIndent(v, "", "  ")
		//fmt.Println(string(txt))
		RunTree(v, env)
	}
}

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
	default:
		panic("Unknown type")
	}
}

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
	}
	return nil
}

func RunMethodCallExpr(expr parser.MethodCallExpr, env *Env) eclaType.Type {
	var args []eclaType.Type
	for _, v := range expr.FunctionCall.Args {
		args = append(args, RunTree(v, env))
	}

	return env.Libs[expr.ObjectName].Call(expr.FunctionCall.Name, args)
}

func RunImportStmt(stmt parser.ImportStmt, env *Env) {
	env.Import(stmt.ModulePath)
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
		}
	} else {
		if eclaType.IsList(tree.Type) {
			l, err := eclaType.NewList(tree.Type)
			if err != nil {
				panic(err)
			}
			// err into this
			t := RunTree(tree.Value, env)
			// check type
			switch t.(type) {
			case *eclaType.List:
				list := t.(*eclaType.List)
				if list.GetFullType() == "empty" {
					list.SetType(tree.Type)
				} else {
					if list.GetFullType() != tree.Type {
						panic(errors.New("type mismatch"))
					}
				}
			default:
				panic(errors.New("cannot assign non-list to list"))
			}
			err = l.SetValue(t)
			// err into this end
			if err != nil {
				panic(err)
			}
			v, err := eclaType.NewVar(tree.Name, tree.Type, l)
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		} else {
			v, err := eclaType.NewVar(tree.Name, tree.Type, RunTree(tree.Value, env))
			if err != nil {
				panic(err)
			}
			env.SetVar(tree.Name, v)
		}
	}
	return nil
}

// RunPrintStmt executes a parser.PrintStmt.
func RunPrintStmt(tree parser.PrintStmt, env *Env) eclaType.Type {
	fmt.Print(RunTree(tree.Expression, env).GetString())
	return nil
}

// RunTypeStmt executes a parser.TypeStmt.
func RunTypeStmt(tree parser.TypeStmt, env *Env) eclaType.Type {
	t := RunTree(tree.Expression, env)
	var typ string

	switch t.(type) {
	case *eclaType.Var:
		t = t.(*eclaType.Var).GetValue().(eclaType.Type)
		switch t.(type) {
		case *eclaType.List:
			typ = t.(*eclaType.List).GetFullType()
		default:
			typ = t.GetType()
		}
	case *eclaType.List:
		typ = t.(*eclaType.List).GetFullType()
	default:
		typ = t.GetType()
	}
	fmt.Println(typ)
	return nil
	//return eclaType.NewString(RunTree(tree.Expression, env).GetType())
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
	v, ok := env.GetVar(tree.Name)
	if !ok {
		panic(errors.New("variable not found"))
	}
	v.SetVar(RunTree(tree.Value, env))
}

// RunWhileStmt
func RunWhileStmt(tree parser.WhileStmt, env *Env) {
	while := eclaKeyWord.NewWhile(tree.Cond, tree.Body)
	for RunTree(while.Condition, env).GetString() == "true" { //TODO add error
		for _, stmt := range while.Body {
			RunTree(stmt, env)
		}
	}
}

func RunForStmt(For parser.ForStmt, env *Env) {
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
			typ = list.(*eclaType.List).GetFullType()[2:]
			l = list.(*eclaType.List).Len()
		case eclaType.String:
			typ = list.GetType()
			l = list.(eclaType.String).Len()
		case *eclaType.Var:
			temp := list.(*eclaType.Var).GetValue()
			//fmt.Printf("%T", temp)
			switch temp.(type) {
			case *eclaType.List:
				typ = temp.(*eclaType.List).GetFullType()[2:]
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
		for _, stmt := range tree.Body {
			RunTree(stmt, env)
		}
	} else if tree.ElseStmt != nil {
		if tree.ElseStmt.IfStmt != nil {
			RunIfStmt(*tree.ElseStmt.IfStmt, env)
		} else {
			for _, stmt := range tree.ElseStmt.Body {
				RunTree(stmt, env)
			}
		}
	}
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
		if values[0].GetType() == "list" {
			switch values[0].(type) {
			case *eclaType.List:
				t := values[0].(*eclaType.List)
				typ = "[]" + t.GetFullType()
			}
		} else {
			typ = "[]" + values[0].GetType()
		}
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
