package day9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const (
	inputDataPath = "2022/9/Вводные данные.txt"
)

type coord struct {
	Y int
	X int
}

func Part1() {
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var head coord
	var headOld coord
	var tail coord
	result := map[coord]bool{tail: true}
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			var x, y int
			switch parts[0] {
			case "U":
				y = 1
				break
			case "D":
				y = -1
				break
			case "L":
				x = -1
				break
			case "R":
				x = 1
				break
			}
			n, _ := strconv.Atoi(parts[1])
			for i := 0; i < n; i++ {
				headOld = head
				head.Y += y
				head.X += x
				if math.Abs(float64(head.X-tail.X)) > 1 || math.Abs(float64(head.Y-tail.Y)) > 1 {
					tail = headOld
					result[tail] = true
				}
			}
		}
		if brkFor {
			break
		}
	}
	fmt.Printf("%d позиций хотя бы один раз посетил хвост веревки", len(result))
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
	var knots []coord
	for i := 0; i < 10; i++ {
		knots = append(knots, coord{})
	}
	result := map[coord]bool{coord{}: true}
	//Для каждой строки:
	//printRopeOnCoordinateSystem(knots)
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			//fmt.Println(parts)
			var x, y int
			switch parts[0] {
			case "U":
				y = 1
				break
			case "D":
				y = -1
				break
			case "L":
				x = -1
				break
			case "R":
				x = 1
				break
			}
			n, _ := strconv.Atoi(parts[1])
			for i := 0; i < n; i++ {
				//Двигаем голову
				knots[0].Y += y
				knots[0].X += x
				//Соответственно двигаем все остальные узлы
				for j := 1; j < len(knots); j++ {
					xDif := knots[j-1].X - knots[j].X
					yDif := knots[j-1].Y - knots[j].Y
					xDifLen := int(math.Abs(float64(xDif)))
					yDifLen := int(math.Abs(float64(yDif)))
					if xDifLen > 1 || yDifLen > 1 {
						if knots[j].Y != knots[j-1].Y {
							knots[j].Y += yDif / yDifLen
						}
						if knots[j].X != knots[j-1].X {
							knots[j].X += xDif / xDifLen
						}
						if j == len(knots)-1 {
							result[knots[j]] = true
						}
					} else {
						break
					}
				}
				printRopeOnCoordinateSystem(knots)
			}
		}
		if brkFor {
			break
		}
	}
	fmt.Printf("%d позиций хотя бы один раз посетил хвост веревки", len(result))
}

func printRopeOnCoordinateSystem(rope []coord) {
	minCoord := coord{
		Y: syscall.INFINITE,
		X: syscall.INFINITE,
	}
	maxCoord := coord{
		Y: -syscall.INFINITE,
		X: -syscall.INFINITE,
	}
	curRopeMap := make(map[coord]rune)
	for i := len(rope) - 1; i >= 0; i-- {
		knot := rope[i]
		if knot.X > maxCoord.X {
			maxCoord.X = knot.X
		}
		if knot.Y > maxCoord.Y {
			maxCoord.Y = knot.Y
		}
		if knot.X < minCoord.X {
			minCoord.X = knot.X
		}
		if knot.Y < minCoord.Y {
			minCoord.Y = knot.Y
		}
		curRopeMap[knot] = rune(i + 'a')
	}
	minCoord.X--
	minCoord.Y--
	maxCoord.X++
	maxCoord.Y++
	for y := maxCoord.Y; y >= minCoord.Y; y-- {
		var curLine []rune
		for x := minCoord.X; x <= maxCoord.X; x++ {
			n, ok := curRopeMap[coord{
				Y: y,
				X: x,
			}]
			if ok {
				curLine = append(curLine, n)
			} else if x == 0 && y == 0 {
				curLine = append(curLine, 'S')
			} else {
				curLine = append(curLine, '.')
			}
		}
		println(string(curLine))
	}
	println()
}
