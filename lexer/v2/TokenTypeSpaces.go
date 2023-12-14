package v2

type TokenTypeSpacesBehavior struct {
	Name   string
	Syntax []string
}

func (t *TokenTypeSpacesBehavior) Resolve(l *TLexer) {
	println("in Resolve Spaces...")
	if (*l).Inquote() {
		println("t.name---", t.Name, "-")
	} else {
		println("t.name--- BACKSPACES", t.Name)
		(*l).isSpaces = true
	}
	if t.Name == "\n" {
		(*l).line++
		(*l).position = 0
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
			"_",
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
