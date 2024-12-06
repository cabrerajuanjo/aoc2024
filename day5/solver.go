package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "./input.txt"
var INPUT_FILE_TEST = "./input_test.txt"

func check(e error) {
	if e != nil {
		fmt.Print("handled panic")
		panic(e)
	}
}

func getRulesGroupedBy(rules []string) (rulesGroupedBy map[int][]int) {
	rulesGroupedBy = make(map[int][]int)
	for i := range rules {
		values := strings.Split(rules[i], "|")
		key, err := strconv.Atoi(values[0])
		check(err)
		value, err := strconv.Atoi(values[1])
		check(err)
		_, ok := rulesGroupedBy[key]
		if !ok {
			rulesGroupedBy[key] = make([]int, 0, 100)
		}
		rulesGroupedBy[key] = append(rulesGroupedBy[key], value)
	}
	return
}

func getRules(inputLines []string) (rules []string, rulesEndIndex int) {
	rulesEndIndex = 0
	for i := range inputLines {
		if len(inputLines[i]) < 2 {
			rulesEndIndex = i
			break
		}
	}
	rules = inputLines[:rulesEndIndex]

	return
}

func getUpdates(inputLines []string, rulesEndIndex int) (updates []string) {
	updates = inputLines[rulesEndIndex+1 : len(inputLines)-1]
	return
}

func getCorrectlyOrderedUpdates(updates []string, rulesGroupedBy map[int][]int) (correctlyOrderedUpdates [][]string, incorrectlyOrderedUpdates [][]string) {
	correctlyOrderedUpdates = make([][]string, 0, 1200)
	incorrectlyOrderedUpdates = make([][]string, 0, 1200)

	for i := range updates {
		pages := strings.Split(updates[i], ",")
		outOfOrderFound := false
		for j := 1; j <= len(pages)-1 && !outOfOrderFound; j++ {
			p := len(pages) - j
			pivot, err := strconv.Atoi(pages[p])
			check(err)
			rules, ok := rulesGroupedBy[pivot]
			if !ok {
				continue
			}
			for k := len(pages) - j - 1; k >= 0 && !outOfOrderFound; k-- {
				page, err := strconv.Atoi(pages[k])
				check(err)
				for r := range rules {
					if page == rules[r] {
						outOfOrderFound = true
						break
					}
				}
			}
		}
		if !outOfOrderFound {
			correctlyOrderedUpdates = append(correctlyOrderedUpdates, pages)
		} else {
			incorrectlyOrderedUpdates = append(incorrectlyOrderedUpdates, pages)
		}

	}
	return
}

func sumMiddlePages(updates [][]string) (sum int) {
	sum = 0
	for i := range updates {
		middlePageIndex := int(len(updates[i]) / 2)
		pageNumber, err := strconv.Atoi(updates[i][middlePageIndex])
		check(err)
		sum += pageNumber
	}
	return
}

func solvePart1(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)

	inputLines := strings.Split(inputString, "\n")

	rules, rulesEndIndex := getRules(inputLines)
	updates := getUpdates(inputLines, rulesEndIndex)
	rulesGroupedBy := getRulesGroupedBy(rules)
	correctlyOrderedUpdates, _ := getCorrectlyOrderedUpdates(updates, rulesGroupedBy)
	sumOfMiddlePages := sumMiddlePages(correctlyOrderedUpdates)

	return sumOfMiddlePages
}

// ===============================================================

func findAllUniquePagesInUpdate(update []string) (uniquePagesSet map[int]bool, uniquePagesCount int) {
	uniquePagesSet = make(map[int]bool, 150)
	for _, page := range update {
		pageAsInt, err := strconv.Atoi(page)
		check(err)
		uniquePagesSet[pageAsInt] = false
	}

	uniquePagesCount = 0
	for range uniquePagesSet {
		uniquePagesCount++
	}
	return

}

func getAppliedRules(rules []string, update []string) (appliedRules []string) {
	appliedRules = make([]string, 0, 1000)

	for _, rule := range rules {
		values := strings.Split(rule, "|")
		leftPageFoundInUpdate := false
		rightPageFoundInUpdate := false
		for _, page := range update {
			if values[0] == page {
				leftPageFoundInUpdate = true
			}
			if values[1] == page {
				rightPageFoundInUpdate = true
			}
			if rightPageFoundInUpdate && leftPageFoundInUpdate {
				break
			}
		}
		if rightPageFoundInUpdate && leftPageFoundInUpdate {
			appliedRules = append(appliedRules, rule)
		}
	}
	return
}

