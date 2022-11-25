package parser

import (
	"github.com/tot0p/Ecla/lexer"
	"log"
	"strings"
)

const (
	LowestPrecedence  = 0
	HighestPrecedence = 6
)

// Parser is the parser for the Ecla language

type Tracer struct {
	DoesTrace   bool
	Nesting     int
	TraceString string
}

func (t *Tracer) Trace(currentNode Node) {
	if t.DoesTrace {
		t.TraceString += strings.Repeat("\t", t.Nesting) + currentNode.String()
	}
}

func (t *Tracer) Reset() {
	t.Nesting = 0
	t.TraceString = ""
}

type Parser struct {
	Tokens       []lexer.Token
	TokenIndex   int
	CurrentToken lexer.Token
	Tracer       Tracer
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
	FinalTrace := ""
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
		FinalTrace += p.Tracer.TraceString
		p.Tracer.Reset()
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

// ParseBinaryExpr parses a binary expression with the given precedence
// func (p *Parser) ParseBinaryExpr(exp Expr, prec int) Expr {
func (p *Parser) ParseBinaryExpr(exp Expr) Expr {
	if exp == nil {
		exp = p.ParseUnaryExpr()
	}
	//for {
	//	opprec := TokenPrecedence(p.CurrentToken)
	//	if opprec < prec {
	//		return exp
	//	}
	//	RightExpr := p.ParseBinaryExpr(nil, opprec+1)
	//	exp = BinaryExpr{LeftExpr: exp, Operator: p.CurrentToken, RightExpr: RightExpr}
	//}
	for p.CurrentToken.TokenType == lexer.ADD || p.CurrentToken.TokenType == lexer.SUB || p.CurrentToken.TokenType == lexer.MULT || p.CurrentToken.TokenType == lexer.DIV || p.CurrentToken.TokenType == lexer.MOD {
		Operator := p.CurrentToken
		p.Step()
		RightExpr := p.ParseBinaryExpr(nil)
		if RightExpr.precedence() > exp.precedence() {
			exp = BinaryExpr{LeftExpr: RightExpr, Operator: Operator, RightExpr: exp}
		} else {
			exp = BinaryExpr{LeftExpr: exp, Operator: Operator, RightExpr: RightExpr}
		}
	}

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
		//tempExpr := p.ParseBinaryExpr(nil, LowestPrecedence+1)
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
			panic("Expected string")
		}
		tempLiteral := Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: p.CurrentToken.Value}
		p.Step()
		if p.CurrentToken.TokenType != lexer.DQUOTE {
			panic("Expected '\"'")
		}
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
