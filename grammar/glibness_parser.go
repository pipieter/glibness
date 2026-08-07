// Code generated from Glibness.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Glibness

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GlibnessParser struct {
	*antlr.BaseParser
}

var GlibnessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func glibnessParserInit() {
	staticData := &GlibnessParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'\\n'", "'dialogue'", "'set'", "'say'", "'}'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "RBRACE", "LBRACE",
		"STRING", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"program", "dialogue", "statements", "statement", "sayStatement", "setStatement",
		"value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 63, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0, 1,
		0, 1, 0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 5, 0, 28, 8, 0, 10, 0,
		12, 0, 31, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 5, 2, 43, 8, 2, 10, 2, 12, 2, 46, 9, 2, 1, 3, 1, 3, 3, 3, 50, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 60, 0, 17, 1, 0, 0, 0, 2, 34,
		1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10,
		55, 1, 0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 16, 5, 2, 0, 0, 15, 14, 1, 0, 0,
		0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 29,
		1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 24, 3, 2, 1, 0, 21, 23, 5, 2, 0, 0,
		22, 21, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1,
		0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 20, 1, 0, 0, 0, 28,
		31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0,
		0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34, 35, 5,
		3, 0, 0, 35, 36, 5, 9, 0, 0, 36, 37, 5, 7, 0, 0, 37, 38, 5, 2, 0, 0, 38,
		39, 3, 4, 2, 0, 39, 40, 5, 6, 0, 0, 40, 3, 1, 0, 0, 0, 41, 43, 3, 6, 3,
		0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45,
		1, 0, 0, 0, 45, 5, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 50, 3, 8, 4, 0,
		48, 50, 3, 10, 5, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 7, 1,
		0, 0, 0, 51, 52, 5, 5, 0, 0, 52, 53, 3, 12, 6, 0, 53, 54, 5, 2, 0, 0, 54,
		9, 1, 0, 0, 0, 55, 56, 5, 4, 0, 0, 56, 57, 5, 9, 0, 0, 57, 58, 3, 12, 6,
		0, 58, 59, 5, 2, 0, 0, 59, 11, 1, 0, 0, 0, 60, 61, 5, 8, 0, 0, 61, 13,
		1, 0, 0, 0, 5, 17, 24, 29, 44, 49,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GlibnessParserInit initializes any static state used to implement GlibnessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGlibnessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GlibnessParserInit() {
	staticData := &GlibnessParserStaticData
	staticData.once.Do(glibnessParserInit)
}

// NewGlibnessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGlibnessParser(input antlr.TokenStream) *GlibnessParser {
	GlibnessParserInit()
	this := new(GlibnessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GlibnessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Glibness.g4"

	return this
}

// GlibnessParser tokens.
const (
	GlibnessParserEOF        = antlr.TokenEOF
	GlibnessParserWHITESPACE = 1
	GlibnessParserNEWLINE    = 2
	GlibnessParserDIALOGUE   = 3
	GlibnessParserSET        = 4
	GlibnessParserSAY        = 5
	GlibnessParserRBRACE     = 6
	GlibnessParserLBRACE     = 7
	GlibnessParserSTRING     = 8
	GlibnessParserVARIABLE   = 9
)

// GlibnessParser rules.
const (
	GlibnessParserRULE_program      = 0
	GlibnessParserRULE_dialogue     = 1
	GlibnessParserRULE_statements   = 2
	GlibnessParserRULE_statement    = 3
	GlibnessParserRULE_sayStatement = 4
	GlibnessParserRULE_setStatement = 5
	GlibnessParserRULE_value        = 6
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllDialogue() []IDialogueContext
	Dialogue(i int) IDialogueContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GlibnessParserEOF, 0)
}

func (s *ProgramContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GlibnessParserNEWLINE)
}

func (s *ProgramContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, i)
}

