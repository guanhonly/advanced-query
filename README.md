# advanced-query
An advanced Google-like query parser based on ANTLR4. Support K-V, logical operator, and parentheses

## Functions
Parse a kind of Google-like query expression.  
Now it is only available in Go, and only the parsed field and value infos can be returned.  
If you hava more custom needs, that visit the ast and do what you want is also ok.

## Syntax
| KeyWord | Meaning                                                                                             | Example                                |
|---------|-----------------------------------------------------------------------------------------------------|----------------------------------------|
| AND     | logical AND, binary operator                                                                        | field1:v1 AND field2:v2                |
| OR      | logical OR, binary operator, lower precedence than AND                                              | field1:v1 OR field2:v2                 |
| NOT     | logic NOT, unary operator, lowest logical precedence                                                | NOT field1:v1                          |
| ()      | parentheses, change the order of expression, the expression inner parentheses will be visited first | field1:v1 AND (field2:v2 OR field3:v3) |

All the keyword is case-insensitive. If some value is conflict with the keywords, use double quote to escape it, just like:
```
field1:"AND"
```

## Examples
Please refer to the code in `runtime/go/examples`
