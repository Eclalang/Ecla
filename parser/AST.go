package parser

import (
	"github.com/tot0p/Ecla/lexer"
)

// The AST struct contains all the information needed for the interpreter to run
type AST struct {
	Operations []Node
}

type File struct {
	ParseTree *AST
	Imports   []string
	// TODO: use a map instead of a slice for better performance
	Dependencies     []string
	VariableDecl     []string
	FunctionDecl     []string
	ConsumedComments []string
	Trace            string
}

// DepChecker checks if all dependencies in the current file are resolved by the specified imports
func (f *File) DepChecker() (bool, []string) {
	var Unresolved []string
	for _, value := range f.Dependencies {
		if !contains(value, f.Imports) {
			Unresolved = append(Unresolved, value)
		}
	}
	if len(Unresolved) > 0 {
		return false, Unresolved
	}
	return true, Unresolved
}

// AddDependency adds a new dependency to the file that is currently being parsed
func (f *File) AddDependency(dep string) {
	if !contains(dep, f.Dependencies) {
		f.Dependencies = append(f.Dependencies, dep)
	}
}

// ConsumeComments examines all the tokens, consumes them and deletes them from the token slice
func (f *File) ConsumeComments(tokens []lexer.Token) []lexer.Token {
	var tempTokens []lexer.Token
	for _, token := range tokens {
		if token.TokenType == lexer.COMMENT || token.TokenType == lexer.COMMENTGROUP {
			f.ConsumedComments = append(f.ConsumedComments, token.Value)
		} else {
			tempTokens = append(tempTokens, token)
		}
	}
	return tempTokens
}

// contains checks if a string is in a slice of strings
func contains(needle string, haystack []string) bool {
	for _, matchValue := range haystack {
		if matchValue == needle {
			return true
		}
	}
	return false
}
