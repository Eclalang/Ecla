package parser

import (
	"fmt"
	"github.com/tot0p/Ecla/errorHandler"
	"github.com/tot0p/Ecla/lexer"
	"log"
)

var (
	EndOfBrace = false
)

// Parser is the parser for the Ecla language

type Parser struct {
	ErrorChannel chan errorHandler.Error
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

// Back moves the parser back one token
func (p *Parser) Back() {
	p.TokenIndex--
	if p.TokenIndex < 0 {
		p.TokenIndex = 0
	}
	p.CurrentToken = p.Tokens[p.TokenIndex]
}

// MultiStep moves the parser forward n times givens as parameter
func (p *Parser) MultiStep(steps int) {
	p.TokenIndex += steps
	if p.TokenIndex >= len(p.Tokens) {
		p.CurrentToken = lexer.Token{}
	} else {
		p.CurrentToken = p.Tokens[p.TokenIndex]
	}
}

// MultiBack moves the parser back n times given as parameter
func (p *Parser) MultiBack(steps int) {
	p.TokenIndex -= steps
	if p.TokenIndex < 0 {
		p.TokenIndex = 0
	}
	p.CurrentToken = p.Tokens[p.TokenIndex]
}

// Peek returns the token n steps ahead of the current token without moving the parser
func (p *Parser) Peek(lookAhead int) lexer.Token {
	if p.TokenIndex+lookAhead >= len(p.Tokens) {
		return lexer.Token{}
	}
	return p.Tokens[p.TokenIndex+lookAhead]
}

func (p *Parser) PrintBacktrace() {
	// print back the 10 last token values
	p.MultiBack(10)
	for i := 0; i < 10; i++ {
		fmt.Print(p.CurrentToken.Value)
		p.Step()
	}
	fmt.Println()
}

func (p *Parser) DisableEOLChecking() {
	EndOfBrace = true
	if p.CurrentToken.TokenType != lexer.EOL {
		p.Back()
	} else {
		log.Println("Warning: Superfluous semicolon at line ", p.CurrentToken.Line, ":", p.CurrentToken.Position)
	}
}

/*
Parse is the main function of the parser.
It parses the tokens within itself and returns a File struct containing the AST and the parsed declarations of variables and functions
It also runs the dependency checker to find any missing dependencies and notifies the user
*/
func (p *Parser) Parse() *File {
	p.TokenIndex = -1
	p.Step()
	tempFile := new(File)
	p.Tokens = tempFile.ConsumeComments(p.Tokens)
	file := p.ParseFile()
	file.ConsumedComments = tempFile.ConsumedComments
	ok, UnresolvedDep := file.DepChecker()
	if !ok {
		Unresolved := ""
		for i, dep := range UnresolvedDep {
			if i < len(UnresolvedDep)-1 {
				Unresolved += dep + ", "
			} else {
				Unresolved += dep
			}
		}
		panic("Unresolved dependencies: " + fmt.Sprintf("%v", Unresolved))
	}
	return file
}

// ParseFile parses the tokens and returns a File struct containing the AST and the parsed declarations of variables and functions
func (p *Parser) ParseFile() *File {
	tempFile := new(File)
	tempFile.ParseTree = new(AST)
	p.CurrentFile = tempFile
	for p.CurrentToken.TokenType != lexer.EOF {
		NewNode := p.ParseNode()
		if NewNode != nil {
			tempFile.ParseTree.Operations = append(tempFile.ParseTree.Operations, NewNode)
		}

		p.Step()
	}
	p.Step()
	return tempFile

}

// ParseNode parses a node from the current token deciding checking if the token is text or other
func (p *Parser) ParseNode() Node {
	if p.CurrentToken.TokenType != lexer.TEXT {
		tempExpr := p.ParseExpr()
		return tempExpr
	} else {
		if p.CurrentToken.Value == "\n" || p.CurrentToken.Value == "\r" {
			log.Println("Warning: nil node added to AST")
			return nil
		} else {
			tempExpr := p.ParseText()
			if p.CurrentToken.TokenType != lexer.EOL && !EndOfBrace {
				p.PrintBacktrace()
				panic("Expected EOL" + p.CurrentToken.Value + " " + p.CurrentToken.TokenType)
			}
			if EndOfBrace {
				EndOfBrace = false
			}
			return tempExpr
		}
	}
}

// ParseBody parses a body of a function,a loop or a conditional statement
func (p *Parser) ParseBody() []Node {
	tempBody := make([]Node, 0)
	for p.CurrentToken.TokenType != lexer.RBRACE {
		tempBody = append(tempBody, p.ParseNode())
		p.Step()
	}
	return tempBody
}

// ParseText parses a text node and checking if it is a keyword from Keywords
func (p *Parser) ParseText() Node {
	if _, ok := Keywords[p.CurrentToken.Value]; ok {
		return p.ParseKeyword()
	} else {
		return p.ParseIdent()
	}
}

// ParseKeyword parses a keyword and calls the appropriate parsing function
func (p *Parser) ParseKeyword() Node {
	if p.CurrentToken.Value == "var" {
		return p.ParseVariableDecl()
	}
	if p.CurrentToken.Value == "function" {
		return p.ParseFunctionDecl()
	}
	if p.CurrentToken.Value == "return" {
		return p.ParseReturnStmt()
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
	if p.CurrentToken.Value == "import" {
		return p.ParseImportStmt()
	}
	panic("Unexpected Keyword " + p.CurrentToken.Value)
	return nil
}

// ParseIdent parses an identifier and checking if it is function or method call,a variable declaration or an indexable variable access
func (p *Parser) ParseIdent() Node {
	if p.Peek(1).TokenType == lexer.PERIOD {
		return p.ParseMethodCallExpr()
	} else if p.Peek(1).TokenType == lexer.LPAREN {
		return p.ParseFunctionCallExpr()
	} else {
		return p.ParseVariableAssign()
	}
}

// ParsePrintStmt 'Deprecated function' parses a print statement
func (p *Parser) ParsePrintStmt() Stmt {
	tempPrint := PrintStmt{PrintToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected Print LPAREN")
	}
	tempPrint.Lparen = p.CurrentToken
	p.Step()
	tempPrint.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected Print RPAREN")
	}
	tempPrint.Rparen = p.CurrentToken
	p.Step()
	return tempPrint
}

// ParseTypeStmt parses a type statement
func (p *Parser) ParseTypeStmt() Stmt {
	tempType := TypeStmt{TypeToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected Type LPAREN")
	}
	tempType.Lparen = p.CurrentToken
	p.Step()
	tempType.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected Type RPAREN")
	}
	tempType.Rparen = p.CurrentToken
	p.Step()
	return tempType
}

