package day10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

const (
	inputDataPath = "2024/10/Вводные данные.txt"
)

func readInput() [][]int {
	file, err := os.Open(inputDataPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result [][]int

	var curLine []int
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if unicode.IsDigit(r) {
			n, _ := strconv.Atoi(string(r))
			curLine = append(curLine, n)
		} else if r == '\n' {
			result = append(result, curLine)
			curLine = nil
		}
	}
	result = append(result, curLine)
	return result
}

func findTops(input [][]int, start СommonElements.Point) []СommonElements.Point {
	curHeight := input[start.Y][start.X]
	if curHeight == 9 {
		return []СommonElements.Point{start}
	}
	neighbours := []СommonElements.Point{
		{start.X - 1, start.Y},
		{start.X + 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var result []СommonElements.Point
	for _, neighbour := range neighbours {
		if neighbour.Inside(СommonElements.Point{}, СommonElements.Point{X: len(input[0]) - 1, Y: len(input) - 1}) &&
			input[neighbour.Y][neighbour.X] == curHeight+1 {
			result = append(result, findTops(input, neighbour)...)
		}
	}
	return result
}

func Part1() {
	input := readInput()
	var result int
	for y, line := range input {
		for x, cell := range line {
			if cell == 0 {
				points := make(map[СommonElements.Point]interface{})
				for _, p := range findTops(input, СommonElements.Point{X: x, Y: y}) {
					points[p] = nil
				}
				result += len(points)
			}
		}
	}
	fmt.Printf("Результат: %d", result)
}

func Part2() {
	input := readInput()
	var result int
	for y, line := range input {
		for x, cell := range line {
			if cell == 0 {
				result += len(findTops(input, СommonElements.Point{X: x, Y: y}))
			}
		}
	}
	fmt.Printf("Результат: %d", result)
}
