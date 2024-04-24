package parser

import (
	"fmt"
	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/lexer"
	"strings"
)

// Parser is the parser for the Ecla language

type Parser struct {
	ErrorHandler *errorHandler.ErrorHandler
	Tokens       []lexer.Token
	TokenIndex   int
	CurrentToken lexer.Token
	CurrentFile  *File
	IsEndOfBrace bool
	VarTypes     map[string]interface{}
}

var selectorDepth int

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
}

// DisableEOLChecking disables the EOL checking for the parser to allow for semicolon-less syntax on conditionals statements and loops
func (p *Parser) DisableEOLChecking() {
	p.IsEndOfBrace = true
	if p.CurrentToken.TokenType != lexer.EOL {
		p.Back()
	} else {
		p.HandleWarning("Semicolon can be omitted")
	}
}

/*
Parse is the main function of the parser.
It parses the tokens within itself and returns a File struct containing the AST and the parsed declarations of variables and functions
It also runs the dependency checker to find any missing dependencies and notifies the user
*/
func (p *Parser) Parse() *File {
	tempFile := new(File)
	// create a deep copy of the VarTypes map
	p.VarTypes = make(map[string]interface{})
	for k, v := range VarTypes {
		p.VarTypes[k] = v
	}
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
		p.HandleFatal("Some dependencies are not satisfied: " + Unresolved)
		return nil
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
	if p.CurrentToken.TokenType == lexer.LBRACE {
		// parse a block
		return p.ParseBlock()
	}
	if p.CurrentToken.TokenType != lexer.TEXT {
		tempExpr := p.ParseExpr()
		return tempExpr
	} else {
		tempExpr := p.ParseText()
		if p.CurrentToken.TokenType != lexer.EOL && !p.IsEndOfBrace {
			p.PrintBacktrace()
			p.HandleFatal("Expected semicolon at the end of the line")
			return nil
		}
		if p.IsEndOfBrace {
			p.IsEndOfBrace = false
		}
		return tempExpr
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
		tempNode := p.ParseIdent()
		// check if there is a period after the identifier to see if it is a selector
		if p.CurrentToken.TokenType == lexer.PERIOD {
			// check if the ident is a Expr
			p.Step()
			if _, ok := tempNode.(Expr); ok {
				selectorDepth++
				exp := p.ParseSelector(tempNode.(Expr))
				selectorDepth--
				// check if exp.Expr is a Literal
				// TODO: remove this line since it is not called and is handled in ParsePrimaryExpr
				if _, ok := exp.(SelectorExpr).Expr.(Literal); ok {
					p.CurrentFile.AddDependency(exp.(SelectorExpr).Expr.(Literal).Token.Value)
				}

				return exp
			} else {
				p.HandleFatal("Cannot use a selector on a non-expression")
				return nil
			}
		}
		return tempNode
	}
}

// ParseBlock parses a block scope
func (p *Parser) ParseBlock() Node {
	tempBlock := BlockScopeStmt{LeftBrace: p.CurrentToken}
	p.Step()
	tempBlock.Body = p.ParseBody()
	// TODO: this check is not needed since ParseBody runs until a right brace
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' at the end of the block")
		return nil
	}
	tempBlock.RightBrace = p.CurrentToken
	return tempBlock
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
		p.HandleFatal("Any cannot be used as a standalone keyword")
		return nil
	}
	if p.CurrentToken.Value == Struct {
		return p.ParseStructDecl()
	}

	p.HandleFatal("Unknown keyword: " + p.CurrentToken.Value)
	return nil
}

// ParseStructDecl parses a struct declaration
func (p *Parser) ParseStructDecl() Node {
	tempStructDecl := StructDecl{StructToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as struct name")
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as struct name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as struct name")
			return nil
		}
	} else {
		p.HandleFatal("Expected struct name instead of " + p.CurrentToken.Value)
		return nil
	}
	tempStructDecl.Name = p.CurrentToken.Value
	// add the struct name to the list of types
	p.VarTypes[tempStructDecl.Name] = nil
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after struct name")
		return nil
	}
	tempStructDecl.LeftBrace = p.CurrentToken
	p.Step()
	for p.CurrentToken.TokenType != lexer.RBRACE {
		tempStructDecl.Fields = append(tempStructDecl.Fields, p.ParseStructField())
		p.Back()
		if p.CurrentToken.TokenType != lexer.EOL && p.CurrentToken.TokenType != lexer.RBRACE {
			p.HandleFatal("Expected semicolon or '}' after struct field")
			return nil
		}
		if p.CurrentToken.TokenType == lexer.EOL {
			p.Step()
		}
	}
	// TODO: remove this condition since the for loop will stop when it reaches the RBRACE
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after struct fields")
		return nil
	}
	p.Step()
	tempStructDecl.RightBrace = p.CurrentToken
	p.DisableEOLChecking()
	return tempStructDecl
}

