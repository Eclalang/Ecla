package v2

type TokenTypeBaseBehavior struct {
	Name   string
	Syntax []string

	Involved []ITokenType
	Result   []TokenTypeCompositeBehavior
}

func (t *TokenTypeBaseBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve base")
	index := -1

	if (*l).TriggerBy != "" {
		l.DEBUGLEXER("in resolve TriggerBy")
		findNameInEveryTokenType(l.TriggerBy, Every).Resolve(l)
	} else {
		if !(*l).isSpaces {
			_, index = t.IsInvolvedWith(l)
		} else {
			(*l).isSpaces = false
		}
		if index == -1 {
			(*l).AddToken(t.Name)
		} else {
			(*l).ComposeToken(t.Result[index].Name)
		}

		l.prevIndex = l.index
	}

}

func (t *TokenTypeBaseBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeBaseBehavior) InvolvedWith() []ITokenType {
	return t.Involved
}

func (t *TokenTypeBaseBehavior) IsInvolvedWith(lexer *TLexer) (ITokenType, int) {
	for i, token := range t.Involved {
		if len(lexer.Ret())-1 >= 0 {
			if token.Get()[len(token.Get())-1] == lexer.Ret()[len(lexer.Ret())-1].TokenType || (token.Get()[len(token.Get())-1] == "SELF" && t.Name == lexer.Ret()[len(lexer.Ret())-1].TokenType) {
				return token, i
			}
		}
	}
	return nil, -1
}

var (
	BPERIOD = TokenTypeBaseBehavior{
		Name: PERIOD,
		Syntax: []string{
			".",
		},
		Involved: []ITokenType{
			&BINT,
		},
		Result: []TokenTypeCompositeBehavior{
			CFLOAT,
		},
	}
	BCOLON = TokenTypeBaseBehavior{
		Name: COLON,
		Syntax: []string{
			":",
		},
	}
	BLBRACE = TokenTypeBaseBehavior{
		Name: LBRACE,
		Syntax: []string{
			"{",
		},
	}
	BRBRACE = TokenTypeBaseBehavior{
		Name: RBRACE,
		Syntax: []string{
			"}",
		},
	}
	BLBRACKET = TokenTypeBaseBehavior{
		Name: LBRACKET,
		Syntax: []string{
			"[",
		},
	}
	BRBRACKET = TokenTypeBaseBehavior{
		Name: RBRACKET,
		Syntax: []string{
			"]",
		},
	}
	BCOMMA = TokenTypeBaseBehavior{
		Name: COMMA,
		Syntax: []string{
			",",
		},
	}
	BBOOL = TokenTypeBaseBehavior{
		Name: BOOL,
		Syntax: []string{
			"true",
			"false",
		},
	}

	BOR = TokenTypeBaseBehavior{
		Name: OR,
		Syntax: []string{
			"||",
		},
	}
	BAND = TokenTypeBaseBehavior{
		Name: AND,
		Syntax: []string{
			"&&",
		},
	}
	BLPAREN = TokenTypeBaseBehavior{
		Name: LPAREN,
		Syntax: []string{
			"(",
		},
	}
	BRPAREN = TokenTypeBaseBehavior{
		Name: RPAREN,
		Syntax: []string{
			")",
		},
	}
	BEOL = TokenTypeBaseBehavior{
		Name: EOL,
		Syntax: []string{
			";",
		},
	}
	BMURLOC = TokenTypeBaseBehavior{
		Name: MURLOC,
		Syntax: []string{
			"mgrlgrl",
		},
	}
	BINT = TokenTypeBaseBehavior{
		Name: INT,
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
		Involved: []ITokenType{
			&SELF, &CFLOAT,
		},
		Result: []TokenTypeCompositeBehavior{
			CINT, CFLOAT,
		},
	}
	BADD = TokenTypeBaseBehavior{
		Name: ADD,
		Syntax: []string{
			"+",
		},
		Involved: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CINC,
		},
	}
	BSUB = TokenTypeBaseBehavior{
		Name: SUB,
		Syntax: []string{
			"-",
		},
		Involved: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CDEC,
		},
	}
	BMOD = TokenTypeBaseBehavior{
		Name: MOD,
		Syntax: []string{
			"%",
		},
		Involved: []ITokenType{},
		Result:   []TokenTypeCompositeBehavior{},
	}
	BDIV = TokenTypeBaseBehavior{
		Name: DIV,
		Syntax: []string{
			"/",
		},
		Involved: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CQOT,
		},
	}
	BMULT = TokenTypeBaseBehavior{
		Name: MULT,
		Syntax: []string{
			"*",
		},
		Involved: []ITokenType{},
		Result:   []TokenTypeCompositeBehavior{},
	}
	BXORBIN = TokenTypeBaseBehavior{
		Name: XORBIN,
		Syntax: []string{
			"^",
		},
		Involved: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CXOR,
		},
	}
	BLSS = TokenTypeBaseBehavior{
		Name: LSS,
		Syntax: []string{
			"<",
		},
		Involved: []ITokenType{},
		Result:   []TokenTypeCompositeBehavior{},
	}
	BGTR = TokenTypeBaseBehavior{
		Name: GTR,
		Syntax: []string{
			">",
		},
		Involved: []ITokenType{},
		Result:   []TokenTypeCompositeBehavior{},
	}
	BNOT = TokenTypeBaseBehavior{
		Name: NOT,
		Syntax: []string{
			"!",
		},
		Involved: []ITokenType{},
		Result:   []TokenTypeCompositeBehavior{},
	}
	BASSIGN = TokenTypeBaseBehavior{
		Name: ASSIGN,
		Syntax: []string{
			"=",
		},
		Involved: []ITokenType{
			&SELF, &BADD, &BSUB, &BMOD, &BDIV, &BMULT, &BNOT, &BGTR, &BLSS, &CQOT,
		},
		Result: []TokenTypeCompositeBehavior{
			CEQUAL, CADDASSIGN, CSUBASSIGN, CMODASSIGN, CDIVASSIGN, CMULTASSIGN, CNEQ, CGEQ, CLEQ, CQOTASSIGN,
		},
	}
	BCOMMENTGROUPEND = TokenTypeBaseBehavior{
		Name: COMMENTGROUP,
		Syntax: []string{
			"#/",
		},
	}
)
