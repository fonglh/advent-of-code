package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	//"strings"
)

func main() {
	p1Cards := initDeck("22-p1.txt")
	p2Cards := initDeck("22-p2.txt")

	fmt.Println(puzzle1(p1Cards, p2Cards))
	//	fmt.Println(puzzle2(input))
}

func puzzle1(p1Cards, p2Cards []int) int {
	for len(p1Cards) > 0 && len(p2Cards) > 0 {
		p1Cards, p2Cards = combat(p1Cards, p2Cards)
	}

	if len(p1Cards) > 0 {
		fmt.Println("Player 1 won")
		return scoreDeck(p1Cards)
	} else {
		fmt.Println("Player 2 won")
		return scoreDeck(p2Cards)
	}
}

func puzzle2(input []string) int {
	return 0
}

func scoreDeck(cards []int) int {
	score := 0
	deckLength := len(cards)

	for i, card := range cards {
		score += (deckLength - i) * card
	}

	return score
}

func initDeck(filename string) []int {
	input := readFile(filename)
	cards := make([]int, len(input))

	for i, line := range input {
		cardNum, _ := strconv.Atoi(line)
		cards[i] = cardNum
	}
	return cards
}

func combat(p1Cards, p2Cards []int) ([]int, []int) {
	p1Card := p1Cards[0]
	p2Card := p2Cards[0]

	if p1Card > p2Card {
		// first element won't be garbage collected until underlying array is reallocated
		p1Cards = p1Cards[1:]
		p1Cards = append(p1Cards, p1Card)
		p1Cards = append(p1Cards, p2Card)
		p2Cards = p2Cards[1:]
	} else {
		p2Cards = p2Cards[1:]
		p2Cards = append(p2Cards, p2Card)
		p2Cards = append(p2Cards, p1Card)
		p1Cards = p1Cards[1:]
	}

	return p1Cards, p2Cards
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
