// Code generated from Python3.g4 by ANTLR 4.9. DO NOT EDIT.

package python // Python3

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 154,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 42, 10, 2, 3, 3, 3, 3, 7, 3, 46,
	10, 3, 12, 3, 14, 3, 49, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 55, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 5, 8, 71, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 5, 11, 79, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 90, 10, 12, 12, 12, 14, 12, 93, 11, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 98, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 6, 14, 109, 10, 14, 13, 14, 14, 14, 110, 3, 14, 3,
	14, 5, 14, 115, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 121, 10, 15,
	12, 15, 14, 15, 124, 11, 15, 3, 16, 3, 16, 3, 16, 5, 16, 129, 10, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 140,
	10, 18, 3, 18, 3, 18, 3, 18, 6, 18, 145, 10, 18, 13, 18, 14, 18, 146, 7,
	18, 149, 10, 18, 12, 18, 14, 18, 152, 11, 18, 3, 18, 2, 3, 34, 19, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 4, 3, 2, 12,
	17, 3, 2, 18, 19, 2, 155, 2, 41, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 54,
	3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2,
	14, 70, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2, 18, 74, 3, 2, 2, 2, 20, 78, 3,
	2, 2, 2, 22, 80, 3, 2, 2, 2, 24, 99, 3, 2, 2, 2, 26, 114, 3, 2, 2, 2, 28,
	116, 3, 2, 2, 2, 30, 128, 3, 2, 2, 2, 32, 130, 3, 2, 2, 2, 34, 139, 3,
	2, 2, 2, 36, 42, 7, 23, 2, 2, 37, 42, 5, 8, 5, 2, 38, 39, 5, 20, 11, 2,
	39, 40, 7, 23, 2, 2, 40, 42, 3, 2, 2, 2, 41, 36, 3, 2, 2, 2, 41, 37, 3,
	2, 2, 2, 41, 38, 3, 2, 2, 2, 42, 3, 3, 2, 2, 2, 43, 46, 7, 23, 2, 2, 44,
	46, 5, 6, 4, 2, 45, 43, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2,
	2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 50, 51, 7, 2, 2, 3, 51, 5, 3, 2, 2, 2, 52, 55, 5, 8, 5, 2,
	53, 55, 5, 20, 11, 2, 54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 7, 3,
	2, 2, 2, 56, 57, 5, 10, 6, 2, 57, 58, 7, 23, 2, 2, 58, 9, 3, 2, 2, 2, 59,
	63, 5, 12, 7, 2, 60, 63, 5, 14, 8, 2, 61, 63, 5, 30, 16, 2, 62, 59, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 11, 3, 2, 2, 2, 64,
	65, 7, 24, 2, 2, 65, 66, 7, 3, 2, 2, 66, 67, 5, 34, 18, 2, 67, 13, 3, 2,
	2, 2, 68, 71, 5, 16, 9, 2, 69, 71, 5, 18, 10, 2, 70, 68, 3, 2, 2, 2, 70,
	69, 3, 2, 2, 2, 71, 15, 3, 2, 2, 2, 72, 73, 7, 4, 2, 2, 73, 17, 3, 2, 2,
	2, 74, 75, 7, 5, 2, 2, 75, 19, 3, 2, 2, 2, 76, 79, 5, 22, 12, 2, 77, 79,
	5, 24, 13, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 21, 3, 2, 2,
	2, 80, 81, 7, 6, 2, 2, 81, 82, 5, 28, 15, 2, 82, 83, 7, 7, 2, 2, 83, 91,
	5, 26, 14, 2, 84, 85, 7, 8, 2, 2, 85, 86, 5, 28, 15, 2, 86, 87, 7, 7, 2,
	2, 87, 88, 5, 26, 14, 2, 88, 90, 3, 2, 2, 2, 89, 84, 3, 2, 2, 2, 90, 93,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 97, 3, 2, 2, 2,
	93, 91, 3, 2, 2, 2, 94, 95, 7, 9, 2, 2, 95, 96, 7, 7, 2, 2, 96, 98, 5,
	26, 14, 2, 97, 94, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 23, 3, 2, 2, 2,
	99, 100, 7, 10, 2, 2, 100, 101, 5, 28, 15, 2, 101, 102, 7, 7, 2, 2, 102,
	103, 5, 26, 14, 2, 103, 25, 3, 2, 2, 2, 104, 115, 5, 8, 5, 2, 105, 106,
	7, 23, 2, 2, 106, 108, 7, 35, 2, 2, 107, 109, 5, 6, 4, 2, 108, 107, 3,
	2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2,
	2, 111, 112, 3, 2, 2, 2, 112, 113, 7, 36, 2, 2, 113, 115, 3, 2, 2, 2, 114,
	104, 3, 2, 2, 2, 114, 105, 3, 2, 2, 2, 115, 27, 3, 2, 2, 2, 116, 122, 5,
	34, 18, 2, 117, 118, 5, 32, 17, 2, 118, 119, 5, 34, 18, 2, 119, 121, 3,
	2, 2, 2, 120, 117, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2,
	2, 122, 123, 3, 2, 2, 2, 123, 29, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125,
	126, 7, 11, 2, 2, 126, 129, 7, 20, 2, 2, 127, 129, 5, 34, 18, 2, 128, 125,
	3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 31, 3, 2, 2, 2, 130, 131, 9, 2,
	2, 2, 131, 33, 3, 2, 2, 2, 132, 133, 8, 18, 1, 2, 133, 140, 7, 24, 2, 2,
	134, 140, 7, 21, 2, 2, 135, 136, 7, 27, 2, 2, 136, 137, 5, 34, 18, 2, 137,
	138, 7, 28, 2, 2, 138, 140, 3, 2, 2, 2, 139, 132, 3, 2, 2, 2, 139, 134,
	3, 2, 2, 2, 139, 135, 3, 2, 2, 2, 140, 150, 3, 2, 2, 2, 141, 144, 12, 6,
	2, 2, 142, 143, 9, 3, 2, 2, 143, 145, 5, 34, 18, 2, 144, 142, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147,
	149, 3, 2, 2, 2, 148, 141, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 35, 3, 2, 2, 2, 152, 150, 3, 2,
	2, 2, 18, 41, 45, 47, 54, 62, 70, 78, 91, 97, 110, 114, 122, 128, 139,
	146, 150,
}
var literalNames = []string{
	"", "'='", "'break'", "'continue'", "'if'", "':'", "'elif'", "'else'",
	"'while'", "'print'", "'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'+'",
	"'-'", "", "", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"STRING", "NUMBER", "INTEGER", "NEWLINE", "NAME", "STRING_LITERAL", "DECIMAL_INTEGER",
	"OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACK", "CLOSE_BRACK", "OPEN_BRACE",
	"CLOSE_BRACE", "SKIP_", "UNKNOWN_CHAR", "INDENT", "DEDENT",
}

