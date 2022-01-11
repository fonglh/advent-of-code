import 'package:collection/collection.dart';
import 'dart:math';

class Burrow {
  List<int> hallway;
  List<List<int>> rooms;
  int cost;

  // The index of the hallway that opens to each room
  List<int> roomPositions = [2, 4, 6, 8];

  Burrow(int roomSize)
      : hallway = List<int>.generate(11, (_) => -1).toList(),
        // 1st element is the space adjacent to the hallway
        rooms = [
          List<int>.generate(roomSize, (_) => -1).toList(),
          List<int>.generate(roomSize, (_) => -1).toList(),
          List<int>.generate(roomSize, (_) => -1).toList(),
          List<int>.generate(roomSize, (_) => -1).toList(),
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

  Burrow clone() {
    Burrow b = Burrow(rooms[0].length);

    b.hallway = List.from(hallway);
    for (int i = 0; i < rooms.length; i++) {
      for (int j = 0; j < rooms[i].length; j++) {
        b.rooms[i][j] = rooms[i][j];
      }
    }
    b.cost = cost;
    return b;
  }

  List<Burrow> nextBurrows() {
    List<Burrow> result = moveHallwayToRoom();
    for (int i = 0; i < 4; i++) {
      result.addAll(moveRoomToHallway(i));
    }
    return result;
  }

  // Returns a list of all possible burrows when an amphipod is moved
  // from the hallway to a room.
  List<Burrow> moveHallwayToRoom() {
    List<Burrow> results = [];

    for (int i = 0; i < hallway.length; i++) {
      // found an amphipod in the hallway
      if (hallway[i] != -1) {
        int amphipod = hallway[i];
        if (isRoomAvailable(amphipod) && isHallwayClear(i)) {
          // Create new burrow with amphipod moved to room
          Burrow newBurrow = clone();
          int destinationIdxInRoom = rooms[amphipod].lastIndexOf(-1);
          newBurrow.hallway[i] = -1;
          newBurrow.rooms[amphipod][destinationIdxInRoom] = amphipod;
          newBurrow.cost += (((i - roomPositions[amphipod]).abs() +
                      destinationIdxInRoom +
                      1) *
                  // amphipod type number corresponds to the cost exponent
                  pow(10, amphipod))
              .toInt();
          results.add(newBurrow);
        }
      }
    }

    return results;
  }

  // Check if there's a path from an amphipod's position in the hallway
  // to its room.
  // All the positions in the hallway between the amphipod and its room
  // should be -1.
  bool isHallwayClear(int amphipodPosition) {
    int roomPosition = roomPositions[hallway[amphipodPosition]];

    if (roomPosition < amphipodPosition) {
      // Space at roomPosition is always empty
      return hallway
          .sublist(roomPosition, amphipodPosition)
          .every((e) => e == -1);
    } else {
      return hallway
          .sublist(amphipodPosition + 1, roomPosition)
          .every((e) => e == -1);
    }
  }

  // Amphipod can only move to its destination room if that room only
  // contains amphipods of the same type.
  // Amphipods should move all the way into the room, so should not need to
  // consider if an amphipod is in the middle of the room.
  // This returns true if the room is full, but since it's meant to be called
  // when checking the hallway, it shouldn't matter since there won't be any
  // amphipods of this type in the hallway if the room is full.
  bool isRoomAvailable(int roomIdx) {
    return rooms[roomIdx]
        .every((element) => element == roomIdx || element == -1);
  }

  // Returns all possible burrows when an amphipod is moved from the given
  // room to the hallway.
  List<Burrow> moveRoomToHallway(int roomIdx) {
    List<Burrow> results = [];

    // find amphipod nearest the hallway
    int amphipodIdx = rooms[roomIdx].indexWhere((e) => e != -1);

    // no amphipod in room, so no new burrows from this room with a move
    // to the hallway
    if (amphipodIdx == -1) {
      return results;
    }

    // Create new burrows with the amphipod moved to all possible hallway spaces

    // Search left of the room
    for (int i = roomPositions[roomIdx] - 1; i >= 0; i--) {
      if (roomPositions.contains(i)) {
        // outside a room, can't stop here
        continue;
      }
      if (hallway[i] != -1) {
        // blocked by another amphipod, stop searching
        break;
      }

      if (hallway[i] == -1) {
        // Create a new burrow with the amphipod moved to the hallway
        Burrow newBurrow = clone();
        newBurrow.hallway[i] = rooms[roomIdx][amphipodIdx];
        newBurrow.rooms[roomIdx][amphipodIdx] = -1;
        newBurrow.cost +=
            (((i - roomPositions[roomIdx]).abs() + amphipodIdx + 1) *
                    // amphipod type number corresponds to the cost exponent
                    pow(10, rooms[roomIdx][amphipodIdx]))
                .toInt();
        results.add(newBurrow);
      }
    }

    // Search right of the room
    for (int i = roomPositions[roomIdx] + 1; i < hallway.length; i++) {
      if (roomPositions.contains(i)) {
        // outside a room, can't stop here
        continue;
      }
      if (hallway[i] != -1) {
        // blocked by another amphipod, stop searching
        break;
      }

      if (hallway[i] == -1) {
        // Create a new burrow with the amphipod moved to the hallway
        Burrow newBurrow = clone();
        newBurrow.hallway[i] = rooms[roomIdx][amphipodIdx];
        newBurrow.rooms[roomIdx][amphipodIdx] = -1;
        newBurrow.cost +=
            (((i - roomPositions[roomIdx]).abs() + amphipodIdx + 1) *
                    // amphipod type number corresponds to the cost exponent
                    pow(10, rooms[roomIdx][amphipodIdx]))
                .toInt();
        results.add(newBurrow);
      }
    }

    return results;
  }

  bool isComplete() {
    bool result = hallway.every((space) => space == -1);

    for (int i = 0; result && i < rooms.length; i++) {
      result &= rooms[i].every((space) => space == i);
    }
    return result;
  }

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
