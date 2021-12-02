import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("02.txt");

  print(part1(input));
  print(part2(input));
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  List<String> commandList = await lines.toList();

  return commandList;
}

int part1(List<String> input) {
  var horizPos = 0;
  var depth = 0;

  for (final line in input) {
    final commands = line.split(" ");

    switch (commands[0]) {
      case 'forward':
        horizPos += int.parse(commands[1]);
        break;
      case 'up':
        depth -= int.parse(commands[1]);
        break;
      case 'down':
        depth += int.parse(commands[1]);
        break;
    }
  }

  return horizPos * depth;
}

int part2(List<String> input) {
  var horizPos = 0;
  var depth = 0;
  var aim = 0;

  for (final line in input) {
    final commands = line.split(" ");

    switch (commands[0]) {
      case 'forward':
        var commandAmount = int.parse(commands[1]);
        horizPos += commandAmount;
        depth += aim * commandAmount;
        break;
      case 'up':
        aim -= int.parse(commands[1]);
        break;
      case 'down':
        aim += int.parse(commands[1]);
        break;
    }
  }

  return horizPos * depth;
}
