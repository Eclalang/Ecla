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

func TestHashtag0(t *testing.T) {
	tLexer(t, testHashtag0, " testHashtag")
}

func TestHashtag1(t *testing.T) {
	tLexer(t, testHashtag1, " testHashtag")
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

func TestHashtag6(t *testing.T) {
	tLexer(t, testHashtag6, " testHashtag6")
}

func TestMurloc(t *testing.T) {
	tLexer(t, testMurloc, " testMurloc")
}

func TestSpeChar(t *testing.T) {
	tLexer(t, testSpeChar, " testSpeChar")
}

func TestChar(t *testing.T) {
	tLexer(t, testCHAR, "testCHar")
}

func TestCharString(t *testing.T) {
	tLexer(t, testCHARSTRING, "testCharString")
}

func TestCharString2(t *testing.T) {
	tLexer(t, testCHARSTRING2, "testCharString2")
}

func TestMultiLigneString(t *testing.T) {
	tLexer(t, testMultiLigneString, "testMultiLigneString")
}

func TestCondInParen(t *testing.T) {
	tLexer(t, testCondInParen, "testCondInParen")
}

func TestCommentGroup(t *testing.T) {
	tLexer(t, testCommentGroup, "testCommentGroup")
}

func TestEmptyString(t *testing.T) {
	tLexer(t, testEmptyString, "testEmptyString")
}

func TestStringWithBSlash(t *testing.T) {
	tLexer(t, testStringWithBSlash, "testStringWithBSlash")
}

func TestBSlashstring(t *testing.T) {
	tLexer(t, testBSlashstring, "testBSlashstring")
}
