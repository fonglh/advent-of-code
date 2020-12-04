package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile()

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func readFile() []string {
	file, err := os.Open("04.txt")
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

func fieldsPresent(passport string) bool {
	neededFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range neededFields {
		if !strings.Contains(passport, field+":") {
			return false
		}
	}
	return true
}

func validateFields(passport string) bool {
	return validateYear(passport, "byr", 1920, 2002) &&
		validateYear(passport, "iyr", 2010, 2020) &&
		validateYear(passport, "eyr", 2020, 2030) &&
		validateHeight(passport) &&
		validateHairColor(passport) &&
		validateEyeColor(passport) &&
		validatePid(passport)
}

func getValue(passport, fieldName string) (string, error) {
	data := strings.Split(passport, " ")

	for _, val := range data {
		if strings.Contains(val, fieldName) {
			return strings.Split(val, ":")[1], nil
		}
	}

	return "", errors.New("field not found")
}

func validateYear(passport, fieldName string, minYear, maxYear int) bool {
	yearStr, err := getValue(passport, fieldName)
	if err != nil {
		return false
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}

	return year >= minYear && year <= maxYear
}

func validateHeight(passport string) bool {
	heightStr, err := getValue(passport, "hgt")
	if err != nil {
		return false
	}

	if strings.HasSuffix(heightStr, "in") {
		// can also do with regex or string indexing
		height, err := strconv.Atoi(strings.Split(heightStr, "i")[0])
		if err != nil {
			return false
		}
		return height >= 59 && height <= 76
	} else if strings.HasSuffix(heightStr, "cm") {
		height, err := strconv.Atoi(strings.Split(heightStr, "c")[0])
		if err != nil {
			return false
		}
		return height >= 150 && height <= 193
	} else {
		return false
	}
}

func validateHairColor(passport string) bool {
	colorStr, err := getValue(passport, "hcl")
	if err != nil {
		return false
	}

	matched, _ := regexp.MatchString(`#[0-9a-f]{6}`, colorStr)
	return matched
}

func validateEyeColor(passport string) bool {
	colorStr, err := getValue(passport, "ecl")
	if err != nil {
		return false
	}

	return strings.Contains("amb blu brn gry grn hzl oth", colorStr)
}

func validatePid(passport string) bool {
	pid, err := getValue(passport, "pid")
	if err != nil {
		return false
	}

	matched, _ := regexp.MatchString(`^[0-9]{9}$`, pid)
	return matched
}

func puzzle1(input []string) int {
	var validCount int
	for _, passport := range input {
		if fieldsPresent(passport) {
			validCount += 1
		}
	}
	return validCount
}

func puzzle2(input []string) int {
	var validCount int
	for _, passport := range input {
		if fieldsPresent(passport) && validateFields(passport) {
			validCount += 1
		}
	}
	return validCount
}
