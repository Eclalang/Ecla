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
	file := &File{
		Dependencies: []string{"a", "b", "c"},
		Imports:      []string{"a", "b", "c"},
	}
	if ok, _ := file.DepChecker(); !ok {
		t.Error("DepChecker failed to detect that all dependencies are satisfied")
	}
	f2 := &File{
		Dependencies: []string{"a", "b", "c"},
		Imports:      []string{"a", "b"},
	}
	if ok, _ := f2.DepChecker(); ok {
		t.Error("DepChecker failed to detect that all dependencies are not satisfied")
	}
	// test remove StructInstances from the Dependencies
	f3 := &File{
		Dependencies:    []string{"a", "b", "c"},
		Imports:         []string{"c"},
		StructInstances: []string{"a", "b"},
	}
	if ok, _ := f3.DepChecker(); !ok {
		t.Error("DepChecker failed to detect that all dependencies are satisfied")
	}
}

func TestFile_AddDependency(t *testing.T) {
	file := &File{
		Dependencies: []string{"b", "c"},
		Imports:      []string{"a", "b", "c"},
	}
	if err := file.AddDependency("a"); err != nil {
		t.Error("AddDependency failed to add a new dependency")
	}
	err := file.AddDependency("a")
	if err != nil {
		t.Error("AddDependency failed to detect that the dependency already exists")
	}
	if len(file.Dependencies) != 3 {
		t.Error("AddDependency failed to detect that the dependency already exists")
	}
}

func TestFile_RemoveDependency(t *testing.T) {
	file := &File{
		Dependencies: []string{"a", "b", "c"},
	}
	file.RemoveDependency("a")
	if contains("a", file.Dependencies) {
		t.Error("RemoveDependency failed to remove the dependency")
	}
	file.RemoveDependency("d")
	if len(file.Dependencies) != 2 {
		t.Error("RemoveDependency failed to detect that the dependency does not exist")
	}

}

func TestFile_AddImport(t *testing.T) {
	file := &File{
		Imports: []string{"a", "b", "c"},
	}
	file.AddImport("a")
	if len(file.Imports) != 3 {
		t.Error("AddImport failed to detect that the import already exists")
	}
	file.AddImport("d")

	if len(file.Imports) != 4 {
		t.Error("AddImport failed to add a new import")
	}
}

func TestFile_IsImported(t *testing.T) {
	file := &File{
		Imports: []string{"a", "b", "c"},
	}
	if !file.IsImported("a") {
		t.Error("IsImported failed to detect that the import exists")
	}
	if file.IsImported("d") {
		t.Error("IsImported failed to detect that the import does not exist")
	}
}

func TestFile_ConsumeComments(t *testing.T) {
	file := &File{}
	tokens := []lexer.Token{
		{TokenType: lexer.COMMENT, Value: "comment"},
		{TokenType: lexer.COMMENTGROUP, Value: "commentgroup"},
		{TokenType: lexer.TEXT, Value: "text"},
	}
	tokens = file.ConsumeComments(tokens)
	if len(tokens) > 1 {
		t.Error("ConsumeComments failed to consume the comments")
	}
}

func TestGetPackageNameByPath(t *testing.T) {
	val := GetPackageNameByPath("/Eclalang/Ecla/lexer")
	if val != "lexer" {
		t.Error("GetPackageNameByPath failed to get the package name")
	}
	val = GetPackageNameByPath("/Eclalang/Ecla/Test/Demo/Release/utils.ecla")
	if val != "utils" {
		t.Error("GetPackageNameByPath failed to get the package name")
	}
	val = GetPackageNameByPath("")
	if val != "" {
		t.Error("GetPackageNameByPath failed to get the package name")
	}

}
