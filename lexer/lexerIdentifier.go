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
		Syntaxe:    []string{},
	},
	{
		Identifier: INT,
		Syntaxe: []string{
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
		Syntaxe: []string{
			"+",
		},
	},
	{
		Identifier: SUB,
		Syntaxe: []string{
			"-",
		},
	},
	{
		Identifier: MULT,
		Syntaxe: []string{
			"*",
		},
	},
	{
		Identifier: DIV,
		Syntaxe: []string{
			"/",
		},
	},
	{
		Identifier: MOD,
		Syntaxe: []string{
			"%",
		},
	},
	{
		Identifier: ASSIGN,
		Syntaxe: []string{
			"=",
		},
	},
	{
		Identifier: GTR,
		Syntaxe: []string{
			">",
		},
	},
	{
		Identifier: LSS,
		Syntaxe: []string{
			"<",
		},
	},
	{
		Identifier: XOR,
		Syntaxe: []string{
			"^",
		},
	},
	{
		Identifier: AND,
		Syntaxe: []string{
			"&&",
		},
	},
	{
		Identifier: OR,
		Syntaxe: []string{
			"||",
		},
	},
	{
		Identifier: LPAREN,
		Syntaxe: []string{
			"(",
		},
	},
	{
		Identifier: RPAREN,
		Syntaxe: []string{
			")",
		},
	},
	{
		Identifier: EOL,
		Syntaxe: []string{
			";",
		},
	},
	{
		Identifier: DQUOTE,
		Syntaxe: []string{
			"\"",
		},
	},
	{
		Identifier: MURLOC,
		Syntaxe: []string{
			"mgrlgrl",
		},
	},
	{
		Identifier: PERIOD,
		Syntaxe: []string{
			".",
		},
	},
	{
		Identifier: COLON,
		Syntaxe: []string{
			":",
		},
	},
	{
		Identifier: COMMA,
		Syntaxe: []string{
			",",
		},
	},
	{
		Identifier: LBRACE,
		Syntaxe: []string{
			"{",
		},
	},
	{
		Identifier: RBRACE,
		Syntaxe: []string{
			"}",
		},
	},
	{
		Identifier: LBRACKET,
		Syntaxe: []string{
			"[",
		},
	},
	{
		Identifier: RBRACKET,
		Syntaxe: []string{
			"]",
		},
	},
	{
		Identifier: COMMENT,
		Syntaxe: []string{
			"#",
		},
	},
	{
		Identifier: "",
		Syntaxe: []string{
			" ",
			"\n",
			"\t",
			"\r",
		},
	},
	{
		Identifier: BOOL,
		Syntaxe: []string{
			"true",
			"false",
		},
	},
	{
		Identifier: NOT,
		Syntaxe: []string{
			"!",
		},
	},
	{
		Identifier: EOF,
		Syntaxe:    []string{},
	},
}

type identifier struct {
	Identifier string
	Syntaxe    []string
}

// IsSyntaxe verify is the string tempVal exist in the current identifier
//
// return a true false value
func (ident identifier) IsSyntaxe(tempVal string) bool {
	for _, syntaxe := range ident.Syntaxe {
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
