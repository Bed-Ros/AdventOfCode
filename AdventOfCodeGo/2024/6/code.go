package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	inputDataPath = "2024/6/Вводные данные.txt"
	wall          = '#'
	guard         = '^'
	space         = '.'
	point         = 'o'
)

type room struct {
	Walls map[СommonElements.Point]interface{}
	Guard СommonElements.Point
	Size  СommonElements.Point
}

func (r room) print(points map[СommonElements.Point]interface{}) {
	for y := 0; y < r.Size.Y; y++ {
		for x := 0; x < r.Size.X; x++ {
			curPoint := СommonElements.Point{X: x, Y: y}
			_, isWall := r.Walls[curPoint]
			if isWall {
				fmt.Printf(string(wall))
				continue
			}
			if r.Guard == curPoint {
				fmt.Printf(string(guard))
				continue
			}
			_, isPoint := points[curPoint]
			if isPoint {
				fmt.Printf(string(point))
				continue
			}
			fmt.Printf(string(space))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (r room) simulateGuard() (map[СommonElements.Point]interface{}, bool) {
	vectors := []СommonElements.Point{
		{0, -1}, //Вверх
		{1, 0},  //Вправо
		{0, 1},  //Вниз
		{-1, 0}, //Влево
	}
	vectorIndex := 0
	steps := make(map[step]interface{})
	points := make(map[СommonElements.Point]interface{})
	for {
		//Создаем следующую предполагаемую позицию охранника
		nextGuardPoint := r.Guard.Add(vectors[vectorIndex])
		//Следующая клетка - стена
		_, nextGuardPointIsWall := r.Walls[nextGuardPoint]
		if nextGuardPointIsWall {
			vectorIndex++
			if vectorIndex >= len(vectors) {
				vectorIndex = 0
			}
			continue
		}
		//Следующая клетка за пределами комнаты
		if nextGuardPoint.X < 0 || nextGuardPoint.X >= r.Size.X ||
			nextGuardPoint.Y < 0 || nextGuardPoint.Y >= r.Size.Y {
			break
		}
		//Проверяем что данного шага не было, иначе охранник в петле
		curStep := step{
			Location: r.Guard,
			Vector:   vectors[vectorIndex],
		}
		_, curStepWasBefore := steps[curStep]
		if curStepWasBefore {
			return points, true
		}
		//Передвигаем охранника и сохраняем результат
		r.Guard = nextGuardPoint
		steps[curStep] = nil
		points[r.Guard] = nil
	}
	return points, false
}

type step struct {
	Location СommonElements.Point
	Vector   СommonElements.Point
}

func readInput() room {
	file, err := os.Open(inputDataPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wallsPoints := make(map[СommonElements.Point]interface{})
	var guardPoint СommonElements.Point

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
		switch r {
		case wall:
			wallsPoints[curPoint] = nil
		case guard:
			guardPoint = curPoint
		case '\n':
			curPoint.Y++
			curPoint.X = -1
		}
		curPoint.X++
	}
	curPoint.Y++
	return room{
		Walls: wallsPoints,
		Guard: guardPoint,
		Size:  curPoint,
	}
}

func Part1() {
	room := readInput()
	points, isLoop := room.simulateGuard()
	if isLoop {
		fmt.Printf("Охранник застрял в петле")
	}
	fmt.Printf("Результат: %d", len(points)+1)
}

func Part2() {
	room := readInput()
	points, isLoop := room.simulateGuard()
	//room.print(points)
	if isLoop {
		fmt.Printf("Охранник застрял в петле")
	}
	var result int
	for point := range points {
		room.Walls[point] = nil
		points, isLoop = room.simulateGuard()
		//room.print(points)
		if isLoop {
			result++
		}
		delete(room.Walls, point)
	}
	fmt.Printf("Результат: %d", result)
}
