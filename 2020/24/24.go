package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
// Use cube coordinates to represent the hex grid
type Hex struct {
	x int
	y int
	z int
}

func main() {
	input := readFile("24.txt")

	//fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

// get cube coordinates from the reference tile after processing an input line
func getHex(line string) Hex {
	var xCoord, yCoord, zCoord int
	for i := 0; i < len(line); i += 1 {
		switch line[i] {
		case 'n':
			zCoord -= 1
			if line[i+1] == 'e' {
				xCoord += 1
			} else {
				yCoord += 1
			}
			i += 1
		case 's':
			zCoord += 1
			if line[i+1] == 'e' {
				yCoord -= 1
			} else {
				xCoord -= 1
			}
			i += 1
		case 'e':
			xCoord += 1
			yCoord -= 1
		case 'w':
			xCoord -= 1
			yCoord += 1
		}
	}
	return Hex{xCoord, yCoord, zCoord}
}

func countFloorBlack(floor map[Hex]bool) int {
	var blackCount int
	for _, v := range floor {
		if v {
			blackCount += 1
		}
	}
	return blackCount
}

func puzzle1(input []string) int {
	// false is white
	floor := make(map[Hex]bool)
	for _, line := range input {
		hex := getHex(line)
		floor[hex] = !floor[hex]
	}

	return countFloorBlack(floor)
}

func puzzle2(input []string) int {
	floor := make(map[Hex]bool)

	// Day 0, setup floor from input
	for _, line := range input {
		hex := getHex(line)
		floor[hex] = !floor[hex]
	}

	minHex, maxHex := searchFloor(floor)

	for day := 1; day <= 100; day += 1 {
		floorCopy := make(map[Hex]bool)

		minHex = Hex{minHex.x - 1, minHex.y - 1, minHex.z - 1}
		maxHex = Hex{maxHex.x + 1, maxHex.y + 1, maxHex.z + 1}
		for x := minHex.x; x <= maxHex.x; x += 1 {
			for y := minHex.y; y <= maxHex.y; y += 1 {
				for z := minHex.z; z <= maxHex.z; z += 1 {
					currHex := Hex{x, y, z}
					numBlackNeighbours := countNeighbourBlack(floor, currHex)
					if floor[currHex] && (numBlackNeighbours == 0 || numBlackNeighbours > 2) {
						floorCopy[currHex] = false
					} else if !floor[currHex] && numBlackNeighbours == 2 {
						floorCopy[currHex] = true
					} else {
						floorCopy[currHex] = floor[currHex]
					}
				}
			}
		}

		// reset floor and copy the new floor over
		floor = make(map[Hex]bool)
		for k, v := range floorCopy {
			floor[k] = v
		}
		fmt.Printf("Day %d: %d\n", day, countFloorBlack(floor))
	}

	return countFloorBlack(floor)
}

// return min and max x,y,z coordinates to search
func searchFloor(floor map[Hex]bool) (Hex, Hex) {
	minCoord := Hex{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	maxCoord := Hex{math.MinInt32, math.MinInt32, math.MinInt32}

	for k, _ := range floor {
		if k.x < minCoord.x {
			minCoord.x = k.x
		}
		if k.y < minCoord.y {
			minCoord.y = k.y
		}
		if k.z < minCoord.z {
			minCoord.z = k.z
		}

		if k.x > maxCoord.x {
			maxCoord.x = k.x
		}
		if k.y > maxCoord.y {
			maxCoord.y = k.y
		}
		if k.z > maxCoord.z {
			maxCoord.z = k.z
		}
	}

	minCoord = Hex{minCoord.x - 1, minCoord.y - 1, minCoord.z - 1}
	maxCoord = Hex{maxCoord.x + 1, maxCoord.y + 1, maxCoord.z + 1}

	return minCoord, maxCoord
}

func countNeighbourBlack(floor map[Hex]bool, hex Hex) int {
	var count int
	// start east, go anticlockwise
	// https://www.redblobgames.com/grids/hexagons/#neighbors-cube
	neighbours := [6]Hex{
		Hex{hex.x + 1, hex.y - 1, hex.z},
		Hex{hex.x + 1, hex.y, hex.z - 1},
		Hex{hex.x, hex.y + 1, hex.z - 1},
		Hex{hex.x - 1, hex.y + 1, hex.z},
		Hex{hex.x - 1, hex.y, hex.z + 1},
		Hex{hex.x, hex.y - 1, hex.z + 1},
	}

	for _, neighbour := range neighbours {
		if floor[neighbour] {
			count += 1
		}
	}

	return count
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
