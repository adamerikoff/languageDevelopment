package loxLanguage;

public class AstPrinter implements Expression.Visitor<String>{
	String print(Expression expr) {
		return expr.accept(this);
	}
	
	@Override
	public String visitBinaryExpr(Expression.Binary expr) {
		return parenthesize(expr.operator.lexeme, expr.left, expr.right);
	}
	
	@Override
	public String visitGroupingExpr(Expression.Grouping expr) {
		return parenthesize("group", expr.expression);
	}
	
	@Override
	public String visitLiteralExpr(Expression.Literal expr) {
		if (expr.value == null) {
			return "nil";
		}
		return expr.value.toString();
	}
	
	@Override
	public String visitUnaryExpr(Expression.Unary expr) {
		return parenthesize(expr.operator.lexeme, expr.right);
	}
	
	private String parenthesize(String name, Expression... exprs) {
		StringBuilder builder = new StringBuilder();
		builder.append("(").append(name);
		for (Expression expr : exprs) {
			builder.append(" ");
			builder.append(expr.accept(this));
		}
		builder.append(")");
		return builder.toString();
	}
}
