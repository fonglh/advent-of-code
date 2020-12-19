package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("18.txt")

	fmt.Println(puzzle1(input))
	fmt.Println(puzzle2(input))
}

// https://algorithms.tutorialhorizon.com/evaluation-of-infix-expressions/
// https://www.geeksforgeeks.org/expression-evaluation/
func puzzle1(input []string) int {
	var totalSum int

	for _, line := range input {
		operandStack := make([]int, 0)
		operatorStack := make([]string, 0)
		var topIndex int
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		tokens := strings.Split(line, " ")

		for _, char := range tokens {
			//fmt.Println(char)
			number, err := strconv.Atoi(char)
			if err == nil { // operand
				operandStack = append(operandStack, number)
			} else { // operator
				if char == "(" || len(operatorStack) == 0 {
					operatorStack = append(operatorStack, char)
				} else if char == ")" {
					// pop until "("

					for operatorStack[len(operatorStack)-1] != "(" {
						// pop operator
						topIndex = len(operatorStack) - 1
						operator := operatorStack[topIndex]
						operatorStack = operatorStack[:topIndex]
						// pop 2 operands
						topIndex = len(operandStack) - 1
						operand1 := operandStack[topIndex]
						operandStack = operandStack[:topIndex]
						topIndex = len(operandStack) - 1
						operand2 := operandStack[topIndex]
						operandStack = operandStack[:topIndex]
						if operator == "+" {
							operandStack = append(operandStack, operand1+operand2)
						} else if operator == "*" {
							operandStack = append(operandStack, operand1*operand2)
						}
					}
					// pop "("
					topIndex = len(operatorStack) - 1
					operatorStack = operatorStack[:topIndex]
				} else if operatorStack[len(operatorStack)-1] != "(" {
					// pop operator
					topIndex = len(operatorStack) - 1
					operator := operatorStack[topIndex]
					operatorStack = operatorStack[:topIndex]
					// pop 2 operands
					topIndex = len(operandStack) - 1
					operand1 := operandStack[topIndex]
					operandStack = operandStack[:topIndex]
					topIndex = len(operandStack) - 1
					operand2 := operandStack[topIndex]
					operandStack = operandStack[:topIndex]
					if operator == "+" {
						operandStack = append(operandStack, operand1+operand2)
					} else if operator == "*" {
						operandStack = append(operandStack, operand1*operand2)
					}
					operatorStack = append(operatorStack, char)
				} else {
					operatorStack = append(operatorStack, char)
				}
			}
			//fmt.Println(operandStack)
			//fmt.Println(operatorStack)
			//fmt.Println("--------------------")
		}

		for len(operatorStack) > 0 {
			// pop operator
			topIndex = len(operatorStack) - 1
			operator := operatorStack[topIndex]
			operatorStack = operatorStack[:topIndex]
			// pop 2 operands
			topIndex = len(operandStack) - 1
			operand1 := operandStack[topIndex]
			operandStack = operandStack[:topIndex]
			topIndex = len(operandStack) - 1
			operand2 := operandStack[topIndex]
			operandStack = operandStack[:topIndex]
			if operator == "+" {
				operandStack = append(operandStack, operand1+operand2)
			} else if operator == "*" {
				operandStack = append(operandStack, operand1*operand2)
			}
		}
		totalSum += operandStack[0]
		//fmt.Println(operandStack[0])
	}

	return totalSum
}

func puzzle2(input []string) int {
	var totalSum int

	for _, line := range input {
		operandStack := make([]int, 0)
		operatorStack := make([]string, 0)
		var topIndex int
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		tokens := strings.Split(line, " ")

		for _, char := range tokens {
			//fmt.Println(char)
			number, err := strconv.Atoi(char)
			if err == nil { // operand
				operandStack = append(operandStack, number)
			} else { // operator
				if char == "(" || len(operatorStack) == 0 {
					operatorStack = append(operatorStack, char)
				} else if char == ")" {
					// pop until "("

					for operatorStack[len(operatorStack)-1] != "(" {
						// pop operator
						topIndex = len(operatorStack) - 1
						operator := operatorStack[topIndex]
						operatorStack = operatorStack[:topIndex]
						// pop 2 operands
						topIndex = len(operandStack) - 1
						operand1 := operandStack[topIndex]
						operandStack = operandStack[:topIndex]
						topIndex = len(operandStack) - 1
						operand2 := operandStack[topIndex]
						operandStack = operandStack[:topIndex]
						if operator == "+" {
							operandStack = append(operandStack, operand1+operand2)
						} else if operator == "*" {
							operandStack = append(operandStack, operand1*operand2)
						}
					}
					// pop "("
					topIndex = len(operatorStack) - 1
					operatorStack = operatorStack[:topIndex]
				} else if operatorStack[len(operatorStack)-1] == "+" {
					// pop operator
					topIndex = len(operatorStack) - 1
					operator := operatorStack[topIndex]
					operatorStack = operatorStack[:topIndex]
					// pop 2 operands
					topIndex = len(operandStack) - 1
					operand1 := operandStack[topIndex]
					operandStack = operandStack[:topIndex]
					topIndex = len(operandStack) - 1
					operand2 := operandStack[topIndex]
					operandStack = operandStack[:topIndex]
					if operator == "+" {
						operandStack = append(operandStack, operand1+operand2)
					} else if operator == "*" {
						operandStack = append(operandStack, operand1*operand2)
					}
					operatorStack = append(operatorStack, char)
				} else {
					operatorStack = append(operatorStack, char)
				}
			}
			//fmt.Println(operandStack)
			//fmt.Println(operatorStack)
			//fmt.Println("--------------------")
		}

		for len(operatorStack) > 0 {
			// pop operator
			topIndex = len(operatorStack) - 1
			operator := operatorStack[topIndex]
			operatorStack = operatorStack[:topIndex]
			// pop 2 operands
			topIndex = len(operandStack) - 1
			operand1 := operandStack[topIndex]
			operandStack = operandStack[:topIndex]
			topIndex = len(operandStack) - 1
			operand2 := operandStack[topIndex]
			operandStack = operandStack[:topIndex]
			if operator == "+" {
				operandStack = append(operandStack, operand1+operand2)
			} else if operator == "*" {
				operandStack = append(operandStack, operand1*operand2)
			}
		}
		totalSum += operandStack[0]
		//fmt.Println(operandStack[0])
	}

	return totalSum
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
