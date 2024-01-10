package parser

import (
	"fmt"

	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/lexer"
)

// Parser is the parser for the Ecla language

type Parser struct {
	ErrorHandler *errorHandler.ErrorHandler
	Tokens       []lexer.Token
	TokenIndex   int
	CurrentToken lexer.Token
	CurrentFile  *File
	IsEndOfBrace bool
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

// PrintBacktrace prints the last 10 tokens for debugging purposes
func (p *Parser) PrintBacktrace() {
	// print back the 10 last token values
	p.MultiBack(10)
	for i := 0; i < 10; i++ {
		fmt.Print(p.CurrentToken.Value)
		p.Step()
	}
	fmt.Println()
}

// HandleWarning handles a warning level error
func (p *Parser) HandleWarning(message string) {
	p.ErrorHandler.HandleError(p.CurrentToken.Line, p.CurrentToken.Position, message, errorHandler.LevelWarning)
}

// HandleError handles an error level error
func (p *Parser) HandleError(message string) {
	p.ErrorHandler.HandleError(p.CurrentToken.Line, p.CurrentToken.Position, message, errorHandler.LevelError)
}

// HandleFatal handles a fatal level error
func (p *Parser) HandleFatal(message string) {
	p.ErrorHandler.HandleError(p.CurrentToken.Line, p.CurrentToken.Position, message, errorHandler.LevelFatal)
	fmt.Println("If you see this message, that the error handler is not working properly :(, or you disabled the panic on fatal errors")
}

// DisableEOLChecking disables the EOL checking for the parser to allow for semicolon-less syntax on conditionals statements and loops
func (p *Parser) DisableEOLChecking() {
	p.IsEndOfBrace = true
	if p.CurrentToken.TokenType != lexer.EOL {
		p.Back()
	} else {
		p.HandleWarning(fmt.Sprintf("Superfluous semicolon at line %v:%v", p.CurrentToken.Line, p.CurrentToken.Position))
	}
}

/*
Parse is the main function of the parser.
It parses the tokens within itself and returns a File struct containing the AST and the parsed declarations of variables and functions
It also runs the dependency checker to find any missing dependencies and notifies the user
*/
func (p *Parser) Parse() *File {
	tempFile := new(File)
	p.Tokens = tempFile.ConsumeComments(p.Tokens)
	p.CurrentToken = p.Tokens[0]
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
		p.HandleFatal("Unresolved dependencies: " + fmt.Sprintf("%v", Unresolved))
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
	if p.CurrentToken.TokenType == lexer.MURLOC {
		tempStmt := MurlocStmt{MurlocToken: p.CurrentToken}
		p.Step()
		return tempStmt
	}
	if p.CurrentToken.TokenType != lexer.TEXT {
		tempExpr := p.ParseExpr()
		return tempExpr
	} else {
		if p.CurrentToken.Value == "\n" || p.CurrentToken.Value == "\r" {
			p.HandleWarning("nil node added to AST")
			return nil
		} else {
			tempExpr := p.ParseText()
			if p.CurrentToken.TokenType != lexer.EOL && !p.IsEndOfBrace {
				p.PrintBacktrace()
				p.HandleFatal("Expected End Of Line")
			}
			if p.IsEndOfBrace {
				p.IsEndOfBrace = false
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
	if p.CurrentToken.Value == Var {
		return p.ParseVariableDecl()
	}
	if p.CurrentToken.Value == Function {
		if p.Peek(1).TokenType == lexer.LPAREN {
			return p.ParseAnonymousFunctionExpr()
		}
		return p.ParseFunctionDecl()
	}
	if p.CurrentToken.Value == Return {
		return p.ParseReturnStmt()
	}
	if p.CurrentToken.Value == Type {
		return p.ParseTypeStmt()
	}
	if p.CurrentToken.Value == If {
		return p.ParseIfStmt()
	}
	if p.CurrentToken.Value == While {
		return p.ParseWhileStmt()
	}
	if p.CurrentToken.Value == For {
		return p.ParseForStmt()
	}
	if p.CurrentToken.Value == Import {
		return p.ParseImportStmt()
	}
	if p.CurrentToken.Value == Null {
		tempLiteral := Literal{Token: p.CurrentToken, Type: "NULL", Value: p.CurrentToken.Value}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.Value == Murloc {
		tempStmt := MurlocStmt{MurlocToken: p.CurrentToken}
		p.Step()
		return tempStmt
	}
	if p.CurrentToken.Value == Any {
		p.HandleFatal("any cannot be used as a keyword")
		return nil
	}
	p.HandleFatal("Unknown keyword: " + p.CurrentToken.Value)
	return nil
}

// ParseIdent parses an identifier and checking if it is function or method call,a variable declaration or an indexable variable access
func (p *Parser) ParseIdent() Node {
	if p.Peek(1).TokenType == lexer.PERIOD {
		return p.ParseMethodCallExpr()
	} else if p.Peek(1).TokenType == lexer.LPAREN {
		return p.ParseFunctionCallExpr()
	} else if p.Peek(1).TokenType == lexer.COLON {
		p.Back()
		return p.ParseImplicitVariableDecl()
	} else {
		return p.ParseVariableAssign()
	}
}

// ParseTypeStmt parses a type statement
func (p *Parser) ParseTypeStmt() Stmt {
	tempType := TypeStmt{TypeToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.HandleFatal("Expected Type LPAREN")
	}
	tempType.Lparen = p.CurrentToken
	p.Step()
	tempType.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected Type RPAREN")
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
		p.HandleFatal("Expected If LPAREN")
	}
	tempIf.LeftParen = p.CurrentToken
	p.Step()
	tempIf.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected If RPAREN")
	}
	tempIf.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected If LBRACE")
	}
	tempIf.LeftBrace = p.CurrentToken
	p.Step()
	tempIf.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected If RBRACE")
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
		p.HandleFatal("Expected Else LBRACE")
	}
	tempElse.LeftBrace = p.CurrentToken
	p.Step()
	tempElse.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected Else RBRACE")
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
		p.HandleFatal("Expected While LPAREN")
	}
	tempWhile.LeftParen = p.CurrentToken
	p.Step()
	tempWhile.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected While RPAREN")
	}
	tempWhile.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected While LBRACE")
	}
	tempWhile.LeftBrace = p.CurrentToken
	p.Step()
	tempWhile.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected While RBRACE")
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
		p.HandleFatal("Expected For LPAREN")
	}
	tempFor.LeftParen = p.CurrentToken
	p.Step()
	lookAhead := p.Peek(1)
	if lookAhead.TokenType != lexer.COMMA {
		tempFor.RangeToken = lexer.Token{}
		if lookAhead.TokenType == lexer.COLON {
			p.Back()
			tempFor.InitDecl = p.ParseImplicitVariableDecl()
		} else {
			tempFor.InitDecl = p.ParseVariableDecl()
		}
		if p.CurrentToken.TokenType != lexer.COMMA {
			p.HandleFatal("Expected Condition Expression instead of " + p.CurrentToken.Value)
		}
		p.Step()
		tempFor.CondExpr = p.ParseExpr()

		if p.CurrentToken.TokenType != lexer.COMMA {
			p.HandleFatal("Expected Post Expression")
		}
		p.Step()
		tempFor.PostAssignStmt = p.ParseVariableAssign()
	} else {
		tempFor.KeyToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.COMMA {
			p.HandleFatal("Expected For COMMA")
		}
		p.Step()
		tempFor.ValueToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.TEXT {
			p.HandleFatal("Unexpected token " + p.CurrentToken.Value)
		}
		if p.CurrentToken.Value != "range" {
			p.HandleFatal("Expected For range instead of " + p.CurrentToken.Value)
		}
		tempFor.RangeToken = p.CurrentToken
		p.Step()
		tempFor.RangeExpr = p.ParseExpr()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected For RPAREN")
	}
	tempFor.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected For LBRACE")
	}
	tempFor.LeftBrace = p.CurrentToken
	p.Step()
	tempFor.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected For RBRACE")
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
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as variable name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
		}
	} else {
		p.HandleFatal("Expected variable name instead of " + p.CurrentToken.Value)
	}
	tempDecl.Name = p.CurrentToken.Value
	typeName, success := p.ParseType()
	if !success {
		p.HandleFatal("Expected variable type instead of " + p.CurrentToken.Value)
	}
	tempDecl.Type = typeName
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.EOL && p.CurrentToken.TokenType != lexer.EOF {
			p.HandleFatal("Expected variable assignment instead of " + p.CurrentToken.Value)
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

func (p *Parser) ParseImplicitVariableDecl() Decl {
	tempDecl := VariableDecl{VarToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as variable name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
		}
	} else {
		p.HandleFatal("Expected variable name instead of " + p.CurrentToken.Value)
	}
	tempDecl.Name = p.CurrentToken.Value
	p.MultiStep(2)
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		p.HandleFatal("Expected implicit variable assignment instead of " + p.CurrentToken.Value)
	}
	p.Step()
	tempDecl.Value = p.ParseExpr()
	p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
	return tempDecl
}

