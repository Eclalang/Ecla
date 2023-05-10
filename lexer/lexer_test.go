package lexer

import (
	"strconv"
	"testing"
)

func tLexer(t *testing.T, tested testList, name string) {
	result := ""
	code := tested.input
	expected := tested.output
	expectedLenth := len(expected)
	result += "\n--------------------------------------------------\n--------------------------------------------------\n\t\t---" + name + "-INPUT---\n--------------------------------------------------\n" + code + "\n--------------------------------------------------"
	result += "\n\t\t---DIFF LIST---\n--------------------------------------------------\n"
	diff := 0
	l := Lexer(code)
	if l == nil {
		result += "Expected a lexer, got nil\n--------------------------------------------------\n"
	} else if len(l) != expectedLenth {
		result += "Expected " + strconv.Itoa(expectedLenth) + " tokens, got " + strconv.Itoa(len(l))
		diff++
	}
	for Position, expct := range expected {
		if Position < len(l) {
			if expct != l[Position] {
				diff++
				result += "\n--------------------------------------------------\nDiff " + strconv.Itoa(diff) + " Expected {" + expct.TokenType + " " + expct.Value + " " + strconv.Itoa(expct.Position) + " " + strconv.Itoa(expct.Line) + "} for the token n°" + strconv.Itoa(Position+1) + "\n       Got \t{" + l[Position].TokenType + " " + l[Position].Value + " " + strconv.Itoa(l[Position].Position) + " " + strconv.Itoa(l[Position].Line) + "}\n--------------------------------------------------\n"
			}
		}
	}
	if diff == 0 {
		result += "\t      ---AUCUNE ERREUR---\n--------------------------------------------------\n"
		t.Log(result)
	} else {
		result += "\n--------------------------------------------------\n"
		t.Error(result, "\ngot :\n", l)
	}
}

//func TestHashtag5(t *testing.T) {
//	tLexer(t, testHashtag5, " testHashtag5")
//}

func TestNoFile(t *testing.T) {
	tLexer(t, testNoFile, " testNoFile")
}

func TestCalc(t *testing.T) {
	tLexer(t, testCalc, " testCalc")
}

func TestBoolOpperand(t *testing.T) {
	tLexer(t, testBoolOpperand, " testBoolOpperand")
}

func TestEOL(t *testing.T) {
	tLexer(t, testEOL, "testEOL")
}

func TestHashtag(t *testing.T) {
	tLexer(t, testHashtag, " testHashtag")
}

func TestHashtag2(t *testing.T) {
	tLexer(t, testHashtag2, " testHashtag2")
}

func TestHashtag3(t *testing.T) {
	tLexer(t, testHashtag3, " testHashtag3")
}

func TestHashtag4(t *testing.T) {
	tLexer(t, testHashtag4, " testHashtag4")
}

func TestMurloc(t *testing.T) {
	tLexer(t, testMurloc, " testMurloc")
}

func TestSpeChar(t *testing.T) {
	tLexer(t, testSpeChar, " testSpeChar")
}

