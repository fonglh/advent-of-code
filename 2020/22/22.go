package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	//"strings"
)

// Scores both hands and use it as a hash key
type HandKey struct {
	p1Score int
	p2Score int
}

// Game number as used in the question description.
// Each recursive call uses a new game number, that's why it's here as a global.
var gameNum = 0

func main() {
	p1Cards := initDeck("22-testp1.txt")
	p2Cards := initDeck("22-testp2.txt")

	//fmt.Println(puzzle1(p1Cards, p2Cards))
	fmt.Println(puzzle2(p1Cards, p2Cards))
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

func puzzle2(p1Cards, p2Cards []int) int {
	winner, score1, score2 := recursiveCombat(p1Cards, p2Cards)

	if winner == 1 {
		return score1
	} else {
		return score2
	}
}

// returns winning player number, score1, score2
func recursiveCombat(p1Cards, p2Cards []int) (int, int, int) {
	gameNum += 1
	myGameNum := gameNum
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Game", myGameNum, p1Cards, p2Cards)
	seenHands := make(map[HandKey]bool)

	for roundNum := 1; len(p1Cards) > 0 && len(p2Cards) > 0; roundNum += 1 {
		// Check that this is a new card configuration
		p1Score := scoreDeck(p1Cards)
		p2Score := scoreDeck(p2Cards)
		p1Card := p1Cards[0]
		p2Card := p2Cards[0]

		if seenHands[HandKey{p1Score, p2Score}] {
			// game ends immediately with player 1 win
			fmt.Println("Game", myGameNum, "Round", roundNum, "seen before, p1 wins", p1Cards, p2Cards)
			return 1, scoreDeck(p1Cards), scoreDeck(p2Cards)
		}

		// mark this configuration as seen
		seenHands[HandKey{p1Score, p2Score}] = true

		if (len(p1Cards)-1) >= p1Card &&
			(len(p2Cards)-1) >= p2Card {
			// another round of recursive combat
			// passing a copy of the slices is crucial
			p1CardsNext := make([]int, p1Card)
			p2CardsNext := make([]int, p2Card)
			copy(p1CardsNext, p1Cards[1:p1Card+1])
			copy(p2CardsNext, p2Cards[1:p2Card+1])

			winner, _, _ := recursiveCombat(p1CardsNext, p2CardsNext)
			if winner == 1 {
				p1Cards = p1Cards[1:]
				p1Cards = append(p1Cards, p1Card)
				p1Cards = append(p1Cards, p2Card)
				p2Cards = p2Cards[1:]
				fmt.Println("Game", myGameNum, "Round", roundNum, "recursive: p1 wins", p1Cards, p2Cards)
				fmt.Println("********************************************************************************")
			} else {
				p2Cards = p2Cards[1:]
				p2Cards = append(p2Cards, p2Card)
				p2Cards = append(p2Cards, p1Card)
				p1Cards = p1Cards[1:]
				fmt.Println("Game", myGameNum, "Round", roundNum, "recursive: p2 wins", p1Cards, p2Cards)
				fmt.Println("********************************************************************************")
			}
		} else {
			// higher value card wins
			if p1Card > p2Card {
				p1Cards = p1Cards[1:]
				p1Cards = append(p1Cards, p1Card)
				p1Cards = append(p1Cards, p2Card)
				p2Cards = p2Cards[1:]
				fmt.Println("Game", myGameNum, "Round", roundNum, "non recursive: p1 wins", p1Cards, p2Cards)
			} else {
				p2Cards = p2Cards[1:]
				p2Cards = append(p2Cards, p2Card)
				p2Cards = append(p2Cards, p1Card)
				p1Cards = p1Cards[1:]
				fmt.Println("Game", myGameNum, "Round", roundNum, "non recursive: p2 wins", p1Cards, p2Cards)
			}
		}

	}

	if len(p1Cards) > 0 {
		fmt.Println("Game", myGameNum, "p1 wins")
		return 1, scoreDeck(p1Cards), scoreDeck(p2Cards)
	} else {
		fmt.Println("Game", myGameNum, "p2 wins")
		return 2, scoreDeck(p1Cards), scoreDeck(p2Cards)
	}
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