// ParseStructField parses a struct field
func (p *Parser) ParseStructField() StructField {
	tempStructField := StructField{}
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as struct field name")
			return tempStructField
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as struct field name")
			return tempStructField
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as struct field name")
			return tempStructField
		}
	} else {
		p.HandleFatal("Expected struct field name instead of " + p.CurrentToken.Value)
		return tempStructField
	}
	tempStructField.Name = p.CurrentToken.Value
	p.Step()
	if p.CurrentToken.TokenType != lexer.COLON {
		p.HandleFatal("Expected struct ':' after struct field name")
		return tempStructField
	}
	var success bool
	tempStructField.Type, success = p.ParseType()
	if !success {
		p.HandleFatal("Expected struct field type instead of " + p.CurrentToken.Value)
		return tempStructField
	}
	p.Step()
	return tempStructField
}

// ParseIdent parses an identifier and checking if it is function or method call,a variable declaration or an indexable variable access
func (p *Parser) ParseIdent() Node {
	if p.Peek(1).TokenType == lexer.LPAREN {
		return p.ParseFunctionCallExpr()
	} else if p.Peek(1).TokenType == lexer.LBRACE {
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			if _, ok2 := DefaultVarTypes[p.CurrentToken.Value]; !ok2 {
				return p.ParseStructInstantiation()
			}
		}
	} else if p.Peek(1).TokenType == lexer.PERIOD {
		tempExpr := p.ParseExpr()
		// check if the selector is before a variable assignement
		if _, ok := AssignOperators[p.CurrentToken.Value]; ok {
			return p.ParseVariableAssign(tempExpr)
		} else {
			return tempExpr
		}
	} else if p.Peek(1).TokenType == lexer.COLON {
		p.Back()
		return p.ParseImplicitVariableDecl()
	} else {
		return p.ParseVariableAssign(nil)
	}
	// TODO: refactor this code to remove the last else statement
	p.HandleFatal("Unknown identifier: " + p.CurrentToken.Value)
	return nil
}

// ParseIfStmt parses an if statement
func (p *Parser) ParseIfStmt() Stmt {
	tempIf := IfStmt{IfToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.HandleFatal("Expected '(' after if")
		return nil
	}
	tempIf.LeftParen = p.CurrentToken
	p.Step()
	tempIf.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after if condition")
		return nil
	}
	tempIf.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after if condition")
		return nil
	}
	tempIf.LeftBrace = p.CurrentToken
	p.Step()
	tempIf.Body = p.ParseBody()
	// TODO: this check is not needed since ParseBody runs until a right brace
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after if body")
		return nil
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
		p.HandleFatal("Expected '{' after else")
		return nil
	}
	tempElse.LeftBrace = p.CurrentToken
	p.Step()
	tempElse.Body = p.ParseBody()
	// TODO: this check is not needed since ParseBody runs until a right brace
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after else body")
		return nil
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
		p.HandleFatal("Expected '(' after while")
		return nil
	}
	tempWhile.LeftParen = p.CurrentToken
	p.Step()
	tempWhile.Cond = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after while condition")
		return nil
	}
	tempWhile.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after while condition")
		return nil
	}
	tempWhile.LeftBrace = p.CurrentToken
	p.Step()
	tempWhile.Body = p.ParseBody()
	// TODO: this check is not needed since ParseBody runs until a right brace
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after while body")
		return nil
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
		p.HandleFatal("Expected '(' after for")
		return nil
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
			p.HandleFatal("Expected a condition expression instead of " + p.CurrentToken.Value)
			return nil
		}
		p.Step()
		tempFor.CondExpr = p.ParseExpr()

		if p.CurrentToken.TokenType != lexer.COMMA {
			p.HandleFatal("Expected a post assignement instead of " + p.CurrentToken.Value)
			return nil
		}
		p.Step()
		tempFor.PostAssignStmt = p.ParseVariableAssign(nil)
	} else {
		tempFor.KeyToken = p.CurrentToken
		p.Step()
		//TODO: remove this check since it is checked in with the if statement above
		if p.CurrentToken.TokenType != lexer.COMMA {
			p.HandleFatal("Expected ',' between key and value instead of " + p.CurrentToken.Value)
			return nil
		}
		p.Step()
		tempFor.ValueToken = p.CurrentToken
		p.Step()
		if p.CurrentToken.TokenType != lexer.TEXT {
			p.HandleFatal("Expected 'range' keyword instead of " + p.CurrentToken.Value)
			return nil
		}
		if p.CurrentToken.Value != "range" {
			p.HandleFatal("Expected 'range' keyword instead of " + p.CurrentToken.Value)
			return nil
		}
		tempFor.RangeToken = p.CurrentToken
		p.Step()
		tempFor.RangeExpr = p.ParseExpr()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after for condition")
		return nil
	}
	tempFor.RightParen = p.CurrentToken
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after for condition")
		return nil
	}
	tempFor.LeftBrace = p.CurrentToken
	p.Step()
	tempFor.Body = p.ParseBody()
	//TODO: this check is not needed since ParseBody runs until a right brace
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after for body")
		return nil
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
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as variable name")
			return nil

		}
	} else {
		p.HandleFatal("Expected variable name instead of " + p.CurrentToken.Value)
		return nil
	}
	tempDecl.Name = p.CurrentToken.Value
	typeName, success := p.ParseType()
	if !success {
		p.HandleFatal("Expected variable type instead of " + p.CurrentToken.Value)
		return nil
	}
	tempDecl.Type = typeName
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.EOL && p.CurrentToken.TokenType != lexer.EOF {
			p.HandleFatal("Expected '=' after variable type")
			return nil
		}
		tempDecl.Value = nil
		p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
		return tempDecl
	}
	p.Step()
	tempDecl.Value = p.ParseExpr()

	// check if the Expr is a StructInstantiationExpr
	if _, ok := tempDecl.Value.(StructInstantiationExpr); ok {
		p.CurrentFile.StructInstances = append(p.CurrentFile.StructInstances, tempDecl.Name)
	}

	p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
	return tempDecl
}

