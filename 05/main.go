package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(puzzle1()) // answer is 4511442
	//	fmt.Println(puzzle2())
}

func puzzle1() int {
	programInput := readFile()
	userInput := 1
	return computer(programInput, userInput)
}

func puzzle2() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := readFile()
			result := computer(input, 5)
			if result == 19690720 {
				return noun*100 + verb
			}
		}
	}
	// made up error code
	return -1
}

func computer(programInput []int, userInput int) int {
	var output int
	// advance instruction pointer 2 places by default, some instructions need more
	for i := 0; i < len(programInput); i += 2 {
		opcode := programInput[i] % 100
		if opcode == 99 {
			fmt.Println("done")
			break
		}
		switch opcode {
		// addition
		case 1:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var operandOne, operandTwo int
			if paramMode1 == 0 {
				operandOne = programInput[programInput[i+1]]
			} else {
				operandOne = programInput[i+1]
			}
			if paramMode2 == 0 {
				operandTwo = programInput[programInput[i+2]]
			} else {
				operandTwo = programInput[i+2]
			}
			dest := programInput[i+3] // position mode
			programInput[dest] = operandOne + operandTwo
			i += 2
		// multiplication
		case 2:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var operandOne, operandTwo int
			if paramMode1 == 0 {
				operandOne = programInput[programInput[i+1]]
			} else {
				operandOne = programInput[i+1]
			}
			if paramMode2 == 0 {
				operandTwo = programInput[programInput[i+2]]
			} else {
				operandTwo = programInput[i+2]
			}
			dest := programInput[i+3]
			programInput[dest] = operandOne * operandTwo
			i += 2
		// programInput
		case 3:
			dest := programInput[i+1]
			programInput[dest] = userInput
		// output
		case 4:
			paramMode1 := (programInput[i] / 100) % 10
			if paramMode1 == 0 {
				output = programInput[programInput[i+1]]
			} else {
				output = programInput[i+1]
			}
			fmt.Println("output", output, "IP", i)
		}
	}
	return output
}

// Read puzzle input of comma separated ints and parse into a slice of ints.
func readFile() []int {
	buffer, err := ioutil.ReadFile("05.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputStringSlice := strings.Split(string(buffer), ",")
	var input []int
	for _, i := range inputStringSlice {
		j, _ := strconv.Atoi(i)
		input = append(input, j)
	}
	return input
}
