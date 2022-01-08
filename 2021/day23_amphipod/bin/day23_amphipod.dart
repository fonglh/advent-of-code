import 'package:day23_amphipod/burrow.dart';

void main(List<String> arguments) {
  // sample input burrow
  //Burrow burrow =
  //    Burrow.fromStrings("...........", ["BDDA", "CCBD", "BBAC", "DACA"]);
  Burrow burrow = Burrow.fromStrings("...........", ["BA", "CD", "BC", "DA"]);
  Burrow burrow1 = Burrow.fromStrings(".....D...A.", [".A", "BB", "CC", ".D"]);

  // Actual input burrow
  //Burrow burrow =
  //    Burrow.fromStrings("...........", ["CDDB", "BCBC", "DBAA", "DACA"]);
  //print(burrow);

  Set<Burrow> uniqueBurrows = {};
  print(uniqueBurrows.length);

  var newBurrows = burrow1.moveHallwayToRoom();
  print(newBurrows.length);
  for (int i = 0; i < newBurrows.length; i++) {
    print(newBurrows[i]);
    print("----------------------------------------");
  }
}
