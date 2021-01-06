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
	matches map[string]int // match direction to tile ID
}

func (t Tile) print() {
	fmt.Println(t.id)
	for _, line := range t.image {
		fmt.Println(line)
	}
	fmt.Println(t.matches)
}

// rotate image 90 degrees anticlockwise
func (t *Tile) rotate90() {
	rotatedImage := make([]string, 0)
	for i := len(t.image) - 1; i >= 0; i -= 1 {
		buf := make([]byte, len(t.image))
		for j := 0; j < len(t.image); j += 1 {
			buf[j] = t.image[j][i]
		}
		rotatedImage = append(rotatedImage, string(buf))
	}
	t.image = rotatedImage
	t.matches["n"], t.matches["w"], t.matches["s"], t.matches["e"] = t.matches["e"], t.matches["n"], t.matches["w"], t.matches["s"]
	// remove 0 values so it's easy to determine the corners
	for k, v := range t.matches {
		if v == 0 {
			delete(t.matches, k)
		}
	}
}

// flip image
func (t *Tile) flip() {
	flippedImage := make([]string, len(t.image))

	for i, line := range t.image {
		flippedImage[i] = reverse(line)
	}

	t.image = flippedImage
	t.matches["e"], t.matches["w"] = t.matches["w"], t.matches["e"]
	// remove 0 values so it's easy to determine the corners
	for k, v := range t.matches {
		if v == 0 {
			delete(t.matches, k)
		}
	}
}

func (t *Tile) stripBorder() {
	newImage := make([]string, 0)
	for i := 1; i < len(t.image)-1; i += 1 {
		newImage = append(newImage, t.image[i][1:len(t.image[i])-1])
	}
	t.image = newImage
}