// ParseIfStmt parses an if statement
func (p *Parser) ParseIfStmt() Stmt {
	tempIf := IfStmt{IfToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected If LPAREN")
	}
	tempIf.LeftParen = p.CurrentToken
	p.Step()
	tempIf.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected If RPAREN")
	}
	tempIf.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		panic("Expected If LBRACE")
	}
	tempIf.LeftBrace = p.CurrentToken
	p.Step()
	tempIf.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected If RBRACE")
	}
	tempIf.RightBrace = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if p.CurrentToken.Value == "else" {
			tempIf.ElseStmt = p.ParseElseStmt()
		} else {
			tempIf.ElseStmt = nil
			p.DisableEOLChecking()
		}
	} else {
		tempIf.ElseStmt = nil
		p.DisableEOLChecking()
	}
	return tempIf
}

// ParseElseStmt parses an else statement
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
		panic("Expected Else LBRACE")
	}
	tempElse.LeftBrace = p.CurrentToken
	p.Step()
	tempElse.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected Else RBRACE")
	}
	tempElse.RightBrace = p.CurrentToken
	p.Step()
	if tempElse.IfStmt == nil {
		p.DisableEOLChecking()
	}
	return tempElse
}

// ParseWhileStmt parses a while statement
func (p *Parser) ParseWhileStmt() Stmt {
	tempWhile := WhileStmt{WhileToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected While LPAREN")
	}
	tempWhile.LeftParen = p.CurrentToken
	p.Step()
	tempWhile.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected While RPAREN")
	}
	tempWhile.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		panic("Expected While LBRACE")
	}
	tempWhile.LeftBrace = p.CurrentToken
	p.Step()
	tempWhile.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected While RBRACE")
	}
	tempWhile.RightBrace = p.CurrentToken
	p.Step()
	p.DisableEOLChecking()
	return tempWhile
}

