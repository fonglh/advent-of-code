import 'dart:convert';
import 'dart:io';
import 'package:day22_reactor_reboot/point3.dart';
import 'package:day22_reactor_reboot/cuboid.dart';

void main(List<String> arguments) async {
  var input = await readFile("22.txt");

  print(part1(input));
  print(part2(input));
}

// Adapted from https://github.com/abeltay/advent-of-code-2021/blob/main/22/part2/runner.go
int part2(input) {
  List<Cuboid> cuboidList = [];

  for (var line in input) {
    Cuboid currentCuboid = Cuboid.fromInput(line);

    if (currentCuboid.on) {
      List<Cuboid> newCuboids = [currentCuboid];
      for (Cuboid cuboid in cuboidList) {
        // check intersection
        Cuboid? intersection = cuboid.getIntersection(currentCuboid);
        // if an existing cuboid is on, add an off cuboid for the intersection
        // to prevent double counting the ONs
        if (cuboid.on) {
          if (intersection != null) {
            newCuboids.add(intersection);
          }
        } else if (cuboid.on == false && intersection != null) {
          intersection.on = true;
          newCuboids.add(intersection);
        }
      }
      cuboidList.addAll(newCuboids);
    } else {
      // don't add an OFF cuboid to the list
      List<Cuboid> newCuboids = [];

      for (Cuboid cuboid in cuboidList) {
        Cuboid? intersection = cuboid.getIntersection(currentCuboid);
        if (cuboid.on && intersection != null) {
          newCuboids.add(intersection);
        } else if (cuboid.on == false && intersection != null) {
          // add an ON here because line 42 added an OFF already when it
          // encountered the corresponding ON cuboid.
          // so for a specific cube that had ON-OFF, when this one sees the ON,
          // it adds OFF, which needs to be negated with another ON.
          // final result ON-OFF-OFF-ON.
          intersection.on = true;
          newCuboids.add(intersection);
        }
      }

      cuboidList.addAll(newCuboids);
    }
  }

  return cuboidList.fold(
      0, (int volumeSoFar, Cuboid cuboid) => volumeSoFar + cuboid.volume());
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
