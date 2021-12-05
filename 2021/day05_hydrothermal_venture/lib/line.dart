import 'dart:math';

class Line {
  Point p1, p2;
  Line(this.p1, this.p2);

  @override
  String toString() {
    return '${p1.x},${p1.y} -> ${p2.x},${p2.y}';
  }

  bool isHorizontalOrVertical() {
    return isHorizontal() || isVertical();
  }

  bool isHorizontal() {
    return p1.y == p2.y;
  }

  bool isVertical() {
    return p1.x == p2.x;
  }

  List<Point> getPoints() {
    List<Point> points = [];

    if (isHorizontal()) {
      num minX = min(p1.x, p2.x);
      num maxX = max(p1.x, p2.x);
      for (num x = minX; x <= maxX; x++) {
        points.add(Point(x, p1.y));
      }
    } else if (isVertical()) {
      num minY = min(p1.y, p2.y);
      num maxY = max(p1.y, p2.y);
      for (num y = minY; y <= maxY; y++) {
        points.add(Point(p1.x, y));
      }
    } else {
      // diagonal lines
      if (p1.x < p2.x) {
        if (p1.y < p2.y) {
          // top left to bottom right
          for (num offset = 0; p1.x + offset <= p2.x; offset++) {
            points.add(Point(p1.x + offset, p1.y + offset));
          }
        } else {
          // bottom left to top right
          for (num offset = 0; p1.x + offset <= p2.x; offset++) {
            points.add(Point(p1.x + offset, p1.y - offset));
          }
        }
      } else {
        if (p1.y < p2.y) {
          // top right to bottom left
          for (num offset = 0; p1.x - offset >= p2.x; offset++) {
            points.add(Point(p1.x - offset, p1.y + offset));
          }
        } else {
          // bottom right to top left
          for (num offset = 0; p1.x - offset >= p2.x; offset++) {
            points.add(Point(p1.x - offset, p1.y - offset));
          }
        }
      }
    }

    return points;
  }
}
