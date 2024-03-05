package parser

import (
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var a_func AnonymousFunctionExpr = AnonymousFunctionExpr{
	FunctionToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "function",
		Position:  0,
		Line:      0,
	},
	Prototype: FunctionPrototype{
		LeftParamParen: lexer.Token{
			TokenType: lexer.LPAREN,
			Value:     "(",
			Position:  1,
			Line:      0,
		},
		RightParamParen: lexer.Token{
			TokenType: lexer.RPAREN,
			Value:     ")",
			Position:  2,
			Line:      0,
		},
		Parameters: []FunctionParams{},
		LeftRetsParen: lexer.Token{
			TokenType: lexer.LPAREN,
			Value:     "(",
			Position:  3,
			Line:      0,
		},
		RightRetsParen: lexer.Token{
			TokenType: lexer.RPAREN,
			Value:     ")",
			Position:  4,
			Line:      0,
		},
		ReturnTypes: []string{},
		LeftBrace: lexer.Token{
			TokenType: lexer.LBRACE,
			Value:     "{",
			Position:  5,
			Line:      0,
		},
		RightBrace: lexer.Token{
			TokenType: lexer.RBRACE,
			Value:     "}",
			Position:  6,
			Line:      3,
		},
	},
	Body: []Node{},
}

var a_call AnonymousFunctionCallExpr = AnonymousFunctionCallExpr{
	AnonymousFunction: a_func,
	LeftParen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  0,
		Line:      0,
	},
	RightParen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  6,
		Line:      3,
	},
	Args: []Expr{},
}

