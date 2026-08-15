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
		"", "", "", "'dialogue'", "'set'", "'say'", "'choose'", "'choice'",
		"'}'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "CHOOSE", "CHOICE",
		"RBRACE", "LBRACE", "STRING", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"program", "dialogue", "statementBlock", "statement", "sayStatement",
		"setStatement", "chooseStatement", "choiceBlock", "choice", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 90, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 5, 0, 29, 8, 0, 10, 0,
		12, 0, 32, 9, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 48, 8, 2, 10, 2, 12, 2, 51,
		9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 58, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 5, 7, 76, 8, 7, 10, 7, 12, 7, 79, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 0, 0, 86, 0, 23, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0,
		6, 57, 1, 0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 63, 1, 0, 0, 0, 12, 68, 1, 0,
		0, 0, 14, 72, 1, 0, 0, 0, 16, 82, 1, 0, 0, 0, 18, 87, 1, 0, 0, 0, 20, 22,
		5, 2, 0, 0, 21, 20, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0,
		23, 24, 1, 0, 0, 0, 24, 35, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 30, 3,
		2, 1, 0, 27, 29, 5, 2, 0, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30,
		28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0,
		0, 33, 26, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1,
		39, 1, 1, 0, 0, 0, 40, 41, 5, 3, 0, 0, 41, 42, 5, 11, 0, 0, 42, 43, 3,
		4, 2, 0, 43, 3, 1, 0, 0, 0, 44, 45, 5, 9, 0, 0, 45, 49, 5, 2, 0, 0, 46,
		48, 3, 6, 3, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0,
		0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 53,
		5, 8, 0, 0, 53, 5, 1, 0, 0, 0, 54, 58, 3, 8, 4, 0, 55, 58, 3, 10, 5, 0,
		56, 58, 3, 12, 6, 0, 57, 54, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 56, 1,
		0, 0, 0, 58, 7, 1, 0, 0, 0, 59, 60, 5, 5, 0, 0, 60, 61, 3, 18, 9, 0, 61,
		62, 5, 2, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 4, 0, 0, 64, 65, 5, 11, 0,
		0, 65, 66, 3, 18, 9, 0, 66, 67, 5, 2, 0, 0, 67, 11, 1, 0, 0, 0, 68, 69,
		5, 6, 0, 0, 69, 70, 3, 14, 7, 0, 70, 71, 5, 2, 0, 0, 71, 13, 1, 0, 0, 0,
		72, 73, 5, 9, 0, 0, 73, 77, 5, 2, 0, 0, 74, 76, 3, 16, 8, 0, 75, 74, 1,
		0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 8, 0, 0, 81, 15, 1, 0, 0,
		0, 82, 83, 5, 7, 0, 0, 83, 84, 5, 10, 0, 0, 84, 85, 3, 4, 2, 0, 85, 86,
		5, 2, 0, 0, 86, 17, 1, 0, 0, 0, 87, 88, 5, 10, 0, 0, 88, 19, 1, 0, 0, 0,
		6, 23, 30, 35, 49, 57, 77,
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
	GlibnessParserCHOOSE     = 6
	GlibnessParserCHOICE     = 7
	GlibnessParserRBRACE     = 8
	GlibnessParserLBRACE     = 9
	GlibnessParserSTRING     = 10
	GlibnessParserVARIABLE   = 11
)