// ParseMethodCallExpr parses a method call expression
func (p *Parser) ParseMethodCallExpr() Expr {
	tempMethodCall := MethodCallExpr{MethodCallToken: p.CurrentToken, ObjectName: p.CurrentToken.Value}
	err := p.CurrentFile.AddDependency(p.CurrentToken.Value)
	if err != nil {
		p.HandleFatal(err.Error())
	}
	p.MultiStep(2)
	tempFunctionCall := p.ParseFunctionCallExpr()
	tempMethodCall.FunctionCall = tempFunctionCall.(FunctionCallExpr)
	return tempMethodCall
}

// ParseFunctionCallExpr parse a function call expression
func (p *Parser) ParseFunctionCallExpr() Expr {
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as function name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
		}
	} else {
		p.HandleFatal("Expected function name instead of " + p.CurrentToken.Value)
	}
	tempFunctionCall := FunctionCallExpr{FunctionCallToken: p.CurrentToken, Name: p.CurrentToken.Value}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.HandleFatal("Expected Function call LPAREN")
	}
	tempFunctionCall.LeftParen = p.CurrentToken
	var exprArray []Expr
	if p.Peek(1).TokenType != lexer.RPAREN {
		for p.CurrentToken.TokenType != lexer.RPAREN {
			p.Step()
			tempExpr := p.ParseExpr()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.PrintBacktrace()
				p.HandleFatal("Expected comma between function call arguments")
			}
			exprArray = append(exprArray, tempExpr)
		}
	} else {
		p.Step()
	}

	tempFunctionCall.Args = exprArray
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected Function call RPAREN")
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
		case ArrayStart:
			tempType = p.ParseArrayType()
		case Map:
			tempType = p.ParseMapType()
		case Function:
			tempType = p.ParseFunctionType()
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
	tempType := Map
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

