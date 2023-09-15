// Code generated from Lecture.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lecture // Lecture
import "github.com/antlr4-go/antlr/v4"

// LectureListener is a complete listener for a parse tree produced by LectureParser.
type LectureListener interface {
	antlr.ParseTreeListener

	// EnterLecture is called when entering the lecture production.
	EnterLecture(c *LectureContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStartClause is called when entering the startClause production.
	EnterStartClause(c *StartClauseContext)

	// EnterEndClause is called when entering the endClause production.
	EnterEndClause(c *EndClauseContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclarationStatement is called when entering the declarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterValueClause is called when entering the valueClause production.
	EnterValueClause(c *ValueClauseContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterLiteralClause is called when entering the literalClause production.
	EnterLiteralClause(c *LiteralClauseContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitLecture is called when exiting the lecture production.
	ExitLecture(c *LectureContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStartClause is called when exiting the startClause production.
	ExitStartClause(c *StartClauseContext)

	// ExitEndClause is called when exiting the endClause production.
	ExitEndClause(c *EndClauseContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclarationStatement is called when exiting the declarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitValueClause is called when exiting the valueClause production.
	ExitValueClause(c *ValueClauseContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitLiteralClause is called when exiting the literalClause production.
	ExitLiteralClause(c *LiteralClauseContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