// GlibnessParser rules.
const (
	GlibnessParserRULE_program         = 0
	GlibnessParserRULE_dialogue        = 1
	GlibnessParserRULE_statementBlock  = 2
	GlibnessParserRULE_statement       = 3
	GlibnessParserRULE_sayStatement    = 4
	GlibnessParserRULE_setStatement    = 5
	GlibnessParserRULE_chooseStatement = 6
	GlibnessParserRULE_choiceBlock     = 7
	GlibnessParserRULE_choice          = 8
	GlibnessParserRULE_value           = 9
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserNEWLINE {
		{
			p.SetState(20)
			p.Match(GlibnessParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserDIALOGUE {
		{
			p.SetState(26)
			p.Dialogue()
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GlibnessParserNEWLINE {
			{
				p.SetState(27)
				p.Match(GlibnessParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
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

	// GetBlock returns the block rule contexts.
	GetBlock() IStatementBlockContext

	// SetBlock sets the block rule contexts.
	SetBlock(IStatementBlockContext)

	// Getter signatures
	DIALOGUE() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	StatementBlock() IStatementBlockContext

	// IsDialogueContext differentiates from other interfaces.
	IsDialogueContext()
}

type DialogueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	block  IStatementBlockContext
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

func (s *DialogueContext) GetBlock() IStatementBlockContext { return s.block }

func (s *DialogueContext) SetBlock(v IStatementBlockContext) { s.block = v }

func (s *DialogueContext) DIALOGUE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserDIALOGUE, 0)
}

func (s *DialogueContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserVARIABLE, 0)
}

func (s *DialogueContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
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
		p.SetState(40)
		p.Match(GlibnessParserDIALOGUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)

		var _m = p.Match(GlibnessParserVARIABLE)

		localctx.(*DialogueContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)

		var _x = p.StatementBlock()

		localctx.(*DialogueContext).block = _x
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

// IStatementBlockContext is an interface to support dynamic dispatch.
type IStatementBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementBlockContext differentiates from other interfaces.
	IsStatementBlockContext()
}

type StatementBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBlockContext() *StatementBlockContext {
	var p = new(StatementBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statementBlock
	return p
}

func InitEmptyStatementBlockContext(p *StatementBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_statementBlock
}

func (*StatementBlockContext) IsStatementBlockContext() {}

func NewStatementBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBlockContext {
	var p = new(StatementBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_statementBlock

	return p
}

func (s *StatementBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserLBRACE, 0)
}

func (s *StatementBlockContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *StatementBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserRBRACE, 0)
}

func (s *StatementBlockContext) AllStatement() []IStatementContext {
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

func (s *StatementBlockContext) Statement(i int) IStatementContext {
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

func (s *StatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterStatementBlock(s)
	}
}

func (s *StatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitStatementBlock(s)
	}
}

func (p *GlibnessParser) StatementBlock() (localctx IStatementBlockContext) {
	localctx = NewStatementBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GlibnessParserRULE_statementBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(GlibnessParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(GlibnessParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0 {
		{
			p.SetState(46)
			p.Statement()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SayStatement() ISayStatementContext
	SetStatement() ISetStatementContext
	ChooseStatement() IChooseStatementContext

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

func (s *StatementContext) ChooseStatement() IChooseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChooseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChooseStatementContext)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GlibnessParserSAY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.SayStatement()
		}

	case GlibnessParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.SetStatement()
		}

	case GlibnessParserCHOOSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.ChooseStatement()
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
		p.SetState(59)
		p.Match(GlibnessParserSAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)

		var _x = p.Value()

		localctx.(*SayStatementContext).val = _x
	}
	{
		p.SetState(61)
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
		p.SetState(63)
		p.Match(GlibnessParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)

		var _m = p.Match(GlibnessParserVARIABLE)

		localctx.(*SetStatementContext).variable = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)

		var _x = p.Value()

		localctx.(*SetStatementContext).val = _x
	}
	{
		p.SetState(66)
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

// IChooseStatementContext is an interface to support dynamic dispatch.
type IChooseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBlock returns the block rule contexts.
	GetBlock() IChoiceBlockContext

	// SetBlock sets the block rule contexts.
	SetBlock(IChoiceBlockContext)

	// Getter signatures
	CHOOSE() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	ChoiceBlock() IChoiceBlockContext

	// IsChooseStatementContext differentiates from other interfaces.
	IsChooseStatementContext()
}

type ChooseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	block  IChoiceBlockContext
}

func NewEmptyChooseStatementContext() *ChooseStatementContext {
	var p = new(ChooseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_chooseStatement
	return p
}

func InitEmptyChooseStatementContext(p *ChooseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_chooseStatement
}

func (*ChooseStatementContext) IsChooseStatementContext() {}

func NewChooseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChooseStatementContext {
	var p = new(ChooseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_chooseStatement

	return p
}

func (s *ChooseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ChooseStatementContext) GetBlock() IChoiceBlockContext { return s.block }

func (s *ChooseStatementContext) SetBlock(v IChoiceBlockContext) { s.block = v }

func (s *ChooseStatementContext) CHOOSE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserCHOOSE, 0)
}

func (s *ChooseStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *ChooseStatementContext) ChoiceBlock() IChoiceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChoiceBlockContext)
}

func (s *ChooseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChooseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChooseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterChooseStatement(s)
	}
}

func (s *ChooseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitChooseStatement(s)
	}
}

