// Code generated from Glibness.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Glibness

import "github.com/antlr4-go/antlr/v4"

// BaseGlibnessListener is a complete listener for a parse tree produced by GlibnessParser.
type BaseGlibnessListener struct{}

var _ GlibnessListener = &BaseGlibnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGlibnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGlibnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGlibnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGlibnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGlibnessListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGlibnessListener) ExitProgram(ctx *ProgramContext) {}

// EnterDialogue is called when production dialogue is entered.
func (s *BaseGlibnessListener) EnterDialogue(ctx *DialogueContext) {}

// ExitDialogue is called when production dialogue is exited.
func (s *BaseGlibnessListener) ExitDialogue(ctx *DialogueContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseGlibnessListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseGlibnessListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGlibnessListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGlibnessListener) ExitStatement(ctx *StatementContext) {}

// EnterSayStatement is called when production sayStatement is entered.
func (s *BaseGlibnessListener) EnterSayStatement(ctx *SayStatementContext) {}

// ExitSayStatement is called when production sayStatement is exited.
func (s *BaseGlibnessListener) ExitSayStatement(ctx *SayStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseGlibnessListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseGlibnessListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterValue is called when production value is entered.
func (s *BaseGlibnessListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseGlibnessListener) ExitValue(ctx *ValueContext) {}
