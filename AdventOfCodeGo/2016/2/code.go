package day2

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2016/2/Вводные данные.txt"
)

func Part1() {
	numpad := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '1', '2', '3', '.'},
		{'.', '4', '5', '6', '.'},
		{'.', '7', '8', '9', '.'},
		{'.', '.', '.', '.', '.'},
	}
	do(numpad, 2, 2)
}

func Part2() {
	numpad := [][]rune{
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '1', '.', '.', '.'},
		{'.', '.', '2', '3', '4', '.', '.'},
		{'.', '5', '6', '7', '8', '9', '.'},
		{'.', '.', 'A', 'B', 'C', '.', '.'},
		{'.', '.', '.', 'D', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
	}
	do(numpad, 1, 3)
}

func do(numpad [][]rune, x, y int) {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	//Переменные
	lines := strings.Split(string(bytes), "\n")
	var result string
	//Для каждой строки:
	for _, line := range lines {
		for _, instruction := range strings.TrimSpace(line) {
			switch instruction {
			case 'U':
				y--
				break
			case 'D':
				y++
				break
			case 'L':
				x--
				break
			case 'R':
				x++
				break
			}
			if numpad[y][x] == '.' {
				switch instruction {
				case 'U':
					y++
					break
				case 'D':
					y--
					break
				case 'L':
					x++
					break
				case 'R':
					x--
					break
				}
			}
		}
		result += string(numpad[y][x])
	}

	fmt.Println("Код от ванной: ", result)
}
