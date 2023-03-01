package lexer

import (
	"strconv"
	"testing"

	"github.com/tot0p/Ecla/lexer"
)

func tLexer(t *testing.T, tested testList, name string) {
	result := ""
	code := tested.input
	expected := tested.output
	expected_lenth := len(expected)
	result += "\n--------------------------------------------------\n--------------------------------------------------\n\t\t---" + name + "-INPUT---\n--------------------------------------------------\n" + code + "\n--------------------------------------------------"
	result += "\n\t\t---DIFF LIST---\n--------------------------------------------------\n"
	diff := 0
	l := lexer.Lexer(code)
	if l == nil {
		result += "Expected a lexer, got nil\n--------------------------------------------------\n"
	} else if len(l) != expected_lenth {
		result += "Expected " + strconv.Itoa(expected_lenth) + " tokens, got " + strconv.Itoa(len(l))
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
