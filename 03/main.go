package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	wires := readWires("03.txt")
	minCoord, maxCoord := findDimensions(wires)
	fmt.Println(minCoord, maxCoord)
	board := make([][]bool, abs(minCoord.Y)+abs(maxCoord.Y)+1)
	for row := range board {
		board[row] = make([]bool, abs(minCoord.X)+abs(maxCoord.X)+1)
	}
	fmt.Println(len(board[0]), len(board))

	var intersections []coordinate

	for i, wire := range wires {
		var currCoord coordinate
		//Offset X and Y coord
		currCoord.X = -minCoord.X
		currCoord.Y = -minCoord.Y
		fmt.Println("Origin: ", currCoord)
		wireSegments := strings.Split(wire, ",")
		for _, segment := range wireSegments {
			wireLength, _ := strconv.Atoi(segment[1:])
			switch segment[0] {
			case 'U':
				for dy := 1; dy <= wireLength; dy++ {
					coord := coordinate{currCoord.X, currCoord.Y + dy}
					if board[coord.Y][coord.X] && i == 1 {
						intersections = append(intersections, coord)
					}
					board[coord.Y][coord.X] = true
				}
				currCoord.Y += wireLength
			case 'D':
				for dy := 1; dy <= wireLength; dy++ {
					coord := coordinate{currCoord.X, currCoord.Y - dy}
					if board[coord.Y][coord.X] && i == 1 {
						intersections = append(intersections, coord)
					}
					board[coord.Y][coord.X] = true
				}
				currCoord.Y -= wireLength
			case 'L':
				for dx := 1; dx <= wireLength; dx++ {
					coord := coordinate{currCoord.X - dx, currCoord.Y}
					if board[coord.Y][coord.X] && i == 1 {
						intersections = append(intersections, coord)
					}
					board[coord.Y][coord.X] = true
				}
				currCoord.X -= wireLength
			case 'R':
				for dx := 1; dx <= wireLength; dx++ {
					coord := coordinate{currCoord.X + dx, currCoord.Y}
					if board[coord.Y][coord.X] && i == 1 {
						intersections = append(intersections, coord)
					}
					board[coord.Y][coord.X] = true
				}
				currCoord.X += wireLength
			}
		}
	}
	//printBoard(board)
	offsetIntersections(intersections, minCoord)
	fmt.Println(intersections)
	return minManhattanDistance(intersections)
}

func puzzle2() int {
	return -1
}

// Offset X and Y coords back
func offsetIntersections(intersections []coordinate, minCoord coordinate) {
	for i := range intersections {
		intersections[i].X += minCoord.X
		intersections[i].Y += minCoord.Y
	}
}

func minManhattanDistance(intersections []coordinate) int {
	// assumes that wires will definitely intersect
	firstCoord := intersections[0]
	minDist := abs(firstCoord.X) + abs(firstCoord.Y)
	for i := 1; i < len(intersections); i++ {
		dist := abs(intersections[i].X) + abs(intersections[i].Y)
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

// Find max dimensions of board from input
// Assume start at (0,0), return values are (minX, minY), (maxX, maxY)
func findDimensions(wires []string) (coordinate, coordinate) {
	var minCoord, maxCoord coordinate

	for _, wire := range wires {
		var currCoord coordinate
		wireSegments := strings.Split(wire, ",")
		for _, segment := range wireSegments {
			wireLength, _ := strconv.Atoi(segment[1:])
			switch segment[0] {
			case 'U':
				currCoord.Y += wireLength
			case 'D':
				currCoord.Y -= wireLength
			case 'L':
				currCoord.X -= wireLength
			case 'R':
				currCoord.X += wireLength
			}
			if currCoord.X > maxCoord.X {
				maxCoord.X = currCoord.X
			} else if currCoord.X < minCoord.X {
				minCoord.X = currCoord.X
			}
			if currCoord.Y > maxCoord.Y {
				maxCoord.Y = currCoord.Y
			} else if currCoord.Y < minCoord.Y {
				minCoord.Y = currCoord.Y
			}
		}
	}
	return minCoord, maxCoord
}

func readWires(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wires []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wires
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(board [][]bool) {
	// Print rows in reverse as +ve Y means up (print earlier)
	for rowNum := len(board) - 1; rowNum >= 0; rowNum-- {
		for colNum := range board[rowNum] {
			if board[rowNum][colNum] {
				fmt.Print("1 ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
