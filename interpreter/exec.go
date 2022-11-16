package interpreter

import (
	"encoding/json"
	"fmt"

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
		txt, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(txt))
		RunTree(v)
	}
}

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func RunTree(tree parser.Node) {

	fmt.Printf("%T ", tree)
	switch Typeof(tree) {
	case "parser.ParenExpr":
		RunTree(tree.(parser.ParenExpr).Expression)
	case "parser.BinaryExpr":
		RunTree(tree.(parser.BinaryExpr).LeftExpr)
		RunTree(tree.(parser.BinaryExpr).RightExpr)
	case "parser.Literal":
		fmt.Print(tree.(parser.Literal).Value)
	}
	fmt.Println()
}
