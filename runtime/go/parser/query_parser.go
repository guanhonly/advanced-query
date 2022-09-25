// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // QueryParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type QueryParser struct {
	*antlr.BaseParser
}

var queryparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func queryparserParserInit() {
	staticData := &queryparserParserStaticData
	staticData.literalNames = []string{
		"", "", "'('", "')'", "'AND'", "'OR'", "'NOT'", "'NULL'", "':'", "",
		"'_'", "'*'", "'-'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "LPAREN", "RPAREN", "AND", "OR", "NOT", "NULLVALUE", "COLON",
		"CHINESE", "UNDERLINE", "STAR", "MINUS", "DOT", "VALUE",
	}
	staticData.ruleNames = []string{
		"run", "expression", "notClause",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 37, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 18, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10, 1, 12, 1,
		32, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 0, 1, 2, 3, 0, 2, 4, 0, 0, 39, 0, 6,
		1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7,
		8, 5, 0, 0, 1, 8, 1, 1, 0, 0, 0, 9, 10, 6, 1, -1, 0, 10, 11, 5, 2, 0, 0,
		11, 12, 3, 2, 1, 0, 12, 13, 5, 3, 0, 0, 13, 18, 1, 0, 0, 0, 14, 18, 3,
		4, 2, 0, 15, 18, 5, 7, 0, 0, 16, 18, 5, 14, 0, 0, 17, 9, 1, 0, 0, 0, 17,
		14, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16, 1, 0, 0, 0, 18, 30, 1, 0, 0,
		0, 19, 20, 10, 6, 0, 0, 20, 21, 5, 8, 0, 0, 21, 29, 3, 2, 1, 7, 22, 23,
		10, 5, 0, 0, 23, 24, 5, 4, 0, 0, 24, 29, 3, 2, 1, 6, 25, 26, 10, 4, 0,
		0, 26, 27, 5, 5, 0, 0, 27, 29, 3, 2, 1, 5, 28, 19, 1, 0, 0, 0, 28, 22,
		1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0,
		30, 31, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 5, 6,
		0, 0, 34, 35, 3, 2, 1, 0, 35, 5, 1, 0, 0, 0, 3, 17, 28, 30,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// QueryParserInit initializes any static state used to implement QueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryParserInit() {
	staticData := &queryparserParserStaticData
	staticData.once.Do(queryparserParserInit)
}

// NewQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryParser(input antlr.TokenStream) *QueryParser {
	QueryParserInit()
	this := new(QueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &queryparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "QueryParser.g4"

	return this
}

// QueryParser tokens.
const (
	QueryParserEOF       = antlr.TokenEOF
	QueryParserSPACE     = 1
	QueryParserLPAREN    = 2
	QueryParserRPAREN    = 3
	QueryParserAND       = 4
	QueryParserOR        = 5
	QueryParserNOT       = 6
	QueryParserNULLVALUE = 7
	QueryParserCOLON     = 8
	QueryParserCHINESE   = 9
	QueryParserUNDERLINE = 10
	QueryParserSTAR      = 11
	QueryParserMINUS     = 12
	QueryParserDOT       = 13
	QueryParserVALUE     = 14
)

// QueryParser rules.
const (
	QueryParserRULE_run        = 0
	QueryParserRULE_expression = 1
	QueryParserRULE_notClause  = 2
)

// IRunContext is an interface to support dynamic dispatch.
type IRunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRunContext differentiates from other interfaces.
	IsRunContext()
}

type RunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunContext() *RunContext {
	var p = new(RunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_run
	return p
}

func (*RunContext) IsRunContext() {}

func NewRunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunContext {
	var p = new(RunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_run

	return p
}

func (s *RunContext) GetParser() antlr.Parser { return s.parser }

func (s *RunContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RunContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryParserEOF, 0)
}

func (s *RunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterRun(s)
	}
}

func (s *RunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitRun(s)
	}
}

func (s *RunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitRun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Run() (localctx IRunContext) {
	this := p
	_ = this

	localctx = NewRunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_run)

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
		p.SetState(6)
		p.expression(0)
	}
	{
		p.SetState(7)
		p.Match(QueryParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualContext struct {
	*ExpressionContext
	leftExpr  IExpressionContext
	operator  antlr.Token
	rightExpr IExpressionContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualContext) GetOperator() antlr.Token { return s.operator }

func (s *EqualContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *EqualContext) GetLeftExpr() IExpressionContext { return s.leftExpr }

func (s *EqualContext) GetRightExpr() IExpressionContext { return s.rightExpr }

func (s *EqualContext) SetLeftExpr(v IExpressionContext) { s.leftExpr = v }

func (s *EqualContext) SetRightExpr(v IExpressionContext) { s.rightExpr = v }

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualContext) COLON() antlr.TerminalNode {
	return s.GetToken(QueryParserCOLON, 0)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitEqual(s)
	}
}

