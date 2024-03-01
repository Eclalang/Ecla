package lexer

type TokenTypeMergerBehavior struct {
	Name   string
	Syntax []string

	CloseBy   []ITokenType
	Result    []TokenTypeCompositeBehavior
	Involved  []ITokenType
	Composite []TokenTypeCompositeBehavior
}

func (t *TokenTypeMergerBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("IN MERGER")
	index := -1
	if l.TriggerBy == "" {
		if !(*l).isSpaces {
			_, index = t.IsInvolvedWith(l)
		} else {
			(*l).isSpaces = false
		}
		if index == -1 {
			(*l).AddToken(t.Name)
			l.TriggerBy = t.Name
			l.prevIndex = l.index
		} else {
			(*l).ComposeToken(t.Composite[index].Name)
			l.TriggerBy = t.Composite[index].Name
			l.prevIndex = l.index
		}
	} else {
		if !(*l).isSpaces {
			_, index = t.IsInvolvedWith(l)
		} else {
			(*l).isSpaces = false
		}
		println("ok", index)
		if index == -1 {
			if l.index > len(l.sentence) {
				l.ComposeToken(t.Result[0].Name)
				l.prevIndex = l.index
			} else if l.sizeOfTokenReversed != -1 {
				triggerByToken := findNameInMergerTokenType(l.TriggerBy)
				indexOfClose := triggerByToken.IsClosedBySyntaxe(NameFromGet(l.indent[0].Get()))
				println("prout", indexOfClose)
				if indexOfClose != -1 {
					//close , donc doit mettre RESULT+CLOSE en token
					l.FindSyntax()
					temp := l.tempVal[:len(l.tempVal)-l.sizeOfTokenReversed]
					if findNameInEveryTokenType(temp) != nil {
						l.tempVal = temp
						l.ComposeToken(t.Result[0].Get()[len(t.Result[0].Get())-1])
					} else {
						l.ComposeToken(l.TriggerBy)
						l.tempVal = temp
						l.TriggerBy = ""
						if l.indent[0].Get()[len(l.indent[0].Get())-1] == "\n" {
							l.line -= 1
						}
						l.indent[0].Resolve(l)
					}
					l.tempVal = ""
					l.TriggerBy = ""
					l.prevIndex = l.index
				} else {
					if triggerByToken.Name == "NULL" {
						findNameInEveryTokenType(l.TriggerBy).Resolve(l)
					} else {
						l.ComposeToken(l.ret[(len(l.ret))-1].TokenType)
						l.tempVal = ""
						l.prevIndex = l.index
					}

				}
			}
		} else {
			t.Resolve(l)
		}
	}
}
func (t *TokenTypeMergerBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeMergerBehavior) InvolvedWith() []ITokenType {
	return t.CloseBy
}

func (t *TokenTypeMergerBehavior) getResult() []TokenTypeCompositeBehavior {
	return t.Result
}
func (t *TokenTypeMergerBehavior) IsInvolvedWith(lexer *TLexer) (ITokenType, int) {
	for i, token := range t.Involved {
		if len(lexer.Ret())-1 >= 0 {
			if token.Get()[len(token.Get())-1] == lexer.Ret()[len(lexer.Ret())-1].TokenType || (token.Get()[len(token.Get())-1] == "SELF" && t.Name == lexer.Ret()[len(lexer.Ret())-1].TokenType) {
				return token, i
			}
		}
	}
	return nil, -1
}
func (t *TokenTypeMergerBehavior) IsClosedByName(other *TokenTypeMergerBehavior) int {
	for i, tokenType := range t.CloseBy {
		if tokenType.Get()[len(tokenType.Get())-1] == other.Name {
			return i
		}
	}
	return -1
}
func (t *TokenTypeMergerBehavior) IsClosedBySyntaxe(otherName string) int {
	for i, tokenType := range t.CloseBy {
		println(tokenType.Get()[len(tokenType.Get())-1], "ok", otherName)
		if tokenType.Get()[len(tokenType.Get())-1] == SELF.Name {
			if t.Name == otherName {
				return i
			}
		} else {
			if tokenType.Get()[len(tokenType.Get())-1] == otherName {
				return i
			}
		}
	}
	return -1
}
func (t *TokenTypeMergerBehavior) IsInvolvedWithLastStep(lexer *TLexer) (ITokenType, int) {
	for i, token := range lexer.indent[0].InvolvedWith() {
		if token.Get()[len(token.Get())-1] == lexer.lastStepToken.Get()[len(lexer.lastStepToken.Get())-1] {
			return token, i
		}

	}
	return nil, -1
}

var (
	TCOMMENT = TokenTypeMergerBehavior{
		Name: COMMENT,
		Syntax: []string{
			"#",
		},
		CloseBy: []ITokenType{
			&RETURN,
		},
		Result: []TokenTypeCompositeBehavior{
			CCOMMENT,
		},
		Involved: []ITokenType{
			&TokenTypeBaseBehavior{
				Name: DIV,
				Syntax: []string{
					"/",
				},
			},
		},
		Composite: []TokenTypeCompositeBehavior{
			CCOMMENTGROUP,
		},
	}
	TCOMMENTGROUP = TokenTypeMergerBehavior{
		Name: COMMENTGROUP,
		Syntax: []string{
			"/#",
		},
		CloseBy: []ITokenType{
			&CCOMMENTGROUPEND,
		},
		Result: []TokenTypeCompositeBehavior{
			CCOMMENT,
		},
	}
)
