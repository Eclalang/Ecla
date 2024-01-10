package main

import v2 "github.com/Eclalang/Ecla/lexer/v2"

func main() {
	tokens := v2.LexerR("\"prout\"ok")
	for i, token := range tokens {
		printToken(i, token)
	}
}

func printToken(i int, token v2.Token) {
	println("\n-------Token", i, "-------")
	println("Tokentype: ", token.TokenType)
	println("Value: ", token.Value)
	println("Line: ", token.Line)
	println("Position: ", token.Position)
}
