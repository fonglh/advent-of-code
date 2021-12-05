# Advent of Code Day 05 Hydrothermal Venture

Remove diagonal points generation code in `line.dart` to get the answer to part 1.

Initial mistake was to generate diagonal points by finding minX and minY.
This wrongly removes the direction information of the line segment.