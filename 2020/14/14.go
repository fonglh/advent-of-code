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
	fmt.Println(puzzle2(input))
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

	return sumMemoryValues(memory)
}

func puzzle2(input []string) int64 {
	var mask string
	memory := make(map[int64]int64)

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			mask = getMaskString(line)
		} else {
			address, value := getMemoryInstruction(line)
			floatingAddress := getFloatingAddress(address, mask)
			allAddresses := getAllAddresses(floatingAddress)
			for _, addr := range allAddresses {
				addrInt, _ := strconv.ParseInt(addr, 2, 64)
				memory[addrInt] = value
			}
		}
	}
	return sumMemoryValues(memory)
}

func sumMemoryValues(memory map[int64]int64) int64 {
	var sumMemory int64
	for _, value := range memory {
		sumMemory += value
	}
	return sumMemory
}

func getFloatingAddress(rawAddress int64, mask string) string {
	addressStr := fmt.Sprintf("%036b", rawAddress)
	floatingAddress := make([]byte, 36)

	for i := 0; i < 36; i += 1 {
		if mask[i] == '1' {
			floatingAddress[i] = '1'
		} else if mask[i] == 'X' {
			floatingAddress[i] = 'X'
		} else {
			floatingAddress[i] = addressStr[i]
		}
	}

	return string(floatingAddress)
}

func getAllAddresses(floatingAddress string) []string {
	allAddresses := make([]string, 0)
	numX := strings.Count(floatingAddress, "X")
	maxCounter := 1 << numX
	formatString := fmt.Sprintf("%%0%dv", numX) //to zero pad all the possible X values

	for i := 0; i < maxCounter; i++ {
		currAddress := floatingAddress
		iStr := fmt.Sprintf(formatString, strconv.FormatInt(int64(i), 2))
		for j := 0; j < len(iStr); j++ {
			// Replace 1 X at a time with the bit values of the counter
			currAddress = strings.Replace(currAddress, "X", string(iStr[j]), 1)
		}
		allAddresses = append(allAddresses, currAddress)
	}

	return allAddresses
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
