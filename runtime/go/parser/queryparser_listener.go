// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // QueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryParserListener is a complete listener for a parse tree produced by QueryParser.
type QueryParserListener interface {
	antlr.ParseTreeListener

	// EnterRun is called when entering the run production.
	EnterRun(c *RunContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterLogicAnd is called when entering the logicAnd production.
	EnterLogicAnd(c *LogicAndContext)

	// EnterLogicOR is called when entering the logicOR production.
	EnterLogicOR(c *LogicORContext)

	// EnterNullValue is called when entering the nullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterLogicNot is called when entering the logicNot production.
	EnterLogicNot(c *LogicNotContext)

	// EnterNotClause is called when entering the notClause production.
	EnterNotClause(c *NotClauseContext)

	// ExitRun is called when exiting the run production.
	ExitRun(c *RunContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitLogicAnd is called when exiting the logicAnd production.
	ExitLogicAnd(c *LogicAndContext)

	// ExitLogicOR is called when exiting the logicOR production.
	ExitLogicOR(c *LogicORContext)

	// ExitNullValue is called when exiting the nullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitLogicNot is called when exiting the logicNot production.
	ExitLogicNot(c *LogicNotContext)

	// ExitNotClause is called when exiting the notClause production.
	ExitNotClause(c *NotClauseContext)
}