var ruleNames = []string{
	"single_input", "file_input", "stmt", "simple_stmt", "small_stmt", "assignment_stmt",
	"flow_stmt", "break_stmt", "continue_stmt", "compound_stmt", "if_stmt",
	"while_stmt", "suite", "test", "print_stmt", "comp_op", "expr",
}

type Python3Parser struct {
	*antlr.BaseParser
}

// NewPython3Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Python3Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPython3Parser(input antlr.TokenStream) *Python3Parser {
	this := new(Python3Parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Python3.g4"

	return this
}

// Python3Parser tokens.
const (
	Python3ParserEOF             = antlr.TokenEOF
	Python3ParserT__0            = 1
	Python3ParserT__1            = 2
	Python3ParserT__2            = 3
	Python3ParserT__3            = 4
	Python3ParserT__4            = 5
	Python3ParserT__5            = 6
	Python3ParserT__6            = 7
	Python3ParserT__7            = 8
	Python3ParserT__8            = 9
	Python3ParserT__9            = 10
	Python3ParserT__10           = 11
	Python3ParserT__11           = 12
	Python3ParserT__12           = 13
	Python3ParserT__13           = 14
	Python3ParserT__14           = 15
	Python3ParserT__15           = 16
	Python3ParserT__16           = 17
	Python3ParserSTRING          = 18
	Python3ParserNUMBER          = 19
	Python3ParserINTEGER         = 20
	Python3ParserNEWLINE         = 21
	Python3ParserNAME            = 22
	Python3ParserSTRING_LITERAL  = 23
	Python3ParserDECIMAL_INTEGER = 24
	Python3ParserOPEN_PAREN      = 25
	Python3ParserCLOSE_PAREN     = 26
	Python3ParserOPEN_BRACK      = 27
	Python3ParserCLOSE_BRACK     = 28
	Python3ParserOPEN_BRACE      = 29
	Python3ParserCLOSE_BRACE     = 30
	Python3ParserSKIP_           = 31
	Python3ParserUNKNOWN_CHAR    = 32
	Python3ParserINDENT          = 33
	Python3ParserDEDENT          = 34
)

