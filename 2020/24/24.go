package main

import (
	"bufio"
	"fmt"
	"log"
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

	fmt.Println(puzzle1(input))
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

func puzzle1(input []string) int {
	// false is white
	floor := make(map[Hex]bool)
	for _, line := range input {
		hex := getHex(line)
		floor[hex] = !floor[hex]
	}

	var blackCount int
	for _, v := range floor {
		if v {
			blackCount += 1
		}
	}
	return blackCount
}

func puzzle2(input []string) int {
	return 0
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
