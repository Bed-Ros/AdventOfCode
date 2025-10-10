package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2024/3/Вводные данные.txt"
)

func mulToResult(str string) int {
	parts := strings.Split(strings.Trim(str, "mul()"), ",")
	var partsInt []int
	for _, part := range parts {
		a, _ := strconv.Atoi(part)
		partsInt = append(partsInt, a)
	}
	return partsInt[0] * partsInt[1]
}

func Part1() {
	inputData, _ := os.ReadFile(inputDataPath)
	//находим все инструкции
	inputStr := string(inputData)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(inputStr, -1)
	//делаем арифметику
	var result int
	for _, match := range matches {
		result += mulToResult(match)
	}
	log.Printf("Результат: %d", result)
}

func Part2() {
	inputData, _ := os.ReadFile(inputDataPath)
	//находим все инструкции
	inputStr := string(inputData)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(inputStr, -1)
	//делаем арифметику
	do := true
	var result int
	for _, match := range matches {
		switch match {
		case "do()":
			do = true
			break
		case "don't()":
			do = false
			break
		default:
			if do {
				result += mulToResult(match)
			}
			break
		}
	}
	log.Printf("Результат: %d", result)
}
