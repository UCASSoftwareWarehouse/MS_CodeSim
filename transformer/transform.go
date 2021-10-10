package transformer

import (
	"code_sim/config"
	"code_sim/transformer/golang"
	"code_sim/transformer/python"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type CodeType int

const (
	Unknown CodeType = 0
	Python  CodeType = 1
	Golang  CodeType = 2
)

func GetSupportedCodeType(codeFileSuffix string) (CodeType, error) {
	switch codeFileSuffix {
	case "py":
		return Python, nil
	case "go":
		return Golang, nil
	default:
		return -1, nil
	}
}

func Transform(code string, codeType CodeType) (string, error) {
	var tokens []string
	switch codeType {
	case Python:
		tokens = transformPython(code)
	case Golang:
		tokens = transformGolang(code)
	default:
		return "", fmt.Errorf("unsupported codeType %v", codeType)
	}
	return strings.Join(tokens, config.Conf.TransformCodeSplitter), nil
}

func collectTokens(tokens []antlr.Token) []string {
	s := make([]string, 0, 8)
	for _, token := range tokens {
		s = append(s, token.GetText())
	}
	return s
}

func transformPython(code string) []string {
	input := antlr.NewInputStream(code)
	lexer := python.NewPython3Lexer(input)
	tokens := lexer.GetAllTokens()
	return collectTokens(tokens)
}

func transformGolang(code string) []string {
	input := antlr.NewInputStream(code)
	lexer := golang.NewGoLexer(input)
	tokens := lexer.GetAllTokens()
	return collectTokens(tokens)
}
