package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuel int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moduleMass, _ := strconv.Atoi(scanner.Text())
		moduleFuel := moduleMass/3 - 2
		totalFuel += moduleFuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return totalFuel
}

func puzzle2() int {
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuel int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moduleMass, _ := strconv.Atoi(scanner.Text())
		var totalModuleFuel, stepModuleFuel int
		stepModuleFuel = moduleMass/3 - 2
		for stepModuleFuel > 0 {
			totalModuleFuel += stepModuleFuel
			stepModuleFuel = stepModuleFuel/3 - 2
		}
		totalFuel += totalModuleFuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return totalFuel
}
