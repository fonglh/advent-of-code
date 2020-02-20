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
	//fmt.Println(puzzle2())
}

func puzzle1() int {
	programInput := readFile("09-1.txt")
	output := computer(programInput, []int{1})
	return output
}

func puzzle2() int {
	programInput := readFile("09.txt")
	output := computer(programInput, []int{1})
	return output
}

func computer(programInput []int, userInputs []int) int {
	var output int
	var userInputIdx int
	var relativeBase int
	programLength := len(programInput)
	// Hack for large memory
	programInput = append(programInput, make([]int, 10000)...)
	// advance instruction pointer 2 places by default, some instructions need more
	for i := 0; i < programLength; i += 2 {
		opcode := programInput[i] % 100
		fmt.Println()
		fmt.Println("IP", i, "opcode", opcode)
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
			} else if paramMode1 == 1{
				operandOne = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				operandOne = programInput[relativeBase+relativeBaseOffset]
			}
			if paramMode2 == 0 {
				operandTwo = programInput[programInput[i+2]]
			} else if paramMode2 == 1{
				operandTwo = programInput[i+2]
			} else {
				relativeBaseOffset := programInput[i+2]
				operandTwo = programInput[relativeBase+relativeBaseOffset]
			}
			dest := programInput[i+3] // position mode
			programInput[dest] = operandOne + operandTwo
			fmt.Println("paramMode1", paramMode1, "paramMode2", paramMode2, "op1", operandOne, "op2", operandTwo, "Dest", dest)
			fmt.Println(programInput[dest])
			i += 2
		// multiplication
		case 2:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var operandOne, operandTwo int
			if paramMode1 == 0 {
				operandOne = programInput[programInput[i+1]]
			} else if paramMode1 == 1{
				operandOne = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				operandOne = programInput[relativeBase+relativeBaseOffset]
			}
			if paramMode2 == 0 {
				operandTwo = programInput[programInput[i+2]]
			} else if paramMode2 == 1{
				operandTwo = programInput[i+2]
			} else {
				relativeBaseOffset := programInput[i+2]
				operandTwo = programInput[relativeBase+relativeBaseOffset]
			}
			dest := programInput[i+3]
			programInput[dest] = operandOne * operandTwo
			i += 2
		// programInput
		case 3:
			dest := programInput[i+1]
			programInput[dest] = userInputs[userInputIdx]
			userInputIdx++
		// output
		case 4:
			paramMode1 := (programInput[i] / 100) % 10
			if paramMode1 == 0 {
				output = programInput[programInput[i+1]]
			} else if paramMode1 == 1{
				output = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				output = programInput[relativeBase+relativeBaseOffset]
			}
			fmt.Println("output", output, "IP", i)
		// jump if true
		case 5:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else if paramMode1 == 1{
				param1 = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				param1 = programInput[relativeBase + relativeBaseOffset]
			}
			//set instruction pointer to 2nd param, minus 2 from the for loop
			if param1 != 0 {
				if paramMode2 == 0 {
					i = programInput[programInput[i+2]] - 2
				} else if paramMode2 == 1{
					i = programInput[i+2] - 2
				} else {
					relativeBaseOffset := programInput[i+2]
					i = programInput[relativeBase + relativeBaseOffset] - 2
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
			} else if paramMode1 == 1{
				param1 = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				param1 = programInput[relativeBase +relativeBaseOffset]
			}
			//set instruction pointer to 2nd param, minus 2 from the for loop
			if param1 != 0 {
				if paramMode2 == 0 {
					i = programInput[programInput[i+2]] - 2
				} else if paramMode2 == 1{
					i = programInput[i+2] - 2
				} else {
					relativeBaseOffset := programInput[i+2]
					i = programInput[relativeBase + relativeBaseOffset] - 2
				}
			} else {
				i++
			}
			fmt.Println("paramMode1", paramMode1, "paramMode2", paramMode2, "param1", param1, "new IP", i)
		// less than
		case 7:
			paramMode1 := (programInput[i] / 100) % 10
			paramMode2 := (programInput[i] / 1000) % 10
			var param1, param2 int
			if paramMode1 == 0 {
				param1 = programInput[programInput[i+1]]
			} else if paramMode1 == 1 {
				param1 = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				param1 = programInput[relativeBase + relativeBaseOffset]
			}
			if paramMode2 == 0 {
				param2 = programInput[programInput[i+2]]
			} else if paramMode2 == 1 {
				param2 = programInput[i+2]
			} else {
				relativeBaseOffset := programInput[i+1]
				param2 = programInput[relativeBase + relativeBaseOffset]
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
			} else if paramMode1 == 1 {
				param1 = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				param1 = programInput[relativeBase + relativeBaseOffset]
			}
			if paramMode2 == 0 {
				param2 = programInput[programInput[i+2]]
			} else if paramMode2 == 1 {
				param2 = programInput[i+2]
			} else {
				relativeBaseOffset := programInput[i+1]
				param2 = programInput[relativeBase + relativeBaseOffset]
			}
			dest := programInput[i+3]
			if param1 == param2 {
				programInput[dest] = 1
			} else {
				programInput[dest] = 0
			}
			fmt.Println("paramMode1", paramMode1, "paramMode2", paramMode2, "param1", param1, "param2", param2, "dest", dest, "value", programInput[dest])
			i += 2
		// relative base offset
		case 9:
			var offset int
			paramMode1 := (programInput[i] / 100) % 10
			if paramMode1 == 0 {
				offset = programInput[programInput[i+1]]
			} else if paramMode1 == 1{
				offset = programInput[i+1]
			} else {
				relativeBaseOffset := programInput[i+1]
				offset = programInput[relativeBase+relativeBaseOffset]
			}
			relativeBase += offset
			fmt.Println("relative base: ", relativeBase)
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
