import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("09.txt");

  List<List<int>> heightMap = [];
  for (var line in input) {
    heightMap.add(line.split("").map((e) => int.parse(e)).toList());
  }

  print(part1(heightMap));
  print(part2(heightMap));
}

int part1(List<List<int>> heightMap) {
  int riskSum = 0;

  for (int row = 0; row < heightMap.length; row++) {
    for (int col = 0; col < heightMap[0].length; col++) {
      if (isLowPoint(row, col, heightMap)) {
        riskSum += 1 + heightMap[row][col];
      }
    }
  }

  return riskSum;
}

int part2(List<List<int>> heightMap) {
  List<int> allBasinSizes = [];
  for (int row = 0; row < heightMap.length; row++) {
    for (int col = 0; col < heightMap[0].length; col++) {
      if (isLowPoint(row, col, heightMap)) {
        var checked = List.generate(heightMap.length,
            (i) => List.generate(heightMap[0].length, (j) => false));

        allBasinSizes.add(basinSize(row, col, heightMap, checked));
      }
    }
  }

  allBasinSizes.sort();
  allBasinSizes = allBasinSizes.reversed.toList();

  return allBasinSizes[0] * allBasinSizes[1] * allBasinSizes[2];
}

int basinSize(
    int row, int col, List<List<int>> heightMap, List<List<bool>> checked) {
  if (row < 0 ||
      col < 0 ||
      row >= heightMap.length ||
      col >= heightMap[0].length) {
    return 0;
  } else if (heightMap[row][col] == 9 || checked[row][col]) {
    return 0;
  } else {
    checked[row][col] = true;
    return 1 +
        basinSize(row + 1, col, heightMap, checked) +
        basinSize(row - 1, col, heightMap, checked) +
        basinSize(row, col - 1, heightMap, checked) +
        basinSize(row, col + 1, heightMap, checked);
  }
}

bool isLowPoint(int row, int col, List<List<int>> heightMap) {
  for (int i = -1; i <= 1; i++) {
    for (int j = -1; j <= 1; j++) {
      int rowCheck = row + i;
      int colCheck = col + j;

      // Check out of bounds
      if (rowCheck < 0 ||
          colCheck < 0 ||
          rowCheck >= heightMap.length ||
          colCheck >= heightMap[0].length) {
        continue;
      }
      // Diagonals are not considered adjacent, this also skips 0, 0
      if (i.abs() == j.abs()) {
        continue;
      }

      // Higher or equal to something adjacent, so not the lowest point
      if (heightMap[row][col] >= heightMap[rowCheck][colCheck]) {
        return false;
      }
    }
  }
  return true;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
