package kv

import (
	"errors"
	"sort"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/guanhonly/advanced-query/runtime/go/entity"
	"github.com/guanhonly/advanced-query/runtime/go/parser"
)

func NewKVHandler() Handler {
	return &kvVisitor{
		kvs: make(map[string]map[entity.ValueInfo]bool),
	}
}

type kvVisitor struct {
	*parser.BaseQueryParserVisitor
	kvs   map[string]map[entity.ValueInfo]bool
	field string
	err   error
}

func (v *kvVisitor) VisitRun(ctx *parser.RunContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *kvVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *kvVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	value, ok := ctx.GetLeftExpr().(*parser.ValueContext)
	if !ok {
		v.err = errors.New("word before colon must be field name")
		return v.err
	}
	v.field = value.GetText()
	res := v.Visit(ctx.GetRightExpr())
	v.field = ""
	return res
}

func (v *kvVisitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *kvVisitor) VisitLogicAnd(ctx *parser.LogicAndContext) interface{} {
	return v.visitBinaryLogicOperator(ctx.GetLeftExpr(), ctx.GetRightExpr())
}

func (v *kvVisitor) VisitLogicOR(ctx *parser.LogicORContext) interface{} {
	return v.visitBinaryLogicOperator(ctx.GetLeftExpr(), ctx.GetRightExpr())
}

func (v *kvVisitor) visitBinaryLogicOperator(l, r parser.IExpressionContext) interface{} {
	lValue := v.Visit(l)
	rValue := v.Visit(r)
	_, lOk := lValue.(string)
	_, rOk := rValue.(string)
	if lOk && rOk {
		return nil
	}
	if !lOk {
		return v.Visit(l)
	}
	return v.Visit(r)
}

func (v *kvVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	v.fillKVs(ctx.GetStart().GetStart(), ctx.GetStop().GetStop(), ctx.GetText(), false)
	return ctx.GetText()
}

func (v *kvVisitor) VisitNotClause(ctx *parser.NotClauseContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *kvVisitor) VisitLogicNot(ctx *parser.LogicNotContext) interface{} {
	return v.Visit(ctx.NotClause())
}

func (v *kvVisitor) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	v.fillKVs(ctx.GetStart().GetStart(), ctx.GetStop().GetStop(), ctx.GetText(), true)
	return ctx.GetText()
}

func (v *kvVisitor) fillKVs(start, end int, value string, isNull bool) {
	if v.field == "" {
		return
	}
	m, ok := v.kvs[v.field]
	if !ok {
		m = make(map[entity.ValueInfo]bool)
		v.kvs[v.field] = m
	}
	vi := entity.ValueInfo{
		StartIndex: start,
		EndIndex:   end,
		IsNull:     isNull,
		Value:      value,
	}
	m[vi] = true
}

func (v *kvVisitor) Err() error {
	return v.err
}

func (v *kvVisitor) KVs() map[string][]*entity.ValueInfo {
	res := make(map[string][]*entity.ValueInfo)
	for field, values := range v.kvs {
		for val := range values {
			shadow := val
			res[field] = append(res[field], &shadow)
		}
	}
	for k, vs := range res {
		sort.Slice(vs, func(i, j int) bool {
			return vs[i].StartIndex < vs[j].StartIndex
		})
		res[k] = vs
	}
	return res
}