func (s *ProgramContext) AllDialogue() []IDialogueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDialogueContext); ok {
			len++
		}
	}

	tst := make([]IDialogueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDialogueContext); ok {
			tst[i] = t.(IDialogueContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Dialogue(i int) IDialogueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDialogueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDialogueContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GlibnessParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GlibnessParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserNEWLINE {
		{
			p.SetState(14)
			p.Match(GlibnessParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserDIALOGUE {
		{
			p.SetState(20)
			p.Dialogue()
		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GlibnessParserNEWLINE {
			{
				p.SetState(21)
				p.Match(GlibnessParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(GlibnessParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDialogueContext is an interface to support dynamic dispatch.
type IDialogueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	DIALOGUE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	Statements() IStatementsContext
	RBRACE() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode

	// IsDialogueContext differentiates from other interfaces.
	IsDialogueContext()
}

type DialogueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyDialogueContext() *DialogueContext {
	var p = new(DialogueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_dialogue
	return p
}

func InitEmptyDialogueContext(p *DialogueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_dialogue
}

func (*DialogueContext) IsDialogueContext() {}

func NewDialogueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DialogueContext {
	var p = new(DialogueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_dialogue

	return p
}

func (s *DialogueContext) GetParser() antlr.Parser { return s.parser }

func (s *DialogueContext) GetName() antlr.Token { return s.name }

func (s *DialogueContext) SetName(v antlr.Token) { s.name = v }

func (s *DialogueContext) DIALOGUE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserDIALOGUE, 0)
}

func (s *DialogueContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserLBRACE, 0)
}

func (s *DialogueContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *DialogueContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *DialogueContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserRBRACE, 0)
}

func (s *DialogueContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserVARIABLE, 0)
}

func (s *DialogueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DialogueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DialogueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterDialogue(s)
	}
}

func (s *DialogueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitDialogue(s)
	}
}

func (p *GlibnessParser) Dialogue() (localctx IDialogueContext) {
	localctx = NewDialogueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GlibnessParserRULE_dialogue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(GlibnessParserDIALOGUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)

		var _m = p.Match(GlibnessParserVARIABLE)

		localctx.(*DialogueContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(GlibnessParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(GlibnessParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Statements()
	}
	{
		p.SetState(39)
		p.Match(GlibnessParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GlibnessParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GlibnessParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserSET || _la == GlibnessParserSAY {
		{
			p.SetState(41)
			p.Statement()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SayStatement() ISayStatementContext
	SetStatement() ISetStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) SayStatement() ISayStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISayStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISayStatementContext)
}

func (s *StatementContext) SetStatement() ISetStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GlibnessParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GlibnessParserRULE_statement)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GlibnessParserSAY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.SayStatement()
		}

	case GlibnessParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.SetStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISayStatementContext is an interface to support dynamic dispatch.
type ISayStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val rule contexts.
	GetVal() IValueContext

	// SetVal sets the val rule contexts.
	SetVal(IValueContext)

	// Getter signatures
	SAY() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	Value() IValueContext

	// IsSayStatementContext differentiates from other interfaces.
	IsSayStatementContext()
}

type SayStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    IValueContext
}

func NewEmptySayStatementContext() *SayStatementContext {
	var p = new(SayStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_sayStatement
	return p
}

func InitEmptySayStatementContext(p *SayStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_sayStatement
}

func (*SayStatementContext) IsSayStatementContext() {}

func NewSayStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SayStatementContext {
	var p = new(SayStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_sayStatement

	return p
}

func (s *SayStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SayStatementContext) GetVal() IValueContext { return s.val }

func (s *SayStatementContext) SetVal(v IValueContext) { s.val = v }

func (s *SayStatementContext) SAY() antlr.TerminalNode {
	return s.GetToken(GlibnessParserSAY, 0)
}

func (s *SayStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *SayStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SayStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SayStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SayStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterSayStatement(s)
	}
}

func (s *SayStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitSayStatement(s)
	}
}

func (p *GlibnessParser) SayStatement() (localctx ISayStatementContext) {
	localctx = NewSayStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GlibnessParserRULE_sayStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(GlibnessParserSAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)

		var _x = p.Value()

		localctx.(*SayStatementContext).val = _x
	}
	{
		p.SetState(53)
		p.Match(GlibnessParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetStatementContext is an interface to support dynamic dispatch.
type ISetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable token.
	GetVariable() antlr.Token

	// SetVariable sets the variable token.
	SetVariable(antlr.Token)

	// GetVal returns the val rule contexts.
	GetVal() IValueContext

	// SetVal sets the val rule contexts.
	SetVal(IValueContext)

	// Getter signatures
	SET() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	Value() IValueContext

	// IsSetStatementContext differentiates from other interfaces.
	IsSetStatementContext()
}

type SetStatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable antlr.Token
	val      IValueContext
}

func NewEmptySetStatementContext() *SetStatementContext {
	var p = new(SetStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_setStatement
	return p
}

func InitEmptySetStatementContext(p *SetStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_setStatement
}

func (*SetStatementContext) IsSetStatementContext() {}

func NewSetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetStatementContext {
	var p = new(SetStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_setStatement

	return p
}

func (s *SetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetStatementContext) GetVariable() antlr.Token { return s.variable }

func (s *SetStatementContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *SetStatementContext) GetVal() IValueContext { return s.val }

func (s *SetStatementContext) SetVal(v IValueContext) { s.val = v }

func (s *SetStatementContext) SET() antlr.TerminalNode {
	return s.GetToken(GlibnessParserSET, 0)
}

func (s *SetStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *SetStatementContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserVARIABLE, 0)
}

func (s *SetStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterSetStatement(s)
	}
}

func (s *SetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitSetStatement(s)
	}
}

func (p *GlibnessParser) SetStatement() (localctx ISetStatementContext) {
	localctx = NewSetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GlibnessParserRULE_setStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(GlibnessParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)

		var _m = p.Match(GlibnessParserVARIABLE)

		localctx.(*SetStatementContext).variable = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)

		var _x = p.Value()

		localctx.(*SetStatementContext).val = _x
	}
	{
		p.SetState(58)
		p.Match(GlibnessParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GlibnessParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *GlibnessParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GlibnessParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(GlibnessParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