func (t Tile) getEdge(rotation string) string {
	var edgeString string

	switch rotation {
	case "n": // top, left to right
		edgeString = t.image[0]
	case "w": // left, top to bottom
		buf := make([]byte, 0, 10)
		for i := 0; i < 10; i += 1 {
			buf = append(buf, t.image[i][0])
		}
		edgeString = string(buf)
	case "s": // bottom, left to right
		edgeString = t.image[len(t.image)-1]
	case "e": // right, top to bottom
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
	fmt.Println(puzzle2(tiles))
}

func puzzle1(tiles map[int]*Tile) int {
	directions := [4]string{"n", "w", "s", "e"}

	for tileId1, tile1 := range tiles {
		for tileId2, tile2 := range tiles {
			if tileId1 == tileId2 {
				continue
			}
			for i := 0; i < 4; i += 1 {
				for j := 0; j < 4; j += 1 {
					if edgeMatch(tile1.getEdge(directions[i]), tile2.getEdge(directions[j])) {
						tile1.matches[directions[i]] = tileId2
						tile2.matches[directions[j]] = tileId1
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

func puzzle2(tiles map[int]*Tile) int {
	// hardcode this to match the sample
	//cornerId := 1951
	//tiles[cornerId].flip()

	// find corner tile
	var cornerId int
	for id, tile := range tiles {
		if len(tile.matches) == 2 {
			cornerId = id
			break
		}
	}

	// rotate until the south and east sides match (top left tile)
	for !(tiles[cornerId].matches["s"] > 0 && tiles[cornerId].matches["e"] > 0) {
		tiles[cornerId].rotate90()
	}

	jigsaw := buildImage(tiles, cornerId, make([][]Tile, 0))

	// strip borders from the images in the tiles in the jigsaw
	for i, row := range jigsaw {
		for j, _ := range row {
			jigsaw[i][j].stripBorder()
		}
	}
	// convert 2D tile array into an array of strings to facilitate searching for sea monsters
	// use that to initalise a Tile so the rotation and flip functions can be used
	jigsawArray := strings.Split(jigsawString(jigsaw, false), "\n")
	// remove the trailing newline
	jigsawTile := Tile{id: 0, image: jigsawArray[:len(jigsawArray)-1], matches: make(map[string]int)}
	var waterCells int
	for _, line := range jigsawTile.image {
		waterCells += strings.Count(line, "#")
	}

	var monsterCount int
	var monsterCells int
	monster := []string{"                  # ", "#    ##    ##    ###", " #  #  #  #  #  #   "}
	for _, line := range monster {
		monsterCells += strings.Count(line, "#")
	}

	for rotationCount := 0; rotationCount < 4; rotationCount += 1 {
		for i := 0; i < len(jigsawTile.image); i += 1 {
			for j := 0; j < len(jigsawTile.image); j += 1 {
				if search(jigsawTile.image, monster, i, j) {
					monsterCount += 1
				}
			}
		}
		jigsawTile.rotate90()
	}
	jigsawTile.flip()
	for rotationCount := 0; rotationCount < 4; rotationCount += 1 {
		for i := 0; i < len(jigsawTile.image); i += 1 {
			for j := 0; j < len(jigsawTile.image); j += 1 {
				if search(jigsawTile.image, monster, i, j) {
					monsterCount += 1
				}
			}
		}
		jigsawTile.rotate90()
	}

	return waterCells - monsterCount*monsterCells
}

// search for the pattern in the image, starting from (rowCoord, colCoord) in the image
func search(image []string, pattern []string, rowCoord, colCoord int) bool {
	patternHeight := len(pattern)
	patternWidth := len(pattern[0])
	imageSize := len(image)

	// out of bounds, don't bother searching
	if rowCoord+patternHeight > imageSize ||
		colCoord+patternWidth > imageSize {
		return false
	}

	for i := 0; i < patternHeight; i += 1 {
		for j := 0; j < patternWidth; j += 1 {
			if pattern[i][j] == '#' && image[rowCoord+i][colCoord+j] != '#' {
				return false
			}
		}
	}
	return true
}

func buildImage(tiles map[int]*Tile, srcId int, accumulatedImage [][]Tile) [][]Tile {
	// base case, cannot go south any more, add the current row and return
	_, ok := tiles[srcId].matches["s"]
	if !ok {
		return append(accumulatedImage, matchRow(tiles, srcId, make([]Tile, 0)))
	}
	thisRow := matchRow(tiles, srcId, make([]Tile, 0))
	accumulatedImage = append(accumulatedImage, thisRow)
	// orient the south tile to match
	nextRowTile := matchTile(tiles, srcId, "s")
	return buildImage(tiles, nextRowTile.id, accumulatedImage)
}

func matchRow(tiles map[int]*Tile, srcId int, accumulatedRow []Tile) []Tile {
	// base case, cannot go east any more, add the current tile and return
	_, ok := tiles[srcId].matches["e"]
	if !ok {
		return append(accumulatedRow, *tiles[srcId])
	}
	nextTile := matchTile(tiles, srcId, "e")
	accumulatedRow = append(accumulatedRow, *tiles[srcId])
	return matchRow(tiles, nextTile.id, accumulatedRow)
}

// try all orientations of the other tile until the corresponding side matches
// return the matched tile
func matchTile(tiles map[int]*Tile, srcId int, side string) Tile {
	destinationId := tiles[srcId].matches[side]
	var destinationSide string
	switch side {
	case "n":
		destinationSide = "s"
	case "w":
		destinationSide = "e"
	case "s":
		destinationSide = "n"
	case "e":
		destinationSide = "w"
	}

	// try rotations
	for i := 0; i < 4; i += 1 {
		if tiles[destinationId].matches[destinationSide] == srcId &&
			tiles[srcId].getEdge(side) == tiles[destinationId].getEdge(destinationSide) {
			return *tiles[destinationId]
		}
		tiles[destinationId].rotate90()
	}
	// flip, then rotate again
	tiles[destinationId].flip()
	for i := 0; i < 4; i += 1 {
		if tiles[destinationId].matches[destinationSide] == srcId &&
			tiles[srcId].getEdge(side) == tiles[destinationId].getEdge(destinationSide) {
			return *tiles[destinationId]
		}
		tiles[destinationId].rotate90()
	}
	return Tile{}
}

func jigsawString(jigsaw [][]Tile, withSpaces bool) string {
	var result string
	for _, jigsawRow := range jigsaw {
		// for each row in the tile
		for i := 0; i < len(jigsawRow[0].image); i += 1 {
			screenRow := ""
			// for each column in the jigsaw
			for _, col := range jigsawRow {
				screenRow += col.image[i]
				if withSpaces {
					screenRow += " "
				}
			}
			result += screenRow + "\n"
		}
		if withSpaces {
			result += "\n"
		}
	}

	return result
}

func printJigsawIds(jigsaw [][]Tile) {
	for _, jigsawRow := range jigsaw {
		for _, col := range jigsawRow {
			fmt.Printf("%d ", col.id)
		}
		fmt.Println()
	}
}

func getTiles(input []string) map[int]*Tile {
	tiles := make(map[int]*Tile)

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
		tiles[idNum] = &Tile{id: idNum, image: image, matches: make(map[string]int)}

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
