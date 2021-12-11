import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("11.txt");

  var energyLevels = initEnergyLevels(input);
  print(part1(energyLevels));

  // Need to reset the energyLevels as the functions modify the array directly.
  energyLevels = initEnergyLevels(input);
  print(part2(energyLevels));
}

int part1(List<List<int>> energyLevels) {
  int flashCount = 0;
  for (int step = 0; step < 100; step++) {
    flashCount += takeStep(energyLevels);
  }
  return flashCount;
}

int part2(List<List<int>> energyLevels) {
  int step = 0;
  while (!allFlashed(energyLevels)) {
    takeStep(energyLevels);
    step++;
  }
  return step;
}

bool allFlashed(List<List<int>> energyLevels) {
  return energyLevels.every((row) => row.every((element) => element == 0));
}

int takeStep(List<List<int>> energyLevels) {
  int flashCount = 0;
  increaseByOne(energyLevels);
  for (int i = 0; i < energyLevels.length; i++) {
    for (int j = 0; j < energyLevels[i].length; j++) {
      flashCount += flash(i, j, energyLevels);
    }
  }

  return flashCount;
}

int flash(int row, int col, List<List<int>> energyLevels) {
  // out of bounds
  if (row < 0 ||
      col < 0 ||
      row >= energyLevels.length ||
      col >= energyLevels[row].length) {
    return 0;
  }

  // already flashed, do nothing.
  if (energyLevels[row][col] == 0) {
    return 0;
  }

  // not enough energy to flash
  if (energyLevels[row][col] <= 9) {
    return 0;
  }

  // flash
  int flashCount = 1;
  energyLevels[row][col] = 0;
  for (var i = -1; i <= 1; i++) {
    for (var j = -1; j <= 1; j++) {
      if (shouldIncreaseEnergy(row + i, col + j, energyLevels)) {
        energyLevels[row + i][col + j] += 1;
        flashCount += flash(row + i, col + j, energyLevels);
      }
    }
  }

  return flashCount;
}

bool shouldIncreaseEnergy(int row, int col, List<List<int>> energyLevels) {
  return row >= 0 &&
      col >= 0 &&
      row < energyLevels.length &&
      col < energyLevels[row].length &&
      energyLevels[row][col] != 0;
}

void increaseByOne(List<List<int>> energyLevels) {
  for (int i = 0; i < energyLevels.length; i++) {
    for (int j = 0; j < energyLevels[i].length; j++) {
      energyLevels[i][j] = energyLevels[i][j] + 1;
    }
  }
}

List<List<int>> initEnergyLevels(List<String> input) {
  List<List<int>> energyLevels = [];
  for (var line in input) {
    energyLevels.add(line.split("").map((e) => int.parse(e)).toList());
  }
  return energyLevels;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
