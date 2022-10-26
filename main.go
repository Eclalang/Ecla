package main

import (
	"flag"
	"fmt"
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
	// Run the interpreter
}
