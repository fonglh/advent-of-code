package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var slope []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slope = append(slope, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(puzzle1(slope))
	fmt.Println(puzzle2(slope))
}

func puzzle1(slope []string) int {
	return countTrees(slope, 3, 1)
}

func countTrees(slope []string, horizontalOffset, verticalOffset int) int {
	var numTrees, hPos, vPos int
	for vPos < len(slope) {
		if slope[vPos][hPos] == '#' {
			numTrees += 1
		}
		hPos = (hPos + horizontalOffset) % len(slope[vPos])
		vPos += verticalOffset
	}
	return numTrees
}

func puzzle2(slope []string) int {
	product := 1
	var gradients = []struct {
		hOffset int
		vOffset int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, gradient := range gradients {
		product *= countTrees(slope, gradient.hOffset, gradient.vOffset)
	}

	return product
}
