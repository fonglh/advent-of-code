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
	input := readFile("07.txt")
	bagRules := buildBagRules(input)

	fmt.Println(puzzle1(bagRules))
	//fmt.Println(puzzle2(input))
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

func puzzle1(bagRules map[string]map[string]int) int {
	answer := make(map[string]bool)
	queue := []string{"shiny gold"}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for bag, contents := range bagRules {
			if answer[bag] { //already counted this bag
				continue
			}
			if contents[current] > 0 {
				queue = append(queue, bag)
				answer[bag] = true
			}
		}
	}
	return len(answer)
}

func puzzle2(input []string) int {
	return 0
}
