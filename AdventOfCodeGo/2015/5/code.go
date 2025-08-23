package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2015/5/Вводные данные.txt"
	vowels        = "aeiou"
)

var (
	forbiddenWords = []string{"ab", "cd", "pq", "xy"}
)

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var result int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			curStr = strings.TrimSpace(curStr)
			var vowelsNum int
			containsDouble := false
			containsForbiddenWord := false
			var prevRune rune
			for _, c := range curStr {
				//Проверка на наличие гласных
				if strings.Contains(vowels, string(c)) {
					vowelsNum++
				}
				//Проверка на двойную букву
				if prevRune == c {
					containsDouble = true
				}
				//Проверка на плохие слова
				for _, word := range forbiddenWords {
					if prevRune == rune(word[0]) && c == rune(word[1]) {
						containsForbiddenWord = true
						break
					}
				}
				if containsForbiddenWord {
					break
				}
				prevRune = c
			}
			//Соединяем все условия воедино
			if vowelsNum >= 3 && containsDouble && !containsForbiddenWord {
				result++
			}
		} else {
			break
		}
	}
	fmt.Println("Всего хороших строк: ", result)
}

func Part2() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var result int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			curStr = strings.TrimSpace(curStr)
			//Проверяем повторяющуюся букву через один символ
			duplicateSymbols := false
			for i := 2; i < len(curStr); i++ {
				if curStr[i] == curStr[i-2] {
					duplicateSymbols = true
					break
				}
			}
			//Проверяем наличие двух одинаковых пар букв
			duplicateWords := false
			for i := 0; i < len(curStr)-1; i++ {
				if strings.Contains(curStr[:i], curStr[i:i+2]) ||
					strings.Contains(curStr[i+2:], curStr[i:i+2]) {
					duplicateWords = true
					break
				}
			}
			if duplicateSymbols && duplicateWords {
				result++
			}
		} else {
			break
		}
	}
	fmt.Println("Всего хороших строк: ", result)
}
