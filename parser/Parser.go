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
	CurrentFile  *File
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

func (p *Parser) Peek() lexer.Token {
	if p.TokenIndex+1 >= len(p.Tokens) {
		return lexer.Token{}
	}
	return p.Tokens[p.TokenIndex+1]
}

//Recursive descent parser

func (p *Parser) Parse() *File {
	p.TokenIndex = -1
	p.Step()
	return p.ParseFile()
}

func (p *Parser) ParseFile() *File {
	tempFile := new(File)
	tempFile.ParseTree = new(AST)
	p.CurrentFile = tempFile
	for p.CurrentToken.TokenType != lexer.EOF {
		tempFile.ParseTree.Operations = append(tempFile.ParseTree.Operations, p.ParseNode())
		p.Step()
	}
	p.Step()
	return tempFile

}

func (p *Parser) ParseNode() Node {
	if p.CurrentToken.TokenType != lexer.TEXT {
		tempExpr := p.ParseExpr()

		return tempExpr
	} else {
		if p.CurrentToken.Value == "\n" || p.CurrentToken.Value == "\r" {

		} else {
			tempExpr := p.ParseText()
			if p.CurrentToken.TokenType != lexer.EOL {
				log.Fatal("Expected EOL")
			}
			return tempExpr
		}
	}
	return nil
}

func (p *Parser) ParseBody() []Node {
	tempBody := make([]Node, 0)
	for p.CurrentToken.TokenType != lexer.RBRACE {
		tempBody = append(tempBody, p.ParseNode())
		p.Step()
	}
	return tempBody
}

func (p *Parser) ParseText() Node {
	if _, ok := Keywords[p.CurrentToken.Value]; ok {
		return p.ParseKeyword()
	} else {
		return p.ParseIdent()
	}
}

func (p *Parser) ParseKeyword() Node {
	if p.CurrentToken.Value == "var" {
		return p.ParseVariableDecl()
	}
	if p.CurrentToken.Value == "print" {
		return p.ParsePrintStmt()
	}
	if p.CurrentToken.Value == "type" {
		return p.ParseTypeStmt()
	}
	if p.CurrentToken.Value == "if" {
		return p.ParseIfStmt()
	}
	if p.CurrentToken.Value == "while" {
		return p.ParseWhileStmt()
	}
	return nil
}

func (p *Parser) ParseIdent() Node {
	return p.ParseVariableAssign()
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

func (p *Parser) ParseTypeStmt() Stmt {
	tempType := TypeStmt{TypeToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		log.Fatal("Expected Type LPAREN")
	}
	tempType.Lparen = p.CurrentToken
	p.Step()
	tempType.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected Type RPAREN")
	}
	tempType.Rparen = p.CurrentToken
	p.Step()
	return tempType
}

func (p *Parser) ParseIfStmt() Stmt {
	return nil
}

func (p *Parser) ParseWhileStmt() Stmt {
	tempWhile := WhileStmt{WhileToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		log.Fatal("Expected While LPAREN")
	}
	tempWhile.LeftParen = p.CurrentToken
	p.Step()
	tempWhile.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected While RPAREN")
	}
	tempWhile.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		log.Fatal("Expected While LBRACE")
	}
	tempWhile.LeftBrace = p.CurrentToken
	p.Step()
	tempWhile.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		log.Fatal("Expected While RBRACE")
	}
	tempWhile.RightBrace = p.CurrentToken
	p.Step()
	return tempWhile
}

func (p *Parser) ParseVariableDecl() Decl {
	tempDecl := VariableDecl{VarToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			log.Fatal("Variable name cannot be a keyword")
		}
		log.Fatal("Expected variable name")
	}
	tempDecl.Name = p.CurrentToken.Value
	p.Step()
	_, ok := VarTypes[p.CurrentToken.Value]
	if !ok || p.CurrentToken.TokenType != lexer.TEXT {
		log.Fatal("Expected variable type")
	}
	tempDecl.Type = p.CurrentToken.Value
	p.Step()
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		if p.CurrentToken.TokenType != lexer.EOL {
			if p.CurrentToken.TokenType != lexer.EOF {
				log.Fatal("Expected variable assignment instead of " + p.CurrentToken.Value)
			}
		}
		tempDecl.Value = nil
		p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
		return tempDecl
	}
	p.Step()
	tempDecl.Value = p.ParseExpr()
	p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
	return tempDecl

}

func (p *Parser) ParseVariableAssign() Stmt {
	Var := p.CurrentToken
	VarName := Var.Value
	p.Step()
	switch p.CurrentToken.TokenType {
	case lexer.ASSIGN:
		p.Step()
		return VariableAssignStmt{VarToken: Var, Name: VarName, Value: p.ParseExpr()}
	case lexer.INC:
		p.Step()
		return VariableIncrementStmt{VarToken: Var, Name: VarName, IncToken: p.CurrentToken}
	case lexer.DEC:
		p.Step()
		return VariableDecrementStmt{VarToken: Var, Name: VarName, DecToken: p.CurrentToken}
	}
	return nil
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
	if p.CurrentToken.TokenType == lexer.ADD || p.CurrentToken.TokenType == lexer.SUB || p.CurrentToken.TokenType == lexer.NOT {
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
	p.Step()
	return tempParentExpr
}

// ParseLiteral parses a literal
func (p *Parser) ParseLiteral() Expr {
	if p.CurrentToken.TokenType == lexer.INT || p.CurrentToken.TokenType == lexer.FLOAT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.TEXT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: "VAR", Value: p.CurrentToken.Value}
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