// Python3Parser rules.
const (
	Python3ParserRULE_single_input    = 0
	Python3ParserRULE_file_input      = 1
	Python3ParserRULE_stmt            = 2
	Python3ParserRULE_simple_stmt     = 3
	Python3ParserRULE_small_stmt      = 4
	Python3ParserRULE_assignment_stmt = 5
	Python3ParserRULE_flow_stmt       = 6
	Python3ParserRULE_break_stmt      = 7
	Python3ParserRULE_continue_stmt   = 8
	Python3ParserRULE_compound_stmt   = 9
	Python3ParserRULE_if_stmt         = 10
	Python3ParserRULE_while_stmt      = 11
	Python3ParserRULE_suite           = 12
	Python3ParserRULE_test            = 13
	Python3ParserRULE_print_stmt      = 14
	Python3ParserRULE_comp_op         = 15
	Python3ParserRULE_expr            = 16
)

// ISingle_inputContext is an interface to support dynamic dispatch.
type ISingle_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_inputContext differentiates from other interfaces.
	IsSingle_inputContext()
}

type Single_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_inputContext() *Single_inputContext {
	var p = new(Single_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_single_input
	return p
}

func (*Single_inputContext) IsSingle_inputContext() {}

func NewSingle_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_inputContext {
	var p = new(Single_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_single_input

	return p
}

func (s *Single_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_inputContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(Python3ParserNEWLINE, 0)
}

func (s *Single_inputContext) Simple_stmt() ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *Single_inputContext) Compound_stmt() ICompound_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_stmtContext)
}

func (s *Single_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterSingle_input(s)
	}
}

func (s *Single_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitSingle_input(s)
	}
}

func (p *Python3Parser) Single_input() (localctx ISingle_inputContext) {
	localctx = NewSingle_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Python3ParserRULE_single_input)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserNEWLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(Python3ParserNEWLINE)
		}

	case Python3ParserT__1, Python3ParserT__2, Python3ParserT__8, Python3ParserNUMBER, Python3ParserNAME, Python3ParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Simple_stmt()
		}

	case Python3ParserT__3, Python3ParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Compound_stmt()
		}
		{
			p.SetState(37)
			p.Match(Python3ParserNEWLINE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFile_inputContext is an interface to support dynamic dispatch.
type IFile_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_inputContext differentiates from other interfaces.
	IsFile_inputContext()
}

type File_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_inputContext() *File_inputContext {
	var p = new(File_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_file_input
	return p
}

func (*File_inputContext) IsFile_inputContext() {}

func NewFile_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_inputContext {
	var p = new(File_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_file_input

	return p
}

func (s *File_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *File_inputContext) EOF() antlr.TerminalNode {
	return s.GetToken(Python3ParserEOF, 0)
}

func (s *File_inputContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Python3ParserNEWLINE)
}

func (s *File_inputContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Python3ParserNEWLINE, i)
}

func (s *File_inputContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *File_inputContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *File_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterFile_input(s)
	}
}

func (s *File_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitFile_input(s)
	}
}

func (p *Python3Parser) File_input() (localctx IFile_inputContext) {
	localctx = NewFile_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Python3ParserRULE_file_input)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Python3ParserT__1)|(1<<Python3ParserT__2)|(1<<Python3ParserT__3)|(1<<Python3ParserT__7)|(1<<Python3ParserT__8)|(1<<Python3ParserNUMBER)|(1<<Python3ParserNEWLINE)|(1<<Python3ParserNAME)|(1<<Python3ParserOPEN_PAREN))) != 0 {
		p.SetState(43)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Python3ParserNEWLINE:
			{
				p.SetState(41)
				p.Match(Python3ParserNEWLINE)
			}

		case Python3ParserT__1, Python3ParserT__2, Python3ParserT__3, Python3ParserT__7, Python3ParserT__8, Python3ParserNUMBER, Python3ParserNAME, Python3ParserOPEN_PAREN:
			{
				p.SetState(42)
				p.Stmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(Python3ParserEOF)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Simple_stmt() ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *StmtContext) Compound_stmt() ICompound_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *Python3Parser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Python3ParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserT__1, Python3ParserT__2, Python3ParserT__8, Python3ParserNUMBER, Python3ParserNAME, Python3ParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Simple_stmt()
		}

	case Python3ParserT__3, Python3ParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Compound_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_stmtContext is an interface to support dynamic dispatch.
