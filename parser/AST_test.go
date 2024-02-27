package parser

import (
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !contains("a", slice) {
		t.Error("contains failed to detect that the slice contains the element")
	}
	if contains("d", slice) {
		t.Error("contains failed to detect that the slice does not contain the element")
	}
}

func TestFile_DepChecker(t *testing.T) {
	f := &File{
		Dependencies: []string{"a", "b", "c"},
		Imports:      []string{"a", "b", "c"},
	}
	if ok, _ := f.DepChecker(); !ok {
		t.Error("DepChecker failed to detect that all dependencies are satisfied")
	}
	f2 := &File{
		Dependencies: []string{"a", "b", "c"},
		Imports:      []string{"a", "b"},
	}
	if ok, _ := f2.DepChecker(); ok {
		t.Error("DepChecker failed to detect that all dependencies are not satisfied")
	}
}

func TestFile_AddDependency(t *testing.T) {
	f := &File{
		Dependencies: []string{"b", "c"},
		Imports:      []string{"a", "b", "c"},
	}
	if err := f.AddDependency("a"); err != nil {
		t.Error("AddDependency failed to add a new dependency")
	}
	if err := f.AddDependency("d"); err == nil {
		t.Error("AddDependency failed to detect that the dependency is not satisfied by previous imports")
	}
	err := f.AddDependency("a")
	if err != nil {
		t.Error("AddDependency failed to detect that the dependency already exists")
	}
	if len(f.Dependencies) != 3 {
		t.Error("AddDependency failed to detect that the dependency already exists")
	}
}

func TestFile_AddImport(t *testing.T) {
	f := &File{
		Imports: []string{"a", "b", "c"},
	}
	err := f.AddImport("a")
	if err != nil {
		t.Error("AddImport failed to detect that the import already exists")
	}
	if len(f.Imports) != 3 {
		t.Error("AddImport failed to detect that the import already exists")
	}
	err = f.AddImport("d")
	if err != nil {
		t.Error("AddImport failed to add a new import")
	}
	if len(f.Imports) != 4 {
		t.Error("AddImport failed to add a new import")
	}
	err = f.AddImport("")
	if err == nil {
		t.Error("AddImport failed to detect that the import is empty")
	}
	if len(f.Imports) != 4 {
		t.Error("AddImport failed to detect that the import is empty")
	}
}

func TestFile_IsImported(t *testing.T) {
	f := &File{
		Imports: []string{"a", "b", "c"},
	}
	if !f.IsImported("a") {
		t.Error("IsImported failed to detect that the import exists")
	}
	if f.IsImported("d") {
		t.Error("IsImported failed to detect that the import does not exist")
	}
}

func TestFile_ConsumeComments(t *testing.T) {
	f := &File{}
	tokens := []lexer.Token{
		{TokenType: lexer.COMMENT, Value: "comment"},
		{TokenType: lexer.COMMENTGROUP, Value: "commentgroup"},
		{TokenType: lexer.TEXT, Value: "text"},
	}
	tokens = f.ConsumeComments(tokens)
	if len(tokens) > 1 {
		t.Error("ConsumeComments failed to consume the comments")
	}
}

func TestGetPackageNameByPath(t *testing.T) {
	val, err := GetPackageNameByPath("/Eclalang/Ecla/lexer")
	if val != "lexer" || err != nil {
		t.Error("GetPackageNameByPath failed to get the package name")
	}
	val, err = GetPackageNameByPath("/Eclalang/Ecla/Test/Demo/Release/utils.ecla")
	if val != "utils" || err != nil {
		t.Error("GetPackageNameByPath failed to get the package name")
	}
	val, err = GetPackageNameByPath("")
	if val != "" || err == nil {
		t.Error("GetPackageNameByPath failed to get the package name")
	}

}
