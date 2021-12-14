import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("14.txt");
  // List<String> polymerTemplate = "NNCB".split("").toList();
  List<String> polymerTemplate = "VPPHOPVVSFSVFOCOSBKF".split("").toList();

  Map<String, String> insertionRules = {};
  for (String line in input) {
    var ruleParts = line.split(" -> ");
    insertionRules[ruleParts[0]] = ruleParts[1];
  }

  print(part1(polymerTemplate, insertionRules));

  // --------------------- Part 2 ----------------------------------------
  // polymerTemplate = "NNCB".split("").toList();
  polymerTemplate = "VPPHOPVVSFSVFOCOSBKF".split("").toList();
  var polymerPairFrequencies = pairFrequencyMap(polymerTemplate);
  var charCounts = frequencyMap(polymerTemplate);

  for (int step = 0; step < 40; step++) {
    polymerPairFrequencies =
        processPolymerMap(polymerPairFrequencies, charCounts, insertionRules);
  }
  print(findDifference(charCounts));
}

int part1(List<String> polymerTemplate, Map<String, String> insertionRules) {
  for (int step = 0; step < 10; step++) {
    processPolymer(polymerTemplate, insertionRules);
  }

  var polymerFrequencies = frequencyMap(polymerTemplate);
  return findDifference(polymerFrequencies);
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

Map<String, int> processPolymerMap(Map<String, int> polymerPairFrequencies,
    Map<String, int> charCounts, Map<String, String> insertionRules) {
  Map<String, int> newMap = {...polymerPairFrequencies};

  for (String key in polymerPairFrequencies.keys) {
    if (polymerPairFrequencies[key]! > 0 && insertionRules.containsKey(key)) {
      // Remove the original pair
      newMap[key] = newMap[key]! - polymerPairFrequencies[key]!;

      // Add new pairs
      String pair1 = key[0] + insertionRules[key]!;
      String pair2 = insertionRules[key]! + key[1];

      //print("${key} -> ${insertionRules[key]}");
      if (charCounts.containsKey(insertionRules[key])) {
        charCounts[insertionRules[key]!] =
            charCounts[insertionRules[key]!]! + polymerPairFrequencies[key]!;
      } else {
        charCounts[insertionRules[key]!] = polymerPairFrequencies[key]!;
      }

      if (newMap.containsKey(pair1)) {
        newMap[pair1] = newMap[pair1]! + polymerPairFrequencies[key]!;
      } else {
        newMap[pair1] = polymerPairFrequencies[key]!;
      }
      if (newMap.containsKey(pair2)) {
        newMap[pair2] = newMap[pair2]! + polymerPairFrequencies[key]!;
      } else {
        newMap[pair2] = polymerPairFrequencies[key]!;
      }
    }
  }

  return newMap;
}

int findDifference(Map<String, int> charCounts) {
  var counts = charCounts.values.toList();
  counts.sort();
  return counts.last - counts.first;
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

Map<String, int> pairFrequencyMap(List<String> polymerTemplate) {
  Map<String, int> result = {};

  for (int i = 0; i < polymerTemplate.length - 1; i++) {
    String pair = polymerTemplate[i] + polymerTemplate[i + 1];
    if (result.containsKey(pair)) {
      result[pair] = result[pair]! + 1;
    } else {
      result[pair] = 1;
    }
  }

  return result;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
