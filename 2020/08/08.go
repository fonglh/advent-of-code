package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("08.txt")

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input []string) (int, error) {
	var accumulator int
	seen := make(map[int]bool)

	for instructionPtr := 0; instructionPtr < len(input); {
		if seen[instructionPtr] == true {
			return accumulator, errors.New("Infinite loop")
		}
		instruction, number := parseInstruction(input[instructionPtr])
		seen[instructionPtr] = true

		switch instruction {
		case "nop":
			instructionPtr += 1
		case "acc":
			accumulator += number
			instructionPtr += 1
		case "jmp":
			instructionPtr += number
		}
	}
	return accumulator, nil
}

func parseInstruction(line string) (string, int) {
	parts := strings.Split(line, " ")
	number, _ := strconv.Atoi(parts[1])
	return parts[0], number
}

func puzzle2(input []string) int {
	var testPtr int

	for {
		instruction, number := parseInstruction(input[testPtr])
		switch instruction {
		case "acc":
			testPtr += 1
		case "nop":
			input[testPtr] = "jmp " + strconv.Itoa(number)
			accumulator, err := puzzle1(input)
			if err == nil {
				return accumulator
			}
			input[testPtr] = "nop " + strconv.Itoa(number)
			testPtr += 1
		case "jmp":
			input[testPtr] = "nop " + strconv.Itoa(number)
			accumulator, err := puzzle1(input)
			if err == nil {
				return accumulator
			}
			input[testPtr] = "jmp " + strconv.Itoa(number)
			testPtr += 1
		}
	}
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
