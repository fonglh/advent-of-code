package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type nav struct {
	move      string
	magnitude int
}

type ship struct {
	heading int
	xPos    int
	yPos    int
}

func (ship *ship) moveShip(nav nav) {
	// convert ship heading to direction
	if nav.move == "F" {
		switch ship.heading {
		case 0:
			nav.move = "N"
		case 90:
			nav.move = "E"
		case 180:
			nav.move = "S"
		case 270:
			nav.move = "W"
		}
	}

	switch nav.move {
	case "N":
		ship.yPos += nav.magnitude
	case "S":
		ship.yPos -= nav.magnitude
	case "E":
		ship.xPos += nav.magnitude
	case "W":
		ship.xPos -= nav.magnitude
	}
}

func (ship *ship) rotateShip(nav nav) {
	switch nav.move {
	case "L":
		ship.heading -= nav.magnitude
		if ship.heading < 0 {
			ship.heading += 360
		}
	case "R":
		ship.heading = (ship.heading + nav.magnitude) % 360
	}
}

func (ship *ship) doNav(nav nav) {
	switch nav.move {
	case "L", "R":
		ship.rotateShip(nav)
	default:
		ship.moveShip(nav)
	}
}

func (ship ship) manhattanDistance() int {
	return int(math.Abs(float64(ship.xPos)) + math.Abs(float64(ship.yPos)))
}

func main() {
	input := readFile("12.txt")

	fmt.Println(puzzle1(input))
	//fmt.Println(puzzle2(input))
}

func puzzle1(input []nav) int {
	ship := ship{heading: 90, xPos: 0, yPos: 0}

	for _, instruction := range input {
		ship.doNav(instruction)
	}

	return ship.manhattanDistance()
}

func puzzle2(input []nav) int {
	return 0
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
