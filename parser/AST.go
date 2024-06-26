package parser

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/Eclalang/Ecla/lexer"
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
	StructInstances  []string
	FunctionDecl     []string
	ConsumedComments []string
	Trace            string
}

// DepChecker checks if all dependencies in the current file are resolved by the specified imports
func (f *File) DepChecker() (bool, []string) {
	// remove StructInstances from the Dependencies
	for _, value := range f.StructInstances {
		f.RemoveDependency(value)
	}

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
func (f *File) AddDependency(dep string) error {
	// check if the dependency is satisfied by the imports
	if !contains(dep, f.Dependencies) {
		f.Dependencies = append(f.Dependencies, dep)
	}
	return nil
}

func (f *File) RemoveDependency(dep string) {
	if contains(dep, f.Dependencies) {
		index := slices.Index(f.Dependencies, dep)
		f.Dependencies = append(f.Dependencies[:index], f.Dependencies[index+1:]...)
	}
}

func (f *File) AddImport(imp string) {
	imp = GetPackageNameByPath(imp)
	if !contains(imp, f.Imports) {
		f.Imports = append(f.Imports, imp)
	}
}

func (f *File) IsImported(imp string) bool {
	return contains(imp, f.Imports)
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

func GetPackageNameByPath(path string) string {
	_, fPath := filepath.Split(path)
	// since spe in split function is not empty the function cannot return a slice of length 0
	temp := strings.Split(fPath, ".")
	return temp[0]
}
