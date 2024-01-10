package v2

type TokenTypeTriggerBehavior struct {
	Name   string
	Syntax []string

	CloseBy []ITokenType
	Result  []TokenTypeCompositeBehavior
}

func (t *TokenTypeTriggerBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve trigger")

	if l.TriggerBy == "" {
		l.TriggerBy = t.Name
		(*l).AddToken(t.Name)
		l.prevIndex = l.index
	} else {
		identified := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
		indexOfClose := t.IsClosedBySyntaxe(identified)
		if indexOfClose != -1 {
			//close , donc doit mettre RESULT | CLOSE en token
			l.FindSyntax()
			temp := identified
			l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
			l.AddToken(t.Result[indexOfClose].Name)
			l.DEBUGLEXER("in resolve trigger AFTER ADD")
			l.tempVal = temp
			l.TriggerBy = ""
			l.indent[0].Resolve(l)
			l.TriggerBy = ""
			l.prevIndex = l.index
		}

	}
}
func (t *TokenTypeTriggerBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeTriggerBehavior) InvolvedWith() []ITokenType {
	return t.CloseBy
}

func (t *TokenTypeTriggerBehavior) IsClosedByName(otherName string) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == otherName {
			return i
		}
	}
	return -1
}
func (t *TokenTypeTriggerBehavior) IsClosedBySyntaxe(otherName string) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[0] == SELF.Name {
			tokenType = t
		}
		for y := 0; y < len(tokenType.Get()); y++ {
			println(tokenType.Get()[y])
			if tokenType.Get()[y] == otherName {
				return i
			}
		}

	}
	return -1
}

var (
	TDQUOTE = TokenTypeTriggerBehavior{
		Name: DQUOTE,
		Syntax: []string{
			"\"",
		},
		CloseBy: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CSTRING,
		},
	}
	TSQUOTE = TokenTypeTriggerBehavior{
		Name: SQUOTE,
		Syntax: []string{
			"'",
		},
		CloseBy: []ITokenType{
			&SELF,
		},
		Result: []TokenTypeCompositeBehavior{
			CCHAR,
		},
	}
)
