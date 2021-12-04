import 'dart:convert';
import 'dart:io';
import 'package:day04_giant_squid/board.dart';

void main(List<String> arguments) async {
  final input = await readFile("04.txt");

  List<int> drawnNumbers =
      input[0].split(",").map((s) => int.parse(s)).toList();

  final boardsInput = input.sublist(2);
  List<Board> boards = [];

  // 5 lines for the board, 1 line for the empty line between boards.
  for (int i = 0; i < boardsInput.length; i += 6) {
    final currentBoardList = boardsInput.sublist(i, i + 5);
    boards.add(Board(currentBoardList));
  }

  print(firstWinningBoard(boards, drawnNumbers));
  print(lastWinningBoard(boards, drawnNumbers));
}

int firstWinningBoard(List<Board> boards, List<int> drawnNumbers) {
  for (int num in drawnNumbers) {
    for (Board board in boards) {
      board.mark(num);
      if (board.anyWin()) {
        return board.score(num);
      }
    }
  }

  return -1;
}

int lastWinningBoard(List<Board> boards, List<int> drawnNumbers) {
  int latestScore = -1;

  for (int num in drawnNumbers) {
    for (Board board in boards) {
      if (!board.anyWin()) {
        board.mark(num);
        if (board.anyWin()) {
          latestScore = board.score(num);
        }
      }
    }
  }

  return latestScore;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