func (p *Parser) ParseFunctionType() string {
	tempType := Function
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		return ""
	}
	tempType += p.CurrentToken.Value
	// parse arguments types using the parseType function
	for p.CurrentToken.TokenType != lexer.RPAREN {
		temp, success := p.ParseType()
		if !success {
			return ""
		}
		tempType += temp
		if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
			return ""
		}
		tempType += p.CurrentToken.Value
	}
	p.Step()
	// check if the function has a return type
	if p.CurrentToken.TokenType == lexer.LPAREN {
		tempType += p.CurrentToken.Value
		// parse return type using the parseType function
		for p.CurrentToken.TokenType != lexer.RPAREN {
			temp, success := p.ParseType()
			if !success {
				return ""
			}
			tempType += temp
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				return ""
			}
		}
		tempType += p.CurrentToken.Value
	}
	return tempType
}

// ParseVariableAssign parses a variable assignment
func (p *Parser) ParseVariableAssign() Stmt {

	Var := p.CurrentToken
	Opp := ""
	toAssign := p.ParseVariableAssignSide()
	if _, ok := AssignOperators[p.CurrentToken.Value]; ok {
		Opp = p.CurrentToken.Value
	} else {
		p.HandleFatal("Unknown opperator\" " + p.CurrentToken.Value + " in variable assignement")
	}
	p.Step()
	if p.CurrentToken.TokenType == lexer.EOL || p.CurrentToken.TokenType == lexer.EOF || p.CurrentToken.TokenType == lexer.RPAREN || p.CurrentToken.TokenType == lexer.RBRACKET || p.CurrentToken.TokenType == lexer.RBRACE {
		return VariableAssignStmt{
			VarToken: Var,
			Names:    toAssign,
			Operator: Opp,
			Values:   []Expr{nil},
		}
	}
	rhs := p.ParseVariableAssignSide()
	return VariableAssignStmt{
		VarToken: Var,
		Names:    toAssign,
		Operator: Opp,
		Values:   rhs,
	}
}

