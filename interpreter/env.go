package interpreter

import (
	"encoding/json"
	"fmt"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
	"os"
	"runtime"

	"github.com/tot0p/Ecla/interpreter/eclaType"
)

// Env is the environment in which the code is executed.
type Env struct {
	Vars       map[string]eclaType.Type
	OS         string
	ARCH       string
	SyntaxTree any
	File       string
	Code       string
}

// NewEnv returns a new Env.
func NewEnv() *Env {
	return &Env{
		OS:   runtime.GOOS,
		ARCH: runtime.GOARCH,
		Vars: make(map[string]eclaType.Type),
	}
}

func (env Env) String() string {
	return fmt.Sprintf("Env{OS: %s, ARCH: %s , CODE: %s}", env.OS, env.ARCH, env.Code)
}

// SetCode sets the code to be executed.
func (env *Env) SetCode(code string) {
	env.Code = code
}

// SetFile sets the file to be executed.
func (env *Env) SetFile(file string) {
	env.File = file
}

// SetVar sets the value of the variable with the given name.
func (env *Env) SetVar(name string, value eclaType.Type) {
	env.Vars[name] = value
}

// GetVar returns the value of the variable with the given name.
func (env *Env) GetVar(name string) eclaType.Type {
	return env.Vars[name]
}

// Execute executes Env.Code or Env.File.
func (env *Env) Execute() {
	if env.File != "" {
		env.Code = readFile(env.File)
	}
	// Lexing
	// TODO: SUPPORT FOR MULTIPLE FILES
	tokens := lexer.Lexer(env.Code)
	fmt.Println("TOKENS:", tokens)
	// Parsing
	// TODO: SUPPORT FOR MULTIPLE FILES
	pars := parser.Parser{Tokens: tokens}
	env.SyntaxTree = pars.Parse()
	//DEBUG
	txt, _ := json.MarshalIndent(env.SyntaxTree, "", "  ")
	fmt.Println("SYNTAX TREE:", string(txt))
	//TODO: execute code
}

// readFile reads the file at the given path and returns its contents as a string.
func readFile(file string) string {
	v, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(v)
}
