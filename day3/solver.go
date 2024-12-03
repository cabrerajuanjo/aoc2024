package main

import (
	"fmt"
	"os"
	"strconv"
)

var INPUT_FILE1 = "./input1.txt"
var INPUT_FILE2 = "./input2.txt"

type EnablerParseResult struct {
	result     bool
	errorIndex int
}

type MulParseResult struct {
	result     int
	errorIndex int
}

func check(e error) {
	if e != nil {
		fmt.Print("handled panic")
		panic(e)
	}
}

func parseValue(input *string, index int) (int, int) {
	digits := ""
	for i := index; i < len(*input); i++ {
		if (*input)[i] >= '0' && (*input)[i] <= '9' {
			digits += string((*input)[i])
		} else {
			break
		}
	}
	number, err := strconv.Atoi(digits)
	if err != nil {
		fmt.Print("Not a valid value found")
		panic(err)
	}
	return number, len(digits) - 1

}

func parseEnabler(input *string, currentIndex int, template string) (result EnablerParseResult) {
	for i, j := currentIndex, 0; i < len(*input) && j < len(template); i, j = i+1, j+1 {
		if (*input)[i] != template[j] {
			result.result = false
			result.errorIndex = i
			return
		}
	}
	result.result = true
	result.errorIndex = -1
	return
}

func parseMul(input *string, currentIndex int) (result MulParseResult) {
	result = MulParseResult{0, -1}
	resultAcc := 1
	const TEMPLATE string = "mul(_,_)"

	var lastIndex int
	for i, j := currentIndex, 0; i < len(*input) && j < len(TEMPLATE); i, j = i+1, j+1 {
		lastIndex = i
		if TEMPLATE[j] == '_' {
			currentValue, offset := parseValue(input, i)
			i += offset
			resultAcc *= currentValue
			continue
		}

		if (*input)[i] != TEMPLATE[j] {
			result.result = 0
			result.errorIndex = i
			return result
		}
	}

	result.result = resultAcc
	result.errorIndex = lastIndex
	return
}

func calculateSumOfMuls(input *string) int {
	result := 0
	isMulEnabled := true
	for i := 0; i < len(*input); i++ {
		doParseResult := parseEnabler(input, i, "do()")
		dontParseResult := parseEnabler(input, i, "don't()")
		if doParseResult.result {
			isMulEnabled = true
		}
		if dontParseResult.result {
			isMulEnabled = false
		}
		if !isMulEnabled {
			continue
		}
		parseMulResult := parseMul(input, i)
		i = parseMulResult.errorIndex
		result += parseMulResult.result
	}
	return result
}

func main() {
	dat1, err := os.ReadFile(INPUT_FILE1)
	check(err)

	inputString1 := string(dat1)
	sumOfMuls := calculateSumOfMuls(&inputString1)
	fmt.Print("Sum of muls ", sumOfMuls)
	fmt.Print("\n")
}
