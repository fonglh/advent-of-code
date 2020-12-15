package main

import "fmt"

func main() {
	input := []int{14, 8, 16, 0, 1, 17}
	//input := []int{0, 3, 6}

	fmt.Println(puzzle1(input))
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
