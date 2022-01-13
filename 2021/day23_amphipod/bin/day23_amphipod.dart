import 'package:day23_amphipod/burrow.dart';
import 'package:collection/collection.dart';

void main(List<String> arguments) {
  // sample input burrow
  // Burrow burrow = Burrow.fromStrings("...........", ["BA", "CD", "BC", "DA"]);
  //Burrow burrow =
  //    Burrow.fromStrings("...........", ["BDDA", "CCBD", "BBAC", "DACA"]);

  // Actual input burrow
  //Burrow burrow = Burrow.fromStrings("...........", ["CB", "BC", "DA", "DA"]);
  Burrow burrow =
      Burrow.fromStrings("...........", ["CDDB", "BCBC", "DBAA", "DACA"]);

  print(minCost(burrow));
}

int minCost(Burrow burrow) {
  Set<Burrow> visited = {};
  Map<Burrow, int> costs = {};
  Map<Burrow, Burrow> prev = {};

  final queue = PriorityQueue<Burrow>((a, b) {
    return a.cost.compareTo(b.cost);
  });

  visited.add(burrow);
  queue.add(burrow);
  costs[burrow] = 0;

  while (queue.isNotEmpty) {
    Burrow current = queue.removeFirst();
    if (current.isComplete()) {
      viewMinCostPath(prev, current);
      return current.cost;
    }

    visited.add(current);

    List<Burrow> nextBurrows = current.nextBurrows();
    for (int i = 0; i < nextBurrows.length; i++) {
      if (!visited.contains(nextBurrows[i])) {
        if (nextBurrows[i].cost < (costs[nextBurrows[i]] ?? 9999999)) {
          costs.remove(nextBurrows[i]);
          costs[nextBurrows[i]] = nextBurrows[i].cost;
          queue.add(nextBurrows[i]);

          prev.remove(nextBurrows[i]);
          prev[nextBurrows[i]] = current;
        }
      }
    }
  }

  return 0;
}

void viewMinCostPath(Map<Burrow, Burrow> prev, Burrow destination) {
  Burrow? currPathBurrow = destination;
  while (currPathBurrow != null) {
    print(currPathBurrow);
    print("----------------------------------------");
    currPathBurrow = prev[currPathBurrow];
  }
}
