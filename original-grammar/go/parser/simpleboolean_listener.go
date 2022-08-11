// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleBooleanListener is a complete listener for a parse tree produced by SimpleBooleanParser.
type SimpleBooleanListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterDecimalExpression is called when entering the decimalExpression production.
	EnterDecimalExpression(c *DecimalExpressionContext)

	// EnterBoolExpression is called when entering the boolExpression production.
	EnterBoolExpression(c *BoolExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterCalcExpression is called when entering the calcExpression production.
	EnterCalcExpression(c *CalcExpressionContext)

	// EnterSqrtExpression is called when entering the sqrtExpression production.
	EnterSqrtExpression(c *SqrtExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterDeprecatedConditionalExpression is called when entering the deprecatedConditionalExpression production.
	EnterDeprecatedConditionalExpression(c *DeprecatedConditionalExpressionContext)

	// EnterComparatorExpression is called when entering the comparatorExpression production.
	EnterComparatorExpression(c *ComparatorExpressionContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitDecimalExpression is called when exiting the decimalExpression production.
	ExitDecimalExpression(c *DecimalExpressionContext)

	// ExitBoolExpression is called when exiting the boolExpression production.
	ExitBoolExpression(c *BoolExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitCalcExpression is called when exiting the calcExpression production.
	ExitCalcExpression(c *CalcExpressionContext)

	// ExitSqrtExpression is called when exiting the sqrtExpression production.
	ExitSqrtExpression(c *SqrtExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitDeprecatedConditionalExpression is called when exiting the deprecatedConditionalExpression production.
	ExitDeprecatedConditionalExpression(c *DeprecatedConditionalExpressionContext)

	// ExitComparatorExpression is called when exiting the comparatorExpression production.
	ExitComparatorExpression(c *ComparatorExpressionContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)
}
