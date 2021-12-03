import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("03.txt");

  print(powerConsumption(input));
  print(lifeSupportRating(input));
}

int powerConsumption(List<String> input) {
  int gamma = 0;
  int epsilon = 0;

  for (int pos = 0; pos < input.first.length; pos++) {
    int nextDigit = mostCommonBit(input, pos);
    gamma = gamma << 1;
    gamma += nextDigit;

    nextDigit = nextDigit == 1 ? 0 : 1;
    epsilon = epsilon << 1;
    epsilon += nextDigit;
  }

  return gamma * epsilon;
}

int lifeSupportRating(List<String> input) {
  return lifeSupportSystemRating(input, mostCommonBit) *
      lifeSupportSystemRating(input, leastCommonBit);
}

int lifeSupportSystemRating(
    List<String> input, int Function(List<String>, int) bitCriteriaFn) {
  List<String> candidates = input;

  for (int pos = 0; pos < input.first.length; pos++) {
    int criteriaBit = bitCriteriaFn(candidates, pos);

    candidates = candidates.where((candidate) {
      return int.parse(candidate[pos]) == criteriaBit;
    }).toList();

    if (candidates.length == 1) {
      return toDecimal(candidates[0]);
    }
  }
  return 0;
}

int toDecimal(String input) {
  int result = 0;

  input.split('').forEach((ch) {
    result = result << 1;
    result += ch == '1' ? 1 : 0;
  });

  return result;
}

int mostCommonBit(List<String> input, int position) {
  int oneCount = 0;
  int zeroCount = 0;

  for (String line in input) {
    if (line[position] == '1') {
      oneCount++;
    } else {
      zeroCount++;
    }
  }
  return oneCount >= zeroCount ? 1 : 0;
}

int leastCommonBit(List<String> input, int position) {
  return mostCommonBit(input, position) == 1 ? 0 : 1;
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
