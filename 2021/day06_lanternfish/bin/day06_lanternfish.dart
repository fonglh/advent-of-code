import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("06.txt");

  List<int> fishList = input[0].split(",").map((e) => int.parse(e)).toList();

  var fishes = initFishes(fishList);

  for (int day = 0; day < 256; day++) {
    fishes = advanceDay(fishes);
  }

  int totalFishes = 0;
  for (int timer = 0; timer <= 8; timer++) {
    totalFishes += fishes[timer]!;
  }

  print(totalFishes);
}

Map<int, int> initFishes(List<int> fishList) {
  var fishes = {0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0};

  for (int fishTimer in fishList) {
    fishes[fishTimer] = fishes[fishTimer]! + 1;
  }

  return fishes;
}

Map<int, int> advanceDay(Map<int, int> fishes) {
  var nextFishes = {0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0};

  for (int timer = 8; timer >= 0; timer--) {
    if (timer == 0) {
      nextFishes[8] = fishes[0]!;
      nextFishes[6] = nextFishes[6]! + fishes[0]!;
    } else {
      nextFishes[timer - 1] = fishes[timer]!;
    }
  }
  return nextFishes;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
