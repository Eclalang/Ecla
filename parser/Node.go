package parser

// Node is the interface that all nodes in the AST must implement
type Node interface {
	StartPos() int
	EndPos() int
}

// Expr is an expression
type Expr interface {
	Node
	precedence() int
	exprNode()
}

// Stmt is a statement
type Stmt interface {
	Node
	stmtNode()
}

// Decl is a declaration
type Decl interface {
	Node
	declNode()
}
