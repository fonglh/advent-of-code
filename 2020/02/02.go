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
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	file, err := os.Open("02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		minCharCount, _ := strconv.Atoi(strings.Split(line, "-")[0])
		maxCharCount, _ := strconv.Atoi(strings.Split(strings.Split(line, "-")[1], " ")[0])
		checkLetter := strings.Split(line, " ")[1][0]
		password := strings.Split(line, " ")[2]
		if isValid(minCharCount, maxCharCount, checkLetter, password) {
			validCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return validCount
}

func isValid(minCharCount, maxCharCount int, checkLetter byte, password string) bool {
	var letterCount int
	for i := 0; i < len(password); i++ {
		if password[i] == checkLetter {
			letterCount += 1
		}
	}
	return letterCount >= minCharCount && letterCount <= maxCharCount
}

func puzzle2() int {
	file, err := os.Open("02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pos1, _ := strconv.Atoi(strings.Split(line, "-")[0])
		pos2, _ := strconv.Atoi(strings.Split(strings.Split(line, "-")[1], " ")[0])
		checkLetter := strings.Split(line, " ")[1][0]
		password := strings.Split(line, " ")[2]
		if isValid2(pos1, pos2, checkLetter, password) {
			validCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return validCount
}

func isValid2(pos1, pos2 int, checkLetter byte, password string) bool {
	return (password[pos1-1] == checkLetter && password[pos2-1] != checkLetter) !=
		(password[pos1-1] != checkLetter && password[pos2-1] == checkLetter)
}
