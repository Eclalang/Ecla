package lexer

type TokenTypeNullableBehavior struct {
	Name   string
	Syntax []string
}

func (t *TokenTypeNullableBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("IN NULLABLE")
	if NameFromGet(l.lastStepToken.Get()) == t.Name {
		l.lastStepToken = &SELF
		l.prevIndex++
		l.tempVal = l.sentence[l.prevIndex:l.index]
	}
}

func (t *TokenTypeNullableBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}
func (t *TokenTypeNullableBehavior) InvolvedWith() []ITokenType {
	return []ITokenType{}
}

var (
	NBSLASH = TokenTypeNullableBehavior{
		Name: "BSLASH",
		Syntax: []string{
			`\`,
		},
	}
)
