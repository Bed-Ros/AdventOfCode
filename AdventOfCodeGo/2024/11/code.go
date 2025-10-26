package day11

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	inputData = "6571 0 5851763 526746 23 69822 9 989"
)

var (
	rocksDictionary = make(map[rockStep]int)
)

type rockStep struct {
	Rock   int
	Blinks int
}

func getInput() []int {
	strNums := strings.Fields(inputData)
	result := make([]int, len(strNums))
	for i, strNum := range strNums {
		result[i], _ = strconv.Atoi(strNum)
	}
	return result
}

func numberOfRocks(rock int, blinks int) int {
	if blinks <= 0 {
		return 1
	}

	curStep := rockStep{
		Rock:   rock,
		Blinks: blinks,
	}

	result, ok := rocksDictionary[curStep]
	if ok {
		return result
	}

	doRules := func() int {
		if rock == 0 {
			return numberOfRocks(1, blinks-1)
		}
		strRock := strconv.Itoa(rock)
		if len(strRock)%2 == 0 {
			halfIndex := len(strRock) / 2
			strRock1 := strRock[:halfIndex]
			strRock2 := strRock[halfIndex:]
			rock1, _ := strconv.Atoi(strRock1)
			rock2, _ := strconv.Atoi(strRock2)
			return numberOfRocks(rock1, blinks-1) + numberOfRocks(rock2, blinks-1)
		}
		return numberOfRocks(rock*2024, blinks-1)
	}
	result = doRules()
	rocksDictionary[curStep] = result
	return result
}

func mainFunc(blinks int) {
	input := getInput()
	var result int

	for _, n := range input {
		result += numberOfRocks(n, blinks)
	}

	fmt.Printf("Результат: %d", result)
}

func Part1() {
	mainFunc(25)
}

func Part2() {
	mainFunc(75)
}
