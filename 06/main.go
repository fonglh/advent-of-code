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
	return 0
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
