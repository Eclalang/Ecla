package lexer

// each type of token
var (
	TEXT              = "TEXT"
	STRING            = "STRING"
	PRINT             = "PRINT"
	INT               = "INT"
	FLOAT             = "FLOAT"
	ADD               = "ADD"
	SUB               = "SUB"
	MULT              = "MULT"
	DIV               = "DIV"
	QOT               = "QOT"
	MOD               = "MOD"
	INC               = "INC"
	DEC               = "DEC"
	ASSIGN            = "ASSIGN"
	LSS               = "LSS"
	GTR               = "GTR"
	NEQ               = "NEQ"
	LEQ               = "LEQ"
	GEQ               = "GEQ"
	XOR               = "XOR"
	OR                = "OR"
	AND               = "AND"
	EQUAL             = "EQUAL"
	LPAREN            = "LPAREN"
	RPAREN            = "RPAREN"
	EOL               = "EOL"
	DQUOTE            = "DQUOTE"
	PERIOD            = "PERIOD"
	COLON             = "COLON"
	LBRACE            = "LBRACE"
	RBRACE            = "RBRACE"
	LBRACKET          = "LBRACKET"
	RBRACKET          = "RBRACKET"
	COMMA             = "COMMA"
	NOT               = "NOT"
	BOOL              = "BOOL"
	MURLOC            = "MURLOC"
	EOF               = "EOF"
	COMMENT           = "COMMENT"
	COMMENTGROUP      = "COMMENTGROUP"
	COMMENTGROUPIDENT = "COMMENTGROUPIDENT"
)

//--------------------------------------------//
// Need to be optimized
// Change for a switch case of a hashmap
//--------------------------------------------//

// link between syntax and token
var Identifier []identifier = []identifier{
	{
		Identifier: TEXT,
		Syntax:    []string{},
	},
	{
		Identifier: INT,
		Syntax: []string{
			"0",
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
		},
	},
	{
		Identifier: ADD,
		Syntax: []string{
			"+",
		},
	},
	{
		Identifier: SUB,
		Syntax: []string{
			"-",
		},
	},
	{
		Identifier: MULT,
		Syntax: []string{
			"*",
		},
	},
	{
		Identifier: DIV,
		Syntax: []string{
			"/",
		},
	},
	{
		Identifier: MOD,
		Syntax: []string{
			"%",
		},
	},
	{
		Identifier: ASSIGN,
		Syntax: []string{
			"=",
		},
	},
	{
		Identifier: GTR,
		Syntax: []string{
			">",
		},
	},
	{
		Identifier: LSS,
		Syntax: []string{
			"<",
		},
	},
	{
		Identifier: XOR,
		Syntax: []string{
			"^",
		},
	},
	{
		Identifier: AND,
		Syntax: []string{
			"&&",
		},
	},
	{
		Identifier: OR,
		Syntax: []string{
			"||",
		},
	},
	{
		Identifier: LPAREN,
		Syntax: []string{
			"(",
		},
	},
	{
		Identifier: RPAREN,
		Syntax: []string{
			")",
		},
	},
	{
		Identifier: EOL,
		Syntax: []string{
			";",
		},
	},
	{
		Identifier: DQUOTE,
		Syntax: []string{
			"\"",
		},
	},
	{
		Identifier: MURLOC,
		Syntax: []string{
			"mgrlgrl",
		},
	},
	{
		Identifier: PERIOD,
		Syntax: []string{
			".",
		},
	},
	{
		Identifier: COLON,
		Syntax: []string{
			":",
		},
	},
	{
		Identifier: COMMA,
		Syntax: []string{
			",",
		},
	},
	{
		Identifier: LBRACE,
		Syntax: []string{
			"{",
		},
	},
	{
		Identifier: RBRACE,
		Syntax: []string{
			"}",
		},
	},
	{
		Identifier: LBRACKET,
		Syntax: []string{
			"[",
		},
	},
	{
		Identifier: RBRACKET,
		Syntax: []string{
			"]",
		},
	},
	{
		Identifier: COMMENT,
		Syntax: []string{
			"#",
		},
	},
	{
		Identifier: "",
		Syntax: []string{
			" ",
			"\n",
			"\t",
			"\r",
		},
	},
	{
		Identifier: BOOL,
		Syntax: []string{
			"true",
			"false",
		},
	},
	{
		Identifier: NOT,
		Syntax: []string{
			"!",
		},
	},
	{
		Identifier: EOF,
		Syntax:    []string{},
	},
}

type identifier struct {
	Identifier string
	Syntax    []string
}

// IsSyntaxe verify is the string tempVal exist in the current identifier
//
// return a true false value
func (ident identifier) IsSyntaxe(tempVal string) bool {
	for _, syntaxe := range ident.Syntax {
		if syntaxe == tempVal {
			return true
		}
	}
	return false
}

// concatEqual verify is the string str is equal with one of the
// mathematical opperand
//
// return a true false value
func concatEqual(str string) bool {
	switch str {
	case ADD:
		return true
	case MOD:
		return true
	case SUB:
		return true
	case QOT:
		return true
	case MULT:
		return true
	case ASSIGN:
		return true
	case DIV:
		return true
	case GTR:
		return true
	case LSS:
		return true
	case NOT:
		return true
	default:
		return false
	}
}
