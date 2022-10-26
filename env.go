package ecla

import (
	"os"
	"runtime"
)

type Env struct {
	Vars       map[string]any
	OS         string
	SyntaxTree any
	File       string
	Code       string
}

func CreateEnv() *Env {
	return &Env{
		OS:   runtime.GOOS,
		Vars: make(map[string]any),
	}
}

func (env *Env) SetCode(code string) {
	env.Code = code
}

func (env *Env) SetFile(file string) {
	env.File = file
}

func (env *Env) SetVar(name string, value any) {
	env.Vars[name] = value
}

func (env *Env) GetVar(name string) any {
	return env.Vars[name]
}

func (env *Env) Execute() {
	if env.File != "" {
		env.Code = readFile(env.File)
	}
	//TODO: lexer
	//TODO: parse code
	//TODO: execute code
}

func readFile(file string) string {
	v, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(v)
}
