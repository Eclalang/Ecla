package lexer

import (
	"strconv"
	"testing"

	"github.com/tot0p/Ecla/lexer"
)

func TestLexer(t *testing.T) {
	var GlobalTest = []testList{
		testCalc,
		testDQuote,
		testSpeChar,
		testEOL,
	}
	for num, tested := range GlobalTest {

		code := tested.input
		expected := tested.output
		expected_lenth := len(expected)
		t.Log("\n\t\t---TEST" + strconv.Itoa(num) + "-INPUT---\n\t\t---\"" + code + "\"---")
		t.Log("\n\t\t---DIFF LIST---")
		diff := 0
		l := lexer.Lexer(code)
		if l == nil {
			t.Error("Expected a lexer, got nil")
		} else if len(l) != expected_lenth {
			t.Error("Expected "+strconv.Itoa(expected_lenth)+" tokens, got ", len(l))
			diff++
		}
		for Position, expct := range expected {
			if Position < len(l) {
				if expct != l[Position] {
					diff++
					t.Error("Diff ", diff, " Expected ", expct, " for the token n°", Position+1, " , got ", l[Position])
				}
			}
		}
		if diff == 0 {
			t.Log("\n\t\t\t---AUCUNE ERREUR---")
		}
		t.Log("\n\t\t---GLOBAL RESULT---\n")
		t.Log("Expected \t", expected)
		t.Log("Got \t", l)
		t.Log("diff total ", diff)
	}
}
