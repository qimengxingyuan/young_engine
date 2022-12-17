grammar Engine;

expr: logOr EOF;
logOr: logOr '||' logAnd | logAnd;
logAnd: logAnd '&&' logNot | logNot;
logNot: '!' logNot | cmp;
cmp: cmp '>' add | cmp '>=' add | cmp '<' add | cmp '<=' add | cmp '==' add | cmp '!=' add | add;
add: add '+' mul | add '-' mul | mul;
mul: mul '*' pri | mul '/' pri | mul '%' pri | pri;
pri: BooleanLiteral|IntegerLiteral|FloatLiteral|StringLiteral|Identifier|'('expr')';



Addition                   : '+';
Subtraction                : '-';
Multiply                   : '*';
Divide                     : '/';
Modulus                    : '%';

GreaterThan                : '>';
LessThan                   : '<';
GreaterEqual               : '>=';
LessEqual                  : '<=';
Equal                      : '==';
NotEqual                   : '!=';

And                        : '&&';
Or                         : '||';
Not                        : '!';

OpenParen                  : '(';
CloseParen                 : ')';


BooleanLiteral: 'true' | 'false';
IntegerLiteral: '0'| [1-9] [0-9]*;
FloatLiteral: IntegerLiteral '.' Digit+| '.' Digit+;
StringLiteral: '"' ~["]* '"'| '\'' ~[\\']* '\''| '`' ~[`]* '`';
Identifier: IdentifierStart IdentifierPart*;

fragment IdentifierStart
    : Letter
    | [_]
    ;

fragment IdentifierPart
    : Letter
    | Digit
    | [_]
    ;

fragment Letter
    : 'A'..'Z'
    | 'a'..'z'
    ;

fragment Digit
    : '0'..'9'
    ;
