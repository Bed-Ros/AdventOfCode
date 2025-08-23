package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2015/8/Вводные данные.txt"
)

func countTechSymbols(str string) int {
	str = strings.TrimSpace(str)
	//Проверяем, что строка вводные данные соответствуют задаче
	if str[0] != '"' || str[len(str)-1] != '"' {
		log.Fatalln("wtf", str)
	}
	//И убираем лишние двойные кавычки в начале и конце
	result := 2
	str = str[1 : len(str)-1]
	//Пропускаем маленькие строки, в которых не может быть спец символов
	if len(str) <= 1 {
		return result
	}
	//Проходимся по строке и считаем спец символы
	for i := 1; i < len(str); i++ {
		if str[i-1] == '\\' {
			switch str[i] {
			case '\\', '"':
				result++
				str = str[:i] + "a" + str[i+1:]
				break
			case 'x':
				result += 3
				i += 2
				break
			}
		}
	}
	return result
}

func mainThing(f func(string) int) {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}

	//Результат
	var resultSum int
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			resultSum += f(curStr)
		} else {
			break
		}
	}

	fmt.Println("Всего спец символов: ", resultSum)
}

func Part1() {
	mainThing(countTechSymbols)
}

func Part2() {
	mainThing(func(str string) int {
		str = strings.ReplaceAll(str, "\\", "\\\\")
		str = strings.ReplaceAll(str, "\"", "\\\"")
		return countTechSymbols("\"" + str + "\"")
	})
}
