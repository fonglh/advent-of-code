package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(puzzle1(171309, 643603))
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

func checkDouble(input int) bool {
	hasDouble := false
	inputString := strconv.Itoa(input)
	for i := 1; i < len(inputString); i++ {
		if inputString[i] == inputString[i-1] {
			hasDouble = true
			break
		}
	}
	return hasDouble
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
