package lexer

import (
	"github.com/tot0p/Ecla/lexer"
)

type testList struct {
	input  string
	output []lexer.Token
}

var (
	testCalc = testList{
		input: `+= -= /= *= %= ++ -- a++ b-- //=`,
		output: []lexer.Token{
			{
				TokenType: lexer.ADD + lexer.ASSIGN,
				Value:     `+=`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.SUB + lexer.ASSIGN,
				Value:     `-=`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: lexer.DIV + lexer.ASSIGN,
				Value:     `/=`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: lexer.MULT + lexer.ASSIGN,
				Value:     `*=`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: lexer.MOD + lexer.ASSIGN,
				Value:     `%=`,
				Position:  13,
				Line:      1,
			},
			{
				TokenType: lexer.ADD,
				Value:     `+`,
				Position:  16,
				Line:      1,
			},
			{
				TokenType: lexer.ADD,
				Value:     `+`,
				Position:  17,
				Line:      1,
			},
			{
				TokenType: lexer.SUB,
				Value:     `-`,
				Position:  19,
				Line:      1,
			},
			{
				TokenType: lexer.SUB,
				Value:     `-`,
				Position:  20,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `a`,
				Position:  22,
				Line:      1,
			},
			{
				TokenType: lexer.INC,
				Value:     `++`,
				Position:  23,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `b`,
				Position:  26,
				Line:      1,
			},
			{
				TokenType: lexer.DEC,
				Value:     `--`,
				Position:  27,
				Line:      1,
			},
			{
				TokenType: lexer.QOT + lexer.ASSIGN,
				Value:     `//=`,
				Position:  30,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     "",
				Position:  33,
				Line:      1,
			},
		},
	}
	testDQuote = testList{
		input: `"    "a"be"`,
		output: []lexer.Token{
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.STRING,
				Value:     `    `,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `a`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: lexer.STRING,
				Value:     `be`,
				Position:  9,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  12,
				Line:      1,
			},
		},
	}
	testMurloc = testList{
		input: `shrek is love, mgrlgrl is life`,
		output: []lexer.Token{
			{
				TokenType: lexer.TEXT,
				Value:     `shrek`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `is`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `love`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: lexer.COMMA,
				Value:     `,`,
				Position:  14,
				Line:      1,
			},
			{
				TokenType: lexer.MURLOC,
				Value:     `mgrlgrl`,
				Position:  16,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `is`,
				Position:  24,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `life`,
				Position:  27,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  31,
				Line:      1,
			},
		},
	}
	testSpeChar = testList{
		input: ":;\n.,()[]{}",
		output: []lexer.Token{
			{
				TokenType: lexer.COLON,
				Value:     `:`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.PERIOD,
				Value:     `.`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: lexer.COMMA,
				Value:     `,`,
				Position:  2,
				Line:      2,
			},
			{
				TokenType: lexer.LPAREN,
				Value:     `(`,
				Position:  3,
				Line:      2,
			},
			{
				TokenType: lexer.RPAREN,
				Value:     `)`,
				Position:  4,
				Line:      2,
			},
			{
				TokenType: lexer.LBRACKET,
				Value:     `[`,
				Position:  5,
				Line:      2,
			},
			{
				TokenType: lexer.RBRACKET,
				Value:     `]`,
				Position:  6,
				Line:      2,
			},
			{
				TokenType: lexer.LBRACE,
				Value:     `{`,
				Position:  7,
				Line:      2,
			},
			{
				TokenType: lexer.RBRACE,
				Value:     `}`,
				Position:  8,
				Line:      2,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  9,
				Line:      2,
			}},
	}
	testEOL = testList{
		input: "();\nmgrlgrl;\n_aa_",
		output: []lexer.Token{
			{
				TokenType: lexer.LPAREN,
				Value:     `(`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.RPAREN,
				Value:     `)`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  3,
				Line:      1,
			},
			{
				TokenType: lexer.MURLOC,
				Value:     `mgrlgrl`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  8,
				Line:      2,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `aa`,
				Position:  2,
				Line:      3,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  5,
				Line:      3,
			},
		},
	}
	testHashtag = testList{
		input: "prout# in comment\n#/ in commentgroup\n and next ligne\n and test for / and #/ /#\nOutside of the comment group",
		output: []lexer.Token{
			{
				TokenType: lexer.TEXT,
				Value:     "prout",
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.COMMENT,
				Value:     " in comment",
				Position:  6,
				Line:      1,
			},
			{
				TokenType: lexer.COMMENTGROUP,
				Value:     " in commentgroup\n and next ligne\n and test for / and #/ ",
				Position:  1,
				Line:      2,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `Outside`,
				Position:  1,
				Line:      5,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `of`,
				Position:  9,
				Line:      5,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `the`,
				Position:  12,
				Line:      5,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `comment`,
				Position:  16,
				Line:      5,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `group`,
				Position:  24,
				Line:      5,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  29,
				Line:      5,
			},
		},
	}
	testHashtag2 = testList{
		input: "prout# in comment\n#/ in commentgroup\n and next ligne\n and test for / and #/",
		output: []lexer.Token{
			{
				TokenType: lexer.TEXT,
				Value:     `prout`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.COMMENT,
				Value:     ` in comment`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: lexer.COMMENTGROUP,
				Value:     " in commentgroup\n and next ligne\n and test for / and #/",
				Position:  1,
				Line:      2,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  22,
				Line:      4,
			},
		},
	}
	testHashtag3 = testList{
		input: "\"#prout\"",
		output: []lexer.Token{
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.STRING,
				Value:     `#prout`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  9,
				Line:      1,
			},
		},
	}
	testHashtag4 = testList{
		input: "\"#/prout/#\"",
		output: []lexer.Token{
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.STRING,
				Value:     `#/prout/#`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  12,
				Line:      1,
			},
		},
	}
	testBoolOpperand = testList{
		input: "&& || ^ \"&& || ^\"",
		output: []lexer.Token{
			{
				TokenType: lexer.AND,
				Value:     `&&`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.OR,
				Value:     `||`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: lexer.XOR,
				Value:     `^`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  9,
				Line:      1,
			},
			{
				TokenType: lexer.STRING,
				Value:     `&& || ^`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  17,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  18,
				Line:      1,
			},
		},
	}
)
