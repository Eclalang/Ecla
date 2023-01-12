package parser

import (
	"github.com/tot0p/Ecla/lexer"
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

func (p ArrayLiteral) precedence() int {
	return HighestPrecedence
}

func (p ArrayLiteral) exprNode() {}

type PrintStmt struct {
	PrintToken lexer.Token
	Lparen     lexer.Token
	Rparen     lexer.Token
	Expression Expr
}

func (p PrintStmt) StartPos() int {
	return p.PrintToken.Position
}

func (p PrintStmt) EndPos() int {
	return p.Rparen.Position
}

func (p PrintStmt) stmtNode() {}

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

func (v VariableDecl) declNode() {}

type VariableAssignStmt struct {
	VarToken lexer.Token
	Name     string
	Value    Expr
}

func (v VariableAssignStmt) StartPos() int {
	return v.VarToken.Position
}

func (v VariableAssignStmt) EndPos() int {
	return v.Value.EndPos()
}

func (v VariableAssignStmt) stmtNode() {}

type VariableIncrementStmt struct {
	VarToken lexer.Token
	Name     string
	IncToken lexer.Token
}

func (v VariableIncrementStmt) StartPos() int {
	return v.VarToken.Position
}

func (v VariableIncrementStmt) EndPos() int {
	return v.IncToken.Position
}

func (v VariableIncrementStmt) stmtNode() {}

type VariableDecrementStmt struct {
	VarToken lexer.Token
	Name     string
	DecToken lexer.Token
}

func (v VariableDecrementStmt) StartPos() int {
	return v.VarToken.Position
}

func (v VariableDecrementStmt) EndPos() int {
	return v.DecToken.Position
}

func (v VariableDecrementStmt) stmtNode() {}

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

func (f FunctionCallExpr) precedence() int {
	return HighestPrecedence
}

func (f FunctionCallExpr) exprNode() {}

type MethodCallExpr struct {
	MethodCallToken lexer.Token
	ObjectName      string
	FunctionCall    FunctionCallExpr
}

func (m MethodCallExpr) StartPos() int {
	return m.MethodCallToken.Position
}

func (m MethodCallExpr) EndPos() int {
	return m.FunctionCall.RightParen.Position
}

func (m MethodCallExpr) precedence() int {
	return HighestPrecedence
}

func (m MethodCallExpr) exprNode() {}

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

func (i ImportStmt) stmtNode() {}

type FunctionDecl struct {
	FunctionToken   lexer.Token
	Name            string
	LeftParamParen  lexer.Token
	RightParamParen lexer.Token
	Parameters      map[string]string
	LeftRetsParen   lexer.Token
	RightRetsParen  lexer.Token
	ReturnTypes     []string
	LeftBrace       lexer.Token
	RightBrace      lexer.Token
	Body            []Node
}

func (f FunctionDecl) StartPos() int {
	return f.FunctionToken.Position
}

func (f FunctionDecl) EndPos() int {
	return f.RightBrace.Position
}

func (f FunctionDecl) declNode() {}

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

func (r ReturnStmt) stmtNode() {}

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

func (a IndexableAccessExpr) precedence() int {
	return HighestPrecedence
}

func (a IndexableAccessExpr) exprNode() {}
