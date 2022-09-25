// QueryLexer.g4
lexer grammar QueryLexer;
options { caseInsensitive = true; }

SPACE: [ \t\r\n]+ -> channel(HIDDEN);

LPAREN: '(';
RPAREN: ')';
AND: 'AND';
OR: 'OR';
NOT: 'NOT';
NULLVALUE: 'NULL';

COLON: ':';

CHINESE: '\u4E00'..'\u9FA5';
UNDERLINE: '_';
STAR: '*';
MINUS: '-';
DOT: '.';

fragment LETTER: [a-z];
fragment DIGIT: [0-9];
fragment DQUOTE_STRING: '"'('\\'.|~('"'|'\\'))*'"';

VALUE: (CHINESE|LETTER|DIGIT|DQUOTE_STRING|UNDERLINE|MINUS|STAR|DOT)+;
