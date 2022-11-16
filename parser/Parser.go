package parser

import (
	"github.com/tot0p/Ecla/lexer"
)

// Parser is the parser for the Ecla language
type Parser struct {
	Tokens       []lexer.Token
	TokenIndex   int
	CurrentToken lexer.Token
}

// Step moves the parser to the next token
func (p *Parser) Step() {
	p.TokenIndex++
	if p.TokenIndex >= len(p.Tokens) {
		p.CurrentToken = lexer.Token{}
	} else {
		p.CurrentToken = p.Tokens[p.TokenIndex]
	}
}

//Recursive descent parser

func (p *Parser) Parse() File {
	p.TokenIndex = -1
	p.Step()
	return p.ParseFile()
}

func (p *Parser) ParseFile() File {
	tempFile := File{ParseTree: new(AST)}
	if p.CurrentToken.TokenType == lexer.LPAREN {
		tempFile.ParseTree.Operations = append(tempFile.ParseTree.Operations, p.ParseParenExpr())
	}

	p.Step()
	return tempFile

}

func (p *Parser) ParseParenExpr() Node {
	tempParentExpr := ParenExpr{}
	tempParentExpr.Lparen = p.CurrentToken
	p.Step()
	tempParentExpr.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected ')'")
	}
	tempParentExpr.Rparen = p.CurrentToken
	return tempParentExpr
}

func (p *Parser) ParseExpr() Expr {
	return p.ParseBinaryExpr(nil)
}

// ParseBinaryExpr parses a binary expression with the given precedence
func (p *Parser) ParseBinaryExpr(exp Expr) Expr {
	if exp == nil {
		exp = p.ParseUnaryExpr()
	}
	for p.CurrentToken.TokenType == lexer.ADD || p.CurrentToken.TokenType == lexer.SUB || p.CurrentToken.TokenType == lexer.MULT || p.CurrentToken.TokenType == lexer.DIV || p.CurrentToken.TokenType == lexer.MOD {
		Operator := p.CurrentToken
		p.Step()
		RightExpr := p.ParseUnaryExpr()
		exp = BinaryExpr{LeftExpr: exp, Operator: Operator, RightExpr: RightExpr}
	}

	//-----------------------------------//
	// for operation priority
	//-----------------------------------//

	//for p.CurrentToken.TokenType == lexer.MULT || p.CurrentToken.TokenType == lexer.DIV || p.CurrentToken.TokenType == lexer.MOD {
	//	Operator := p.CurrentToken
	//	p.Step()
	//	RightExpr := p.ParseUnaryExpr()
	//	exp = BinaryExpr{LeftExpr: exp, Operator: Operator, RightExpr: RightExpr}
	//}
	return exp

}

// ParseUnaryExpr parses a unary expression
func (p *Parser) ParseUnaryExpr() Expr {
	if p.CurrentToken.TokenType == lexer.ADD || p.CurrentToken.TokenType == lexer.SUB {
		Operator := p.CurrentToken
		p.Step()
		return UnaryExpr{Operator: Operator, RightExpr: p.ParseUnaryExpr()}
	}
	return p.ParsePrimaryExpr()

}

// ParsePrimaryExpr parses a primary expression
func (p *Parser) ParsePrimaryExpr() Expr {
	if p.CurrentToken.TokenType == lexer.LPAREN {
		tempLPAREN := p.CurrentToken
		p.Step()
		tempExpr := p.ParseBinaryExpr(nil)
		if p.CurrentToken.TokenType != lexer.RPAREN {
			panic("Expected ')'")
		}
		tempRPAREN := p.CurrentToken
		p.Step()
		return ParenExpr{Lparen: tempLPAREN, Expression: tempExpr, Rparen: tempRPAREN}
	}
	return p.ParseLiteral()

}

// ParseLiteral parses a literal
func (p *Parser) ParseLiteral() Expr {
	if p.CurrentToken.TokenType == lexer.INT || p.CurrentToken.TokenType == lexer.FLOAT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	panic("Expected literal")

}
