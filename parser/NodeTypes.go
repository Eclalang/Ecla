package parser

import (
	"github.com/Eclalang/Ecla/lexer"
)

// Literal is a struct that defines a literal value for all types
type Literal struct {
	Token lexer.Token
	Type  string
	Value string
}

func (l Literal) StartPos() int {
	return l.Token.Position
}

func (l Literal) EndPos() int {
	return l.Token.Position
}

func (l Literal) StartLine() int {
	return l.Token.Line
}

func (l Literal) EndLine() int {
	return l.Token.Line
}

func (l Literal) precedence() int {
	return TokenPrecedence(l.Token)
}

func (l Literal) exprNode() {}

// BinaryExpr is a struct that defines a binary operation between two expressions
type BinaryExpr struct {
	LeftExpr  Expr
	Operator  lexer.Token
	RightExpr Expr
}

func (b BinaryExpr) StartPos() int {
	return b.LeftExpr.StartPos()
}

func (b BinaryExpr) EndPos() int {
	return b.RightExpr.EndPos()
}

func (b BinaryExpr) StartLine() int {
	return b.LeftExpr.StartLine()
}

func (b BinaryExpr) EndLine() int {
	return b.RightExpr.EndLine()
}

func (b BinaryExpr) precedence() int {
	return TokenPrecedence(b.Operator)
}

func (b BinaryExpr) exprNode() {}

// UnaryExpr is a struct that defines a unary operation on an expression
type UnaryExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (u UnaryExpr) StartPos() int {
	return u.Operator.Position
}

func (u UnaryExpr) EndPos() int {
	return u.RightExpr.EndPos()
}

func (u UnaryExpr) StartLine() int {
	return u.Operator.Line
}

func (u UnaryExpr) EndLine() int {
	return u.RightExpr.EndLine()
}

func (u UnaryExpr) precedence() int {
	return TokenPrecedence(u.Operator)
}

func (u UnaryExpr) exprNode() {}

// ParenExpr is a struct that defines a parenthesized expression
type ParenExpr struct {
	Lparen     lexer.Token
	Expression Expr
	Rparen     lexer.Token
}

func (p ParenExpr) StartPos() int {
	return p.Lparen.Position
}

func (p ParenExpr) EndPos() int {
	return p.Rparen.Position
}

func (p ParenExpr) StartLine() int {
	return p.Lparen.Line
}

func (p ParenExpr) EndLine() int {
	return p.Rparen.Line
}

func (p ParenExpr) precedence() int {
	return HighestPrecedence
}

func (p ParenExpr) exprNode() {}

type ArrayLiteral struct {
	LBRACKET lexer.Token
	Values   []Expr
	RBRACKET lexer.Token
}

func (p ArrayLiteral) StartPos() int {
	return p.LBRACKET.Position
}

func (p ArrayLiteral) EndPos() int {
	return p.RBRACKET.Position
}

func (p ArrayLiteral) StartLine() int {
	return p.LBRACKET.Line
}

func (p ArrayLiteral) EndLine() int {
	return p.RBRACKET.Line
}

func (p ArrayLiteral) precedence() int {
	return HighestPrecedence
}

func (p ArrayLiteral) exprNode() {}

type MapLiteral struct {
	LBRACE lexer.Token
	Keys   []Expr
	Values []Expr
	RBRACE lexer.Token
}

func (m MapLiteral) StartPos() int {
	return m.LBRACE.Position
}

func (m MapLiteral) EndPos() int {
	return m.RBRACE.Position
}

func (m MapLiteral) StartLine() int {
	return m.LBRACE.Line
}

func (m MapLiteral) EndLine() int {
	return m.RBRACE.Line
}

func (m MapLiteral) precedence() int {
	return HighestPrecedence
}

func (m MapLiteral) exprNode() {}

type IndexableAccessExpr struct {
	VariableToken lexer.Token
	VariableName  string
	Indexes       []Expr
	LastBracket   lexer.Token
}

func (a IndexableAccessExpr) StartPos() int {
	return a.VariableToken.Position
}

func (a IndexableAccessExpr) EndPos() int {
	return a.LastBracket.Position
}

func (a IndexableAccessExpr) StartLine() int {
	return a.VariableToken.Line
}

func (a IndexableAccessExpr) EndLine() int {
	return a.LastBracket.Line
}

func (a IndexableAccessExpr) precedence() int {
	return HighestPrecedence
}

func (a IndexableAccessExpr) exprNode() {}

type TypeStmt struct {
	TypeToken  lexer.Token
	Lparen     lexer.Token
	Rparen     lexer.Token
	Expression Expr
}

func (p TypeStmt) StartPos() int {
	return p.TypeToken.Position
}

func (p TypeStmt) EndPos() int {
	return p.Rparen.Position
}

func (p TypeStmt) StartLine() int {
	return p.TypeToken.Line
}

func (p TypeStmt) EndLine() int {
	return p.Rparen.Line
}

func (p TypeStmt) stmtNode() {}

type VariableDecl struct {
	VarToken lexer.Token
	Name     string
	Type     string
	Value    Expr
}

func (v VariableDecl) StartPos() int {
	return v.VarToken.Position
}

func (v VariableDecl) EndPos() int {
	return v.Value.EndPos()
}

func (v VariableDecl) StartLine() int {
	return v.VarToken.Line
}

func (v VariableDecl) EndLine() int {
	return v.Value.EndLine()
}

func (v VariableDecl) declNode() {}

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

