package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(puzzle1(171309, 643603))
	fmt.Println(puzzle2(171309, 643603))
}

func puzzle1(minimum, maximum int) int {
	var validCount int
	for i := minimum; i <= maximum; i++ {
		if checkDouble(i) && checkDigitIncrease(i) {
			validCount++
		}
	}
	return validCount
}

func puzzle2(minimum, maximum int) int {
	validCount := 0
	for i := minimum; i <= maximum; i++ {
		if checkDouble(i) && checkDigitIncrease(i) && checkDoubleOnly(i) {
			validCount++
		}
	}
	return validCount
}

func checkDouble(input int) bool {
	inputString := strconv.Itoa(input)
	for i := 1; i < len(inputString); i++ {
		if inputString[i] == inputString[i-1] {
			return true
		}
	}
	return false
}

func checkDigitIncrease(input int) bool {
	inputString := strconv.Itoa(input)
	for i := 1; i < len(inputString); i++ {
		if inputString[i] < inputString[i-1] {
			return false
		}
	}
	return true
}

// By this time, only numbers with doubles and increasing digits will enter this function.
// So index positions for a given digit will be continuous.
func checkDoubleOnly(input int) bool {
	inputString := strconv.Itoa(input)
	digitMap := make(map[rune][]int)
	for i, digit := range inputString {
		digitMap[digit] = append(digitMap[digit], i)
	}
	for _, positions := range digitMap {
		// checkDigitIncrease is called before this so the digits are already monotonically increasing
		// and the same digit will be grouped together.
		// Thus a valid double not part of a larger group should only have 2 index positions.
		if len(positions) == 2 {
			return true
		}
	}
	return false
}