func (s *EqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicAndContext struct {
	*ExpressionContext
	leftExpr  IExpressionContext
	operator  antlr.Token
	rightExpr IExpressionContext
}

func NewLogicAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicAndContext {
	var p = new(LogicAndContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicAndContext) GetOperator() antlr.Token { return s.operator }

func (s *LogicAndContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *LogicAndContext) GetLeftExpr() IExpressionContext { return s.leftExpr }

func (s *LogicAndContext) GetRightExpr() IExpressionContext { return s.rightExpr }

func (s *LogicAndContext) SetLeftExpr(v IExpressionContext) { s.leftExpr = v }

func (s *LogicAndContext) SetRightExpr(v IExpressionContext) { s.rightExpr = v }

func (s *LogicAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicAndContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicAndContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicAndContext) AND() antlr.TerminalNode {
	return s.GetToken(QueryParserAND, 0)
}

func (s *LogicAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterLogicAnd(s)
	}
}

func (s *LogicAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitLogicAnd(s)
	}
}

func (s *LogicAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitLogicAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicORContext struct {
	*ExpressionContext
	leftExpr  IExpressionContext
	operator  antlr.Token
	rightExpr IExpressionContext
}

func NewLogicORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicORContext {
	var p = new(LogicORContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicORContext) GetOperator() antlr.Token { return s.operator }

func (s *LogicORContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *LogicORContext) GetLeftExpr() IExpressionContext { return s.leftExpr }

func (s *LogicORContext) GetRightExpr() IExpressionContext { return s.rightExpr }

func (s *LogicORContext) SetLeftExpr(v IExpressionContext) { s.leftExpr = v }

func (s *LogicORContext) SetRightExpr(v IExpressionContext) { s.rightExpr = v }

func (s *LogicORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicORContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicORContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicORContext) OR() antlr.TerminalNode {
	return s.GetToken(QueryParserOR, 0)
}

func (s *LogicORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterLogicOR(s)
	}
}

func (s *LogicORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitLogicOR(s)
	}
}

func (s *LogicORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitLogicOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullValueContext struct {
	*ExpressionContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) NULLVALUE() antlr.TerminalNode {
	return s.GetToken(QueryParserNULLVALUE, 0)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContext struct {
	*ExpressionContext
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(QueryParserVALUE, 0)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	*ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QueryParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QueryParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicNotContext struct {
	*ExpressionContext
}

func NewLogicNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicNotContext {
	var p = new(LogicNotContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicNotContext) NotClause() INotClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotClauseContext)
}

func (s *LogicNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterLogicNot(s)
	}
}

func (s *LogicNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitLogicNot(s)
	}
}

func (s *LogicNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitLogicNot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *QueryParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, QueryParserRULE_expression, _p)

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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(10)
			p.Match(QueryParserLPAREN)
		}
		{
			p.SetState(11)
			p.expression(0)
		}
		{
			p.SetState(12)
			p.Match(QueryParserRPAREN)
		}

	case QueryParserNOT:
		localctx = NewLogicNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.NotClause()
		}

	case QueryParserNULLVALUE:
		localctx = NewNullValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(QueryParserNULLVALUE)
		}

	case QueryParserVALUE:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.Match(QueryParserVALUE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).leftExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_expression)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(20)

					var _m = p.Match(QueryParserCOLON)

					localctx.(*EqualContext).operator = _m
				}
				{
					p.SetState(21)

					var _x = p.expression(7)

					localctx.(*EqualContext).rightExpr = _x
				}

			case 2:
				localctx = NewLogicAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LogicAndContext).leftExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_expression)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(23)

					var _m = p.Match(QueryParserAND)

					localctx.(*LogicAndContext).operator = _m
				}
				{
					p.SetState(24)

					var _x = p.expression(6)

					localctx.(*LogicAndContext).rightExpr = _x
				}

			case 3:
				localctx = NewLogicORContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LogicORContext).leftExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_expression)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(26)

					var _m = p.Match(QueryParserOR)

					localctx.(*LogicORContext).operator = _m
				}
				{
					p.SetState(27)

					var _x = p.expression(5)

					localctx.(*LogicORContext).rightExpr = _x
				}

			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// INotClauseContext is an interface to support dynamic dispatch.
type INotClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotClauseContext differentiates from other interfaces.
	IsNotClauseContext()
}

type NotClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotClauseContext() *NotClauseContext {
	var p = new(NotClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_notClause
	return p
}

func (*NotClauseContext) IsNotClauseContext() {}

func NewNotClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotClauseContext {
	var p = new(NotClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_notClause

	return p
}

func (s *NotClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *NotClauseContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *NotClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.EnterNotClause(s)
	}
}

func (s *NotClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryParserListener); ok {
		listenerT.ExitNotClause(s)
	}
}

func (s *NotClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryParserVisitor:
		return t.VisitNotClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) NotClause() (localctx INotClauseContext) {
	this := p
	_ = this

	localctx = NewNotClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_notClause)

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
		p.SetState(33)
		p.Match(QueryParserNOT)
	}
	{
		p.SetState(34)
		p.expression(0)
	}

	return localctx
}

func (p *QueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
