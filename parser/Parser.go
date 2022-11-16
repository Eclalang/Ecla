package parser

import (
	"fmt"
	"github.com/tot0p/Ecla/lexer"
	"log"
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
	fmt.Println(p.Tokens)
	tempFile := File{ParseTree: new(AST)}
	for p.CurrentToken.TokenType != lexer.EOF {
		if p.CurrentToken.TokenType != lexer.TEXT {
			tempFile.ParseTree.Operations = append(tempFile.ParseTree.Operations, p.ParseExpr())
			if p.CurrentToken.TokenType != lexer.EOL {
				log.Fatal("Expected EOL")
			}
		} else {
			if p.CurrentToken.Value == "\n" || p.CurrentToken.Value == "\r" {

			} else {
				tempFile.ParseTree.Operations = append(tempFile.ParseTree.Operations, p.ParseKeyword())
			}
		}
		p.Step()
	}
	p.Step()
	return tempFile

}

func (p *Parser) ParseKeyword() Stmt {
	if p.CurrentToken.TokenType == lexer.TEXT {
		if p.CurrentToken.Value == "print" {
			return p.ParsePrintStmt()
		}

	}
	panic("Expected keyword")
}

func (p *Parser) ParsePrintStmt() Stmt {
	tempPrint := PrintStmt{PrintToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		log.Fatal("Expected Print LPAREN")
	}
	tempPrint.Lparen = p.CurrentToken
	p.Step()
	tempPrint.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected Print RPAREN")
	}
	tempPrint.Rparen = p.CurrentToken
	p.Step()
	return tempPrint
}

func (p *Parser) ParseExpr() Expr {
	return p.ParseBinaryExpr(nil)
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
	if p.CurrentToken.TokenType == lexer.TEXT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: "STRING", Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.BOOL {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	panic("Expected literal")

}
