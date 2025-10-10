package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2024/1/Вводные данные.txt"
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
			for i, part := range parts {
				n, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalln(err)
				}
				if len(result) > i {
					result[i] = append(result[i], n)
				} else {
					result = append(result, []int{n})
				}
			}
		} else {
			break
		}
	}
	return result
}

func Part1() {
	inputData := readInputData()
	//сортируем каждый список
	for _, curList := range inputData {
		sort.Ints(curList)
	}
	//считаем
	var resultSum int
	for i := 0; i < len(inputData[0]); i++ {
		resultSum += int(math.Abs(float64(inputData[0][i] - inputData[1][i])))

	}
	log.Printf("Итоговая сумма: %d", resultSum)
}

func Part2() {
	inputData := readInputData()
	//считаем
	var resultSum int
	for _, left := range inputData[0] {
		for _, right := range inputData[1] {
			if left == right {
				resultSum += left
			}
		}
	}
	log.Printf("Итоговая сумма: %d", resultSum)
}
