package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputDataPath = "2022/6/Вводные данные.txt"
)

func mainPart(carriageLength int) {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	result := carriageLength
	var curRunes []rune
	var carriageIndexToReplace int
	//Собираем первый набор
	for a := 0; a < carriageLength; a++ {
		curRune, _, _ := scanner.ReadRune()
		curRunes = append(curRunes, curRune)
	}
	//Для каждого последующего символа:
	for {
		curRune, _, err := scanner.ReadRune()
		if err != nil {
			brkFor = true
		} else {
			//Проверяем текущее состояние
			if allRunesDifferent(curRunes) {
				brkFor = true
			}
			//Заменяем старый символ на новый
			curRunes[carriageIndexToReplace] = curRune
			carriageIndexToReplace++
			if carriageIndexToReplace == carriageLength {
				carriageIndexToReplace = 0
			}
		}
		if brkFor {
			break
		}
		result++
	}
	fmt.Printf("%d символов необходимо обработать, прежде чем будет обнаружен первый маркер начала пакета", result)
}

func allRunesDifferent(runes []rune) bool {
	helpMap := make(map[rune]bool)
	for _, r := range runes {
		_, exist := helpMap[r]
		if exist {
			return false
		}
		helpMap[r] = true
	}
	return true
}

func Part1() {
	mainPart(4)
}

func Part2() {
	mainPart(14)
}