func findFirstPage(rules []string, uniquePages map[int]bool) (firstPage int) {
	uniquePagesCopy := make(map[int]bool, 150)
	for k, v := range uniquePages {
		uniquePagesCopy[k] = v
	}
	for i := range rules {
		values := strings.Split(rules[i], "|")
		second, err := strconv.Atoi(values[1])
		check(err)
		// found := uniquePagesSet[first]
		uniquePagesCopy[second] = true
	}

	for k, v := range uniquePagesCopy {
		if !v {
			firstPage = k
		}
	}
	return
}

func findLastPage(rules []string, uniquePages map[int]bool) (lastPage int) {
	uniquePagesCopy := make(map[int]bool, 100)
	for k, v := range uniquePages {
		uniquePagesCopy[k] = v
	}
	for i := range rules {
		values := strings.Split(rules[i], "|")
		first, err := strconv.Atoi(values[0])
		check(err)
		// found := uniquePagesCopySet[first]
		uniquePagesCopy[first] = true
	}

	for k, v := range uniquePagesCopy {
		if !v {
			lastPage = k
		}
	}
	return
}

func hasRepetedValues(slice []int) bool {
	for i, v := range slice {
		for j := i + 1; j < len(slice); j++ {
			if v == slice[j] {
				return true
			}
		}
	}
	return false
}

func pop(slice []int) {
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}
}

func recursePath(rulesGroupedBy map[int][]int, currentPage int, lastPage int, path *[]int, uniquePagesCount int) bool {

	for _, v := range rulesGroupedBy[currentPage] {
		*path = append(*path, v)
		if hasRepetedValues(*path) {
			if len(*path) > 0 {
				*path = (*path)[:len(*path)-1]
			}
			continue
		}
		if v == lastPage && len(*path) == uniquePagesCount {
			return true
		}
		if v == lastPage {
			if len(*path) > 0 {
				*path = (*path)[:len(*path)-1]
			}
			continue
		}
		result := recursePath(rulesGroupedBy, v, lastPage, path, uniquePagesCount)
		if result {
			return true
		} else {
			if len(*path) > 0 {
				*path = (*path)[:len(*path)-1]
			}
		}
	}

	return false
}

func findCompletePath(rulesGroupedBy map[int][]int, firstPage int, lastPage int, uniquePagesCount int) (path []int) {
	path = make([]int, 0, 100)

	path = append(path, firstPage)
	recursePath(rulesGroupedBy, firstPage, lastPage, &path, uniquePagesCount)
	return
}

func correctUpdates(rules []string, updates [][]string) (correctedUpdates [][]string) {
	correctedUpdates = make([][]string, 0, 1200)
	for i, update := range updates {
        appliedRules := getAppliedRules(rules, update)
		uniquePages, uniquePagesCount := findAllUniquePagesInUpdate(update)
		firstPage := findFirstPage(appliedRules, uniquePages)
		lastPage := findLastPage(appliedRules, uniquePages)
		rulesGroupedBy := getRulesGroupedBy(appliedRules)
		compeltePath := findCompletePath(rulesGroupedBy, firstPage, lastPage, uniquePagesCount)
		correctedUpdates = append(correctedUpdates, make([]string, 0, 100))
		for _, pageInPath := range compeltePath {
			for _, pageInUpdate := range update {
				pageNumber, err := strconv.Atoi(pageInUpdate)
				check(err)
				if pageNumber == pageInPath {
					correctedUpdates[i] = append(correctedUpdates[i], pageInUpdate)
				}
			}
		}
	}

	return
}

func solvePart2(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)

	inputLines := strings.Split(inputString, "\n")

	rules, rulesEndIndex := getRules(inputLines)
    rulesGroupedBy := getRulesGroupedBy(rules)
	updates := getUpdates(inputLines, rulesEndIndex)
	_, incorrectlyOrderedUpdates := getCorrectlyOrderedUpdates(updates, rulesGroupedBy)
	correctlyOrderedUpdates := correctUpdates(rules, incorrectlyOrderedUpdates)
	sumOfMiddlePages := sumMiddlePages(correctlyOrderedUpdates)

	return sumOfMiddlePages
	return 0
}

func main() {
	// Test solution: 143
	// solution1 := solvePart1(INPUT_FILE_TEST)
	solution1 := solvePart1(INPUT_FILE)
	// solution2 := solvePart2(INPUT_FILE_TEST)
	solution2 := solvePart2(INPUT_FILE)
	fmt.Print("Solution part 1: ", solution1)
	fmt.Print("\n")
	fmt.Print("Solution part 2: ", solution2)
	fmt.Print("\n")
}
