import 'dart:convert';
import 'dart:io';
import 'dart:math';

void main(List<String> arguments) async {
  var input = await readFile("20.txt");

  List<bool> algo = input[0].split("").map((e) => e == '#').toList();
  input = input.sublist(2);

  // init image
  Map<Point, bool> image = {};
  for (int i = 0; i < input.length; i++) {
    var linePixels = input[i].split("").map((e) => e == '#').toList();

    for (int j = 0; j < linePixels.length; j++) {
      if (linePixels[j]) {
        // Define points using x and y coordinates, instead of row and column
        image[Point(j, i)] = true;
      }
    }
  }

  // Frame the image exactly, this will expand by 1 pixel around the whole
  // perimeter for each step.
  Rectangle canvas = Rectangle.fromPoints(
      Point(0, 0), Point(input[0].length - 1, input.length - 1));
  bool background = false;

  // Part 1 is 2 steps
  for (int step = 0; step < 50; step++) {
    Map<Point, bool> nextImage = {};
    // Start processing outside the canvas because the border pixels will affect
    // the values for the new border.
    for (int y = canvas.topLeft.y.toInt() - 1;
        y <= canvas.bottomRight.y + 1;
        y++) {
      for (int x = canvas.topLeft.x.toInt() - 1;
          x <= canvas.bottomRight.x + 1;
          x++) {
        var currPixel = outputPixel(canvas, image, algo, x, y, background);
        if (currPixel) {
          nextImage[Point(x, y)] = true;
        }
      }
    }

    // expand canvas, update image
    canvas = Rectangle(canvas.topLeft.x - 1, canvas.topLeft.y - 1,
        canvas.width + 2, canvas.height + 2);
    image = Map<Point, bool>.from(nextImage);
    // update background by picking some point far outside the canvas
    background = outputPixel(canvas, image, algo, canvas.topLeft.x.toInt() - 10,
        canvas.topLeft.y.toInt(), background);
  }

  //printImage(canvas, image);
  print(image.length);
}

// print image for debugging
void printImage(Rectangle canvas, Map<Point, bool> image) {
  for (int y = canvas.topLeft.y.toInt(); y <= canvas.bottomRight.y; y++) {
    var line = "";
    for (int x = canvas.topLeft.x.toInt(); x <= canvas.bottomRight.x; x++) {
      line += image.containsKey(Point(x, y)) ? "#" : ".";
    }
    print(line);
  }
}

// Background refers to the colour of the "infinite" canvas.
// From the input, it alternates between dark and light, but this can be handled
// more generally.
bool outputPixel(Rectangle canvas, Map<Point, bool> image, List<bool> algo,
    int x, int y, bool background) {
  String lookupBitString = "";

  for (int j = y - 1; j <= y + 1; j++) {
    for (int i = x - 1; i <= x + 1; i++) {
      if (canvas.containsPoint(Point(i, j))) {
        lookupBitString += image.containsKey(Point(i, j)) ? "1" : "0";
      } else {
        lookupBitString += background ? "1" : "0";
      }
    }
  }

  return algo[int.parse(lookupBitString, radix: 2)];
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
