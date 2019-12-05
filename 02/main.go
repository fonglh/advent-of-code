package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	input := readFile(12, 2)
	return computer(input)
}

func puzzle2() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := readFile(noun, verb)
			result := computer(input)
			if result == 19690720 {
				return noun*100 + verb
			}
		}
	}
	// made up error code
	return -1
}

func computer(input []int) int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		if opcode == 99 {
			break
		}
		operandOne := input[input[i+1]]
		operandTwo := input[input[i+2]]
		dest := input[i+3]
		switch opcode {
		case 1:
			input[dest] = operandOne + operandTwo
		case 2:
			input[dest] = operandOne * operandTwo
		}
	}
	return input[0]
}

// Noun is placed in address 1, verb is placed in address 2, as explained in part 2 of the puzzle.
// Read puzzle input of comma separated ints and parse into a slice of ints.
func readFile(noun, verb int) []int {
	buffer, err := ioutil.ReadFile("02.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputStringSlice := strings.Split(string(buffer), ",")
	var input []int
	for _, i := range inputStringSlice {
		j, _ := strconv.Atoi(i)
		input = append(input, j)
	}
	// Program init as defined in puzzle
	input[1] = noun
	input[2] = verb
	return input
}
