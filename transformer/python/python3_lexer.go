// Code generated from Python3.g4 by ANTLR 4.9. DO NOT EDIT.

package python

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 246,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 5, 22, 153, 10, 22, 3,
	22, 3, 22, 5, 22, 157, 10, 22, 3, 22, 5, 22, 160, 10, 22, 3, 23, 3, 23,
	7, 23, 164, 10, 23, 12, 23, 14, 23, 167, 11, 23, 3, 24, 3, 24, 7, 24, 171,
	10, 24, 12, 24, 14, 24, 174, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25,
	180, 10, 25, 12, 25, 14, 25, 183, 11, 25, 3, 25, 6, 25, 186, 10, 25, 13,
	25, 14, 25, 187, 5, 25, 190, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 5,
	32, 207, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 6, 36, 218, 10, 36, 13, 36, 14, 36, 219, 3, 37, 3, 37, 7, 37, 224,
	10, 37, 12, 37, 14, 37, 227, 11, 37, 3, 38, 3, 38, 5, 38, 231, 10, 38,
	3, 38, 5, 38, 234, 10, 38, 3, 38, 3, 38, 5, 38, 238, 10, 38, 3, 39, 5,
	39, 241, 10, 39, 3, 40, 3, 40, 5, 40, 245, 10, 40, 3, 172, 2, 41, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2,
	3, 2, 7, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12,
	14, 15, 5, 2, 67, 92, 97, 97, 99, 124, 2, 254, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3,
	81, 3, 2, 2, 2, 5, 83, 3, 2, 2, 2, 7, 89, 3, 2, 2, 2, 9, 98, 3, 2, 2, 2,
	11, 101, 3, 2, 2, 2, 13, 103, 3, 2, 2, 2, 15, 108, 3, 2, 2, 2, 17, 113,
	3, 2, 2, 2, 19, 119, 3, 2, 2, 2, 21, 125, 3, 2, 2, 2, 23, 127, 3, 2, 2,
	2, 25, 129, 3, 2, 2, 2, 27, 132, 3, 2, 2, 2, 29, 135, 3, 2, 2, 2, 31, 138,
	3, 2, 2, 2, 33, 141, 3, 2, 2, 2, 35, 143, 3, 2, 2, 2, 37, 145, 3, 2, 2,
	2, 39, 147, 3, 2, 2, 2, 41, 149, 3, 2, 2, 2, 43, 156, 3, 2, 2, 2, 45, 161,
	3, 2, 2, 2, 47, 168, 3, 2, 2, 2, 49, 189, 3, 2, 2, 2, 51, 191, 3, 2, 2,
	2, 53, 193, 3, 2, 2, 2, 55, 195, 3, 2, 2, 2, 57, 197, 3, 2, 2, 2, 59, 199,
	3, 2, 2, 2, 61, 201, 3, 2, 2, 2, 63, 206, 3, 2, 2, 2, 65, 210, 3, 2, 2,
	2, 67, 212, 3, 2, 2, 2, 69, 214, 3, 2, 2, 2, 71, 217, 3, 2, 2, 2, 73, 221,
	3, 2, 2, 2, 75, 228, 3, 2, 2, 2, 77, 240, 3, 2, 2, 2, 79, 244, 3, 2, 2,
	2, 81, 82, 7, 63, 2, 2, 82, 4, 3, 2, 2, 2, 83, 84, 7, 100, 2, 2, 84, 85,
	7, 116, 2, 2, 85, 86, 7, 103, 2, 2, 86, 87, 7, 99, 2, 2, 87, 88, 7, 109,
	2, 2, 88, 6, 3, 2, 2, 2, 89, 90, 7, 101, 2, 2, 90, 91, 7, 113, 2, 2, 91,
	92, 7, 112, 2, 2, 92, 93, 7, 118, 2, 2, 93, 94, 7, 107, 2, 2, 94, 95, 7,
	112, 2, 2, 95, 96, 7, 119, 2, 2, 96, 97, 7, 103, 2, 2, 97, 8, 3, 2, 2,
	2, 98, 99, 7, 107, 2, 2, 99, 100, 7, 104, 2, 2, 100, 10, 3, 2, 2, 2, 101,
	102, 7, 60, 2, 2, 102, 12, 3, 2, 2, 2, 103, 104, 7, 103, 2, 2, 104, 105,
	7, 110, 2, 2, 105, 106, 7, 107, 2, 2, 106, 107, 7, 104, 2, 2, 107, 14,
	3, 2, 2, 2, 108, 109, 7, 103, 2, 2, 109, 110, 7, 110, 2, 2, 110, 111, 7,
	117, 2, 2, 111, 112, 7, 103, 2, 2, 112, 16, 3, 2, 2, 2, 113, 114, 7, 121,
	2, 2, 114, 115, 7, 106, 2, 2, 115, 116, 7, 107, 2, 2, 116, 117, 7, 110,
	2, 2, 117, 118, 7, 103, 2, 2, 118, 18, 3, 2, 2, 2, 119, 120, 7, 114, 2,
	2, 120, 121, 7, 116, 2, 2, 121, 122, 7, 107, 2, 2, 122, 123, 7, 112, 2,
	2, 123, 124, 7, 118, 2, 2, 124, 20, 3, 2, 2, 2, 125, 126, 7, 62, 2, 2,
	126, 22, 3, 2, 2, 2, 127, 128, 7, 64, 2, 2, 128, 24, 3, 2, 2, 2, 129, 130,
	7, 63, 2, 2, 130, 131, 7, 63, 2, 2, 131, 26, 3, 2, 2, 2, 132, 133, 7, 64,
	2, 2, 133, 134, 7, 63, 2, 2, 134, 28, 3, 2, 2, 2, 135, 136, 7, 62, 2, 2,
	136, 137, 7, 63, 2, 2, 137, 30, 3, 2, 2, 2, 138, 139, 7, 35, 2, 2, 139,
	140, 7, 63, 2, 2, 140, 32, 3, 2, 2, 2, 141, 142, 7, 45, 2, 2, 142, 34,
	3, 2, 2, 2, 143, 144, 7, 47, 2, 2, 144, 36, 3, 2, 2, 2, 145, 146, 5, 47,
	24, 2, 146, 38, 3, 2, 2, 2, 147, 148, 5, 41, 21, 2, 148, 40, 3, 2, 2, 2,
	149, 150, 5, 49, 25, 2, 150, 42, 3, 2, 2, 2, 151, 153, 7, 15, 2, 2, 152,
	151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 157,
	7, 12, 2, 2, 155, 157, 4, 14, 15, 2, 156, 152, 3, 2, 2, 2, 156, 155, 3,
	2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 160, 5, 71, 36, 2, 159, 158, 3, 2,
	2, 2, 159, 160, 3, 2, 2, 2, 160, 44, 3, 2, 2, 2, 161, 165, 5, 77, 39, 2,
	162, 164, 5, 79, 40, 2, 163, 162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 46, 3, 2, 2, 2, 167, 165, 3,
	2, 2, 2, 168, 172, 7, 36, 2, 2, 169, 171, 11, 2, 2, 2, 170, 169, 3, 2,
	2, 2, 171, 174, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	173, 175, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 176, 7, 36, 2, 2, 176,
	48, 3, 2, 2, 2, 177, 181, 5, 67, 34, 2, 178, 180, 5, 69, 35, 2, 179, 178,
	3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2,
	2, 2, 182, 190, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 186, 7, 50, 2, 2,
	185, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2, 189, 177, 3, 2, 2, 2, 189, 185,
	3, 2, 2, 2, 190, 50, 3, 2, 2, 2, 191, 192, 7, 42, 2, 2, 192, 52, 3, 2,
	2, 2, 193, 194, 7, 43, 2, 2, 194, 54, 3, 2, 2, 2, 195, 196, 7, 93, 2, 2,
	196, 56, 3, 2, 2, 2, 197, 198, 7, 95, 2, 2, 198, 58, 3, 2, 2, 2, 199, 200,
	7, 125, 2, 2, 200, 60, 3, 2, 2, 2, 201, 202, 7, 127, 2, 2, 202, 62, 3,
	2, 2, 2, 203, 207, 5, 71, 36, 2, 204, 207, 5, 73, 37, 2, 205, 207, 5, 75,
	38, 2, 206, 203, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 205, 3, 2, 2, 2,
	207, 208, 3, 2, 2, 2, 208, 209, 8, 32, 2, 2, 209, 64, 3, 2, 2, 2, 210,
	211, 11, 2, 2, 2, 211, 66, 3, 2, 2, 2, 212, 213, 9, 2, 2, 2, 213, 68, 3,
	2, 2, 2, 214, 215, 9, 3, 2, 2, 215, 70, 3, 2, 2, 2, 216, 218, 9, 4, 2,
	2, 217, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219,
	220, 3, 2, 2, 2, 220, 72, 3, 2, 2, 2, 221, 225, 7, 37, 2, 2, 222, 224,
	10, 5, 2, 2, 223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2,
	2, 2, 225, 226, 3, 2, 2, 2, 226, 74, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2,
	228, 230, 7, 94, 2, 2, 229, 231, 5, 71, 36, 2, 230, 229, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 237, 3, 2, 2, 2, 232, 234, 7, 15, 2, 2, 233, 232,
	3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 238, 7, 12,
	2, 2, 236, 238, 4, 14, 15, 2, 237, 233, 3, 2, 2, 2, 237, 236, 3, 2, 2,
	2, 238, 76, 3, 2, 2, 2, 239, 241, 9, 6, 2, 2, 240, 239, 3, 2, 2, 2, 241,
	78, 3, 2, 2, 2, 242, 245, 5, 77, 39, 2, 243, 245, 9, 3, 2, 2, 244, 242,
	3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 80, 3, 2, 2, 2, 19, 2, 152, 156,
	159, 165, 172, 181, 187, 189, 206, 219, 225, 230, 233, 237, 240, 244, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'break'", "'continue'", "'if'", "':'", "'elif'", "'else'",
	"'while'", "'print'", "'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'+'",
	"'-'", "", "", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"STRING", "NUMBER", "INTEGER", "NEWLINE", "NAME", "STRING_LITERAL", "DECIMAL_INTEGER",
	"OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACK", "CLOSE_BRACK", "OPEN_BRACE",
	"CLOSE_BRACE", "SKIP_", "UNKNOWN_CHAR",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"STRING", "NUMBER", "INTEGER", "NEWLINE", "NAME", "STRING_LITERAL", "DECIMAL_INTEGER",
	"OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACK", "CLOSE_BRACK", "OPEN_BRACE",
	"CLOSE_BRACE", "SKIP_", "UNKNOWN_CHAR", "NON_ZERO_DIGIT", "DIGIT", "SPACES",
	"COMMENT", "LINE_JOINING", "ID_START", "ID_CONTINUE",
}

