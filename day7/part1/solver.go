package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "./input.txt"
var INPUT_FILE_TEST = "./input_test.txt"

type Operation struct {
	result   int
	operands []int
}

func powInt(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func check(e error) {
	if e != nil {
		fmt.Print("handled panic")
		panic(e)
	}
}

func parseOperations(inputLines []string) (operations []Operation) {
	operations = make([]Operation, 0, 1000)
	for _, equation := range inputLines {
		sections := strings.Split(equation, " ")
		result, err := strconv.Atoi(strings.Trim(sections[0], ":"))
		check(err)
		operands := make([]int, 0, 50)
		for i := 1; i < len(sections); i++ {
			operand, err := strconv.Atoi(sections[i])
			check(err)
			operands = append(operands, operand)
		}
		operations = append(operations, Operation{result: result, operands: operands})
	}
	return
}

func getCurrentOperandsCombinations(currentCombinationBits int, numberOfOperators int) (currentOperandsCombination []uint8) {
	currentOperandsCombination = make([]uint8, 0, numberOfOperators)
	mask := 0x0001
	for i := 0; i < numberOfOperators; i++ {
		if (currentCombinationBits & mask) == 0 {
			currentOperandsCombination = append(currentOperandsCombination, '+')
		} else {
			currentOperandsCombination = append(currentOperandsCombination, '*')
		}
		mask = mask << 1
	}

	return
}

func testOperations(equation Operation) bool {
	numberOfOperators := len(equation.operands) - 1
	numberOfPossibleCombinations := powInt(2, numberOfOperators)
	for currentCombinationBits := 0; currentCombinationBits < numberOfPossibleCombinations; currentCombinationBits++ {
		currentOperandsCombination := getCurrentOperandsCombinations(currentCombinationBits, numberOfOperators)
		result := equation.operands[0]
		for i := 1; i < len(equation.operands); i++ {
			if currentOperandsCombination[i-1] == '+' {
				result += equation.operands[i]
			} else {
				result *= equation.operands[i]
			}
		}
		if result == equation.result {
			return true
		}
	}
	return false

}

func solvePart1(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)

	inputLines := strings.Split(inputString, "\n")
	if len(inputLines) > 0 {
		inputLines = inputLines[:len(inputLines)-1]
	}
	parsedEquations := parseOperations(inputLines)
	sumOfValidEquations := 0
	for _, equation := range parsedEquations {
		isValidEquation := testOperations(equation)
		if isValidEquation {
			sumOfValidEquations += equation.result
		}
	}

	return sumOfValidEquations
}

func main() {
	// Test solution: 143
	// solution1 := solvePart1(INPUT_FILE_TEST)
	solution1 := solvePart1(INPUT_FILE)
	fmt.Print("Solution part 1: ", solution1)
	fmt.Print("\n")
}
