import 'dart:convert';
import 'dart:io';
import 'package:day04_giant_squid/board.dart';

void main(List<String> arguments) async {
  final input = await readFile("04.txt");

  List<int> drawnNumbers =
      input[0].split(",").map((s) => int.parse(s)).toList();

  final boardsInput = input.sublist(2);
  List<Board> boards = [];

  for (int i = 0; i < boardsInput.length; i += 6) {
    final currentBoardList = boardsInput.sublist(i, i + 5);
    boards.add(Board(currentBoardList));
  }

  for (int num in drawnNumbers) {
    for (Board board in boards) {
      board.mark(num);
      if (board.anyWin()) {
        print("WINNER");
        print(board.score(num));
        exit(0);
      }
    }
  }

  print(boards[0]);
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
