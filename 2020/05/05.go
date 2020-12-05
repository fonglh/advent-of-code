package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile()

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func readFile() []string {
	file, err := os.Open("05.txt")
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

func getSeatId2(boardingPass string) int {
	var binStr string
	for _, char := range boardingPass {
		if char == 'F' || char == 'L' {
			binStr += "0"
		} else {
			binStr += "1"
		}
	}
	seatId, _ := strconv.ParseInt(binStr, 2, 32)
	return int(seatId)
}

func getSeatId(boardingPass string) int {
	var seatRow int
	minRow := 0
	maxRow := 127
	for i := 0; i < 7; i++ {
		seatRow = (minRow + maxRow) / 2
		if boardingPass[i] == 'F' {
			maxRow = seatRow
		} else {
			minRow = seatRow + 1
		}
	}
	seatRow = (minRow + maxRow) / 2

	var seatCol int
	minCol := 0
	maxCol := 7
	for i := 7; i < 10; i++ {
		seatCol = (minCol + maxCol) / 2
		if boardingPass[i] == 'L' {
			maxCol = seatCol
		} else {
			minCol = seatCol + 1
		}
	}
	seatCol = (minCol + maxCol) / 2
	return seatRow*8 + seatCol
}

func getSeatRange(input []string) (int, int) {
	var maxSeatId int
	minSeatId := 127 * 8
	for _, boardingPass := range input {
		seatId := getSeatId2(boardingPass)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
		if seatId < minSeatId {
			minSeatId = seatId
		}
	}
	return minSeatId, maxSeatId
}

func puzzle1(input []string) int {
	_, maxSeatId := getSeatRange(input)
	return maxSeatId
}

func puzzle2(input []string) int {
	takenSeats := make(map[int]bool)
	for _, boardingPass := range input {
		takenSeats[getSeatId2(boardingPass)] = true
	}

	minSeatId, maxSeatId := getSeatRange(input)

	for i := minSeatId; i <= maxSeatId; i++ {
		if takenSeats[i] == false {
			return i
		}
	}
	return 0
}
