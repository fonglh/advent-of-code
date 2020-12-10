package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readFile("10.txt")
	fmt.Println(puzzle1(input))

	// input is modified in the solutions, so the file should be read again.
	input = readFile("10.txt")
	fmt.Println(puzzle2(input))
}

func puzzle1(input []int) int {
	var diffOne int
	var diffThree int

	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	for i := 1; i < len(input); i++ {
		if input[i]-input[i-1] == 1 {
			diffOne += 1
		} else if input[i]-input[i-1] == 3 {
			diffThree += 1
		}
	}
	fmt.Println(diffOne, diffThree)
	return diffOne * diffThree
}

func puzzle2(input []int) int64 {
	// do not prepend 0 so the loop construct is simpler.
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	// numWays is the number of ways to reach a particular jolt level
	numWays := make(map[int]int64)
	numWays[0] = 1 // there is 1 way to reach the '0' (power source)

	for _, adapter := range input {
		// because each adapter can take 1, 2 or 3 jolts lower, the number of ways to reach each
		// adapter is the sum of the ways to reach the adapters 1, 2 or 3 jolts below it.
		//
		// since only adapter values in the input are assigned values, the adapters which don't exist
		// have 0 ways to reach them, so we can just include them in the summation.
		numWays[adapter] = numWays[adapter-1] + numWays[adapter-2] + numWays[adapter-3]
	}

	// pass in the device jolt level (last item in the modified input) to get the number of ways to reach it.
	return numWays[input[len(input)-1]]
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
