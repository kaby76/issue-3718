// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleBooleanParser struct {
	*antlr.BaseParser
}

var simplebooleanParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplebooleanParserInit() {
	staticData := &simplebooleanParserStaticData
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
		"parse", "expression", "comparator", "binary", "boolean",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 69, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 49, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 56,
		8, 1, 5, 1, 58, 8, 1, 10, 1, 12, 1, 61, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 0, 1, 2, 5, 0, 2, 4, 6, 8, 0, 5, 1, 0, 19, 20, 1, 0, 21,
		22, 1, 0, 9, 14, 1, 0, 1, 2, 1, 0, 4, 5, 75, 0, 10, 1, 0, 0, 0, 2, 27,
		1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0, 10,
		11, 3, 2, 1, 0, 11, 12, 5, 0, 0, 1, 12, 1, 1, 0, 0, 0, 13, 14, 6, 1, -1,
		0, 14, 15, 5, 15, 0, 0, 15, 16, 3, 2, 1, 0, 16, 17, 5, 16, 0, 0, 17, 28,
		1, 0, 0, 0, 18, 19, 5, 3, 0, 0, 19, 28, 3, 2, 1, 10, 20, 21, 5, 18, 0,
		0, 21, 22, 5, 15, 0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 16, 0, 0, 24, 28,
		1, 0, 0, 0, 25, 28, 3, 8, 4, 0, 26, 28, 5, 17, 0, 0, 27, 13, 1, 0, 0, 0,
		27, 18, 1, 0, 0, 0, 27, 20, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 26, 1,
		0, 0, 0, 28, 59, 1, 0, 0, 0, 29, 30, 10, 9, 0, 0, 30, 31, 3, 4, 2, 0, 31,
		32, 3, 2, 1, 10, 32, 58, 1, 0, 0, 0, 33, 34, 10, 8, 0, 0, 34, 35, 3, 6,
		3, 0, 35, 36, 3, 2, 1, 9, 36, 58, 1, 0, 0, 0, 37, 38, 10, 4, 0, 0, 38,
		39, 7, 0, 0, 0, 39, 58, 3, 2, 1, 5, 40, 41, 10, 3, 0, 0, 41, 42, 7, 1,
		0, 0, 42, 58, 3, 2, 1, 4, 43, 44, 10, 7, 0, 0, 44, 45, 5, 7, 0, 0, 45,
		48, 3, 2, 1, 0, 46, 47, 5, 8, 0, 0, 47, 49, 3, 2, 1, 0, 48, 46, 1, 0, 0,
		0, 48, 49, 1, 0, 0, 0, 49, 58, 1, 0, 0, 0, 50, 51, 10, 6, 0, 0, 51, 52,
		5, 6, 0, 0, 52, 55, 3, 2, 1, 0, 53, 54, 5, 8, 0, 0, 54, 56, 3, 2, 1, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 29, 1,
		0, 0, 0, 57, 33, 1, 0, 0, 0, 57, 37, 1, 0, 0, 0, 57, 40, 1, 0, 0, 0, 57,
		43, 1, 0, 0, 0, 57, 50, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0,
		0, 59, 60, 1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 7,
		2, 0, 0, 63, 5, 1, 0, 0, 0, 64, 65, 7, 3, 0, 0, 65, 7, 1, 0, 0, 0, 66,
		67, 7, 4, 0, 0, 67, 9, 1, 0, 0, 0, 5, 27, 48, 55, 57, 59,
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

// SimpleBooleanParserInit initializes any static state used to implement SimpleBooleanParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleBooleanParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleBooleanParserInit() {
	staticData := &simplebooleanParserStaticData
	staticData.once.Do(simplebooleanParserInit)
}

// NewSimpleBooleanParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleBooleanParser(input antlr.TokenStream) *SimpleBooleanParser {
	SimpleBooleanParserInit()
	this := new(SimpleBooleanParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &simplebooleanParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SimpleBoolean.g4"

	return this
}

// SimpleBooleanParser tokens.
const (
	SimpleBooleanParserEOF     = antlr.TokenEOF
	SimpleBooleanParserAND     = 1
	SimpleBooleanParserOR      = 2
	SimpleBooleanParserNOT     = 3
	SimpleBooleanParserTRUE    = 4
	SimpleBooleanParserFALSE   = 5
	SimpleBooleanParserIF      = 6
	SimpleBooleanParserTHEN    = 7
	SimpleBooleanParserELSE    = 8
	SimpleBooleanParserGT      = 9
	SimpleBooleanParserGE      = 10
	SimpleBooleanParserLT      = 11
	SimpleBooleanParserLE      = 12
	SimpleBooleanParserEQ      = 13
	SimpleBooleanParserNEQ     = 14
	SimpleBooleanParserLPAREN  = 15
	SimpleBooleanParserRPAREN  = 16
	SimpleBooleanParserDECIMAL = 17
	SimpleBooleanParserSQRT    = 18
	SimpleBooleanParserMULT    = 19
	SimpleBooleanParserDIV     = 20
	SimpleBooleanParserADD     = 21
	SimpleBooleanParserSUB     = 22
	SimpleBooleanParserWS      = 23
)

