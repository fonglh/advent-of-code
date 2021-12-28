class Point3 {
  int x, y, z;

  Point3(this.x, this.y, this.z);

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
