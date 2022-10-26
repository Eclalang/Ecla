package lexer

// each type of token
var (
	TEXT      = "TEXT"
	PRINT     = "PRINT"
	INT       = "INT"
	FLOAT     = "FLOAT"
	OPPERATOR = "OPPERATOR"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	EOL       = "EOL"
	EOF       = "EOF"
)

// link between syntaxe and token
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
		Identifier: OPPERATOR,
		Syntaxe: []string{
			"+",
			"/",
			"//",
			"%",
			"-",
			"*",
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
