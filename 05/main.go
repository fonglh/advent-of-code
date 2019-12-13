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
	fmt.Println(puzzle2()) // answer is 12648139
}

func puzzle1() int {
	programInput := readFile("05.txt")
	userInput := 1
	return computer(programInput, userInput)
}

func puzzle2() int {
	programInput := readFile("05.txt")
	userInput := 5
	return computer(programInput, userInput)
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
		// jump if true
		case 5:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else {
				param1 = programInput[i+1]
			}
			//set instruction pointer to 2nd param, minus 2 from the for loop
			if param1 != 0 {
				if paramMode2 == 0 {
					i = programInput[programInput[i+2]] - 2
				} else {
					i = programInput[i+2] - 2
				}
			} else {
				i++
			}
		// jump if false
		case 6:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else {
				param1 = programInput[i+1]
			}
			//set instruction pointer to 2nd param, minus 2 from the for loop
			if param1 == 0 {
				if paramMode2 == 0 {
					i = programInput[programInput[i+2]] - 2

				} else {
					i = programInput[i+2] - 2
				}
			} else {
				i++
			}
		// less than
		case 7:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1, param2 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else {
				param1 = programInput[i+1]
			}
			if paramMode2 == 0 {
				param2 = programInput[programInput[i+2]]
			} else {
				param2 = programInput[i+2]
			}
			dest := programInput[i+3]
			if param1 < param2 {
				programInput[dest] = 1
			} else {
				programInput[dest] = 0
			}
			i += 2
		// equals
		case 8:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1, param2 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else {
				param1 = programInput[i+1]
			}
			if paramMode2 == 0 {
				param2 = programInput[programInput[i+2]]
			} else {
				param2 = programInput[i+2]
			}
			dest := programInput[i+3]
			if param1 == param2 {
				programInput[dest] = 1
			} else {
				programInput[dest] = 0
			}
			i += 2
		default:
			fmt.Println("no instruction", opcode)
		}
	}
	return output
}

// Read puzzle input of comma separated ints and parse into a slice of ints.
func readFile(filename string) []int {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Trim ending newline or the last instruction conversion will fail.
	inputStringSlice := strings.Split(strings.TrimSuffix(string(buffer), "\n"), ",")
	var input []int
	for _, i := range inputStringSlice {
		j, _ := strconv.Atoi(i)
		input = append(input, j)
	}
	return input
}
