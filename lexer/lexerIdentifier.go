package lexer

// each type of token
var (
	TEXT   = "TEXT"
	PRINT  = "PRINT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	ADD    = "ADD"
	SUB    = "SUB"
	MULT   = "MULT"
	DIV    = "DIV"
	QOT    = "QOT"
	MOD    = "MOD"
	INC    = "INC"
	DEC    = "DEC"
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	EOL    = "EOL"
	EOF    = "EOF"
)

// link between syntax and token
var Identifier []identifier = []identifier{
	{
		Identifier: TEXT,
		Syntaxe:    []string{},
	},
	{
		Identifier: PRINT,
		Syntaxe: []string{
			"print",
		},
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
		Identifier: QOT,
		Syntaxe: []string{
			"//",
		},
	},
	{
		Identifier: MOD,
		Syntaxe: []string{
			"%",
		},
	},
	{
		Identifier: INC,
		Syntaxe: []string{
			"++",
		},
	},
	{
		Identifier: DEC,
		Syntaxe: []string{
			"--",
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
		Identifier: EOF,
		Syntaxe:    []string{},
	},
}

type identifier struct {
	Identifier string
	Syntaxe    []string
}

// IsSyntaxe verify is the string tempVal exist in the current identifier
func (ident identifier) IsSyntaxe(tempVal string) bool {
	for _, syntaxe := range ident.Syntaxe {
		if syntaxe == tempVal {
			return true
		}
	}
	return false
}
