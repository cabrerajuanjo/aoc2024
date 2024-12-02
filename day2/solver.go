package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
)

var INPUT_FILE1 = "./input1.txt"
var INPUT_FILE2 = "./input2.txt"

var VALUES_SEPARATOR byte = 0x20
var MIN_DIFF int = 1
var MAX_DIFF int = 3

func check(e error) {
	if e != nil {
		fmt.Print("handled panic")
		panic(e)
	}
}

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func isDiffWithinRange(a, b int) bool {
	absDiff := intAbs(a - b)
	return absDiff >= MIN_DIFF && absDiff <= MAX_DIFF
}

func removeItemAtIndex(line []int, index int) []int {
	result := make([]int, 0, len(line))
	for i := 0; i < len(line); i++ {
		if i != index {
			result = append(result, line[i])
		}
	}

	return result
}

func isLineSafe(line *string, deleteIndex int) bool {

	// reportsLines := strings.Split(*line, VALUES_SEPARATOR)
	if deleteIndex >= len(*line) {
		return false
	}

	lineValues := make([]int, 0, 10)
	currentLevel := ""
	for i := 0; i <= len(*line); i++ {
		if i != len(*line) && (*line)[i] != VALUES_SEPARATOR {
			currentLevel += string((*line)[i])
		} else {
			level, err := strconv.Atoi(currentLevel)
			if err != nil {
				fmt.Print("Not a valid value found")
				panic(err)
			}
			lineValues = append(lineValues, level)
			currentLevel = ""
		}
	}

	processedLineValues := removeItemAtIndex(lineValues, deleteIndex)

	if len(processedLineValues) < 2 {
		return true
	}

	const (
		Increassing string = "increassing"
		Decreassing        = "decreassing"
		Unknown            = "unknown"
	)

	previousSlope := Unknown
	for i := 0; i < len(processedLineValues)-1; i++ {
		currentValue := processedLineValues[i]
		nextValue := processedLineValues[i+1]
		currentSlope := Unknown
		if !isDiffWithinRange(currentValue, nextValue) {
			return isLineSafe(line, deleteIndex+1)
		}

		if currentValue < nextValue {
			currentSlope = Increassing
		} else {
			currentSlope = Decreassing
		}

		if previousSlope == Unknown {
			previousSlope = currentSlope
			continue
		}

		if currentSlope != previousSlope {
			return isLineSafe(line, deleteIndex+1)
		}
	}

	return true
}

func getSafeCount(reports *string) int {
	reportsLines := strings.Split(*reports, "\n")
	safeCount := 0
	for i := 0; i < len(reportsLines)-1; i++ {
		if isLineSafe(&(reportsLines[i]), -1) {
			safeCount++
		}
	}
	return safeCount
}

func main() {
	dat1, err := os.ReadFile(INPUT_FILE1)
	check(err)

	inputString1 := string(dat1)
	safeCount := getSafeCount(&inputString1)
	fmt.Print("safeCount ", safeCount)
	fmt.Print("\n")
}
