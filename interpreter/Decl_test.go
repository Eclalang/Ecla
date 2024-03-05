package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

// New

func TestNewINT(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.INT, Value: "0"}, &Env{})
	if bus.GetVal().GetType() != "int" {
		t.Error("Expected INT, got", bus.GetVal().GetType())
	}
}

func TestNewSTRING(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.STRING, Value: "test"}, &Env{})
	if bus.GetVal().GetType() != "string" {
		t.Error("Expected STRING, got", bus.GetVal().GetType())
	}
	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.STRING, Value: "\n"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewBool(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.BOOL, Value: "true"}, &Env{})
	if bus.GetVal().GetType() != "bool" {
		t.Error("Expected BOOL, got", bus.GetVal().GetType())
	}

	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.BOOL, Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewFLOAT(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.FLOAT, Value: "0.0"}, &Env{})
	if bus.GetVal().GetType() != "float" {
		t.Error("Expected FLOAT, got", bus.GetVal().GetType())
	}
}

func TestNewCHAR(t *testing.T) {
	bus := New(parser.Literal{Type: lexer.CHAR, Value: "a"}, &Env{})
	if bus.GetVal().GetType() != "char" {
		t.Error("Expected CHAR, got", bus.GetVal().GetType())
	}

	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus = New(parser.Literal{Type: lexer.CHAR, Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewVAR(t *testing.T) {
	s, err := eclaType.NewString("test")
	if err != nil {
		t.Error("Error creating string: ", err.Error())
	}
	env := NewEnv()
	v, err := eclaType.NewVar("test", "", s)
	if err != nil {
		t.Error("Error creating var: ", err.Error())
	}
	env.SetVar("test", v)

	bus := New(parser.Literal{Type: "VAR", Value: "test"}, env)
	if bus.GetVal().GetType() != "string" {
		t.Error("Expected STRING, got", bus.GetVal().GetType())
	}

	var errCheck = false
	env.ErrorHandle.HookExit(
		func(i int) {
			errCheck = true
		})

	bus = New(parser.Literal{Type: "VAR", Value: "test2"}, env)
	if !errCheck {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

func TestNewNull(t *testing.T) {
	bus := New(parser.Literal{Type: "NULL", Value: "null"}, &Env{})
	if bus.GetVal().GetType() != "" {
		t.Error("Expected NULL, got", bus.GetVal().GetType())
	}
}

func TestNewIfNodeNotExist(t *testing.T) {
	env := NewEnv()
	var err = false
	env.ErrorHandle.HookExit(
		func(i int) {
			err = true
		})
	bus := New(parser.Literal{Type: "NODE", Value: "test"}, env)
	if !err {
		t.Error("Expected error, got", bus.GetVal().GetType())
	}
}

// RunVariableDecl

func TestRunVariableDeclIfValueIsNil(t *testing.T) {

	env := NewEnv()
	s, err := eclaType.NewString("test")
	if err != nil {
		t.Error("Error creating string: ", err.Error())
	}
	v, err := eclaType.NewVar("test", "", s)
	if err != nil {
		t.Error("Error creating var: ", err.Error())
	}
	env.SetVar("test", v)

	decl := parser.VariableDecl{
		Name:  "test",
		Value: nil,
		Type:  "string",
	}

	var errCheck = false
	env.ErrorHandle.HookExit(
		func(i int) {
			errCheck = true
		})
	RunVariableDecl(decl, env)
	if !errCheck {
		t.Error("Expected error, got", decl.Name)
	}
	env.ErrorHandle.RestoreExit()

	decl = parser.VariableDecl{
		Name:  "testInt",
		Value: nil,
		Type:  parser.Int,
	}

	RunVariableDecl(decl, env)
	if !env.Vars.CheckIfVarExistsInCurrentScope("testInt") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testInt"))
	}

	decl = parser.VariableDecl{
		Name:  "testString",
		Value: nil,
		Type:  parser.String,
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testString") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testString"))
	}

	decl = parser.VariableDecl{
		Name:  "testBool",
		Value: nil,
		Type:  parser.Bool,
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testBool") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testBool"))
	}

	decl = parser.VariableDecl{
		Name:  "testFloat",
		Value: nil,
		Type:  parser.Float,
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testFloat") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testFloat"))
	}

	decl = parser.VariableDecl{
		Name:  "testChar",
		Value: nil,
		Type:  parser.Char,
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testChar") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testChar"))
	}

	decl = parser.VariableDecl{
		Name:  "testAny",
		Value: nil,
		Type:  parser.Any,
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testAny") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testAny"))
	}

	decl = parser.VariableDecl{
		Name:  "testArray",
		Value: nil,
		Type:  "[]int",
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testArray") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testArray"))
	}

	decl = parser.VariableDecl{
		Name:  "testMap",
		Value: nil,
		Type:  "map[int]string",
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testMap") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testMap"))
	}

	env.AddTypeDecl(eclaDecl.TypeDecl(
		&eclaDecl.StructDecl{
			Name:   "Test",
			Fields: map[string]string{},
		}))

	decl = parser.VariableDecl{
		Name:  "testStruct",
		Value: nil,
		Type:  "Test",
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testStruct") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testStruct"))
	}
}

