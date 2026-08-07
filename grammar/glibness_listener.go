// Code generated from Glibness.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Glibness

import "github.com/antlr4-go/antlr/v4"

// GlibnessListener is a complete listener for a parse tree produced by GlibnessParser.
type GlibnessListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDialogue is called when entering the dialogue production.
	EnterDialogue(c *DialogueContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSayStatement is called when entering the sayStatement production.
	EnterSayStatement(c *SayStatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDialogue is called when exiting the dialogue production.
	ExitDialogue(c *DialogueContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSayStatement is called when exiting the sayStatement production.
	ExitSayStatement(c *SayStatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
