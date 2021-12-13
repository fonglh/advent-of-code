import 'dart:convert';
import 'dart:io';

import 'package:day13_transparent_origami/transparency.dart';

void main(List<String> arguments) async {
  final input = await readFile("13.txt");

  // find blank line between dot coordinates and fold instructions
  int separatorIndex = 0;
  for (int i = 0; i < input.length; i++) {
    if (input[i] == "") {
      separatorIndex = i;
      break;
    }
  }

  List<String> dotCoordinates = input.sublist(0, separatorIndex);

  var transparency = Transparency(dotCoordinates);
  transparency.foldAlongX(655);
  // Part 1
  print(transparency.numDots());
  transparency.foldAlongY(447);
  transparency.foldAlongX(327);
  transparency.foldAlongY(223);
  transparency.foldAlongX(163);
  transparency.foldAlongY(111);
  transparency.foldAlongX(81);
  transparency.foldAlongY(55);
  transparency.foldAlongX(40);
  transparency.foldAlongY(27);
  transparency.foldAlongY(13);
  transparency.foldAlongY(6);
  // Part 2
  print(transparency);
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
