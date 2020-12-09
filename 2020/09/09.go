package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile("09.txt")

	puzzle1Answer := puzzle1(input, 25)
	fmt.Println(puzzle1Answer)
	fmt.Println(puzzle2(input, puzzle1Answer))
}

func puzzle1(input []int, preambleLength int) int {
	var rangeStart int
	rangeEnd := preambleLength

	for ; rangeEnd < len(input); rangeEnd += 1 {
		if !validNumber(input[rangeStart:rangeEnd], input[rangeEnd]) {
			return input[rangeEnd]
		}
		rangeStart += 1
	}

	return -1
}

func validNumber(inputSlice []int, num int) bool {
	for i := 0; i < len(inputSlice); i++ {
		for j := i + 1; j < len(inputSlice); j++ {
			if i == j {
				continue
			}
			if inputSlice[i]+inputSlice[j] == num {
				return true
			}
		}
	}
	return false
}

func puzzle2(input []int, answer int) int {
	for i := 0; i < len(input); i += 1 {
		currSum := input[i]
		for j := i + 1; j < len(input); j += 1 {
			currSum += input[j]
			if currSum == answer {
				return addMinMax(input[i : j+1])
			}
		}
	}
	return 0
}

func addMinMax(inputSlice []int) int {
	var max int = inputSlice[0]
	var min int = inputSlice[0]
	for _, value := range inputSlice {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min + max
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