// ParseForStmt parses a for statement
func (p *Parser) ParseForStmt() Stmt {
	tempFor := ForStmt{}
	tempFor.ForToken = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected For LPAREN")
	}
	tempFor.LeftParen = p.CurrentToken
	p.Step()
	lookAhead := p.Peek(1)
	if lookAhead.TokenType != lexer.COMMA {
		tempFor.RangeToken = lexer.Token{}
		tempFor.InitDecl = p.ParseVariableDecl()
		if p.CurrentToken.TokenType != lexer.COMMA {
			panic("Expected Condition Expression instead of " + p.CurrentToken.Value)
		}
		p.Step()
		tempFor.CondExpr = p.ParseExpr()

		if p.CurrentToken.TokenType != lexer.COMMA {
			panic("Expected Post Expression")
		}
		p.Step()
		tempFor.PostAssignStmt = p.ParseVariableAssign()
	} else {
		tempFor.KeyToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.COMMA {
			panic("Expected For COMMA")
		}
		p.Step()
		tempFor.ValueToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.TEXT {
			panic("Unexpected Token : " + p.CurrentToken.Value)
		}
		if p.CurrentToken.Value != "range" {
			panic("Expected range")
		}
		tempFor.RangeToken = p.CurrentToken
		p.Step()
		tempFor.RangeExpr = p.ParseExpr()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected For RPAREN instead of " + p.CurrentToken.Value)
	}
	tempFor.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		panic("Expected For LBRACE instead of " + p.CurrentToken.Value)
	}
	tempFor.LeftBrace = p.CurrentToken
	p.Step()
	tempFor.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected For RBRACE instead of " + p.CurrentToken.Value)
	}
	tempFor.RightBrace = p.CurrentToken
	p.Step()
	p.DisableEOLChecking()
	return tempFor
}