// ParseVariableAssignSide parses the right hand side or the left hand side of a variable assignment
func (p *Parser) ParseVariableAssignSide() []Expr {
	var tempArray []Expr
	tempArray = append(tempArray, p.ParseExpr())
	for p.CurrentToken.TokenType == lexer.COMMA {
		p.Step()
		tempArray = append(tempArray, p.ParseExpr())
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
		// check if it is an anonymous function or an anonymous struct
		if p.CurrentToken.Value == Function {
			return p.ParseAnonymousFunctionExpr()
		} //else if p.CurrentToken.Value == Struct {
		//	return p.ParseAnonymousStructExpr()
		//}
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
		p.HandleFatal("Expected ')'")
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
		p.HandleFatal("Expected ']'")
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
			p.HandleFatal("Expected ':'")
		}
		p.Step()
		tempMapLiteral.Values = append(tempMapLiteral.Values, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COMMA {
			break
		}
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}'")
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
	if p.CurrentToken.TokenType != lexer.DQUOTE {
		p.HandleFatal("Missing string opening double quote")
	}
	p.Step()
	if p.CurrentToken.TokenType != lexer.STRING {
		p.HandleFatal("Expected package name")
	}
	tempImportStmt.ModulePath = p.CurrentToken.Value
	p.CurrentFile.AddImport(p.CurrentToken.Value)
	p.Step()
	if p.CurrentToken.TokenType != lexer.DQUOTE {
		p.HandleFatal("Missing string closing double quote")
	}

	p.Step()
	return tempImportStmt
}

func (p *Parser) ParseAnonymousFunctionExpr() Expr {
	tempAnonymousFunctionDecl := AnonymousFunctionExpr{FunctionToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.PrintBacktrace()
		p.HandleFatal("Expected '('")
	}
	tempAnonymousFunctionDecl.Prototype = p.ParsePrototype()
	p.Step()
	tempAnonymousFunctionDecl.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}'")
	}
	p.Step()
	//check if it is a call
	if p.CurrentToken.TokenType == lexer.LPAREN {
		tempAnonymousFunctionCall := AnonymousFunctionCallExpr{AnonymousFunction: tempAnonymousFunctionDecl, LeftParen: p.CurrentToken}
		var exprArray []Expr
		if p.Peek(1).TokenType != lexer.RPAREN {
			for p.CurrentToken.TokenType != lexer.RPAREN {
				p.Step()
				tempExpr := p.ParseExpr()
				if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
					p.PrintBacktrace()
					p.HandleFatal("Expected comma between Anonymous function call arguments")
				}
				exprArray = append(exprArray, tempExpr)
			}
		} else {
			p.Step()
		}

		tempAnonymousFunctionCall.Args = exprArray
		if p.CurrentToken.TokenType != lexer.RPAREN {
			p.HandleFatal("Expected Anonymous Function call RPAREN")
		}
		tempAnonymousFunctionCall.RightParen = p.CurrentToken
		p.Step()
		return tempAnonymousFunctionCall
	}
	return tempAnonymousFunctionDecl
}

