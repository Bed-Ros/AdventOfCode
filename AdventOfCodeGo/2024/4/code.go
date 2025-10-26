package day4

import (
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2024/4/Вводные данные.txt"
)

func readInput() []string {
	result, _ := os.ReadFile(inputDataPath)
	return strings.Split(string(result), "\n")
}

func thereIsWord(start СommonElements.Point, word string, vector СommonElements.Point, lines []string) bool {
	minPoint := СommonElements.Point{}
	maxPoint := СommonElements.Point{
		X: len(lines[0]) - 1,
		Y: len(lines) - 1,
	}
	curPoint := start
	curSymbol := lines[start.Y][start.X]
	for i := 0; i < len(word); i++ {
		if curSymbol != word[i] {
			return false
		}
		if i == len(word)-1 {
			break
		}
		curPoint = curPoint.Add(vector)
		if !curPoint.Inside(minPoint, maxPoint) {
			return false
		}
		curSymbol = lines[curPoint.Y][curPoint.X]
	}
	return true
}

func Part1() {
	inputData := readInput()
	word := "XMAS"
	vectors := СommonElements.Point{}.Neighbours8().Inside(СommonElements.Point{X: -1, Y: -1}, СommonElements.Point{X: 1, Y: 1})
	var result int
	for y := 0; y < len(inputData); y++ {
		for x := 0; x < len(inputData[y]); x++ {
			if inputData[y][x] == word[0] {
				for _, vector := range vectors {
					if thereIsWord(СommonElements.Point{X: x, Y: y}, word, vector, inputData) {
						result++
					}
				}
			}
		}
	}
	log.Printf("Результат: %d", result)
}

func Part2() {
	inputData := readInput()
	toLeftTop := СommonElements.Point{X: -1, Y: -1}
	toRightTop := СommonElements.Point{X: 1, Y: -1}
	toRightBottom := СommonElements.Point{X: 1, Y: 1}
	toLeftBottom := СommonElements.Point{X: -1, Y: 1}
	var result int
	for y := 0; y < len(inputData); y++ {
		for x := 0; x < len(inputData[y]); x++ {
			if inputData[y][x] == 'A' {
				d1a := thereIsWord(СommonElements.Point{X: x, Y: y}, "AS", toLeftTop, inputData) &&
					thereIsWord(СommonElements.Point{X: x, Y: y}, "AM", toRightBottom, inputData)
				d1b := thereIsWord(СommonElements.Point{X: x, Y: y}, "AM", toLeftTop, inputData) &&
					thereIsWord(СommonElements.Point{X: x, Y: y}, "AS", toRightBottom, inputData)
				d2a := thereIsWord(СommonElements.Point{X: x, Y: y}, "AS", toRightTop, inputData) &&
					thereIsWord(СommonElements.Point{X: x, Y: y}, "AM", toLeftBottom, inputData)
				d2b := thereIsWord(СommonElements.Point{X: x, Y: y}, "AM", toRightTop, inputData) &&
					thereIsWord(СommonElements.Point{X: x, Y: y}, "AS", toLeftBottom, inputData)
				if (d1a || d1b) && (d2a || d2b) {
					result++
				}
			}
		}
	}
	log.Printf("Результат: %d", result)
}
