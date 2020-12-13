package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := readFile()

	fmt.Println(puzzle1(input))
	//fmt.Println(puzzle2(input))
}

func readFile() []string {
	file, err := os.Open("07.txt")
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

// Example entry: "light orange": { "dark maroon": 1, "dim maroon": 3, ... }
func buildBagRules(input []string) map[string]map[string]int {
	bagRules := make(map[string]map[string]int)
	containingBagRegex := regexp.MustCompile(`(?P<bag>(.*)) bags contain`)
	contentsRegex := regexp.MustCompile(`(?P<count>\d+) (?P<colour>\w+ \w+) bag`)

	for _, line := range input {
		bag := containingBagRegex.FindStringSubmatch(line)[1]

		contentsMap := make(map[string]int)
		contents := contentsRegex.FindAllStringSubmatch(line, -1)
		for _, containedBag := range contents {
			count, _ := strconv.Atoi(containedBag[1])
			contentsMap[containedBag[2]] = count
		}

		bagRules[bag] = contentsMap
	}

	return bagRules
}

func puzzle1(input []string) int {
	bagRules := buildBagRules(input)

	for bag, contents := range bagRules {
		fmt.Println(bag, contents)
	}
	return 0
}

func puzzle2(input []string) int {
	return 0
}
