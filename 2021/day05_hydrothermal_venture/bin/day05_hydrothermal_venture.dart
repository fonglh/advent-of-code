import 'dart:convert';
import 'dart:io';
import 'dart:math';

import 'package:day05_hydrothermal_venture/line.dart';
import 'package:day05_hydrothermal_venture/board.dart';

void main(List<String> arguments) async {
  final input = await readFile("05.txt");

  List<Line> lines = [];

  for (var inputLine in input) {
    final inputLineData = inputLine.split(' ');
    Point p1 = Point(int.parse(inputLineData[0].split(',')[0]),
        int.parse(inputLineData[0].split(',')[1]));
    Point p2 = Point(int.parse(inputLineData[2].split(',')[0]),
        int.parse(inputLineData[2].split(',')[1]));

    lines.add(Line(p1, p2));
  }

  var board = Board();
  for (var line in lines) {
    board.addLine(line);
  }

  // part 2's answer. Delete diagonals in Line class to get part 1's answer.
  print(board.numDangerPoints());
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
