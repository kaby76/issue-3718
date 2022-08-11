// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleBooleanListener is a complete listener for a parse tree produced by SimpleBooleanParser.
type BaseSimpleBooleanListener struct{}

var _ SimpleBooleanListener = &BaseSimpleBooleanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleBooleanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleBooleanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleBooleanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleBooleanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseSimpleBooleanListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseSimpleBooleanListener) ExitParse(ctx *ParseContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseSimpleBooleanListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseSimpleBooleanListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterDecimalExpression is called when production decimalExpression is entered.
func (s *BaseSimpleBooleanListener) EnterDecimalExpression(ctx *DecimalExpressionContext) {}

// ExitDecimalExpression is called when production decimalExpression is exited.
func (s *BaseSimpleBooleanListener) ExitDecimalExpression(ctx *DecimalExpressionContext) {}

// EnterBoolExpression is called when production boolExpression is entered.
func (s *BaseSimpleBooleanListener) EnterBoolExpression(ctx *BoolExpressionContext) {}

// ExitBoolExpression is called when production boolExpression is exited.
func (s *BaseSimpleBooleanListener) ExitBoolExpression(ctx *BoolExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseSimpleBooleanListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseSimpleBooleanListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterCalcExpression is called when production calcExpression is entered.
func (s *BaseSimpleBooleanListener) EnterCalcExpression(ctx *CalcExpressionContext) {}

// ExitCalcExpression is called when production calcExpression is exited.
func (s *BaseSimpleBooleanListener) ExitCalcExpression(ctx *CalcExpressionContext) {}

// EnterSqrtExpression is called when production sqrtExpression is entered.
func (s *BaseSimpleBooleanListener) EnterSqrtExpression(ctx *SqrtExpressionContext) {}

// ExitSqrtExpression is called when production sqrtExpression is exited.
func (s *BaseSimpleBooleanListener) ExitSqrtExpression(ctx *SqrtExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseSimpleBooleanListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseSimpleBooleanListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseSimpleBooleanListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseSimpleBooleanListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterDeprecatedConditionalExpression is called when production deprecatedConditionalExpression is entered.
func (s *BaseSimpleBooleanListener) EnterDeprecatedConditionalExpression(ctx *DeprecatedConditionalExpressionContext) {
}

// ExitDeprecatedConditionalExpression is called when production deprecatedConditionalExpression is exited.
func (s *BaseSimpleBooleanListener) ExitDeprecatedConditionalExpression(ctx *DeprecatedConditionalExpressionContext) {
}

// EnterComparatorExpression is called when production comparatorExpression is entered.
func (s *BaseSimpleBooleanListener) EnterComparatorExpression(ctx *ComparatorExpressionContext) {}

// ExitComparatorExpression is called when production comparatorExpression is exited.
func (s *BaseSimpleBooleanListener) ExitComparatorExpression(ctx *ComparatorExpressionContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseSimpleBooleanListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseSimpleBooleanListener) ExitComparator(ctx *ComparatorContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseSimpleBooleanListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseSimpleBooleanListener) ExitBinary(ctx *BinaryContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseSimpleBooleanListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseSimpleBooleanListener) ExitBoolean(ctx *BooleanContext) {}
