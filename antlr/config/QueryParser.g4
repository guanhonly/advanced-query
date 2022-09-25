// QueryParser.g4
parser grammar QueryParser;

options {
    // the lexer is from QueryLexer
    tokenVocab = QueryLexer;
}

run : expression EOF;

expression:
    // ()
    LPAREN expression RPAREN                                                #parenExpr
    // K:V
    | leftExpr = expression operator = COLON rightExpr = expression         #equal
    // logical AND OR. AND has higher precendence than OR
    | leftExpr = expression operator = AND rightExpr = expression           #logicAnd
    | leftExpr = expression operator = OR rightExpr = expression            #logicOR
    // NOT
    | notClause                                                             #logicNot
    // NULLVALUE flag
    | NULLVALUE                                                             #nullValue
    | VALUE                                                                 #value
;

notClause:
    NOT expression
;
