class DeterministicDice {
  int value = 1;
  int rollCount = 0;

  DeterministicDice() {}

  int roll() {
    int result = value;
    value += 1;
    if (value > 100) value = 1;

    rollCount++;
    return result;
  }
}
