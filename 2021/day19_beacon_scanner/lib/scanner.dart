import 'package:day19_beacon_scanner/point3.dart';
import 'dart:math';

class Scanner {
  List<Point3> beacons = [];
  Map<int, List<Point3>> beaconDistances = {};
  // whether beacons have been normalized to scanner 0
  bool normalized = false;
  // relative to scanner 0
  Point3? position;

  Scanner(List<String> beaconCoordinates) {
    for (var coord in beaconCoordinates) {
      var coordList = coord.split(",");
      Point3 beaconCoord = Point3(int.parse(coordList[0]),
          int.parse(coordList[1]), int.parse(coordList[2]));

      beacons.add(beaconCoord);
    }
    updateBeaconDistances();
  }

  // find pair-wise distances for all the beacons
  void updateBeaconDistances() {
    for (int i = 0; i < beacons.length; i++) {
      for (int j = i + 1; j < beacons.length; j++) {
        int distance = (pow(beacons[i].x - beacons[j].x, 2) +
                pow(beacons[i].y - beacons[j].y, 2) +
                pow(beacons[i].z - beacons[j].z, 2))
            .toInt();

        beaconDistances[distance] = [beacons[i], beacons[j]];
      }
    }
  }

  // find distances between beacons which are the same for this scanner
  // and the other scanner.
  List<int> getMatchingDistances(Scanner other) {
    List<int> matchingDistances = [];
    for (int d in beaconDistances.keys) {
      if (other.beaconDistances.containsKey(d)) {
        matchingDistances.add(d);
      }
    }

    return matchingDistances;
  }

  int? getTransformIndex(Scanner other) {
    List<int> matchingDistances = getMatchingDistances(other);

    // 12 beacons need to match, so that's 12 choose 2 = 66 distance pairs.
    if (matchingDistances.length < 66) return null;

    // Iterate through all the possible rotations/orientations.
    for (int transformIndex = 0; transformIndex < 24; transformIndex++) {
      bool allMatch = true;

      for (int dist in matchingDistances) {
        Point3 myVector = beaconDistances[dist]![0] - beaconDistances[dist]![1];
        // Orient the beacons from the other scanner and see if the direction matches
        // the direction of the pair of beacons from this scanner.
        Point3 otherVector = other.beaconDistances[dist]![0]
                .allTransformations()[transformIndex] -
            other.beaconDistances[dist]![1]
                .allTransformations()[transformIndex];

        // No match for either same or opposite direction
        if (myVector != otherVector && myVector != otherVector * -1) {
          allMatch = false;
        }
      }

      // With this rotation/orientation transformation index, all pairs of beacons
      // the same distance apart have the same direction.
      if (allMatch) {
        return transformIndex;
      }
    }
  }

  // Using the transformIndex returned by getTransformIndex,
  // rotate/orientate all the beacons in this scanner.
  void updateBeaconOrientations(int transformIndex) {
    List<Point3> newBeacons = [];

    for (Point3 beacon in beacons) {
      newBeacons.add(beacon.allTransformations()[transformIndex]);
    }

    beacons = newBeacons;
    updateBeaconDistances();
  }

  // Using the vector (represented by Point3) returned by getRelativePositionFromScanner0,
  // move all the beacons so they're relative to Scanner 0.
  // Beacon orientations MUST be updated first.
  void translateBeacons(Point3 relativePosition) {
    List<Point3> newBeacons = [];

    for (Point3 beacon in beacons) {
      newBeacons.add(beacon + relativePosition);
    }

    beacons = newBeacons;
    updateBeaconDistances();
  }

  // only call this when the other scanner is properly oriented and at least
  // 12 beacons overlap
  // use updateBeacons with getTransformIndex to update the other scanner's
  // orientation
  // Because this is called after the original scanner is normalized to scanner 0,
  // this actually returns the relative position from scanner 0.
  Point3 getRelativePositionFromScanner0(Scanner other) {
    List<int> matchingDistances = getMatchingDistances(other);

    for (int dist in matchingDistances) {
      Point3 myVector = beaconDistances[dist]![0] - beaconDistances[dist]![1];
      Point3 otherVector =
          other.beaconDistances[dist]![0] - other.beaconDistances[dist]![1];

      // same direction, assume there's something in the matching distances
      // that have vectors in the same direction so this code path is reached.
      // works with the test input.
      if (myVector == otherVector) {
        return beaconDistances[dist]![0] - other.beaconDistances[dist]![0];
      }
    }

    throw NullThrownError();
  }
}
