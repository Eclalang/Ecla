package lexer

type TokenTypeSpacesBehavior struct {
	Name   string
	Syntax []string
}

func (t *TokenTypeSpacesBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve spaces")
	if (*l).TriggerBy != "" {
		l.DEBUGLEXER("in resolve TriggerBy")
		findNameInEveryTokenType(l.TriggerBy, Every).Resolve(l)
	} else {
		if (*l).Inquote() {
			findNameInEveryTokenType((*l).TriggerBy, Every).Resolve(l)

		} else {
			(*l).isSpaces = true
		}
		if t.Name == "\n" {
			(*l).line++
			(*l).position = 1
		}

		l.prevIndex = l.index
	}

}

func (t *TokenTypeSpacesBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeSpacesBehavior) InvolvedWith() []ITokenType {
	return []ITokenType{}
}

var (
	EMPTY = TokenTypeSpacesBehavior{
		Name: "",
		Syntax: []string{
			" ",
			"\t",
			"\r",
		},
	}
	RETURN = TokenTypeSpacesBehavior{
		Name: "\n",
		Syntax: []string{
			"\n",
		},
	}
	SELF = TokenTypeSpacesBehavior{
		Name:   "SELF",
		Syntax: []string{},
	}
)