type Python3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPython3Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Python3Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPython3Lexer(input antlr.CharStream) *Python3Lexer {
	l := new(Python3Lexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Python3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Python3Lexer tokens.
const (
	Python3LexerT__0            = 1
	Python3LexerT__1            = 2
	Python3LexerT__2            = 3
	Python3LexerT__3            = 4
	Python3LexerT__4            = 5
	Python3LexerT__5            = 6
	Python3LexerT__6            = 7
	Python3LexerT__7            = 8
	Python3LexerT__8            = 9
	Python3LexerT__9            = 10
	Python3LexerT__10           = 11
	Python3LexerT__11           = 12
	Python3LexerT__12           = 13
	Python3LexerT__13           = 14
	Python3LexerT__14           = 15
	Python3LexerT__15           = 16
	Python3LexerT__16           = 17
	Python3LexerSTRING          = 18
	Python3LexerNUMBER          = 19
	Python3LexerINTEGER         = 20
	Python3LexerNEWLINE         = 21
	Python3LexerNAME            = 22
	Python3LexerSTRING_LITERAL  = 23
	Python3LexerDECIMAL_INTEGER = 24
	Python3LexerOPEN_PAREN      = 25
	Python3LexerCLOSE_PAREN     = 26
	Python3LexerOPEN_BRACK      = 27
	Python3LexerCLOSE_BRACK     = 28
	Python3LexerOPEN_BRACE      = 29
	Python3LexerCLOSE_BRACE     = 30
	Python3LexerSKIP_           = 31
	Python3LexerUNKNOWN_CHAR    = 32
)