type ISimple_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_stmtContext differentiates from other interfaces.
	IsSimple_stmtContext()
}

type Simple_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_stmtContext() *Simple_stmtContext {
	var p = new(Simple_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_simple_stmt
	return p
}

func (*Simple_stmtContext) IsSimple_stmtContext() {}

func NewSimple_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_stmtContext {
	var p = new(Simple_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_simple_stmt

	return p
}

func (s *Simple_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_stmtContext) Small_stmt() ISmall_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISmall_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISmall_stmtContext)
}

func (s *Simple_stmtContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(Python3ParserNEWLINE, 0)
}

func (s *Simple_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterSimple_stmt(s)
	}
}

func (s *Simple_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitSimple_stmt(s)
	}
}

func (p *Python3Parser) Simple_stmt() (localctx ISimple_stmtContext) {
	localctx = NewSimple_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Python3ParserRULE_simple_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Small_stmt()
	}
	{
		p.SetState(55)
		p.Match(Python3ParserNEWLINE)
	}

	return localctx
}

// ISmall_stmtContext is an interface to support dynamic dispatch.
type ISmall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmall_stmtContext differentiates from other interfaces.
	IsSmall_stmtContext()
}

type Small_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmall_stmtContext() *Small_stmtContext {
	var p = new(Small_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_small_stmt
	return p
}

func (*Small_stmtContext) IsSmall_stmtContext() {}

func NewSmall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Small_stmtContext {
	var p = new(Small_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_small_stmt

	return p
}

func (s *Small_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Small_stmtContext) Assignment_stmt() IAssignment_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_stmtContext)
}

func (s *Small_stmtContext) Flow_stmt() IFlow_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlow_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlow_stmtContext)
}

func (s *Small_stmtContext) Print_stmt() IPrint_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint_stmtContext)
}

func (s *Small_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Small_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Small_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterSmall_stmt(s)
	}
}

func (s *Small_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitSmall_stmt(s)
	}
}

func (p *Python3Parser) Small_stmt() (localctx ISmall_stmtContext) {
	localctx = NewSmall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Python3ParserRULE_small_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Assignment_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Flow_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Print_stmt()
		}

	}

	return localctx
}

// IAssignment_stmtContext is an interface to support dynamic dispatch.
type IAssignment_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_stmtContext differentiates from other interfaces.
	IsAssignment_stmtContext()
}

type Assignment_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_stmtContext() *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_assignment_stmt
	return p
}

func (*Assignment_stmtContext) IsAssignment_stmtContext() {}

func NewAssignment_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_assignment_stmt

	return p
}

func (s *Assignment_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_stmtContext) NAME() antlr.TerminalNode {
	return s.GetToken(Python3ParserNAME, 0)
}

func (s *Assignment_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Assignment_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterAssignment_stmt(s)
	}
}

func (s *Assignment_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitAssignment_stmt(s)
	}
}

func (p *Python3Parser) Assignment_stmt() (localctx IAssignment_stmtContext) {
	localctx = NewAssignment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Python3ParserRULE_assignment_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(Python3ParserNAME)
	}
	{
		p.SetState(63)
		p.Match(Python3ParserT__0)
	}
	{
		p.SetState(64)
		p.expr(0)
	}

	return localctx
}

// IFlow_stmtContext is an interface to support dynamic dispatch.
type IFlow_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlow_stmtContext differentiates from other interfaces.
	IsFlow_stmtContext()
}

type Flow_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlow_stmtContext() *Flow_stmtContext {
	var p = new(Flow_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_flow_stmt
	return p
}

func (*Flow_stmtContext) IsFlow_stmtContext() {}

func NewFlow_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flow_stmtContext {
	var p = new(Flow_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_flow_stmt

	return p
}

func (s *Flow_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Flow_stmtContext) Break_stmt() IBreak_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *Flow_stmtContext) Continue_stmt() IContinue_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
}

func (s *Flow_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flow_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Flow_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterFlow_stmt(s)
	}
}

func (s *Flow_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitFlow_stmt(s)
	}
}

func (p *Python3Parser) Flow_stmt() (localctx IFlow_stmtContext) {
	localctx = NewFlow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Python3ParserRULE_flow_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Break_stmt()
		}

	case Python3ParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Continue_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_break_stmt
	return p
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterBreak_stmt(s)
	}
}

