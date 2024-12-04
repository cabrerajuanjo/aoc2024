package main

import (
	"fmt"
	"os"
	"strings"
	// "strconv"
)

var INPUT_FILE = "./input.txt"
var INPUT_FILE_TEST = "./input_test.txt"

type Direction struct {
	x int8
	y int8
}

func check(e error) {
	if e != nil {
		fmt.Print("handled panic")
		panic(e)
	}
}

func getWordSearchAsArray(input string) [][]uint8 {
	inputLines := strings.Split(input, "\n")
	inputWidth := len(inputLines[0])
	inputHight := len(inputLines)

	if len(inputLines) > 0 {
		inputLines = inputLines[:len(inputLines)-1]
	}

	result := make([][]uint8, 0, inputHight)
	for i := range inputLines {
		result = append(result, make([]uint8, 0, inputWidth))
		for j := 0; j < len(inputLines[i]); j++ {
			result[i] = append(result[i], inputLines[i][j])
		}
	}

	return result
}

func xmas(i, j int, dir Direction, input [][]uint8, word string) bool {
	for idex, jdex, w := i, j, 0; idex >= 0 && idex < len(input[0]) && jdex >= 0 && jdex < len(input); {

		if input[idex][jdex] != word[w] {
			return false
		}

		w++
		if w == len(word) {
			return true
		}
		idex += int(dir.x)
		jdex += int(dir.y)
	}

	return false
}

func solvePart1(filePath string) int {
	const SEARCH_WORD = "XMAS"
	var DIRECTIONS [8]Direction = [8]Direction{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	data, err := os.ReadFile(filePath)
	check(err)

	inputString := string(data)
	inputAsSlice := getWordSearchAsArray(inputString)

	matchCounter := 0
	for i := range inputAsSlice[0] {
		for j := range inputAsSlice {
			for d := range DIRECTIONS {
				if xmas(i, j, DIRECTIONS[d], inputAsSlice, SEARCH_WORD) {
					matchCounter++
				}
			}
		}
	}

	return matchCounter
}

func xdmas(i, j int, dirStart int, input [][]uint8) bool {
	const OUTER_SEQUENCE = "MSSM"
    const CENTER = 'A'
	var DIRECTIONS [7]Direction = [7]Direction{
		{-1, -1},
		{1, -1},
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}
	if input[i][j] != CENTER {
		return false
	}
	for w, d := 0, dirStart; w < len(OUTER_SEQUENCE) && d < len(DIRECTIONS); w, d = w+1, d+1 {
		if input[i+int(DIRECTIONS[d].x)][j+int(DIRECTIONS[d].y)] != OUTER_SEQUENCE[w] {
			return false
		}
	}
	return true
}

func solvePart2(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)

	inputString := string(data)
	inputAsSlice := getWordSearchAsArray(inputString)

	matchCounter := 0
	for i := 1; i < len(inputAsSlice[0])-1; i++ {
		for j := 1; j < len(inputAsSlice)-1; j++ {
			// magic number :(
			for d := 0; d < 4; d++ {
				if xdmas(i, j, d, inputAsSlice) {
					matchCounter++
				}
			}
		}
	}

	return matchCounter
}

func main() {
    // solution1 := solvePart1(INPUT_FILE_TEST)
	solution1 := solvePart1(INPUT_FILE)
	// solution2 := solvePart2(INPUT_FILE_TEST)
	solution2 := solvePart2(INPUT_FILE)
    fmt.Print("Solution part 1: ", solution1)
	fmt.Print("\n")
    fmt.Print("Solution part 2: ", solution2)
	fmt.Print("\n")
}
