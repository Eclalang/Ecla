package lexer

type TokenTypeSpacesBehavior struct {
	Name   string
	Syntax []string
}

func (t *TokenTypeSpacesBehavior) Resolve(l *TLexer) {
	if (*l).TriggerBy != "" {
		findNameInEveryTokenType(l.TriggerBy).Resolve(l)
	} else {
		if (*l).Inquote() {
			findNameInEveryTokenType((*l).TriggerBy).Resolve(l)

		} else {
			(*l).isSpaces = true
		}

		l.prevIndex = l.index
	}
	if t.Name == "\n" {
		(*l).line++
		(*l).position = 1
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
