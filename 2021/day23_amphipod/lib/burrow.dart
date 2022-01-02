import 'package:collection/collection.dart';

class Burrow {
  List<int> hallway;
  List<List<int>> rooms;
  int cost;

  Burrow()
      : hallway = List<int>.generate(11, (_) => -1).toList(),
        // 1st element is the space adjacent to the hallway
        rooms = [
          [-1, -1, -1, -1],
          [-1, -1, -1, -1],
          [-1, -1, -1, -1],
          [-1, -1, -1, -1]
        ],
        cost = 0;

  Burrow.fromStrings(String hallway, List<String> rooms)
      : hallway = hallway
            .split('')
            .map((c) => c == '.' ? -1 : c.codeUnitAt(0) - 65)
            .toList(),
        rooms = rooms
            .map((r) => r
                .split('')
                .map((c) => c == '.' ? -1 : c.codeUnitAt(0) - 65)
                .toList())
            .toList(),
        cost = 0;

  @override
  bool operator ==(other) {
    if (other is Burrow) {
      return ListEquality().equals(hallway, other.hallway) &&
          DeepCollectionEquality().equals(rooms, other.rooms);
    }
    return false;
  }

  @override
  int get hashCode =>
      DeepCollectionEquality().hash(rooms) +
      DeepCollectionEquality().hash(hallway);

  @override
  String toString() {
    String result = "#############\n#";

    for (int hallSpace in hallway) {
      if (hallSpace == -1) {
        result += ".";
      } else {
        result += String.fromCharCode(65 + hallSpace);
      }
    }

    result += "#\n###";

    for (int i = 0; i < rooms[0].length; i++) {
      for (int j = 0; j < 4; j++) {
        if (rooms[j][i] == -1) {
          result += ".#";
        } else {
          result += String.fromCharCode(65 + rooms[j][i]) + "#";
        }
      }
      result += i == 0 ? "##\n  #" : "  \n  #";
    }
    result += "########\n";
    result += "cost: $cost";
    return result;
  }
}
