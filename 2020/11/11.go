package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile("11.txt")

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input []string) int {
	// Deep copy slice or changing 1 slice will change the other as well
	currLayout := append(make([]string, 0, len(input)), input...)
	nextLayout := append(make([]string, 0, len(currLayout)), currLayout...)

	for {
		currLayout = append(make([]string, 0, len(nextLayout)), nextLayout...)
		nextLayout = append(make([]string, 0, len(currLayout)), currLayout...)

		for rowIndex, row := range currLayout {
			for colIndex, col := range row {
				adjacentOccupied := countOccupied(currLayout, rowIndex, colIndex)

				if col == 'L' && adjacentOccupied == 0 {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + "#" + nextLayout[rowIndex][colIndex+1:]
				} else if col == '#' && adjacentOccupied >= 4 {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + "L" + nextLayout[rowIndex][colIndex+1:]
				} else {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + string(col) + nextLayout[rowIndex][colIndex+1:]
				}
			}
		}

		if countTotalOccupied(currLayout) == countTotalOccupied(nextLayout) {
			return countTotalOccupied(nextLayout)
		}
	}
}

func countOccupied(seatLayout []string, rowIndex, colIndex int) int {
	var numOccupied int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (rowIndex+i) >= 0 && (rowIndex+i) < len(seatLayout) &&
				(colIndex+j) >= 0 && (colIndex+j) < len(seatLayout[rowIndex]) &&
				!(i == 0 && j == 0) {
				if seatLayout[rowIndex+i][colIndex+j] == '#' {
					numOccupied += 1
				}
			}
		}
	}

	return numOccupied
}

func countTotalOccupied(seatLayout []string) int {
	var numOccupied int

	for _, line := range seatLayout {
		numOccupied += strings.Count(line, "#")
	}

	return numOccupied
}

func puzzle2(input []string) int {
	// Deep copy slice or changing 1 slice will change the other as well
	currLayout := append(make([]string, 0, len(input)), input...)
	nextLayout := append(make([]string, 0, len(currLayout)), currLayout...)

	for {
		currLayout = append(make([]string, 0, len(nextLayout)), nextLayout...)
		nextLayout = append(make([]string, 0, len(currLayout)), currLayout...)

		for rowIndex, row := range currLayout {
			for colIndex, col := range row {
				adjacentOccupied := countOccupied2(currLayout, rowIndex, colIndex)

				if col == 'L' && adjacentOccupied == 0 {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + "#" + nextLayout[rowIndex][colIndex+1:]
				} else if col == '#' && adjacentOccupied >= 5 {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + "L" + nextLayout[rowIndex][colIndex+1:]
				} else {
					nextLayout[rowIndex] = nextLayout[rowIndex][:colIndex] + string(col) + nextLayout[rowIndex][colIndex+1:]
				}
			}
		}

		if countTotalOccupied(currLayout) == countTotalOccupied(nextLayout) {
			return countTotalOccupied(nextLayout)
		}
	}
}

func countOccupied2(seatLayout []string, rowIndex, colIndex int) int {
	var numOccupied int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			step := 1
			for {
				// multiply all the offsets for each direction by `step` and check if out of bounds
				if (rowIndex+i*step) >= 0 && (rowIndex+i*step) < len(seatLayout) &&
					(colIndex+j*step) >= 0 && (colIndex+j*step) < len(seatLayout[rowIndex]) &&
					!(i*step == 0 && j*step == 0) {

					// break out of the stepper loop when a seat (occupied or empty) is encountered
					if seatLayout[rowIndex+i*step][colIndex+j*step] == '#' {
						numOccupied += 1
						break
					} else if seatLayout[rowIndex+i*step][colIndex+j*step] == 'L' {
						break
					}
				} else { // stop stepping when out of bounds
					break
				}
				step += 1
			}
		}
	}

	return numOccupied
}

func printLayout(seatLayout []string) {
	for _, line := range seatLayout {
		fmt.Println(line)
	}
	fmt.Println("--------------------------------------------------------------------------------")
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
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
