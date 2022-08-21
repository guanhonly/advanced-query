package main

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/guanhonly/advanced-query/antlr/parser"
)

type visitor struct {
	*parser.BaseQueryParserVisitor
	err    error
	fields []string
	values []string
}

func (v *visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *visitor) VisitEqExpr(ctx *parser.EqExprContext) interface{} {
	val, ok := ctx.GetLeftExpr().(*parser.OnlyValueContext)
	if !ok {
		v.err = errors.New("field name only support text")
		return nil
	}
	field := val.GetText()
	v.fields = append(v.fields, field)
	return v.Visit(ctx.GetRightExpr())
}

func (v *visitor) VisitLrExpr(ctx *parser.LrExprContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	left := ctx.GetLeftExpr()
	right := ctx.GetRightExpr()
	leftValue, lOk := left.(*parser.OnlyValueContext)
	rightValue, rOk := right.(*parser.OnlyValueContext)
	if lOk && rOk {
		v.values = append(v.values, leftValue.GetText())
		v.values = append(v.values, rightValue.GetText())
		return nil
	}
	if !lOk {
		return v.Visit(left)
	}
	return v.Visit(right)
}

func (v *visitor) VisitOnlyValue(ctx *parser.OnlyValueContext) interface{} {
	v.values = append(v.values, ctx.GetText())
	return ctx.GetText()
}

func (v *visitor) VisitNotClause(ctx *parser.NotClauseContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return v.Visit(ctx.NotClause())
}
