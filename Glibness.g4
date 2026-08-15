grammar Glibness;

WHITESPACE: [ \t\f\r] -> skip;
NEWLINE: ('\n' WHITESPACE*)+;

DIALOGUE: 'dialogue';
SET: 'set';
SAY: 'say';
CHOOSE: 'choose';
CHOICE: 'choice';

RBRACE: '}';
LBRACE: '{';
STRING: '"' (~["\n])* '"' ;
VARIABLE: [A-Za-z_] [A-Za-z_0-9]*;

program: NEWLINE* (dialogue NEWLINE*)* EOF;

dialogue:
	DIALOGUE name=VARIABLE block=statementBlock;

statementBlock: LBRACE NEWLINE statement* RBRACE;

statement: sayStatement | setStatement | chooseStatement;

sayStatement: SAY val=value NEWLINE;
setStatement: SET variable=VARIABLE val=value NEWLINE;
chooseStatement: CHOOSE block=choiceBlock NEWLINE;

choiceBlock: LBRACE NEWLINE choices=choice* RBRACE;
choice: CHOICE name=STRING block=statementBlock NEWLINE;

value: STRING; // Only support strings for now