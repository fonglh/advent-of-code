package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
}

func puzzle1() {
	currentPhase := readFile("16.txt")
	for phaseNum := 0; phaseNum < 100; phaseNum++ {
		nextPhase := make([]int, len(currentPhase))
		for i := range currentPhase {
			nextPhase[i] = nthElement(currentPhase, i+1)
		}
		currentPhase = nextPhase
	}
	fmt.Println(currentPhase[:8])
}

// generate the nth element. start list numbering from 1
func nthElement(inputList []int, n int) int {
	var fftSum int
	pattern := generatePattern(n, len(inputList))
	for i := range inputList {
		fftSum += inputList[i] * pattern[i]
	}
	return abs(fftSum % 10)
}

// generate the base pattern depending on which element it is
func generatePattern(numRepeat, length int) []int {
	pattern := make([]int, length+1)
	basePattern := []int{0, 1, 0, -1}
	for patternIdx, baseIdx := 0, 0; patternIdx < length+1; {
		for currCount := 0; currCount < numRepeat && patternIdx < length+1; currCount++ {
			pattern[patternIdx] = basePattern[baseIdx]
			patternIdx++
		}
		baseIdx = (baseIdx + 1) % 4
	}
	// offset by 1, as given in instructions
	return pattern[1:]
}

// https://www.golangprograms.com/go-program-to-read-a-text-file-character-by-character.html
// There might be a better way to do this without all the conversions
func readFile(filename string) []int {
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	output := make([]int, len(inputdata)-1)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	var idx int
	for data.Scan() {
		digit, err := strconv.Atoi(data.Text())
		if err != nil {
			fmt.Println(err)
		} else {
			output[idx] = digit
		}
		idx++
	}
	return output
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
