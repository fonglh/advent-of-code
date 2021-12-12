import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  final input = await readFile("12.txt");

  Map<String, List<String>> passageMap = {};
  for (String line in input) {
    final List<String> path = line.split("-");

    if (passageMap.containsKey(path[0])) {
      passageMap[path[0]]!.add(path[1]);
      if (passageMap.containsKey(path[1])) {
        passageMap[path[1]]!.add(path[0]);
      } else {
        passageMap[path[1]] = [path[0]];
      }
    } else {
      passageMap[path[0]] = [path[1]];
      if (passageMap.containsKey(path[1])) {
        passageMap[path[1]]!.add(path[0]);
      } else {
        passageMap[path[1]] = [path[0]];
      }
    }
  }

  //print(part1(passageMap));
  print(part2(passageMap));
}

int part1(Map<String, List<String>> passageMap) {
  List<String> visited = [];
  List<List<String>> possiblePaths = [];
  List<String> currentPath = [];

  search(passageMap, visited, possiblePaths, currentPath, "start");

  return possiblePaths.length;
}

int part2(Map<String, List<String>> passageMap) {
  Map<String, int> visited = {};
  List<List<String>> possiblePaths = [];
  List<String> currentPath = [];

  search2(passageMap, visited, possiblePaths, currentPath, "start");

  return possiblePaths.length;
}

// https://www.baeldung.com/cs/simple-paths-between-two-vertices
void search(Map<String, List<String>> passageMap, List<String> visited,
    List<List<String>> possiblePaths, List<String> currentPath, String start) {
  // already visited, return
  if (visited.contains(start)) {
    return;
  }
  // mark this small cave as visited
  if (!isBigCave(start)) {
    visited.add(start);
  }
  currentPath.add(start);

  // Found the end. Add the path to the list of possible paths, backtrack.
  if (start == "end") {
    possiblePaths.add(List.from(currentPath));
    visited.remove(start);
    currentPath.removeLast();
    return;
  }

  // Recursively search all the connected caves
  for (String next in passageMap[start]!) {
    search(passageMap, visited, possiblePaths, currentPath, next);
  }
  currentPath.removeLast();
  visited.remove(start);
}

void search2(Map<String, List<String>> passageMap, Map<String, int> visited,
    List<List<String>> possiblePaths, List<String> currentPath, String start) {
  // cannot visit, already visited this small cave and another small cave has
  // been visited twice
  if (!canVisit(visited, start)) {
    return;
  }
  // mark this small cave as visited
  if (!isBigCave(start)) {
    if (visited.containsKey(start)) {
      visited[start] = visited[start]! + 1;
    } else {
      visited[start] = 1;
    }
  }
  currentPath.add(start);

  // Found the end. Add the path to the list of possible paths, backtrack.
  if (start == "end") {
    possiblePaths.add(List.from(currentPath));
    if (visited.containsKey(start) && visited[start] == 2) {
      visited[start] = visited[start]! - 1;
    } else {
      visited.remove(start);
    }
    currentPath.removeLast();
    return;
  }

  // Recursively search all the connected caves
  for (String next in passageMap[start]!) {
    search2(passageMap, visited, possiblePaths, currentPath, next);
  }
  currentPath.removeLast();
  if (visited.containsKey(start) && visited[start] == 2) {
    visited[start] = visited[start]! - 1;
  } else {
    visited.remove(start);
  }
}

bool isBigCave(String cave) {
  return cave.toUpperCase() == cave;
}

bool canVisit(Map<String, int> visited, String cave) {
  if (cave == "start" && visited.containsKey("start")) {
    return false;
  }

  return !visited.containsKey(cave) || visited.values.every((v) => v == 1);
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
