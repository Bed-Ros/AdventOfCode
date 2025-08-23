package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/1/Вводные данные.txt"
)

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var maxCaloriesSum int
	var curCaloriesSum int
	brkFor := false
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			number, err := strconv.Atoi(curStr)
			if err != nil {
				log.Fatalln(err)
			}
			curCaloriesSum += number
		} else {
			if maxCaloriesSum < curCaloriesSum {
				maxCaloriesSum = curCaloriesSum
			}
			curCaloriesSum = 0
			if brkFor {
				break
			}
		}
	}
	fmt.Println("Наибольшее количество калорий: ", maxCaloriesSum)
}

func Part2() {
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var caloriesSums []int
	var curCaloriesSum int
	brkFor := false
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			number, err := strconv.Atoi(curStr)
			if err != nil {
				log.Fatalln(err)
			}
			curCaloriesSum += number
		} else {
			caloriesSums = append(caloriesSums, curCaloriesSum)
			curCaloriesSum = 0
			if brkFor {
				break
			}
		}
	}
	//Сортируем суммарные каллории в обьратном порядке
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesSums)))
	//Складываем первые 3
	result := caloriesSums[0] + caloriesSums[1] + caloriesSums[2]
	fmt.Println("Наибольшее суммарное количество каллорий у трех эльфов: ", result)
}
