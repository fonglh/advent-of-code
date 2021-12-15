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

  // part 1
  print(findShortestPath(riskMap));

  // part 2
  // generate the full map
  List<List<int>> fullRiskMap = [];
  for (int row = 0; row < riskMap.length * 5; row++) {
    fullRiskMap.add(List.filled(riskMap[0].length * 5, 0));
  }
  for (int horizOffset = 0; horizOffset <= 4; horizOffset++) {
    for (int vertOffset = 0; vertOffset <= 4; vertOffset++) {
      for (int row = 0; row < riskMap.length; row++) {
        for (int col = 0; col < riskMap[0].length; col++) {
          fullRiskMap[row + vertOffset * riskMap.length]
                  [col + horizOffset * riskMap[0].length] =
              offsetRiskLevel(horizOffset, vertOffset, riskMap[row][col]);
        }
      }
    }
  }

  print(findShortestPath(fullRiskMap));
}

int offsetRiskLevel(int horizOffset, int vertOffset, int originalRisk) {
  int newRisk = originalRisk + horizOffset + vertOffset;
  return newRisk > 9 ? newRisk - 9 : newRisk;
}

// Be careful. Point() is represented as x, y coordinates, but the
// list of list maps are y, x.
// Dijkstra's algorithm translated from https://stackabuse.com/dijkstras-algorithm-in-python/
// Note that in the example on the webpage, vertex 3 should have been updated to 17.
int findShortestPath(List<List<int>> riskMap) {
  // Cumulative lowest risk to reach each point on the map.
  List<List<int>> costMap = [];
  for (int row = 0; row < riskMap.length; row++) {
    costMap.add(List.filled(riskMap[0].length, 99999));
  }
  Set<Point> visited = {};

  // This comparison puts the lowest risk at the front of the queue.
  final queue = PriorityQueue<Point>((a, b) {
    return costMap[a.y.toInt()][a.x.toInt()]
        .compareTo(costMap[b.y.toInt()][b.x.toInt()]);
  });

  // Start from the top left corner, but don't include its risk.
  visited.add(Point(0, 0));
  costMap[0][0] = 0;
  queue.add(Point(0, 0));

  while (queue.isNotEmpty) {
    Point currentPoint = queue.removeFirst();
    visited.add(currentPoint);

    int riskSoFar = costMap[currentPoint.y.toInt()][currentPoint.x.toInt()];

    for (int offsetX = -1; offsetX <= 1; offsetX++) {
      for (int offsetY = -1; offsetY <= 1; offsetY++) {
        // Omit diagonals and no offsets
        if (offsetX.abs() == offsetY.abs() || (offsetX == 0 && offsetY == 0)) {
          continue;
        }

        int newX = currentPoint.x.toInt() + offsetX;
        int newY = currentPoint.y.toInt() + offsetY;
        // Check that the new point is still on the map.
        if (newX < 0 ||
            newY < 0 ||
            newX >= riskMap[0].length ||
            newY >= riskMap.length) {
          continue;
        }

        if (!visited.contains(Point(newX, newY))) {
          // Cumulative risk of reaching this neighbour
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
