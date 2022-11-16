package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

/*
type Env struct {
	Vars       map[string]eclaType.Type
	OS         string
	ARCH       string
	SyntaxTree any
	File       string
	Code       string
}
*/

func Run(env *Env) {
	for _, v := range env.SyntaxTree.ParseTree.Operations {
		//txt, _ := json.MarshalIndent(v, "", "  ")
		//fmt.Println(string(txt))
		RunTree(v)
	}
}

func RunTree(tree parser.Node) eclaType.Type {
	//fmt.Printf("%T\n", tree)
	switch tree.(type) {
	case parser.Literal:
		t := tree.(parser.Literal)
		switch t.Type {
		case lexer.INT:
			return eclaType.NewInt(t.Value)
		case "STRING":
			return eclaType.NewString(t.Value)
		}
	case parser.BinaryExpr:
		return RunBinaryExpr(tree.(parser.BinaryExpr))
	case parser.UnaryExpr:
		return RunUnaryExpr(tree.(parser.UnaryExpr))
	case parser.ParenExpr:
		return RunTree(tree.(parser.ParenExpr).Expression)
	case parser.PrintStmt:
		return RunPrintStmt(tree.(parser.PrintStmt))
	}
	return nil
}

func RunPrintStmt(tree parser.PrintStmt) eclaType.Type {
	fmt.Print(RunTree(tree.Expression))
	return nil
}

func RunBinaryExpr(tree parser.BinaryExpr) eclaType.Type {
	//fmt.Printf("%T\n", tree)
	left := RunTree(tree.LeftExpr)
	right := RunTree(tree.RightExpr)
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

func RunUnaryExpr(tree parser.UnaryExpr) eclaType.Type {
	return nil
}
