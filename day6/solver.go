package main

import (
	"fmt"
	"os"
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

type Direction struct {
	x int
	y int
}

type DirectionsMap map[uint8]Direction

type RightRotationMap map[uint8]uint8

type Stance struct {
	x         int
	y         int
	direction Direction
	stance    uint8
	obstacled bool
}

var DIRECTIONS_MAP DirectionsMap = DirectionsMap{
	'^': Direction{x: 0, y: -1},
	'v': Direction{x: 0, y: 1},
	'>': Direction{x: 1, y: 0},
	'<': Direction{x: -1, y: 0},
}

var RIGHT_ROTATION_MAP RightRotationMap = RightRotationMap{
	'^': '>',
	'v': '<',
	'>': 'v',
	'<': '^',
}

func getMapAsArray(input string) [][]uint8 {
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
func findCurrentStance(mapa [][]uint8) (stance Stance) {
	for i, row := range mapa {
		for j, position := range row {
			if position == '^' || position == 'v' || position == '>' || position == '<' {
				stance.x = j
				stance.y = i
				stance.stance = position
				stance.direction = DIRECTIONS_MAP[position]
				return
			}
		}
	}
	return
}

func positionOutOfBound(stance Stance, width int, hight int) bool {
	return stance.x < 0 || stance.y < 0 || stance.x > width-1 || stance.y > hight-1
}

func foundObstacle(stance Stance, mapa [][]uint8) bool {
	return mapa[stance.y][stance.x] == '#' || mapa[stance.y][stance.x] == 'O'
}

func getNextStance(currentStance *Stance, mapa [][]uint8) Stance {
	nextStance := Stance{
		x:         currentStance.x + currentStance.direction.x,
		y:         currentStance.y + currentStance.direction.y,
		direction: currentStance.direction,
		stance:    currentStance.stance,
		obstacled: currentStance.obstacled,
	}

	if positionOutOfBound(nextStance, len(mapa[0]), len(mapa)) {
		return nextStance
	}
	if foundObstacle(nextStance, mapa) {
		rotatedDirection := DIRECTIONS_MAP[RIGHT_ROTATION_MAP[currentStance.stance]]
		currentStance.obstacled = true
		nextStance = Stance{x: currentStance.x, y: currentStance.y, direction: rotatedDirection, stance: RIGHT_ROTATION_MAP[currentStance.stance], obstacled: false}
	}
	return nextStance
}

func sweepMap(mapa [][]uint8, initialStance Stance) (mappedMap [][]uint8) {

	mappedMap = make([][]uint8, len(mapa))
	for i := range mapa {
		mappedMap[i] = make([]uint8, len(mapa[i]))
		copy(mappedMap[i], mapa[i])
	}

	currentStance := initialStance
	for !positionOutOfBound(currentStance, len(mapa[0]), len(mapa)) {
		mappedMap[currentStance.y][currentStance.x] = 'X'
		currentStance = getNextStance(&currentStance, mapa)
	}
	return
}

func countSeenPlaces(mapa [][]uint8) (count int) {
	count = 0
	for _, row := range mapa {
		for _, element := range row {
			if element == 'X' {
				count++
			}
		}
	}
	return
}

func solvePart1(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)
	inputAsArray := getMapAsArray(inputString)
	initialStance := findCurrentStance(inputAsArray)
	mappedMap := sweepMap(inputAsArray, initialStance)
	seenPlacesCount := countSeenPlaces(mappedMap)

	return seenPlacesCount
}

// =========================================================

func printArray(array [][]uint8) {
	for y, row := range array {
		for x := range row {
			fmt.Print(string(array[y][x]))
		}
		fmt.Println()
	}
}

func checkLoopForObstruction(mapa [][]uint8, initialStance Stance) (isLoop bool) {
	isLoop = false

	currentStance := initialStance
	previousObstaculatedStances := make([]Stance, 0, 500)
	for !positionOutOfBound(currentStance, len(mapa[0]), len(mapa)) {
		nextStance := getNextStance(&currentStance, mapa)
		if currentStance.obstacled {
			for _, obstacle := range previousObstaculatedStances {
				if obstacle.x == currentStance.x && obstacle.y == currentStance.y && obstacle.stance == currentStance.stance {
					return true
				}
			}
			previousObstaculatedStances = append(previousObstaculatedStances, currentStance)
		}
		currentStance = nextStance
	}
	return
}

func sweepMapWithObstructions(mapa [][]uint8, initialStance Stance) (loopCasesCount int) {
	loopCasesCount = 0
	frontOfGuardInitialStance := getNextStance(&initialStance, mapa)

	for y, row := range mapa {
		for x, _ := range row {
			if foundObstacle(Stance{x: x, y: y}, mapa) || (x == frontOfGuardInitialStance.x && y == frontOfGuardInitialStance.y) {
				continue
			}
			temp := mapa[y][x]
			mapa[y][x] = 'O'
			// printArray(mapa)
			if checkLoopForObstruction(mapa, initialStance) {
				loopCasesCount++
			}
			mapa[y][x] = temp
		}
	}
	return
}

func solvePart2(filePath string) int {
	data, err := os.ReadFile(filePath)
	check(err)
	inputString := string(data)
	inputAsArray := getMapAsArray(inputString)
	initialStance := findCurrentStance(inputAsArray)
	loopCasesCount := sweepMapWithObstructions(inputAsArray, initialStance)

	return loopCasesCount
}

func main() {
	// solution1 := solvePart1(INPUT_FILE_TEST)
	// solution1 := solvePart1(INPUT_FILE)
	// solution2 := solvePart2(INPUT_FILE_TEST)
	solution2 := solvePart2(INPUT_FILE)
	// fmt.Print("Solution part 1: ", solution1)
	fmt.Print("\n")
	fmt.Print("Solution part 2: ", solution2)
	fmt.Print("\n")
}
