class Cuboid {
  late bool on;
  late int minX;
  late int maxX;
  late int minY;
  late int maxY;
  late int minZ;
  late int maxZ;

  Cuboid(this.on, this.minX, this.maxX, this.minY, this.maxY, this.minZ,
      this.maxZ);

  Cuboid.fromInput(String line) {
    var commandAndRanges = line.split(" ");
    var command = commandAndRanges[0];
    var ranges = commandAndRanges[1].split(",");

    var xRange =
        ranges[0].substring(2).split("..").map((e) => int.parse(e)).toList();
    var yRange =
        ranges[1].substring(2).split("..").map((e) => int.parse(e)).toList();
    var zRange =
        ranges[2].substring(2).split("..").map((e) => int.parse(e)).toList();

    minX = xRange[0];
    maxX = xRange[1];
    minY = yRange[0];
    maxY = yRange[1];
    minZ = zRange[0];
    maxZ = zRange[1];
    on = command == "on";
  }

  int volume() {
    int vol = (maxX - minX + 1) * (maxY - minY + 1) * (maxZ - minZ + 1);
    return on ? vol : -vol;
  }

  Cuboid? getIntersection(Cuboid otherCuboid) {
    int x1 = minX;
    if (minX < otherCuboid.minX) {
      x1 = otherCuboid.minX;
    }
    int x2 = maxX;
    if (maxX > otherCuboid.maxX) {
      x2 = otherCuboid.maxX;
    }

    int y1 = minY;
    if (minY < otherCuboid.minY) {
      y1 = otherCuboid.minY;
    }
    int y2 = maxY;
    if (maxY > otherCuboid.maxY) {
      y2 = otherCuboid.maxY;
    }

    int z1 = minZ;
    if (minZ < otherCuboid.minZ) {
      z1 = otherCuboid.minZ;
    }
    int z2 = maxZ;
    if (maxZ > otherCuboid.maxZ) {
      z2 = otherCuboid.maxZ;
    }

    if (x1 <= x2 && y1 <= y2 && z1 <= z2) {
      return Cuboid(false, x1, x2, y1, y2, z1, z2);
    }
  }

  @override
  String toString() {
    return "$on x=$minX..$maxX, y=$minY..$maxY, z=$minZ..$maxZ";
  }
}
