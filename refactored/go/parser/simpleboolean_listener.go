// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleBooleanListener is a complete listener for a parse tree produced by SimpleBooleanParser.
type SimpleBooleanListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)
}
