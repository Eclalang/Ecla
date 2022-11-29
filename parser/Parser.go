package parser

import (
	"github.com/tot0p/Ecla/lexer"
	"log"
)

const (
	LowestPrecedence  = 0
	HighestPrecedence = 6
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
	tempFile := File{ParseTree: new(AST), Trace: ""}
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
	log.Fatal("Expected keyword")
	return nil
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
	return p.ParseBinaryExpr(nil, LowestPrecedence+1)
}

// ParseBinaryExpr parses a binary expression with the given precedence
func (p *Parser) ParseBinaryExpr(Lhs Expr, precedence int) Expr {
	if Lhs == nil {
		Lhs = p.ParseUnaryExpr()
	}
	var n int
	for n = 1; ; n++ {
		opprec := TokenPrecedence(p.CurrentToken)
		operator := p.CurrentToken
		if opprec < precedence {
			return Lhs
		}
		p.Step()
		Rhs := p.ParseBinaryExpr(nil, opprec+1)
		Lhs = BinaryExpr{LeftExpr: Lhs, Operator: operator, RightExpr: Rhs}
	}

}

// ParseUnaryExpr parses a unary expression
func (p *Parser) ParseUnaryExpr() Expr {
	if p.CurrentToken.TokenType == lexer.ADD || p.CurrentToken.TokenType == lexer.SUB {
		Operator := p.CurrentToken
		p.Step()
		Rhs := p.ParseUnaryExpr()
		return UnaryExpr{Operator: Operator, RightExpr: Rhs}
	}
	return p.ParsePrimaryExpr(nil)

}

// ParsePrimaryExpr parses a primary expression
func (p *Parser) ParsePrimaryExpr(exp Expr) Expr {
	if exp == nil {
		exp = p.ParseOperand()
	}
	return exp
}

func (p *Parser) ParseOperand() Expr {
	if p.CurrentToken.TokenType == lexer.LPAREN {
		return p.ParseParenExpr()
	}
	return p.ParseLiteral()
}

func (p *Parser) ParseParenExpr() Expr {
	tempParentExpr := ParenExpr{}
	tempParentExpr.Lparen = p.CurrentToken
	p.Step()
	tempParentExpr.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected ')'")
	}
	tempParentExpr.Rparen = p.CurrentToken
	return tempParentExpr
}

// ParseLiteral parses a literal
func (p *Parser) ParseLiteral() Expr {
	if p.CurrentToken.TokenType == lexer.INT || p.CurrentToken.TokenType == lexer.FLOAT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.DQUOTE {
		p.Step()
		if p.CurrentToken.TokenType != lexer.STRING {
			log.Fatal("Expected string")
		}
		tempLiteral := Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: p.CurrentToken.Value}
		p.Step()
		if p.CurrentToken.TokenType != lexer.DQUOTE {
			log.Fatal("Expected '\"'")
		}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.BOOL {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	log.Fatal("Expected literal instead of " + p.CurrentToken.Value)
	return nil
}
