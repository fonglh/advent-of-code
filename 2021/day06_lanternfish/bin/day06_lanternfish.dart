import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("06-test.txt");

  List<int> fishes = input[0].split(",").map((e) => int.parse(e)).toList();

  for (int day = 0; day < 80; day++) {
    fishes = advanceDay(fishes);
  }

  print(fishes.length);
}

List<int> advanceDay(List<int> fishes) {
  List<int> newFishes = [];
  int toAdd = 0;

  for (int fish in fishes) {
    if (fish == 0) {
      newFishes.add(6);
      toAdd += 1;
    } else {
      newFishes.add(fish - 1);
    }
  }

  // Implement this way for easier comparison with sample output
  for (int i = 0; i < toAdd; i++) {
    newFishes.add(8);
  }

  return newFishes;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
