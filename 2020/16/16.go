package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"regexp"
	"strconv"
	"strings"
)

func main() {
	tickets := getTickets("tickets-test.txt")

	fmt.Println(puzzle1(tickets))
	//	fmt.Println(puzzle2(input))
}

func puzzle1(tickets [][]int) int {
	validNumbers := getValidNumbers("ranges-test.txt")
	var ticketErrorRate int //actually the sum of invalid numbers encountered

	for _, ticket := range tickets {
		for _, number := range ticket {
			if validNumbers[number] == false {
				ticketErrorRate += number
			}
		}
	}
	return ticketErrorRate
}

func puzzle2(input []string) int {
	return 0
}

func getTickets(filename string) [][]int {
	input := readFile(filename)
	tickets := make([][]int, len(input))

	for i, ticket := range input {
		fields := strings.Split(ticket, ",")
		ticketFields := make([]int, len(fields))
		for j, field := range fields {
			fieldInt, _ := strconv.Atoi(field)
			ticketFields[j] = fieldInt
		}
		tickets[i] = ticketFields
	}

	return tickets
}

func getValidNumbers(filename string) []bool {
	input := readFile(filename)
	validNumbers := make([]bool, 1000) //from visual inspection of the input

	for _, line := range input {
		data := strings.Split(line, ":")[1]
		data = strings.Trim(data, " ")
		lineRanges := strings.Split(data, " or ")

		for _, numRange := range lineRanges {
			minInt, maxInt := getMinMax(numRange)
			for i := minInt; i <= maxInt; i += 1 {
				validNumbers[i] = true
			}
		}
	}

	return validNumbers
}

func getMinMax(input string) (int, int) {
	minMaxStr := strings.Split(input, "-")
	minInt, _ := strconv.Atoi(minMaxStr[0])
	maxInt, _ := strconv.Atoi(minMaxStr[1])
	return minInt, maxInt
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
