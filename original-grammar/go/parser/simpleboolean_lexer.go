// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleBooleanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var simplebooleanlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplebooleanlexerLexerInit() {
	staticData := &simplebooleanlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "", "", "", "'IF'", "", "", "", "", "", "", "", "", "'('",
		"')'", "", "'SQRT'", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "AND", "OR", "NOT", "TRUE", "FALSE", "IF", "THEN", "ELSE", "GT",
		"GE", "LT", "LE", "EQ", "NEQ", "LPAREN", "RPAREN", "DECIMAL", "SQRT",
		"MULT", "DIV", "ADD", "SUB", "WS",
	}
	staticData.ruleNames = []string{
		"AND", "OR", "NOT", "TRUE", "FALSE", "IF", "THEN", "ELSE", "GT", "GE",
		"LT", "LE", "EQ", "NEQ", "LPAREN", "RPAREN", "DECIMAL", "SQRT", "MULT",
		"DIV", "ADD", "SUB", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 207, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 52, 8, 0,
		1, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 77, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 94, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 113, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 127, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 132,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 139, 8, 9, 1, 10, 1, 10, 1, 10,
		3, 10, 144, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 151, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 157, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 166, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 3, 16, 173, 8, 16, 1, 16, 4, 16, 176, 8, 16, 11, 16, 12, 16, 177,
		1, 16, 1, 16, 4, 16, 182, 8, 16, 11, 16, 12, 16, 183, 3, 16, 186, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 4, 22, 202, 8, 22, 11, 22, 12, 22, 203, 1, 22,
		1, 22, 0, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 1, 0, 2, 1, 0, 48, 57, 3, 0,
		9, 10, 12, 13, 32, 32, 232, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 56, 1, 0, 0, 0, 5, 62, 1,
		0, 0, 0, 7, 76, 1, 0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 95, 1, 0, 0, 0, 13,
		112, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17, 131, 1, 0, 0, 0, 19, 138, 1,
		0, 0, 0, 21, 143, 1, 0, 0, 0, 23, 150, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0,
		27, 165, 1, 0, 0, 0, 29, 167, 1, 0, 0, 0, 31, 169, 1, 0, 0, 0, 33, 172,
		1, 0, 0, 0, 35, 187, 1, 0, 0, 0, 37, 192, 1, 0, 0, 0, 39, 194, 1, 0, 0,
		0, 41, 196, 1, 0, 0, 0, 43, 198, 1, 0, 0, 0, 45, 201, 1, 0, 0, 0, 47, 48,
		5, 65, 0, 0, 48, 49, 5, 78, 0, 0, 49, 52, 5, 68, 0, 0, 50, 52, 5, 38, 0,
		0, 51, 47, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5,
		79, 0, 0, 54, 57, 5, 82, 0, 0, 55, 57, 5, 124, 0, 0, 56, 53, 1, 0, 0, 0,
		56, 55, 1, 0, 0, 0, 57, 4, 1, 0, 0, 0, 58, 59, 5, 78, 0, 0, 59, 60, 5,
		79, 0, 0, 60, 63, 5, 84, 0, 0, 61, 63, 5, 33, 0, 0, 62, 58, 1, 0, 0, 0,
		62, 61, 1, 0, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65, 5, 84, 0, 0, 65, 66, 5,
		82, 0, 0, 66, 67, 5, 85, 0, 0, 67, 77, 5, 69, 0, 0, 68, 69, 5, 84, 0, 0,
		69, 70, 5, 114, 0, 0, 70, 71, 5, 117, 0, 0, 71, 77, 5, 101, 0, 0, 72, 73,
		5, 116, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5, 117, 0, 0, 75, 77, 5, 101,
		0, 0, 76, 64, 1, 0, 0, 0, 76, 68, 1, 0, 0, 0, 76, 72, 1, 0, 0, 0, 77, 8,
		1, 0, 0, 0, 78, 79, 5, 70, 0, 0, 79, 80, 5, 65, 0, 0, 80, 81, 5, 76, 0,
		0, 81, 82, 5, 83, 0, 0, 82, 94, 5, 69, 0, 0, 83, 84, 5, 70, 0, 0, 84, 85,
		5, 97, 0, 0, 85, 86, 5, 108, 0, 0, 86, 87, 5, 115, 0, 0, 87, 94, 5, 101,
		0, 0, 88, 89, 5, 102, 0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 108, 0, 0,
		91, 92, 5, 115, 0, 0, 92, 94, 5, 101, 0, 0, 93, 78, 1, 0, 0, 0, 93, 83,
		1, 0, 0, 0, 93, 88, 1, 0, 0, 0, 94, 10, 1, 0, 0, 0, 95, 96, 5, 73, 0, 0,
		96, 97, 5, 70, 0, 0, 97, 12, 1, 0, 0, 0, 98, 99, 5, 84, 0, 0, 99, 100,
		5, 72, 0, 0, 100, 101, 5, 69, 0, 0, 101, 113, 5, 78, 0, 0, 102, 103, 5,
		84, 0, 0, 103, 104, 5, 104, 0, 0, 104, 105, 5, 101, 0, 0, 105, 113, 5,
		110, 0, 0, 106, 107, 5, 116, 0, 0, 107, 108, 5, 104, 0, 0, 108, 109, 5,
		101, 0, 0, 109, 113, 5, 110, 0, 0, 110, 111, 5, 45, 0, 0, 111, 113, 5,
		62, 0, 0, 112, 98, 1, 0, 0, 0, 112, 102, 1, 0, 0, 0, 112, 106, 1, 0, 0,
		0, 112, 110, 1, 0, 0, 0, 113, 14, 1, 0, 0, 0, 114, 115, 5, 69, 0, 0, 115,
		116, 5, 76, 0, 0, 116, 117, 5, 83, 0, 0, 117, 127, 5, 69, 0, 0, 118, 119,
		5, 69, 0, 0, 119, 120, 5, 108, 0, 0, 120, 121, 5, 115, 0, 0, 121, 127,
		5, 101, 0, 0, 122, 123, 5, 101, 0, 0, 123, 124, 5, 108, 0, 0, 124, 125,
		5, 115, 0, 0, 125, 127, 5, 101, 0, 0, 126, 114, 1, 0, 0, 0, 126, 118, 1,
		0, 0, 0, 126, 122, 1, 0, 0, 0, 127, 16, 1, 0, 0, 0, 128, 132, 5, 62, 0,
		0, 129, 130, 5, 71, 0, 0, 130, 132, 5, 84, 0, 0, 131, 128, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 132, 18, 1, 0, 0, 0, 133, 134, 5, 62, 0, 0, 134,
		139, 5, 61, 0, 0, 135, 136, 5, 71, 0, 0, 136, 139, 5, 69, 0, 0, 137, 139,
		5, 8805, 0, 0, 138, 133, 1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 137, 1,
		0, 0, 0, 139, 20, 1, 0, 0, 0, 140, 144, 5, 60, 0, 0, 141, 142, 5, 76, 0,
		0, 142, 144, 5, 84, 0, 0, 143, 140, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144,
		22, 1, 0, 0, 0, 145, 146, 5, 60, 0, 0, 146, 151, 5, 61, 0, 0, 147, 148,
		5, 76, 0, 0, 148, 151, 5, 69, 0, 0, 149, 151, 5, 8804, 0, 0, 150, 145,
		1, 0, 0, 0, 150, 147, 1, 0, 0, 0, 150, 149, 1, 0, 0, 0, 151, 24, 1, 0,
		0, 0, 152, 153, 5, 61, 0, 0, 153, 157, 5, 61, 0, 0, 154, 155, 5, 69, 0,
		0, 155, 157, 5, 81, 0, 0, 156, 152, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157,
		26, 1, 0, 0, 0, 158, 159, 5, 33, 0, 0, 159, 166, 5, 61, 0, 0, 160, 161,
		5, 60, 0, 0, 161, 166, 5, 62, 0, 0, 162, 163, 5, 78, 0, 0, 163, 164, 5,
		69, 0, 0, 164, 166, 5, 81, 0, 0, 165, 158, 1, 0, 0, 0, 165, 160, 1, 0,
		0, 0, 165, 162, 1, 0, 0, 0, 166, 28, 1, 0, 0, 0, 167, 168, 5, 40, 0, 0,
		168, 30, 1, 0, 0, 0, 169, 170, 5, 41, 0, 0, 170, 32, 1, 0, 0, 0, 171, 173,
		5, 45, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0,
		0, 0, 174, 176, 7, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 185, 1, 0, 0, 0, 179,
		181, 5, 46, 0, 0, 180, 182, 7, 0, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0,
		0, 0, 185, 179, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 34, 1, 0, 0, 0,
		187, 188, 5, 83, 0, 0, 188, 189, 5, 81, 0, 0, 189, 190, 5, 82, 0, 0, 190,
		191, 5, 84, 0, 0, 191, 36, 1, 0, 0, 0, 192, 193, 5, 42, 0, 0, 193, 38,
		1, 0, 0, 0, 194, 195, 5, 47, 0, 0, 195, 40, 1, 0, 0, 0, 196, 197, 5, 43,
		0, 0, 197, 42, 1, 0, 0, 0, 198, 199, 5, 45, 0, 0, 199, 44, 1, 0, 0, 0,
		200, 202, 7, 1, 0, 0, 201, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206,
		6, 22, 0, 0, 206, 46, 1, 0, 0, 0, 19, 0, 51, 56, 62, 76, 93, 112, 126,
		131, 138, 143, 150, 156, 165, 172, 177, 183, 185, 203, 1, 6, 0, 0,
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

// SimpleBooleanLexerInit initializes any static state used to implement SimpleBooleanLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleBooleanLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleBooleanLexerInit() {
	staticData := &simplebooleanlexerLexerStaticData
	staticData.once.Do(simplebooleanlexerLexerInit)
}

// NewSimpleBooleanLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleBooleanLexer(input antlr.CharStream) *SimpleBooleanLexer {
	SimpleBooleanLexerInit()
	l := new(SimpleBooleanLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &simplebooleanlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SimpleBoolean.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleBooleanLexer tokens.
const (
	SimpleBooleanLexerAND     = 1
	SimpleBooleanLexerOR      = 2
	SimpleBooleanLexerNOT     = 3
	SimpleBooleanLexerTRUE    = 4
	SimpleBooleanLexerFALSE   = 5
	SimpleBooleanLexerIF      = 6
	SimpleBooleanLexerTHEN    = 7
	SimpleBooleanLexerELSE    = 8
	SimpleBooleanLexerGT      = 9
	SimpleBooleanLexerGE      = 10
	SimpleBooleanLexerLT      = 11
	SimpleBooleanLexerLE      = 12
	SimpleBooleanLexerEQ      = 13
	SimpleBooleanLexerNEQ     = 14
	SimpleBooleanLexerLPAREN  = 15
	SimpleBooleanLexerRPAREN  = 16
	SimpleBooleanLexerDECIMAL = 17
	SimpleBooleanLexerSQRT    = 18
	SimpleBooleanLexerMULT    = 19
	SimpleBooleanLexerDIV     = 20
	SimpleBooleanLexerADD     = 21
	SimpleBooleanLexerSUB     = 22
	SimpleBooleanLexerWS      = 23
)
