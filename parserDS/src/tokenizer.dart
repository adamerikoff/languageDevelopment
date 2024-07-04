const List<dynamic> spec = [
  // Implement here...
];

class Tokenizer {
  String string = '';
  int cursor = 0;

  void init(String string) {
    string = string;
    cursor = 0;
  }

  bool isEOF() {
    return cursor == string.length;
  }

  bool hasMoreTokens() {
    return cursor < string.length;
  }

  dynamic getNextToken() {
    // Implement here...
  }
}
