import '../lib/deterministic_dice.dart';

void main(List<String> arguments) {
  // puzzle input
  int p1Pos = 6;
  int p2Pos = 10;

  // test input
  //p1Pos = 4;
  //p2Pos = 8;

  int p1Score = 0;
  int p2Score = 0;

  DeterministicDice dice = DeterministicDice();

  while (p1Score <= 1000 && p2Score <= 1000) {
    int p1Move = dice.roll() + dice.roll() + dice.roll();
    p1Pos = movePlayer(p1Pos, p1Move);
    p1Score += p1Pos;

    print("player 1: $p1Pos, $p1Score");

    if (p1Score >= 1000) break;

    int p2Move = dice.roll() + dice.roll() + dice.roll();
    p2Pos = movePlayer(p2Pos, p2Move);
    p2Score += p2Pos;

    print("player 2: $p2Pos, $p2Score");
  }

  // player 1 won
  if (p1Score >= 1000) {
    print("player 1 won");
    print(p2Score * dice.rollCount);
  } else {
    print("player 2 won");
    print(p1Score * dice.rollCount);
  }
}

int movePlayer(int currPos, int amount) {
  return ((currPos - 1) + amount) % 10 + 1;
}
