package lexer

// Token is a struct that contains all the information about a token
type Token struct {
	TokenType string
	Value     string
	Position  int
	Line      int
}

// Lexer do a lexical analysis of the string sentence to separate each element,
// and associate each element with a token
func Lexer(sentence string) []Token {

	// ret is the []Token that the lexer will return
	var ret []Token

	// prevIndex is index of the start of the element that we want to compare
	// with the known syntaxe
	var prevIndex int = 0
	// line will be increase each time a ";" is founded
	var line int = 0
	// canBeText is false when an element is already considered as a known
	// syntaxe, and true elsewhere
	var canBeText bool

	// tempVal is the current element that we want to compare with the known
	// syntaxe
	var tempVal string

	for i := 0; i <= len(sentence); i++ {
		// we assign tempVal as an element in the interval [prevIndex:i]
		tempVal = sentence[prevIndex:i]

		// we assign canBeText to true, because we actually don't know if the
		// current element is a text or not
		canBeText = true

		for _, ident := range Identifier {
			// for each element of Identifier, we compare all the known
			// syntaxes with our tempVal, If the comparison is true,
			// tempVal is now syntaxes, and then a token
			if ident.IsSyntaxe(tempVal) {
				canBeText = false
				// if the type of the known syntaxe is INT, we want to
				// concat each subsequent INT to the same token
				if ident.Identifier == INT {
					if len(ret) > 1 {
						if ret[len(ret)-1].TokenType == INT {
							ret[len(ret)-1].Value += tempVal
							tempVal = ""
							prevIndex = i
							break
						}
					}
				}

				// append a new Token to the variable ret
				ret = append(ret, addToken(ident.Identifier, tempVal, prevIndex, line))

				tempVal = ""
				prevIndex = i

				// if the current created token is EOL,
				// Line ++
				if ident.Identifier == EOL {
					line += 1
				}
				break
			}
		}
		// if after checking all the known syntaxe, the tempValue can still
		// be a TEXT, we parse the tempValue backward to verifies if
		// a substring of tempValue can also be a known syntaxe
		if canBeText {
			for y := len(tempVal) - 1; y >= 0; y-- {
				for _, ident := range Identifier {
					if ident.IsSyntaxe(tempVal[y:]) {
						canBeText = false
						ret = append(ret, addToken(Identifier[0].Identifier, tempVal[:y], prevIndex, line))

						i += len(tempVal[y:]) - 2
						prevIndex = i
					}
				}
			}
		}
	}
	// if at the end of the sentence parse, tempVal is not "", it means that
	// a last token of type TEXT must be appended to the return value
	if tempVal != "" {
		ret = append(ret, addToken(Identifier[0].Identifier, tempVal, prevIndex, line))
	}

	// created a last token of type EOF (EndOfFile)
	ret = append(ret, addToken(Identifier[len(Identifier)-1].Identifier, "", prevIndex, line))

	return ret
}

// addToken create a new token with the given parameters
func addToken(TokenType string, Value string, Position int, Line int) Token {
	var ret Token

	ret.TokenType = TokenType
	ret.Value = Value
	ret.Position = Position
	ret.Line = Line

	return ret
}
