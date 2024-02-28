package parser

import "github.com/Eclalang/Ecla/lexer"

type FunctionParams struct {
	Name string
	Type string
}

type FunctionPrototype struct {
	LeftParamParen  lexer.Token
	RightParamParen lexer.Token
	Parameters      []FunctionParams
	LeftRetsParen   lexer.Token
	RightRetsParen  lexer.Token
	ReturnTypes     []string
	LeftBrace       lexer.Token
	RightBrace      lexer.Token
}

type StructField struct {
	Name string
	Type string
}
