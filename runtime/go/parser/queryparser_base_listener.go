// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // QueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryParserListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryParserListener struct{}

var _ QueryParserListener = &BaseQueryParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRun is called when production run is entered.
func (s *BaseQueryParserListener) EnterRun(ctx *RunContext) {}

// ExitRun is called when production run is exited.
func (s *BaseQueryParserListener) ExitRun(ctx *RunContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseQueryParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseQueryParserListener) ExitEqual(ctx *EqualContext) {}

// EnterLogicAnd is called when production logicAnd is entered.
func (s *BaseQueryParserListener) EnterLogicAnd(ctx *LogicAndContext) {}

// ExitLogicAnd is called when production logicAnd is exited.
func (s *BaseQueryParserListener) ExitLogicAnd(ctx *LogicAndContext) {}

// EnterLogicOR is called when production logicOR is entered.
func (s *BaseQueryParserListener) EnterLogicOR(ctx *LogicORContext) {}

// ExitLogicOR is called when production logicOR is exited.
func (s *BaseQueryParserListener) ExitLogicOR(ctx *LogicORContext) {}

// EnterNullValue is called when production nullValue is entered.
func (s *BaseQueryParserListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production nullValue is exited.
func (s *BaseQueryParserListener) ExitNullValue(ctx *NullValueContext) {}

// EnterValue is called when production value is entered.
func (s *BaseQueryParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseQueryParserListener) ExitValue(ctx *ValueContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseQueryParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseQueryParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterLogicNot is called when production logicNot is entered.
func (s *BaseQueryParserListener) EnterLogicNot(ctx *LogicNotContext) {}

// ExitLogicNot is called when production logicNot is exited.
func (s *BaseQueryParserListener) ExitLogicNot(ctx *LogicNotContext) {}

// EnterNotClause is called when production notClause is entered.
func (s *BaseQueryParserListener) EnterNotClause(ctx *NotClauseContext) {}

// ExitNotClause is called when production notClause is exited.
func (s *BaseQueryParserListener) ExitNotClause(ctx *NotClauseContext) {}
