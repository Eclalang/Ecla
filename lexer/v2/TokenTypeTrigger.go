package v2

type TokenTypeTriggerBehavior struct {
	Name   string
	Syntax []string

	CloseBy []ITokenType
	Result  []TokenTypeCompositeBehavior
}

func (t *TokenTypeTriggerBehavior) Resolve(l *TLexer) {
	println("in Resolve trigger...")

	if l.inTriggerof.GetITokenType().Get()[len(l.inTriggerof.GetITokenType().Get())-1] == "" {
		println("Set up trigger...")
		(*l).AddToken(t.Name)
		l.inTriggerof = t
		l.prevIndex = l.index
	} else {
		println("in Resolve trigger already triggerd...")
		closer := l.inTriggerof.IsClosedBy(t)
		if closer != -1 {
			//pas de close, donc doit continuer
		} else {
			//close , donc doit mettre RESULT | CLOSE en token
			l.FindSyntax()
			temp := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
			l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
			l.AddToken(l.inTriggerof.Result[0].Name)
			l.tempVal = temp
			l.inTriggerof = &TokenTypeTriggerBehavior{
				Name: "",
			}
			l.indent[0].Resolve(l)

			l.inTriggerof = &TokenTypeTriggerBehavior{
				Name: "",
			}
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

func (t *TokenTypeTriggerBehavior) IsClosedBy(other *IActionToken) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == (*other).GetITokenType().Get()[len((*other).GetITokenType().Get())-1] {
			return i
		}
	}
	return -1
}
func (t *TokenTypeTriggerBehavior) GetITokenType() ITokenType {
	return t
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
)
