package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	orbitMap := readFile("06.txt")
	var orbitCount int
	for k := range orbitMap {
		curr := k
		for curr != "" {
			curr = orbitMap[curr]
			orbitCount++
		}
		orbitCount-- // Subtract out the extra step from the center to nothing
	}
	return orbitCount
}

func puzzle2() int {
	orbitMap := readFile("06.txt")

	// Build chains from our locations to COM
	var myChain []string
	var santaChain []string
	for curr := orbitMap["YOU"]; curr != ""; curr = orbitMap[curr] {
		myChain = append(myChain, curr)
	}
	for curr := orbitMap["SAN"]; curr != ""; curr = orbitMap[curr] {
		santaChain = append(santaChain, curr)
	}
	fmt.Println(myChain)
	fmt.Println(santaChain)

	// Find first ancestor node
	var ancestor string
	for _, object := range myChain {
		for _, santaObject := range santaChain {
			if object == santaObject {
				ancestor = object
				break
			}
		}
		if ancestor != "" {
			break
		}
	}
	fmt.Println(ancestor)

	// Count number of steps from each of us to the common ancestor
	var transferCount int
	for _, object := range myChain {
		if object == ancestor {
			break
		}
		transferCount++
	}
	for _, object := range santaChain {
		if object == ancestor {
			break
		}
		transferCount++
	}
	return transferCount
}

func readFile(filename string) map[string]string {
	orbitMap := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		objects := strings.Split(scanner.Text(), ")")
		orbitMap[objects[1]] = objects[0]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return orbitMap
}
