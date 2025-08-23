package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2022/2/Вводные данные.txt"
)

func Part1() {
	//Создадим карту очков
	pointsMap := map[string]map[string]int{
		"A": {"X": 3, "Y": 6, "Z": 0},
		"B": {"X": 0, "Y": 3, "Z": 6},
		"C": {"X": 6, "Y": 0, "Z": 3},
	}
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var points int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			//Добавляем очки за выбор фигуры
			switch parts[1] {
			case "X":
				points += 1
				break
			case "Y":
				points += 2
				break
			case "Z":
				points += 3
				break
			}
			//Добавляем очки за исход игры
			points += pointsMap[parts[0]][parts[1]]
		}
		if brkFor {
			break
		}
	}
	fmt.Println("Всего получу очков: ", points)
}

func Part2() {
	//Создадим карту очков
	figureMap := map[string]map[string]int{
		"A": {"X": 3, "Y": 1, "Z": 2},
		"B": {"X": 1, "Y": 2, "Z": 3},
		"C": {"X": 2, "Y": 3, "Z": 1},
	}
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var points int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			//Добавляем очки за исход игры
			switch parts[1] {
			case "X":
				points += 0
				break
			case "Y":
				points += 3
				break
			case "Z":
				points += 6
				break
			}
			//Добавляем очки за выбор фигуры
			points += figureMap[parts[0]][parts[1]]
		}
		if brkFor {
			break
		}
	}
	fmt.Println("Всего получу очков: ", points)
}
