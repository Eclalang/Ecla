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
		testHashtag,
	}
	for num, tested := range GlobalTest {
		result := ""
		code := tested.input
		expected := tested.output
		expected_lenth := len(expected)
		result += "\n--------------------------------------------------\n--------------------------------------------------\n\t\t---TEST" + strconv.Itoa(num) + "-INPUT---\n--------------------------------------------------\n" + code + "\n--------------------------------------------------"
		result += "\n\t\t---DIFF LIST---"
		diff := 0
		l := lexer.Lexer(code)
		if l == nil {
			result += "\n--------------------------------------------------\nExpected a lexer, got nil"
		} else if len(l) != expected_lenth {
			result += "\n--------------------------------------------------\nExpected " + strconv.Itoa(expected_lenth) + " tokens, got " + strconv.Itoa(len(l)) + "\n--------------------------------------------------\n"
			diff++
		}
		for Position, expct := range expected {
			if Position < len(l) {
				if expct != l[Position] {
					diff++
					result += "\n--------------------------------------------------\nDiff " + strconv.Itoa(diff) + " Expected {" + expct.TokenType + " " + expct.Value + " " + strconv.Itoa(expct.Line) + " " + strconv.Itoa(expct.Position) + "} for the token n°" + strconv.Itoa(Position+1) + "\n       Got \t{" + l[Position].TokenType + " " + l[Position].Value + " " + strconv.Itoa(l[Position].Line) + " " + strconv.Itoa(l[Position].Position) + "}\n--------------------------------------------------\n"
				}
			}
		}
		if diff == 0 {
			result += "\n\n--------------------------------------------------\n\t      ---AUCUNE ERREUR---\n--------------------------------------------------\n"
			t.Log(result)
		} else {
			result += "\n--------------------------------------------------\n"
			t.Error(result)
			t.Log("\ngot :\n", l)
		}
	}
}
