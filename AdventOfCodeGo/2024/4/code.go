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

type Point struct {
	X, Y int
}

func (p Point) Inside(min, max Point) bool {
	return p.X >= min.X && p.Y >= min.Y && p.X <= max.X && p.Y <= max.Y
}

func (p Point) GetNeighbours(min, max Point) []Point {
	var result []Point
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			newPoint := Point{X: p.X + offsetX, Y: p.Y + offsetY}
			if newPoint.Inside(min, max) {
				result = append(result, newPoint)
			}
		}
	}
	return result
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

func thereIsWord(start Point, word string, vector Point, lines []string) bool {
	minPoint := Point{}
	maxPoint := Point{
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
	vectors := Point{}.GetNeighbours(Point{-1, -1}, Point{1, 1})
	var result int
	for y := 0; y < len(inputData); y++ {
		for x := 0; x < len(inputData[y]); x++ {
			if inputData[y][x] == word[0] {
				for _, vector := range vectors {
					if thereIsWord(Point{X: x, Y: y}, word, vector, inputData) {
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
	toLeftTop := Point{X: -1, Y: -1}
	toRightTop := Point{X: 1, Y: -1}
	toRightBottom := Point{X: 1, Y: 1}
	toLeftBottom := Point{X: -1, Y: 1}
	var result int
	for y := 0; y < len(inputData); y++ {
		for x := 0; x < len(inputData[y]); x++ {
			if inputData[y][x] == 'A' {
				d1a := thereIsWord(Point{X: x, Y: y}, "AS", toLeftTop, inputData) && thereIsWord(Point{X: x, Y: y}, "AM", toRightBottom, inputData)
				d1b := thereIsWord(Point{X: x, Y: y}, "AM", toLeftTop, inputData) && thereIsWord(Point{X: x, Y: y}, "AS", toRightBottom, inputData)
				d2a := thereIsWord(Point{X: x, Y: y}, "AS", toRightTop, inputData) && thereIsWord(Point{X: x, Y: y}, "AM", toLeftBottom, inputData)
				d2b := thereIsWord(Point{X: x, Y: y}, "AM", toRightTop, inputData) && thereIsWord(Point{X: x, Y: y}, "AS", toLeftBottom, inputData)
				if (d1a || d1b) && (d2a || d2b) {
					result++
				}
			}
		}
	}
	log.Printf("Результат: %d", result)
}
