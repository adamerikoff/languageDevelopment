package loxLanguage;

public abstract class Expression {
	interface Visitor<R> {
        R visitBinaryExpr(Binary expr);
        R visitGroupingExpr(Grouping expr);
        R visitLiteralExpr(Literal expr);
        R visitUnaryExpr(Unary expr);
    }
	
    abstract <R> R accept(Visitor<R> visitor);

	static class Binary extends Expression {
		final Expression left;
	    final Token operator;
	    final Expression right;
	    
	    Binary(Expression left, Expression right, Token operator) {
			this.left = left;
			this.operator = operator;
			this.right = right;
	    }
	    
	    @Override
        <R> R accept(Visitor<R> visitor) {
            return visitor.visitBinaryExpr(this);
        }
	}

	static class Grouping extends Expression {
		final Expression expression;
	
		Grouping(Expression expression) {
			this.expression = expression;
		}
		
		@Override
        <R> R accept(Visitor<R> visitor) {
            return visitor.visitGroupingExpr(this);
        }
	}

	static class Literal extends Expression {
		final Object value;
		
		Literal(Object value) {
			this.value = value;
		}

		@Override
        <R> R accept(Visitor<R> visitor) {
            return visitor.visitLiteralExpr(this);
        }
		
	}

	static class Unary extends Expression {
		final Token operator;
		final Expression right;
		
		Unary(Token operator, Expression right) {
			this.operator = operator;
			this.right = right;
		}
		
		@Override
        <R> R accept(Visitor<R> visitor) {
            return visitor.visitUnaryExpr(this);
        }
	}
}
