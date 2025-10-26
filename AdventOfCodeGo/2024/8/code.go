package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	inputDataPath = "2024/8/Вводные данные.txt"
)

func readInput() (map[rune][]СommonElements.Point, СommonElements.Point) {
	file, err := os.Open(inputDataPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make(map[rune][]СommonElements.Point)

	var curPoint СommonElements.Point
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			cur, ok := result[r]
			if ok {
				result[r] = append(cur, curPoint)
			} else {
				result[r] = []СommonElements.Point{curPoint}
			}
		} else if r == '\n' {
			curPoint.Y++
			curPoint.X = -1
		}
		curPoint.X++
	}
	curPoint.Y++
	return result, curPoint
}

func Part1() {
	antennas, diagramSize := readInput()
	result := make(map[СommonElements.Point]interface{})
	for _, points := range antennas {
		for i1, p1 := range points {
			for i2, p2 := range points {
				if i1 == i2 {
					continue
				}
				vector := p2.Sub(p1)
				x := p1.Sub(vector)
				if x.X >= 0 && x.Y >= 0 && x.X < diagramSize.X && x.Y < diagramSize.Y {
					result[x] = nil
				}
			}
		}
	}
	fmt.Printf("Результат: %d", len(result))
}

func Part2() {
	antennas, diagramSize := readInput()
	result := make(map[СommonElements.Point]interface{})
	for _, points := range antennas {
		for i1, p1 := range points {
			for i2, p2 := range points {
				if i1 == i2 {
					continue
				}
				result[p1] = nil
				vector := p2.Sub(p1)
				x := p1.Sub(vector)
				for {
					if x.X >= 0 && x.Y >= 0 && x.X < diagramSize.X && x.Y < diagramSize.Y {
						result[x] = nil
						x = x.Sub(vector)
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Printf("Результат: %d", len(result))
}
