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

  print(part1(outputValues));
  print(part2(signalPatterns, outputValues));
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

int part2(List<List<String>> signalPatterns, List<List<String>> outputValues) {
  int outputSum = 0;

  for (int i = 0; i < signalPatterns.length; i++) {
    var signalPattern = signalPatterns[i];
    var outputValue = outputValues[i];

    Map<String, int> segmentCount = getSegmentCounts(signalPattern);

    Map<String, int> outputLookup =
        buildOutputLookup(segmentCount, signalPattern);

    outputSum += getNumericalOutputValue(outputLookup, outputValue);
  }

  return outputSum;
}

Map<String, int> getSegmentCounts(List<String> signalPattern) {
  Map<String, int> segmentCount = {
    "a": 0,
    "b": 0,
    "c": 0,
    "d": 0,
    "e": 0,
    "f": 0,
    "g": 0
  };

  signalPattern.join("").split("").forEach((char) {
    segmentCount[char] = segmentCount[char]! + 1;
  });

  return segmentCount;
}

// Maps each signal pattern to its display value on a 7-segment display
Map<String, int> buildOutputLookup(
    Map<String, int> segmentCount, List<String> signalPattern) {
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
  Map<String, int> outputLookup = {};

  signalPattern.forEach((pattern) {
    List<int> digitKey = [];
    pattern.split("").forEach((char) {
      digitKey.add(segmentCount[char]!);
    });
    digitKey.sort();
    outputLookup[pattern] = countsToDigitsMap[digitKey.join("")]!;
  });
  // print(outputLookup);
  return outputLookup;
}

int getNumericalOutputValue(
    Map<String, int> outputLookup, List<String> outputValue) {
  String outputDigits = "";
  outputValue.forEach((element) {
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

  return int.parse(outputDigits);
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
