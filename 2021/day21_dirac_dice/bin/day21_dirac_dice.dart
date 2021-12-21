import '../lib/deterministic_dice.dart';
import 'dart:math';

void main(List<String> arguments) {
  // puzzle input
  int p1Pos = 6;
  int p2Pos = 10;

  // test input
  // p1Pos = 4;
  // p2Pos = 8;

//  print(part1(p1Pos, p2Pos));
  Point result = Point(0, 0);

  // p2 goes 1st with a dummy roll of 0 value, so the actual game has p1 going first.
  result = diracGame(p1Pos, p2Pos, 0, 0, 0, false);
  if (result.x > result.y) {
    print(result.x);
  } else {
    print(result.y);
  }
}

Point diracGame(
    int p1Pos, int p2Pos, int p1Score, int p2Score, int roll, bool p1Turn) {
  //print("player 1: $p1Pos, $p1Score");
  //print("player 2: $p2Pos, $p2Score");
  //print("roll: $roll, p1Turn: $p1Turn");
  // move players, check for win
  if (roll > 0) {
    if (p1Turn) {
      p1Pos = movePlayer(p1Pos, roll);
      p1Score += p1Pos;

      if (p1Score >= 21) {
        return Point(1, 0);
      }
    } else {
      p2Pos = movePlayer(p2Pos, roll);
      p2Score += p2Pos;

      if (p2Score >= 21) {
        return Point(0, 1);
      }
    }
  }

  // roll amounts, and the number of times they appear
  // this optimization, instead of a recursive call for every roll result,
  // is necessary for the program to finish running
  List<Point> allRolls = [
    Point(3, 1),
    Point(4, 3),
    Point(5, 6),
    Point(6, 7),
    Point(7, 6),
    Point(8, 3),
    Point(9, 1)
  ];

  //print("-----------trying all rolls ------------");
  Point result = Point(0, 0);
  // try all possible outcomes
  for (Point roll in allRolls) {
    Point rollResult =
        diracGame(p1Pos, p2Pos, p1Score, p2Score, roll.x.toInt(), !p1Turn);
    result += Point(rollResult.x.toInt() * roll.y.toInt(),
        rollResult.y.toInt() * roll.y.toInt());
  }

  return result;
}

int part1(int p1Pos, int p2Pos) {
  int p1Score = 0;
  int p2Score = 0;

  DeterministicDice dice = DeterministicDice();

  while (p1Score <= 1000 && p2Score <= 1000) {
    int p1Move = dice.roll() + dice.roll() + dice.roll();
    p1Pos = movePlayer(p1Pos, p1Move);
    p1Score += p1Pos;

    //print("player 1: $p1Pos, $p1Score");

    if (p1Score >= 1000) break;

    int p2Move = dice.roll() + dice.roll() + dice.roll();
    p2Pos = movePlayer(p2Pos, p2Move);
    p2Score += p2Pos;

    // print("player 2: $p2Pos, $p2Score");
  }

  if (p1Score >= 1000) {
    print("player 1 won");
    return p2Score * dice.rollCount;
  } else {
    print("player 2 won");
    return p1Score * dice.rollCount;
  }
}

int movePlayer(int currPos, int amount) {
  return ((currPos - 1) + amount) % 10 + 1;
}
