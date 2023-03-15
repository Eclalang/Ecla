package interpreter

import (
	"fmt"
	"github.com/tot0p/Ecla/errorHandler"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
	"os"
	"runtime"
)

// Env is the environment in which the code is executed.
type Env struct {
	Vars       *Scope
	OS         string
	ARCH       string
	SyntaxTree *parser.File
	Tokens     []lexer.Token
	File       string
	Code       string
	Libs       map[string]libs.Lib
	//Func        map[string]*eclaKeyWord.Function
	ErrorHandle *errorHandler.ErrorHandler
}

// NewEnv returns a new Env.
func NewEnv() *Env {
	return &Env{
		OS:   runtime.GOOS,
		ARCH: runtime.GOARCH,
		Vars: NewScopeMain(),
		Libs: make(map[string]libs.Lib),
		//Func:        make(map[string]*eclaKeyWord.Function),
		ErrorHandle: errorHandler.NewHandler(),
	}
}

func (env *Env) String() string {
	return fmt.Sprintf("Env{OS: %s, ARCH: %s , CODE: %s , VAR : %s}", env.OS, env.ARCH, env.Code, env.Vars)
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
func (env *Env) SetVar(name string, value *eclaType.Var) {
	env.Vars.Set(name, value)
}

// GetVar returns the value of the variable with the given name.
func (env *Env) GetVar(name string) (*eclaType.Var, bool) {
	v, ok := env.Vars.Get(name)
	return v, ok
}

func (env *Env) NewScope(Type ScopeType) {
	env.Vars.GoDeep(Type)
}

func (env *Env) EndScope() {
	env.Vars.GoUp()
}

// SetFunction sets the function with the given name.
func (env *Env) SetFunction(name string, f *eclaType.Function) {
	v, err := eclaType.NewVar(name, f.GetType(), f)
	if err != nil {
		panic(err)
	}
	env.Vars.Set(name, v)
}

// GetFunction returns the function with the given name.
func (env *Env) GetFunction(name string) (*eclaType.Function, bool) {
	f, ok := env.Vars.Get(name)
	if !ok {
		return nil, false
	}
	if f.IsFunction() {
		fn := f.GetFunction()
		if fn == nil {
			panic("function is nil")
		}
		return fn, true
	}
	return nil, false
}

// Execute executes Env.Code or Env.File.
func (env *Env) Execute() {
	if env.File != "" {
		env.Code = readFile(env.File)
	}
	// Lexing
	env.Tokens = lexer.Lexer(env.Code)

	// Parsing
	pars := parser.Parser{Tokens: env.Tokens, ErrorHandler: env.ErrorHandle}
	env.SyntaxTree = pars.Parse()

	// Execute
	Run(env)
}

func (env *Env) Import(file string) {
	env.Libs[file] = libs.Import(file)
}

// readFile reads the file at the given path and returns its contents as a string.
func readFile(file string) string {
	v, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(v)
}
