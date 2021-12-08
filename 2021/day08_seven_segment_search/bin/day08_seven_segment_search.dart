import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("08.txt");

  List<List<String>> signalPatterns = [];
  List<List<String>> outputValues = [];

  // Parse input
  for (var line in input) {
    var lineParts = line.split(" | ");

    var signalPattern = lineParts[0].split(" ");
    signalPatterns.add(signalPattern);

    var outputValue = lineParts[1].split(" ");
    outputValues.add(outputValue);
  }

  // https://www.reddit.com/r/adventofcode/comments/rbj87a/2021_day_8_solutions/hnpad75/
  Map<String, int> countsToDigitsMap = {
    '467889': 0,
    '89': 1,
    '47788': 2,
    '77889': 3,
    '6789': 4,
    '67789': 5,
    '467789': 6,
    '889': 7,
    '4677889': 8,
    '677889': 9
  };

  //print(part1(outputValues));

  // Very messy part 2
  int outputSum = 0;

  for (int i = 0; i < signalPatterns.length; i++) {
    var signalPattern = signalPatterns[i];
    var outputValue = outputValues[i];
    Map<String, int> segmentCount = {
      "a": 0,
      "b": 0,
      "c": 0,
      "d": 0,
      "e": 0,
      "f": 0,
      "g": 0
    };

    Map<String, int> outputLookup = {};

    signalPatterns[i].join("").split("").forEach((char) {
      segmentCount[char] = segmentCount[char]! + 1;
    });
    signalPatterns[i].forEach((pattern) {
      List<int> digitKey = [];
      pattern.split("").forEach((char) {
        digitKey.add(segmentCount[char]!);
      });
      digitKey.sort();
      outputLookup[pattern] = countsToDigitsMap[digitKey.join("")]!;
    });
    // print(outputLookup);

    String outputDigits = "";
    outputValues[i].forEach((element) {
      Set elementSet = Set();
      elementSet.addAll(element.split(""));

      for (var key in outputLookup.keys) {
        Set keySet = Set();
        keySet.addAll(key.split(""));
        if (setEquals(keySet, elementSet)) {
          outputDigits += outputLookup[key].toString();
        }
      }
    });

    outputSum += int.parse(outputDigits);
  }

  print(outputSum);
}

int part1(List<List<String>> outputValues) {
  int countDigitsWithUniqueNumberOfSegments = 0;
  for (var output in outputValues) {
    for (var digit in output) {
      if (digit.length == 2 ||
          digit.length == 3 ||
          digit.length == 4 ||
          digit.length == 7) {
        countDigitsWithUniqueNumberOfSegments++;
      }
    }
  }
  return countDigitsWithUniqueNumberOfSegments;
}

bool setEquals(Set a, Set b) {
  if (a.length != b.length) return false;

  return a.containsAll(b.toList());
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
