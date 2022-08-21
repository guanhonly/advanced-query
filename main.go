package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/guanhonly/advanced-query/antlr/parser"
)

func main() {
	// Create the input stream for antlr
	is := antlr.NewInputStream("f1:NOT(vl OR v2)")

	// Create lexer
	lexer := parser.NewQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewQueryParser(stream)
	visit(p)
}

func visit(p *parser.QueryParser) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("panic:", r)
		}
	}()
	v := &visitor{}
	p.Start().Accept(v)
}
