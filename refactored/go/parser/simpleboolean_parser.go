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
		"parse", "expression", "boolean",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 59, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 24, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 43, 8, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 5, 1, 52, 8, 1, 10, 1, 12, 1,
		55, 9, 1, 1, 2, 1, 2, 1, 2, 0, 1, 2, 3, 0, 2, 4, 0, 5, 1, 0, 9, 14, 1,
		0, 1, 2, 1, 0, 19, 20, 1, 0, 21, 22, 1, 0, 4, 5, 67, 0, 6, 1, 0, 0, 0,
		2, 23, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 5, 0, 0,
		1, 8, 1, 1, 0, 0, 0, 9, 10, 6, 1, -1, 0, 10, 11, 5, 15, 0, 0, 11, 12, 3,
		2, 1, 0, 12, 13, 5, 16, 0, 0, 13, 24, 1, 0, 0, 0, 14, 15, 5, 3, 0, 0, 15,
		24, 3, 2, 1, 10, 16, 17, 5, 18, 0, 0, 17, 18, 5, 15, 0, 0, 18, 19, 3, 2,
		1, 0, 19, 20, 5, 16, 0, 0, 20, 24, 1, 0, 0, 0, 21, 24, 3, 4, 2, 0, 22,
		24, 5, 17, 0, 0, 23, 9, 1, 0, 0, 0, 23, 14, 1, 0, 0, 0, 23, 16, 1, 0, 0,
		0, 23, 21, 1, 0, 0, 0, 23, 22, 1, 0, 0, 0, 24, 53, 1, 0, 0, 0, 25, 26,
		10, 9, 0, 0, 26, 27, 7, 0, 0, 0, 27, 52, 3, 2, 1, 10, 28, 29, 10, 8, 0,
		0, 29, 30, 7, 1, 0, 0, 30, 52, 3, 2, 1, 9, 31, 32, 10, 4, 0, 0, 32, 33,
		7, 2, 0, 0, 33, 52, 3, 2, 1, 5, 34, 35, 10, 3, 0, 0, 35, 36, 7, 3, 0, 0,
		36, 52, 3, 2, 1, 4, 37, 38, 10, 7, 0, 0, 38, 39, 5, 7, 0, 0, 39, 42, 3,
		2, 1, 0, 40, 41, 5, 8, 0, 0, 41, 43, 3, 2, 1, 0, 42, 40, 1, 0, 0, 0, 42,
		43, 1, 0, 0, 0, 43, 52, 1, 0, 0, 0, 44, 45, 10, 6, 0, 0, 45, 46, 5, 6,
		0, 0, 46, 49, 3, 2, 1, 0, 47, 48, 5, 8, 0, 0, 48, 50, 3, 2, 1, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 25, 1, 0, 0, 0,
		51, 28, 1, 0, 0, 0, 51, 31, 1, 0, 0, 0, 51, 34, 1, 0, 0, 0, 51, 37, 1,
		0, 0, 0, 51, 44, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53,
		54, 1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 7, 4, 0,
		0, 57, 5, 1, 0, 0, 0, 5, 23, 42, 49, 51, 53,
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
	SimpleBooleanParserRULE_boolean    = 2
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
		p.SetState(6)
		p.expression(0)
	}
	{
		p.SetState(7)
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

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLPAREN, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserRPAREN, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserNOT, 0)
}

func (s *ExpressionContext) SQRT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserSQRT, 0)
}

func (s *ExpressionContext) Boolean() IBooleanContext {
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

func (s *ExpressionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserDECIMAL, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserGT, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserGE, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLT, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserLE, 0)
}

func (s *ExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserEQ, 0)
}

func (s *ExpressionContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserNEQ, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserOR, 0)
}

func (s *ExpressionContext) MULT() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserMULT, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserDIV, 0)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserSUB, 0)
}

func (s *ExpressionContext) THEN() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserTHEN, 0)
}

func (s *ExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserELSE, 0)
}

func (s *ExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(SimpleBooleanParserIF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleBooleanListener); ok {
		listenerT.ExitExpression(s)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleBooleanParserLPAREN:
		{
			p.SetState(10)
			p.Match(SimpleBooleanParserLPAREN)
		}
		{
			p.SetState(11)
			p.expression(0)
		}
		{
			p.SetState(12)
			p.Match(SimpleBooleanParserRPAREN)
		}

	case SimpleBooleanParserNOT:
		{
			p.SetState(14)
			p.Match(SimpleBooleanParserNOT)
		}
		{
			p.SetState(15)
			p.expression(10)
		}

	case SimpleBooleanParserSQRT:
		{
			p.SetState(16)
			p.Match(SimpleBooleanParserSQRT)
		}
		{
			p.SetState(17)
			p.Match(SimpleBooleanParserLPAREN)
		}
		{
			p.SetState(18)
			p.expression(0)
		}
		{
			p.SetState(19)
			p.Match(SimpleBooleanParserRPAREN)
		}

	case SimpleBooleanParserTRUE, SimpleBooleanParserFALSE:
		{
			p.SetState(21)
			p.Boolean()
		}

	case SimpleBooleanParserDECIMAL:
		{
			p.SetState(22)
			p.Match(SimpleBooleanParserDECIMAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(26)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleBooleanParserGT)|(1<<SimpleBooleanParserGE)|(1<<SimpleBooleanParserLT)|(1<<SimpleBooleanParserLE)|(1<<SimpleBooleanParserEQ)|(1<<SimpleBooleanParserNEQ))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(27)
					p.expression(10)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(29)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleBooleanParserAND || _la == SimpleBooleanParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(30)
					p.expression(9)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(32)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleBooleanParserMULT || _la == SimpleBooleanParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(33)
					p.expression(5)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(35)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleBooleanParserADD || _la == SimpleBooleanParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expression(4)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(38)
					p.Match(SimpleBooleanParserTHEN)
				}
				{
					p.SetState(39)
					p.expression(0)
				}
				p.SetState(42)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(40)
						p.Match(SimpleBooleanParserELSE)
					}
					{
						p.SetState(41)
						p.expression(0)
					}

				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleBooleanParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(45)
					p.Match(SimpleBooleanParserIF)
				}
				{
					p.SetState(46)
					p.expression(0)
				}
				p.SetState(49)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(47)
						p.Match(SimpleBooleanParserELSE)
					}
					{
						p.SetState(48)
						p.expression(0)
					}

				}

			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 4, SimpleBooleanParserRULE_boolean)
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
		p.SetState(56)
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
