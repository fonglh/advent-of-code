package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func puzzle1() int {
	picture := readPic("08.txt")
	var minZeroLayer int
	minZeros := countDigit(picture[0], '0')
	for i, layer := range picture {
		countZero := countDigit(layer, '0')
		if countZero < minZeros {
			minZeroLayer = i
			minZeros = countZero
		}
	}
	return countDigit(picture[minZeroLayer], '1') * countDigit(picture[minZeroLayer], '2')
}

func puzzle2() int {
	return 0
}

func readPic(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fullPicture string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fullPicture = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var layeredPicture []string
	layerLength := 25 * 6
	for i := 0; i < len(fullPicture); i += layerLength {
		layer := fullPicture[i : i+layerLength]
		layeredPicture = append(layeredPicture, layer)
	}
	return layeredPicture
}

func countDigit(layer string, digit rune) int {
	var count int
	for _, char := range layer {
		if char == digit {
			count++
		}
	}
	return count
}
