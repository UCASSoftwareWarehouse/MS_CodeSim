// Code generated from Python3.g4 by ANTLR 4.9. DO NOT EDIT.

package python // Python3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Python3Listener is a complete listener for a parse tree produced by Python3Parser.
type Python3Listener interface {
	antlr.ParseTreeListener

	// EnterSingle_input is called when entering the single_input production.
	EnterSingle_input(c *Single_inputContext)

	// EnterFile_input is called when entering the file_input production.
	EnterFile_input(c *File_inputContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterSmall_stmt is called when entering the small_stmt production.
	EnterSmall_stmt(c *Small_stmtContext)

	// EnterAssignment_stmt is called when entering the assignment_stmt production.
	EnterAssignment_stmt(c *Assignment_stmtContext)

	// EnterFlow_stmt is called when entering the flow_stmt production.
	EnterFlow_stmt(c *Flow_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterCompound_stmt is called when entering the compound_stmt production.
	EnterCompound_stmt(c *Compound_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterSuite is called when entering the suite production.
	EnterSuite(c *SuiteContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitSingle_input is called when exiting the single_input production.
	ExitSingle_input(c *Single_inputContext)

	// ExitFile_input is called when exiting the file_input production.
	ExitFile_input(c *File_inputContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitSmall_stmt is called when exiting the small_stmt production.
	ExitSmall_stmt(c *Small_stmtContext)

	// ExitAssignment_stmt is called when exiting the assignment_stmt production.
	ExitAssignment_stmt(c *Assignment_stmtContext)

	// ExitFlow_stmt is called when exiting the flow_stmt production.
	ExitFlow_stmt(c *Flow_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitCompound_stmt is called when exiting the compound_stmt production.
	ExitCompound_stmt(c *Compound_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitSuite is called when exiting the suite production.
	ExitSuite(c *SuiteContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
