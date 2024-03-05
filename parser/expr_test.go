package parser

import (
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var aFunc = AnonymousFunctionExpr{
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

var aCall = AnonymousFunctionCallExpr{
	AnonymousFunction: aFunc,
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
	if aCall.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_EndPos(t *testing.T) {
	if aCall.EndPos() != 6 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_StartLine(t *testing.T) {
	if aCall.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_EndLine(t *testing.T) {
	if aCall.EndLine() != 3 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_precedence(t *testing.T) {
	if aCall.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestAnonymousFunctionCallExpr_exprNode(t *testing.T) {
	aCall.exprNode()
}

func TestAnonymousFunctionExpr_StartPos(t *testing.T) {
	if aFunc.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_EndPos(t *testing.T) {
	if aFunc.EndPos() != 6 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_StartLine(t *testing.T) {
	if aFunc.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_EndLine(t *testing.T) {
	if aFunc.EndLine() != 3 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_precedence(t *testing.T) {
	if aFunc.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestAnonymousFunctionExpr_exprNode(t *testing.T) {
	aFunc.exprNode()
}

var arrLiteral = ArrayLiteral{
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
	if arrLiteral.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestArrayLiteral_EndPos(t *testing.T) {
	if arrLiteral.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestArrayLiteral_StartLine(t *testing.T) {
	if arrLiteral.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestArrayLiteral_EndLine(t *testing.T) {
	if arrLiteral.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestArrayLiteral_precedence(t *testing.T) {
	if arrLiteral.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestArrayLiteral_exprNode(t *testing.T) {
	arrLiteral.exprNode()
}

var bExpr = BinaryExpr{
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
	if bExpr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestBinaryExpr_EndPos(t *testing.T) {
	if bExpr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestBinaryExpr_StartLine(t *testing.T) {
	if bExpr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestBinaryExpr_EndLine(t *testing.T) {
	if bExpr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestBinaryExpr_precedence(t *testing.T) {
	if bExpr.precedence() != TokenPrecedence(bExpr.Operator) {
		t.Error("precedence failed to return the correct value")
	}
}

func TestBinaryExpr_exprNode(t *testing.T) {
	bExpr.exprNode()
}

var fCall = FunctionCallExpr{
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
	if fCall.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestFunctionCallExpr_EndPos(t *testing.T) {
	if fCall.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestFunctionCallExpr_StartLine(t *testing.T) {
	if fCall.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestFunctionCallExpr_EndLine(t *testing.T) {
	if fCall.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestFunctionCallExpr_precedence(t *testing.T) {
	if fCall.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestFunctionCallExpr_exprNode(t *testing.T) {
	fCall.exprNode()
}

var iAccess = IndexableAccessExpr{
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
	if iAccess.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestIndexableAccessExpr_EndPos(t *testing.T) {
	if iAccess.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestIndexableAccessExpr_StartLine(t *testing.T) {
	if iAccess.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestIndexableAccessExpr_EndLine(t *testing.T) {
	if iAccess.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestIndexableAccessExpr_precedence(t *testing.T) {
	if iAccess.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestIndexableAccessExpr_exprNode(t *testing.T) {
	iAccess.exprNode()
}

var l = Literal{
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

var mLiteral = MapLiteral{
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
	if mLiteral.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestMapLiteral_EndPos(t *testing.T) {
	if mLiteral.EndPos() != 4 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestMapLiteral_StartLine(t *testing.T) {
	if mLiteral.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestMapLiteral_EndLine(t *testing.T) {
	if mLiteral.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestMapLiteral_precedence(t *testing.T) {
	if mLiteral.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestMapLiteral_exprNode(t *testing.T) {
	mLiteral.exprNode()
}

var pExpr = ParenExpr{
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
	if pExpr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestParenExpr_EndPos(t *testing.T) {
	if pExpr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestParenExpr_StartLine(t *testing.T) {
	if pExpr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestParenExpr_EndLine(t *testing.T) {
	if pExpr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestParenExpr_precedence(t *testing.T) {
	if pExpr.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestParenExpr_exprNode(t *testing.T) {
	pExpr.exprNode()
}

var sExpr = SelectorExpr{
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
	if sExpr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestSelectorExpr_EndPos(t *testing.T) {
	if sExpr.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestSelectorExpr_StartLine(t *testing.T) {
	if sExpr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestSelectorExpr_EndLine(t *testing.T) {
	if sExpr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestSelectorExpr_precedence(t *testing.T) {
	if sExpr.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestSelectorExpr_exprNode(t *testing.T) {
	sExpr.exprNode()
}

var sInstance = StructInstantiationExpr{
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
	if sInstance.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestStructInstantiationExpr_EndPos(t *testing.T) {
	if sInstance.EndPos() != 2 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestStructInstantiationExpr_StartLine(t *testing.T) {
	if sInstance.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestStructInstantiationExpr_EndLine(t *testing.T) {
	if sInstance.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestStructInstantiationExpr_precedence(t *testing.T) {
	if sInstance.precedence() != HighestPrecedence {
		t.Error("precedence failed to return the correct value")
	}
}

func TestStructInstantiationExpr_exprNode(t *testing.T) {
	sInstance.exprNode()
}

var uExpr = UnaryExpr{
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
	if uExpr.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestUnaryExpr_EndPos(t *testing.T) {
	if uExpr.EndPos() != 1 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestUnaryExpr_StartLine(t *testing.T) {
	if uExpr.StartLine() != 0 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestUnaryExpr_EndLine(t *testing.T) {
	if uExpr.EndLine() != 0 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestUnaryExpr_precedence(t *testing.T) {
	if uExpr.precedence() != TokenPrecedence(uExpr.Operator) {
		t.Error("precedence failed to return the correct value")
	}
}

func TestUnaryExpr_exprNode(t *testing.T) {
	uExpr.exprNode()
}
