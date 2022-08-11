// Generated from SimpleBoolean.g4 by ANTLR 4.10.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SimpleBooleanParser}.
 */
public interface SimpleBooleanListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SimpleBooleanParser#parse}.
	 * @param ctx the parse tree
	 */
	void enterParse(SimpleBooleanParser.ParseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleBooleanParser#parse}.
	 * @param ctx the parse tree
	 */
	void exitParse(SimpleBooleanParser.ParseContext ctx);

	/**
	 * Enter a parse tree produced by {@link SimpleBooleanParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(SimpleBooleanParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleBooleanParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(SimpleBooleanParser.ExpressionContext ctx);

	/**
	 * Enter a parse tree produced by {@link SimpleBooleanParser#boolean}.
	 * @param ctx the parse tree
	 */
	void enterBoolean(SimpleBooleanParser.BooleanContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleBooleanParser#boolean}.
	 * @param ctx the parse tree
	 */
	void exitBoolean(SimpleBooleanParser.BooleanContext ctx);
}