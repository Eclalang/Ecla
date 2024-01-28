package v2

type TokenTypeMergerBehavior struct {
	Name   string
	Syntax []string

	CloseBy   []ITokenType
	Result    []TokenTypeCompositeBehavior
	Involved  []ITokenType
	Composite []TokenTypeMergerBehavior
}

func (t *TokenTypeMergerBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve merger")
	index := -1
	if l.TriggerBy == "" {
		if !(*l).isSpaces {
			_, index = t.IsInvolvedWith(l)
		} else {
			(*l).isSpaces = false
		}
		if index == -1 {
			(*l).AddToken(t.Name)
			l.TriggerBy = t.Name
			l.prevIndex = l.index
		} else {
			l.DEBUGLEXER("COMMENTGROUP TEST")
			(*l).ComposeToken(t.Composite[index].Name)
			l.TriggerBy = t.Composite[index].Name
			l.prevIndex = l.index
		}
	} else {
		if !(*l).isSpaces {
			_, index = t.IsInvolvedWith(l)
		} else {
			(*l).isSpaces = false
		}
		if index == -1 {
			if l.sizeOfTokenReversed != -1 {
				identified := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
				indexOfClose := t.IsClosedBySyntaxe(identified)
				if indexOfClose != -1 {
					//close , donc doit mettre RESULT+CLOSE en token
					l.FindSyntax()
					temp := l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
					if findNameInEveryTokenType(temp, Every) != nil {
						l.tempVal = temp
						l.ComposeToken(t.Result[0].Get()[len(t.Result[0].Get())-1])
					} else {
						l.ComposeToken(l.TriggerBy)
						l.tempVal = temp
						l.TriggerBy = ""
						l.indent[0].Resolve(l)

					}
					l.tempVal = ""
					l.TriggerBy = ""
					l.prevIndex = l.index
				}
			}
		} else {
			t.Resolve(l)
		}

	}
}
func (t *TokenTypeMergerBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeMergerBehavior) InvolvedWith() []ITokenType {
	return t.CloseBy
}

func (t *TokenTypeMergerBehavior) getResult() []TokenTypeCompositeBehavior {
	return t.Result
}
func (t *TokenTypeMergerBehavior) IsInvolvedWith(lexer *TLexer) (ITokenType, int) {
	for i, token := range t.Involved {
		if len(lexer.Ret())-1 >= 0 {
			if token.Get()[len(token.Get())-1] == lexer.Ret()[len(lexer.Ret())-1].TokenType || (token.Get()[len(token.Get())-1] == "SELF" && t.Name == lexer.Ret()[len(lexer.Ret())-1].TokenType) {
				return token, i
			}
		}
	}
	return nil, -1
}
func (t *TokenTypeMergerBehavior) IsClosedByName(other *TokenTypeMergerBehavior) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == other.Name {
			return i
		}
	}
	return -1
}
func (t *TokenTypeMergerBehavior) IsClosedBySyntaxe(otherName string) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[0] == SELF.Name {
			tokenType = t
		}
		for y := 0; y < len(tokenType.Get()); y++ {
			if tokenType.Get()[y] == otherName {
				return i
			}
		}

	}
	return -1
}

var (
	TCOMMENT = TokenTypeMergerBehavior{
		Name: COMMENT,
		Syntax: []string{
			"#",
		},
		CloseBy: []ITokenType{
			&RETURN,
		},
		Result: []TokenTypeCompositeBehavior{
			CCOMMENT,
		},
		Involved: []ITokenType{
			&BDIV,
		},
		Composite: []TokenTypeMergerBehavior{
			TCOMMENTGROUP,
		},
	}
	TCOMMENTGROUP = TokenTypeMergerBehavior{
		Name: COMMENTGROUP,
		Syntax: []string{
			"/#",
		},
		CloseBy: []ITokenType{
			&CCOMMENTGROUPEND,
		},
		Result: []TokenTypeCompositeBehavior{
			CCOMMENT,
		},
	}
)
