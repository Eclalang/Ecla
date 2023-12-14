package v2

type ReadingType int

const (
	INLINE ReadingType = iota
	REVERSED
	NOTFOUND
)

type TLexer struct {
	// ret is the []Token that the lexer will return
	ret []Token

	index  int
	indent []ITokenType

	stepFind ReadingType
	sentence string
	tempVal  string
	// prevIndex is actualIndex of the start of the element that we want to compare
	// with the known syntaxe
	prevIndex   int
	actualIndex int
	// line will be increase each time a ";" is founded
	line     int
	position int
	// canBeText is false when an element is already considered as a known
	// syntaxe, and true elsewhere
	canBeText bool
	isSpaces  bool

	inTriggerof IActionToken

	allTokenType        []ITokenType
	maxLen              int
	sizeOfTokenReversed int
}

func (l *TLexer) SetSentence(sentence string) {
	l.sentence = sentence
}

func (l *TLexer) Ret() []Token {
	return l.ret
}

func (l *TLexer) Step() {

	if l.index < len(l.sentence) {
		l.position++
		l.index++
		l.tempVal = l.sentence[l.prevIndex:l.index]
		l.stepFind = l.IsSyntax()
		println("TempVal---", l.tempVal, "-")
		switch l.stepFind {
		case NOTFOUND:

		case INLINE:
			l.FindSyntax()
			println(l.index, " l.indent[0].Get()-----", l.indent[0].Get()[0])
			l.indent[0].Resolve(l)
			l.prevIndex = l.index

		case REVERSED:
			if l.Inquote() {
				l.FindSyntax()
			} else {
				l.FindSyntax()
				temp := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
				l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
				l.AddToken(TEXT)
				l.tempVal = temp
			}
			l.indent[0].Resolve(l)
			l.prevIndex = l.index
		}
		l.indent = []ITokenType{}
		l.sizeOfTokenReversed = -1
		l.Step()
	}

}

func (l *TLexer) GetTempVal() string {
	return l.tempVal
}

func (l *TLexer) IsSyntax() ReadingType {
	for _, elem := range l.allTokenType {
		if l.maxLen != -1 && len(l.tempVal) > l.maxLen {
			break
		}
		for i := 0; i < len(elem.Get())-1; i++ {
			if l.maxLen < len(elem.Get()[i]) {
				l.maxLen = len(elem.Get()[i])
			}
			if l.tempVal == elem.Get()[i] {
				return INLINE
			}
		}
	}
	for y := len(l.tempVal) - 1; y >= len(l.tempVal)-l.maxLen && y >= 0; y-- {
		for _, elem := range l.allTokenType {
			for i := 0; i < len(elem.Get())-1; i++ {
				if l.maxLen < len(elem.Get()[i]) {
					l.maxLen = len(elem.Get()[i])
				}
				if l.tempVal[y:] == elem.Get()[i] {
					return REVERSED
				}
			}
		}
	}

	return NOTFOUND
}

func (l *TLexer) FindSyntax() {
	switch l.stepFind {
	case INLINE:
		for _, elem := range l.allTokenType {
			if l.maxLen != -1 && len(l.tempVal) > l.maxLen {
				break
			}
			for i := 0; i < len(elem.Get())-1; i++ {
				if l.maxLen < len(elem.Get()[i]) {
					l.maxLen = len(elem.Get()[i])
				}
				if l.tempVal == elem.Get()[i] {
					l.indent = append(l.indent, elem)
				}
			}
		}
	case REVERSED:
		for y := len(l.tempVal) - 1; y >= len(l.tempVal)-l.maxLen && y >= 0; y-- {
			for _, elem := range l.allTokenType {
				for i := 0; i < len(elem.Get())-1; i++ {
					if l.maxLen < len(elem.Get()[i]) {
						l.maxLen = len(elem.Get()[i])
					}
					if l.tempVal[y:] == elem.Get()[i] {
						l.sizeOfTokenReversed = len(elem.Get()[i])
						l.indent = append(l.indent, elem)
					}
				}
			}
		}
	}
}

func (l *TLexer) AddToken(TokenType string) {
	var tmp = Token{
		TokenType: TokenType,
		Value:     l.tempVal,
		Position:  l.position,
		Line:      l.line,
	}
	l.ret = append(l.ret, tmp)

}

func (l *TLexer) ComposeToken(NewName string) {
	l.ret[len(l.ret)-1].TokenType = NewName
	l.ret[len(l.ret)-1].Value += l.tempVal

}

func (l *TLexer) Inquote() bool {
	return l.inTriggerof.GetITokenType().Get()[len(l.inTriggerof.GetITokenType().Get())-1] != ""

}

var (
	Lex = &TLexer{
		ret:      []Token{},
		index:    0,
		sentence: "",
		tempVal:  "",

		prevIndex:   0,
		actualIndex: -1,

		line:     1,
		position: 0,

		canBeText: false,
		isSpaces:  false,

		inTriggerof: &TokenTypeTriggerBehavior{
			Name: "",
		},

		allTokenType:        Every,
		maxLen:              -1,
		sizeOfTokenReversed: -1,
	}
)

func LexerR(sentence string) []Token {
	println("\n---------------------\n-----PRINT DEBUG-----\n---------------------\n")
	Lex.SetSentence(sentence)
	Lex.Step()
	if Lex.tempVal != "" && Lex.inTriggerof.Name != "" {
		Lex.AddToken(Lex.inTriggerof.GetITokenType().InvolvedWith().Result[0].Name)
	}
	Lex.tempVal = ""
	Lex.AddToken(EOF)
	println("\n---------------------\n---FIN PRINT DEBUG---\n---------------------")
	return Lex.Ret()
}
