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

func (p *Parser) Peek(lookAhead int) lexer.Token {
	if p.TokenIndex+lookAhead >= len(p.Tokens) {
		return lexer.Token{}
	}
	return p.Tokens[p.TokenIndex+lookAhead]
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
	if p.CurrentToken.Value == "for" {
		return p.ParseForStmt()
	}
	log.Fatal("Unexpected Keyword ", p.CurrentToken.Value)
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
	tempIf := IfStmt{IfToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		log.Fatal("Expected If LPAREN")
	}
	tempIf.LeftParen = p.CurrentToken
	p.Step()
	tempIf.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected If RPAREN")
	}
	tempIf.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		log.Fatal("Expected If LBRACE")
	}
	tempIf.LeftBrace = p.CurrentToken
	p.Step()
	tempIf.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		log.Fatal("Expected If RBRACE")
	}
	tempIf.RightBrace = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if p.CurrentToken.Value == "else" {
			tempIf.ElseStmt = p.ParseElseStmt()
		} else {
			tempIf.ElseStmt = nil
		}
	} else {
		tempIf.ElseStmt = nil
	}
	return tempIf
}

func (p *Parser) ParseElseStmt() *ElseStmt {
	tempElse := new(ElseStmt)
	tempElse.ElseToken = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if p.CurrentToken.Value == "if" {
			parsedIf := p.ParseIfStmt()
			point := parsedIf.(IfStmt)
			tempElse.IfStmt = &point
			return tempElse
		} else {
			tempElse.IfStmt = nil
		}
	} else {
		tempElse.IfStmt = nil
	}
	if p.CurrentToken.TokenType != lexer.LBRACE {
		log.Fatal("Expected Else LBRACE")
	}
	tempElse.LeftBrace = p.CurrentToken
	p.Step()
	tempElse.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		log.Fatal("Expected Else RBRACE")
	}
	tempElse.RightBrace = p.CurrentToken
	p.Step()
	return tempElse
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

func (p *Parser) ParseForStmt() Stmt {
	tempFor := ForStmt{}
	tempFor.ForToken = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		log.Fatal("Expected For LPAREN")
	}
	tempFor.LeftParen = p.CurrentToken
	p.Step()
	lookAhead := p.Peek(1)
	if lookAhead.TokenType != lexer.COMMA {
		tempFor.RangeToken = lexer.Token{}
		tempFor.InitDecl = p.ParseVariableDecl()
		if p.CurrentToken.TokenType != lexer.COMMA {
			log.Fatal("Expected Condition Expression instead of ", p.CurrentToken.Value)
		}
		p.Step()
		tempFor.CondExpr = p.ParseExpr()

		if p.CurrentToken.TokenType != lexer.COMMA {
			log.Fatal("Expected Post Expression")
		}
		p.Step()
		tempFor.PostAssignStmt = p.ParseVariableAssign()
	} else {
		tempFor.KeyToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.COMMA {
			log.Fatal("Expected For COMMA")
		}
		p.Step()
		tempFor.ValueToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.TEXT {
			log.Fatal("Unexpected Token : ", p.CurrentToken.Value)
		}
		if p.CurrentToken.Value != "range" {
			log.Fatal("Expected range")
		}
		tempFor.RangeToken = p.CurrentToken
		p.Step()
		tempFor.RangeExpr = p.ParseExpr()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		log.Fatal("Expected For RPAREN instead of ", p.CurrentToken.Value)
	}
	tempFor.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		log.Fatal("Expected For LBRACE instead of ", p.CurrentToken.Value)
	}
	tempFor.LeftBrace = p.CurrentToken
	p.Step()
	tempFor.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		log.Fatal("Expected For RBRACE instead of ", p.CurrentToken.Value)
	}
	tempFor.RightBrace = p.CurrentToken
	p.Step()
	return tempFor
}

func (p *Parser) ParseVariableDecl() Decl {
	tempDecl := VariableDecl{VarToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			log.Fatal("Variable name cannot be a keyword")
		}
		log.Fatal("Expected variable name instead of ", p.CurrentToken.Value)
	}
	tempDecl.Name = p.CurrentToken.Value
	typeName, isType := p.ParseType()
	if !isType {
		log.Fatal("Expected variable type instead of ", p.CurrentToken.Value)
	}
	tempDecl.Type = typeName
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		if p.CurrentToken.TokenType != lexer.COMMA {
			if p.CurrentToken.TokenType != lexer.EOL {
				if p.CurrentToken.TokenType != lexer.EOF {
					log.Fatal("Expected variable assignment instead of " + p.CurrentToken.Value)
				}
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

func (p *Parser) ParseType() (string, bool) {
	p.Step()
	if _, ok := VarTypes[p.CurrentToken.Value]; ok {
		tempType := ""
		for {
			if _, ok2 := VarTypes[p.CurrentToken.Value]; !ok2 {
				break
			}
			if p.CurrentToken.TokenType == lexer.LBRACKET {
				Peek := p.Peek(1)
				if Peek.TokenType != lexer.RBRACKET {
					return "", false
				}
				p.Step()
				p.Step()
				tempType += "[]"
				continue
			}
			tempType += p.CurrentToken.Value
			p.Step()
		}
		if tempType == "" {
			return "", false
		}
		return tempType, true
	}
	return "", false
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

func (p *Parser) ParseArrayLiteral() Expr {
	tempArrayExpr := ArrayLiteral{}
	tempArrayExpr.LBRACKET = p.CurrentToken
	p.Step()
	for {
		tempArrayExpr.Values = append(tempArrayExpr.Values, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COMMA {
			break
		}
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RBRACKET {
		log.Fatal("Expected ']'")
	}
	tempArrayExpr.RBRACKET = p.CurrentToken
	p.Step()
	return tempArrayExpr
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
	if p.CurrentToken.TokenType == lexer.LBRACKET {
		return p.ParseArrayLiteral()
	}
	log.Fatal("Expected literal instead of " + p.CurrentToken.Value)
	return nil
}
