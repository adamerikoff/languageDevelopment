package loxLanguage;

import java.util.List;

public class Parser {
	private static class ParseError extends RuntimeException {}
	
	private final List<Token> tokens;
	private int current = 0;
	
	Parser(List<Token> tokens) {
		this.tokens= tokens;
	}
	
	Expression parse() {
		try {
			return expression();
		} catch (ParseError error) {
			return null;
		}
	}

	private Expression expression() {
		return equality();
	}
	
	private boolean match(TokenType... types) {
		for (TokenType type : types) {
			if (check(type)) {
				advance();
				return true;
			}
		}
		return false;
	}
	
	private boolean check(TokenType type) {
		if (isAtEnd()) {
			return false;
		}
		return peek().type == type;
	}
	
	private Token advance() {
		if (!isAtEnd()) {
			current++;
		}
		return previous();
	}
	
	private boolean isAtEnd() {
		return peek().type == TokenType.EOF;
	}
	
	private Token peek() {
		return tokens.get(current);
	}
	
	private Token previous() {
		return tokens.get(current - 1);
	}
	
	private Expression equality() {
		Expression expr = comparison();
		while (match(TokenType.BANG_EQUAL, TokenType.EQUAL_EQUAL)) {
			Token operator = previous();
			Expression right = comparison();
			expr = new Expression.Binary(expr, right, operator);
		}
		return expr;
	}
	
	private Expression comparison() {
		Expression expr = term();
		while (match(TokenType.GREATER, TokenType.GREATER_EQUAL, TokenType.LESS, TokenType.LESS_EQUAL)) {
			Token operator = previous();
			Expression right = term();
			expr = new Expression.Binary(expr, right, operator);
		}
		return expr;
	}
	
	private Expression term() {
		Expression expr = factor();
		while (match(TokenType.MINUS, TokenType.PLUS)) {
			Token operator = previous();
			Expression right = factor();
			expr = new Expression.Binary(expr, right, operator);
		}
		return expr;
	}
	
	private Expression factor() {
		Expression expr = unary();
		while (match(TokenType.SLASH, TokenType.STAR)) {
			Token operator = previous();
			Expression right = unary();
			expr = new Expression.Binary(expr, right, operator);
		}
		return expr;
	}
	
	private Expression unary() {
		while (match(TokenType.BANG, TokenType.MINUS)) {
			Token operator = previous();
			Expression right = unary();
			return new Expression.Unary(operator, right);
		}
		return primary();
	}
	
	private Expression primary() {
		if (match(TokenType.FALSE)) return new Expression.Literal(false);
		if (match(TokenType.TRUE)) return new Expression.Literal(true);
		if (match(TokenType.NIL)) return new Expression.Literal(null);
		
		if (match(TokenType.NUMBER, TokenType.STRING)) {
			return new Expression.Literal(previous().literal);
		}
		
		if (match(TokenType.LEFT_PARENTHESIS)) {
			Expression expr = unary();
			consume(TokenType.RIGHT_PARENTHESIS, "Expect ')' after expression.");
			return new Expression.Grouping(expr);
		}
		
		throw error(peek(), "Expect expression.");
	}
	
	private Token consume(TokenType type, String message) {
		if (check(type)) return advance();
		throw error(peek(), message);
	}
	
	private ParseError error(Token token, String message) {
		Lox.error(token, message);
		return new ParseError();
	}
	
	private void synchronize() {
		advance();
		
		while(!isAtEnd()) {
			if (previous().type == TokenType.SEMICOLON) return;
			
			switch (peek().type) {
				case TokenType.CLASS:
				case TokenType.FUN:
				case TokenType.VAR:
				case TokenType.FOR:
				case TokenType.IF:
				case TokenType.WHILE:
				case TokenType.PRINT:
				case TokenType.RETURN:
					return;
			}
			advance();
		}
	}
}

