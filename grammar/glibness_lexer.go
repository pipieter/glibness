// Code generated from Glibness.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GlibnessLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GlibnessLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func glibnesslexerLexerInit() {
	staticData := &GlibnessLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'\\n'", "'dialogue'", "'set'", "'say'", "'}'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "RBRACE", "LBRACE",
		"STRING", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "RBRACE", "LBRACE",
		"STRING", "VARIABLE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 62, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 5, 7, 49, 8, 7, 10, 7, 12, 7, 52, 9, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 5, 8, 58, 8, 8, 10, 8, 12, 8, 61, 9, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 4, 3, 0, 9, 9, 12, 13, 32,
		32, 2, 0, 10, 10, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 63, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3,
		23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 34, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0,
		11, 42, 1, 0, 0, 0, 13, 44, 1, 0, 0, 0, 15, 46, 1, 0, 0, 0, 17, 55, 1,
		0, 0, 0, 19, 20, 7, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 6, 0, 0, 0, 22,
		2, 1, 0, 0, 0, 23, 24, 5, 10, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 100,
		0, 0, 26, 27, 5, 105, 0, 0, 27, 28, 5, 97, 0, 0, 28, 29, 5, 108, 0, 0,
		29, 30, 5, 111, 0, 0, 30, 31, 5, 103, 0, 0, 31, 32, 5, 117, 0, 0, 32, 33,
		5, 101, 0, 0, 33, 6, 1, 0, 0, 0, 34, 35, 5, 115, 0, 0, 35, 36, 5, 101,
		0, 0, 36, 37, 5, 116, 0, 0, 37, 8, 1, 0, 0, 0, 38, 39, 5, 115, 0, 0, 39,
		40, 5, 97, 0, 0, 40, 41, 5, 121, 0, 0, 41, 10, 1, 0, 0, 0, 42, 43, 5, 125,
		0, 0, 43, 12, 1, 0, 0, 0, 44, 45, 5, 123, 0, 0, 45, 14, 1, 0, 0, 0, 46,
		50, 5, 34, 0, 0, 47, 49, 8, 1, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0,
		0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53, 1, 0, 0, 0, 52, 50,
		1, 0, 0, 0, 53, 54, 5, 34, 0, 0, 54, 16, 1, 0, 0, 0, 55, 59, 7, 2, 0, 0,
		56, 58, 7, 3, 0, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 18, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 3,
		0, 50, 59, 1, 6, 0, 0,
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

// GlibnessLexerInit initializes any static state used to implement GlibnessLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGlibnessLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GlibnessLexerInit() {
	staticData := &GlibnessLexerLexerStaticData
	staticData.once.Do(glibnesslexerLexerInit)
}

// NewGlibnessLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGlibnessLexer(input antlr.CharStream) *GlibnessLexer {
	GlibnessLexerInit()
	l := new(GlibnessLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GlibnessLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Glibness.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GlibnessLexer tokens.
const (
	GlibnessLexerWHITESPACE = 1
	GlibnessLexerNEWLINE    = 2
	GlibnessLexerDIALOGUE   = 3
	GlibnessLexerSET        = 4
	GlibnessLexerSAY        = 5
	GlibnessLexerRBRACE     = 6
	GlibnessLexerLBRACE     = 7
	GlibnessLexerSTRING     = 8
	GlibnessLexerVARIABLE   = 9
)
