// Code generated from C:/Users/honly/Documents/GitHub/advanced-query/antlr/config\QueryLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var querylexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylexerLexerInit() {
	staticData := &querylexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'('", "')'", "'AND'", "'OR'", "'NOT'", "'NULL'", "':'", "",
		"'_'", "'*'", "'-'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "LPAREN", "RPAREN", "AND", "OR", "NOT", "NULLVALUE", "COLON",
		"CHINESE", "UNDERLINE", "STAR", "MINUS", "DOT", "VALUE",
	}
	staticData.ruleNames = []string{
		"SPACE", "LPAREN", "RPAREN", "AND", "OR", "NOT", "NULLVALUE", "COLON",
		"CHINESE", "UNDERLINE", "STAR", "MINUS", "DOT", "LETTER", "DIGIT", "DQUOTE_STRING",
		"VALUE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 101, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 4, 0, 37, 8, 0, 11, 0, 12, 0, 38, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 83, 8, 15, 10, 15, 12, 15,
		86, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 4, 16, 98, 8, 16, 11, 16, 12, 16, 99, 0, 0, 17, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 0, 29, 0, 31, 0, 33, 14, 1, 0, 12, 3, 0, 9, 10, 13, 13, 32, 32, 2,
		0, 65, 65, 97, 97, 2, 0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0,
		79, 79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 84, 84, 116, 116, 2, 0,
		85, 85, 117, 117, 2, 0, 76, 76, 108, 108, 2, 0, 65, 90, 97, 122, 1, 0,
		48, 57, 2, 0, 34, 34, 92, 92, 108, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 1, 36, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 44, 1, 0, 0, 0, 7, 46, 1,
		0, 0, 0, 9, 50, 1, 0, 0, 0, 11, 53, 1, 0, 0, 0, 13, 57, 1, 0, 0, 0, 15,
		62, 1, 0, 0, 0, 17, 64, 1, 0, 0, 0, 19, 66, 1, 0, 0, 0, 21, 68, 1, 0, 0,
		0, 23, 70, 1, 0, 0, 0, 25, 72, 1, 0, 0, 0, 27, 74, 1, 0, 0, 0, 29, 76,
		1, 0, 0, 0, 31, 78, 1, 0, 0, 0, 33, 97, 1, 0, 0, 0, 35, 37, 7, 0, 0, 0,
		36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1,
		0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 6, 0, 0, 0, 41, 2, 1, 0, 0, 0, 42,
		43, 5, 40, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 41, 0, 0, 45, 6, 1, 0, 0,
		0, 46, 47, 7, 1, 0, 0, 47, 48, 7, 2, 0, 0, 48, 49, 7, 3, 0, 0, 49, 8, 1,
		0, 0, 0, 50, 51, 7, 4, 0, 0, 51, 52, 7, 5, 0, 0, 52, 10, 1, 0, 0, 0, 53,
		54, 7, 2, 0, 0, 54, 55, 7, 4, 0, 0, 55, 56, 7, 6, 0, 0, 56, 12, 1, 0, 0,
		0, 57, 58, 7, 2, 0, 0, 58, 59, 7, 7, 0, 0, 59, 60, 7, 8, 0, 0, 60, 61,
		7, 8, 0, 0, 61, 14, 1, 0, 0, 0, 62, 63, 5, 58, 0, 0, 63, 16, 1, 0, 0, 0,
		64, 65, 2, 19968, 40869, 0, 65, 18, 1, 0, 0, 0, 66, 67, 5, 95, 0, 0, 67,
		20, 1, 0, 0, 0, 68, 69, 5, 42, 0, 0, 69, 22, 1, 0, 0, 0, 70, 71, 5, 45,
		0, 0, 71, 24, 1, 0, 0, 0, 72, 73, 5, 46, 0, 0, 73, 26, 1, 0, 0, 0, 74,
		75, 7, 9, 0, 0, 75, 28, 1, 0, 0, 0, 76, 77, 7, 10, 0, 0, 77, 30, 1, 0,
		0, 0, 78, 84, 5, 34, 0, 0, 79, 80, 5, 92, 0, 0, 80, 83, 9, 0, 0, 0, 81,
		83, 8, 11, 0, 0, 82, 79, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 87, 88, 5, 34, 0, 0, 88, 32, 1, 0, 0, 0, 89, 98, 3, 17, 8,
		0, 90, 98, 3, 27, 13, 0, 91, 98, 3, 29, 14, 0, 92, 98, 3, 31, 15, 0, 93,
		98, 3, 19, 9, 0, 94, 98, 3, 23, 11, 0, 95, 98, 3, 21, 10, 0, 96, 98, 3,
		25, 12, 0, 97, 89, 1, 0, 0, 0, 97, 90, 1, 0, 0, 0, 97, 91, 1, 0, 0, 0,
		97, 92, 1, 0, 0, 0, 97, 93, 1, 0, 0, 0, 97, 94, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 34, 1, 0, 0, 0, 6, 0, 38, 82, 84, 97, 99, 1, 0, 1,
		0,
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

// QueryLexerInit initializes any static state used to implement QueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLexerInit() {
	staticData := &querylexerLexerStaticData
	staticData.once.Do(querylexerLexerInit)
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	QueryLexerInit()
	l := new(QueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &querylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "QueryLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerSPACE     = 1
	QueryLexerLPAREN    = 2
	QueryLexerRPAREN    = 3
	QueryLexerAND       = 4
	QueryLexerOR        = 5
	QueryLexerNOT       = 6
	QueryLexerNULLVALUE = 7
	QueryLexerCOLON     = 8
	QueryLexerCHINESE   = 9
	QueryLexerUNDERLINE = 10
	QueryLexerSTAR      = 11
	QueryLexerMINUS     = 12
	QueryLexerDOT       = 13
	QueryLexerVALUE     = 14
)
