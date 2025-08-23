package day1

import (
	"fmt"
	"log"
	"os"
)

const (
	inputDataPath = "2015/1/Вводные данные.txt"
)

func Part1() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var curFloor int
	//Для каждого символа:
	for _, b := range bytes {
		if b == '(' {
			curFloor++
		} else if b == ')' {
			curFloor--
		}
	}
	fmt.Println("Итоговый этаж: ", curFloor)
}

func Part2() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var curFloor int
	//Для каждого символа:
	for i, b := range bytes {
		if b == '(' {
			curFloor++
		} else if b == ')' {
			curFloor--
		}
		if curFloor < 0 {
			fmt.Println("Санта впервые будет в подвале на символе: ", i+1)
			break
		}
	}
}
