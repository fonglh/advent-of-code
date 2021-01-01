package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	input := readFile("21.txt")

	p1Answer, allergenToIngredients := puzzle1(input)

	fmt.Println(p1Answer)
	fmt.Println(puzzle2(allergenToIngredients))
}

func puzzle1(input []string) (int, map[string][]string) {
	allergenToIngredients := make(map[string][]string)
	// just a combined list of all the ingredients to count the number of times ingredients cannot have allergens.
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

	return count, allergenToIngredients
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

func puzzle2(allergenToIngredients map[string][]string) string {
	allergenMap := make(map[string]string)

	for len(allergenToIngredients) > 0 {
		var knownIngredient string
		for allergen, ing := range allergenToIngredients {
			// when there is 1 ingredient for that allergen, the mapping is certain.
			// add it to the map
			if len(ing) == 1 {
				knownIngredient = ing[0]
				allergenMap[allergen] = knownIngredient
				delete(allergenToIngredients, allergen)
				break
			}
		}

		// remove known dangerous ingredient from all the remaining lists
		// this should result in another allergen having only 1 possible ingredient
		for allergen, ingList := range allergenToIngredients {
			newIngList := make([]string, 0)
			for _, ing := range ingList {
				if ing != knownIngredient {
					newIngList = append(newIngList, ing)
				}
			}
			allergenToIngredients[allergen] = newIngList
		}
	}

	// get all the allergens from the map and sort the allergen list
	allergenList := make([]string, 0)
	for allergen, _ := range allergenMap {
		allergenList = append(allergenList, allergen)
	}
	sort.Strings(allergenList)

	// create the list of dangerous ingredients sorted by the allergen list
	dangerousList := make([]string, 0)
	for _, allergen := range allergenList {
		dangerousList = append(dangerousList, allergenMap[allergen])
	}

	return strings.Join(dangerousList, ",")
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
