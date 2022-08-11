grammar SimpleBoolean;

parse: expression EOF;

expression:
	LPAREN expression RPAREN									# parenExpression
	| NOT expression											# notExpression
	| left = expression op = comparator right = expression		# comparatorExpression
	| left = expression op = binary right = expression			# binaryExpression
	| condition = expression THEN affirmative = expression (ELSE negative = expression)? 	# deprecatedConditionalExpression
	| affirmative = expression IF condition = expression (ELSE negative = expression)? 		# conditionalExpression
	| SQRT LPAREN expression RPAREN									# sqrtExpression
	| left = expression op = (MULT|DIV) right = expression  		# calcExpression
    | left = expression op = (ADD|SUB) right = expression  			# calcExpression
	| boolean														# boolExpression
	| DECIMAL														# decimalExpression
;

comparator:
	GT
	| GE
	| LT
	| LE
	| EQ
	| NEQ;

binary: AND | OR;

boolean: TRUE | FALSE;

AND: 'AND' | '&';
OR: 'OR' | '|';
NOT: 'NOT' | '!';
TRUE: 'TRUE' | 'True' | 'true';
FALSE: 'FALSE' | 'False' | 'false';
IF: 'IF';
THEN: 'THEN' | 'Then' | 'then' | '->';
ELSE: 'ELSE' | 'Else' | 'else';
GT: '>' | 'GT';
GE: '>=' | 'GE' | '≥';
LT: '<' | 'LT';
LE: '<=' | 'LE' | '≤';
EQ: '==' | 'EQ';
NEQ: '!=' | '<>' | 'NEQ';
LPAREN: '(';
RPAREN: ')';
DECIMAL: '-'? [0-9]+ ( '.' [0-9]+)?;
SQRT: 'SQRT';
MULT: '*';
DIV: '/';
ADD: '+';
SUB: '-';
WS: [ \r\t\u000C\n]+ -> skip;