func TestRunVariableDeclIfValueIsNotNil(t *testing.T) {
	env := NewEnv()

	decl := parser.VariableDecl{
		Name: "test",
		Value: parser.AnonymousFunctionExpr{
			Prototype: parser.FunctionPrototype{
				Parameters:  make([]parser.FunctionParams, 0),
				ReturnTypes: make([]string, 0),
			},
			Body: make([]parser.Node, 0),
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("test") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("test"))
	}

	//Add Overload

	decl = parser.VariableDecl{
		Name: "test",
		Value: parser.AnonymousFunctionExpr{
			Prototype: parser.FunctionPrototype{
				Parameters:  make([]parser.FunctionParams, 0),
				ReturnTypes: []string{parser.Int},
			},
			Body: make([]parser.Node, 0),
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("test") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("test"))
	}

	decl = parser.VariableDecl{
		Name: "testInt",
		Value: parser.Literal{
			Type:  lexer.INT,
			Value: "0",
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("test") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("test"))
	}

	decl = parser.VariableDecl{
		Name: "testBool",
		Value: parser.Literal{
			Type:  lexer.BOOL,
			Value: "true",
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testBool") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testBool"))
	}

	decl = parser.VariableDecl{
		Name: "testStr",
		Value: parser.Literal{
			Type:  lexer.STRING,
			Value: "test",
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testStr") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testStr"))
	}

	decl = parser.VariableDecl{
		Name: "testFloat",
		Value: parser.Literal{
			Type:  lexer.FLOAT,
			Value: "1.1",
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testFloat") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testFloat"))
	}

	decl = parser.VariableDecl{
		Name: "testChar",
		Value: parser.Literal{
			Type:  lexer.CHAR,
			Value: "a",
		},
	}

	RunVariableDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("testChar") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("testChar"))
	}

}

func TestRunArrayLiteral(t *testing.T) {
	env := NewEnv()

	decl := parser.ArrayLiteral{
		Values: []parser.Expr{
			parser.Literal{
				Type:  lexer.INT,
				Value: "0",
			},
			parser.Literal{
				Type:  lexer.INT,
				Value: "1",
			},
		},
	}

	b := RunArrayLiteral(decl, env)
	if b.GetVal().GetType() != "[]int" {
		t.Error("Expected []int, got", b.GetVal().GetType())
	}

	decl = parser.ArrayLiteral{
		Values: []parser.Expr{},
	}

	b = RunArrayLiteral(decl, env)
	if b.GetVal().GetType() != "empty" {
		t.Error("Expected empty, got", b.GetVal().GetType())
	}
}

func TestRunFunctionDecl(t *testing.T) {
	env := NewEnv()

	decl := parser.FunctionDecl{
		Name: "test",
		Prototype: parser.FunctionPrototype{
			Parameters:  make([]parser.FunctionParams, 0),
			ReturnTypes: make([]string, 0),
		},
	}

	RunFunctionDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("test") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("test"))
	}

	s, err := eclaType.NewString("test")
	if err != nil {
		t.Error("Error creating string: ", err.Error())
	}
	v, err := eclaType.NewVar("testString", "", s)
	env.SetVar("testString", v)

	decl = parser.FunctionDecl{
		Name: "testString",
		Prototype: parser.FunctionPrototype{
			Parameters:  make([]parser.FunctionParams, 0),
			ReturnTypes: []string{parser.String},
		},
	}

	var errCheck = false
	env.ErrorHandle.HookExit(
		func(i int) {
			errCheck = true
		})

	RunFunctionDecl(decl, env)

	if !errCheck {
		t.Error("Expected error, got", decl.Name)
	}

	env.ErrorHandle.RestoreExit()

	decl = parser.FunctionDecl{
		Name: "test",
		Prototype: parser.FunctionPrototype{
			Parameters:  []parser.FunctionParams{{Name: "test", Type: parser.String}},
			ReturnTypes: []string{parser.String},
		},
	}

	RunFunctionDecl(decl, env)

	if !env.Vars.CheckIfVarExistsInCurrentScope("test") {
		t.Error("Expected true, got", env.Vars.CheckIfVarExistsInCurrentScope("test"))
	}
}
