package main

import (
	"errors"
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

func parseOperations(inputLines []string) (operations []Operation, longest int) {
	longest = 0
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
		if len(sections)-1 > longest {
			longest = len(sections) - 1
		}
	}
	return
}

func getCurrentOperandsCombinations(currentCombinationBits int, numberOfOperators int) (currentOperandsCombination []uint8, err error) {
	currentOperandsCombination = make([]uint8, 0, numberOfOperators)
	mask := 0x00000001
	for i := 0; i < numberOfOperators; i++ {
		// fmt.Println("mask", mask)
		// fmt.Println(currentCombinationBits&mask)
		first := (currentCombinationBits & mask)
		second := (currentCombinationBits & (mask << 1))
		if first == 0 && second == 0 {
			currentOperandsCombination = append(currentOperandsCombination, '+')
		}
		if first != 0 && second == 0 {
			currentOperandsCombination = append(currentOperandsCombination, '*')
		}
		if first == 0 && second != 0 {
			currentOperandsCombination = append(currentOperandsCombination, '|')
		}
		if first != 0 && second != 0 {
			err = errors.New("Not used")
			break
		}
		mask = mask << 2
	}

	return
}

func concatenateOperators(a, b int) (result int) {
	result, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	check(err)
	return
}

func testOperations(equation Operation, allOperandsCombinations [][]uint8) bool {
	for currentCombinationBits := 0; currentCombinationBits < len(allOperandsCombinations); currentCombinationBits++ {
        currentOperandsCombination := allOperandsCombinations[currentCombinationBits]
		result := equation.operands[0]
		for i := 1; i < len(equation.operands); i++ {
			if currentOperandsCombination[i-1] == '+' {
				result += equation.operands[i]
			} else if currentOperandsCombination[i-1] == '*' {
				result *= equation.operands[i]
			} else {
				result = concatenateOperators(result, equation.operands[i])
			}
		}
		if result == equation.result {
			return true
		}
	}
	return false

}

func solve(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)

	inputLines := strings.Split(inputString, "\n")
	if len(inputLines) > 0 {
		inputLines = inputLines[:len(inputLines)-1]
	}

	parsedEquations, longest := parseOperations(inputLines)
	numberOfPossibleCombinations := powInt(4, longest-1)
    allOperandsCombinations := make([][]uint8, 0, numberOfPossibleCombinations)
	for currentCombinationBits := 0; currentCombinationBits < numberOfPossibleCombinations; currentCombinationBits++ {
		currentOperandsCombination, err := getCurrentOperandsCombinations(currentCombinationBits, longest)
        if err != nil {
            continue
        }
        allOperandsCombinations = append(allOperandsCombinations, currentOperandsCombination)
	}

	sumOfValidEquations := 0
	for _, equation := range parsedEquations {
		isValidEquation := testOperations(equation, allOperandsCombinations)
		if isValidEquation {
			sumOfValidEquations += equation.result
		}
	}

	return sumOfValidEquations
}

func main() {
	// Test solution: 143
	// solution := solve(INPUT_FILE_TEST)
	solution := solve(INPUT_FILE)
	fmt.Print("Solution part 2: ", solution)
	fmt.Print("\n")
	// fmt.Print("Solution part 2: ", solution2)
	fmt.Print("\n")
}