func (p *Parser) ParseImplicitVariableDecl() Decl {
	tempDecl := VariableDecl{VarToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as variable name")
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as variable name")
			return nil

		}
	} else {
		p.HandleFatal("Expected variable name instead of " + p.CurrentToken.Value)
		return nil
	}
	tempDecl.Name = p.CurrentToken.Value
	p.MultiStep(2)
	if p.CurrentToken.TokenType != lexer.ASSIGN {
		p.HandleFatal("Expected ':=' after variable name")
		return nil
	}
	p.Step()
	tempDecl.Value = p.ParseExpr()
	// check if the Expr is a StructInstantiationExpr
	if _, ok := tempDecl.Value.(StructInstantiationExpr); ok {
		p.CurrentFile.StructInstances = append(p.CurrentFile.StructInstances, tempDecl.Name)
	}
	p.CurrentFile.VariableDecl = append(p.CurrentFile.VariableDecl, tempDecl.Name)
	return tempDecl
}

// ParseFunctionCallExpr parse a function call expression
func (p *Parser) ParseFunctionCallExpr() Expr {
	if p.CurrentToken.TokenType == lexer.TEXT {
		if _, ok := Keywords[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use keyword " + p.CurrentToken.Value + " as function name")
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			if _, ok2 := DefaultVarTypes[p.CurrentToken.Value]; ok2 {
				p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
				return nil
			}
		}
	} else {
		p.HandleFatal("Expected function name instead of " + p.CurrentToken.Value)
		return nil
	}
	tempFunctionCall := FunctionCallExpr{FunctionCallToken: p.CurrentToken, Name: p.CurrentToken.Value}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.HandleFatal("Expected '(' after function name")
		return nil
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
				return nil
			}
			exprArray = append(exprArray, tempExpr)
		}
	} else {
		p.Step()
	}

	tempFunctionCall.Args = exprArray
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after function call arguments")
		return nil
	}
	tempFunctionCall.RightParen = p.CurrentToken
	p.Step()
	return tempFunctionCall
}

func (p *Parser) ParseStructInstantiation() StructInstantiationExpr {
	tempStructInstantiation := StructInstantiationExpr{StructNameToken: p.CurrentToken, Name: p.CurrentToken.Value}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after struct name")
		return tempStructInstantiation
	}
	tempStructInstantiation.LeftBrace = p.CurrentToken
	if p.Peek(1).TokenType != lexer.RBRACE {
		for p.CurrentToken.TokenType != lexer.RBRACE {
			p.Step()
			tempExpr := p.ParseExpr()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RBRACE {
				p.PrintBacktrace()
				p.HandleFatal("Expected comma between function call arguments")
				return tempStructInstantiation
			}
			tempStructInstantiation.Args = append(tempStructInstantiation.Args, tempExpr)
		}
	} else {
		p.Step()
	}
	// check if the args list is empty
	if len(tempStructInstantiation.Args) == 0 {
		// if it is empty, it means that the struct is instantiated without with the default values
		tempStructInstantiation.Args = nil
	}
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after struct instantiation arguments")
		return tempStructInstantiation
	}
	tempStructInstantiation.RightBrace = p.CurrentToken
	p.Step()
	return tempStructInstantiation
}

