package lexer

type testList struct {
	input  string
	output []Token
}

var (
	testCalc = testList{
		input: `+= -= /= *= %= ++ -- a++ b-- //= ^^ ^`,
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
				TokenType: XOR,
				Value:     `^^`,
				Position:  34,
				Line:      1,
			},
			{
				TokenType: XORBIN,
				Value:     `^`,
				Position:  37,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     "",
				Position:  38,
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
	testCHAR = testList{
		input: "'Ok'",
		output: []Token{
			{
				TokenType: SQUOTE,
				Value:     `'`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: CHAR,
				Value:     `Ok`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: SQUOTE,
				Value:     `'`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  5,
				Line:      1,
			},
		},
	}
	testCHARSTRING = testList{
		input: "\"Ok'\"\"'Ok\"",
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `Ok'`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  5,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `'Ok`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  10,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  11,
				Line:      1,
			},
		},
	}
	testCHARSTRING2 = testList{
		input: `'Ok'"'Ok"`,
		output: []Token{
			{
				TokenType: SQUOTE,
				Value:     `'`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: CHAR,
				Value:     `Ok`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: SQUOTE,
				Value:     `'`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  5,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `'Ok`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  9,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  10,
				Line:      1,
			},
		},
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
	testHashtag0 = testList{
		input: "# in comment ",
		output: []Token{
			{
				TokenType: COMMENT,
				Value:     "# in comment ",
				Position:  1,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  14,
				Line:      1,
			},
		},
	}
	testHashtag1 = testList{
		input: "#comment\n# comment\n# comment ##\n\n",
		output: []Token{
			{
				TokenType: COMMENT,
				Value:     "#comment",
				Position:  1,
				Line:      1,
			},
			{
				TokenType: COMMENT,
				Value:     `# comment`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: COMMENT,
				Value: `# comment ##

`,
				Position: 1,
				Line:     3,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  1,
				Line:      5,
			},
		},
	}
	testHashtag2 = testList{
		input: "/# #/ \n /# /#",
		output: []Token{
			{
				TokenType: DIV,
				Value:     `/`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: COMMENTGROUP,
				Value:     "# #/ \n /#",
				Position:  2,
				Line:      1,
			},
			{
				TokenType: DIV,
				Value:     `/`,
				Position:  5,
				Line:      2,
			},
			{
				TokenType: COMMENT,
				Value:     `#`,
				Position:  6,
				Line:      2,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  7,
				Line:      2,
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
	testHashtag6 = testList{
		input: " #/ #/# ok",
		output: []Token{
			{
				TokenType: COMMENTGROUP,
				Value:     `#/ #/#`,
				Position:  2,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     "ok",
				Position:  9,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  11,
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
				TokenType: XORBIN,
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
	testMultiLigneString = testList{
		input: "b := \"hello\nworld\";",
		output: []Token{
			{
				TokenType: TEXT,
				Value:     `b`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: COLON,
				Value:     `:`,
				Position:  3,
				Line:      1,
			},
			{
				TokenType: ASSIGN,
				Value:     `=`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     "hello",
				Position:  7,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `world`,
				Position:  1,
				Line:      2,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  6,
				Line:      2,
			},
			{
				TokenType: STRING,
				Value:     `;`,
				Position:  7,
				Line:      2,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  8,
				Line:      2,
			},
		},
	}
	testCondInParen = testList{
		input: "import \"console\";\n\n# this is a comment\n\n#/ this is a block comment /#\n\nconsole.println(\"Hello, World!\");",
		output: []Token{
			{
				TokenType: EOF,
				Value:     ``,
				Position:  8,
				Line:      2,
			},
		},
	}
	testCommentGroup = testList{
		input: "#//# ok",
		output: []Token{
			{
				TokenType: COMMENTGROUP,
				Value:     `#//#`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: TEXT,
				Value:     `ok`,
				Position:  6,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  8,
				Line:      1,
			},
		},
	}
	testEmptyString = testList{
		input: "\"\";",
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
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
				TokenType: EOF,
				Value:     ``,
				Position:  4,
				Line:      1,
			},
		},
	}
	testStringWithBSlash = testList{
		input: "\"\\\"\\\"\"",
		output: []Token{
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `\"\"`,
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
				TokenType: EOF,
				Value:     ``,
				Position:  7,
				Line:      1,
			},
		},
	}
	testBSlashstring = testList{
		input: "\\\\ \"\\\\\"",
		output: []Token{
			{
				TokenType: TEXT,
				Value:     `\\`,
				Position:  1,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  4,
				Line:      1,
			},
			{
				TokenType: STRING,
				Value:     `\\`,
				Position:  5,
				Line:      1,
			},
			{
				TokenType: DQUOTE,
				Value:     `"`,
				Position:  7,
				Line:      1,
			},
			{
				TokenType: EOF,
				Value:     ``,
				Position:  8,
				Line:      1,
			},
		},
	}

	testHshTagminus = testList{
		input: "import \"console\";\n\n\n\n#test Print\nconsole.println(\"Test Print\");\nconsole.println(\"Hello World\");\nconsole.println(\"Test Print Past\");\n\n\n# test add\nconsole.println(\"Test Add\");\nconsole.println(\"1+1 = \",1+1);\nconsole.println(\"1+\\\"hello\\\" = \",1+\"hello\");\nconsole.println(\"\\\"hello\\\"+1 = \",\"hello\"+1);\nconsole.println(\"\\\"hello\\\"+\\\"world\\\" = \",\"hello\"+\"world\");\nconsole.println(\"1.3+0.2 = \",1.3+0.2);\nconsole.println(\"1.3+\\\"hello\\\" = \",1.3+\"hello\");\nconsole.println(\"\\\"hello\\\"+1.3 = \",\"hello\"+1.3);\nconsole.println(\"1 + 1.3 = \",1+1.3);\nconsole.println(\"1.3 + 1 = \",1.3+1);\nconsole.println(\"\\\"hello\\\"+true = \",\"hello\"+true);\nconsole.println(\"true+\\\"hello\\\" = \",true+\"hello\");\n\n\n# test sub\nconsole.println(\"Test sub\");\nconsole.println(\"1-1 = \",1-1);\nconsole.println(\"1.3-0.2 = \",1.3-0.2);\nconsole.println(\"2.0-1 = \",2.0-1);\nconsole.println(\"1-2.0 = \",1-2.0);\n\n# test div\nconsole.println(\"Test div\");\nconsole.println(\"10/5 = \",10/5);\nconsole.println(\"2.5/2.0 = \",2.5/2.0);\nconsole.println(\"2.5/2 = \",2.5/2);\nconsole.println(\"2/2.5 = \",2/2.5);\n\n# test qot\nconsole.println(\"Test qot\");\nconsole.println(\"10//3 = \",10//3);\n\n# test mod\nconsole.println(\"Test mod\");\nconsole.println(\"10%3 = \",10%3);\n\n# test mult\nconsole.println(\"Test mult\");\nconsole.println(\"10*3 = \",10*3);\nconsole.println(\"2.5*2.0 = \",2.5*2.0);\nconsole.println(\"2.5*2 = \",2.5*2);\nconsole.println(\"2*2.5 = \",2*2.5);\nconsole.println(\"\\\"a\\\"*2 = \",\"a\"*2);\nconsole.println(\"2*\\\"a\\\" = \",2*\"a\");\n\n\n# var test\nconsole.println(\"Test Var int\");\nvar a int;\nconsole.println(\"a is :\",a);\nvar b int = 10;\nconsole.println(\"b is :\",b);\nconsole.println(\"Test Var float\");\nvar c float;\nconsole.println(\"c is :\",c);\nvar d float = 20.1;\nconsole.println(\"d is :\",d);\nconsole.println(\"Test Var string\");\nvar e string;\nconsole.println(\"e is :\",e);\nvar f string = \"Hello World\";\nconsole.println(\"f is :\",f);\nconsole.println(\"Test Var bool\");\nvar g bool;\nconsole.println(\"g is :\",g);\nvar h bool = true;\nconsole.println(\"h is :\",h);\n\n# test add var\nconsole.println(\"Test add var\");\nconsole.println(\"Test add int\");\nconsole.println(\"a+b = \",a+b,\"result should be 10\");\nconsole.println(\"b+1 = \",b+1,\"result should be 11\");\nconsole.println(\"1+b = \",1+b,\"result should be 11\");\nconsole.println(\"b+1.1 = \",b+1.1,\"result should be 11.1\");\nconsole.println(\"1.1+b = \",1.1+b,\"result should be 11.1\");\nconsole.println(\"b+\\\"hello\\\" = \",b+\"hello\",\"result should be 10hello\");\nconsole.println(\"\\\"hello\\\"+b = \",\"hello\"+b,\"result should be hello10\");\n\nconsole.println(\"Test add float\");\nconsole.println(\"c+d = \",c+d,\"result should be 20.1\");\nconsole.println(\"d+1 = \",d+1,\"result should be 21.1\");\nconsole.println(\"1+d = \",1+d,\"result should be 21.1\");\nconsole.println(\"d+1.1 = \",d+1.1,\"result should be 21.2\");\nconsole.println(\"1.1+d = \",1.1+d,\"result should be 21.2\");\nconsole.println(\"d+\\\"hello\\\" = \",d+\"hello\",\"result should be 20.1hello\");\nconsole.println(\"\\\"hello\\\"+d = \",\"hello\"+d,\"result should be hello20.1\");\n\nconsole.println(\"Test add string\");\nconsole.println(\"e+f = \",e+f,\"result should be Hello World\");\nconsole.println(\"f+\\\"hello\\\" = \",f+\"hello\",\"result should be Hello Worldhello\");\nconsole.println(\"\\\"hello\\\"+f = \",\"hello\"+f,\"result should be helloHello World\");\n\nconsole.println(\"Test add bool\");\nconsole.println(\"h+\\\"hello\\\" = \",h+\"hello\",\"result should be truehello\");\nconsole.println(\"\\\"hello\\\"+h = \",\"hello\"+h,\"result should be hellotrue\");\n\n# test sub var\nconsole.println(\"Test sub var\");\nconsole.println(\"Test sub int\");\nconsole.println(\"b-a = \",b-a,\"result should be 10\");\nconsole.println(\"b-1 = \",b-1,\"result should be 9\");\nconsole.println(\"1-b = \",1-b,\"result should be -9\");\nconsole.println(\"b-1.1 = \",b-1.1,\"result should be 8.9\");\nconsole.println(\"1.1-b = \",1.1-b,\"result should be -8.9\");\n\nconsole.println(\"Test sub float\");\nconsole.println(\"d-c = \",d-c,\"result should be 20.1\");\nconsole.println(\"d-1 = \",d-1,\"result should be 19.1\");\nconsole.println(\"1-d = \",1-d,\"result should be -19.1\");\nconsole.println(\"d-1.1 = \",d-1.1,\"result should be 19\");\nconsole.println(\"1.1-d = \",1.1-d,\"result should be -19\");\n\n# test div var\nconsole.println(\"Test div var\");\nconsole.println(\"Test div int\");\nconsole.println(\"a/b = \",a/b,\"result should be 0\");\nconsole.println(\"b/1 = \",b/1,\"result should be 10\");\nconsole.println(\"1/b = \",1/b,\"result should be 0.1\");\nconsole.println(\"b/1.1 = \",b/1.1,\"result should be 9.09\");\nconsole.println(\"1.1/b = \",1.1/b,\"result should be 0.11\");\n\nconsole.println(\"Test div float\");\nconsole.println(\"c/d = \",c/d,\"result should be 0\");\nconsole.println(\"d/1 = \",d/1,\"result should be 20.1\");\nconsole.println(\"1/d = \",1/d,\"result should be 0.049\");\nconsole.println(\"d/1.1 = \",d/1.1,\"result should be 18.2\");\nconsole.println(\"1.1/d = \",1.1/d,\"result should be 0.054\");\n\n# test qot var\nconsole.println(\"Test qot var\");\nconsole.println(\"Test qot int\");\nconsole.println(\"a//b = \",a//b,\"result should be 0\");\nconsole.println(\"b//1 = \",b//1,\"result should be 10\");\nconsole.println(\"1//b = \",1//b,\"result should be 0\");\n\n# test mod var\nconsole.println(\"Test mod var\");\nconsole.println(\"Test mod int\");\nconsole.println(\"a%b = \",a%b,\"result should be 0\");\nconsole.println(\"b%1 = \",b%1,\"result should be 0\");\nconsole.println(\"1%b = \",1%b,\"result should be 1\");\n\n# test mult var\nconsole.println(\"Test mult var\");\nconsole.println(\"Test mult int\");\nconsole.println(\"a*b = \",a*b,\"result should be 0\");\nconsole.println(\"b*1 = \",b*1,\"result should be 10\");\nconsole.println(\"1*b = \",1*b,\"result should be 10\");\nconsole.println(\"b*1.1 = \",b*1.1,\"result should be 11\");\nconsole.println(\"1.1*b = \",1.1*b,\"result should be 11\");\n\nconsole.println(\"Test mult float\");\nconsole.println(\"c*d = \",c*d,\"result should be 0\");\nconsole.println(\"d*1 = \",d*1,\"result should be 20.1\");\nconsole.println(\"1*d = \",1*d,\"result should be 20.1\");\nconsole.println(\"d*1.1 = \",d*1.1,\"result should be 22.11\");\nconsole.println(\"1.1*d = \",1.1*d,\"result should be 22.11\");\n\n# test Array\n# ----------------------------------------------------------------------------------------------------------------\n",
		output: []Token{
			{
				TokenType: EOF,
				Value:     ``,
				Position:  1,
				Line:      1,
			},
		},
	}

	testDump = testList{
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
)
