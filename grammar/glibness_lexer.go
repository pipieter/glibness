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
		"", "", "", "'dialogue'", "'set'", "'say'", "'choose'", "'choice'",
		"'}'", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "CHOOSE", "CHOICE",
		"RBRACE", "LBRACE", "STRING", "INTEGER", "BOOLEAN", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "DIALOGUE", "SET", "SAY", "CHOOSE", "CHOICE",
		"RBRACE", "LBRACE", "STRING", "INTEGER", "BOOLEAN", "VARIABLE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 109, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 5, 1, 34, 8, 1, 10, 1, 12, 1, 37, 9, 1, 4, 1, 39, 8, 1, 11, 1, 12, 1,
		40, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 5, 9, 80, 8, 9, 10, 9, 12, 9, 83, 9, 9, 1, 9, 1, 9, 1, 10, 4,
		10, 88, 8, 10, 11, 10, 12, 10, 89, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 3, 11, 101, 8, 11, 1, 12, 1, 12, 5, 12, 105, 8,
		12, 10, 12, 12, 12, 108, 9, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 5, 3,
		0, 9, 9, 12, 13, 32, 32, 2, 0, 10, 10, 34, 34, 1, 0, 48, 57, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 114, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 38, 1, 0, 0, 0, 5, 42, 1, 0, 0,
		0, 7, 51, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 59, 1, 0, 0, 0, 13, 66, 1,
		0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0, 21,
		87, 1, 0, 0, 0, 23, 100, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 28, 7, 0,
		0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 6, 0, 0, 0, 30, 2, 1, 0, 0, 0, 31, 35,
		5, 10, 0, 0, 32, 34, 3, 1, 0, 0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1,
		0, 0, 0, 38, 31, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40,
		41, 1, 0, 0, 0, 41, 4, 1, 0, 0, 0, 42, 43, 5, 100, 0, 0, 43, 44, 5, 105,
		0, 0, 44, 45, 5, 97, 0, 0, 45, 46, 5, 108, 0, 0, 46, 47, 5, 111, 0, 0,
		47, 48, 5, 103, 0, 0, 48, 49, 5, 117, 0, 0, 49, 50, 5, 101, 0, 0, 50, 6,
		1, 0, 0, 0, 51, 52, 5, 115, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 116,
		0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5, 115, 0, 0, 56, 57, 5, 97, 0, 0, 57,
		58, 5, 121, 0, 0, 58, 10, 1, 0, 0, 0, 59, 60, 5, 99, 0, 0, 60, 61, 5, 104,
		0, 0, 61, 62, 5, 111, 0, 0, 62, 63, 5, 111, 0, 0, 63, 64, 5, 115, 0, 0,
		64, 65, 5, 101, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 99, 0, 0, 67, 68,
		5, 104, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 105, 0, 0, 70, 71, 5, 99,
		0, 0, 71, 72, 5, 101, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 125, 0, 0, 74,
		16, 1, 0, 0, 0, 75, 76, 5, 123, 0, 0, 76, 18, 1, 0, 0, 0, 77, 81, 5, 34,
		0, 0, 78, 80, 8, 1, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79,
		1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		84, 85, 5, 34, 0, 0, 85, 20, 1, 0, 0, 0, 86, 88, 7, 2, 0, 0, 87, 86, 1,
		0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90,
		22, 1, 0, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 114, 0, 0, 93, 94, 5,
		117, 0, 0, 94, 101, 5, 101, 0, 0, 95, 96, 5, 102, 0, 0, 96, 97, 5, 97,
		0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 115, 0, 0, 99, 101, 5, 101, 0, 0,
		100, 91, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 101, 24, 1, 0, 0, 0, 102, 106,
		7, 3, 0, 0, 103, 105, 7, 4, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 26, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 7, 0, 35, 40, 81, 89, 100, 106, 1, 6, 0, 0,
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
	GlibnessLexerINTEGER    = 11
	GlibnessLexerBOOLEAN    = 12
	GlibnessLexerVARIABLE   = 13
)
