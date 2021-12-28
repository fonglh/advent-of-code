import 'dart:convert';
import 'dart:io';
import 'package:day22_reactor_reboot/point3.dart';

void main(List<String> arguments) async {
  var input = await readFile("22.txt");

  print(part1(input));
}

int part1(input) {
  Set<Point3> reactor = {};

  // for part 1, manually eyeball the list and see that only the first
  // twenty lines are between -50 to 50.
  for (var line in input.sublist(0, 20)) {
    var commandAndRanges = line.split(" ");
    var command = commandAndRanges[0];
    var ranges = commandAndRanges[1].split(",");

    var xRange =
        ranges[0].substring(2).split("..").map((e) => int.parse(e)).toList();
    var yRange =
        ranges[1].substring(2).split("..").map((e) => int.parse(e)).toList();
    var zRange =
        ranges[2].substring(2).split("..").map((e) => int.parse(e)).toList();

    for (int i = xRange[0]; i <= xRange[1]; i++) {
      for (int j = yRange[0]; j <= yRange[1]; j++) {
        for (int k = zRange[0]; k <= zRange[1]; k++) {
          Point3 cube = Point3(i, j, k);

          if (command == 'on') {
            reactor.add(cube);
          } else {
            reactor.remove(cube);
          }
        }
      }
    }
  }

  return reactor.length;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
