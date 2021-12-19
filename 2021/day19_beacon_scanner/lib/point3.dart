class Point3 {
  int x, y, z;

  Point3(this.x, this.y, this.z);

  // Return a list of the point rotated/oriented to all 24 possible orientations
  // https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
  // See the matrices at the bottom of the page, with the airplane diagrams.
  List<Point3> allTransformations() {
    return [
      Point3(x, y, z),
      Point3(-y, x, z),
      Point3(-x, -y, z),
      Point3(y, -x, z),
      Point3(-x, y, -z),
      Point3(y, x, -z),
      Point3(x, -y, -z),
      Point3(-y, -x, -z),
      Point3(x, -z, y),
      Point3(x, z, -y),
      Point3(-x, -z, -y),
      Point3(-x, z, y),
      Point3(z, x, y),
      Point3(-z, x, -y),
      Point3(z, -x, -y),
      Point3(-z, -x, y),
      Point3(-z, y, x),
      Point3(y, z, x),
      Point3(z, -y, x),
      Point3(-y, -z, x),
      Point3(-z, -y, -x),
      Point3(-y, z, -x),
      Point3(z, y, -x),
      Point3(y, -z, -x),
    ];
  }

  List<Point3> allTransformationsTranslated(Point3 translate) {
    return allTransformations()
        .map((element) => Point3(element.x + translate.x,
            element.y + translate.y, element.z + translate.z))
        .toList();
  }

  Point3 operator +(Point3 other) {
    return Point3(x + other.x, y + other.y, z + other.z);
  }

  Point3 operator -(Point3 other) {
    return Point3(x - other.x, y - other.y, z - other.z);
  }

  Point3 operator *(int scale) {
    return Point3(x * scale, y * scale, z * scale);
  }

  @override
  bool operator ==(other) {
    return (other is Point3) && other.x == x && other.y == y && other.z == z;
  }

  @override
  int get hashCode => Object.hash(x, y, z);

  @override
  String toString() {
    return "($x, $y, $z)";
  }
}
