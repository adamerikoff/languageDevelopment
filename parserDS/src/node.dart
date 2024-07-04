enum NodeType { PROGRAM, NUMERICLITERAL }

class Node {
  final NodeType type;
  final Object value;

  Node(this.type, this.value);
}
