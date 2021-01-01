package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile("21.txt")

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

func puzzle1(input []string) int {
	allergenToIngredients := make(map[string][]string)
	allIngredients := make([]string, 0)

	for _, line := range input {
		ingredients, allergens := parseLine(line)
		for _, allergen := range allergens {
			ingredientList, ok := allergenToIngredients[allergen]
			if !ok {
				allergenToIngredients[allergen] = ingredients
			} else {
				// find intersection and keep that
				allergenToIngredients[allergen] = intersect(ingredientList, ingredients)
			}
		}
		allIngredients = append(allIngredients, ingredients...)
	}

	// find set of all ingredients with allergens
	allAllergenIngredients := make(map[string]bool)
	for _, ingredients := range allergenToIngredients {
		for _, ing := range ingredients {
			allAllergenIngredients[ing] = true
		}
	}

	var count int
	for _, ing := range allIngredients {
		_, ok := allAllergenIngredients[ing]
		if !ok {
			count += 1
		}
	}
	return count
}

func intersect(ingredientList, ingredients []string) []string {
	commonIngredients := make([]string, 0)
	for _, ing := range ingredientList {
		if contains(ingredients, ing) {
			commonIngredients = append(commonIngredients, ing)
		}
	}
	return commonIngredients
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func puzzle2(input []string) int {
	return 0
}

// return list of ingredients in each line, followed by list of allergens
func parseLine(line string) ([]string, []string) {
	separated := strings.Split(line, " (contains ")
	//remove trailing comma from allergens half
	separated[1] = separated[1][0 : len(separated[1])-1]

	ingredients := strings.Split(separated[0], " ")
	allergens := strings.Split(separated[1], ", ")

	return ingredients, allergens
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
