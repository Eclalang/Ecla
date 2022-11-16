package interpreter

import (
	"fmt"
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
		fmt.Println(v)
		fmt.Printf("type of Node %T\n", v)
	}
}
