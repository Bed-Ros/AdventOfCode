package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2015/6/Вводные данные.txt"
)

type Point struct {
	X, Y int
}

type LightsGrid struct {
	Lights [1000][1000]bool
}

func (g *LightsGrid) Toggle(begin, end Point) {
	for x := begin.X; x <= end.X; x++ {
		for y := begin.Y; y <= end.Y; y++ {
			g.Lights[y][x] = !g.Lights[y][x]
		}
	}
}

func (g *LightsGrid) Set(value bool, begin, end Point) {
	for x := begin.X; x <= end.X; x++ {
		for y := begin.Y; y <= end.Y; y++ {
			g.Lights[y][x] = value
		}
	}
}

func (g *LightsGrid) NumberOfLitLights() int {
	var result int
	for x := 0; x < len(g.Lights[0]); x++ {
		for y := 0; y < len(g.Lights); y++ {
			if g.Lights[y][x] {
				result++
			}
		}
	}
	return result
}

type RightLightsGrid struct {
	Lights [1000][1000]int
}

func (g *RightLightsGrid) Add(value int, begin, end Point) {
	for x := begin.X; x <= end.X; x++ {
		for y := begin.Y; y <= end.Y; y++ {
			g.Lights[y][x] += value
			if g.Lights[y][x] < 0 {
				g.Lights[y][x] = 0
			}
		}
	}
}

func (g *RightLightsGrid) TotalBrightness() int {
	var result int
	for x := 0; x < len(g.Lights[0]); x++ {
		for y := 0; y < len(g.Lights); y++ {
			result += g.Lights[y][x]
		}
	}
	return result
}

func WordToPoint(word string) Point {
	parts := strings.Split(word, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{
		X: x,
		Y: y,
	}
}

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var field LightsGrid
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			if parts[0] == "toggle" {
				field.Toggle(WordToPoint(parts[1]), WordToPoint(parts[3]))
			} else if parts[1] == "on" {
				field.Set(true, WordToPoint(parts[2]), WordToPoint(parts[4]))
			} else if parts[1] == "off" {
				field.Set(false, WordToPoint(parts[2]), WordToPoint(parts[4]))
			}
		} else {
			break
		}
	}
	fmt.Println("Всего огней загорится: ", field.NumberOfLitLights())
}

func Part2() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var field RightLightsGrid
	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			if parts[0] == "toggle" {
				field.Add(2, WordToPoint(parts[1]), WordToPoint(parts[3]))
			} else if parts[1] == "on" {
				field.Add(1, WordToPoint(parts[2]), WordToPoint(parts[4]))
			} else if parts[1] == "off" {
				field.Add(-1, WordToPoint(parts[2]), WordToPoint(parts[4]))
			}
		} else {
			break
		}
	}
	fmt.Println("Всего огней загорится: ", field.TotalBrightness())
}
