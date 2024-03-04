package lexer

type TokenTypeBaseBehavior struct {
	Name   string
	Syntax []string

	Involved []ITokenType
	Result   []TokenTypeCompositeBehavior
}

// Resolve Base Behavior Token
//
// imply no special behavior
func (t *TokenTypeBaseBehavior) Resolve(l *TLexer) {
	index := -1
	l.DEBUGLEXER("IN BASE")
	if (*l).TriggerBy != "" {
		// in Trigger By Behavior

		// verify the previous token created (even in a merging situation) to detect for exemple the end of a
		// COMMENTGROUP "trigger by" behavior.
		_, index = t.IsInvolvedWithLastStep(l)

		// if no matching, classic behavior, otherwise, COMMENTGROUPEND style behavior.
		if index == -1 {
			// Classic TriggerBy Behavior
			findNameInEveryTokenType(l.TriggerBy).Resolve(l)
		} else {
			// Special TriggerBy Behavior

			// related = what is the possible merged or composed result token if the actual token and the previous
			// one merge or compose together.

			related := t.Result[index]
			triggerByToken := findNameInEveryTokenType((*l).TriggerBy)
			finded := -1
			// update the lexer to acknoledge the new token to work with.
			for i, v := range triggerByToken.InvolvedWith() {
				if v.Get()[len(v.Get())-1] == related.Name {
					finded = i
					l.indent[0] = &related
				}
			}
			if finded != -1 {
				// compose the token BUT end the triggerBy
				(*l).ComposeToken(triggerByToken.Get()[len(triggerByToken.Get())-1])
				l.TriggerBy = ""
				// reset the reading head of our lexer.
				l.prevIndex = l.index
			} else {
				triggerByToken.Resolve(l)
			}

		}
	} else {
		// classic Behavior

		// spaces must be ignored for all merge or compose behavior.
		if !(*l).isSpaces {
			// try to find some Involved token if possible.
			// if none find, classic behavior, otherwise compose behavior.
			_, index = t.IsInvolvedWith(l)
		} else {
			// this tokentype is not a spaces.
			(*l).isSpaces = false
		}
		if index == -1 {
			// classic behavior
			(*l).AddToken(t.Name)
		} else {
			// compose behavior
			(*l).ComposeToken(t.Result[index].Name)
		}

		// reset the reading head of our lexer.
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

func (t *TokenTypeBaseBehavior) IsInvolvedWithLastStep(lexer *TLexer) (ITokenType, int) {
	for i, token := range t.Involved {
		if token.Get()[len(token.Get())-1] == lexer.lastStepToken.Get()[len(lexer.lastStepToken.Get())-1] {
			return token, i
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
			&SELF, &TokenTypeMergerBehavior{
				Name: COMMENT, Syntax: []string{"#"},
				CloseBy: []ITokenType{&RETURN},
				Result:  []TokenTypeCompositeBehavior{CCOMMENT}},
		},
		Result: []TokenTypeCompositeBehavior{
			CQOT, CCOMMENTGROUPEND,
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
)
