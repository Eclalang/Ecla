package lexer

type testList struct {
	input  string
	output []Token
}

var (
	testCalc = testList{
		input: `+= -= /= *= %= ++ -- a++ b-- //=`,
		output: []Token{
			{
				TokenType: ADD + ASSIGN,
				Value:     `+=`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: SUB + ASSIGN,
				Value:     `-=`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: DIV + ASSIGN,
				Value:     `/=`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: MULT + ASSIGN,
				Value:     `*=`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: MOD + ASSIGN,
				Value:     `%=`,
				Position:  13,
				Line:      1,
			},
			{
				TokenType: INC,
				Value:     `++`,
				Position:  16,
				Line:      1,
			},
			{
				TokenType: DEC,
				Value:     `--`,
				Position:  19,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `a`,
				Position:  22,
				Line:      1,
			},
			{
				TokenType: INC,
				Value:     `++`,
				Position:  23,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `b`,
				Position:  26,
				Line:      1,
			},
			{
				TokenType: DEC,
				Value:     `--`,
				Position:  27,
				Line:      1,
			},
			{
				TokenType: QOT + ASSIGN,
				Value:     `//=`,
				Position:  30,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     "",
				Position:  33,
				Line:      1,
			},
		},
	}
	testDQuote = testList{
		input: `"    "a"be"`,
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `    `,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `a`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `be`,
				Position:  9,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  12,
				Line:      1,
			},
		},
	}
	testMurloc = testList{
		input: `shrek is love, mgrlgrl is life`,
		output: []Token{
			{
				TokenType: TEXT,
				Value:     `shrek`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `is`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `love`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: COMMA,
				Value:     `,`,
				Position:  14,
				Line:      1,
			},
			{
				TokenType: MURLOC,
				Value:     `mgrlgrl`,
				Position:  16,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `is`,
				Position:  24,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `life`,
				Position:  27,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  31,
				Line:      1,
			},
		},
	}
	testSpeChar = testList{
		input: ":;\n.,()[]{}",
		output: []Token{
			{
				TokenType: COLON,
				Value:     `:`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: EOL,
				Value:     `;`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: PERIOD,
				Value:     `.`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: COMMA,
				Value:     `,`,
				Position:  2,
				Line:      2,
			},
			{
				TokenType: LPAREN,
				Value:     `(`,
				Position:  3,
				Line:      2,
			},
			{
				TokenType: RPAREN,
				Value:     `)`,
				Position:  4,
				Line:      2,
			},
			{
				TokenType: LBRACKET,
				Value:     `[`,
				Position:  5,
				Line:      2,
			},
			{
				TokenType: RBRACKET,
				Value:     `]`,
				Position:  6,
				Line:      2,
			},
			{
				TokenType: LBRACE,
				Value:     `{`,
				Position:  7,
				Line:      2,
			},
			{
				TokenType: RBRACE,
				Value:     `}`,
				Position:  8,
				Line:      2,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  9,
				Line:      2,
			}},
	}
	testEOL = testList{
		input: "();\nmgrlgrl;\n_aa_",
		output: []Token{
			{
				TokenType: LPAREN,
				Value:     `(`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: RPAREN,
				Value:     `)`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: EOL,
				Value:     `;`,
				Position:  3,
				Line:      1,
			},
			{
				TokenType: MURLOC,
				Value:     `mgrlgrl`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: EOL,
				Value:     `;`,
				Position:  8,
				Line:      2,
			},
			{
				TokenType: TEXT,
				Value:     `_aa_`,
				Position:  1,
				Line:      3,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  5,
				Line:      3,
			},
		},
	}
	testNoFile = testList{
		input: "",
		output: []Token{
			{
				TokenType: EOF,
				Value:     ``,
				Position:  1,
				Line:      1,
			},
		},
	}
	testHashtag = testList{
		input: "prout# in comment\n#/ in commentgroup\n and next ligne\n and test for / and #/ /#\nOutside of the comment group",
		output: []Token{
			{
				TokenType: TEXT,
				Value:     "prout",
				Position:  1,
				Line:      1,
			},
			{
				TokenType: COMMENT,
				Value:     " in comment",
				Position:  6,
				Line:      1,
			},
			{
				TokenType: COMMENTGROUP,
				Value:     " in commentgroup\n and next ligne\n and test for / and #/ ",
				Position:  1,
				Line:      2,
			},
			{
				TokenType: TEXT,
				Value:     `Outside`,
				Position:  1,
				Line:      5,
			},
			{
				TokenType: TEXT,
				Value:     `of`,
				Position:  9,
				Line:      5,
			},
			{
				TokenType: TEXT,
				Value:     `the`,
				Position:  12,
				Line:      5,
			},
			{
				TokenType: TEXT,
				Value:     `comment`,
				Position:  16,
				Line:      5,
			},
			{
				TokenType: TEXT,
				Value:     `group`,
				Position:  24,
				Line:      5,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  29,
				Line:      5,
			},
		},
	}
	testHashtag2 = testList{
		input: "prout# in comment\n#/ in commentgroup\n and next ligne\n and test for / and #/",
		output: []Token{
			{
				TokenType: TEXT,
				Value:     `prout`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: COMMENT,
				Value:     ` in comment`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: COMMENTGROUP,
				Value:     " in commentgroup\n and next ligne\n and test for / and #/",
				Position:  1,
				Line:      2,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  22,
				Line:      4,
			},
		},
	}
	testHashtag3 = testList{
		input: "\"#prout\"",
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `#prout`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  8,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  9,
				Line:      1,
			},
		},
	}
	testHashtag4 = testList{
		input: "\"#/prout/#\"",
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `#/prout/#`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  12,
				Line:      1,
			},
		},
	}
	testHashtag5 = testList{
		input: "#ok\nimport console;\n\n\n#test Print\nconsole.println(\"Test Print\");\nconsole.println(\"Hello World\");\nconsole.println(\"Test Print Past\");",
		output: []Token{
			{
				TokenType: COMMENTGROUP,
				Value:     `#prout`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  11,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  12,
				Line:      1,
			},
		},
	}
	testBoolOpperand = testList{
		input: "&& || ^ \"&& || ^\"",
		output: []Token{
			{
				TokenType: AND,
				Value:     `&&`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: OR,
				Value:     `||`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: XOR,
				Value:     `^`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  9,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `&& || ^`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  17,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  18,
				Line:      1,
			},
		},
	}
)
