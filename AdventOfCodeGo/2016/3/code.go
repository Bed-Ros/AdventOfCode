package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2016/3/Вводные данные.txt"
)

func do(numbers []int) bool {
	for i, number := range numbers {
		var sum int
		for j, n := range numbers {
			if i != j {
				sum += n
			}
		}
		if sum <= number {
			return false
		}
	}
	return true
}

func Part1() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	//Переменные
	lines := strings.Split(string(bytes), "\n")
	var result int
	//Для каждой строки:
	for _, line := range lines {
		//берем числа
		parts := strings.Split(line, " ")
		var numbers []int
		for i := 0; i < len(parts); i++ {
			str := strings.TrimSpace(parts[i])
			if len(str) > 0 {
				num, err := strconv.Atoi(str)
				if err != nil {
					log.Fatalln(err)
				}
				numbers = append(numbers, num)
			}
		}
		//проверяем стороны
		if do(numbers) {
			result++
		}
	}

	fmt.Println("Возможных треугольников: ", result)
}

func Part2() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	//Переменные
	lines := strings.Split(string(bytes), "\n")
	var result int

	//Для каждых трех строк:
	for l := 0; l < len(lines); l += 3 {
		//берем числа
		parts := strings.Split(line, " ")
		var numbers []int
		for i := 0; i < len(parts); i++ {
			str := strings.TrimSpace(parts[i])
			if len(str) > 0 {
				num, err := strconv.Atoi(str)
				if err != nil {
					log.Fatalln(err)
				}
				numbers = append(numbers, num)
			}
		}
		//проверяем стороны
		if do(numbers) {
			result++
		}
	}

	fmt.Println("Возможных треугольников: ", result)
}
