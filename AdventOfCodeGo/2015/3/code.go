package day3

import (
	"fmt"
	"log"
	"maps"
	"os"
)

const (
	inputDataPath = "2015/3/Вводные данные.txt"
)

type Point struct {
	X, Y int
}

type Santa struct {
	Points   map[Point]bool
	curPoint Point
}

func NewSanta() Santa {
	result := Santa{
		Points:   make(map[Point]bool),
		curPoint: Point{0, 0}}
	result.Points[result.curPoint] = true
	return result
}

func (s *Santa) Step(symbol byte) {
	switch symbol {
	case '^':
		s.curPoint.Y++
		break
	case 'v':
		s.curPoint.Y--
		break
	case '>':
		s.curPoint.X++
		break
	case '<':
		s.curPoint.X--
		break
	}
	s.Points[s.curPoint] = true
}

func (s *Santa) GetUniqueHouses() int {
	return len(s.Points)
}

func Part1() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	santa := NewSanta()
	//для каждого символа
	for _, b := range bytes {
		santa.Step(b)
	}
	fmt.Println("Количество домов: ", santa.GetUniqueHouses())
}

func Part2() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	santas := []Santa{
		NewSanta(),
		NewSanta(),
	}
	//для каждого символа
	for i, b := range bytes {
		if i%2 == 0 {
			santas[0].Step(b)
		} else {
			santas[1].Step(b)
		}
	}
	maps.Copy(santas[0].Points, santas[1].Points)
	fmt.Println("Количество домов: ", santas[0].GetUniqueHouses())
}
