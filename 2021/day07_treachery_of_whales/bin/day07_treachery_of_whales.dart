import 'dart:convert';
import 'dart:io';
import 'dart:math';

void main(List<String> arguments) async {
  final input = await readFile("07.txt");

  List<int> crabPositions =
      input[0].split(",").map((e) => int.parse(e)).toList();

  int maxPosition = crabPositions.reduce(max);
  // Initialise assuming all align on position 0.
  //int minFuel = crabPositions.reduce((a, b) => a + b);
  int minFuel =
      crabPositions.map((pos) => partTwoFuel(pos, 0)).reduce((a, b) => a + b);
  int minPos = 0;

  for (int position = 1; position <= maxPosition; position++) {
    int fuelRequired = crabPositions
        .map((pos) => partTwoFuel(pos, position))
        .reduce((a, b) => a + b);
    //print('$position: $fuelRequired');
    if (fuelRequired < minFuel) {
      minFuel = fuelRequired;
      minPos = position;
    }
  }

  print(minFuel);
  print(minPos);
}

int partOneFuel(int currPos, int target) {
  return (currPos - target).abs();
}

int partTwoFuel(int currPos, int target) {
  int steps = (currPos - target).abs();
  return (steps * (steps + 1) ~/ 2);
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
