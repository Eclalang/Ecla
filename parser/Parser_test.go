package parser

import (
	"fmt"
	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var tok = lexer.Lexer("import \"console\";")
var unresolved1Tokens = lexer.Lexer("console.println(\"not working\");")
var unresolved2Tokens = lexer.Lexer("console.println(\"not working\");math.abs(-10);")
var helloWorld = lexer.Lexer("import \"console\";console.println(\"Hello, World!\");")
var eol = lexer.Lexer("{};")

var nodesTok = [][]lexer.Token{
	lexer.Lexer(`mgrlgrl;`),
	lexer.Lexer(`{}`),
	lexer.Lexer(`1+1;`),
	lexer.Lexer(`var test string;`),
	lexer.Lexer(`var test int`),
}

var bodyTok = lexer.Lexer(`{
console.println("Hello, World!");
console.println("This is a test");

var test string;
var test2 int;
var test3 bool;}`)

var textTok = [][]lexer.Token{
	// keyword
	lexer.Lexer(`import "console";`),
	// identifier
	lexer.Lexer(`test();`),
	// selector
	lexer.Lexer(`console.println("Hello, World!");`),
	// selector of a selector
	lexer.Lexer(`data.transform.x;`),
	// selector of a value returned by a function
	lexer.Lexer(`getPoint(p).printSelf(p);`),
	lexer.Lexer(`getPoint(p).x;`),
	lexer.Lexer(`getPoint(p).x.y;`),
}

var blockTok = lexer.Lexer(`{
console.println("Hello, World!");
}`)

var keywords = [][]lexer.Token{
	lexer.Lexer(Var),
	lexer.Lexer(Function + " test(){}"),
	lexer.Lexer(Function + "(){}"),
	lexer.Lexer(Return),
	lexer.Lexer(If + " (true){}"),
	lexer.Lexer(While + " (true){}"),
	lexer.Lexer(For + " (i:=0,i<10,i++){}"),
	lexer.Lexer(Import + " \"console\";"),
	lexer.Lexer(Null),
	lexer.Lexer(Murloc),
	lexer.Lexer(Any),
	lexer.Lexer(Struct + " test{}"),
	lexer.Lexer("randomName"),
}

var e = errorHandler.NewHandler()
var TestParser = Parser{Tokens: tok, ErrorHandler: e}

func resetWithTokens(parser *Parser, tokens []lexer.Token) {
	parser.VarTypes = make(map[string]interface{})
	for k, v := range VarTypes {
		parser.VarTypes[k] = v
	}
	tempFile := new(File)
	tempFile.ParseTree = new(AST)
	parser.CurrentFile = tempFile
	parser.Tokens = tokens
	parser.TokenIndex = 0
	parser.CurrentToken = parser.Tokens[0]
}

func TestParser_Step(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	lastIndex := par.TokenIndex
	par.Step()
	if par.TokenIndex == lastIndex {
		t.Errorf("Step() did not advance the token index")
	}
	// step back to the end
	for par.TokenIndex < len(par.Tokens) {
		par.Step()
	}
	par.Step()
	// check if the current token is empty
	var temp = lexer.Token{}
	if par.CurrentToken != temp {
		t.Errorf("Step() did not set the current token to an empty token")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_Back(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	lastIndex := par.TokenIndex
	par.Step()
	par.Back()
	if par.TokenIndex != lastIndex {
		t.Errorf("Back() did not restore the token index")
	}
	lastIndex = par.TokenIndex
	par.Back()
	if par.TokenIndex != lastIndex {
		t.Errorf("Back() backed into an invalid state")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_MultiStep(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	lastIndex := par.TokenIndex
	par.MultiStep(2)
	if par.TokenIndex != lastIndex+2 {
		t.Errorf("MultiStep() did not advance the token index by 2")
	}
	// step back to the end
	for par.TokenIndex < len(par.Tokens) {
		par.Step()
	}
	par.MultiStep(2)
	// check if the current token is empty
	var temp = lexer.Token{}
	if par.CurrentToken != temp {
		t.Errorf("MultiStep() did not set the current token to an empty token")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_MultiBack(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	lastIndex := par.TokenIndex
	par.MultiStep(2)
	par.MultiBack(2)
	if par.TokenIndex != lastIndex {
		t.Errorf("MultiBack() did not restore the token index")
	}
	lastIndex = par.TokenIndex
	par.MultiBack(2)
	if par.TokenIndex != lastIndex {
		t.Errorf("MultiBack() backed into an invalid state")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_Peek(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	lastIndex := par.TokenIndex
	if par.Peek(1) != par.Tokens[lastIndex+1] {
		t.Errorf("Peek() did not return the correct token")
	}
	var temp = lexer.Token{}
	if par.Peek(1000) != temp {
		t.Errorf("Peek() did not return an empty token")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_PrintBacktrace(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.PrintBacktrace()
	// restore the parser to the original state
	par = TestParser
}

func TestParser_HandleWarning(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.HandleWarning("test")
	// restore the parser to the original state
	par = TestParser
}

func TestParser_HandleError(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.HandleError("test")
	// restore the parser to the original state
	par = TestParser
}

func TestParser_HandleFatal(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	par.ErrorHandler.HookExit(f)
	par.HandleFatal("test")
	if !ok {
		t.Errorf("HandleFatal() did not call the exit hook")
	}
	par.ErrorHandler.RestoreExit()
	// restore the parser to the original state
	par = TestParser
}

func TestParser_DisableEOLChecking(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.Tokens = eol
	par.TokenIndex = 0
	par.Step()
	lastToken := par.CurrentToken
	par.DisableEOLChecking()
	if par.CurrentToken == lastToken {
		t.Errorf("DisableEOLChecking() did not change the current token")
	}
	par.MultiStep(2)
	lastToken = par.CurrentToken
	par.DisableEOLChecking()
	if par.CurrentToken != lastToken {
		t.Errorf("DisableEOLChecking() changed the current token")
	}
	// restore the parser to the original state
	par = TestParser
}

func TestParser_Parse(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.Tokens = unresolved1Tokens
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	par.ErrorHandler.HookExit(f)
	par.Parse()
	if !ok {
		t.Errorf("Parse() did not raise the unsatisfied dependancy error")
	}

	par = TestParser
	ok = false
	par.Tokens = unresolved2Tokens
	par.Parse()
	if !ok {
		t.Errorf("Parse() did not raise the unsatisfied dependancy error")
	}

	par = TestParser
	ok = false
	par.Tokens = helloWorld
	par.Parse()
	if ok {
		t.Errorf("Parse() raised an error when it should not")
	}
	e.RestoreExit()
}

func TestParser_ParseFile(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	resetWithTokens(&par, helloWorld)
	par.ParseFile()
}

func TestParser_ParseNode(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	// test the different nodes
	// murloc node
	resetWithTokens(&par, nodesTok[0])
	par.ParseNode()
	// block node
	resetWithTokens(&par, nodesTok[1])
	par.ParseNode()
	// expression node
	resetWithTokens(&par, nodesTok[2])
	par.ParseNode()
	// text node
	resetWithTokens(&par, nodesTok[3])
	par.ParseNode()
	// text node without semicolon
	// hook the error handler
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)
	resetWithTokens(&par, nodesTok[4])
	par.ParseNode()
	if !ok {
		t.Errorf("ParseNode() did not raise the missing semicolon error")
	}
	e.RestoreExit()

	// text node with endOfBrace turned On
	resetWithTokens(&par, nodesTok[3])
	par.IsEndOfBrace = true
	par.ParseNode()
	// check if the endOfBrace is turned off
	if par.IsEndOfBrace {
		t.Errorf("ParseNode() did not turn off the endOfBrace")
	}
}

func TestParser_ParseBody(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	resetWithTokens(&par, bodyTok)
	// skip the first token since it is a brace
	par.Step()
	par.ParseBody()
}

func TestParser_ParseText(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	for _, v := range textTok {
		resetWithTokens(&par, v)
		par.ParseText()
	}

}

func TestParser_ParseBlock(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	resetWithTokens(&par, blockTok)
	par.ParseBlock()
}

func TestParser_ParseKeyword(t *testing.T) {
	// save the current state of the parser
	// hook the error handler to avoid the fatal errors from the keywords not completing

	var f = func(i int) {
	}
	e.HookExit(f)

	par := TestParser
	for _, v := range keywords {
		resetWithTokens(&par, v)
		par.ParseKeyword()
	}
	e.RestoreExit()
}

func TestParser_ParseStructDecl(t *testing.T) {
	// save the current state of the parser
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	par := TestParser
	resetWithTokens(&par, lexer.Lexer("struct test{a : int;}"))
	par.ParseStructDecl()
	if ok {
		t.Errorf("ParseStructDecl() raised an error when it should not")
	}
	ok = false
	// test the struct with type as name
	resetWithTokens(&par, lexer.Lexer("struct int{}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the invalid name error")
	}
	ok = false
	// test the struct with built-in function as name
	resetWithTokens(&par, lexer.Lexer("struct len{}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the invalid name error")
	}
	ok = false
	// test the struct with keyword as name
	resetWithTokens(&par, lexer.Lexer("struct var{}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the invalid name error")
	}
	ok = false
	// test the struct with no name
	resetWithTokens(&par, lexer.Lexer("struct ;{}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the invalid name error")
	}
	ok = false
	// test the struct without left brace
	resetWithTokens(&par, lexer.Lexer("struct test 1 2{}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the missing left brace error")

	}
	ok = false
	// test the struct without semicolon after the fields
	resetWithTokens(&par, lexer.Lexer("struct test{test : int test2 : int}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the missing semicolon error")

	}
	ok = false
	// test the struct without semicolon after the fields 2
	resetWithTokens(&par, lexer.Lexer("struct test{test : int; test2 : int}"))
	par.ParseStructDecl()
	if !ok {
		t.Errorf("ParseStructDecl() did not raise the missing semicolon error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseStructField(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	resetWithTokens(&par, lexer.Lexer("test : int;"))
	par.ParseStructField()
	if ok {
		t.Errorf("ParseStructField() raised an error when it should not")
	}
	ok = false
	// test the field with no name
	resetWithTokens(&par, lexer.Lexer(": int;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the invalid name error")
	}
	ok = false
	// test the field with no type
	resetWithTokens(&par, lexer.Lexer("test : ;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the invalid type error")
	}
	ok = false
	// test the field with no colon
	resetWithTokens(&par, lexer.Lexer("test int;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the missing colon error")
	}
	ok = false
	// test the field with keyword as name
	resetWithTokens(&par, lexer.Lexer("var : int;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the invalid name error")
	}
	ok = false
	// test the field with built-in function as name
	resetWithTokens(&par, lexer.Lexer("len : int;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the invalid name error")
	}
	ok = false
	// test the field with type as name
	resetWithTokens(&par, lexer.Lexer("int : int;"))
	par.ParseStructField()
	if !ok {
		t.Errorf("ParseStructField() did not raise the invalid name error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseIdent(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var f = func(i int) {
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different identifiers
	// a function call
	resetWithTokens(&par, lexer.Lexer("test();"))
	par.ParseIdent()
	// a struct instanciation
	resetWithTokens(&par, lexer.Lexer("TEST{}"))
	par.VarTypes["TEST"] = "struct"
	par.ParseIdent()
	// a variable assignment with a selector
	resetWithTokens(&par, lexer.Lexer("test.a = 1;"))
	par.ParseIdent()
	resetWithTokens(&par, lexer.Lexer("test.a;"))
	par.ParseIdent()
	// an implicit variable declaration
	resetWithTokens(&par, lexer.Lexer("test := 1;"))
	par.ParseIdent()
	// a variable assignment
	resetWithTokens(&par, lexer.Lexer("test = 1;"))
	par.ParseIdent()

	e.RestoreExit()
}

func TestParser_ParseIfStmt(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different if statements
	// normal if statement
	resetWithTokens(&par, lexer.Lexer("if (true){console.println(\"Hello, World!\");}"))
	par.ParseIfStmt()
	if ok {
		t.Errorf("ParseIfStmt() raised an error when it should not")
	}
	ok = false
	// if statement with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("if true){console.println(\"Hello, World!\");}"))
	par.ParseIfStmt()
	if !ok {
		t.Errorf("ParseIfStmt() did not raise the missing left parenthesis error")
	}
	ok = false
	// if statement with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("if (true{console.println(\"Hello, World!\");}"))
	par.ParseIfStmt()
	if !ok {
		t.Errorf("ParseIfStmt() did not raise the missing right parenthesis error")
	}
	ok = false
	// if statement with missing left brace
	resetWithTokens(&par, lexer.Lexer("if (true)console.println(\"Hello, World!\");}"))
	par.ParseIfStmt()
	if !ok {
		t.Errorf("ParseIfStmt() did not raise the missing left brace error")
	}
	ok = false
	// if statement with random text after the right brace
	resetWithTokens(&par, lexer.Lexer("if (true){console.println(\"Hello, World!\");}randomText"))
	par.ParseIfStmt()
	if ok {
		t.Errorf("ParseIfStmt() raised an error when it should not")
	}
	ok = false
	// if statement with else statement
	resetWithTokens(&par, lexer.Lexer("if (true){console.println(\"Hello, World!\");}else{console.println(\"Goodbye, World!\");}"))
	par.ParseIfStmt()
	if ok {
		t.Errorf("ParseIfStmt() raised an error when it should not")
	}
	ok = false
	// if statement with else statement and missing left brace
	resetWithTokens(&par, lexer.Lexer("if (true){console.println(\"Hello, World!\");}else console.println(\"Goodbye, World!\");}"))
	par.ParseIfStmt()
	if !ok {
		t.Errorf("ParseIfStmt() did not raise the missing left brace error")
	}
	ok = false
	// if statement with else if statement
	resetWithTokens(&par, lexer.Lexer("if (true){console.println(\"Hello, World!\");}else if (true){console.println(\"Goodbye, World!\");}"))
	par.ParseIfStmt()
	if ok {
		t.Errorf("ParseIfStmt() raised an error when it should not")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseWhileStmt(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different while statements
	// normal while statement
	resetWithTokens(&par, lexer.Lexer("while (true){console.println(\"Hello, World!\");}"))
	par.ParseWhileStmt()
	if ok {
		t.Errorf("ParseWhileStmt() raised an error when it should not")
	}
	ok = false
	// while statement with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("while true){console.println(\"Hello, World!\");}"))
	par.ParseWhileStmt()
	if !ok {
		t.Errorf("ParseWhileStmt() did not raise the missing left parenthesis error")
	}
	ok = false
	// while statement with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("while (true{console.println(\"Hello, World!\");}"))
	par.ParseWhileStmt()
	if !ok {
		t.Errorf("ParseWhileStmt() did not raise the missing right parenthesis error")
	}
	ok = false
	// while statement with missing left brace
	resetWithTokens(&par, lexer.Lexer("while (true)console.println(\"Hello, World!\");}"))
	par.ParseWhileStmt()
	if !ok {
		t.Errorf("ParseWhileStmt() did not raise the missing left brace error")
	}
	ok = false

}

func TestParser_ParseForStmt(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different for statements
	// normal for statement
	resetWithTokens(&par, lexer.Lexer("for (i:=0,i<10,i++){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if ok {
		t.Errorf("ParseForStmt() raised an error when it should not")
	}
	ok = false
	// normal for statement with variable declaration
	resetWithTokens(&par, lexer.Lexer("for (var i int,i<10,i++){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if ok {
		t.Errorf("ParseForStmt() raised an error when it should not")
	}
	ok = false
	// for statement with range
	resetWithTokens(&par, lexer.Lexer("for (x,y range z){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if ok {
		t.Errorf("ParseForStmt() raised an error when it should not")
	}
	ok = false
	// for statement with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("for i:=0,i<10,i++){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing left parenthesis error")
	}
	ok = false
	// for statement with range and missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("for x,y range z){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing left parenthesis error")
	}
	ok = false
	// for statement with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("for (i:=0,i<10,i{console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing right parenthesis error")
	}
	ok = false
	// for statement with range and missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("for (x,y range z{console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing right parenthesis error")
	}
	ok = false
	// for statement with comma missing between the declaration and the condition
	resetWithTokens(&par, lexer.Lexer("for (i:=0 i<10,i++){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing colon error")
	}
	ok = false
	// for statement with comma missing between the condition and the increment
	resetWithTokens(&par, lexer.Lexer("for (i:=0,i<10 i++){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing colon error")
	}
	ok = false
	// for statement with missing comma between the key and the value
	resetWithTokens(&par, lexer.Lexer("for (x 1 y range z){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing colon error")
	}
	ok = false
	// for statement with an int instead of a range
	resetWithTokens(&par, lexer.Lexer("for (x,y 10){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the invalid range error")
	}
	ok = false
	// for statement with an text other than range
	resetWithTokens(&par, lexer.Lexer("for (x,y test z){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the invalid range error")
	}
	ok = false
	// for statement with missing comma between the key and the value
	resetWithTokens(&par, lexer.Lexer("for (x y range z){console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		fmt.Println("check here")
		t.Errorf("ParseForStmt() did not raise the missing colon error")
	}
	ok = false
	// for statement with missing left brace
	resetWithTokens(&par, lexer.Lexer("for (i:=0,i<10,i++)console.println(\"Hello, World!\");}"))
	par.ParseForStmt()
	if !ok {
		t.Errorf("ParseForStmt() did not raise the missing left brace error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseVariableDecl(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different variable declarations
	// normal variable declaration
	resetWithTokens(&par, lexer.Lexer("var test string;"))
	par.ParseVariableDecl()
	if ok {
		t.Errorf("ParseVariableDecl() raised an error when it should not")
	}
	ok = false
	// variable declaration with missing type
	resetWithTokens(&par, lexer.Lexer("var test;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the missing type error")
	}
	ok = false
	// variable declaration with missing name
	resetWithTokens(&par, lexer.Lexer("var string;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the missing name error")
	}
	ok = false
	// variable declaration with keyword as name
	resetWithTokens(&par, lexer.Lexer("var var string;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// variable declaration with built-in function as name
	resetWithTokens(&par, lexer.Lexer("var len string;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// variable declaration with type as name
	resetWithTokens(&par, lexer.Lexer("var int string;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// variable declaration with somthing other than = or ; after the type
	resetWithTokens(&par, lexer.Lexer("var test string test;"))
	par.ParseVariableDecl()
	if !ok {
		t.Errorf("ParseVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// variable declaration with =
	resetWithTokens(&par, lexer.Lexer("var test string = \"hello\";"))
	par.ParseVariableDecl()
	if ok {
		t.Errorf("ParseVariableDecl() raised an error when it should not")
	}
	ok = false
	// variable declaration with a struct instanciation as value
	resetWithTokens(&par, lexer.Lexer("var t Test = Test{};"))
	par.VarTypes["Test"] = "struct"
	par.ParseVariableDecl()
	if ok {
		t.Errorf("ParseVariableDecl() raised an error when it should not")
	}

	e.RestoreExit()
}

func TestParser_ParseImplicitVariableDecl(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different implicit variable declarations
	// normal implicit variable declaration
	resetWithTokens(&par, lexer.Lexer("test := 1;"))
	par.ParseImplicitVariableDecl()
	if ok {
		t.Errorf("ParseImplicitVariableDecl() raised an error when it should not")
	}
	ok = false
	// implicit variable declaration with missing value
	resetWithTokens(&par, lexer.Lexer("test := ;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the missing value error")
	}
	ok = false
	// implicit variable declaration with missing name
	resetWithTokens(&par, lexer.Lexer(":= 1;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the missing name error")
	}
	ok = false
	// implicit variable declaration with keyword as name
	resetWithTokens(&par, lexer.Lexer("var := 1;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// implicit variable declaration with built-in function as name
	resetWithTokens(&par, lexer.Lexer("len := 1;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// implicit variable declaration with type as name
	resetWithTokens(&par, lexer.Lexer("int := 1;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// implicit variable declaration with something other than := after the name
	resetWithTokens(&par, lexer.Lexer("test = 1;"))
	par.ParseImplicitVariableDecl()
	if !ok {
		t.Errorf("ParseImplicitVariableDecl() did not raise the invalid name error")
	}
	ok = false
	// implicit variable declaration with a struct instanciation as value
	resetWithTokens(&par, lexer.Lexer("t := Test{};"))
	par.VarTypes["Test"] = "struct"
	par.ParseImplicitVariableDecl()
	if ok {
		t.Errorf("ParseImplicitVariableDecl() raised an error when it should not")
	}

	e.RestoreExit()
}

func TestParser_ParseFunctionCallExpr(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different function calls
	// normal function call
	resetWithTokens(&par, lexer.Lexer("test();"))
	par.ParseFunctionCallExpr()
	if ok {
		t.Errorf("ParseFunctionCallExpr() raised an error when it should not")
	}
	ok = false
	// function call with keyword as name
	resetWithTokens(&par, lexer.Lexer("var();"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the invalid name error")
	}
	ok = false
	// function call with type as name
	resetWithTokens(&par, lexer.Lexer("int();"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the invalid name error")
	}
	ok = false
	// function call with something other than a name before the parenthesis
	resetWithTokens(&par, lexer.Lexer("1();"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the invalid name error")
	}
	ok = false
	// function call with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("test);"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the missing left parenthesis error")
	}
	ok = false
	// function call with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("test(1;"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the missing right parenthesis error")
	}
	ok = false
	// function call with missing comma between the arguments
	resetWithTokens(&par, lexer.Lexer("test(1 2);"))
	par.ParseFunctionCallExpr()
	if !ok {
		t.Errorf("ParseFunctionCallExpr() did not raise the missing comma error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseStructInstantiation(t *testing.T) {
	// hook the error handler to avoid the fatal errors from the keywords not completing
	var ok bool
	var f = func(i int) {
		ok = i == 1
	}
	e.HookExit(f)

	// save the current state of the parser
	par := TestParser

	// test the different struct instantiations
	// normal struct instantiation
	resetWithTokens(&par, lexer.Lexer("Test{}"))
	par.ParseStructInstantiation()
	if ok {
		t.Errorf("ParseStructInstantiation() raised an error when it should not")
	}
	ok = false
	// struct instantiation with missing left brace
	resetWithTokens(&par, lexer.Lexer("Test}"))
	par.ParseStructInstantiation()
	if !ok {
		t.Errorf("ParseStructInstantiation() did not raise the missing left brace error")
	}
	ok = false
	// struct instantiation with missing right brace
	resetWithTokens(&par, lexer.Lexer("Test{"))
	par.ParseStructInstantiation()
	if !ok {
		t.Errorf("ParseStructInstantiation() did not raise the missing right brace error")
	}
	ok = false
	// struct instantiation with some parameters
	resetWithTokens(&par, lexer.Lexer("Test{1,\"hello\"}"))
	par.ParseStructInstantiation()
	if ok {
		t.Errorf("ParseStructInstantiation() raised an error when it should not")
	}
	ok = false
	// struct instantiation with missing arguments and missing right brace
	resetWithTokens(&par, lexer.Lexer("Test{1,\"hello\""))
	par.ParseStructInstantiation()
	if !ok {
		t.Errorf("ParseStructInstantiation() did not raise the missing right brace error")
	}
	ok = false
	// struct instantiation with struct instanciation as argument and missing right brace
	resetWithTokens(&par, lexer.Lexer("Test{Test{"))
	par.ParseStructInstantiation()
	if !ok {
		t.Errorf("ParseStructInstantiation() did not raise the missing right brace error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseArrayType(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	// test the different array types
	// normal array type
	resetWithTokens(&par, lexer.Lexer("[]int"))
	if par.ParseArrayType() == "" {
		t.Errorf("ParseArrayType() did not return the correct type")
	}
	// array type with missing type
	resetWithTokens(&par, lexer.Lexer("[]"))
	if par.ParseArrayType() != "" {
		t.Errorf("ParseArrayType() did not return an empty string")
	}
	// array type with missing right bracket
	resetWithTokens(&par, lexer.Lexer("[int"))
	if par.ParseArrayType() != "" {
		t.Errorf("ParseArrayType() did not return an empty string")
	}

	e.RestoreExit()
}

func TestParser_ParseMapType(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	// test the different map types
	// normal map type
	resetWithTokens(&par, lexer.Lexer("map[int]string"))
	if par.ParseMapType() == "" {
		t.Errorf("ParseMapType() did not return the correct type")
	}
	// map type with missing key type
	resetWithTokens(&par, lexer.Lexer("map[string"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}
	// map type with missing value type
	resetWithTokens(&par, lexer.Lexer("map[int"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}
	// map type with missing right bracket
	resetWithTokens(&par, lexer.Lexer("map[intstring"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}
	// map type with missing left bracket
	resetWithTokens(&par, lexer.Lexer("map int]string"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}
	// map type with invalid key type
	resetWithTokens(&par, lexer.Lexer("map[1]string"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}
	// map type with invalid value type
	resetWithTokens(&par, lexer.Lexer("map[int]1"))
	if par.ParseMapType() != "" {
		t.Errorf("ParseMapType() did not return an empty string")
	}

	e.RestoreExit()
}

func TestParser_ParseFunctionType(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	// test the different function types
	// normal function type
	resetWithTokens(&par, lexer.Lexer("functions(int)(string)"))
	if par.ParseFunctionType() == "" {
		t.Errorf("ParseFunctionType() did not return the correct type")
	}
	// function type with missing arguments
	resetWithTokens(&par, lexer.Lexer("functions)(string)"))
	if par.ParseFunctionType() != "" {
		t.Errorf("ParseFunctionType() did not return an empty string")
	}
	// function type with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("functions(intstring"))
	if par.ParseFunctionType() != "" {
		t.Errorf("ParseFunctionType() did not return an empty string")
	}
	// function type with something other than comma between the arguments
	resetWithTokens(&par, lexer.Lexer("functions(int 1 string)(string)"))
	if par.ParseFunctionType() != "" {
		t.Errorf("ParseFunctionType() did not return an empty string")
	}
	// function type with something other than comma between the arguments
	resetWithTokens(&par, lexer.Lexer("functions(int)(string 1 string)"))
	if par.ParseFunctionType() != "" {
		t.Errorf("ParseFunctionType() did not return an empty string")
	}
	// function type with invalid return type
	resetWithTokens(&par, lexer.Lexer("functions(int)(1)"))
	if par.ParseFunctionType() != "" {
		t.Errorf("ParseFunctionType() did not return an empty string")
	}
	// function type with missing return type left parenthesis
	resetWithTokens(&par, lexer.Lexer("functions(int)string)"))
	if par.ParseFunctionType() == "" {
		t.Errorf("ParseFunctionType() did not return the correct type")
	}

	e.RestoreExit()
}

func TestParser_ParseType(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var success bool
	// test the different types and add a random token at the start since the function expects a token before the type by stepping
	// normal type
	resetWithTokens(&par, lexer.Lexer("a int"))
	_, success = par.ParseType()
	if !success {
		t.Errorf("ParseType() did not return the correct type")
	}
	// array type
	resetWithTokens(&par, lexer.Lexer("a []int"))
	_, success = par.ParseType()
	if !success {
		t.Errorf("ParseType() did not return the correct type")
	}
	// map type
	resetWithTokens(&par, lexer.Lexer("a map[int]string"))
	_, success = par.ParseType()
	if !success {
		t.Errorf("ParseType() did not return the correct type")
	}
	// function type
	resetWithTokens(&par, lexer.Lexer("a function(int)(string)"))
	_, success = par.ParseType()
	if !success {
		t.Errorf("ParseType() did not return the correct type")
	}
	// invalid type
	resetWithTokens(&par, lexer.Lexer("a notAVaildType"))
	res, success := par.ParseType()
	fmt.Println(res, success)
	if success {
		t.Errorf("ParseType() did return a type")
	}

	e.RestoreExit()
}

func TestParser_ParseVariableAssign(t *testing.T) {
	// TODO: implement the test later
}

func TestParser_ParseVariableAssignSide(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var expr []Expr

	// test the different variable assign sides
	// normal variable assign side
	resetWithTokens(&par, lexer.Lexer("1"))
	expr = par.ParseVariableAssignSide()
	if len(expr) != 1 {
		t.Errorf("ParseVariableAssignSide() did not return the correct number of expressions")
	}
	// normal variable assign side with multiple expressions
	resetWithTokens(&par, lexer.Lexer("1,2,3"))
	expr = par.ParseVariableAssignSide()
	if len(expr) != 3 {
		t.Errorf("ParseVariableAssignSide() did not return the correct number of expressions")
	}
}

func TestParser_ParseExpr(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var expr Expr

	// test the different expressions
	// normal expression
	resetWithTokens(&par, lexer.Lexer("1"))
	expr = par.ParseExpr()
	if expr == nil {
		t.Errorf("ParseExpr() did not return an expression")
	}
	// normal expression with multiple expressions
	resetWithTokens(&par, lexer.Lexer("1+2"))
	expr = par.ParseExpr()
	if expr == nil {
		t.Errorf("ParseExpr() did not return an expression")
	}
}

func TestParser_ParseBinaryExpr(t *testing.T) {
	// TODO: implement the test later
}

func TestParser_ParseUnaryExpr(t *testing.T) {
	// TODO: implement the test later
}

func TestParser_ParsePrimaryExpr(t *testing.T) {
	// TODO: implement the test later
}

func TestParser_ParseOperand(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var expr Expr

	// test the different operands
	// normal operand
	resetWithTokens(&par, lexer.Lexer("1"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a function call
	resetWithTokens(&par, lexer.Lexer("test()"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a struct instantiation
	resetWithTokens(&par, lexer.Lexer("Test{}"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a variable
	resetWithTokens(&par, lexer.Lexer("test"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a paren expression
	resetWithTokens(&par, lexer.Lexer("(1)"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with an array literal
	resetWithTokens(&par, lexer.Lexer("[1]"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a map literal
	resetWithTokens(&par, lexer.Lexer("{1:\"hello\"}"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a selector
	resetWithTokens(&par, lexer.Lexer("test.test"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operant with anonymous function
	resetWithTokens(&par, lexer.Lexer("function(int)(string){return \"hello\";}"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}
	// operand with a type cast with missing type
	resetWithTokens(&par, lexer.Lexer("(1)"))
	expr = par.ParseOperand()
	if expr == nil {
		t.Errorf("ParseOperand() did not return an expression")
	}

}

func TestParser_ParseSelector(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different selectors
	// normal selector
	resetWithTokens(&par, lexer.Lexer("test.test"))
	par.MultiStep(2)
	par.ParseSelector(nil)
	if ok {
		t.Errorf("ParseSelector() raised an error when it should not")
	}
	ok = false
	// selector with none text field
	resetWithTokens(&par, lexer.Lexer("test.1"))
	par.MultiStep(2)
	par.ParseSelector(nil)
	if !ok {
		t.Errorf("ParseSelector() did not raise the invalid name error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseParenExpr(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different paren expressions
	// normal paren expression
	resetWithTokens(&par, lexer.Lexer("(1)"))
	par.ParseParenExpr()
	if ok {
		t.Errorf("ParseParenExpr() raised an error when it should not")
	}
	ok = false
	// paren expression with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("(1"))
	par.ParseParenExpr()
	if !ok {
		t.Errorf("ParseParenExpr() did not raise the missing right parenthesis error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseArrayLiteral(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different array literals
	// normal array literal
	resetWithTokens(&par, lexer.Lexer("[1]"))
	par.ParseArrayLiteral()
	if ok {
		t.Errorf("ParseArrayLiteral() raised an error when it should not")
	}
	ok = false
	// array literal with missing right bracket
	resetWithTokens(&par, lexer.Lexer("[1"))
	par.ParseArrayLiteral()
	if !ok {
		t.Errorf("ParseArrayLiteral() did not raise the missing right bracket error")
	}
	ok = false
	// array literal with multiple elements
	resetWithTokens(&par, lexer.Lexer("[1,2,3]"))
	par.ParseArrayLiteral()
	if ok {
		t.Errorf("ParseArrayLiteral() raised an error when it should not")
	}

	e.RestoreExit()
}

func TestParser_ParseMapLiteral(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different map literals
	// normal map literal
	resetWithTokens(&par, lexer.Lexer("{1:\"hello\"}"))
	par.ParseMapLiteral()
	if ok {
		t.Errorf("ParseMapLiteral() raised an error when it should not")
	}
	ok = false
	// map literal with missing right brace
	resetWithTokens(&par, lexer.Lexer("{1:\"hello\""))
	par.ParseMapLiteral()
	if !ok {
		t.Errorf("ParseMapLiteral() did not raise the missing right brace error")
	}
	ok = false
	// map literal with multiple elements
	resetWithTokens(&par, lexer.Lexer("{1:\"hello\",2:\"world\"}"))
	par.ParseMapLiteral()
	if ok {
		t.Errorf("ParseMapLiteral() raised an error when it should not")
	}
	// map literal with missing colon
	resetWithTokens(&par, lexer.Lexer("{1\"hello\"}"))
	par.ParseMapLiteral()
	if !ok {
		t.Errorf("ParseMapLiteral() did not raise the missing colon error")
	}

	e.RestoreExit()
}

func TestParser_ParseImportStmt(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different import statements
	// normal import statement
	resetWithTokens(&par, lexer.Lexer("import \"console\""))
	par.ParseImportStmt()
	if ok {
		t.Errorf("ParseImportStmt() raised an error when it should not")
	}
	ok = false
	// import statement with missing string
	resetWithTokens(&par, lexer.Lexer("import"))
	par.ParseImportStmt()
	if !ok {
		t.Errorf("ParseImportStmt() did not raise the missing string error")
	}
	// import statement with something other than a string as package name
	resetWithTokens(&par, lexer.Lexer("import \"\""))
	par.ParseImportStmt()
	if !ok {
		t.Errorf("ParseImportStmt() did not raise the invalid package name error")
	}
	// import statement with missing opening double quote
	resetWithTokens(&par, lexer.Lexer("import console\""))
	par.ParseImportStmt()
	if !ok {
		t.Errorf("ParseImportStmt() did not raise the missing opening double quote error")
	}
	// import statement with missing closing double quote
	resetWithTokens(&par, lexer.Lexer("import \"console;"))
	par.ParseImportStmt()
	if !ok {
		t.Errorf("ParseImportStmt() did not raise the missing closing double quote error")
	}

	e.RestoreExit()
}

func TestParser_ParseAnonymousFunctionExpr(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different anonymous function expressions
	// normal anonymous function expression
	resetWithTokens(&par, lexer.Lexer("function (a : int, b : int) {console.println(a + b);}"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false
	// anonymous function expression with missing arguments
	resetWithTokens(&par, lexer.Lexer("function () {console.println(\"hello\");}"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false
	// anonymous function expression with missing return type
	resetWithTokens(&par, lexer.Lexer("function (a : int, b : int) {console.println(a + b);}"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false
	// anonymous function expression with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("function a : int)(string){return \"hello\";}"))
	par.ParseAnonymousFunctionExpr()
	if !ok {
		t.Errorf("ParseAnonymousFunctionExpr() did not raise the missing left parenthesis error")
	}
	ok = false
	// anonymous function expression with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("function(a : int(string){return \"hello\";}"))
	par.ParseAnonymousFunctionExpr()
	if !ok {
		t.Errorf("ParseAnonymousFunctionExpr() did not raise the missing right parenthesis error")
	}
	ok = false
	// anonymous function expression with missing left brace
	resetWithTokens(&par, lexer.Lexer("function(a : int)(string)return \"hello\";}"))
	par.ParseAnonymousFunctionExpr()
	if !ok {
		t.Errorf("ParseAnonymousFunctionExpr() did not raise the missing left brace error")
	}
	ok = false

	// anonymous function expression with call just after the function
	resetWithTokens(&par, lexer.Lexer("function(a : int)(string){return \"hello\";}(1)"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false
	// anonymous function expression with call just after the function with multiple arguments
	resetWithTokens(&par, lexer.Lexer("function(a : int)(string){return \"hello\";}(1,2,3)"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false
	// anonymous function expression with call just after the function with missing comma
	resetWithTokens(&par, lexer.Lexer("function(a : int)(string){return \"hello\";}(1 2 3)"))
	par.ParseAnonymousFunctionExpr()
	if !ok {
		t.Errorf("ParseAnonymousFunctionExpr() did not raise the missing comma error")
	}
	ok = false
	// anonymous function expression with call just after the function with no arguments
	resetWithTokens(&par, lexer.Lexer("function(a : int)(string){return \"hello\";}()"))
	par.ParseAnonymousFunctionExpr()
	if ok {
		t.Errorf("ParseAnonymousFunctionExpr() raised an error when it should not")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseFunctionDecl(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different function declarations
	// normal function declaration
	resetWithTokens(&par, lexer.Lexer("function test(a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if ok {
		t.Errorf("ParseFunctionDecl() raised an error when it should not")
	}
	ok = false
	// function declaration with keyword as name
	resetWithTokens(&par, lexer.Lexer("function var(a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if !ok {
		t.Errorf("ParseFunctionDecl() did not raise the invalid name error")
	}
	ok = false
	// function declaration with type as name
	resetWithTokens(&par, lexer.Lexer("function int(a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if !ok {
		t.Errorf("ParseFunctionDecl() did not raise the invalid name error")
	}
	ok = false
	// function declaration with built-in function as name
	resetWithTokens(&par, lexer.Lexer("function len(a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if !ok {
		t.Errorf("ParseFunctionDecl() did not raise the invalid name error")
	}
	ok = false
	// function declaration with something other than a name before the parenthesis
	resetWithTokens(&par, lexer.Lexer("function 1(a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if !ok {
		t.Errorf("ParseFunctionDecl() did not raise the invalid name error")
	}
	ok = false
	// function declaration with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("function test a : int, b : int)( string) {return a + b;}"))
	par.ParseFunctionDecl()
	if !ok {
		t.Errorf("ParseFunctionDecl() did not raise the missing left parenthesis error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseReturnStmt(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different return statements
	// normal return statement
	resetWithTokens(&par, lexer.Lexer("return 1;"))
	par.ParseReturnStmt()
	if ok {
		t.Errorf("ParseReturnStmt() raised an error when it should not")
	}
	ok = false
	// return statement with missing expression
	resetWithTokens(&par, lexer.Lexer("return;"))
	par.ParseReturnStmt()
	if ok {
		t.Errorf("ParseReturnStmt() raised an error when it should not")
	}
	ok = false
	// return statement with multiple expressions
	resetWithTokens(&par, lexer.Lexer("return 1,2,3;"))
	par.ParseReturnStmt()
	if ok {
		t.Errorf("ParseReturnStmt() raised an error when it should not")
	}

	e.RestoreExit()
}

func TestParser_ParseIndexableAccessExpr(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different indexable access expressions
	// normal indexable access expression
	resetWithTokens(&par, lexer.Lexer("test[1]"))
	par.ParseIndexableAccessExpr()
	if ok {
		t.Errorf("ParseIndexableAccessExpr() raised an error when it should not")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseVariableAccess(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different variable accesses
	// normal variable access
	resetWithTokens(&par, lexer.Lexer("test"))
	par.ParseVariableAccess()
	if ok {
		t.Errorf("ParseVariableAccess() raised an error when it should not")
	}
	ok = false
	// variable access with keyword as name
	resetWithTokens(&par, lexer.Lexer("var"))
	par.ParseVariableAccess()
	if !ok {
		t.Errorf("ParseVariableAccess() did not raise the invalid name error")
	}
	ok = false
	// variable access with type as name
	resetWithTokens(&par, lexer.Lexer("int"))
	par.ParseVariableAccess()
	if !ok {
		t.Errorf("ParseVariableAccess() did not raise the invalid name error")
	}
	ok = false
	// variable access with built-in function as name
	resetWithTokens(&par, lexer.Lexer("len"))
	par.ParseVariableAccess()
	if !ok {
		t.Errorf("ParseVariableAccess() did not raise the invalid name error")
	}
	ok = false
	// variable access with indexable access
	resetWithTokens(&par, lexer.Lexer("test[1]"))
	par.ParseVariableAccess()
	if ok {
		t.Errorf("ParseVariableAccess() raised an error when it should not")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParseLiteral(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different literals
	// int literal
	resetWithTokens(&par, lexer.Lexer("1"))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// text literal
	resetWithTokens(&par, lexer.Lexer("hello"))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// text with keyword as name
	resetWithTokens(&par, lexer.Lexer("var"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the invalid name error")
	}
	ok = false
	// text name null
	resetWithTokens(&par, lexer.Lexer("null"))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// text with type as name
	resetWithTokens(&par, lexer.Lexer("int"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the invalid name error")
	}
	ok = false
	// text with built-in function as name
	resetWithTokens(&par, lexer.Lexer("len"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the invalid name error")
	}
	ok = false
	// string literal
	resetWithTokens(&par, lexer.Lexer("\"hello\""))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// string literal with only double quote
	resetWithTokens(&par, lexer.Lexer("\"\""))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// string literal with missing double quote at the end
	resetWithTokens(&par, lexer.Lexer("\"hello;"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the missing double quote error")
	}
	ok = false
	// error with multiline string
	resetWithTokens(&par, lexer.Lexer("\"hello\nworld\""))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the invalid string error")
	}
	ok = false

	// char literal
	resetWithTokens(&par, lexer.Lexer("'a'"))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// char literal with only single quote
	resetWithTokens(&par, lexer.Lexer("''"))
	par.ParseLiteral()
	if ok {
		t.Errorf("ParseLiteral() raised an error when it should not")
	}
	ok = false
	// char literal with missing single quote at the end
	resetWithTokens(&par, []lexer.Token{{TokenType: lexer.SQUOTE, Value: "'"}, {TokenType: lexer.TEXT, Value: "a"}})
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the missing single quote error")
	}
	ok = false
	// char literal with multiline
	resetWithTokens(&par, lexer.Lexer("'a\nb'"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the invalid char error")
	}
	ok = false
	// char literal with missing ending single quote
	resetWithTokens(&par, lexer.Lexer("'a;"))
	par.ParseLiteral()
	if !ok {
		t.Errorf("ParseLiteral() did not raise the missing single quote error")
	}
	ok = false

	e.RestoreExit()
}

func TestParser_ParsePrototype(t *testing.T) {
	// save the current state of the parser
	par := TestParser

	var ok bool

	e.HookExit(func(i int) {
		ok = i == 1
	})

	// test the different prototypes
	// normal prototype
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)(string){return \"hello\";}"))
	par.ParsePrototype()
	if ok {
		t.Errorf("ParsePrototype() raised an error when it should not")
	}
	ok = false
	// prototype with missing arguments
	resetWithTokens(&par, lexer.Lexer("()(string){return \"hello\";}"))
	par.ParsePrototype()
	if ok {
		t.Errorf("ParsePrototype() raised an error when it should not")
	}
	ok = false
	// prototype with missing return type
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int){return \"hello\";}"))
	par.ParsePrototype()
	if ok {
		t.Errorf("ParsePrototype() raised an error when it should not")
	}
	ok = false
	// prototype with missing left parenthesis
	resetWithTokens(&par, lexer.Lexer("a : int, b : int)(string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing left parenthesis error")
	}
	ok = false
	// prototype with missing colon
	resetWithTokens(&par, lexer.Lexer("(a int, b : int)(string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing colon error")
	}
	ok = false
	// prototype with wrong type
	resetWithTokens(&par, lexer.Lexer("(a : test, b : int)(string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the invalid type error")
	}
	ok = false
	// prototype with duplicate argument name
	resetWithTokens(&par, lexer.Lexer("(a : int, a : int)(string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the duplicate argument error")
	}
	ok = false
	// prototype with missing right parenthesis
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int(string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing right parenthesis error")
	}
	ok = false
	// prototype with wrong return type
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)(test){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the invalid return type error")
	}
	ok = false
	// prototype with missing return left parenthesis
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing left parenthesis error")
	}
	ok = false
	// prototype with missing return right parenthesis
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)(string{return \"hello\"};"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing right parenthesis error")
	}
	ok = false
	// prototype with multiple return types
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)(string, string){return \"hello\";}"))
	par.ParsePrototype()
	if ok {
		t.Errorf("ParsePrototype() raised an error when it should not")
	}
	ok = false
	// prototype with missing colon in return type
	resetWithTokens(&par, lexer.Lexer("(a : int, b : int)(string string){return \"hello\";}"))
	par.ParsePrototype()
	if !ok {
		t.Errorf("ParsePrototype() did not raise the missing colon in return type error")
	}
	ok = false

	e.RestoreExit()

}
