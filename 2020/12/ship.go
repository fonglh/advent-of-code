package main

import (
	"math"
)

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

func (ship *ship) moveToWaypoint(w waypoint, magnitude int) {
	if w.xPos > 0 { // move east
		ship.moveShip(nav{move: "E", magnitude: w.xPos * magnitude})
	}
	if w.xPos < 0 { //move west
		ship.moveShip(nav{move: "W", magnitude: -w.xPos * magnitude})
	}
	if w.yPos > 0 { //move north
		ship.moveShip(nav{move: "N", magnitude: w.yPos * magnitude})
	}
	if w.yPos < 0 { //move south
		ship.moveShip(nav{move: "S", magnitude: -w.yPos * magnitude})
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
