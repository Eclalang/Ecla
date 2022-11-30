package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaKeyWord"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
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
		return env.GetVar(t.Value)
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
	case parser.VariableDecl:
		return RunVariableDecl(tree.(parser.VariableDecl), env)
	}
	return nil
}

// RunVariableDecl executes a parser.VariableDecl.
func RunVariableDecl(tree parser.VariableDecl, env *Env) eclaType.Type {
	if tree.Value == nil {
		switch tree.Type {
		case parser.Int:
			env.SetVar(tree.Name, eclaKeyWord.NewVar(tree.Name, eclaType.NewInt("0")))
		case parser.String:
			env.SetVar(tree.Name, eclaKeyWord.NewVar(tree.Name, eclaType.NewString("")))
		case parser.Bool:
			env.SetVar(tree.Name, eclaKeyWord.NewVar(tree.Name, eclaType.NewBool("false")))
		case parser.Float:
			env.SetVar(tree.Name, eclaKeyWord.NewVar(tree.Name, eclaType.NewFloat("0")))
		}
	} else {
		env.Vars[tree.Name] = eclaKeyWord.NewVar(tree.Name, RunTree(tree.Value, env))
	}
	return nil
}

// RunPrintStmt executes a parser.PrintStmt.
func RunPrintStmt(tree parser.PrintStmt, env *Env) eclaType.Type {
	fmt.Print(RunTree(tree.Expression, env).GetString())
	return nil
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
	}
	return nil
}
