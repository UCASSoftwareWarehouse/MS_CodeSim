package python

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestNewPython3Lexer(t *testing.T) {
	input, _ := antlr.NewFileStream("/Users/purchaser/PycharmProjects/A3/main.py")
	lexer := NewPython3Lexer(input)
	tokens := lexer.GetAllTokens()
	for _, token := range tokens {
		t.Log(token.GetText())
	}
}