// ParseVariableDecl parses a variable declaration
func (p *Parser) ParseVariableDecl() Decl {
	tempDecl := VariableDecl{VarToken: p.CurrentToken}
	p.Step()
	if _, ok := Keywords[p.CurrentToken.Value]; ok {
		panic("Variable name cannot be a keyword")
	}
	if p.CurrentToken.TokenType != lexer.TEXT {
		panic("Expected variable name instead of " + p.CurrentToken.Value)
	}
	tempDecl.Name = p.CurrentToken.Value
	typeName, success := p.ParseType()
	if !success {
		panic("Expected variable type instead of " + p.CurrentToken.Value)
	}
	tempDecl.Type = typeName
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		if p.CurrentToken.TokenType != lexer.COMMA {
			if p.CurrentToken.TokenType != lexer.EOL {
				if p.CurrentToken.TokenType != lexer.EOF {
					panic("Expected variable assignment instead of " + p.CurrentToken.Value)
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

// ParseMethodCallExpr parses a method call expression
func (p *Parser) ParseMethodCallExpr() Expr {
	tempMethodCall := MethodCallExpr{MethodCallToken: p.CurrentToken, ObjectName: p.CurrentToken.Value}
	p.CurrentFile.AddDependency(p.CurrentToken.Value)
	p.MultiStep(2)
	tempFunctionCall := p.ParseFunctionCallExpr()
	tempMethodCall.FunctionCall = tempFunctionCall.(FunctionCallExpr)
	return tempMethodCall
}

// ParseFunctionCallExpr parse a function call expression
func (p *Parser) ParseFunctionCallExpr() Expr {
	tempFunctionCall := FunctionCallExpr{FunctionCallToken: p.CurrentToken, Name: p.CurrentToken.Value}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected Function call LPAREN" + p.CurrentToken.Value)
	}
	tempFunctionCall.LeftParen = p.CurrentToken
	var exprArray []Expr
	if p.Peek(1).TokenType != lexer.RPAREN {
		for p.CurrentToken.TokenType != lexer.RPAREN {
			p.Step()
			tempExpr := p.ParseExpr()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.PrintBacktrace()
				panic("Expected comma between function call arguments" + p.CurrentToken.Value)
			}
			exprArray = append(exprArray, tempExpr)
		}
	} else {
		p.Step()
	}

	tempFunctionCall.Args = exprArray
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected Function call RPAREN")
	}
	tempFunctionCall.RightParen = p.CurrentToken
	p.Step()
	return tempFunctionCall
}

// ParseType parses a valid type
func (p *Parser) ParseType() (string, bool) {
	p.Step()
	if _, ok := VarTypes[p.CurrentToken.Value]; ok {
		tempType := ""
		switch p.CurrentToken.Value {
		case "[":
			tempType = p.ParseArrayType()
		case "map":
			tempType = p.ParseMapType()
		default:
			tempType = p.CurrentToken.Value
		}
		p.Step()
		if tempType == "" {
			return "", false
		}
		return tempType, true
	}
	return "", false
}

// ParseArrayType parses an array type
func (p *Parser) ParseArrayType() string {
	tempType := ""
	for p.CurrentToken.TokenType == lexer.LBRACKET {
		tempType += p.CurrentToken.Value
		p.Step()
		if p.CurrentToken.TokenType != lexer.RBRACKET {
			return ""
		}
		tempType += p.CurrentToken.Value
		p.Step()
	}
	p.Back()
	arrayType, success := p.ParseType()
	if !success {
		return ""
	}
	p.Back()
	tempType += arrayType
	return tempType
}

// ParseMapType parses a map type
func (p *Parser) ParseMapType() string {
	tempType := "map"
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACKET {
		return ""
	}
	tempType += p.CurrentToken.Value
	keyType, success := p.ParseType()
	if !success {
		return ""
	}
	tempType += keyType
	if p.CurrentToken.TokenType != lexer.RBRACKET {
		return ""
	}
	tempType += p.CurrentToken.Value
	valueType, success := p.ParseType()
	if !success {
		return ""
	}
	p.Back()
	tempType += valueType
	return tempType
}

// ParseVariableAssign parses a variable assignment
func (p *Parser) ParseVariableAssign() Stmt {

	Var := p.CurrentToken
	Opp := ""
	toAssign := p.ParseVariableAssignLHS()
	p.Back()
	if _, ok := AssignOperators[p.CurrentToken.Value]; ok {
		Opp = p.CurrentToken.Value
	} else {
		panic("Unknown opperator\" " + p.CurrentToken.Value + " in variable assignement")
	}
	p.Step()
	if p.CurrentToken.TokenType == lexer.EOL || p.CurrentToken.TokenType == lexer.EOF {
		return VariableAssignStmt{
			VarToken: Var,
			Names:    toAssign,
			Operator: Opp,
			Values:   []Expr{nil},
		}
	}
	rhs := p.ParseVariableAssignRHS()
	p.Back()
	return VariableAssignStmt{
		VarToken: Var,
		Names:    toAssign,
		Operator: Opp,
		Values:   rhs,
	}
}

func (p *Parser) ParseVariableAssignLHS() []Expr {
	var tempArray []Expr
	tempArray = append(tempArray, p.ParseExpr())
	for p.CurrentToken.TokenType == lexer.COMMA {
		p.Step()
		tempArray = append(tempArray, p.ParseExpr())
	}
	p.Step()
	return tempArray
}

func (p *Parser) ParseVariableAssignRHS() []Expr {
	entered := false
	var tempArray []Expr
	tempArray = append(tempArray, p.ParseExpr())
	for p.CurrentToken.TokenType == lexer.COMMA {
		entered = true
		p.Step()
		tempArray = append(tempArray, p.ParseExpr())
		p.Step()
	}
	if !entered {
		p.Step()
	}
	return tempArray
}

// ParseExpr parse an expression
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
		opPrecedence := TokenPrecedence(p.CurrentToken)
		operator := p.CurrentToken
		if opPrecedence < precedence {
			return Lhs
		}
		p.Step()
		Rhs := p.ParseBinaryExpr(nil, opPrecedence+1)

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

// ParseOperand parses an operand
func (p *Parser) ParseOperand() Expr {
	if p.CurrentToken.TokenType == lexer.LPAREN {
		return p.ParseParenExpr()
	} else if p.CurrentToken.TokenType == lexer.TEXT {
		lookAhead := p.Peek(1)
		if lookAhead.TokenType == lexer.PERIOD {
			return p.ParseMethodCallExpr()
		} else if lookAhead.TokenType == lexer.LPAREN {
			return p.ParseFunctionCallExpr()
		}
	}
	return p.ParseLiteral()
}

// ParseParenExpr parses a parenthesized expression
func (p *Parser) ParseParenExpr() Expr {
	tempParentExpr := ParenExpr{}
	tempParentExpr.Lparen = p.CurrentToken
	p.Step()
	tempParentExpr.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected ')'")
	}
	tempParentExpr.Rparen = p.CurrentToken
	p.Step()
	return tempParentExpr
}

