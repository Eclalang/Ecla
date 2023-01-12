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
				TokenType: lexer.INC,
				Value:     `+=`,
				Position:  1,
				Line:      0,
			},
			{
				TokenType: lexer.DEC,
				Value:     `-=`,
				Position:  4,
				Line:      0,
			},
			{
				TokenType: lexer.DIV,
				Value:     `/=`,
				Position:  7,
				Line:      0,
			},
			{
				TokenType: lexer.MULT,
				Value:     `*=`,
				Position:  10,
				Line:      0,
			},
			{
				TokenType: lexer.MOD,
				Value:     `%=`,
				Position:  13,
				Line:      0,
			},
			{
				TokenType: lexer.ADD,
				Value:     `+`,
				Position:  16,
				Line:      0,
			},
			{
				TokenType: lexer.ADD,
				Value:     `+`,
				Position:  17,
				Line:      0,
			},
			{
				TokenType: lexer.SUB,
				Value:     `-`,
				Position:  19,
				Line:      0,
			},
			{
				TokenType: lexer.SUB,
				Value:     `-`,
				Position:  20,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `a`,
				Position:  22,
				Line:      0,
			},
			{
				TokenType: lexer.INC,
				Value:     `++`,
				Position:  23,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `b`,
				Position:  26,
				Line:      0,
			},
			{
				TokenType: lexer.DEC,
				Value:     `--`,
				Position:  27,
				Line:      0,
			},
			{
				TokenType: lexer.QOT,
				Value:     `//=`,
				Position:  30,
				Line:      0,
			},
			{
				TokenType: lexer.EOF,
				Value:     "",
				Position:  33,
				Line:      0,
			},
		},
	}
	testDQuote = testList{
		input: `"a\"b"a"be"`,
		output: []lexer.Token{
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      0,
			},
			{
				TokenType: lexer.STRING,
				Value:     `a\"b`,
				Position:  2,
				Line:      0,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `a`,
				Position:  7,
				Line:      0,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  8,
				Line:      0,
			},
			{
				TokenType: lexer.STRING,
				Value:     `be`,
				Position:  9,
				Line:      0,
			},
			{
				TokenType: lexer.DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      0,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  12,
				Line:      0,
			},
		},
	}
	testMurloc = testList{
		input: `shrek is love, mgrlgrl is life`,
		output: []lexer.Token{
			{
				TokenType: lexer.TEXT,
				Value:     `shrek`,
				Position:  0,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `is`,
				Position:  6,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `love`,
				Position:  9,
				Line:      0,
			},
			{
				TokenType: lexer.COMMA,
				Value:     `,`,
				Position:  13,
				Line:      0,
			},
			{
				TokenType: lexer.MURLOC,
				Value:     `mgrlgrl`,
				Position:  15,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `is`,
				Position:  23,
				Line:      0,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `life`,
				Position:  26,
				Line:      0,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  30,
				Line:      0,
			},
		},
	}
	testSpeChar = testList{
		input: ":;.,()[]{}",
		output: []lexer.Token{
			{
				TokenType: lexer.COLON,
				Value:     `:`,
				Position:  1,
				Line:      0,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  2,
				Line:      0,
			},
			{
				TokenType: lexer.PERIOD,
				Value:     `.`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.COMMA,
				Value:     `,`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: lexer.LPAREN,
				Value:     `(`,
				Position:  3,
				Line:      1,
			},
			{
				TokenType: lexer.RPAREN,
				Value:     `)`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: lexer.LBRACKET,
				Value:     `[`,
				Position:  5,
				Line:      1,
			},
			{
				TokenType: lexer.RBRACKET,
				Value:     `]`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: lexer.LBRACE,
				Value:     `{`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: lexer.RBRACE,
				Value:     `}`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  9,
				Line:      1,
			}},
	}
	testEOL = testList{
		input: "();mgrlgrl;_aa_",
		output: []lexer.Token{
			{
				TokenType: lexer.LPAREN,
				Value:     `(`,
				Position:  1,
				Line:      0,
			},
			{
				TokenType: lexer.RPAREN,
				Value:     `)`,
				Position:  2,
				Line:      0,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  3,
				Line:      0,
			},
			{
				TokenType: lexer.MURLOC,
				Value:     `mgrlgrl`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: lexer.EOL,
				Value:     `;`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: lexer.TEXT,
				Value:     `aa`,
				Position:  2,
				Line:      2,
			},
			{
				TokenType: lexer.EOF,
				Value:     ``,
				Position:  5,
				Line:      2,
			},
		},
	}
)
