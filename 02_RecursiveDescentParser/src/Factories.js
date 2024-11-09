const DefaultFactory = {
    Program(body) {
        return {
            type: "Program",
            body,
        };
    },

    EmptyStatement() {
        return {
            type: "EmptyStatement"
        };
    },

    BlockStatement(body) {
        return {
            type: "BlockStatement",
            body,
        };
    },

    ExpressionStatement(expression) {
        return {
            type: "ExpressionStatement",
            expression,
        };
    },

    VariableStatement(declarations) {
        return {
            type: "VariableStatement",
            declarations
        };
    },

    VariableDeclaration(id, init) {
        return {
            type: "VariableDeclaration",
            id,
            init
        };
    },

    AssignmentExpression(operator, left, right) {
        return {
            type: "AssignmentExpression",
            operator,
            left,
            right
        };
    },

    Identifier(name) {
        return {
            type: "Identifier",
            name,
        };
    },

    BinaryExpression(operator, left, right) {
        return {
            type: "BinaryExpression",
            operator,
            left,
            right
        };
    },

    StringLiteral(value) {
        return {
            type: "StringLiteral",
            value,
        };
    },

    NumericLiteral(value) {
        return {
            type: "NumericLiteral",
            value,
        };
    },

    // New method for IfStatement
    IfStatement(test, consequent, alternate) {
        return {
            type: "IfStatement",
            test,
            consequent,
            alternate
        };
    }
};

const SExpressionFactory = {
    Program(body) {
        return ["begin", body];
    },

    EmptyStatement() {
        return ["empty"];
    },

    BlockStatement(body) {
        return ["begin", body];
    },

    ExpressionStatement(expression) {
        return expression;
    },

    VariableStatement(declarations) {
        return ["var", declarations];
    },

    VariableDeclaration(id, init) {
        return init ? ["assign", id, init] : ["var", id];
    },

    AssignmentExpression(operator, left, right) {
        return [operator, left, right];
    },

    Identifier(name) {
        return name;
    },

    BinaryExpression(operator, left, right) {
        return [operator, left, right];
    },

    StringLiteral(value) {
        return `"${value}"`;
    },

    NumericLiteral(value) {
        return value;
    },

    // New method for IfStatement
    IfStatement(test, consequent, alternate) {
        const ifExpression = ["if", test, consequent];
        return alternate ? [...ifExpression, alternate] : ifExpression;
    }
};

module.exports = { DefaultFactory, SExpressionFactory };
