import 'dart:convert';
import 'dart:io';
import 'package:day19_beacon_scanner/point3.dart';
import 'package:day19_beacon_scanner/scanner.dart';

void main(List<String> arguments) async {
  final input = await readFile("19.txt");

  List<Scanner> scanners = [];
  List<String> beaconCoords = [];
  for (int i = 0; i < input.length; i++) {
    if (input[i].startsWith("---")) {
      beaconCoords = [];
    } else if (input[i] == "" || i == input.length - 1) {
      // for adding the very last line of input
      if (i == input.length - 1) {
        beaconCoords.add(input[i]);
      }
      Scanner scanner = Scanner(beaconCoords);
      scanners.add(scanner);
    } else {
      beaconCoords.add(input[i]);
    }
  }

  scanners[0].normalized = true;
  scanners[0].position = Point3(0, 0, 0);

  List<int> normalizedIndices = [0];
  List<int> notNormalizedIndices =
      List.generate(scanners.length - 1, (i) => i + 1);

  while (notNormalizedIndices.isNotEmpty) {
    // find a scanner pair with enough overlapping beacons
    for (int i = 0; i < normalizedIndices.length; i++) {
      for (int j = 0; j < notNormalizedIndices.length; j++) {
        int currNormIdx = normalizedIndices[i];
        int currNotNormIdx = notNormalizedIndices[j];

        int? transformIndex =
            scanners[currNormIdx].getTransformIndex(scanners[currNotNormIdx]);

        // normalize scanner
        if (transformIndex != null) {
          scanners[currNotNormIdx].updateBeaconOrientations(transformIndex);
          Point3 relativePosition = scanners[currNormIdx]
              .getRelativePositionFromScanner0(scanners[currNotNormIdx]);
          scanners[currNotNormIdx].translateBeacons(relativePosition);
          scanners[currNotNormIdx].position = relativePosition;
          scanners[currNotNormIdx].normalized = true;

          normalizedIndices.add(currNotNormIdx);
          notNormalizedIndices.remove(currNotNormIdx);
        }
      }
    }
  }

  // Part 1
  print(countBeacons(scanners));
}

// Assumes all the beacons in scanners have been normalized to scanner 0
int countBeacons(List<Scanner> scanners) {
  Set<Point3> fullBeaconList = {};
  for (var scanner in scanners) {
    for (var beacon in scanner.beacons) {
      fullBeaconList.add(beacon);
    }
  }

  return fullBeaconList.length;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
