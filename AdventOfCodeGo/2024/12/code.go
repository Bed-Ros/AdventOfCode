package day12

import (
	"AdventOfCode/CommonElements"
	"fmt"
	"os"
	"strings"
)

const (
	inputDataPath = "2024/12/Вводные данные.txt"
)

func readInput() []string {
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(bytes))
}

func Part1() {
	lines := readInput()
	h := len(lines)
	w := len(lines[0])
	completedSectors := make(map[CommonElements.Point]any)
	curPoints := []CommonElements.Point{{0, 0}}
	var nextPoints []CommonElements.Point
	differentPoints := make(map[CommonElements.Point]any)
	var curPerimeter int
	curSize := 1
	var result int

	for { //len(completedSectors) < h*w {
		size := curSize
		for _, point := range curPoints {
			_, ok := completedSectors[point]
			if ok {
				continue
			}
			possibleNeighbours := point.Neighbours4()
			neighbours := possibleNeighbours.Inside(CommonElements.Point{}, CommonElements.Point{X: h - 1, Y: w - 1})
			curPerimeter += len(possibleNeighbours) - len(neighbours)
			for _, neighbour := range neighbours {
				if lines[point.Y][point.X] == lines[neighbour.Y][neighbour.X] {
					curSize++
					nextPoints = append(nextPoints, neighbour)
				} else {
					curPerimeter++
					differentPoints[neighbour] = nil
				}
			}
			completedSectors[point] = nil
		}
		if size == curSize {
			curPoints = nil
			for point, _ := range differentPoints {
				_, ok := completedSectors[point]
				if ok {
					continue
				}
				curPoints = append(curPoints, point)
			}
			if len(curPoints) == 0 {
				break
			}
			differentPoints = make(map[CommonElements.Point]any)
			for _, point := range curPoints {
				differentPoints[point] = nil
			}
			curPoints = []CommonElements.Point{curPoints[0]}
			result += curPerimeter * curSize
			curPerimeter = 0
			curSize = 0
		} else {
			curPoints = nextPoints
		}
		nextPoints = nil
	}

	fmt.Printf("Результат: %d", result)
}

func Part2() {
	//lines := readInput()
	//
	//fmt.Printf("Результат: %d", result)
}