// ParseArrayLiteral parses an array literal
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
		panic("Expected ']'")
	}
	tempArrayExpr.RBRACKET = p.CurrentToken
	p.Step()
	return tempArrayExpr
}

// ParseMapLiteral parses a map literal
func (p *Parser) ParseMapLiteral() Expr {
	tempMapLiteral := MapLiteral{}
	tempMapLiteral.LBRACE = p.CurrentToken
	p.Step()
	for {
		tempMapLiteral.Keys = append(tempMapLiteral.Keys, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COLON {
			panic("Expected ':'")
		}
		p.Step()
		tempMapLiteral.Values = append(tempMapLiteral.Values, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COMMA {
			break
		}
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected '}'")
	}
	tempMapLiteral.RBRACE = p.CurrentToken
	p.Step()
	return tempMapLiteral
}

// ParseImportStmt parses an import statement
func (p *Parser) ParseImportStmt() Stmt {
	tempImportStmt := ImportStmt{}
	tempImportStmt.ImportToken = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.TEXT {
		panic("Expected import path")
	}
	tempImportStmt.ModulePath = p.CurrentToken.Value
	p.CurrentFile.Imports = append(p.CurrentFile.Imports, tempImportStmt.ModulePath)
	p.Step()
	return tempImportStmt
}

// ParseFunctionDecl parses a function declaration
func (p *Parser) ParseFunctionDecl() Node {
	tempFunctionDecl := FunctionDecl{FunctionToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.TEXT {
		panic("Expected function name")
	}
	tempFunctionDecl.Name = p.CurrentToken.Value
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		panic("Expected '('")
	}
	tempFunctionDecl.LeftParamParen = p.CurrentToken
	isParen := p.Peek(1)
	if isParen.TokenType != lexer.RPAREN {
		for p.CurrentToken.TokenType != lexer.RPAREN {
			p.Step()
			// parameter in the form of "a : int, b : int"
			ParamName := ""
			ParamType := ""
			ParamName = p.CurrentToken.Value
			p.Step()
			if p.CurrentToken.TokenType != lexer.COLON {
				panic("Expected ':'")
			}
			ParamType, succes := p.ParseType()
			if !succes {
				panic("Type does not exist")
			}
			p.Back()
			if !(DuplicateParam(tempFunctionDecl.Parameters, ParamName)) {
				newParams := FunctionParams{Name: ParamName, Type: ParamType}
				tempFunctionDecl.Parameters = append(tempFunctionDecl.Parameters, newParams)
			} else {
				panic("Duplicate argument name")
			}
			p.Step()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				panic("Expected ','")
			}
		}
	} else {
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		panic("Expected ')'")
	}
	tempFunctionDecl.RightParamParen = p.CurrentToken
	if p.Peek(1).TokenType == lexer.LPAREN {
		p.Step()
		tempFunctionDecl.LeftRetsParen = p.CurrentToken
		for p.CurrentToken.TokenType != lexer.RPAREN {
			retType, success := p.ParseType()
			if !success {
				panic("Type does not exist")
			}
			tempFunctionDecl.ReturnTypes = append(tempFunctionDecl.ReturnTypes, retType)
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				panic("Expected ','" + p.CurrentToken.Value + " " + p.CurrentToken.TokenType)
			}
		}
		if p.CurrentToken.TokenType != lexer.RPAREN {
			panic("Expected ')'")
		}
	}
	p.Step()
	tempFunctionDecl.RightRetsParen = p.CurrentToken
	if p.CurrentToken.TokenType != lexer.LBRACE {
		panic("Expected '{' instead of " + p.CurrentToken.Value)
	}
	p.Step()
	tempFunctionDecl.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		panic("Expected '}'" + p.CurrentToken.Value)
	}
	p.Step()
	p.CurrentFile.FunctionDecl = append(p.CurrentFile.FunctionDecl, tempFunctionDecl.Name)
	p.DisableEOLChecking()
	return tempFunctionDecl
}

