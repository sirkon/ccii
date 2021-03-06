// Code generated from CCII.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 34, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 6, 8, 31, 10, 8, 13, 8, 14, 8, 32, 2, 2, 9, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 3, 12, 2, 36, 36, 38, 38,
	40, 41, 46, 46, 60, 60, 62, 64, 93, 95, 98, 98, 125, 125, 127, 127, 2,
	34, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2,
	2, 2, 5, 19, 3, 2, 2, 2, 7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 25,
	3, 2, 2, 2, 13, 27, 3, 2, 2, 2, 15, 30, 3, 2, 2, 2, 17, 18, 7, 125, 2,
	2, 18, 4, 3, 2, 2, 2, 19, 20, 7, 127, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22,
	7, 63, 2, 2, 22, 8, 3, 2, 2, 2, 23, 24, 7, 46, 2, 2, 24, 10, 3, 2, 2, 2,
	25, 26, 7, 93, 2, 2, 26, 12, 3, 2, 2, 2, 27, 28, 7, 95, 2, 2, 28, 14, 3,
	2, 2, 2, 29, 31, 10, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 16, 3, 2, 2, 2, 4, 2, 32, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'", "'='", "','", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "SCALAR",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "SCALAR",
}

type CCIILexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCCIILexer(input antlr.CharStream) *CCIILexer {

	l := new(CCIILexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CCII.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CCIILexer tokens.
const (
	CCIILexerT__0   = 1
	CCIILexerT__1   = 2
	CCIILexerT__2   = 3
	CCIILexerT__3   = 4
	CCIILexerT__4   = 5
	CCIILexerT__5   = 6
	CCIILexerSCALAR = 7
)
