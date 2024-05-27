package parser

import (
	"fmt"
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var bStmt = BlockScopeStmt{
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  1,
		Line:      1,
	},
	Body: []Node{
		aCall,
		aFunc,
	},
}

func TestBlockScopeStmt_StartPos(t *testing.T) {
	if bStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestBlockScopeStmt_EndPos(t *testing.T) {
	if bStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestBlockScopeStmt_StartLine(t *testing.T) {
	if bStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestBlockScopeStmt_EndLine(t *testing.T) {
	if bStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestBlockScopeStmt_stmtNode(t *testing.T) {
	bStmt.stmtNode()
}

var eStmt = ElseStmt{
	ElseToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "else",
		Position:  1,
		Line:      1,
	},
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  1,
		Line:      1,
	},
	Body: []Node{
		aCall,
		aFunc,
	},
	IfStmt: nil,
}

func TestElseStmt_StartPos(t *testing.T) {
	if eStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestElseStmt_EndPos(t *testing.T) {
	if eStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestElseStmt_StartLine(t *testing.T) {
	if eStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestElseStmt_EndLine(t *testing.T) {
	if eStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestElseStmt_stmtNode(t *testing.T) {
	eStmt.stmtNode()
}

var fStmt = ForStmt{
	ForToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "for",
		Position:  1,
		Line:      1,
	},
	LeftParen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  1,
		Line:      1,
	},
	RightParen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  1,
		Line:      1,
	},
	InitDecl:       nil,
	CondExpr:       nil,
	PostAssignStmt: nil,
	KeyToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "key",
		Position:  1,
		Line:      1,
	},
	ValueToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "value",
		Position:  1,
		Line:      1,
	},
	RangeToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "range",
		Position:  1,
		Line:      1,
	},
	RangeExpr: nil,
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  1,
		Line:      1,
	},
	Body: []Node{
		aCall,
		aFunc,
	},
}

func TestForStmt_StartPos(t *testing.T) {
	if fStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestForStmt_EndPos(t *testing.T) {
	if fStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestForStmt_StartLine(t *testing.T) {
	if fStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestForStmt_EndLine(t *testing.T) {
	if fStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestForStmt_stmtNode(t *testing.T) {
	fStmt.stmtNode()
}

var iStmt = IfStmt{
	IfToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "if",
		Position:  1,
		Line:      1,
	},
	LeftParen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  1,
		Line:      1,
	},
	RightParen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  1,
		Line:      1,
	},
	Cond: nil,
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  1,
		Line:      1,
	},
	Body: []Node{
		aCall,
		aFunc,
	},
	ElseStmt: &eStmt,
}

func TestIfStmt_StartPos(t *testing.T) {
	if iStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestIfStmt_EndPos(t *testing.T) {
	if iStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestIfStmt_StartLine(t *testing.T) {
	if iStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestIfStmt_EndLine(t *testing.T) {
	if iStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestIfStmt_stmtNode(t *testing.T) {
	iStmt.stmtNode()
}

var impStmt = ImportStmt{
	ImportToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "import",
		Position:  1,
		Line:      1,
	},
	ModulePath: "console",
}

func TestImportStmt_StartPos(t *testing.T) {
	if impStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestImportStmt_EndPos(t *testing.T) {
	if impStmt.EndPos() != 1+len(impStmt.ModulePath) {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestImportStmt_StartLine(t *testing.T) {
	if impStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestImportStmt_EndLine(t *testing.T) {
	if impStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestImportStmt_stmtNode(t *testing.T) {
	impStmt.stmtNode()
}

var mStmt = MurlocStmt{
	MurlocToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "murloc",
		Position:  1,
		Line:      1,
	},
}

func TestMurlocStmt_StartPos(t *testing.T) {
	if mStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestMurlocStmt_EndPos(t *testing.T) {
	if mStmt.EndPos() != 1+len(lexer.MURLOC) {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestMurlocStmt_StartLine(t *testing.T) {
	if mStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestMurlocStmt_EndLine(t *testing.T) {
	if mStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestMurlocStmt_stmtNode(t *testing.T) {
	mStmt.stmtNode()
}

var rStmt = ReturnStmt{
	ReturnToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "return",
		Position:  1,
		Line:      1,
	},
	ReturnValues: []Expr{
		aCall,
		aFunc,
	},
}

func TestReturnStmt_StartPos(t *testing.T) {
	if rStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestReturnStmt_EndPos(t *testing.T) {
	if rStmt.EndPos() != rStmt.ReturnValues[len(rStmt.ReturnValues)-1].EndPos() {
		t.Error("EndPos failed to return the correct value")
	}
	back := rStmt.ReturnValues
	rStmt.ReturnValues = nil
	if rStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
	rStmt.ReturnValues = back

}

func TestReturnStmt_StartLine(t *testing.T) {
	if rStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestReturnStmt_EndLine(t *testing.T) {
	fmt.Println(len(rStmt.ReturnValues))
	if rStmt.EndLine() != rStmt.ReturnValues[len(rStmt.ReturnValues)-1].EndLine() {
		t.Error("EndLine failed to return the correct value")
	}
	back := rStmt.ReturnValues
	rStmt.ReturnValues = nil
	if rStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
	rStmt.ReturnValues = back
}

func TestReturnStmt_stmtNode(t *testing.T) {
	rStmt.stmtNode()
}

var vAsStmt = VariableAssignStmt{
	VarToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "var",
		Position:  1,
		Line:      1,
	},
	Names: []Expr{
		aCall,
		aFunc,
	},
	Values: []Expr{
		aCall,
		aFunc,
	},
	Operator: "=",
}

func TestVariableAssignStmt_StartPos(t *testing.T) {
	if vAsStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestVariableAssignStmt_EndPos(t *testing.T) {
	if vAsStmt.EndPos() != vAsStmt.Values[len(vAsStmt.Values)-1].EndPos() {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestVariableAssignStmt_StartLine(t *testing.T) {
	if vAsStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestVariableAssignStmt_EndLine(t *testing.T) {
	if vAsStmt.EndLine() != vAsStmt.Values[len(vAsStmt.Values)-1].EndLine() {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestVariableAssignStmt_stmtNode(t *testing.T) {
	vAsStmt.stmtNode()
}

var wStmt = WhileStmt{
	WhileToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "while",
		Position:  1,
		Line:      1,
	},
	LeftParen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  1,
		Line:      1,
	},
	RightParen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  1,
		Line:      1,
	},
	Cond: nil,
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  1,
		Line:      1,
	},
	Body: []Node{
		aCall,
		aFunc,
	},
}

func TestWhileStmt_StartPos(t *testing.T) {
	if wStmt.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestWhileStmt_EndPos(t *testing.T) {
	if wStmt.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestWhileStmt_StartLine(t *testing.T) {
	if wStmt.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestWhileStmt_EndLine(t *testing.T) {
	if wStmt.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestWhileStmt_stmtNode(t *testing.T) {
	wStmt.stmtNode()
}
