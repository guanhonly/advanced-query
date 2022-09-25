// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // QueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#run.
	VisitRun(ctx *RunContext) interface{}

	// Visit a parse tree produced by QueryParser#equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by QueryParser#logicAnd.
	VisitLogicAnd(ctx *LogicAndContext) interface{}

	// Visit a parse tree produced by QueryParser#logicOR.
	VisitLogicOR(ctx *LogicORContext) interface{}

	// Visit a parse tree produced by QueryParser#nullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by QueryParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by QueryParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by QueryParser#logicNot.
	VisitLogicNot(ctx *LogicNotContext) interface{}

	// Visit a parse tree produced by QueryParser#notClause.
	VisitNotClause(ctx *NotClauseContext) interface{}
}
