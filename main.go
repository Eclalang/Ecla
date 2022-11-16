package main

import (
	"flag"
	"fmt"
	"github.com/tot0p/Ecla/interpreter"
	"strings"
	"time"
)

var (
	Debug    = false
	Time     = false
	TimeExec = time.Now()
)

func init() {
	flag.BoolVar(&Debug, "debug", Debug, "enable debug mode")
	flag.BoolVar(&Debug, "d", Debug, "enable debug mode (shorthand)")
	flag.BoolVar(&Time, "time", Time, "enable time mode")
	flag.BoolVar(&Time, "t", Time, "enable time mode (shorthand)")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		println("ecla: no input")
		return
	}
	fmt.Println("//--- RUN", args[0], "---")
	Env := interpreter.NewEnv()
	if t := strings.Split(args[0], "."); t[len(t)-1] == "ecla" || t[len(t)-1] == "eclaw" {
		Env.SetFile(args[0])
	} else if args[0][len(args[0])-1] == ';' {
		Env.SetCode(args[0])
	} else {
		println("ecla: invalid input file")
		return
	}
	Env.Execute()
	fmt.Println("//--- END", args[0], "---")
	if Debug {
		fmt.Println("ENV:", Env)
	}
	if Time {
		fmt.Println("TIME EXEC:", time.Since(TimeExec))
	}
	/*
		tokenList := []lexer.Token{
			lexer.Token{TokenType: lexer.LPAREN, Value: "(", Position: 0, Line: 0},
			lexer.Token{TokenType: lexer.INT, Value: "8", Position: 1, Line: 0},
			lexer.Token{TokenType: lexer.MULT, Value: "*", Position: 2, Line: 0},
			lexer.Token{TokenType: lexer.INT, Value: "5", Position: 3, Line: 0},
			lexer.Token{TokenType: lexer.RPAREN, Value: ")", Position: 4, Line: 0},
			lexer.Token{TokenType: lexer.ADD, Value: "+", Position: 5, Line: 0},
			lexer.Token{TokenType: lexer.INT, Value: "2", Position: 6, Line: 0},
		}
		parser := parser.Parser{Tokens: tokenList}
		text, _ := json.MarshalIndent(parser.Parse(), "", "  ")
		fmt.Println(string(text))
	*/
	// Run the interpreter
}
