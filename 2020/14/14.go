package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile("14.txt")

	fmt.Println(puzzle1(input))
}

func puzzle1(input []string) int64 {
	var mask string
	memory := make(map[int64]int64)

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			mask = getMaskString(line)
		} else {
			address, rawValue := getMemoryInstruction(line)
			memory[address] = getModifiedValue(rawValue, mask)
		}
	}

	var sumMemory int64
	for _, value := range memory {
		sumMemory += value
	}
	return sumMemory
}

func getMaskString(line string) string {
	r := regexp.MustCompile(`^mask = ([X10]{36})`)
	return r.FindStringSubmatch(line)[1]
}

func replaceMaskXWith(maskInput, xReplacement string) int64 {
	maskString := strings.ReplaceAll(maskInput, "X", xReplacement)
	mask, _ := strconv.ParseInt(maskString, 2, 64)
	return mask
}

func getMemoryInstruction(line string) (int64, int64) {
	r := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)`)
	matches := r.FindAllStringSubmatch(line, -1)
	address, _ := strconv.Atoi(matches[0][1])
	value, _ := strconv.Atoi(matches[0][2])
	return int64(address), int64(value)
}

func getModifiedValue(value int64, mask string) int64 {
	value = value | replaceMaskXWith(mask, "0")
	value = value & replaceMaskXWith(mask, "1")

	return value
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
