import 'dart:convert';
import 'dart:io';
import 'package:collection/collection.dart';

void main(List<String> arguments) async {
  var input = await readFile("25.txt");
  List<String> seafloor = [];
  List<String> nextSeafloor = [];

  // read seafloor init state from input
  for (var line in input) {
    seafloor.add(line);
  }

  Function listEq = ListEquality().equals;

  int step = 0;

  while (true) {
    nextSeafloor = takeStep(seafloor);
    step++;

    if (listEq(seafloor, nextSeafloor)) {
      break;
    } else {
      seafloor = nextSeafloor;
    }
  }
  print(step);
}

List<String> takeStep(List<String> seafloor) {
  List<String> nextSeafloor = [];

  nextSeafloor = moveEast(seafloor);
  nextSeafloor = moveSouth(nextSeafloor);

  return nextSeafloor;
}

List<String> moveEast(List<String> seafloor) {
  List<String> nextSeafloor = [];

  for (int i = 0; i < seafloor.length; i++) {
    String thisRow = "";
    for (int j = 0; j < seafloor[i].length; j++) {
      int nextColumn = (j + 1) % seafloor[i].length;

      // move east
      if (seafloor[i][j] == ">" && seafloor[i][nextColumn] == ".") {
        // this column is now empty
        thisRow += ".";

        // wrap around if the next column is 0
        if (nextColumn == 0) {
          thisRow = ">" + thisRow.substring(1);
        } else {
          // else append the sea cucumber to the end of the row
          thisRow += ">";
        }
        j += 1;
      } else {
        thisRow += seafloor[i][j];
      }
    }
    nextSeafloor.add(thisRow);
  }

  return nextSeafloor;
}

List<String> moveSouth(List<String> seafloor) {
  List<String> nextSeafloor = [];

  for (int i = 0; i < seafloor.length; i++) {
    String thisRow = "";
    for (int j = 0; j < seafloor[i].length; j++) {
      int prevRow = i - 1 >= 0 ? i - 1 : seafloor.length - 1;
      int nextRow = (i + 1) % seafloor.length;

      // sea cucumber from the previous row is moving down
      if (seafloor[prevRow][j] == "v" && seafloor[i][j] == ".") {
        thisRow += "v";
      } else if (seafloor[i][j] == "v" && seafloor[nextRow][j] == ".") {
        // sea cucumber from this row is moving down
        thisRow += ".";
      } else {
        thisRow += seafloor[i][j];
      }
    }
    nextSeafloor.add(thisRow);
  }

  return nextSeafloor;
}

void printSeafloor(List<String> seafloor) {
  for (var line in seafloor) {
    print(line);
  }
  print(
      "--------------------------------------------------------------------------------");
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