func TestAnonymousFunctionCallExpr_StartPos(t *testing.T) {
	if a_call.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_EndPos(t *testing.T) {
	if a_call.EndPos() != 6 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_StartLine(t *testing.T) {
	if a_call.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_EndLine(t *testing.T) {
	if a_call.EndLine() != 3 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_precedence(t *testing.T) {
	if a_call.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_exprNode(t *testing.T) {
	a_call.exprNode()
}

func TestAnonymousFunctionExpr_StartPos(t *testing.T) {
	if a_func.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_EndPos(t *testing.T) {
	if a_func.EndPos() != 6 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_StartLine(t *testing.T) {
	if a_func.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_EndLine(t *testing.T) {
	if a_func.EndLine() != 3 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_precedence(t *testing.T) {
	if a_func.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_exprNode(t *testing.T) {
	a_func.exprNode()
}

var arr_Literal ArrayLiteral = ArrayLiteral{
	LBRACKET: lexer.Token{
		TokenType: lexer.LBRACKET,
		Value:     "[",
		Position:  1,
		Line:      1,
	},
	RBRACKET: lexer.Token{
		TokenType: lexer.RBRACKET,
		Value:     "]",
		Position:  2,
		Line:      1,
	},
	Values: []Expr{},
}

func TestArrayLiteral_StartPos(t *testing.T) {
	if arr_Literal.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestArrayLiteral_EndPos(t *testing.T) {
	if arr_Literal.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestArrayLiteral_StartLine(t *testing.T) {
	if arr_Literal.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestArrayLiteral_EndLine(t *testing.T) {
	if arr_Literal.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestArrayLiteral_precedence(t *testing.T) {
	if arr_Literal.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestArrayLiteral_exprNode(t *testing.T) {
	arr_Literal.exprNode()
}

var b_expr BinaryExpr = BinaryExpr{
	LeftExpr: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "a",
			Position:  0,
			Line:      0,
		},
		Type:  "VAR",
		Value: "a",
	},
	Operator: lexer.Token{
		TokenType: lexer.ADD,
		Value:     "+",
		Position:  1,
		Line:      0,
	},
	RightExpr: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "b",
			Position:  2,
			Line:      0,
		},
	},
}

func TestBinaryExpr_StartPos(t *testing.T) {
	if b_expr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestBinaryExpr_EndPos(t *testing.T) {
	if b_expr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestBinaryExpr_StartLine(t *testing.T) {
	if b_expr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestBinaryExpr_EndLine(t *testing.T) {
	if b_expr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestBinaryExpr_precedence(t *testing.T) {
	if b_expr.precedence() != TokenPrecedence(b_expr.Operator) {
		t.Error("precedence failed to return the correct value")
	}
}

func TestBinaryExpr_exprNode(t *testing.T) {
	b_expr.exprNode()
}

var f_call FunctionCallExpr = FunctionCallExpr{
	FunctionCallToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "f",
		Position:  0,
		Line:      0,
	},
	Name: "f",
	LeftParen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  1,
		Line:      0,
	},
	RightParen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  2,
		Line:      0,
	},
	Args: []Expr{},
}

func TestFunctionCallExpr_StartPos(t *testing.T) {
	if f_call.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestFunctionCallExpr_EndPos(t *testing.T) {
	if f_call.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestFunctionCallExpr_StartLine(t *testing.T) {
	if f_call.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestFunctionCallExpr_EndLine(t *testing.T) {
	if f_call.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestFunctionCallExpr_precedence(t *testing.T) {
	if f_call.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestFunctionCallExpr_exprNode(t *testing.T) {
	f_call.exprNode()
}

var i_access IndexableAccessExpr = IndexableAccessExpr{
	VariableToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "a",
		Position:  0,
		Line:      0,
	},
	VariableName: "a",
	Indexes: []Expr{
		Literal{
			Token: lexer.Token{
				TokenType: lexer.TEXT,
				Value:     "0",
				Position:  1,
				Line:      0,
			},
			Type:  "INT",
			Value: "0",
		},
	},
	LastBracket: lexer.Token{
		TokenType: lexer.RBRACKET,
		Value:     "]",
		Position:  2,
		Line:      0,
	},
}

func TestIndexableAccessExpr_StartPos(t *testing.T) {
	if i_access.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestIndexableAccessExpr_EndPos(t *testing.T) {
	if i_access.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestIndexableAccessExpr_StartLine(t *testing.T) {
	if i_access.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestIndexableAccessExpr_EndLine(t *testing.T) {
	if i_access.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestIndexableAccessExpr_precedence(t *testing.T) {
	if i_access.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestIndexableAccessExpr_exprNode(t *testing.T) {
	i_access.exprNode()
}

var l Literal = Literal{
	Token: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "a",
		Position:  0,
		Line:      0,
	},
	Type:  "VAR",
	Value: "a",
}

func TestLiteral_StartPos(t *testing.T) {
	if l.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestLiteral_EndPos(t *testing.T) {
	if l.EndPos() != 0 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestLiteral_StartLine(t *testing.T) {
	if l.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestLiteral_EndLine(t *testing.T) {
	if l.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestLiteral_precedence(t *testing.T) {
	if l.precedence() != TokenPrecedence(l.Token) {
		t.Error("precedence failed to return the correct value")
	}
}

func TestLiteral_exprNode(t *testing.T) {
	l.exprNode()
}

var m_literal MapLiteral = MapLiteral{
	LBRACE: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  0,
		Line:      0,
	},
	Keys: []Expr{
		Literal{
			Token: lexer.Token{
				TokenType: lexer.TEXT,
				Value:     "a",
				Position:  1,
				Line:      0,
			},
			Type:  "VAR",
			Value: "a",
		},
	},
	Values: []Expr{
		Literal{
			Token: lexer.Token{
				TokenType: lexer.TEXT,
				Value:     "b",
				Position:  3,
				Line:      0,
			},
			Type:  "VAR",
			Value: "b",
		},
	},
	RBRACE: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  4,
		Line:      0,
	},
}

func TestMapLiteral_StartPos(t *testing.T) {
	if m_literal.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestMapLiteral_EndPos(t *testing.T) {
	if m_literal.EndPos() != 4 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestMapLiteral_StartLine(t *testing.T) {
	if m_literal.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestMapLiteral_EndLine(t *testing.T) {
	if m_literal.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestMapLiteral_precedence(t *testing.T) {
	if m_literal.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestMapLiteral_exprNode(t *testing.T) {
	m_literal.exprNode()
}

var p_expr ParenExpr = ParenExpr{
	Lparen: lexer.Token{
		TokenType: lexer.LPAREN,
		Value:     "(",
		Position:  0,
		Line:      0,
	},
	Expression: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "a",
			Position:  1,
			Line:      0,
		},
		Type:  "VAR",
		Value: "a",
	},
	Rparen: lexer.Token{
		TokenType: lexer.RPAREN,
		Value:     ")",
		Position:  2,
		Line:      0,
	},
}

func TestParenExpr_StartPos(t *testing.T) {
	if p_expr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestParenExpr_EndPos(t *testing.T) {
	if p_expr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestParenExpr_StartLine(t *testing.T) {
	if p_expr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestParenExpr_EndLine(t *testing.T) {
	if p_expr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestParenExpr_precedence(t *testing.T) {
	if p_expr.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestParenExpr_exprNode(t *testing.T) {
	p_expr.exprNode()
}

var s_expr SelectorExpr = SelectorExpr{
	Field: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "a",
		Position:  0,
		Line:      0,
	},
	Expr: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "b",
			Position:  1,
			Line:      0,
		},
		Type:  "VAR",
		Value: "b",
	},
	Sel: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "c",
			Position:  2,
			Line:      0,
		},
		Type:  "VAR",
		Value: "c",
	},
}

func TestSelectorExpr_StartPos(t *testing.T) {
	if s_expr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestSelectorExpr_EndPos(t *testing.T) {
	if s_expr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestSelectorExpr_StartLine(t *testing.T) {
	if s_expr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestSelectorExpr_EndLine(t *testing.T) {
	if s_expr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestSelectorExpr_precedence(t *testing.T) {
	if s_expr.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestSelectorExpr_exprNode(t *testing.T) {
	s_expr.exprNode()
}

var s_instance StructInstantiationExpr = StructInstantiationExpr{
	StructNameToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "a",
		Position:  0,
		Line:      0,
	},
	Name: "a",
	LeftBrace: lexer.Token{
		TokenType: lexer.LBRACE,
		Value:     "{",
		Position:  1,
		Line:      0,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  2,
		Line:      0,
	},
	Args: []Expr{},
}

func TestStructInstantiationExpr_StartPos(t *testing.T) {
	if s_instance.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestStructInstantiationExpr_EndPos(t *testing.T) {
	if s_instance.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestStructInstantiationExpr_StartLine(t *testing.T) {
	if s_instance.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestStructInstantiationExpr_EndLine(t *testing.T) {
	if s_instance.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestStructInstantiationExpr_precedence(t *testing.T) {
	if s_instance.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestStructInstantiationExpr_exprNode(t *testing.T) {
	s_instance.exprNode()
}

var u_expr UnaryExpr = UnaryExpr{
	Operator: lexer.Token{
		TokenType: lexer.SUB,
		Value:     "-",
		Position:  0,
		Line:      0,
	},
	RightExpr: Literal{
		Token: lexer.Token{
			TokenType: lexer.TEXT,
			Value:     "a",
			Position:  1,
			Line:      0,
		},
		Type:  "VAR",
		Value: "a",
	},
}

func TestUnaryExpr_StartPos(t *testing.T) {
	if u_expr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestUnaryExpr_EndPos(t *testing.T) {
	if u_expr.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestUnaryExpr_StartLine(t *testing.T) {
	if u_expr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestUnaryExpr_EndLine(t *testing.T) {
	if u_expr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestUnaryExpr_precedence(t *testing.T) {
	if u_expr.precedence() != TokenPrecedence(u_expr.Operator) {
		t.Error("precedence failed to return the correct value")
	}
}

func TestUnaryExpr_exprNode(t *testing.T) {
	u_expr.exprNode()
}
