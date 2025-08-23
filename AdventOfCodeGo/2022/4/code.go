package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/4/Вводные данные.txt"
)

func Part1() {
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var resultSum int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			pair := strings.Split(curStr, ",")
			if isRangeContainsAnotherRange(pair[0], pair[1]) || isRangeContainsAnotherRange(pair[1], pair[0]) {
				resultSum++
			}
		}
		if brkFor {
			break
		}
	}
	fmt.Printf("В %d парах один диапазон полностью содержит другой", resultSum)
}

func isRangeContainsAnotherRange(bigRange, smallRange string) bool {
	bigRangeParts := strings.Split(bigRange, "-")
	smallRangeParts := strings.Split(smallRange, "-")
	minBigRange, _ := strconv.Atoi(bigRangeParts[0])
	maxBigRange, _ := strconv.Atoi(bigRangeParts[1])
	minSmallRange, _ := strconv.Atoi(smallRangeParts[0])
	maxSmallRange, _ := strconv.Atoi(smallRangeParts[1])
	if minBigRange <= minSmallRange && maxBigRange >= maxSmallRange {
		return true
	}
	return false
}

func Part2() {
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var resultSum int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			pair := strings.Split(curStr, ",")
			if !(isRangeFullyBeforeAnotherRange(pair[0], pair[1]) || isRangeFullyBeforeAnotherRange(pair[1], pair[0])) {
				resultSum++
			}
		}
		if brkFor {
			break
		}
	}
	fmt.Printf("В %d парах один диапазон хотя бы немного содержит другой", resultSum)
}

func isRangeFullyBeforeAnotherRange(range1, range2 string) bool {
	range1Parts := strings.Split(range1, "-")
	range2Parts := strings.Split(range2, "-")
	maxRange1, _ := strconv.Atoi(range1Parts[1])
	minRange2, _ := strconv.Atoi(range2Parts[0])
	return maxRange1 < minRange2
}
