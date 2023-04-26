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
	var endOfComm int = -1
	var endOfCommGroup int = -1

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
			// -----------Is Known Syntaxe Part-------------
			//
			// for each element of Identifier, we compare all the known
			// syntaxes with our tempVal, If the comparison is true,
			// tempVal is a known syntaxes, and then a token
			if ident.IsSyntaxe(tempVal) {
				// canot be a text now
				canBeText = false
				if (ident.Identifier == COMMENT || ident.Identifier == COMMENTGROUP) && inQuote && len(ret) != 0 {
					break
				}
				// -----------Previous Token COMMENT Part-------------
				if len(ret) > 1 {
					// if the previous token is a COMMENT, we must concat the actual value to the previous
					// token instead of create a new one.
					// we don't concat only if it's the end of the COMMENT token.
					// if we stop concat the value with the COMMENT, we keep the COMMENT token index,
					// like that we can be sure to not append something else in.
					if ret[len(ret)-1].TokenType == COMMENT && len(ret)-1 != endOfComm && !inQuote {
						if tempVal == "/" && ret[len(ret)-1].Value == "" {
							ret[len(ret)-1].TokenType = COMMENTGROUP
							break
						} else if tempVal != "\n" && tempVal != "\r" {
							ret[len(ret)-1].Value += tempVal
						} else {
							endOfComm = len(ret) - 1
						}
						tempVal = ""
						prevIndex = i
						break
					}
					// same things for the COMMENTGROUP, but with a different ending close.
					if ret[len(ret)-1].TokenType == COMMENTGROUP && len(ret)-1 != endOfCommGroup && !inQuote {
						println("in com")
						if tempVal[len(tempVal)-1] == '/' {
							break
						} else if len(tempVal) > 1 {
							if tempVal[len(tempVal)-2:] == "/#" {
								endOfCommGroup = len(ret) - 1
							} else {
								ret[len(ret)-1].Value += tempVal
							}
						} else {
							ret[len(ret)-1].Value += tempVal
						}
						tempVal = ""
						prevIndex = i
						break
					}
				}
				// -----------Previous Token COMMENT Part END-------------
				// -----------Quote Token Part-------------
				// we ignore the "spaces" TOKEN, wo include " ","\r","\n" and more to
				// only be include in string TOKEN
				if ident.Identifier == "" && !inQuote {
					isSpaces = true
					prevIndex = i
					tempVal = sentence[prevIndex:i]
					break
				}
				if ident.Identifier == DQUOTE {
					// if the current lecture head in inside of a string, we we be carefull about \", cause
					// it does not end the current string.
					// if we have a token DQUOTE without being in a string, its the start of a new string
					if inQuote {
						if ret[len(ret)-1].Value[len(ret[len(ret)-1].Value)-1] != '\\' {
							inQuote = false
						}
					} else {
						inQuote = true
						inQuoteStep = true
					}
				}
				// -----------Quote Token Part END-------------

				// -----------Special Token Part-------------
				//
				// we change our behavior in case of some special TOKEN, like ASSIGN, cause he can be a part of
				// a complexe token (ADDASSIGN for exemple)
				beforeChangeVal := tempVal
				ret = tokenCommentGroup(ident, ret, &prevIndex, &tempVal, i)
				ret = tokenComment(ident, ret, &prevIndex, &tempVal, i, line, sentence)
				ret = tokenInt(ident, ret, &prevIndex, &tempVal, i, isSpaces)
				ret = tokenPeriod(ident, ret, &prevIndex, &tempVal, i, isSpaces, inQuote, sentence)
				ret = tokenAssign(ident, ret, &prevIndex, &tempVal, i)
				ret = tokenDiv(ident, ret, &prevIndex, &tempVal, i)
				ret = tokenAddSub(ident.Identifier == ADD, ret, &prevIndex, &tempVal, i, ADD, INC)
				ret = tokenAddSub(ident.Identifier == SUB, ret, &prevIndex, &tempVal, i, SUB, DEC)
				if beforeChangeVal != tempVal {
					break
				}
				// ---------Special Token Part END-----------

				// ---------Normal Token Part END-----------
				//
				// append a new Token to the variable ret
				ret = inQuoteChange(inQuote && !inQuoteStep, ret, ident, tempVal, prevIndex, sentence)

				isSpaces = false

				tempVal = ""
				prevIndex = i
				// ---------Normal Token Part END-----------
			}
			// -----------Is Known Syntaxe Part END-------------

			// -----------Previous Token COMMENT Part Again-------------
			//
			// we must check again if the previous one is a COMMENT in case its a text, or the end of the lexing.
			if len(ret) >= 1 {
				if ret[len(ret)-1].TokenType == COMMENT && len(ret)-1 != endOfComm && !inQuote {
					if tempVal == "/" && ret[len(ret)-1].Value == "" {
						ret[len(ret)-1].TokenType = COMMENTGROUP
					} else if tempVal != "\n" && tempVal != "\r" {
						ret[len(ret)-1].Value += tempVal
					} else {
						endOfComm = len(ret) - 1
					}
					canBeText = false
					tempVal = ""
					prevIndex = i
					break
				}
				if ret[len(ret)-1].TokenType == COMMENTGROUP && len(ret)-1 != endOfCommGroup && !inQuote {
					canBeText = false
					if tempVal[len(tempVal)-1] == '/' {
						break
					} else if len(tempVal) > 1 {
						if tempVal[len(tempVal)-2:] == "/#" {
							endOfCommGroup = len(ret) - 1
						} else {
							ret[len(ret)-1].Value += tempVal
						}
					} else {
						ret[len(ret)-1].Value += tempVal
					}
					tempVal = ""
					prevIndex = i
					break
				}
			}
			// -----------Previous Token COMMENT Part Again END-------------
		}
		// -----------Can still be text Part-------------
		//
		// if after checking all the known syntaxe, the tempValue can still
		// be a TEXT, we parse the tempValue backward to verifies if
		// a substring of tempValue can also be a known syntaxe
		if canBeText {
			for y := len(tempVal) - 1; y >= 0; y-- {
				for _, ident := range Identifier {
					if ident.Identifier != INT {
						if ident.IsSyntaxe(tempVal[y:]) {
							canBeText = false
							ret = inQuoteChange(inQuote && !inQuoteStep, ret, Identifier[0], tempVal[:y], prevIndex, sentence)
							i += y - len(tempVal)
							isSpaces = false
							prevIndex = i
						}
					}

				}
			}
		}
		// -----------Can still be text Part END-------------
	}
	// -----------End of lexer Part-------------
	//
	// if at the end of the sentence parse, tempVal is not "", it means that
	// a last token of type TEXT must be appended to the return value
	if len(ret) > 0 {
		if ret[len(ret)-1].TokenType == COMMENTGROUP && tempVal == "/" {
			ret[len(ret)-1].Value += tempVal
		} else if tempVal != "" {
			actualIndex, line = positionDetector(prevIndex, sentence)
			ret = append(ret, addToken(Identifier[0].Identifier, tempVal, actualIndex, line))

			prevIndex += len(tempVal)
		}
	} else if tempVal != "" {
		actualIndex, line = positionDetector(prevIndex, sentence)
		ret = append(ret, addToken(Identifier[0].Identifier, tempVal, actualIndex, line))

		prevIndex += len(tempVal)
	}

	// created a last token of type EOF (EndOfFile)
	actualIndex, line = positionDetector(prevIndex, sentence)
	ret = append(ret, addToken(Identifier[len(Identifier)-1].Identifier, "", actualIndex, line))

	return ret
	// -----------End of lexer Part END-------------
}

