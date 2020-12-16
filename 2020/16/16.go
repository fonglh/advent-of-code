package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ValidRange struct {
	min int
	max int
}

func main() {
	ticketFilename := "tickets.txt"
	rulesFilename := "ranges.txt"
	myTicket := []int{223, 139, 211, 131, 113, 197, 151, 193, 127, 53, 89, 167, 227, 79, 163, 199, 191, 83, 137, 149}
	//myTicket := []int{11, 12, 13}

	tickets := getTickets(ticketFilename)
	validTickets := getValidTickets(tickets, rulesFilename)
	validTickets = append(validTickets, myTicket)
	rules := getRules(rulesFilename)

	fmt.Println(puzzle1(tickets, rulesFilename))
	fmt.Println(puzzle2(validTickets, rules))
}

func puzzle1(tickets [][]int, rulesFilename string) int {
	validNumbers := getValidNumbers(rulesFilename)
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

func puzzle2(validTickets [][]int, rules map[string][]ValidRange) int {
	fieldOrder := make([]string, len(rules))
	myTicket := validTickets[len(validTickets)-1]

	findFieldOrder(validTickets, rules, fieldOrder, 0)

	answer := 1
	for i, field := range fieldOrder {
		if strings.HasPrefix(field, "departure") {
			answer *= myTicket[i]
		}
	}

	printTicket(myTicket, fieldOrder)
	return answer
}

func printTicket(ticket []int, fieldOrder []string) {
	for i := 0; i < len(ticket); i += 1 {
		fmt.Println(fieldOrder[i], ":", ticket[i])
	}
}

// Adapation of 8 queens problem
func findFieldOrder(validTickets [][]int, rules map[string][]ValidRange, fieldOrder []string, fieldOrderProgress int) bool {
	// base case, have reached the last field
	if fieldOrderProgress == len(rules) {
		return true
	}

	// map key iteration order is undefined, although it seems to follow insertion order
	rulesFields := make([]string, 0, len(rules))
	for k := range rules {
		rulesFields = append(rulesFields, k)
	}
	sort.Strings(rulesFields)

	// Try all the rules for this field
	for _, rule := range rulesFields {
		if !contains(fieldOrder, rule) && checkField(validTickets, rules, fieldOrderProgress, rule) {
			fieldOrder[fieldOrderProgress] = rule
			if findFieldOrder(validTickets, rules, fieldOrder, fieldOrderProgress+1) { // can reach the end, i.e. valid solution
				return true
			} else {
				// remove the rule for backtracking
				fieldOrder[fieldOrderProgress] = ""
			}
		}
	}

	return false
}

// fields cannot be reused for more than 1 column, so this is used to check if it has already been used.
func contains(input []string, e string) bool {
	for _, item := range input {
		if item == e {
			return true
		}
	}
	return false
}

// returns true if the fieldName is valid for a given position, false if not
func checkField(validTickets [][]int, rules map[string][]ValidRange, fieldIndex int, fieldName string) bool {
	fieldValid := true
	fieldRules := rules[fieldName]

	for _, ticket := range validTickets {
		fieldValue := ticket[fieldIndex]

		if !((fieldValue >= fieldRules[0].min && fieldValue <= fieldRules[0].max) ||
			(fieldValue >= fieldRules[1].min && fieldValue <= fieldRules[1].max)) {
			fieldValid = false
			break
		}
	}

	return fieldValid
}

// filter the data set
func getValidTickets(tickets [][]int, filename string) [][]int {
	validNumbers := getValidNumbers(filename)
	validTickets := make([][]int, 0)

	for _, ticket := range tickets {
		valid := true
		for _, number := range ticket {
			if validNumbers[number] == false {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
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

// for part 1 only
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

func getRules(filename string) map[string][]ValidRange {
	rules := make(map[string][]ValidRange)
	input := readFile(filename)

	for _, line := range input {
		lineSplit := strings.Split(line, ":")
		fieldName := lineSplit[0]
		data := lineSplit[1]
		data = strings.Trim(data, " ")
		lineRanges := strings.Split(data, " or ")

		for _, numRange := range lineRanges {
			minInt, maxInt := getMinMax(numRange)
			rules[fieldName] = append(rules[fieldName], ValidRange{minInt, maxInt})
		}
	}

	return rules
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
