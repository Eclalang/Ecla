package v2

type TokenTypeCompositeBehavior struct {
	Name string
}

func (t *TokenTypeCompositeBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve composite")
	if l.ret[len(l.ret)-1].TokenType == t.Name {
		l.ComposeToken(t.Name)
		l.prevIndex = l.index
		l.tempVal = ""
	}
}
func (t *TokenTypeCompositeBehavior) Get() []string {
	return append([]string{}, t.Name)
}
func (t *TokenTypeCompositeBehavior) InvolvedWith() []ITokenType {
	return []ITokenType{}
}

var (
	CINT = TokenTypeCompositeBehavior{
		Name: INT,
	}
	CFLOAT = TokenTypeCompositeBehavior{
		Name: FLOAT,
	}
	CADDASSIGN = TokenTypeCompositeBehavior{
		Name: ADD + ASSIGN,
	}
	CSUBASSIGN = TokenTypeCompositeBehavior{
		Name: SUB + ASSIGN,
	}
	CMODASSIGN = TokenTypeCompositeBehavior{
		Name: MOD + ASSIGN,
	}
	CDIVASSIGN = TokenTypeCompositeBehavior{
		Name: DIV + ASSIGN,
	}
	CMULTASSIGN = TokenTypeCompositeBehavior{
		Name: MULT + ASSIGN,
	}
	CINC = TokenTypeCompositeBehavior{
		Name: INC,
	}
	CDEC = TokenTypeCompositeBehavior{
		Name: DEC,
	}
	CEQUAL = TokenTypeCompositeBehavior{
		Name: EQUAL,
	}
	CNEQ = TokenTypeCompositeBehavior{
		Name: NEQ,
	}
	CQOT = TokenTypeCompositeBehavior{
		Name: QOT,
	}
	CLEQ = TokenTypeCompositeBehavior{
		Name: LEQ,
	}
	CGEQ = TokenTypeCompositeBehavior{
		Name: GEQ,
	}
	CSTRING = TokenTypeCompositeBehavior{
		Name: STRING,
	}
	CCHAR = TokenTypeCompositeBehavior{
		Name: CHAR,
	}
	CCOMMENT = TokenTypeCompositeBehavior{
		Name: COMMENT,
	}
	CXOR = TokenTypeCompositeBehavior{
		Name: XOR,
	}
	CQOTASSIGN = TokenTypeCompositeBehavior{
		Name: QOT + ASSIGN,
	}
	CCOMMENTGROUP = TokenTypeCompositeBehavior{
		Name: COMMENTGROUP,
	}
	CCOMMENTGROUPEND = TokenTypeCompositeBehavior{
		Name: COMMENTGROUP + "END",
	}
)
