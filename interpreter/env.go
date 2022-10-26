package interpreter

import (
	"os"
	"runtime"

	"github.com/tot0p/Ecla/interpreter/eclaType"
)

// Env is the environment in which the code is executed.
type Env struct {
	Vars       map[string]eclaType.Type
	OS         string
	SyntaxTree any
	File       string
	Code       string
}

// NewEnv returns a new Env.
func NewEnv() *Env {
	return &Env{
		OS:   runtime.GOOS,
		Vars: make(map[string]eclaType.Type),
	}
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
	//TODO: lexer
	//TODO: parse code
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
