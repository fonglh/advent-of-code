import 'dart:convert';
import 'dart:io';
import 'dart:math';

// All operations need to be done on List<String> so double digit numbers
// that happen during the 'explode' operation are easier to handle.
// Keep the commas for easier debugging, can just use 'join'.
// https://github.com/fridokus/advent-of-code/blob/master/2021/18.py
void main(List<String> arguments) async {
  final input = await readFile("18.txt");

  List<String> finalSum = input[0].split("");

  // Part 1
  for (int i = 1; i < input.length; i++) {
    finalSum = add(finalSum, input[i].split(""));
  }
  print(magnitude(finalSum.join()));

  // Part 2
  int maxMagnitude = 0;
  for (int i = 0; i < input.length; i++) {
    for (int j = i + 1; j < input.length; j++) {
      int sum1 = magnitude(add(input[i].split(""), input[j].split("")).join());
      int sum2 = magnitude(add(input[j].split(""), input[i].split("")).join());

      maxMagnitude = max(maxMagnitude, max(sum1, sum2));
    }
  }
  print(maxMagnitude);
}

List<String> reduce(List<String> snail) {
  while (true) {
    int? explodeIndex = shouldExplode(snail);
    while (explodeIndex != null) {
      snail = explode(snail, explodeIndex);
      explodeIndex = shouldExplode(snail);
    }

    int? splitIndex = shouldSplit(snail);
    if (splitIndex != null) {
      snail = split(snail, splitIndex);
    } else {
      break;
    }
  }

  return snail;
}

List<String> split(List<String> snail, int index) {
  int value = int.parse(snail[index]);
  int left = value ~/ 2;
  int right = value % 2 == 1 ? left + 1 : left;
  List<String> pair =
      ["["] + [left.toString()] + [","] + [right.toString()] + ["]"];
  return snail.sublist(0, index) + pair + snail.sublist(index + 1);
}

int? shouldSplit(List<String> snail) {
  for (int i = 0; i < snail.length; i++) {
    int? value = int.tryParse(snail[i]);

    if (value != null && value >= 10) {
      return i;
    }
  }

  return null;
}

List<String> explode(List<String> snail, int index) {
  // search left
  for (int i = index - 1; i >= 0; i--) {
    if (!"[],".contains(snail[i])) {
      snail[i] = (int.parse(snail[i]) + int.parse(snail[index])).toString();
      break;
    }
  }
  // search right after right of pair, will always be a regular number
  for (int i = index + 3; i < snail.length; i++) {
    if (!"[],".contains(snail[i])) {
      snail[i] = (int.parse(snail[i]) + int.parse(snail[index + 2])).toString();
      break;
    }
  }

  // sublist from start to just before the left pair regular number
  // put a regular 0
  // sublist after ",<right pair regular number>]" i.e. +4 elements
  return snail.sublist(0, index - 1) + ["0"] + snail.sublist(index + 4);
}

int? shouldExplode(List<String> snail) {
  int depth = 0;
  for (int i = 0; i < snail.length; i++) {
    if (depth > 4) return i;

    if (snail[i] == "[") {
      depth++;
    } else if (snail[i] == "]") {
      depth--;
    }
  }
  return null;
}

List<String> add(List<String> snail1, List<String> snail2) {
  return reduce(["["] + snail1 + [","] + snail2 + ["]"]);
}

// https://github.com/Myxcil/AdventOfCode2021/blob/main/Day18/day18_main.py
// A solution doesn't use python's eval to convert the string to nested lists.
// Looks similar to the algo for evaluating postfix expressions.
// In the py solution, the input is a list without the commas.
int magnitude(String snail) {
  List<int> stack = [];
  snail.split("").forEach((element) {
    // everything is a single digit number after reduce
    if (!"[],".contains(element)) {
      stack.add(int.parse(element));
    } else if (element == "]") {
      // the "right" number of the pair is top of the stack, followed by the "left" number of the pair.
      // push the result back until popped and processed by the next encounter of "]".
      stack.add(2 * stack.removeLast() + 3 * stack.removeLast());
    }
  });
  return stack.first;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
