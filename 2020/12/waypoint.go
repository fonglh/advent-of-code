package main

type waypoint struct {
	// relative position from ship
	// i.e. treat ship as origin
	xPos int
	yPos int
}

func (waypoint *waypoint) moveWaypoint(nav nav) {
	switch nav.move {
	case "N":
		waypoint.yPos += nav.magnitude
	case "S":
		waypoint.yPos -= nav.magnitude
	case "E":
		waypoint.xPos += nav.magnitude
	case "W":
		waypoint.xPos -= nav.magnitude
	}
}

func (waypoint *waypoint) rotateWaypoint(navInstruction nav) {
	if navInstruction.magnitude == 0 {
		return
	}
	if navInstruction.move == "L" {
		waypoint.xPos, waypoint.yPos = -waypoint.yPos, waypoint.xPos
	} else {
		waypoint.xPos, waypoint.yPos = waypoint.yPos, -waypoint.xPos
	}
	waypoint.rotateWaypoint(nav{move: navInstruction.move, magnitude: navInstruction.magnitude - 90})
}

func (waypoint *waypoint) doNav(nav nav) {
	switch nav.move {
	case "L", "R":
		waypoint.rotateWaypoint(nav)
	default:
		waypoint.moveWaypoint(nav)
	}
}
