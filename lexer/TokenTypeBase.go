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
		var tempToken ITokenType
		tempToken, index = t.IsInvolvedWithLastStep(l)
		// if no matching, classic behavior, otherwise, COMMENTGROUPEND style behavior.
		if index == -1 {
			// Classic TriggerBy Behavior
			finded := findNameInEveryTokenType(l.TriggerBy)

			if NameFromGet(finded.Get()) != "NULL" {
				findNameInEveryTokenType(l.TriggerBy).Resolve(l)
			}
		} else {
			// Special TriggerBy Behavior

			// related = what is the possible merged or composed result token if the actual token and the previous
			// one merge or compose together.
			related := t.Result[index]
			// update the lexer to acknoledge the new token to work with.

			if NameFromGet(findNameInTriggerTokenType(NameFromGet(tempToken.Get())).Get()) != "NULL" &&
				NameFromGet(findNameInTriggerTokenType(l.TriggerBy).Get()) != "NULL" {

				triggerByToken := findNameInTriggerTokenType(NameFromGet(tempToken.Get()))
				l.indent[0] = &related
				// compose the token BUT end the triggerBy
				(*l).ComposeToken(triggerByToken.Get()[len(triggerByToken.Get())-1])
				l.TriggerBy = ""
				// reset the reading head of our lexer.
				l.prevIndex = l.index
			} else if NameFromGet(findNameInMergerTokenType(NameFromGet(tempToken.Get())).Get()) != "NULL" &&
				NameFromGet(findNameInMergerTokenType(l.TriggerBy).Get()) != "NULL" {

				triggerByToken := findNameInMergerTokenType(NameFromGet(tempToken.Get()))

				if NameFromGet(triggerByToken.Get()) == l.TriggerBy {

					(*l).ComposeToken(NameFromGet(related.Get()))
					l.TriggerBy = NameFromGet(related.Get())
					l.prevIndex = l.index
				} else {
					if findNameInMergerTokenType(l.TriggerBy).IsClosedBySyntaxe(NameFromGet(tempToken.Get())) == -1 {

						(*l).ComposeToken(NameFromGet(related.Get()))
						l.TriggerBy = NameFromGet(related.Get())
						l.prevIndex = l.index
					} else {
						l.indent[0] = &related
						// compose the token BUT end the triggerBy
						(*l).ComposeToken(NameFromGet(triggerByToken.Get()))
						l.TriggerBy = ""
						// reset the reading head of our lexer.
						l.prevIndex = l.index
					}
				}

			} else {

				findNameInEveryTokenType(l.TriggerBy).Resolve(l)
			}
		}
	} else {
		// classic Behavior

		// spaces must be ignored for all merge or compose behavior.
		if !(*l).isSpaces {
			var finded ITokenType
			// try to find some Involved token if possible.
			// if none find, classic behavior, otherwise compose behavior.
			finded, index = t.IsInvolvedWith(l)
			if index != -1 {
				// avoid compose with an ended merger or trigger token
				if NameFromGet(findNameInMergerTokenType(NameFromGet(finded.Get())).Get()) != "NULL" {
					index = -1
				} else if NameFromGet(findNameInTriggerTokenType(NameFromGet(finded.Get())).Get()) != "NULL" {
					index = -1
				}
			}
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
				Result:  []TokenTypeCompositeBehavior{CCOMMENTGROUP}},
		},
		Result: []TokenTypeCompositeBehavior{
			CQOT, CCOMMENTGROUP,
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
		Name:   COMMENTGROUP + "END",
		Syntax: []string{},
	}
)
