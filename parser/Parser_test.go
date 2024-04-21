package parser

import (
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
	resetWithTokens(&par, lexer.Lexer("struct test{}"))
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

}
