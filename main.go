package main

import (
	"flag"
	"fmt"
	"github.com/tot0p/Ecla/interpreter"
)

var (
	Debug = false
)

func init() {
	flag.BoolVar(&Debug, "debug", Debug, "enable debug mode")
	flag.BoolVar(&Debug, "d", Debug, "enable debug mode (shorthand)")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		println("ecla: no input files")
		return
	}
	fmt.Println("//--- RUN", args[0], "---")
	Env := interpreter.NewEnv()
	Env.SetFile(args[0])
	Env.Execute()
	// Run the interpreter
}