// SimpleBooleanParser rules.
const (
	SimpleBooleanParserRULE_parse      = 0
	SimpleBooleanParserRULE_expression = 1
	SimpleBooleanParserRULE_comparator = 2
	SimpleBooleanParserRULE_binary     = 3
	SimpleBooleanParserRULE_boolean    = 4
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleBooleanParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleBooleanParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *SimpleBooleanParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleBooleanParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.expression(0)
	}
	{
		p.SetState(11)
		p.Match(SimpleBooleanParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleBooleanParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleBooleanParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    IBinaryContext
	right IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryExpressionContext) GetOp() IBinaryContext { return s.op }

func (s *BinaryExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryExpressionContext) SetOp(v IBinaryContext) { s.op = v }

func (s *BinaryExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) Binary() IBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

type DecimalExpressionContext struct {
	*ExpressionContext
}

func NewDecimalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalExpressionContext {
	var p = new(DecimalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DecimalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalExpressionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserDECIMAL, 0)
}

func (s *DecimalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterDecimalExpression(s)
	}
}

func (s *DecimalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitDecimalExpression(s)
	}
}

type BoolExpressionContext struct {
	*ExpressionContext
}

func NewBoolExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) Boolean() IBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanContext)
}

func (s *BoolExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterBoolExpression(s)
	}
}

func (s *BoolExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitBoolExpression(s)
	}
}

type ConditionalExpressionContext struct {
	*ExpressionContext
	affirmative IExpressionContext
	condition   IExpressionContext
	negative    IExpressionContext
}

func NewConditionalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalExpressionContext) GetAffirmative() IExpressionContext { return s.affirmative }

func (s *ConditionalExpressionContext) GetCondition() IExpressionContext { return s.condition }

func (s *ConditionalExpressionContext) GetNegative() IExpressionContext { return s.negative }

func (s *ConditionalExpressionContext) SetAffirmative(v IExpressionContext) { s.affirmative = v }

func (s *ConditionalExpressionContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *ConditionalExpressionContext) SetNegative(v IExpressionContext) { s.negative = v }

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserIF, 0)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserELSE, 0)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

type CalcExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewCalcExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalcExpressionContext {
	var p = new(CalcExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CalcExpressionContext) GetOp() antlr.Token { return s.op }

func (s *CalcExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *CalcExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *CalcExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *CalcExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *CalcExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *CalcExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CalcExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CalcExpressionContext) MULT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserMULT, 0)
}

func (s *CalcExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserDIV, 0)
}

func (s *CalcExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserADD, 0)
}

func (s *CalcExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserSUB, 0)
}

func (s *CalcExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterCalcExpression(s)
	}
}

func (s *CalcExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitCalcExpression(s)
	}
}

type SqrtExpressionContext struct {
	*ExpressionContext
}

func NewSqrtExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrtExpressionContext {
	var p = new(SqrtExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SqrtExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtExpressionContext) SQRT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserSQRT, 0)
}

func (s *SqrtExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLPAREN, 0)
}

func (s *SqrtExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqrtExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserRPAREN, 0)
}

func (s *SqrtExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterSqrtExpression(s)
	}
}

func (s *SqrtExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitSqrtExpression(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type ParenExpressionContext struct {
	*ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLPAREN, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserRPAREN, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

type DeprecatedConditionalExpressionContext struct {
	*ExpressionContext
	condition   IExpressionContext
	affirmative IExpressionContext
	negative    IExpressionContext
}

func NewDeprecatedConditionalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeprecatedConditionalExpressionContext {
	var p = new(DeprecatedConditionalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DeprecatedConditionalExpressionContext) GetCondition() IExpressionContext {
	return s.condition
}

func (s *DeprecatedConditionalExpressionContext) GetAffirmative() IExpressionContext {
	return s.affirmative
}

func (s *DeprecatedConditionalExpressionContext) GetNegative() IExpressionContext { return s.negative }

func (s *DeprecatedConditionalExpressionContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *DeprecatedConditionalExpressionContext) SetAffirmative(v IExpressionContext) {
	s.affirmative = v
}

func (s *DeprecatedConditionalExpressionContext) SetNegative(v IExpressionContext) { s.negative = v }

func (s *DeprecatedConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeprecatedConditionalExpressionContext) THEN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserTHEN, 0)
}

func (s *DeprecatedConditionalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DeprecatedConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *DeprecatedConditionalExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserELSE, 0)
}

func (s *DeprecatedConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterDeprecatedConditionalExpression(s)
	}
}

func (s *DeprecatedConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitDeprecatedConditionalExpression(s)
	}
}

type ComparatorExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    IComparatorContext
	right IExpressionContext
}

func NewComparatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorExpressionContext {
	var p = new(ComparatorExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparatorExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ComparatorExpressionContext) GetOp() IComparatorContext { return s.op }

func (s *ComparatorExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ComparatorExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ComparatorExpressionContext) SetOp(v IComparatorContext) { s.op = v }

func (s *ComparatorExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ComparatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparatorExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ComparatorExpressionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ComparatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterComparatorExpression(s)
	}
}

func (s *ComparatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitComparatorExpression(s)
	}
}

func (p *SimpleBooleanParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SimpleBooleanParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, SimpleBooleanParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleBooleanParserLPAREN:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Match(SimpleBooleanParserLPAREN)
		}
		{
			p.SetState(15)
			p.expression(0)
		}
		{
			p.SetState(16)
			p.Match(SimpleBooleanParserRPAREN)
		}

	case SimpleBooleanParserNOT:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(SimpleBooleanParserNOT)
		}
		{
			p.SetState(19)
			p.expression(10)
		}

	case SimpleBooleanParserSQRT:
		localctx = NewSqrtExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(SimpleBooleanParserSQRT)
		}
		{
			p.SetState(21)
			p.Match(SimpleBooleanParserLPAREN)
		}
		{
			p.SetState(22)
			p.expression(0)
		}
		{
			p.SetState(23)
			p.Match(SimpleBooleanParserRPAREN)
		}

	case SimpleBooleanParserTRUE, SimpleBooleanParserFALSE:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Boolean()
		}

	case SimpleBooleanParserDECIMAL:
		localctx = NewDecimalExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(SimpleBooleanParserDECIMAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ComparatorExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(30)

					var _x = p.Comparator()

					localctx.(*ComparatorExpressionContext).op = _x
				}
				{
					p.SetState(31)

					var _x = p.expression(10)

					localctx.(*ComparatorExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(34)

					var _x = p.Binary()

					localctx.(*BinaryExpressionContext).op = _x
				}
				{
					p.SetState(35)

					var _x = p.expression(9)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 3:
				localctx = NewCalcExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CalcExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(38)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CalcExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleBooleanParserMULT || _la == SimpleBooleanParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CalcExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)

					var _x = p.expression(5)

					localctx.(*CalcExpressionContext).right = _x
				}

			case 4:
				localctx = NewCalcExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CalcExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(41)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CalcExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleBooleanParserADD || _la == SimpleBooleanParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CalcExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(42)

					var _x = p.expression(4)

					localctx.(*CalcExpressionContext).right = _x
				}

			case 5:
				localctx = NewDeprecatedConditionalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*DeprecatedConditionalExpressionContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(44)
					p.Match(SimpleBooleanParserTHEN)
				}
				{
					p.SetState(45)

					var _x = p.expression(0)

					localctx.(*DeprecatedConditionalExpressionContext).affirmative = _x
				}
				p.SetState(48)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(46)
						p.Match(SimpleBooleanParserELSE)
					}
					{
						p.SetState(47)

						var _x = p.expression(0)

						localctx.(*DeprecatedConditionalExpressionContext).negative = _x
					}

				}

			case 6:
				localctx = NewConditionalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ConditionalExpressionContext).affirmative = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(51)
					p.Match(SimpleBooleanParserIF)
				}
				{
					p.SetState(52)

					var _x = p.expression(0)

					localctx.(*ConditionalExpressionContext).condition = _x
				}
				p.SetState(55)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(53)
						p.Match(SimpleBooleanParserELSE)
					}
					{
						p.SetState(54)

						var _x = p.expression(0)

						localctx.(*ConditionalExpressionContext).negative = _x
					}

				}

			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleBooleanParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleBooleanParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserGT, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserGE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLE, 0)
}

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserEQ, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserNEQ, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *SimpleBooleanParser) Comparator() (localctx IComparatorContext) {
	this := p
	_ = this

	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleBooleanParserRULE_comparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleBooleanParserGT)|(1<<SimpleBooleanParserGE)|(1<<SimpleBooleanParserLT)|(1<<SimpleBooleanParserLE)|(1<<SimpleBooleanParserEQ)|(1<<SimpleBooleanParserNEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleBooleanParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleBooleanParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserOR, 0)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *SimpleBooleanParser) Binary() (localctx IBinaryContext) {
	this := p
	_ = this

	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleBooleanParserRULE_binary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleBooleanParserAND || _la == SimpleBooleanParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleBooleanParserRULE_boolean
	return p
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleBooleanParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserFALSE, 0)
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (p *SimpleBooleanParser) Boolean() (localctx IBooleanContext) {
	this := p
	_ = this

	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleBooleanParserRULE_boolean)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleBooleanParserTRUE || _la == SimpleBooleanParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SimpleBooleanParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleBooleanParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
