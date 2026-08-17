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
		4, 0, 13, 118, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 4, 0, 29, 8, 0, 11, 0, 12,
		0, 30, 1, 0, 1, 0, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 1,
		1, 1, 5, 1, 43, 8, 1, 10, 1, 12, 1, 46, 9, 1, 4, 1, 48, 8, 1, 11, 1, 12,
		1, 49, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 89, 8, 9, 10, 9, 12, 9, 92, 9, 9, 1, 9, 1, 9, 1, 10,
		4, 10, 97, 8, 10, 11, 10, 12, 10, 98, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 110, 8, 11, 1, 12, 1, 12, 5, 12, 114,
		8, 12, 10, 12, 12, 12, 117, 9, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 5,
		3, 0, 9, 9, 12, 13, 32, 32, 2, 0, 10, 10, 34, 34, 1, 0, 48, 57, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 125, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 1, 28, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 51, 1, 0, 0,
		0, 7, 60, 1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 75, 1,
		0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 19, 86, 1, 0, 0, 0, 21,
		96, 1, 0, 0, 0, 23, 109, 1, 0, 0, 0, 25, 111, 1, 0, 0, 0, 27, 29, 7, 0,
		0, 0, 28, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31,
		1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 6, 0, 0, 0, 33, 2, 1, 0, 0, 0,
		34, 36, 3, 1, 0, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1,
		0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40,
		44, 5, 10, 0, 0, 41, 43, 3, 1, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0,
		0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 47, 37, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 100, 0, 0, 52, 53, 5,
		105, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 111, 0,
		0, 56, 57, 5, 103, 0, 0, 57, 58, 5, 117, 0, 0, 58, 59, 5, 101, 0, 0, 59,
		6, 1, 0, 0, 0, 60, 61, 5, 115, 0, 0, 61, 62, 5, 101, 0, 0, 62, 63, 5, 116,
		0, 0, 63, 8, 1, 0, 0, 0, 64, 65, 5, 115, 0, 0, 65, 66, 5, 97, 0, 0, 66,
		67, 5, 121, 0, 0, 67, 10, 1, 0, 0, 0, 68, 69, 5, 99, 0, 0, 69, 70, 5, 104,
		0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 111, 0, 0, 72, 73, 5, 115, 0, 0,
		73, 74, 5, 101, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 5, 99, 0, 0, 76, 77,
		5, 104, 0, 0, 77, 78, 5, 111, 0, 0, 78, 79, 5, 105, 0, 0, 79, 80, 5, 99,
		0, 0, 80, 81, 5, 101, 0, 0, 81, 14, 1, 0, 0, 0, 82, 83, 5, 125, 0, 0, 83,
		16, 1, 0, 0, 0, 84, 85, 5, 123, 0, 0, 85, 18, 1, 0, 0, 0, 86, 90, 5, 34,
		0, 0, 87, 89, 8, 1, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		93, 94, 5, 34, 0, 0, 94, 20, 1, 0, 0, 0, 95, 97, 7, 2, 0, 0, 96, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		22, 1, 0, 0, 0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 114, 0, 0, 102, 103,
		5, 117, 0, 0, 103, 110, 5, 101, 0, 0, 104, 105, 5, 102, 0, 0, 105, 106,
		5, 97, 0, 0, 106, 107, 5, 108, 0, 0, 107, 108, 5, 115, 0, 0, 108, 110,
		5, 101, 0, 0, 109, 100, 1, 0, 0, 0, 109, 104, 1, 0, 0, 0, 110, 24, 1, 0,
		0, 0, 111, 115, 7, 3, 0, 0, 112, 114, 7, 4, 0, 0, 113, 112, 1, 0, 0, 0,
		114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		26, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 9, 0, 30, 37, 44, 49, 90, 98, 109,
		115, 1, 6, 0, 0,
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
