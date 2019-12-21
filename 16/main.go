package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	inputList := readFile("16.txt")
	fmt.Println(fft(inputList)[:8])
}

func puzzle2() {
	inputList := readFile("16.txt")
	offset := messageOffset(inputList)
	var inputSignal []int
	for i := 0; i < 10000; i++ {
		inputSignal = append(inputSignal, inputList...)
	}
	inputSignal = inputSignal[offset:]
	fmt.Println(fft2(inputSignal)[:8])
}

// Everything before the offset doesn't matter as they're all zeroed out
// Since the offset is more than halfway through, the matrix is an upside
// down triangle of 1s.
// Thus the matrix multiplication just involves adding each element to the nth + 1
// element from the back.
func fft2(inputList []int) []int {
	currentPhase := inputList
	for phaseNum := 0; phaseNum < 100; phaseNum++ {
		nextPhase := make([]int, len(currentPhase))
		nextPhase[len(currentPhase)-1] = currentPhase[len(currentPhase)-1]
		for i := len(currentPhase) - 2; i >= 0; i-- {
			nextPhase[i] = abs(nextPhase[i+1]+currentPhase[i]) % 10
		}
		currentPhase = nextPhase
	}
	return currentPhase
}

func fft(inputList []int) []int {
	currentPhase := inputList
	for phaseNum := 0; phaseNum < 100; phaseNum++ {
		nextPhase := make([]int, len(currentPhase))
		for i := range currentPhase {
			nextPhase[i] = nthElement(currentPhase, i+1)
		}
		currentPhase = nextPhase
	}
	return currentPhase
}

// generate the nth element. start list numbering from 1
func nthElement(inputList []int, n int) int {
	var fftSum int
	//pattern := generatePattern(n, len(inputList))
	for i := range inputList {
		fftSum += inputList[i] * patternDigit(n, i)
		//fftSum += inputList[i] * pattern[i]
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

func patternDigit(numRepeat, idx int) int {
	basePattern := []int{0, 1, 0, -1}
	return basePattern[((idx+1)/numRepeat)%4]
}

func messageOffset(inputList []int) int {
	digits := inputList[:7]
	var offset int
	for i, digit := range digits {
		offset += digit * int(math.Pow(10, float64(6-i)))
	}
	return offset
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
