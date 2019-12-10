package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	picture := readPic("08.txt")
	fmt.Println(puzzle1(picture))
	puzzle2(picture)
}

func puzzle1(picture []string) int {
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

func puzzle2(picture []string) {
	visibleImage := make([]byte, 25*6)
	for i := 0; i < (25 * 6); i++ {
		for layerNum := 0; layerNum < len(picture); layerNum++ {
			if picture[layerNum][i] != '2' {
				visibleImage[i] = picture[layerNum][i]
				break
			}
		}
	}
	for i := range visibleImage {
		if i%25 == 0 {
			fmt.Println()
		}
		if visibleImage[i] == '1' {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
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
