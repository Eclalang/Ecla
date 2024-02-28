package parser

import "github.com/Eclalang/Ecla/lexer"

type BlockScopeStmt struct {
	LeftBrace  lexer.Token
	RightBrace lexer.Token
	Body       []Node
}

func (b BlockScopeStmt) StartPos() int {
	return b.LeftBrace.Position
}

func (b BlockScopeStmt) EndPos() int {
	return b.RightBrace.Position
}

func (b BlockScopeStmt) StartLine() int {
	return b.LeftBrace.Line
}

func (b BlockScopeStmt) EndLine() int {
	return b.RightBrace.Line
}

func (b BlockScopeStmt) stmtNode() {}

type ElseStmt struct {
	ElseToken  lexer.Token
	LeftBrace  lexer.Token
	RightBrace lexer.Token
	Body       []Node
	IfStmt     *IfStmt
}

func (e ElseStmt) StartPos() int {
	return e.ElseToken.Position
}

func (e ElseStmt) EndPos() int {
	return e.RightBrace.Position
}

func (e ElseStmt) StartLine() int {
	return e.ElseToken.Line
}

func (e ElseStmt) EndLine() int {
	return e.RightBrace.Line
}

func (e ElseStmt) stmtNode() {}

type ForStmt struct {
	ForToken             lexer.Token
	LeftParen            lexer.Token
	RightParen           lexer.Token
	InitDecl             Decl
	CondExpr             Expr
	PostAssignStmt       Stmt
	KeyToken, ValueToken lexer.Token
	RangeToken           lexer.Token
	RangeExpr            Expr
	LeftBrace            lexer.Token
	RightBrace           lexer.Token
	Body                 []Node
}

func (f ForStmt) StartPos() int {
	return f.ForToken.Position
}

func (f ForStmt) EndPos() int {
	return f.RightBrace.Position
}

func (f ForStmt) StartLine() int {
	return f.ForToken.Line
}

func (f ForStmt) EndLine() int {
	return f.RightBrace.Line
}

func (f ForStmt) stmtNode() {}

type IfStmt struct {
	IfToken    lexer.Token
	LeftParen  lexer.Token
	RightParen lexer.Token
	Cond       Expr
	LeftBrace  lexer.Token
	RightBrace lexer.Token
	Body       []Node
	ElseStmt   *ElseStmt
}

func (i IfStmt) StartPos() int {
	return i.IfToken.Position
}

func (i IfStmt) EndPos() int {
	return i.RightBrace.Position
}

func (i IfStmt) StartLine() int {
	return i.IfToken.Line
}

func (i IfStmt) EndLine() int {
	return i.RightBrace.Line
}

func (i IfStmt) stmtNode() {}

type ImportStmt struct {
	ImportToken lexer.Token
	ModulePath  string
}

func (i ImportStmt) StartPos() int {
	return i.ImportToken.Position
}

func (i ImportStmt) EndPos() int {
	return i.ImportToken.Position + len(i.ModulePath)
}

func (i ImportStmt) StartLine() int {
	return i.ImportToken.Line
}

func (i ImportStmt) EndLine() int {
	return i.ImportToken.Line
}

func (i ImportStmt) stmtNode() {}

type ReturnStmt struct {
	ReturnToken  lexer.Token
	ReturnValues []Expr
}

func (r ReturnStmt) StartPos() int {
	return r.ReturnToken.Position
}

func (r ReturnStmt) EndPos() int {
	if len(r.ReturnValues) > 0 {
		return r.ReturnValues[len(r.ReturnValues)-1].EndPos()
	}
	return r.ReturnToken.Position
}

func (r ReturnStmt) StartLine() int {
	return r.ReturnToken.Line
}

func (r ReturnStmt) EndLine() int {
	if len(r.ReturnValues) > 0 {
		return r.ReturnValues[len(r.ReturnValues)-1].EndLine()
	}
	return r.ReturnToken.Line
}

func (r ReturnStmt) stmtNode() {}

type VariableAssignStmt struct {
	VarToken lexer.Token
	Names    []Expr
	Operator string
	Values   []Expr
}

func (v VariableAssignStmt) StartPos() int {
	return v.VarToken.Position
}

func (v VariableAssignStmt) EndPos() int {
	return v.Values[len(v.Values)-1].EndPos()
}

func (v VariableAssignStmt) StartLine() int {
	return v.VarToken.Line
}

func (v VariableAssignStmt) EndLine() int {
	return v.Values[len(v.Values)-1].EndLine()
}

func (v VariableAssignStmt) stmtNode() {}

type WhileStmt struct {
	WhileToken lexer.Token
	LeftParen  lexer.Token
	RightParen lexer.Token
	Cond       Expr
	LeftBrace  lexer.Token
	RightBrace lexer.Token
	Body       []Node
}

func (w WhileStmt) StartPos() int {
	return w.WhileToken.Position
}

func (w WhileStmt) EndPos() int {
	return w.RightBrace.Position
}

func (w WhileStmt) StartLine() int {
	return w.WhileToken.Line
}

func (w WhileStmt) EndLine() int {
	return w.RightBrace.Line
}

func (w WhileStmt) stmtNode() {}
