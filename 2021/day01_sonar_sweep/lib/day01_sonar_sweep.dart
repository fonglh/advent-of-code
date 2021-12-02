import 'dart:convert';
import 'dart:io';

Future<List<int>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  List<String> depthList = await lines.toList();

  return depthList.map(int.parse).toList();
}

int countIncreases(List<int> depthInts) {
  var currentDepth = depthInts.first;
  var numIncreases = 0;

  for (final depth in depthInts) {
    if (depth > currentDepth) {
      numIncreases++;
    }
    currentDepth = depth;
  }

  return numIncreases;
}

int windowCountIncreases(List<int> depthInts) {
  List<int> windowDepths = [];

  for (int i = 1; i < depthInts.length - 1; i++) {
    windowDepths.add(depthInts[i - 1] + depthInts[i] + depthInts[i + 1]);
  }
  return countIncreases(windowDepths);
}
