import 'dart:convert';
import 'dart:collection';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("10.txt");

  print(part1(input));

  List<String> incompleteLines =
      input.where((line) => findIllegalCharacter(line) == null).toList();

  List<int> completionScores = [];

  for (String line in incompleteLines) {
    completionScores.add(scoreCompletionString(getCompletionString(line)));
  }
  completionScores.sort();
  print(completionScores[completionScores.length ~/ 2]);
}

int scoreCompletionString(String completionString) {
  int score = 0;

  for (String char in completionString.split("")) {
    score *= 5;
    switch (char) {
      case ")":
        score += 1;
        break;
      case "]":
        score += 2;
        break;
      case "}":
        score += 3;
        break;
      case ">":
        score += 4;
    }
  }

  return score;
}

String getCompletionString(String incompleteLine) {
  ListQueue<String> stack = ListQueue();
  String result = "";
  Map<String, String> bracketPairs = {
    "(": ")",
    "{": "}",
    "[": "]",
    "<": ">",
  };

  for (var char in incompleteLine.split("")) {
    if (char == "(" || char == "[" || char == "{" || char == "<") {
      stack.addLast(char);
    } else if (char == ")" || char == "]" || char == "}" || char == ">") {
      stack.removeLast();
    }
  }

  while (stack.isNotEmpty) {
    result += bracketPairs[stack.removeLast()]!;
  }

  return result;
}

int part1(List<String> input) {
  Map<String, int> scoringTable = {")": 3, "]": 57, "}": 1197, ">": 25137};
  int score = 0;

  for (String line in input) {
    String? illegalChar = findIllegalCharacter(line);

    if (illegalChar != null) {
      score += scoringTable[illegalChar]!;
    }
  }

  return score;
}

String? findIllegalCharacter(String line) {
  ListQueue<String> stack = ListQueue();

  for (var char in line.split("")) {
    if (char == "(" || char == "[" || char == "{" || char == "<") {
      stack.addLast(char);
    } else if (char == ")" || char == "]" || char == "}" || char == ">") {
      if (stack.isEmpty) {
        return char;
      } else {
        String last = stack.removeLast();
        if (last == "(" && char != ")") {
          return char;
        } else if (last == "[" && char != "]") {
          return char;
        } else if (last == "{" && char != "}") {
          return char;
        } else if (last == "<" && char != ">") {
          return char;
        }
      }
    }
  }

  return null;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
