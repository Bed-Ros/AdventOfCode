package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/8/Вводные данные.txt"
)

func readTreeMap() [][]int {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var forestMap [][]int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		curStr = strings.TrimSpace(curStr)
		if len(curStr) > 0 {
			var curLine []int
			for _, symbol := range curStr {
				n, _ := strconv.Atoi(string(rune(symbol)))
				curLine = append(curLine, n)
			}
			forestMap = append(forestMap, curLine)
		}
		if brkFor {
			break
		}
	}
	return forestMap
}

func Part1() {
	forestMap := readTreeMap()
	h := len(forestMap)
	w := len(forestMap[0])
	result := make(map[СommonElements.Point]bool)
	//слева | справа
	for lineIndex := 1; lineIndex < h-1; lineIndex++ {
		line := forestMap[lineIndex]
		leftMax := line[0]
		for i := 1; i < w-1; i++ {
			if line[i] > leftMax {
				result[СommonElements.Point{X: lineIndex, Y: i}] = true
				leftMax = line[i]
			}
		}
		rightMax := line[w-1]
		for i := w - 2; i > 0; i-- {
			if line[i] > rightMax {
				result[СommonElements.Point{X: lineIndex, Y: i}] = true
				rightMax = line[i]
			}
		}
	}
	//сверху | снизу
	for columnIndex := 1; columnIndex < w-1; columnIndex++ {
		topMax := forestMap[0][columnIndex]
		for i := 1; i < h-1; i++ {
			if forestMap[i][columnIndex] > topMax {
				result[СommonElements.Point{X: i, Y: columnIndex}] = true
				topMax = forestMap[i][columnIndex]
			}
		}
		bottomMax := forestMap[h-1][columnIndex]
		for i := h - 2; i > 0; i-- {
			if forestMap[i][columnIndex] > bottomMax {
				result[СommonElements.Point{X: i, Y: columnIndex}] = true
				bottomMax = forestMap[i][columnIndex]
			}
		}
	}
	fmt.Printf("%d деревьев видно за пределами сетки", len(result)+2*w+2*h-4)
}

func Part2() {
	forestMap := readTreeMap()
	h := len(forestMap)
	w := len(forestMap[0])
	var maxScore int
	var test123 []int
	var test456 []int
	for lineIndex := 1; lineIndex < h-1; lineIndex++ {
		for columnIndex := 1; columnIndex < w-1; columnIndex++ {
			curTree := forestMap[lineIndex][columnIndex]
			curScore := 1
			//вниз
			i1 := 1
			for lineIndex+i1 < h {
				if forestMap[lineIndex+i1][columnIndex] >= curTree {
					i1++
					break
				}
				i1++
			}
			curScore *= i1 - 1
			//вверх
			i2 := 1
			for lineIndex-i2 >= 0 {
				if forestMap[lineIndex-i2][columnIndex] >= curTree {
					i2++
					break
				}
				i2++
			}
			curScore *= i2 - 1
			//влево
			i3 := 1
			for columnIndex-i3 >= 0 {
				if forestMap[lineIndex][columnIndex-i3] >= curTree {
					i3++
					break
				}
				i3++
			}
			curScore *= i3 - 1
			//вправо
			i4 := 1
			for columnIndex+i4 < w {
				if forestMap[lineIndex][columnIndex+i4] >= curTree {
					i4++
					break
				}
				i4++
			}
			curScore *= i4 - 1
			if curScore > maxScore {
				maxScore = curScore
				test123 = []int{i1 - 1, i2 - 1, i3 - 1, i4 - 1}
				test456 = []int{lineIndex, columnIndex}
			}
		}
	}
	fmt.Printf("Максимально возможная оценка живописности для данной карты это %d %v %v", maxScore, test123, test456)
}
