// QueryParser.g4
parser grammar QueryParser;

options {
    // the lexer is from QueryLexer
    tokenVocab = QueryLexer;
}

start : expression EOF;

expression:
    // Support ()
    LPAREN expression RPAREN                                                #lrExpr
    // Support K:V
    | leftExpr = expression operator = COLON rightExpr = expression         #eqExpr
    // Support logical AND OR
    | leftExpr = expression operator = (AND | OR) rightExpr = expression    #boolExpr
    // Support NOT
    | notClause                                                             #notExpr
    // Support any value
    | VALUE                                                                 #onlyValue
;

notClause:
    NOT expression
;
