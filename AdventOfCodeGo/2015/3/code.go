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

type santa struct {
	Points   map[СommonElements.Point]bool
	curPoint СommonElements.Point
}

func newSanta() santa {
	result := santa{
		Points:   make(map[СommonElements.Point]bool),
		curPoint: СommonElements.Point{}}
	result.Points[result.curPoint] = true
	return result
}

func (s *santa) Step(symbol byte) {
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

func (s *santa) GetUniqueHouses() int {
	return len(s.Points)
}

func Part1() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	santa := newSanta()
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
	santas := []santa{
		newSanta(),
		newSanta(),
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
