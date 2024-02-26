package interpreter

import (
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Eclalang/Ecla/errorHandler"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
	met "github.com/Eclalang/Ecla/metrics"
	"github.com/Eclalang/Ecla/parser"
	libs "github.com/Eclalang/LibraryController"
)

func InitBuildIn() *Scope {
	vars := NewScopeMain()
	v, err := eclaType.NewVar("typeof", "", eclaType.NewTypeOf())
	if err != nil {
		panic(err)
	}
	vars.Set("typeOf", v)
	v, err = eclaType.NewVar("sizeof", "", eclaType.NewSizeOf())
	if err != nil {
		panic(err)
	}
	vars.Set("sizeOf", v)
	v, err = eclaType.NewVar("len", "", eclaType.NewLen())
	if err != nil {
		panic(err)
	}
	vars.Set("len", v)
	v, err = eclaType.NewVar("append", "", eclaType.NewAppend())
	if err != nil {
		panic(err)
	}
	vars.Set("append", v)
	return vars
}

// Env is the environment in which the code is executed.
type Env struct {
	Vars         *Scope
	OS           string
	ARCH         string
	SyntaxTree   *parser.File
	Tokens       []lexer.Token
	File         string
	Code         string
	Libs         map[string]libs.Lib
	ErrorHandle  *errorHandler.ErrorHandler
	ExecutedFunc []*eclaType.Function
	TypeDecl     []eclaDecl.TypeDecl
}

// NewEnv returns a new Env.
func NewEnv() *Env {
	return &Env{
		OS:           runtime.GOOS,
		ARCH:         runtime.GOARCH,
		Vars:         InitBuildIn(),
		Libs:         make(map[string]libs.Lib),
		ErrorHandle:  errorHandler.NewHandler(),
		ExecutedFunc: []*eclaType.Function{},
	}
}

// NewTemporaryEnv returns a new temporary Env.
// Temporary Env are used to import modules.
func NewTemporaryEnv(ErrorHandler *errorHandler.ErrorHandler) *Env {
	return &Env{
		OS:           runtime.GOOS,
		ARCH:         runtime.GOARCH,
		Vars:         InitBuildIn(),
		Libs:         make(map[string]libs.Lib),
		ErrorHandle:  ErrorHandler,
		ExecutedFunc: []*eclaType.Function{},
	}
}

