package main

import "fmt"

func main() {
	input := []int{14, 8, 16, 0, 1, 17}
	//input := []int{1, 3, 2}

	//fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input []int) int {
	spoken := input

	for i := len(spoken); i < 2020; i += 1 {
		lastIndex := len(spoken) - 1
		searchIndex := search(spoken[:lastIndex], spoken[lastIndex])
		if searchIndex == -1 {
			// first time the last number has been spoken, so append 0
			spoken = append(spoken, 0)
		} else {
			spoken = append(spoken, lastIndex-searchIndex)
		}
	}

	return spoken[len(spoken)-1]
}

func search(arr []int, item int) int {
	maxIndex := -1

	for index, arrItem := range arr {
		if item == arrItem {
			maxIndex = index
		}
	}
	return maxIndex
}

func puzzle2(input []int) int {
	// Map number to the turn number. Turn numbers start from 1
	spoken := make(map[int]int)
	lastNumber := input[len(input)-1]

	// init map
	for index := 0; index < len(input)-1; index += 1 {
		spoken[input[index]] = index + 1 // remember turn numbers start from 1
	}

	for turn := len(input) + 1; turn <= 30000000; turn += 1 {
		// do the lookup first, then add the previous turn's number to the map
		index, ok := spoken[lastNumber]
		spoken[lastNumber] = turn - 1

		if ok {
			// found, update current turn's number to the difference between prev turn number and index
			// index is the turn where it was most recently spoken before the previous turn
			lastNumber = turn - 1 - index
		} else {
			// not found, current turn's number is 0
			lastNumber = 0
		}
		// don't add to the map yet. the lookup needs to be done first.
	}

	return lastNumber
}
