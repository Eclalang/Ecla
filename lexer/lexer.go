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
	var actualIndex int = 1
	// line will be increase each time a ";" is founded
	var line int = 0
	// canBeText is false when an element is already considered as a known
	// syntaxe, and true elsewhere
	var canBeText bool
	var isSpaces bool
	var inQuote bool
	var inQuoteStep bool

	// tempVal is the current element that we want to compare with the known
	// syntaxe
	var tempVal string

	for i := 0; i <= len(sentence); i++ {
		// we assign tempVal as an element in the interval [prevIndex:i]
		tempVal = sentence[prevIndex:i]

		// we assign canBeText to true, because we actually don't know if the
		// current element is a text or not
		canBeText = true
		inQuoteStep = false

		for _, ident := range Identifier {
			// for each element of Identifier, we compare all the known
			// syntaxes with our tempVal, If the comparison is true,
			// tempVal is a known syntaxes, and then a token
			if ident.IsSyntaxe(tempVal) {
				canBeText = false

				if ident.Identifier == "" && !inQuote {
					isSpaces = true
					prevIndex = i
					tempVal = sentence[prevIndex:i]
					break
				}

				if ident.Identifier == DQUOTE {
					if inQuote {
						if ret[len(ret)-1].Value[len(ret[len(ret)-1].Value)-1] != '\\' {
							inQuote = false
						}
					} else {
						inQuote = true
						inQuoteStep = true
					}
				}

				// if the type of the known syntaxe is INT, we want to
				// concat each subsequent INT to the same token
				if ident.Identifier == INT {
					if !isSpaces {
						if len(ret) >= 1 {
							if ret[len(ret)-1].TokenType == INT || ret[len(ret)-1].TokenType == TEXT || ret[len(ret)-1].TokenType == FLOAT {
								ret[len(ret)-1].Value += tempVal
								tempVal = ""
								prevIndex = i
								break
							}
						}
					}
					// if the type is ASSIGN, we want to concat it with all the other
					// mathematic operand place just before ( ex: "+=", "==")
				} else if ident.Identifier == ASSIGN {
					if len(ret) >= 1 {
						if concatEqual(ret[len(ret)-1].TokenType) {
							if ret[len(ret)-1].TokenType == ASSIGN {
								ret[len(ret)-1].TokenType = EQUAL
							} else if ret[len(ret)-1].TokenType == ADD {
								ret[len(ret)-1].TokenType = INC
							} else if ret[len(ret)-1].TokenType == SUB {
								ret[len(ret)-1].TokenType = DEC
							} else if ret[len(ret)-1].TokenType == LSS {
								ret[len(ret)-1].TokenType = LEQ
							} else if ret[len(ret)-1].TokenType == GTR {
								ret[len(ret)-1].TokenType = GEQ
							} else if ret[len(ret)-1].TokenType == NOT {
								ret[len(ret)-1].TokenType = NEQ
							}
							ret[len(ret)-1].Value += tempVal
							tempVal = ""
							prevIndex = i
							break
						}
					}
				} else if ident.Identifier == PERIOD {
					if len(ret) >= 1 {
						if ret[len(ret)-1].TokenType == INT && !isSpaces {
							ret[len(ret)-1].Value += tempVal
							ret[len(ret)-1].TokenType = FLOAT
							tempVal = ""
							prevIndex = i
							break
						} else if !inQuote {
							actualIndex = positionDetector(ret, prevIndex)
							ret = append(ret, addToken(ident.Identifier, tempVal, actualIndex, line))
							tempVal = ""
							prevIndex = i
							break
						}
					}
				} else if ident.Identifier == SUB {
					if len(ret) >= 1 {
						if ret[len(ret)-1].TokenType == SUB {
							if len(ret) >= 2 {
								if ret[len(ret)-2].TokenType == TEXT {
									ret[len(ret)-1].TokenType = DEC
									ret[len(ret)-1].Value += tempVal
									tempVal = ""
									prevIndex = i
									break
								}
							}

						}
					}
				} else if ident.Identifier == ADD {
					if len(ret) >= 1 {
						if ret[len(ret)-1].TokenType == ADD {
							if len(ret) >= 2 {
								if ret[len(ret)-2].TokenType == TEXT {
									ret[len(ret)-1].TokenType = INC
									ret[len(ret)-1].Value += tempVal
									tempVal = ""
									prevIndex = i
									break
								}
							}
						}
					}
				} else if ident.Identifier == DIV {
					if len(ret) >= 1 {
						if ret[len(ret)-1].TokenType == DIV {
							ret[len(ret)-1].TokenType = QOT
							ret[len(ret)-1].Value += tempVal
							tempVal = ""
							prevIndex = i
							break
						}
					}
				}
				// append a new Token to the variable ret
				ret = inQuoteChange(inQuote && !inQuoteStep, ret, ident, tempVal, prevIndex, line)

				isSpaces = false

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
						ret = inQuoteChange(inQuote && !inQuoteStep, ret, Identifier[0], tempVal[:y], prevIndex, line)
						isSpaces = false

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
		actualIndex = positionDetector(ret, prevIndex)
		ret = append(ret, addToken(Identifier[0].Identifier, tempVal, actualIndex, line))

		prevIndex += len(tempVal)

	}

	// created a last token of type EOF (EndOfFile)
	actualIndex = positionDetector(ret, prevIndex)
	ret = append(ret, addToken(Identifier[len(Identifier)-1].Identifier, "", actualIndex, line))

	return ret
}

func inQuoteChange(inQuote bool, ret []Token, identi identifier, val string, prevIndex int, line int) []Token {
	actualIndex := positionDetector(ret, prevIndex)
	if inQuote {
		if len(ret) >= 1 {
			if ret[len(ret)-1].TokenType == STRING {
				ret[len(ret)-1].Value += val
			} else {
				ret = append(ret, addToken(STRING, val, actualIndex, line))
			}
		} else {
			ret = append(ret, addToken(STRING, val, actualIndex, line))
		}
	} else {
		ret = append(ret, addToken(identi.Identifier, val, actualIndex, line))
	}
	return ret
}

func positionDetector(ret []Token, prevIndex int) int {
	for _, v := range ret {
		if v.TokenType == EOL {
			prevIndex -= v.Position
		}
	}
	return prevIndex + 1
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
