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
  print(countDigitsWithUniqueNumberOfSegments);
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
