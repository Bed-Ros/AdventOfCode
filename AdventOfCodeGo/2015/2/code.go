package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2015/2/Вводные данные.txt"
)

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var resultSum int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			curStr = strings.TrimSpace(curStr)
			parts := strings.Split(curStr, "x")
			//преобразование из строки в число
			measurements := make([]int, len(parts))
			for i, part := range parts {
				number, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalln(err)
				}
				measurements[i] = number
			}
			//вычисление сторон
			sides := []int{
				measurements[0] * measurements[1], //l*w
				measurements[1] * measurements[2], //w*h
				measurements[2] * measurements[0], //h*l
			}
			//поиск минимума
			minSide := slices.Min(sides)
			resultSum += minSide
			//добавление к результату
			for _, side := range sides {
				resultSum += 2 * side
			}

		} else {
			break
		}
	}
	fmt.Println("Всего понадобится: ", resultSum)
}

func Part2() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var resultSum int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			curStr = strings.TrimSpace(curStr)
			parts := strings.Split(curStr, "x")
			//преобразование из строки в число
			measurements := make([]int, len(parts))
			for i, part := range parts {
				number, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalln(err)
				}
				measurements[i] = number
			}
			sort.Ints(measurements)
			//добавление мин периметра к результату и объем соотв
			resultSum += 2 * measurements[0]
			resultSum += 2 * measurements[1]
			resultSum += measurements[0] * measurements[1] * measurements[2]
		} else {
			break
		}
	}
	fmt.Println("Всего понадобится: ", resultSum)
}
