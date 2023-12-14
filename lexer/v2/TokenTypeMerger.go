package v2

type TokenTypeMergerBehavior struct {
	Name   string
	Syntax []string

	CloseBy []ITokenType
	Result  []TokenTypeCompositeBehavior
}

func (t *TokenTypeMergerBehavior) Resolve(l *TLexer) {
	println("in Resolve trigger...")

	if l.inTriggerof.Name == "" {
		println("Set up trigger...")
		(*l).AddToken(t.Name)
		l.inTriggerof = *t
		l.prevIndex = l.index
	} else {
		println("in Resolve trigger already triggerd...")
		closer := l.inTriggerof.IsCLosedBy(t)
		if closer != -1 {
			//pas de close, donc doit continuer
		} else {
			//close , donc doit mettre RESULT | CLOSE en token
			l.FindSyntax()
			temp := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
			l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
			println("in Resolve trigger find by name..")
			if findNameInEveryTokenType(l.inTriggerof.Result[0].Name, Every) != nil {
				findNameInEveryTokenType(l.inTriggerof.Result[0].Name, Every).Resolve(l)
				l.AddToken(l.inTriggerof.Result[0].Name)
				l.tempVal = temp
				l.inTriggerof = TokenTypeTriggerBehavior{
					Name: "",
				}
				l.indent[0].Resolve(l)
			} else {
				l.AddToken(l.inTriggerof.Result[0].Name)
				l.tempVal = temp
				l.inTriggerof = TokenTypeTriggerBehavior{
					Name: "",
				}
				l.indent[0].Resolve(l)

			}
			l.inTriggerof = TokenTypeTriggerBehavior{
				Name: "",
			}
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

func (t *TokenTypeMergerBehavior) IsClosedBy(other *IActionToken) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == (*other).GetITokenType().Get()[len((*other).GetITokenType().Get())-1] {
			return i
		}
	}
	return -1
}

func (t *TokenTypeMergerBehavior) GetITokenType() ITokenType {
	return t
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