func TestPositionDetector_1(t *testing.T) {
	var Expected1 = 1 //position
	var Expected2 = 1 //line

	var sentence = "test"
	var previndex = 0

	var got1, got2 = positionDetector(previndex, sentence)

	var errorString = "\n"

	if got1 != Expected1 {
		errorString += "Failed at getting Position\n expected\t: " + strconv.Itoa(Expected1) + "\n got \t\t: " + strconv.Itoa(got1) + "\n"
	}
	if got2 != Expected2 {
		errorString += "Failed at getting Line\n expected\t: " + strconv.Itoa(Expected2) + "\n got \t\t: " + strconv.Itoa(got2) + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestPositionDetector_2(t *testing.T) {
	var Expected1 = 1 //position
	var Expected2 = 2 //line

	var sentence = "test\ntest"
	var previndex = 5

	var got1, got2 = positionDetector(previndex, sentence)

	var errorString = "\n"

	if got1 != Expected1 {
		errorString += "Failed at getting Position\n expected\t: " + strconv.Itoa(Expected1) + "\n got \t\t: " + strconv.Itoa(got1) + "\n"
	}
	if got2 != Expected2 {
		errorString += "Failed at getting Line\n expected\t: " + strconv.Itoa(Expected2) + "\n got \t\t: " + strconv.Itoa(got2) + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestPositionDetector_3(t *testing.T) {
	var Expected1 = 3 //position
	var Expected2 = 1 //line

	var sentence = "test\ntest"
	var previndex = 2

	var got1, got2 = positionDetector(previndex, sentence)

	var errorString = "\n"

	if got1 != Expected1 {
		errorString += "Failed at getting Position\n expected\t: " + strconv.Itoa(Expected1) + "\n got \t\t: " + strconv.Itoa(got1) + "\n"
	}
	if got2 != Expected2 {
		errorString += "Failed at getting Line\n expected\t: " + strconv.Itoa(Expected2) + "\n got \t\t: " + strconv.Itoa(got2) + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestAddToken(t *testing.T) {
	var Expected Token = Token{
		TokenType: INT,
		Value:     "42",
		Position:  4,
		Line:      2,
	}

	var got = addToken(INT, "42", 4, 2)

	if got != Expected {
		t.Error("Failed at getting Position\n expected\t: ", Expected, "\n got \t\t: ", got, "\n")
	}

}

func TestTokenAddSub_1(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: INC,
			Value:     "++",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: ADD,
			Value:     "+",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "+"

	var got = tokenAddSub(true, ret, &prevIndex, &tempVal, 2, ADD, INC)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to INC\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestTokenAddSub_2(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: TEXT,
			Value:     "pi",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: TEXT,
			Value:     "pi",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "+"

	var got = tokenAddSub(true, ret, &prevIndex, &tempVal, 2, ADD, INC)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to INC\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 1 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "+" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestTokenAddSub_3(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: DEC,
			Value:     "--",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: SUB,
			Value:     "-",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "-"

	var got = tokenAddSub(true, ret, &prevIndex, &tempVal, 2, SUB, DEC)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to DEC\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestTokenAssign_EQUAL(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: EQUAL,
			Value:     "==",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: ASSIGN,
			Value:     "=",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to DEC\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestTokenAssign_LEQ(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: LEQ,
			Value:     "<=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: LSS,
			Value:     "<",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to DEC\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}
}

func TestTokenAssign_GEQ(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: GEQ,
			Value:     ">=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: GTR,
			Value:     ">",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to GEQ\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_ADDASSIGN(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: ADD + ASSIGN,
			Value:     "+=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: ADD,
			Value:     "+",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to ADDASSIGN\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_SUBASSIGN(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: SUB + ASSIGN,
			Value:     "-=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: SUB,
			Value:     "-",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to SUBASSIGN\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_MULTASSIGN(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: MULT + ASSIGN,
			Value:     "*=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: MULT,
			Value:     "*",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to MULTASSIGN\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_DIVASSIGN(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: DIV + ASSIGN,
			Value:     "/=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: DIV,
			Value:     "/",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to DIVASSIGN\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_MODASSIGN(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: MOD + ASSIGN,
			Value:     "%=",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: MOD,
			Value:     "%",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "="

	var got = tokenAssign(identifier{Identifier: ASSIGN, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at changing to MODASSIGN\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenAssign_BAD(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: MOD,
			Value:     "%",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: MOD,
			Value:     "%",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "+"

	var got = tokenAssign(identifier{Identifier: ADD, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 1 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "+" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenDiv_OK(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: QOT,
			Value:     "//",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: DIV,
			Value:     "/",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "/"

	var got = tokenDiv(identifier{Identifier: DIV, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at modifying  to MOD\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestTokenDiv_BAD(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: DIV,
			Value:     "/",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: DIV,
			Value:     "/",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "+"

	var got = tokenDiv(identifier{Identifier: ADD, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2)

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 1 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(1) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "+" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestPeriod_FLOAT(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: FLOAT,
			Value:     "1.",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: INT,
			Value:     "1",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 1
	var tempVal = "."

	var got = tokenPeriod(identifier{Identifier: PERIOD, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 2, false, false, "1.")

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed at modifying  to FLOAT\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 2 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}

func TestPeriod_BAD(t *testing.T) {
	var Expected = []Token{
		Token{
			TokenType: INT,
			Value:     "1",
			Position:  1,
			Line:      1,
		},
	}

	var ret = []Token{
		Token{
			TokenType: INT,
			Value:     "1",
			Position:  1,
			Line:      1,
		},
	}
	var prevIndex = 2
	var tempVal = "."

	var got = tokenPeriod(identifier{Identifier: PERIOD, Syntaxe: []string{}}, ret, &prevIndex, &tempVal, 3, true, false, "1 .")

	var errorString = "\n"

	if got[0] != Expected[0] {
		errorString += "Failed\n expected\t: " + Expected[0].TokenType + "\n got \t\t: " + got[0].TokenType + "\n"
	}
	if prevIndex != 3 {
		errorString += "Failed at modifying pos\n expected\t: " + strconv.Itoa(2) + "\n got \t\t: " + strconv.Itoa(prevIndex) + "\n"
	}
	if tempVal != "" {
		errorString += "Failed at modifying val\n expected\t: " + "\n got \t\t: " + tempVal + "\n"
	}
	if len(errorString) != 1 {
		t.Error(errorString)
	}

}
