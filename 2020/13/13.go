package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var one = big.NewInt(1)

// Taken from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main() {
	input := readFile("13.txt")

	fmt.Println(puzzle2a(input))
}

func puzzle2a(input []string) big.Int {
	busOffsetMap := make(map[int]int)

	for offset, busId := range input {
		if busId != "x" {
			busIdInt, _ := strconv.Atoi(busId)
			// Calculate remainder from offset
			// Subtract divisor (the mod number) from offset until it's less than the mod,
			// then take the "complement".
			for ; offset >= busIdInt; offset -= busIdInt {
			}
			busOffsetMap[busIdInt] = busIdInt - offset
		}
	}

	n := make([]*big.Int, len(busOffsetMap))
	a := make([]*big.Int, len(busOffsetMap))

	index := 0
	for divisor, remainder := range busOffsetMap {
		n[index] = big.NewInt(int64(divisor))
		a[index] = big.NewInt(int64(remainder))
		index += 1
	}

	answer, _ := crt(a, n)
	return *answer
}

//full brute force
func puzzle2(input []string) int64 {
	var timestamp int64

	busOffsetMap := make(map[int64]int64)

	for offset, busId := range input {
		if busId != "x" {
			busIdInt, _ := strconv.Atoi(busId)
			busOffsetMap[int64(busIdInt)] = int64(offset)
		}
	}

	for {
		valid := true
		for busId, offset := range busOffsetMap {
			if (timestamp+offset)%busId != 0 {
				valid = false
				break
			}
		}
		if valid {
			return timestamp
		}
		timestamp += 1
	}
	return timestamp
}

// This readFile function is different from most days.
// It takes only the 2nd line of the input file and splits it by ","
func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rawInput []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawInput = append(rawInput, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var input []string
	input = strings.Split(rawInput[1], ",")

	return input
}
