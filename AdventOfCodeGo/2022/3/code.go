package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	inputDataPath = "2022/3/Вводные данные.txt"
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
	var prioritiesSum int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			prioritiesSum += bagsScore(curStr)
		}
		if brkFor {
			break
		}
	}
	fmt.Println("Итоговая сумма приоритетов: ", prioritiesSum)
}

func bagsScore(bag string) int {
	bag1 := bag[:len(bag)/2]
	bag2 := bag[len(bag)/2:]
	for _, item1 := range bag1 {
		for _, item2 := range bag2 {
			if item1 == item2 {
				if unicode.IsLower(item1) {
					return int(item1) - 'a' + 1
				} else {
					return int(item1) - 'A' + 27
				}
			}
		}
	}
	return 0
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
	var prioritiesSum int
	var curGroup []string
	i := 1
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			curGroup = append(curGroup, curStr)
			if i%3 == 0 {
				prioritiesSum += elfGroupBadge(curGroup)
				curGroup = nil
			}
		}
		if brkFor {
			break
		}
		i++
	}
	fmt.Println("Итоговая сумма приоритетов значков: ", prioritiesSum)
}

func elfGroupBadge(bags []string) int {
	var checkLists [][]int
	for _, bag := range bags {
		curCheckList := make([]int, 54)
		for _, item := range bag {
			if unicode.IsLower(item) {
				curCheckList[int(item)-'a'] = 1
			} else {
				curCheckList[int(item)-'A'+26] = 1
			}
		}
		checkLists = append(checkLists, curCheckList)
	}
	for i := 0; i < 54; i++ {
		isBadge := 0
		for _, list := range checkLists {
			isBadge += list[i]
		}
		if isBadge == len(checkLists) {
			return i + 1
		}
	}
	return 0
}
