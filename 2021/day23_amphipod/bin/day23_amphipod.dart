import 'package:day23_amphipod/burrow.dart';

void main(List<String> arguments) {
  // sample input burrow
  //Burrow burrow =
  //    Burrow.fromStrings("...........", ["BDDA", "CCBD", "BBAC", "DACA"]);
  Burrow burrow = Burrow.fromStrings("...........", ["BA", "CD", "BC", "DA"]);
  Burrow burrow1 = Burrow.fromStrings("...........", ["BA", "CD", "BC", "DA"]);

  // Actual input burrow
  //Burrow burrow =
  //    Burrow.fromStrings("...........", ["CDDB", "BCBC", "DBAA", "DACA"]);
  //print(burrow);

  Set<Burrow> uniqueBurrows = {};
  uniqueBurrows.add(burrow);
  uniqueBurrows.add(burrow1);
  print(uniqueBurrows.length);
  print(burrow == burrow1);
}
