import 'tokenizer.dart';
import 'node.dart';

class Parser {
  String string = '';
  Tokenizer tokenizer = Tokenizer();

  Node parse(expression) {
    this.string = expression;

    return this.program();
  }

  Node program() {
    return this.numericLiteral();
  }

  Node numericLiteral() {
    NodeType type = NodeType.NUMERICLITERAL;
    num value = int.parse(string);
    return Node(type, value);
  }
}