func (s *Break_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitBreak_stmt(s)
	}
}

func (p *Python3Parser) Break_stmt() (localctx IBreak_stmtContext) {
	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Python3ParserRULE_break_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(Python3ParserT__1)
	}

	return localctx
}

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_continue_stmt
	return p
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterContinue_stmt(s)
	}
}

func (s *Continue_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitContinue_stmt(s)
	}
}

func (p *Python3Parser) Continue_stmt() (localctx IContinue_stmtContext) {
	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Python3ParserRULE_continue_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(Python3ParserT__2)
	}

	return localctx
}

// ICompound_stmtContext is an interface to support dynamic dispatch.
type ICompound_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompound_stmtContext differentiates from other interfaces.
	IsCompound_stmtContext()
}

type Compound_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_stmtContext() *Compound_stmtContext {
	var p = new(Compound_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_compound_stmt
	return p
}

func (*Compound_stmtContext) IsCompound_stmtContext() {}

func NewCompound_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_stmtContext {
	var p = new(Compound_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_compound_stmt

	return p
}

func (s *Compound_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_stmtContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *Compound_stmtContext) While_stmt() IWhile_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *Compound_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterCompound_stmt(s)
	}
}

func (s *Compound_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitCompound_stmt(s)
	}
}

func (p *Python3Parser) Compound_stmt() (localctx ICompound_stmtContext) {
	localctx = NewCompound_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Python3ParserRULE_compound_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.If_stmt()
		}

	case Python3ParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.While_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) AllTest() []ITestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITestContext)(nil)).Elem())
	var tst = make([]ITestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITestContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Test(i int) ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *If_stmtContext) AllSuite() []ISuiteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISuiteContext)(nil)).Elem())
	var tst = make([]ISuiteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISuiteContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Suite(i int) ISuiteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuiteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISuiteContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (p *Python3Parser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Python3ParserRULE_if_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(Python3ParserT__3)
	}
	{
		p.SetState(79)
		p.Test()
	}
	{
		p.SetState(80)
		p.Match(Python3ParserT__4)
	}
	{
		p.SetState(81)
		p.Suite()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Python3ParserT__5 {
		{
			p.SetState(82)
			p.Match(Python3ParserT__5)
		}
		{
			p.SetState(83)
			p.Test()
		}
		{
			p.SetState(84)
			p.Match(Python3ParserT__4)
		}
		{
			p.SetState(85)
			p.Suite()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Python3ParserT__6 {
		{
			p.SetState(92)
			p.Match(Python3ParserT__6)
		}
		{
			p.SetState(93)
			p.Match(Python3ParserT__4)
		}
		{
			p.SetState(94)
			p.Suite()
		}

	}

	return localctx
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_while_stmt
	return p
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *While_stmtContext) Suite() ISuiteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuiteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuiteContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *Python3Parser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Python3ParserRULE_while_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(Python3ParserT__7)
	}
	{
		p.SetState(98)
		p.Test()
	}
	{
		p.SetState(99)
		p.Match(Python3ParserT__4)
	}
	{
		p.SetState(100)
		p.Suite()
	}

	return localctx
}

// ISuiteContext is an interface to support dynamic dispatch.
type ISuiteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuiteContext differentiates from other interfaces.
	IsSuiteContext()
}

type SuiteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuiteContext() *SuiteContext {
	var p = new(SuiteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_suite
	return p
}

func (*SuiteContext) IsSuiteContext() {}

func NewSuiteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuiteContext {
	var p = new(SuiteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_suite

	return p
}

func (s *SuiteContext) GetParser() antlr.Parser { return s.parser }

func (s *SuiteContext) Simple_stmt() ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *SuiteContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(Python3ParserNEWLINE, 0)
}

func (s *SuiteContext) INDENT() antlr.TerminalNode {
	return s.GetToken(Python3ParserINDENT, 0)
}

func (s *SuiteContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(Python3ParserDEDENT, 0)
}

func (s *SuiteContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *SuiteContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SuiteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuiteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuiteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterSuite(s)
	}
}

func (s *SuiteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitSuite(s)
	}
}

