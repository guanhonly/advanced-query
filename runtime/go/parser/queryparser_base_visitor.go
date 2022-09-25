// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // QueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQueryParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryParserVisitor) VisitRun(ctx *RunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitEqual(ctx *EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitLogicAnd(ctx *LogicAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitLogicOR(ctx *LogicORContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitLogicNot(ctx *LogicNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryParserVisitor) VisitNotClause(ctx *NotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}
