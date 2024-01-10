package v2

type TokenTypeMergerBehavior struct {
	Name   string
	Syntax []string

	CloseBy []ITokenType
	Result  []TokenTypeCompositeBehavior
}

func (t *TokenTypeMergerBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve merger")

	if l.TriggerBy == "" {
		(*l).AddToken(t.Name)
		l.TriggerBy = t.Name
		l.prevIndex = l.index
	} else {
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
			println(tokenType.Get()[y])
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
	}
)
