package loxLanguage;

public class Interpreter implements Expression.Visitor<Object> {
	
	void interpret(Expression expression) {
		try {
			Object value = evaluate(expression);
			System.out.println(stringify(value));
		} catch (RunTimeError error) {
			Lox.runtimeError(error);
		}
	}
	
	@Override
	public Object visitLiteralExpr(Expression.Literal expr) {
		return expr.value;
	}
	
	@Override
	public Object visitUnaryExpr(Expression.Unary expr) {
		Object right = evaluate(expr.right);
		switch (expr.operator.type) {
			case TokenType.MINUS:
				checkNumberOperand(expr.operator, right);
				return -(double)right;
			case TokenType.BANG:
				return !isTruthy(right);
		}
		// Unreachable.
		return null;
	}
	
	@Override
	public Object visitBinaryExpr(Expression.Binary expr) {
		Object left = evaluate(expr.left);
		Object right = evaluate(expr.right);
		switch (expr.operator.type) {
			case TokenType.GREATER:
				checkNumberOperands(expr.operator, left, right);
				return (double)left > (double)right;
			case TokenType.GREATER_EQUAL:
				checkNumberOperands(expr.operator, left, right);
				return (double)left >= (double)right;
			case TokenType.LESS:
				checkNumberOperands(expr.operator, left, right);
				return (double)left < (double)right;
			case TokenType.LESS_EQUAL:
				checkNumberOperands(expr.operator, left, right);
				return (double)left <= (double)right;
			case TokenType.MINUS:
				checkNumberOperands(expr.operator, left, right);
				return (double)left - (double)right;
			case TokenType.PLUS:
				if (left instanceof Double && right instanceof Double) {
					return (double)left + (double)right;
				}
				if (left instanceof String && right instanceof String) {
					return (String)left + (String)right;
				}
				throw new RunTimeError(expr.operator, "Operands must be two numbers or two strings.");
			case TokenType.SLASH:
				checkNumberOperands(expr.operator, left, right);
				return (double)left / (double)right;
			case TokenType.STAR:
				checkNumberOperands(expr.operator, left, right);
				return (double)left * (double)right;
			case TokenType.BANG_EQUAL: return !isEqual(left, right);
			case TokenType.EQUAL_EQUAL: return isEqual(left, right); 
		}
		// Unreachable.
		return null;
	}
	
	@Override
	public Object visitGroupingExpr(Expression.Grouping expr) {
		return evaluate(expr.expression);
	}
	
	private Object evaluate(Expression expr) {
		return expr.accept(this);
	}
	
	private boolean isTruthy(Object object) {
		if (object == null) return false;
		if (object instanceof Boolean) return (boolean)object;
		return true;
	}
	
	private boolean isEqual(Object a, Object b) {
		if (a == null && b == null) return true;
		if (a == null) return false;
		return a.equals(b);
	}
	
	private void checkNumberOperand(Token operator, Object operand) {
		if (operand instanceof Double) return;
		throw new RunTimeError(operator, "Operand must be a number.");
	}
	
	private void checkNumberOperands(Token operator, Object left, Object right) {
		if (left instanceof Double && right instanceof Double) return;
		throw new RunTimeError(operator, "Operands must be numbers");
	}
	
	private String stringify(Object object) {
		if (object == null) return "nil";
		if (object instanceof Double) {
			String text = object.toString();
			if (text.endsWith(".0")) {
				text = text.substring(0, text.length() - 2);
			}
			return text;
		}
		return object.toString();
	}
}
