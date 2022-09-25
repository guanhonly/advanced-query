package kv

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/guanhonly/advanced-query/runtime/go/entity"
	"github.com/guanhonly/advanced-query/runtime/go/handler/errlistener"
	"github.com/guanhonly/advanced-query/runtime/go/parser"
)

type Handler interface {
	parser.QueryParserVisitor
	Err() error
	KVs() map[string][]*entity.ValueInfo
}

func GetExpressionKVs(in string, kvHandler Handler, el errlistener.ErrorListener) (map[string][]*entity.ValueInfo, error) {
	is := antlr.NewInputStream(in)

	lexer := parser.NewQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewQueryParser(stream)

	lexer.RemoveErrorListeners()
	p.RemoveErrorListeners()

	if el != nil {
		lexer.AddErrorListener(el)
		p.AddErrorListener(el)
	}

	p.Run().Accept(kvHandler)

	if el.Err() != nil {
		return nil, el.Err()
	}
	if kvHandler.Err() != nil {
		return nil, kvHandler.Err()
	}
	return kvHandler.KVs(), nil
}
