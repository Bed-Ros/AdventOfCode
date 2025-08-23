package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/5/Вводные данные.txt"
)

var (
	boxStacks = [][]string{
		{"S", "Z", "P", "D", "L", "B", "F", "C"},
		{"N", "V", "G", "P", "H", "W", "B"},
		{"F", "W", "B", "J", "G"},
		{"G", "J", "N", "F", "L", "W", "C", "S"},
		{"W", "J", "L", "T", "P", "M", "S", "H"},
		{"B", "C", "W", "G", "F", "S"},
		{"H", "T", "P", "M", "Q", "B", "W"},
		{"F", "S", "W", "T"},
		{"N", "C", "R"},
	}
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
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			howMany, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			from--
			to, _ := strconv.Atoi(parts[5])
			to--
			for i := 0; i < howMany; i++ {
				boxStacks[to] = append(boxStacks[to], boxStacks[from][len(boxStacks[from])-1])
				boxStacks[from] = boxStacks[from][:len(boxStacks[from])-1]
			}
		}
		if brkFor {
			break
		}
	}
	//Списываем верхние ящики
	var result string
	for _, stack := range boxStacks {
		result += stack[len(stack)-1]
	}
	fmt.Println("После завершения процедуры перестановки следующие ящики окажутся наверху каждой стопки: " + result)
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

	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			howMany, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			from--
			to, _ := strconv.Atoi(parts[5])
			to--
			boxStacks[to] = append(boxStacks[to], boxStacks[from][len(boxStacks[from])-howMany:]...)
			boxStacks[from] = boxStacks[from][:len(boxStacks[from])-howMany]
		}
		if brkFor {
			break
		}
	}
	//Списываем верхние ящики
	var result string
	for _, stack := range boxStacks {
		result += stack[len(stack)-1]
	}
	fmt.Println("После завершения процедуры перестановки следующие ящики окажутся наверху каждой стопки: " + result)
}
