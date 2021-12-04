class Board {
  List<int> board = [];
  late List<bool> marks;
  int winOrder = 0;

  Board(List<String> boardString) {
    // To handle right aligned single digit numbers.
    var re = RegExp(r"\s+");
    for (var row in boardString) {
      var rowList = row.trim().split(re).map((s) => int.parse(s)).toList();
      board = [...board, ...rowList];
    }
    marks = List<bool>.filled(25, false);
  }

  void mark(int value) {
    for (int i = 0; i < board.length; i++) {
      if (board[i] == value) {
        marks[i] = true;
      }
    }
  }

  bool anyWin() {
    return anyRowWin() || anyColumnWin();
  }

  bool anyRowWin() {
    for (int i = 0; i < marks.length; i += 5) {
      final row = marks.sublist(i, i + 5);
      if (row.every((m) => m)) {
        return true;
      }
    }
    return false;
  }

  bool anyColumnWin() {
    for (int i = 0; i < 5; i++) {
      final column = [
        marks[i],
        marks[i + 5],
        marks[i + 10],
        marks[i + 15],
        marks[i + 20]
      ];
      if (column.every((m) => m)) {
        return true;
      }
    }
    return false;
  }

  int score(int winningNumber) {
    int currSum = 0;
    for (int i = 0; i < board.length; i++) {
      if (marks[i] == false) {
        currSum += board[i];
      }
    }

    return currSum * winningNumber;
  }

  @override
  String toString() {
    String result = "";

    for (int i = 0; i < board.length; i++) {
      if (marks[i]) {
        result += "*" + board[i].toString() + "* ";
      } else {
        result += board[i].toString() + " ";
      }

      if (i % 5 == 4) {
        result += "; ";
      }
    }

    return result;
  }
}
