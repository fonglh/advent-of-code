package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile("25.txt")

	fmt.Println(puzzle1(input))
	//	fmt.Println(puzzle2(input))
}

func puzzle1(input []int) int {
	loopSizes := make([]int, 0)
	for _, publicKey := range input {
		loopSize := findLoopSize(publicKey)
		loopSizes = append(loopSizes, loopSize)
	}

	encryptionKey1 := transform(input[0], loopSizes[1])
	encryptionKey2 := transform(input[1], loopSizes[0])
	if encryptionKey1 == encryptionKey2 {
		return encryptionKey1
	}

	// should never get here
	return -1
}

func puzzle2(input []int) int {
	return 0
}

func findLoopSize(publicKey int) int {
	value := 1
	loopSize := 0
	initialSubjectNumber := 7

	for {
		value = value * initialSubjectNumber
		value = value % 20201227
		loopSize += 1
		if value == publicKey {
			return loopSize
		}
	}

	return 0
}

func transform(initialSubjectNumber, loopSize int) int {
	value := 1

	for i := 0; i < loopSize; i += 1 {
		value = value * initialSubjectNumber
		value = value % 20201227
	}

	return value
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
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
