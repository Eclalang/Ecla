package lexer

import "strconv"

type ReadingType int

const (
	INLINE = iota
	REVERSED
	NOTFOUND
)

type TLexer struct {
	// ret is the []Token that the lexer will return
	ret           []Token
	lastStepToken ITokenType

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
	TriggerBy string

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
		l.DEBUGLEXER("each step " + strconv.Itoa(l.index))
		switch l.stepFind {
		case NOTFOUND:

		case INLINE:
			l.FindSyntax()
			l.indent[0].Resolve(l)
			l.lastStepToken = l.indent[0]

		case REVERSED:
			if l.Inquote() {
				l.FindSyntax()
			} else {
				l.FindSyntax()
				temp := l.tempVal[len(l.tempVal)-l.sizeOfTokenReversed:]
				l.tempVal = l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
				l.position -= 1
				l.AddToken(TEXT)
				l.position += 1
				l.tempVal = temp
			}
			l.indent[0].Resolve(l)
			l.lastStepToken = l.indent[0]
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
	l.DEBUGLEXER("before add " + TokenType)
	var tmp = Token{
		TokenType: TokenType,
		Value:     l.tempVal,
		Position:  l.position - len(l.tempVal),
		Line:      l.line,
	}
	l.ret = append(l.ret, tmp)

}

func (l *TLexer) ComposeToken(NewName string) {
	l.DEBUGLEXER("before compose " + NewName)
	l.ret[len(l.ret)-1].TokenType = NewName
	l.ret[len(l.ret)-1].Value += l.tempVal

}

func (l *TLexer) Inquote() bool {
	return l.TriggerBy != ""
}

var (
	Lex = &TLexer{
		ret:           []Token{},
		lastStepToken: &TokenTypeBaseBehavior{Name: ""},
		index:         0,
		sentence:      "",
		tempVal:       "",

		prevIndex:   0,
		actualIndex: -1,

		line:     1,
		position: 1,

		canBeText: false,
		isSpaces:  false,

		TriggerBy: "",

		allTokenType:        Every,
		maxLen:              -1,
		sizeOfTokenReversed: -1,
	}
)

func LexerR(sentence string) []Token {
	println("\n---------------------\n-----PRINT DEBUG-----\n---------------------\n")
	Lex.SetSentence(sentence)
	Lex.Step()
	Lex.tempVal = Lex.sentence[Lex.prevIndex:Lex.index]
	if Lex.tempVal != "" {

		Lex.AddToken(TEXT)

		Lex.tempVal = ""
	}
	Lex.AddToken(EOF)
	println("\n---------------------\n---FIN PRINT DEBUG---\n---------------------")
	return Lex.Ret()
}

func (l *TLexer) DEBUGLEXER(s string) {
	//println("\n-------------"+s+"-------------\nl.tempVal\t\t:", "\""+l.tempVal+"\"")
	//println("l.TriggerBy\t\t:", l.TriggerBy)
	//if (len(l.indent) - 1) >= 0 {
	//	println("l.indent name\t\t:", l.indent[0].Get()[len(l.indent[0].Get())-1])
	//} else {
	//	println("l.indent name\t\t: None")
	//}
	//
	//println("l.lastSTepToken\t\t:", l.lastStepToken.Get()[len(l.lastStepToken.Get())-1])
	//println("l.isSpaces\t\t:", l.isSpaces)
	//println("l.index\t\t\t:", l.index)
	//println("l.prevIndex\t\t:", l.prevIndex)
	//println("l.position\t\t:", l.position)
	//println("l.line\t\t\t:", l.line)
	//println("l.sizeOfTokenReversed\t:", l.sizeOfTokenReversed)
	//println("l.sentence\t\t:", l.sentence)
	//println("l.sentence readed\t:", l.sentence[:l.prevIndex]+"|")

}
