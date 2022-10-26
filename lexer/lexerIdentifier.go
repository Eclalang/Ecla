package lexer

var Identifier []identifier = []identifier{
	{
		Identifier: "PRINT",
		Syntaxe: []string{
			"print",
		},
	},
	{
		Identifier: "INT",
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
		Identifier: "OPPERATOR",
		Syntaxe: []string{
			"+",
			"=",
			"/",
			"//",
			"%",
			"-",
			"*",
		},
	},
	{
		Identifier: "OPARENT",
		Syntaxe: []string{
			"(",
		},
	},
	{
		Identifier: "CPARENT",
		Syntaxe: []string{
			")",
		},
	},
	{
		Identifier: "EOL",
		Syntaxe: []string{
			";",
		},
	},
	{
		Identifier: "EOF",
		Syntaxe:    []string{},
	},
}

type identifier struct {
	Identifier string
	Syntaxe    []string
}
