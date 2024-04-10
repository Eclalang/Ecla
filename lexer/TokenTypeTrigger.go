package lexer

type TokenTypeTriggerBehavior struct {
	Name   string
	Syntax []string

	CloseBy   []ITokenType
	Result    []TokenTypeCompositeBehavior
	Involved  []ITokenType
	Composite []TokenTypeCompositeBehavior
}

func (t *TokenTypeTriggerBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("IN TRIGGER")
	if l.TriggerBy == EMPTY.Name {
		l.TriggerBy = t.Name
		(*l).AddToken(t.Name)
		l.prevIndex = l.index
	} else {

		if l.index > len(l.sentence) {
			l.AddToken(t.Result[0].Name)
			l.prevIndex = l.index

		} else if l.sizeOfTokenReversed != -1 {
			//println(l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:])
			identified := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
			triggerByToken := findNameInTriggerTokenType(l.TriggerBy)
			indexOfClose := triggerByToken.IsClosedBySyntaxe(NameFromGet(l.indent[0].Get()))
			if indexOfClose != -1 {
				//close , donc doit mettre RESULT | CLOSE en token
				if NameFromGet(l.lastStepToken.Get()) == BSLASH {

				} else {
					l.FindSyntax()
					temp := identified
					l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
					l.position -= 1

					l.AddToken(t.Result[indexOfClose].Name)
					l.position += 1
					l.tempVal = temp
					l.TriggerBy = ""
					if NameFromGet(l.indent[0].Get()) == "\n" {
						l.line -= 1
					}
					l.indent[0].Resolve(l)
					l.TriggerBy = ""
					l.prevIndex = l.index
				}

			}
		} else {
			triggerByToken := findNameInTriggerTokenType(l.TriggerBy)
			indexOfClose := triggerByToken.IsClosedBySyntaxe(NameFromGet(l.indent[0].Get()))

			if indexOfClose != -1 {
				//close , donc doit mettre CLOSE en token

				l.TriggerBy = ""
				l.indent[0].Resolve(l)
				l.TriggerBy = ""
				l.prevIndex = l.index
			}
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
		if tokenType.Get()[len(tokenType.Get())-1] == SELF.Name {
			if t.Name == otherName {
				return i
			}
		} else {
			if tokenType.Get()[len(tokenType.Get())-1] == otherName {
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
			&SELF, &RETURN,
		},
		Result: []TokenTypeCompositeBehavior{
			CSTRING, CSTRING,
		},
	}
	TSQUOTE = TokenTypeTriggerBehavior{
		Name: SQUOTE,
		Syntax: []string{
			"'",
		},
		CloseBy: []ITokenType{
			&SELF, &RETURN,
		},
		Result: []TokenTypeCompositeBehavior{
			CCHAR, CCHAR,
		},
	}
)