// ParseFunctionDecl parses a function declaration
func (p *Parser) ParseFunctionDecl() Node {
	tempFunctionDecl := FunctionDecl{FunctionToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as function name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
		}
	} else {
		p.HandleFatal("Expected function name instead of " + p.CurrentToken.Value)
	}
	tempFunctionDecl.Name = p.CurrentToken.Value
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.PrintBacktrace()
		p.HandleFatal("Expected '('")
	}
	tempFunctionDecl.Prototype = p.ParsePrototype()
	p.Step()
	tempFunctionDecl.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}'")
	}
	p.Step()
	p.CurrentFile.FunctionDecl = append(p.CurrentFile.FunctionDecl, tempFunctionDecl.Name)
	p.DisableEOLChecking()
	return tempFunctionDecl
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
		p.HandleFatal("Unexpected token in return statement")
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

// ParseVariableAccess parses a variable access
func (p *Parser) ParseVariableAccess() Expr {
	// check if the variable name is not in the keywords and a text
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as variable name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
		}
	}
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
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			if p.CurrentToken.Value == Null {
				tempLiteral := Literal{Token: p.CurrentToken, Type: "NULL", Value: p.CurrentToken.Value}
				p.Step()
				return tempLiteral
			}
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as variable name")
		}
		if _, ok := VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
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
				p.HandleFatal("Expected string")
			}
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: p.CurrentToken.Value}
		}

		p.Step()
		if p.CurrentToken.TokenType != lexer.DQUOTE {
			p.HandleFatal("Expected \"")
		}
		p.Step()
		return tempLiteral
	}
	if p.CurrentToken.TokenType == lexer.SQUOTE {
		p.Step()
		tempLiteral := Literal{}
		if p.CurrentToken.TokenType == lexer.SQUOTE {
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.CHAR, Value: ""}
			p.Step()
			return tempLiteral
		} else {
			if p.CurrentToken.TokenType != lexer.CHAR {
				p.PrintBacktrace()
				p.HandleFatal("Expected char")
			}
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.CHAR, Value: p.CurrentToken.Value}
		}

		p.Step()
		if p.CurrentToken.TokenType != lexer.SQUOTE {
			p.HandleFatal("Expected '")
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
	p.HandleFatal("Expected literal")
	return nil
}

// DuplicateParam checks if a parameter is already in the list of function parameters
func DuplicateParam(params []FunctionParams, newParam string) bool {
	for _, paramName := range params {
		if paramName.Name == newParam {
			return true
		}
	}
	return false
}

func (p *Parser) ParsePrototype() FunctionPrototype {
	tempFunctionPrototype := FunctionPrototype{}
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.PrintBacktrace()
		p.HandleFatal("Expected '('")
	}
	tempFunctionPrototype.LeftParamParen = p.CurrentToken
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
				p.HandleFatal("Expected ':'")
			}
			ParamType, succes := p.ParseType()
			if !succes {
				p.HandleFatal("Unknown type")
			}
			p.Back()
			if !(DuplicateParam(tempFunctionPrototype.Parameters, ParamName)) {
				newParams := FunctionParams{Name: ParamName, Type: ParamType}
				tempFunctionPrototype.Parameters = append(tempFunctionPrototype.Parameters, newParams)
			} else {
				p.HandleFatal("Duplicate parameter name")
			}
			p.Step()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.HandleFatal("Expected ','")
			}
		}
	} else {
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')'")
	}
	tempFunctionPrototype.RightParamParen = p.CurrentToken
	if p.Peek(1).TokenType == lexer.LPAREN {
		p.Step()
		tempFunctionPrototype.LeftRetsParen = p.CurrentToken
		for p.CurrentToken.TokenType != lexer.RPAREN {
			retType, success := p.ParseType()
			if !success {
				p.HandleFatal("Unknown type")
			}
			tempFunctionPrototype.ReturnTypes = append(tempFunctionPrototype.ReturnTypes, retType)
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.HandleFatal("Expected ','")
			}
		}
		if p.CurrentToken.TokenType != lexer.RPAREN {
			p.HandleFatal("Expected ')'")
		}
	}
	p.Step()
	tempFunctionPrototype.RightRetsParen = p.CurrentToken
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{'")
	}
	return tempFunctionPrototype
}
