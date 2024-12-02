package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var INPUT_FILE1 = "./input1.txt"
var INPUT_FILE2 = "./input2.txt"
var VALUES_SEPARATOR = "   "

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

func getValuesFromLine(line string) (string, string) {
	values := strings.Split(line, VALUES_SEPARATOR)

	return values[0], values[1]
}

func getLists(inputLines []string) ([]int, []int) {
	leftList := make([]int, 0, 1000)
	rightList := make([]int, 0, 1000)

	for i := 0; i < len(inputLines)-1; i++ {

		stringValueLeft, stringValueRight := getValuesFromLine(inputLines[i])

		leftVal, err := strconv.Atoi(stringValueLeft)
		check(err)
		leftList = append(leftList, leftVal)

		rightVal, err := strconv.Atoi(stringValueRight)
		check(err)
		rightList = append(rightList, rightVal)
	}
	return leftList, rightList
}

func calculateListsDistance(inputString *string) int {
	inputList := strings.Split(*inputString, "\n")

	leftList, rightList := getLists(inputList)

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	diffSum := 0
	for i := 0; i < len(leftList) || i < len(rightList); i++ {
		diffSum += intAbs(leftList[i] - rightList[i])
	}
	return diffSum
}

func getMapFromListAsKeys(list []int) map[int]bool {
	result := make(map[int]bool, len(list))
	for i := 0; i < len(list); i++ {
		result[list[i]] = true
	}
	return result
}

func similarityScore(inputString *string) int {
	inputList := strings.Split(*inputString, "\n")

	leftList, rightList := getLists(inputList)

	// leftListMap := getMapFromListAsKeys(leftList)

	rightListCount := make(map[int]int, len(rightList))
	for i := 0; i < len(rightList); i++ {
        if _, ok := rightListCount[rightList[i]]; ok {
            rightListCount[rightList[i]] += 1
		} else {
            rightListCount[rightList[i]] = 1
		}
	}

    similarityScoreResult := 0;

	for i := 0; i < len(leftList); i++ {
        if _, ok := rightListCount[rightList[i]]; ok {
            similarityScoreResult += leftList[i] * rightListCount[leftList[i]]
		} 
    }

    return similarityScoreResult 
}

func main() {
	dat1, err := os.ReadFile(INPUT_FILE1)
	check(err)
	dat2, err := os.ReadFile(INPUT_FILE2)
	check(err)

	inputString1 := string(dat1)
	inputString2 := string(dat2)

	diffSum := calculateListsDistance(&inputString1)
	fmt.Print("diffSum ", diffSum, "\n")
	similarityScore := similarityScore(&inputString2)
	fmt.Print("similarityScore ", similarityScore, "\n")

}
