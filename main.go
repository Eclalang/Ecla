package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tot0p/Ecla/interpreter"
	"strings"
	"time"
)

var (
	Debug       = false
	Time        = false
	TimeExec    = time.Now()
	lexerDebug  = false
	parserDebug = false
)

func init() {
	flag.BoolVar(&Debug, "debug", Debug, "enable debug mode")
	flag.BoolVar(&Debug, "d", Debug, "enable debug mode (shorthand)")
	flag.BoolVar(&Time, "time", Time, "enable time mode")
	flag.BoolVar(&Time, "t", Time, "enable time mode (shorthand)")
	flag.BoolVar(&lexerDebug, "debugLex", lexerDebug, "enable lexer debug mod")
	flag.BoolVar(&lexerDebug, "dl", lexerDebug, "enable lexer debug mod")
	flag.BoolVar(&parserDebug, "debugParser", parserDebug, "enable parser debug")
	flag.BoolVar(&parserDebug, "dp", parserDebug, "enable parser debug")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("invalid input")
		return
	}
	fmt.Println("//--- RUN", args[0], "---")
	Env := interpreter.NewEnv()
	if t := strings.Split(args[0], "."); t[len(t)-1] == "ecla" || t[len(t)-1] == "eclaw" {
		Env.SetFile(args[0])
	} else if args[0][len(args[0])-1] == ';' {
		Env.SetCode(args[0])
	} else {
		fmt.Print("Ecla: invalid input file")
		return
	}
	Env.Execute()
	fmt.Println("\n//--- END", args[0], "---")
	if Debug {
		fmt.Println("ENV:", Env)
	}
	if Time {
		fmt.Println("TIME EXEC:", time.Since(TimeExec))
	}
	if lexerDebug {
		fmt.Println("Lexer Token:", Env.Tokens)
	}
	if parserDebug {
		txt, _ := json.MarshalIndent(Env.SyntaxTree, "", "  ")
		fmt.Println("SYNTAX TREE:", string(txt))
	}
}
