package errlistener

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener interface {
	antlr.ErrorListener
	Err() error
}

func New() ErrorListener {
	return &errorListener{}
}

type errorListener struct {
	antlr.DefaultErrorListener
	err error
}

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.err = fmt.Errorf("syntax error at line %d, column: %d, error msg: %s", line, column, msg)
}

func (el *errorListener) Err() error {
	return el.err
}
