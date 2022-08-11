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

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleBooleanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleBooleanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseSimpleBooleanListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseSimpleBooleanListener) ExitBoolean(ctx *BooleanContext) {}
