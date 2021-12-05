import 'dart:collection';
import 'dart:math';

import 'line.dart';

class Board {
  var board = HashMap<Point, int>();

  void addLine(Line line) {
    for (var point in line.getPoints()) {
      if (board[point] == null) {
        board[point] = 1;
      } else {
        board[point] = board[point]! + 1;
      }
    }
  }

  int numDangerPoints() {
    var dangerPoints = 0;
    for (var point in board.keys) {
      if (board[point]! >= 2) {
        dangerPoints++;
      }
    }
    return dangerPoints;
  }
}