func (p *Python3Parser) Suite() (localctx ISuiteContext) {
	localctx = NewSuiteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Python3ParserRULE_suite)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserT__1, Python3ParserT__2, Python3ParserT__8, Python3ParserNUMBER, Python3ParserNAME, Python3ParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Simple_stmt()
		}

	case Python3ParserNEWLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(Python3ParserNEWLINE)
		}
		{
			p.SetState(104)
			p.Match(Python3ParserINDENT)
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Python3ParserT__1)|(1<<Python3ParserT__2)|(1<<Python3ParserT__3)|(1<<Python3ParserT__7)|(1<<Python3ParserT__8)|(1<<Python3ParserNUMBER)|(1<<Python3ParserNAME)|(1<<Python3ParserOPEN_PAREN))) != 0) {
			{
				p.SetState(105)
				p.Stmt()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(110)
			p.Match(Python3ParserDEDENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TestContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TestContext) AllComp_op() []IComp_opContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComp_opContext)(nil)).Elem())
	var tst = make([]IComp_opContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComp_opContext)
		}
	}

	return tst
}

func (s *TestContext) Comp_op(i int) IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitTest(s)
	}
}

func (p *Python3Parser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Python3ParserRULE_test)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.expr(0)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Python3ParserT__9)|(1<<Python3ParserT__10)|(1<<Python3ParserT__11)|(1<<Python3ParserT__12)|(1<<Python3ParserT__13)|(1<<Python3ParserT__14))) != 0 {
		{
			p.SetState(115)
			p.Comp_op()
		}
		{
			p.SetState(116)
			p.expr(0)
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrint_stmtContext is an interface to support dynamic dispatch.
type IPrint_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrint_stmtContext differentiates from other interfaces.
	IsPrint_stmtContext()
}

type Print_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_stmtContext() *Print_stmtContext {
	var p = new(Print_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_print_stmt
	return p
}

func (*Print_stmtContext) IsPrint_stmtContext() {}

func NewPrint_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_stmtContext {
	var p = new(Print_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_print_stmt

	return p
}

func (s *Print_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_stmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(Python3ParserSTRING, 0)
}

func (s *Print_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Print_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterPrint_stmt(s)
	}
}

func (s *Print_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitPrint_stmt(s)
	}
}

func (p *Python3Parser) Print_stmt() (localctx IPrint_stmtContext) {
	localctx = NewPrint_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Python3ParserRULE_print_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(Python3ParserT__8)
		}
		{
			p.SetState(124)
			p.Match(Python3ParserSTRING)
		}

	case Python3ParserNUMBER, Python3ParserNAME, Python3ParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComp_opContext is an interface to support dynamic dispatch.
type IComp_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComp_opContext differentiates from other interfaces.
	IsComp_opContext()
}

type Comp_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_opContext() *Comp_opContext {
	var p = new(Comp_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_comp_op
	return p
}

func (*Comp_opContext) IsComp_opContext() {}

func NewComp_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_opContext {
	var p = new(Comp_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_comp_op

	return p
}

func (s *Comp_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Comp_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comp_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterComp_op(s)
	}
}

func (s *Comp_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitComp_op(s)
	}
}

func (p *Python3Parser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Python3ParserRULE_comp_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Python3ParserT__9)|(1<<Python3ParserT__10)|(1<<Python3ParserT__11)|(1<<Python3ParserT__12)|(1<<Python3ParserT__13)|(1<<Python3ParserT__14))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Python3ParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Python3ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NAME() antlr.TerminalNode {
	return s.GetToken(Python3ParserNAME, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Python3ParserNUMBER, 0)
}

func (s *ExprContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(Python3ParserOPEN_PAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(Python3ParserCLOSE_PAREN, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Python3Listener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *Python3Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *Python3Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, Python3ParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Python3ParserNAME:
		{
			p.SetState(131)
			p.Match(Python3ParserNAME)
		}

	case Python3ParserNUMBER:
		{
			p.SetState(132)
			p.Match(Python3ParserNUMBER)
		}

	case Python3ParserOPEN_PAREN:
		{
			p.SetState(133)
			p.Match(Python3ParserOPEN_PAREN)
		}
		{
			p.SetState(134)
			p.expr(0)
		}
		{
			p.SetState(135)
			p.Match(Python3ParserCLOSE_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, Python3ParserRULE_expr)
			p.SetState(139)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(140)
						_la = p.GetTokenStream().LA(1)

						if !(_la == Python3ParserT__15 || _la == Python3ParserT__16) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}
					{
						p.SetState(141)
						p.expr(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(144)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
			}

		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Python3Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Python3Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
