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
	// init crab ring and node map
	crabRing := ring.New(len(input))
	nodeMap := make(map[int]*ring.Ring)
	for _, num := range input {
		crabRing.Value = num
		nodeMap[num] = crabRing
		crabRing = crabRing.Next()
	}

	for count := 0; count < 100; count += 1 {
		fmt.Println("move", count+1)
		fmt.Println("current", crabRing.Value)
		removed := crabRing.Unlink(3)

		// find destination candidate
		destCandidate := crabRing.Value.(int) - 1
		for {
			foundDestCandidate := true
			// hard code the length as Len() runs in O(n) time and the length is fixed anyway
			for i := 0; i < 3; i += 1 {
				if removed.Value.(int) == destCandidate {
					foundDestCandidate = false
				}
				removed = removed.Next()
			}
			if !foundDestCandidate || destCandidate == 0 {
				destCandidate -= 1
				if destCandidate <= 0 {
					destCandidate = 9
				}
			} else {
				break
			}
		}

		//find destination
		destination := nodeMap[destCandidate]

		fmt.Println("destination", destination.Value)
		destination.Link(removed)
		printRing(crabRing)
		crabRing = crabRing.Next()

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
	// init crab ring and nodemap
	crabRing := ring.New(1000000)
	nodeMap := make(map[int]*ring.Ring)
	for _, num := range input {
		crabRing.Value = num
		nodeMap[num] = crabRing
		crabRing = crabRing.Next()
	}

	for num := 10; num <= 1000000; num += 1 {
		crabRing.Value = num
		nodeMap[num] = crabRing
		crabRing = crabRing.Next()
	}

	for count := 0; count < 10000000; count += 1 {
		//if count%10000 == 0 {
		//	fmt.Println("move", count+1)
		//}
		removed := crabRing.Unlink(3)

		// find destination candidate
		destCandidate := crabRing.Value.(int) - 1
		for {
			foundDestCandidate := true
			// hard code the length as Len() runs in O(n) time and the length is fixed anyway
			for i := 0; i < 3; i += 1 {
				if removed.Value.(int) == destCandidate {
					foundDestCandidate = false
				}
				removed = removed.Move(1)
			}
			if !foundDestCandidate || destCandidate == 0 {
				destCandidate -= 1
				if destCandidate <= 0 {
					destCandidate = 1000000
				}
			} else {
				break
			}
		}

		//find destination
		destination := nodeMap[destCandidate]
		destination.Link(removed)
		crabRing = crabRing.Move(1)
	}

	product := 1
	firstCup := nodeMap[1]
	for i := 0; i < 2; i += 1 {
		firstCup = firstCup.Move(1)
		fmt.Println(firstCup.Value.(int))
		product *= firstCup.Value.(int)
	}

	return product
}