func inQuoteChange(inQuote bool, ret []Token, identi identifier, val string, prevIndex int, sentence string) []Token {
	actualIndex, line := positionDetector(prevIndex, sentence)
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

// positionDetector find the current position and line of the token we want to create.
// Take our current []Token, the index of lexing, and the global sentence.
//
// return position, line
func positionDetector(prevIndex int, sentence string) (int, int) {
	var toRet = 0
	var line = 1
	for _, v := range sentence[:prevIndex] {
		if v == '\n' {
			toRet = 0
			line += 1
		} else {
			toRet += 1
		}
	}
	return toRet + 1, line
}

// addToken create a new token with the given parameters
//
// return the created token
func addToken(TokenType string, Value string, Position int, Line int) Token {
	var ret Token

	ret.TokenType = TokenType
	ret.Value = Value
	ret.Position = Position
	ret.Line = Line

	return ret
}

// tokenAddSub replace the previous token.tokenType in ret to toReplace if the current token.tokenType is equal to toFind.
//
// return the changed []Token
func tokenAddSub(isIdentifier bool, ret []Token, prevIndex *int, tempVal *string, index int, toFind string, toReplace string) []Token {
	if isIdentifier {
		if len(ret) >= 1 {
			if ret[len(ret)-1].TokenType == toFind {

				ret[len(ret)-1].TokenType = toReplace
				ret[len(ret)-1].Value += *tempVal
				*tempVal = ""
				*prevIndex = index

			}
		}
	}
	return ret
}

// tokenAssign replace the previous token.tokenType in ret to the compose token of the current token and the previous one
// if the current token.tokenType is ASSIGN, and if the previous one is one of the listed token able to merge with an
// ASSIGN.
//
// return the changed []Token
func tokenAssign(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int) []Token {
	if ident.Identifier == ASSIGN {
		if len(ret) >= 1 {
			if concatEqual(ret[len(ret)-1].TokenType) {
				if ret[len(ret)-1].TokenType == ASSIGN {
					ret[len(ret)-1].TokenType = EQUAL
				} else if ret[len(ret)-1].TokenType == ADD || ret[len(ret)-1].TokenType == SUB || ret[len(ret)-1].TokenType == MULT || ret[len(ret)-1].TokenType == DIV || ret[len(ret)-1].TokenType == QOT || ret[len(ret)-1].TokenType == MOD {
					ret[len(ret)-1].TokenType = ret[len(ret)-1].TokenType + ident.Identifier
				} else if ret[len(ret)-1].TokenType == LSS {
					ret[len(ret)-1].TokenType = LEQ
				} else if ret[len(ret)-1].TokenType == GTR {
					ret[len(ret)-1].TokenType = GEQ
				} else if ret[len(ret)-1].TokenType == NOT {
					ret[len(ret)-1].TokenType = NEQ
				}
				ret[len(ret)-1].Value += *tempVal
				*tempVal = ""
				*prevIndex = index
			}
		}
	}
	return ret
}

// tokenDiv replace the previous token.tokenType in ret to the QOT token
// if the current token.tokenType and the previous token.tokenType are DIV
//
// return the changed []Token
func tokenDiv(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int) []Token {
	if ident.Identifier == DIV {
		if len(ret) >= 1 {
			if ret[len(ret)-1].TokenType == DIV {
				ret[len(ret)-1].TokenType = QOT
				ret[len(ret)-1].Value += *tempVal
				*tempVal = ""
				*prevIndex = index
			}
		}
	}
	return ret
}

// TokenPeriod replace the previous token.tokenType in ret to the FLOAT token
// if the current token.tokenType is PERIOD and the previous token.tokenType is INT
//
// if we are in a string, ignore this behavior
//
// return the changed []Token
func tokenPeriod(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int, isSpaces bool, inQuote bool, sentence string) []Token {
	if ident.Identifier == PERIOD {
		if len(ret) >= 1 {
			if ret[len(ret)-1].TokenType == INT && !isSpaces {
				ret[len(ret)-1].Value += *tempVal
				ret[len(ret)-1].TokenType = FLOAT
				*tempVal = ""
				*prevIndex = index
			} else if !inQuote {
				actualIndex, line := positionDetector(*prevIndex, sentence)
				ret = append(ret, addToken(ident.Identifier, *tempVal, actualIndex, line))
				*tempVal = ""
				*prevIndex = index
			}
		}
	}
	return ret
}

// tokenDiv replace the previous token.tokenType in ret to the QOT tokenType
// if the current token.tokenType and the previous token.tokenType are DIV
//
// return the changed []Token
func tokenInt(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int, isSpaces bool) []Token {
	if ident.Identifier == INT {
		if !isSpaces {
			if len(ret) >= 1 {
				if ret[len(ret)-1].TokenType == INT || ret[len(ret)-1].TokenType == TEXT || ret[len(ret)-1].TokenType == FLOAT {
					ret[len(ret)-1].Value += *tempVal
					*tempVal = ""
					*prevIndex = index
				}
			}
		}

	}
	return ret
}

// tokenComment replace the previous token.tokenType in ret to the COMMENTGROUP tokenType
// if the current token.tokenType is COMMENTGROUPIDENT, and of the previous one is COMMENT
//
// return the changed []Token
func tokenCommentGroup(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int) []Token {
	if ident.Identifier == COMMENTGROUPIDENT {
		if len(ret) >= 1 {
			if ret[len(ret)-1].TokenType == COMMENT {
				ret[len(ret)-1].TokenType = COMMENTGROUP
				ret[len(ret)-1].Value += ""
				*tempVal = ""
				*prevIndex = index
			}
		}
	}
	return ret
}

// tokenComment add a new token if the no value if the identifier is COMMENT
//
// return the changed []Token
func tokenComment(ident identifier, ret []Token, prevIndex *int, tempVal *string, index int, line int, sentence string) []Token {
	if ident.Identifier == COMMENT {
		actualIndex, line := positionDetector(*prevIndex, sentence)
		ret = append(ret, addToken(ident.Identifier, "", actualIndex, line))
		*tempVal = ""
		*prevIndex = index
	}
	return ret
}
