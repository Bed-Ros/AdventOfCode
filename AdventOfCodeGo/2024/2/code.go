package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2024/2/Вводные данные.txt"
)

func readInputData() [][]int {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var result [][]int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			var line []int
			for _, part := range parts {
				n, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalln(err)
				}
				line = append(line, n)
			}
			result = append(result, line)
		} else {
			break
		}
	}
	return result
}

func checkLine(line []int) bool {
	var sign *int
	for i := 1; i < len(line); i++ {
		difference := line[i-1] - line[i]
		differenceAbs := math.Abs(float64(difference))
		if differenceAbs < 1 || differenceAbs > 3 {
			return false
		}
		curSign := int(float64(difference) / differenceAbs)
		if sign == nil {
			sign = &curSign
			continue
		}
		if *sign != curSign {
			return false
		}
	}
	return true
}

func Part1() {
	inputData := readInputData()
	//проверяем
	var result int
	for _, line := range inputData {
		if checkLine(line) {
			result++
		}
	}
	log.Printf("%d отчетов безопасны", result)
}

func Part2() {
	inputData := readInputData()
	//проверяем
	var result int
	for _, line := range inputData {
		curLineIsGood := checkLine(line)
		for i := 0; i < len(line); i++ {
			if curLineIsGood {
				break
			}
			trimLine := make([]int, len(line))
			copy(trimLine, line)
			trimLine = append(trimLine[:i], trimLine[i+1:]...)
			curLineIsGood = checkLine(trimLine)
		}
		if curLineIsGood {
			result++
		}
	}
	log.Printf("%d отчетов безопасны", result)
}