func (p *GlibnessParser) ChooseStatement() (localctx IChooseStatementContext) {
	localctx = NewChooseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GlibnessParserRULE_chooseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(GlibnessParserCHOOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)

		var _x = p.ChoiceBlock()

		localctx.(*ChooseStatementContext).block = _x
	}
	{
		p.SetState(70)
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

// IChoiceBlockContext is an interface to support dynamic dispatch.
type IChoiceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetChoices returns the choices rule contexts.
	GetChoices() IChoiceContext

	// SetChoices sets the choices rule contexts.
	SetChoices(IChoiceContext)

	// Getter signatures
	LBRACE() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllChoice() []IChoiceContext
	Choice(i int) IChoiceContext

	// IsChoiceBlockContext differentiates from other interfaces.
	IsChoiceBlockContext()
}

type ChoiceBlockContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	choices IChoiceContext
}

func NewEmptyChoiceBlockContext() *ChoiceBlockContext {
	var p = new(ChoiceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_choiceBlock
	return p
}

func InitEmptyChoiceBlockContext(p *ChoiceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_choiceBlock
}

func (*ChoiceBlockContext) IsChoiceBlockContext() {}

func NewChoiceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceBlockContext {
	var p = new(ChoiceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_choiceBlock

	return p
}

func (s *ChoiceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceBlockContext) GetChoices() IChoiceContext { return s.choices }

func (s *ChoiceBlockContext) SetChoices(v IChoiceContext) { s.choices = v }

func (s *ChoiceBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserLBRACE, 0)
}

func (s *ChoiceBlockContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *ChoiceBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserRBRACE, 0)
}

func (s *ChoiceBlockContext) AllChoice() []IChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChoiceContext); ok {
			len++
		}
	}

	tst := make([]IChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChoiceContext); ok {
			tst[i] = t.(IChoiceContext)
			i++
		}
	}

	return tst
}

func (s *ChoiceBlockContext) Choice(i int) IChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceContext); ok {
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

	return t.(IChoiceContext)
}

func (s *ChoiceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChoiceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterChoiceBlock(s)
	}
}

func (s *ChoiceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitChoiceBlock(s)
	}
}

func (p *GlibnessParser) ChoiceBlock() (localctx IChoiceBlockContext) {
	localctx = NewChoiceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GlibnessParserRULE_choiceBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(GlibnessParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(GlibnessParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GlibnessParserCHOICE {
		{
			p.SetState(74)

			var _x = p.Choice()

			localctx.(*ChoiceBlockContext).choices = _x
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
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

// IChoiceContext is an interface to support dynamic dispatch.
type IChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetBlock returns the block rule contexts.
	GetBlock() IStatementBlockContext

	// SetBlock sets the block rule contexts.
	SetBlock(IStatementBlockContext)

	// Getter signatures
	CHOICE() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	STRING() antlr.TerminalNode
	StatementBlock() IStatementBlockContext

	// IsChoiceContext differentiates from other interfaces.
	IsChoiceContext()
}

type ChoiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	block  IStatementBlockContext
}

func NewEmptyChoiceContext() *ChoiceContext {
	var p = new(ChoiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_choice
	return p
}

func InitEmptyChoiceContext(p *ChoiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GlibnessParserRULE_choice
}

func (*ChoiceContext) IsChoiceContext() {}

func NewChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceContext {
	var p = new(ChoiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GlibnessParserRULE_choice

	return p
}

func (s *ChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceContext) GetName() antlr.Token { return s.name }

func (s *ChoiceContext) SetName(v antlr.Token) { s.name = v }

func (s *ChoiceContext) GetBlock() IStatementBlockContext { return s.block }

func (s *ChoiceContext) SetBlock(v IStatementBlockContext) { s.block = v }

func (s *ChoiceContext) CHOICE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserCHOICE, 0)
}

func (s *ChoiceContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GlibnessParserNEWLINE, 0)
}

func (s *ChoiceContext) STRING() antlr.TerminalNode {
	return s.GetToken(GlibnessParserSTRING, 0)
}

func (s *ChoiceContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChoiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.EnterChoice(s)
	}
}

func (s *ChoiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GlibnessListener); ok {
		listenerT.ExitChoice(s)
	}
}

func (p *GlibnessParser) Choice() (localctx IChoiceContext) {
	localctx = NewChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GlibnessParserRULE_choice)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(GlibnessParserCHOICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)

		var _m = p.Match(GlibnessParserSTRING)

		localctx.(*ChoiceContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)

		var _x = p.StatementBlock()

		localctx.(*ChoiceContext).block = _x
	}
	{
		p.SetState(85)
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
	p.EnterRule(localctx, 18, GlibnessParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
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