// ParseType parses a valid type
func (p *Parser) ParseType() (string, bool) {
	p.Step()
	if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
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
	} else {
		p.Back()
	}
	return tempType
}

// ParseVariableAssign parses a variable assignment
func (p *Parser) ParseVariableAssign(lhs Expr) Stmt {
	Var := p.CurrentToken
	Opp := ""
	var toAssign []Expr
	if lhs == nil {
		toAssign = p.ParseVariableAssignSide()
	} else {
		toAssign = append(toAssign, lhs)
	}
	if _, ok := AssignOperators[p.CurrentToken.Value]; ok {
		Opp = p.CurrentToken.Value
	} else {
		p.HandleFatal("Unknown assignement opperator \" " + p.CurrentToken.Value + "\n in variable assignement")
		return nil
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

	if p.CurrentToken.TokenType == lexer.PERIOD {
		p.Step()
		selectorDepth++
		exp = p.ParseSelector(exp)
		selectorDepth--
		// check if exp.Expr is a Literal
		if _, ok := exp.(SelectorExpr).Expr.(Literal); ok && selectorDepth == 0 {
			p.CurrentFile.AddDependency(exp.(SelectorExpr).Expr.(Literal).Token.Value)
		}
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
		if lookAhead.TokenType == lexer.LPAREN {
			return p.ParseFunctionCallExpr()
		}
		if lookAhead.TokenType == lexer.LBRACE {
			if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
				if _, ok2 := DefaultVarTypes[p.CurrentToken.Value]; !ok2 {
					return p.ParseStructInstantiation()
				}
			}
		}
	}
	return p.ParseLiteral()
}

func (p *Parser) ParseSelector(x Expr) Expr {
	if p.CurrentToken.TokenType != lexer.TEXT {
		p.HandleFatal("Expected field name after '.'")
		return nil
	}
	selector := p.ParseExpr()
	return SelectorExpr{Field: p.CurrentToken, Expr: x, Sel: selector}
}

// ParseParenExpr parses a parenthesized expression
func (p *Parser) ParseParenExpr() Expr {
	tempParentExpr := ParenExpr{}
	tempParentExpr.Lparen = p.CurrentToken
	p.Step()
	tempParentExpr.Expression = p.ParseExpr()
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after expression")
		return nil
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
		p.HandleFatal("Expected ']' after array literal")
		return nil
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
			p.HandleFatal("Expected ':' after map key")
			return nil
		}
		p.Step()
		tempMapLiteral.Values = append(tempMapLiteral.Values, p.ParseExpr())
		if p.CurrentToken.TokenType != lexer.COMMA {
			break
		}
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after map literal")
		return nil
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
		p.HandleFatal("Expected opening double quote")
		return nil
	}
	p.Step()
	if p.CurrentToken.TokenType != lexer.STRING {
		p.HandleFatal("Expected package name")
		return nil
	}
	tempImportStmt.ModulePath = p.CurrentToken.Value
	p.CurrentFile.AddImport(p.CurrentToken.Value)
	p.Step()
	if p.CurrentToken.TokenType != lexer.DQUOTE {
		p.HandleFatal("Expected closing double quote")
		return nil
	}

	p.Step()
	return tempImportStmt
}

