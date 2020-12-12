package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type nav struct {
	move      string
	magnitude int
}

func main() {
	input := readFile("12.txt")

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input []nav) int {
	ship := ship{heading: 90, xPos: 0, yPos: 0}

	for _, instruction := range input {
		ship.doNav(instruction)
	}

	return ship.manhattanDistance()
}

func puzzle2(input []nav) int {
	ship := ship{heading: 90, xPos: 0, yPos: 0}
	waypoint := waypoint{xPos: 10, yPos: 1}

	for _, instruction := range input {
		if instruction.move == "F" {
			ship.moveToWaypoint(waypoint, instruction.magnitude)
		} else {
			waypoint.doNav(instruction)
		}
	}
	return ship.manhattanDistance()
}

func readFile(filename string) []nav {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []nav

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		move := string(line[0])
		magnitude, _ := strconv.Atoi(line[1:])

		input = append(input, nav{move: move, magnitude: magnitude})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
