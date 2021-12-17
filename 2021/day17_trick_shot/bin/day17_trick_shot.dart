import 'dart:math';

void main(List<String> arguments) {
  //Point targetMin = Point(20, -10);
  //Point targetMax = Point(30, -5);
  Point targetMin = Point(155, -117);
  Point targetMax = Point(182, -67);

  int validInitialVelocityCount = 0;

  for (int currX = 1; currX <= targetMax.x.toInt(); currX++) {
    for (int currY = targetMin.y.toInt();
        currY <= -targetMin.y.toInt();
        currY++) {
      Point velocity = Point(currX, currY);
      if (simulate(velocity, targetMin, targetMax)) {
        validInitialVelocityCount += 1;
      }
    }
  }

  print(validInitialVelocityCount);
}

// Return true if a step is in the target area, false if probe
// falls lower than lowest Y without a step in the target
bool simulate(Point velocity, Point targetMin, Point targetMax) {
  Point position = Point(0, 0);

  while (position.y.toInt() >= targetMin.y.toInt() &&
      !withinTarget(position, targetMin, targetMax)) {
    position = Point(position.x + velocity.x, position.y + velocity.y);

    int newVelocityX = 0;
    if (velocity.x < 0) {
      newVelocityX = velocity.x.toInt() + 1;
    } else if (velocity.x > 0) {
      newVelocityX = velocity.x.toInt() - 1;
    }
    velocity = Point(newVelocityX, velocity.y - 1);

    //print(position);
  }

  return withinTarget(position, targetMin, targetMax);
}

bool withinTarget(Point position, Point targetMin, Point targetMax) {
  return targetMin.x.toInt() <= position.x.toInt() &&
      position.x.toInt() <= targetMax.x.toInt() &&
      targetMin.y.toInt() <= position.y.toInt() &&
      position.y.toInt() <= targetMax.y.toInt();
}
