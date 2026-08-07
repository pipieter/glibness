grammar Glibness;

WHITESPACE: [ \t\f\r] -> skip;
NEWLINE: '\n';

DIALOGUE: 'dialogue';
SET: 'set';
SAY: 'say';

RBRACE: '}';
LBRACE: '{';
STRING: '"' (~["\n])* '"' ;
VARIABLE: [A-Za-z_] [A-Za-z_0-9]*;

program: NEWLINE* (dialogue NEWLINE*)* EOF;

dialogue:
	DIALOGUE name=VARIABLE LBRACE NEWLINE statements RBRACE;

statements: statement*;

statement: sayStatement | setStatement;

sayStatement: SAY val=value NEWLINE;
setStatement: SET variable=VARIABLE val=value NEWLINE;

value: STRING; // Only support strings for now