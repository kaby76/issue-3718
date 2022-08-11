// Generated from SimpleBoolean.g4 by ANTLR 4.10.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SimpleBooleanParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.10.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		AND=1, OR=2, NOT=3, TRUE=4, FALSE=5, IF=6, THEN=7, ELSE=8, GT=9, GE=10, 
		LT=11, LE=12, EQ=13, NEQ=14, LPAREN=15, RPAREN=16, DECIMAL=17, SQRT=18, 
		MULT=19, DIV=20, ADD=21, SUB=22, WS=23;
	public static final int
		RULE_parse = 0, RULE_expression = 1, RULE_boolean = 2;
	private static String[] makeRuleNames() {
		return new String[] {
			"parse", "expression", "boolean"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, "'IF'", null, null, null, null, null, 
			null, null, null, "'('", "')'", null, "'SQRT'", "'*'", "'/'", "'+'", 
			"'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "AND", "OR", "NOT", "TRUE", "FALSE", "IF", "THEN", "ELSE", "GT", 
			"GE", "LT", "LE", "EQ", "NEQ", "LPAREN", "RPAREN", "DECIMAL", "SQRT", 
			"MULT", "DIV", "ADD", "SUB", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SimpleBoolean.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SimpleBooleanParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ParseContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SimpleBooleanParser.EOF, 0); }
		public ParseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parse; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).enterParse(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).exitParse(this);
		}
	}

	public final ParseContext parse() throws RecognitionException {
		ParseContext _localctx = new ParseContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_parse);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(6);
			expression(0);
			setState(7);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(SimpleBooleanParser.LPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(SimpleBooleanParser.RPAREN, 0); }
		public TerminalNode NOT() { return getToken(SimpleBooleanParser.NOT, 0); }
		public TerminalNode SQRT() { return getToken(SimpleBooleanParser.SQRT, 0); }
		public BooleanContext boolean_() {
			return getRuleContext(BooleanContext.class,0);
		}
		public TerminalNode DECIMAL() { return getToken(SimpleBooleanParser.DECIMAL, 0); }
		public TerminalNode GT() { return getToken(SimpleBooleanParser.GT, 0); }
		public TerminalNode GE() { return getToken(SimpleBooleanParser.GE, 0); }
		public TerminalNode LT() { return getToken(SimpleBooleanParser.LT, 0); }
		public TerminalNode LE() { return getToken(SimpleBooleanParser.LE, 0); }
		public TerminalNode EQ() { return getToken(SimpleBooleanParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(SimpleBooleanParser.NEQ, 0); }
		public TerminalNode AND() { return getToken(SimpleBooleanParser.AND, 0); }
		public TerminalNode OR() { return getToken(SimpleBooleanParser.OR, 0); }
		public TerminalNode MULT() { return getToken(SimpleBooleanParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(SimpleBooleanParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(SimpleBooleanParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SimpleBooleanParser.SUB, 0); }
		public TerminalNode THEN() { return getToken(SimpleBooleanParser.THEN, 0); }
		public TerminalNode ELSE() { return getToken(SimpleBooleanParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SimpleBooleanParser.IF, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				{
				setState(10);
				match(LPAREN);
				setState(11);
				expression(0);
				setState(12);
				match(RPAREN);
				}
				break;
			case NOT:
				{
				setState(14);
				match(NOT);
				setState(15);
				expression(10);
				}
				break;
			case SQRT:
				{
				setState(16);
				match(SQRT);
				setState(17);
				match(LPAREN);
				setState(18);
				expression(0);
				setState(19);
				match(RPAREN);
				}
				break;
			case TRUE:
			case FALSE:
				{
				setState(21);
				boolean_();
				}
				break;
			case DECIMAL:
				{
				setState(22);
				match(DECIMAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(53);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(51);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(25);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(26);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << GE) | (1L << LT) | (1L << LE) | (1L << EQ) | (1L << NEQ))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(27);
						expression(10);
						}
						break;

					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(28);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(29);
						_la = _input.LA(1);
						if ( !(_la==AND || _la==OR) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(30);
						expression(9);
						}
						break;

					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(31);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(32);
						_la = _input.LA(1);
						if ( !(_la==MULT || _la==DIV) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(33);
						expression(5);
						}
						break;

					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(34);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(35);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(36);
						expression(4);
						}
						break;

					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(37);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(38);
						match(THEN);
						setState(39);
						expression(0);
						setState(42);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
						case 1:
							{
							setState(40);
							match(ELSE);
							setState(41);
							expression(0);
							}
							break;
						}
						}
						break;

					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(44);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(45);
						match(IF);
						setState(46);
						expression(0);
						setState(49);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
						case 1:
							{
							setState(47);
							match(ELSE);
							setState(48);
							expression(0);
							}
							break;
						}
						}
						break;
					}
					} 
				}
				setState(55);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class BooleanContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(SimpleBooleanParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(SimpleBooleanParser.FALSE, 0); }
		public BooleanContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolean; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).enterBoolean(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleBooleanListener ) ((SimpleBooleanListener)listener).exitBoolean(this);
		}
	}

	public final BooleanContext boolean_() throws RecognitionException {
		BooleanContext _localctx = new BooleanContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_boolean);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 9);

		case 1:
			return precpred(_ctx, 8);

		case 2:
			return precpred(_ctx, 4);

		case 3:
			return precpred(_ctx, 3);

		case 4:
			return precpred(_ctx, 7);

		case 5:
			return precpred(_ctx, 6);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0017;\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001\u0018\b\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001+\b\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0003\u00012\b\u0001\u0005\u00014\b\u0001"+
		"\n\u0001\f\u00017\t\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0000\u0001"+
		"\u0002\u0003\u0000\u0002\u0004\u0000\u0005\u0001\u0000\t\u000e\u0001\u0000"+
		"\u0001\u0002\u0001\u0000\u0013\u0014\u0001\u0000\u0015\u0016\u0001\u0000"+
		"\u0004\u0005C\u0000\u0006\u0001\u0000\u0000\u0000\u0002\u0017\u0001\u0000"+
		"\u0000\u0000\u00048\u0001\u0000\u0000\u0000\u0006\u0007\u0003\u0002\u0001"+
		"\u0000\u0007\b\u0005\u0000\u0000\u0001\b\u0001\u0001\u0000\u0000\u0000"+
		"\t\n\u0006\u0001\uffff\uffff\u0000\n\u000b\u0005\u000f\u0000\u0000\u000b"+
		"\f\u0003\u0002\u0001\u0000\f\r\u0005\u0010\u0000\u0000\r\u0018\u0001\u0000"+
		"\u0000\u0000\u000e\u000f\u0005\u0003\u0000\u0000\u000f\u0018\u0003\u0002"+
		"\u0001\n\u0010\u0011\u0005\u0012\u0000\u0000\u0011\u0012\u0005\u000f\u0000"+
		"\u0000\u0012\u0013\u0003\u0002\u0001\u0000\u0013\u0014\u0005\u0010\u0000"+
		"\u0000\u0014\u0018\u0001\u0000\u0000\u0000\u0015\u0018\u0003\u0004\u0002"+
		"\u0000\u0016\u0018\u0005\u0011\u0000\u0000\u0017\t\u0001\u0000\u0000\u0000"+
		"\u0017\u000e\u0001\u0000\u0000\u0000\u0017\u0010\u0001\u0000\u0000\u0000"+
		"\u0017\u0015\u0001\u0000\u0000\u0000\u0017\u0016\u0001\u0000\u0000\u0000"+
		"\u00185\u0001\u0000\u0000\u0000\u0019\u001a\n\t\u0000\u0000\u001a\u001b"+
		"\u0007\u0000\u0000\u0000\u001b4\u0003\u0002\u0001\n\u001c\u001d\n\b\u0000"+
		"\u0000\u001d\u001e\u0007\u0001\u0000\u0000\u001e4\u0003\u0002\u0001\t"+
		"\u001f \n\u0004\u0000\u0000 !\u0007\u0002\u0000\u0000!4\u0003\u0002\u0001"+
		"\u0005\"#\n\u0003\u0000\u0000#$\u0007\u0003\u0000\u0000$4\u0003\u0002"+
		"\u0001\u0004%&\n\u0007\u0000\u0000&\'\u0005\u0007\u0000\u0000\'*\u0003"+
		"\u0002\u0001\u0000()\u0005\b\u0000\u0000)+\u0003\u0002\u0001\u0000*(\u0001"+
		"\u0000\u0000\u0000*+\u0001\u0000\u0000\u0000+4\u0001\u0000\u0000\u0000"+
		",-\n\u0006\u0000\u0000-.\u0005\u0006\u0000\u0000.1\u0003\u0002\u0001\u0000"+
		"/0\u0005\b\u0000\u000002\u0003\u0002\u0001\u00001/\u0001\u0000\u0000\u0000"+
		"12\u0001\u0000\u0000\u000024\u0001\u0000\u0000\u00003\u0019\u0001\u0000"+
		"\u0000\u00003\u001c\u0001\u0000\u0000\u00003\u001f\u0001\u0000\u0000\u0000"+
		"3\"\u0001\u0000\u0000\u00003%\u0001\u0000\u0000\u00003,\u0001\u0000\u0000"+
		"\u000047\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u000056\u0001\u0000"+
		"\u0000\u00006\u0003\u0001\u0000\u0000\u000075\u0001\u0000\u0000\u0000"+
		"89\u0007\u0004\u0000\u00009\u0005\u0001\u0000\u0000\u0000\u0005\u0017"+
		"*135";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}