package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := readFile("rules.txt")
	messages := readFile("19.txt")

	fmt.Println(puzzle1(rules, messages))
	fmt.Println(puzzle2(rules))
}

func puzzle1(rules, messages []string) int {
	// build rules dictionary
	rulesMap := make(map[int][][]string)
	for _, rule := range rules {
		ruleNum, ruleList := parseRule(rule)
		rulesMap[ruleNum] = ruleList
	}

	// find all possibilites for rule 0
	allPossibilitesDict := buildListDict(findPossibilities(rulesMap, 0))

	var numMatch int
	for _, msg := range messages {
		if allPossibilitesDict[msg] {
			numMatch += 1
		}
	}

	return numMatch
}

func parseRule(rule string) (int, [][]string) {
	ruleParts := strings.Split(rule, ": ")
	ruleNumber, _ := strconv.Atoi(ruleParts[0])

	ruleList := make([][]string, 0)
	subRules := strings.Split(ruleParts[1], " | ")

	for _, subRule := range subRules {
		subRuleComponents := strings.Split(subRule, " ")
		subRuleComponentList := make([]string, 0)
		for _, subRuleComponent := range subRuleComponents {
			subRuleComponentList = append(subRuleComponentList, subRuleComponent)
		}
		ruleList = append(ruleList, subRuleComponentList)
	}

	return ruleNumber, ruleList
}

func findPossibilities(rulesMap map[int][][]string, ruleNum int) []string {
	possibilities := make([]string, 0)

	subRules := rulesMap[ruleNum]

	for _, subRule := range subRules {
		subRulePossibilities := make([]string, 0)

		for _, subRuleComponent := range subRule {
			// base case, "a" or "b"
			if subRuleComponent == "\"a\"" {
				possibilities = append(possibilities, "a")
			} else if subRuleComponent == "\"b\"" {
				possibilities = append(possibilities, "b")
			} else {
				// recurse into another subrule
				subRuleComponentNum, _ := strconv.Atoi(subRuleComponent)
				nextPossibilities := findPossibilities(rulesMap, subRuleComponentNum)

				// "multiply" existing possibilites for this subrule with the new possibilities from the current component
				subRulePossibilities = strSumProduct(subRulePossibilities, nextPossibilities)
			}
		}

		// after the subrule is done, append all its possibilities to the ones already found for the rule.
		// (i.e. previous subrules, but in the input there are a maximum of 2)
		possibilities = append(possibilities, subRulePossibilities...)
	}

	return possibilities
}

// build dictionary with list contents so checking membership is O(1)
func buildListDict(list []string) map[string]bool {
	listDict := make(map[string]bool)

	for _, item := range list {
		listDict[item] = true
	}

	return listDict
}

// append each item in list 2 to each item in list 1
// return the other list if either list is empty
func strSumProduct(list1, list2 []string) []string {
	combined := make([]string, 0)

	if len(list1) == 0 {
		return list2
	}

	if len(list2) == 0 {
		return list1
	}

	for _, str1 := range list1 {
		for _, str2 := range list2 {
			combined = append(combined, str1+str2)
		}
	}
	return combined
}

func puzzle2(input []string) int {
	return 0
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
