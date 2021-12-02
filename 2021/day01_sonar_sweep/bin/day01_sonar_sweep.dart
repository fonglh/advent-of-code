import 'package:day01_sonar_sweep/day01_sonar_sweep.dart';

import 'dart:io';

const lineNumber = 'line-number';

void main(List<String> arguments) async {
  exitCode = 0; // presume success

  List<int> depthInts = await readFile("01.txt");
  stdout.writeln(countIncreases(depthInts));
  stdout.writeln(windowCountIncreases(depthInts));
}
