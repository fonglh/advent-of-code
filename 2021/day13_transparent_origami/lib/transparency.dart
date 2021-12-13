class Transparency {
  List<List<bool>> transparency = [];

  Transparency(List<String> dotCoordinates) {
    int maxX = 0, maxY = 0;

    for (String coordinate in dotCoordinates) {
      var coords = coordinate.split(",");
      int x = int.parse(coords[0]);
      int y = int.parse(coords[1]);

      if (x > maxX) {
        maxX = x;
      }
      if (y > maxY) {
        maxY = y;
      }
    }
    //print(maxX);
    //print(maxY);

    for (int row = 0; row <= maxY; row++) {
      transparency.add(List.filled(maxX + 1, false));
    }

    for (String coordinate in dotCoordinates) {
      var coords = coordinate.split(",");
      int x = int.parse(coords[0]);
      int y = int.parse(coords[1]);
      transparency[y][x] = true;
    }
  }

  void foldAlongY(int yCoord) {
    for (int offset = 1;
        yCoord - offset >= 0 && yCoord + offset < transparency.length;
        offset++) {
      for (int col = 0; col < transparency[0].length; col++) {
        transparency[yCoord - offset][col] = transparency[yCoord - offset]
                [col] |
            transparency[yCoord + offset][col];
      }
    }

    transparency = transparency.sublist(0, yCoord);
  }

  void foldAlongX(int xCoord) {
    for (int offset = 1;
        xCoord - offset >= 0 && xCoord + offset < transparency[0].length;
        offset++) {
      for (int row = 0; row < transparency.length; row++) {
        transparency[row][xCoord - offset] = transparency[row]
                [xCoord - offset] |
            transparency[row][xCoord + offset];
      }
    }

    transparency = transparency.map((row) => row.sublist(0, xCoord)).toList();
  }

  int numDots() {
    return transparency
        .map((row) => row.where((e) => e).length)
        .reduce((value, element) => value + element);
  }

  @override
  String toString() {
    return transparency
        .map((row) => row.map((e) => e ? "#" : ".").join())
        .join("\n");
  }
}