func DuplicateParam(params []FunctionParams, newParam string) bool {
	for _, paramName := range params {
		if paramName.Name == newParam {
			return true
		}
	}
	return false
}

// ParseReturnStmt parses a return statement
func (p *Parser) ParseReturnStmt() Node {
	tempReturnStmt := ReturnStmt{ReturnToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType == lexer.EOL {
		return tempReturnStmt
	}
	err := false
	for {
		tempReturnStmt.ReturnValues = append(tempReturnStmt.ReturnValues, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COMMA {
			if p.CurrentToken.TokenType != lexer.EOL && p.CurrentToken.TokenType != lexer.EOF {
				err = true
			}
			break
		}
		p.Step()
	}
	if err {
		panic("Unexpected token in return statement" + p.CurrentToken.Value + " " + p.CurrentToken.TokenType)
	}
	return tempReturnStmt
}

// ParseIndexableAccessExpr parses an indexable variable access expression
func (p *Parser) ParseIndexableAccessExpr() Expr {
	tempIndexableAccessExpr := IndexableAccessExpr{VariableToken: p.CurrentToken, VariableName: p.CurrentToken.Value}
	p.Step()
	for p.CurrentToken.TokenType == lexer.LBRACKET {
		p.Step()
		tempIndexableAccessExpr.Indexes = append(tempIndexableAccessExpr.Indexes, p.ParseExpr())
		p.Step()
	}
	return tempIndexableAccessExpr
}

func (p *Parser) ParseVariableAccess() Expr {
	if p.Peek(1).TokenType == lexer.LBRACKET {
		temp := p.ParseIndexableAccessExpr()
		p.Back()
		return temp
	} else {
		return Literal{Token: p.CurrentToken, Type: "VAR", Value: p.CurrentToken.Value}
	}
}

// ParseLiteral parses a literal
func (p *Parser) ParseLiteral() Expr {
	if p.CurrentToken.TokenType == lexer.INT || p.CurrentToken.TokenType == lexer.FLOAT {
		tempLiteral := Literal{Token: p.CurrentToken, Type: p.CurrentToken.TokenType, Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.TEXT {
		if p.CurrentToken.Value == "null" {
			tempLiteral := Literal{Token: p.CurrentToken, Type: "NULL", Value: p.CurrentToken.Value}
			p.Step()
			return tempLiteral
		}
		tempLiteral := p.ParseVariableAccess()
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.DQUOTE {
		p.Step()
		tempLiteral := Literal{}
		if p.CurrentToken.TokenType == lexer.DQUOTE {
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: ""}
			p.Step()
			return tempLiteral
		} else {
			if p.CurrentToken.TokenType != lexer.STRING {
				p.PrintBacktrace()
				panic("Expected string")
			}
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: p.CurrentToken.Value}
		}

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
	if p.CurrentToken.TokenType == lexer.LBRACKET {
		return p.ParseArrayLiteral()
	}
	if p.CurrentToken.TokenType == lexer.LBRACE {
		return p.ParseMapLiteral()
	}
	p.PrintBacktrace()
	panic("Expected literal instead of " + p.CurrentToken.Value)
	return nil
}
