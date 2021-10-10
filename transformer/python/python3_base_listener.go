// Code generated from Python3.g4 by ANTLR 4.9. DO NOT EDIT.

package python // Python3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePython3Listener is a complete listener for a parse tree produced by Python3Parser.
type BasePython3Listener struct{}

var _ Python3Listener = &BasePython3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePython3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePython3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePython3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePython3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingle_input is called when production single_input is entered.
func (s *BasePython3Listener) EnterSingle_input(ctx *Single_inputContext) {}

// ExitSingle_input is called when production single_input is exited.
func (s *BasePython3Listener) ExitSingle_input(ctx *Single_inputContext) {}

// EnterFile_input is called when production file_input is entered.
func (s *BasePython3Listener) EnterFile_input(ctx *File_inputContext) {}

// ExitFile_input is called when production file_input is exited.
func (s *BasePython3Listener) ExitFile_input(ctx *File_inputContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePython3Listener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePython3Listener) ExitStmt(ctx *StmtContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BasePython3Listener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BasePython3Listener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterSmall_stmt is called when production small_stmt is entered.
func (s *BasePython3Listener) EnterSmall_stmt(ctx *Small_stmtContext) {}

// ExitSmall_stmt is called when production small_stmt is exited.
func (s *BasePython3Listener) ExitSmall_stmt(ctx *Small_stmtContext) {}

// EnterAssignment_stmt is called when production assignment_stmt is entered.
func (s *BasePython3Listener) EnterAssignment_stmt(ctx *Assignment_stmtContext) {}

// ExitAssignment_stmt is called when production assignment_stmt is exited.
func (s *BasePython3Listener) ExitAssignment_stmt(ctx *Assignment_stmtContext) {}

// EnterFlow_stmt is called when production flow_stmt is entered.
func (s *BasePython3Listener) EnterFlow_stmt(ctx *Flow_stmtContext) {}

// ExitFlow_stmt is called when production flow_stmt is exited.
func (s *BasePython3Listener) ExitFlow_stmt(ctx *Flow_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BasePython3Listener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BasePython3Listener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BasePython3Listener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BasePython3Listener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterCompound_stmt is called when production compound_stmt is entered.
func (s *BasePython3Listener) EnterCompound_stmt(ctx *Compound_stmtContext) {}

// ExitCompound_stmt is called when production compound_stmt is exited.
func (s *BasePython3Listener) ExitCompound_stmt(ctx *Compound_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BasePython3Listener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BasePython3Listener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BasePython3Listener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BasePython3Listener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterSuite is called when production suite is entered.
func (s *BasePython3Listener) EnterSuite(ctx *SuiteContext) {}

// ExitSuite is called when production suite is exited.
func (s *BasePython3Listener) ExitSuite(ctx *SuiteContext) {}

// EnterTest is called when production test is entered.
func (s *BasePython3Listener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasePython3Listener) ExitTest(ctx *TestContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BasePython3Listener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BasePython3Listener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BasePython3Listener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BasePython3Listener) ExitComp_op(ctx *Comp_opContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePython3Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePython3Listener) ExitExpr(ctx *ExprContext) {}
