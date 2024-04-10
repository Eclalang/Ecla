package parser

import (
	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var t = lexer.Lexer("import \"console\";")
var unresolved1Tokens = lexer.Lexer("console.println(\"not working\");")
var unresolved2Tokens = lexer.Lexer("console.println(\"not working\");math.abs(-10);")
var helloWorld = lexer.Lexer("import \"console\";console.println(\"Hello, World!\");")
var eol = lexer.Lexer("{};")
var e = errorHandler.NewHandler()
var TestParser = Parser{Tokens: t, ErrorHandler: e}

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
}

func TestParser_ParseFile(t *testing.T) {
	// save the current state of the parser
	par := TestParser
	par.Tokens = helloWorld
	par.TokenIndex = 0
	par.CurrentToken = par.Tokens[0]
	par.ParseFile()
}
