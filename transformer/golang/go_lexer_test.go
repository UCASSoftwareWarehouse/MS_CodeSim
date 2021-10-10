package golang

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestNewGoLexer(t *testing.T) {
	input, _ := antlr.NewFileStream("/Users/purchaser/go/src/code_sim/config/init.go")
	lexer := NewGoLexer(input)
	tokens := lexer.GetAllTokens()
	for _, token := range tokens {
		t.Log(token.GetText())
	}
}
