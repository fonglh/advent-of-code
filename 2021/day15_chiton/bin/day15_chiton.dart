import 'dart:convert';
import 'dart:io';
import 'dart:math';
import 'package:collection/collection.dart';

void main(List<String> arguments) async {
  final input = await readFile("15.txt");

  List<List<int>> riskMap = [];
  for (var line in input) {
    riskMap.add(line.split("").map((e) => int.parse(e)).toList());
  }

  print(findShortestPath(riskMap));
}

// Be careful. Point() is represented as x, y coordinates, but the
// list of list maps are y, x.
int findShortestPath(List<List<int>> riskMap) {
  List<List<int>> costMap = [];
  for (int row = 0; row < riskMap.length; row++) {
    costMap.add(List.filled(riskMap[0].length, 9999));
  }
  Set<Point> visited = {};

  final queue = PriorityQueue<Point>((a, b) {
    return costMap[a.y.toInt()][a.x.toInt()]
        .compareTo(costMap[b.y.toInt()][b.x.toInt()]);
  });

  visited.add(Point(0, 0));
  costMap[0][0] = 0;
  queue.add(Point(0, 0));

  while (queue.isNotEmpty) {
    Point currentPoint = queue.removeFirst();
    visited.add(currentPoint);

    int riskSoFar = costMap[currentPoint.y.toInt()][currentPoint.x.toInt()];

    for (int offsetX = -1; offsetX <= 1; offsetX++) {
      for (int offsetY = -1; offsetY <= 1; offsetY++) {
        if (offsetX.abs() == offsetY.abs() || (offsetX == 0 && offsetY == 0)) {
          continue;
        }

        int newX = currentPoint.x.toInt() + offsetX;
        int newY = currentPoint.y.toInt() + offsetY;
        if (newX < 0 ||
            newY < 0 ||
            newX >= riskMap[0].length ||
            newY >= riskMap.length) {
          continue;
        }

        if (!visited.contains(Point(newX, newY))) {
          int newCost = riskSoFar + riskMap[newY][newX];
          if (newCost < costMap[newY][newX]) {
            costMap[newY][newX] = newCost;
            queue.add(Point(newX, newY));
          }
        }
      }
    }
  }

  return costMap.last.last;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
