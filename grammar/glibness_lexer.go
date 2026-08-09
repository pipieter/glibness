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
		"", "", "'\\n'", "'dialogue'", "'set'", "'say'", "'choose'", "'choice'",
		"'}'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "CHOOSE", "CHOICE",
		"RBRACE", "LBRACE", "STRING", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "CHOOSE", "CHOICE",
		"RBRACE", "LBRACE", "STRING", "VARIABLE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 80, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 67, 8, 9, 10,
		9, 12, 9, 70, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 76, 8, 10, 10, 10,
		12, 10, 79, 9, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 4, 3, 0, 9, 9, 12, 13, 32, 32, 2, 0,
		10, 10, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 81, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 1, 23, 1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 38, 1,
		0, 0, 0, 9, 42, 1, 0, 0, 0, 11, 46, 1, 0, 0, 0, 13, 53, 1, 0, 0, 0, 15,
		60, 1, 0, 0, 0, 17, 62, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 73, 1, 0, 0,
		0, 23, 24, 7, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 6, 0, 0, 0, 26, 2, 1,
		0, 0, 0, 27, 28, 5, 10, 0, 0, 28, 4, 1, 0, 0, 0, 29, 30, 5, 100, 0, 0,
		30, 31, 5, 105, 0, 0, 31, 32, 5, 97, 0, 0, 32, 33, 5, 108, 0, 0, 33, 34,
		5, 111, 0, 0, 34, 35, 5, 103, 0, 0, 35, 36, 5, 117, 0, 0, 36, 37, 5, 101,
		0, 0, 37, 6, 1, 0, 0, 0, 38, 39, 5, 115, 0, 0, 39, 40, 5, 101, 0, 0, 40,
		41, 5, 116, 0, 0, 41, 8, 1, 0, 0, 0, 42, 43, 5, 115, 0, 0, 43, 44, 5, 97,
		0, 0, 44, 45, 5, 121, 0, 0, 45, 10, 1, 0, 0, 0, 46, 47, 5, 99, 0, 0, 47,
		48, 5, 104, 0, 0, 48, 49, 5, 111, 0, 0, 49, 50, 5, 111, 0, 0, 50, 51, 5,
		115, 0, 0, 51, 52, 5, 101, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 99, 0,
		0, 54, 55, 5, 104, 0, 0, 55, 56, 5, 111, 0, 0, 56, 57, 5, 105, 0, 0, 57,
		58, 5, 99, 0, 0, 58, 59, 5, 101, 0, 0, 59, 14, 1, 0, 0, 0, 60, 61, 5, 125,
		0, 0, 61, 16, 1, 0, 0, 0, 62, 63, 5, 123, 0, 0, 63, 18, 1, 0, 0, 0, 64,
		68, 5, 34, 0, 0, 65, 67, 8, 1, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0,
		0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68,
		1, 0, 0, 0, 71, 72, 5, 34, 0, 0, 72, 20, 1, 0, 0, 0, 73, 77, 7, 2, 0, 0,
		74, 76, 7, 3, 0, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 22, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 3,
		0, 68, 77, 1, 6, 0, 0,
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
	GlibnessLexerCHOOSE     = 6
	GlibnessLexerCHOICE     = 7
	GlibnessLexerRBRACE     = 8
	GlibnessLexerLBRACE     = 9
	GlibnessLexerSTRING     = 10
	GlibnessLexerVARIABLE   = 11
)
