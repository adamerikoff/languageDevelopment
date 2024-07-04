import '../src/parser.dart';

import 'dart:developer';

void main() {
  final parser = Parser();

  const String program = '55';

  var ast = parser.parse(program);

  inspect(ast);
}