type FunctionCallExpr struct {
	FunctionCallToken lexer.Token
	Name              string
	LeftParen         lexer.Token
	RightParen        lexer.Token
	Args              []Expr
}

func (f FunctionCallExpr) StartPos() int {
	return f.FunctionCallToken.Position
}

func (f FunctionCallExpr) EndPos() int {
	return f.RightParen.Position
}

func (f FunctionCallExpr) StartLine() int {
	return f.FunctionCallToken.Line
}

func (f FunctionCallExpr) EndLine() int {
	return f.RightParen.Line
}

func (f FunctionCallExpr) precedence() int {
	return HighestPrecedence
}

func (f FunctionCallExpr) exprNode() {}

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

type FunctionParams struct {
	Name string
	Type string
}

type FunctionPrototype struct {
	LeftParamParen  lexer.Token
	RightParamParen lexer.Token
	Parameters      []FunctionParams
	LeftRetsParen   lexer.Token
	RightRetsParen  lexer.Token
	ReturnTypes     []string
	LeftBrace       lexer.Token
	RightBrace      lexer.Token
}

type FunctionDecl struct {
	FunctionToken lexer.Token
	Name          string
	Prototype     FunctionPrototype
	Body          []Node
}

func (f FunctionDecl) StartPos() int {
	return f.FunctionToken.Position
}

func (f FunctionDecl) EndPos() int {
	return f.Prototype.RightBrace.Position
}

func (f FunctionDecl) StartLine() int {
	return f.FunctionToken.Line
}

func (f FunctionDecl) EndLine() int {
	return f.Prototype.RightBrace.Line
}

func (f FunctionDecl) declNode() {}

type AnonymousFunctionExpr struct {
	FunctionToken lexer.Token
	Prototype     FunctionPrototype
	Body          []Node
}

func (f AnonymousFunctionExpr) StartPos() int {
	return f.FunctionToken.Position
}

func (f AnonymousFunctionExpr) EndPos() int {
	return f.Prototype.RightBrace.Position
}

func (f AnonymousFunctionExpr) StartLine() int {
	return f.FunctionToken.Line
}

func (f AnonymousFunctionExpr) EndLine() int {
	return f.Prototype.RightBrace.Line
}

func (f AnonymousFunctionExpr) precedence() int {
	return HighestPrecedence
}

func (f AnonymousFunctionExpr) exprNode() {}

type AnonymousFunctionCallExpr struct {
	AnonymousFunction AnonymousFunctionExpr
	LeftParen         lexer.Token
	RightParen        lexer.Token
	Args              []Expr
}

func (f AnonymousFunctionCallExpr) StartPos() int {
	return f.AnonymousFunction.StartPos()
}

func (f AnonymousFunctionCallExpr) EndPos() int {
	return f.RightParen.Position
}

func (f AnonymousFunctionCallExpr) StartLine() int {
	return f.AnonymousFunction.StartLine()
}

func (f AnonymousFunctionCallExpr) EndLine() int {
	return f.RightParen.Line
}

func (f AnonymousFunctionCallExpr) precedence() int {
	return HighestPrecedence
}

func (f AnonymousFunctionCallExpr) exprNode() {}

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

type MurlocStmt struct {
	MurlocToken lexer.Token
}

func (m MurlocStmt) StartPos() int {
	return m.MurlocToken.Position
}

func (m MurlocStmt) EndPos() int {
	return m.MurlocToken.Position + len(lexer.MURLOC)
}

func (m MurlocStmt) StartLine() int {
	return m.MurlocToken.Line
}

func (m MurlocStmt) EndLine() int {
	return m.MurlocToken.Line
}

func (m MurlocStmt) stmtNode() {}

type StructField struct {
	Name string
	Type string
}

type StructDecl struct {
	StructToken lexer.Token
	Name        string
	LeftBrace   lexer.Token
	Fields      []StructField
	RightBrace  lexer.Token
}

func (s StructDecl) StartPos() int {
	return s.StructToken.Position
}

func (s StructDecl) EndPos() int {
	return s.RightBrace.Position
}

func (s StructDecl) StartLine() int {
	return s.StructToken.Line
}

func (s StructDecl) EndLine() int {
	return s.RightBrace.Line
}

func (s StructDecl) declNode() {}

type StructInstantiationExpr struct {
	StructNameToken lexer.Token
	Name            string
	LeftBrace       lexer.Token
	RightBrace      lexer.Token
	Args            []Expr
}

func (s StructInstantiationExpr) StartPos() int {
	return s.StructNameToken.Position
}

func (s StructInstantiationExpr) EndPos() int {
	return s.RightBrace.Position
}

func (s StructInstantiationExpr) StartLine() int {
	return s.StructNameToken.Line
}

func (s StructInstantiationExpr) EndLine() int {
	return s.RightBrace.Line
}

func (s StructInstantiationExpr) precedence() int {
	return HighestPrecedence
}

func (s StructInstantiationExpr) exprNode() {}

type SelectorExpr struct {
	Field lexer.Token
	Expr  Expr
	Sel   Expr
}

func (s SelectorExpr) StartPos() int {
	return s.Field.Position
}

func (s SelectorExpr) EndPos() int {
	return s.Sel.EndPos()
}

func (s SelectorExpr) StartLine() int {
	return s.Field.Line
}

func (s SelectorExpr) EndLine() int {
	return s.Sel.EndLine()
}

func (s SelectorExpr) precedence() int {
	return HighestPrecedence
}

func (s SelectorExpr) exprNode() {}

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
