// QueryLexer.g4
lexer grammar QueryLexer;

SPACE: [ \t\r\n]+ -> channel(HIDDEN);

LPAREN: '(';
RPAREN: ')';
AND: 'AND';
OR: 'OR';
NOT: 'NOT';
COLON: ':';

CHINESE: '\u4E00'..'\u9FA5';  //表示所有中文的unicode编码，以支持中文
UNDERLINE: '_';
fragment STRING: [a-zA-Z0-9_];
fragment DEC_DIGIT: [0-9];
MINUS: '-';  //使MINUS和-等价，以下同理
INT: MINUS? DEC_DIGIT+;

FLOAT: (MINUS? DEC_DIGIT+ DOT DEC_DIGIT+)| (MINUS? DOT DEC_DIGIT+);

DOT: '.' ->mode(AFTER_DOT);

VALUE: (CHINESE|STRING|UNDERLINE|INT|FLOAT)+;

mode AFTER_DOT;

//DEFAULT_MODE是Antlr中默认定义好的mode
DOTINTEGER: ( '0' | [1-9] [0-9]*) -> mode(DEFAULT_MODE);
DOTID: [_a-zA-Z] [_a-zA-Z0-9]* -> mode(DEFAULT_MODE);