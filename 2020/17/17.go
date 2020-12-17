package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Coordinate struct {
	x int
	y int
	z int
}

func main() {
	input := readFile("17.txt")

	fmt.Println(puzzle1(input))
	//fmt.Println(puzzle2(input))
}

func puzzle1(input []string) int {
	dimension := initDimension(input)

	for i := 0; i < 6; i++ {
		dimensionCopy := make(map[Coordinate]bool)

		minCoord, maxCoord := searchDimension(dimension)
		for x := minCoord.x; x <= maxCoord.x; x += 1 {
			for y := minCoord.y; y <= maxCoord.y; y += 1 {
				for z := minCoord.z; z <= maxCoord.z; z += 1 {
					currCoord := Coordinate{x, y, z}
					numNeighbours := countActiveNeighbours(dimension, currCoord)
					if dimension[currCoord] && (numNeighbours == 2 || numNeighbours == 3) {
						dimensionCopy[currCoord] = true
					} else if !dimension[currCoord] && numNeighbours == 3 {
						dimensionCopy[currCoord] = true
					} else {
						dimensionCopy[currCoord] = false
					}
				}
			}
		}

		// reset dimension and copy new state over
		dimension = make(map[Coordinate]bool)
		for k, v := range dimensionCopy {
			dimension[k] = v
		}
	}

	activeCount := 0
	for _, v := range dimension {
		if v {
			activeCount += 1
		}
	}

	return activeCount
}

func puzzle2(input []string) int {
	return 0
}

func initDimension(input []string) map[Coordinate]bool {
	dimension := make(map[Coordinate]bool)

	for y, row := range input {
		for x, cell := range row {
			if cell == '#' {
				dimension[Coordinate{x, y, 0}] = true
			}
		}
	}

	return dimension
}

func countActiveNeighbours(dimension map[Coordinate]bool, coord Coordinate) int {
	count := 0

	for dx := -1; dx <= 1; dx += 1 {
		for dy := -1; dy <= 1; dy += 1 {
			for dz := -1; dz <= 1; dz += 1 {
				if !(dx == 0 && dy == 0 && dz == 0) {
					checkCoord := Coordinate{coord.x + dx, coord.y + dy, coord.z + dz}
					if dimension[checkCoord] {
						count += 1
					}
				}
			}
		}
	}

	return count
}

// return min and max x,y,z coordinates to search
func searchDimension(dimension map[Coordinate]bool) (Coordinate, Coordinate) {
	minCoord := Coordinate{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	maxCoord := Coordinate{math.MinInt32, math.MinInt32, math.MinInt32}

	for k, _ := range dimension {
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

	minCoord = Coordinate{minCoord.x - 1, minCoord.y - 1, minCoord.z - 1}
	maxCoord = Coordinate{maxCoord.x + 1, maxCoord.y + 1, maxCoord.z + 1}

	return minCoord, maxCoord
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