func (p *Parser) ParseAnonymousFunctionExpr() Expr {
	tempAnonymousFunctionDecl := AnonymousFunctionExpr{FunctionToken: p.CurrentToken}
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.PrintBacktrace()
		p.HandleFatal("Expected '(' after function")
		return nil
	}
	tempAnonymousFunctionDecl.Prototype = p.ParsePrototype()
	p.Step()
	tempAnonymousFunctionDecl.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after function body")
		return nil
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
					return nil
				}
				exprArray = append(exprArray, tempExpr)
			}
		} else {
			p.Step()
		}

		tempAnonymousFunctionCall.Args = exprArray
		if p.CurrentToken.TokenType != lexer.RPAREN {
			p.HandleFatal("Expected ')' after Anonymous function call arguments")
			return nil
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
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as function name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as function name")
			return nil
		}
	} else {
		p.HandleFatal("Expected function name instead of " + p.CurrentToken.Value)
		return nil
	}
	tempFunctionDecl.Name = p.CurrentToken.Value
	p.Step()
	if p.CurrentToken.TokenType != lexer.LPAREN {
		p.PrintBacktrace()
		p.HandleFatal("Expected '(' after function name")
		return nil
	}
	tempFunctionDecl.Prototype = p.ParsePrototype()
	p.Step()
	tempFunctionDecl.Body = p.ParseBody()
	if p.CurrentToken.TokenType != lexer.RBRACE {
		p.HandleFatal("Expected '}' after function body")
		return nil
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
		p.HandleFatal("Expected comma between return values")
		return nil
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
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as variable name")
			return nil
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
			return nil
		}
		if _, ok := p.VarTypes[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use type name " + p.CurrentToken.Value + " as variable name")
			return nil
		}
		if _, ok := BuiltInFunctions[p.CurrentToken.Value]; ok {
			p.HandleFatal("Cannot use built-in function name " + p.CurrentToken.Value + " as variable name")
			return nil
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
				p.HandleFatal("Expected string value between double quotes")
				return nil
			}
			// check if there is any \n or \r in the string
			if strings.Contains(p.CurrentToken.Value, "\n") || strings.Contains(p.CurrentToken.Value, "\r") {
				p.HandleFatal("strings cannot continue on multiple lines")
				return nil
			}
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.STRING, Value: p.CurrentToken.Value}
		}

		p.Step()
		if p.CurrentToken.TokenType != lexer.DQUOTE {
			p.HandleFatal("Expected '\"' after string value")
			return nil
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
				p.HandleFatal("Expected char value between single quotes")
				return nil
			}
			// check if there is any \n or \r in the char
			if strings.Contains(p.CurrentToken.Value, "\n") || strings.Contains(p.CurrentToken.Value, "\r") {
				p.HandleFatal("chars cannot continue on multiple lines")
				return nil
			}
			tempLiteral = Literal{Token: p.CurrentToken, Type: lexer.CHAR, Value: p.CurrentToken.Value}
		}

		p.Step()
		if p.CurrentToken.TokenType != lexer.SQUOTE {
			p.HandleFatal("Expected ' after char value")
			return nil
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
	p.HandleFatal("Expected a literal instead of " + p.CurrentToken.Value)
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
		p.HandleFatal("Expected '(' after function name")
		return tempFunctionPrototype
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
				p.HandleFatal("Expected ':' after parameter name")
				return tempFunctionPrototype
			}
			ParamType, succes := p.ParseType()
			if !succes {
				p.HandleFatal("Unknown type " + p.CurrentToken.Value + " for parameter " + ParamName)
				return tempFunctionPrototype
			}
			p.Back()
			if !(DuplicateParam(tempFunctionPrototype.Parameters, ParamName)) {
				newParams := FunctionParams{Name: ParamName, Type: ParamType}
				tempFunctionPrototype.Parameters = append(tempFunctionPrototype.Parameters, newParams)
			} else {
				p.HandleFatal("Duplicate parameter " + ParamName)
				return tempFunctionPrototype
			}
			p.Step()
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.HandleFatal("Expected ',' between parameter type")
				return tempFunctionPrototype
			}
		}
	} else {
		p.Step()
	}
	if p.CurrentToken.TokenType != lexer.RPAREN {
		p.HandleFatal("Expected ')' after function parameters")
		return tempFunctionPrototype
	}
	tempFunctionPrototype.RightParamParen = p.CurrentToken
	if p.Peek(1).TokenType == lexer.LPAREN {
		p.Step()
		tempFunctionPrototype.LeftRetsParen = p.CurrentToken
		for p.CurrentToken.TokenType != lexer.RPAREN {
			retType, success := p.ParseType()
			if !success {
				p.HandleFatal("Unknown type " + p.CurrentToken.Value + " for return type")
				return tempFunctionPrototype
			}
			tempFunctionPrototype.ReturnTypes = append(tempFunctionPrototype.ReturnTypes, retType)
			if p.CurrentToken.TokenType != lexer.COMMA && p.CurrentToken.TokenType != lexer.RPAREN {
				p.HandleFatal("Expected ',' between return type")
				return tempFunctionPrototype
			}
		}
		if p.CurrentToken.TokenType != lexer.RPAREN {
			p.HandleFatal("Expected ')' after return types")
			return tempFunctionPrototype
		}
	}
	p.Step()
	tempFunctionPrototype.RightRetsParen = p.CurrentToken
	if p.CurrentToken.TokenType != lexer.LBRACE {
		p.HandleFatal("Expected '{' after function prototype")
		return tempFunctionPrototype
	}
	return tempFunctionPrototype
}
