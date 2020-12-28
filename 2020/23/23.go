package main

import (
	"container/ring"
	"fmt"
)

func main() {
	input := [9]int{7, 1, 6, 8, 9, 2, 5, 4, 3}
	//input = [9]int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input [9]int) int {
	// init crab ring
	crabRing := ring.New(len(input))
	for _, num := range input {
		crabRing.Value = num
		crabRing = crabRing.Next()
	}

	for count := 0; count < 100; count += 1 {
		fmt.Println("move", count+1)
		fmt.Println("current", crabRing.Value)
		removed := crabRing.Unlink(3)

		// find destination
		destCandidate := crabRing.Value.(int) - 1
		destination := crabRing.Move(1)
		foundDestination := false
		for !foundDestination {
			// hard code the length as Len() runs in O(n) time and the length is fixed anyway
			for i := 0; i < 6; i += 1 {
				if destination.Value.(int) == destCandidate {
					foundDestination = true
					break
				}
				destination = destination.Move(1)
			}
			destCandidate -= 1
			if destCandidate <= 0 {
				destCandidate = 9
			}
		}

		fmt.Println("destination", destination.Value)
		destination.Link(removed)
		printRing(crabRing)
		crabRing = crabRing.Move(1)

		fmt.Println("****************************")
	}

	return 0
}

func printRing(r *ring.Ring) {
	r.Do(func(x interface{}) {
		fmt.Print(x)
	})
	fmt.Println()
}

func puzzle2(input [9]int) int {
	return 0
}
