import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("14.txt");
  //List<String> polymerTemplate = "NNCB".split("").toList();
  List<String> polymerTemplate = "VPPHOPVVSFSVFOCOSBKF".split("").toList();

  Map<String, String> insertionRules = {};
  for (String line in input) {
    var ruleParts = line.split(" -> ");
    insertionRules[ruleParts[0]] = ruleParts[1];
  }

  for (int step = 0; step < 10; step++) {
    print(step);
    processPolymer(polymerTemplate, insertionRules);
  }

  var polymerFrequencies = frequencyMap(polymerTemplate);
  print(polymerFrequencies);

  // Manually inspect frequencies, the answer for part 1 is 2233.
}

void processPolymer(
    List<String> polymerTemplate, Map<String, String> insertionRules) {
  for (int i = 0; i < polymerTemplate.length - 1; i++) {
    String pair = polymerTemplate[i] + polymerTemplate[i + 1];
    if (insertionRules.containsKey(pair)) {
      polymerTemplate.insert(i + 1, insertionRules[pair]!);
      i++;
    }
  }
}

Map<String, int> frequencyMap(List<String> polymerTemplate) {
  Map<String, int> result = {};

  for (String ch in polymerTemplate) {
    if (result.containsKey(ch)) {
      result[ch] = result[ch]! + 1;
    } else {
      result[ch] = 1;
    }
  }

  return result;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
