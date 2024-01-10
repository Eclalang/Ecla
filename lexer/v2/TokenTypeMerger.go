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
		closer := findNameInEveryTokenType(t.CloseBy[0].Get()[len(t.CloseBy[0].Get())-1], Every)
		if closer == nil {
			//pas de close, donc doit continuer
		} else {
			//close , donc doit mettre RESULT+CLOSE en token
			l.FindSyntax()
			temp := l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
			if findNameInEveryTokenType(temp, Every) != nil {
				l.tempVal = temp
				l.AddToken(t.Result[0].Get()[len(t.Result[0].Get())-1])
			} else {
				l.AddToken(l.TriggerBy)
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

func (t *TokenTypeMergerBehavior) IsClosedBy(other *TokenTypeMergerBehavior) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == other.Name {
			return i
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