func (env *Env) String() string {
	return fmt.Sprintf("Env{OS: %s, ARCH: %s , CODE: %s }", env.OS, env.ARCH, env.Code)
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

// CheckIfVarExistsInCurrentScope returns true if the variable exists in the current scope.
func (env *Env) CheckIfVarExistsInCurrentScope(name string) bool {
	return env.Vars.CheckIfVarExistsInCurrentScope(name)
}

// NewScope creates a new scope.
func (env *Env) NewScope(Type ScopeType) {
	env.Vars.GoDeep(Type)
}

// SetScope sets the most deep scope.
func (env *Env) SetScope(s *Scope) {
	env.Vars.GoDeepWithSpecificScope(s)
}

// EndScope ends the current scope.
func (env *Env) EndScope() {
	env.Vars.GoUp()
}

// SetFunction sets the function with the given name.
func (env *Env) SetFunction(name string, f *eclaType.Function) {
	v, err := eclaType.NewVar(name, f.GetType(), f)
	if err != nil {
		env.ErrorHandle.HandleError(0, 0, err.Error(), errorHandler.LevelFatal)
	}
	env.Vars.Set(name, v)
}

// Execute executes Env.Code or Env.File.
func (env *Env) Execute() {
	// catch all panics
	defer func() {
		if r := recover(); r != nil {
			env.ErrorHandle.HandleError(0, 0,
				fmt.Sprintf("an internal error occured please report it to the developers on https://github.com/Eclalang/Ecla/issues : %v", r),
				errorHandler.LevelFatal)
		}
	}()

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

func (env *Env) ExecuteMetrics() met.Metrics {
	if env.File != "" {
		env.Code = readFile(env.File)
	}
	m := met.NewMetrics()
	m.StartTimers()
	// Lexing
	m.StartLexerTimer()
	env.Tokens = lexer.Lexer(env.Code)
	m.StopLexerTimer()

	// Parsing
	m.StartParserTimer()
	pars := parser.Parser{Tokens: env.Tokens, ErrorHandler: env.ErrorHandle}
	env.SyntaxTree = pars.Parse()
	m.StopParserTimer()

	// Execute
	m.StartInterpreterTimer()
	Run(env)
	m.StopInterpreterTimer()
	m.StopTotalTimer()
	return *m
}

// Load the file
func (env *Env) Load() {
	env.Code = readFile(env.File)
	// Lexing
	env.Tokens = lexer.Lexer(env.Code)

	// Parsing
	pars := parser.Parser{Tokens: env.Tokens, ErrorHandler: env.ErrorHandle}
	env.SyntaxTree = pars.Parse()

	Load(env)
}

// Import executes an import statement.
func (env *Env) Import(stmt parser.ImportStmt) {
	file := stmt.ModulePath
	temp := libs.Import(file)
	if temp == nil {
		if !filepath.IsAbs(file) {
			file = filepath.Join(filepath.Dir(env.File), stmt.ModulePath)
		}
		if _, err := os.Stat(file); os.IsNotExist(err) {
			env.ErrorHandle.HandleError(stmt.ImportToken.Line, 0, fmt.Sprintf("module '%s' not found", file), errorHandler.LevelFatal)
		} else if err != nil {
			env.ErrorHandle.HandleError(stmt.ImportToken.Line, 0, err.Error(), errorHandler.LevelFatal)
		}
		tempsEnv := NewTemporaryEnv(env.ErrorHandle)
		tempsEnv.SetFile(file)
		tempsEnv.Load()

		temp = tempsEnv.ConvertToLib(env)
	}
	name := parser.GetPackageNameByPath(file)
	env.Libs[name] = temp
	v, err := eclaType.NewVar(name, "", eclaType.NewLib(name))
	if err != nil {
		env.ErrorHandle.HandleError(stmt.ImportToken.Line, 0, err.Error(), errorHandler.LevelFatal)
	}
	env.Vars.Set(name, v)
}

// AddFunctionExecuted adds a function to the pile of executed functions.
func (env *Env) AddFunctionExecuted(f *eclaType.Function) {
	env.ExecutedFunc = append(env.ExecutedFunc, f)
}

// GetFunctionExecuted returns the last function executed.
func (env *Env) GetFunctionExecuted() *eclaType.Function {
	return env.ExecutedFunc[len(env.ExecutedFunc)-1]
}

// RemoveFunctionExecuted removes the last function executed.
func (env *Env) RemoveFunctionExecuted() {
	env.ExecutedFunc = env.ExecutedFunc[:len(env.ExecutedFunc)-1]
}

// envLib represents a library and that uses to compartiment the scope of the library and the scope of the main program.
type envLib struct {
	Var  *Scope
	Libs map[string]libs.Lib
	env  *Env
}

// Call calls the function with the given name and arguments.
func (lib *envLib) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	function, ok := lib.Var.Get(name)
	if !ok {
		lib.env.ErrorHandle.HandleError(0, 0, fmt.Sprintf("function '%s' not found", name), errorHandler.LevelFatal)
	}
	if !function.IsFunction() {
		lib.env.ErrorHandle.HandleError(0, 0, fmt.Sprintf("'%s' is not a function", name), errorHandler.LevelFatal)
	}
	f := function.GetFunction()
	if f == nil {
		lib.env.ErrorHandle.HandleError(0, 0, fmt.Sprintf("function '%s' is nil", name), errorHandler.LevelFatal)
	}

	// TODO : Change this to more clean code

	// Save the current libs
	temps := lib.env.Libs
	// Set the libs of the lib
	lib.env.Libs = lib.Libs
	// Run the function
	r1, r2 := RunFunctionCallExprWithArgs(name, lib.env, f, args)
	// Restore the libs
	lib.env.Libs = temps
	return r1, r2
}

// ConvertToLib converts the Env to a Lib.
func (env *Env) ConvertToLib(MainEnv *Env) libs.Lib {
	return &envLib{
		Var:  env.Vars,
		Libs: env.Libs,
		env:  MainEnv,
	}
}

func (env *Env) AddTypeDecl(t eclaDecl.TypeDecl) {
	env.TypeDecl = append(env.TypeDecl, t)
}

func (env *Env) GetTypeDecl(name string) (eclaDecl.TypeDecl, bool) {
	for _, t := range env.TypeDecl {
		if t.GetName() == name {
			return t, true
		}
	}
	return nil, false
}

// readFile reads the file at the given path and returns its contents as a string.
func readFile(file string) string {
	v, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(v)
}
