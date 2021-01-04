package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Tile struct {
	id      int
	image   []string
	matches map[int]bool
}

func (t Tile) getEdge(rotation int) string {
	var edgeString string

	switch rotation {
	case 0: // top, left to right
		edgeString = t.image[0]
	case 1: // left, top to bottom
		buf := make([]byte, 0, 10)
		for i := 0; i < 10; i += 1 {
			buf = append(buf, t.image[i][0])
		}
		edgeString = string(buf)
	case 2: // bottom, left to right
		edgeString = t.image[len(t.image)-1]
	case 3: // right, top to bottom
		buf := make([]byte, 0, 10)
		for i := 0; i < 10; i += 1 {
			buf = append(buf, t.image[i][len(t.image[i])-1])
		}
		edgeString = string(buf)
	}
	return edgeString
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func edgeMatch(str1, str2 string) bool {
	return str1 == str2 || reverse(str1) == str2
}

func main() {
	input := readFile("20.txt")
	tiles := getTiles(input)

	fmt.Println(puzzle1(tiles))
	//fmt.Println(puzzle2(rules, messages))
}

func puzzle1(tiles map[int]Tile) int {
	for tileId1, tile1 := range tiles {
		for tileId2, tile2 := range tiles {
			if tileId1 == tileId2 {
				continue
			}
			for i := 0; i < 4; i += 1 {
				for j := 0; j < 4; j += 1 {
					if edgeMatch(tile1.getEdge(i), tile2.getEdge(j)) {
						tile1.matches[tileId2] = true
						tile2.matches[tileId1] = true
					}
				}
			}
		}
	}

	idProduct := 1
	for id, tile := range tiles {
		if len(tile.matches) == 2 {
			idProduct *= id
		}
	}
	return idProduct
}

func puzzle2(rules, messages []string) int {
	return 0
}

func getTiles(input []string) map[int]Tile {
	tiles := make(map[int]Tile)

	for lineNumber := 0; lineNumber < len(input); lineNumber += 1 {
		image := make([]string, 0)
		line := input[lineNumber]
		idLine := line[:len(line)-1]
		idNumStr := strings.Split(idLine, " ")[1]
		idNum, _ := strconv.Atoi(idNumStr)

		lineNumber += 1
		for i := 0; i < 10; i += 1 {
			image = append(image, input[lineNumber+i])
		}
		tiles[idNum] = Tile{id: idNum, image: image, matches: make(map[int]bool)}

		lineNumber += 10
	}

	return tiles
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
