package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/10/Вводные данные.txt"
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
	X := 1
	var result int
	var curCycle int
	nextCycle := 20
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			var lastNum int
			switch parts[0] {
			case "addx":
				curCycle += 2
				lastNum, _ = strconv.Atoi(parts[1])
				X += lastNum
				break
			case "noop":
				curCycle += 1
				break
			}
			if curCycle >= nextCycle {
				if curCycle == nextCycle {
					result += nextCycle * X
				} else if curCycle > nextCycle {
					result += nextCycle * (X - lastNum)
				}
				println(nextCycle, result)
				nextCycle += 40
			}
		}
		if brkFor {
			break
		}
	}
	fmt.Printf("Сумма уровней сигнала равна %d", result)
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
	X := 1
	var result []rune
	var curCycle int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			var lastNum int
			var addCycles int
			switch parts[0] {
			case "addx":
				addCycles = 2
				lastNum, _ = strconv.Atoi(parts[1])
				break
			case "noop":
				addCycles = 1
				break
			}
			for i := 0; i < addCycles; i++ {
				if X-1 <= curCycle && curCycle <= X+1 {
					result = append(result, '#')
				} else {
					result = append(result, '.')
				}
				curCycle++
				if curCycle >= 40 {
					curCycle = curCycle % 40
				}
			}
			X += lastNum
		}
		if brkFor {
			break
		}
	}
	for lineIndex := 0; lineIndex < 6; lineIndex++ {
		println(string(result[lineIndex*40 : (lineIndex+1)*40]))
	}
}
