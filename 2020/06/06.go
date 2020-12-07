package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile()

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func readFile() []string {
	file, err := os.Open("06.txt")
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

func puzzle1(input []string) int {
	var sumCounts int
	for _, line := range input {
		var lineCount int
		for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
			if strings.Count(line, string(letter)) > 0 {
				lineCount += 1
			}
		}
		sumCounts += lineCount
	}
	return sumCounts
}

func puzzle2(input []string) int {
	var sumCounts int
	for _, line := range input {
		var lineCount int
		numPeople := len(strings.Split(line, " "))
		for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
			if strings.Count(line, string(letter)) == numPeople {
				lineCount += 1
			}
		}
		sumCounts += lineCount
	}
	return sumCounts
}
