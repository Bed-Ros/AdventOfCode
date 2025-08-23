package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2016/1/Вводные данные.txt"

	north = iota
	east
	south
	west
)

type point struct {
	X int
	Y int
}

func Part1() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	//Берем шаги
	steps := strings.Split(string(bytes), ",")
	var x, y int
	direction := north
	//Для каждого шага:
	for _, step := range steps {
		step = strings.TrimSpace(step)
		if len(step) == 0 {
			continue
		}
		//поворот
		rotate := step[0]
		if rotate == 'R' {
			direction++
		} else if rotate == 'L' {
			direction--
		}
		if direction > west {
			direction = north
		} else if direction < north {
			direction = west
		}
		//перемещение
		n, err := strconv.Atoi(step[1:])
		if err != nil {
			log.Fatalln(err)
		}
		switch direction {
		case north:
			y += n
			break
		case east:
			x += n
			break
		case south:
			y -= n
			break
		case west:
			x -= n
			break
		}
	}

	result := int(math.Abs(float64(x)) + math.Abs(float64(y)))
	fmt.Println("Кварталов от начальной позиции: ", result)
}

func Part2() {
	//Открываем файл
	bytes, err := os.ReadFile(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	//Берем шаги
	steps := strings.Split(string(bytes), ",")
	var curCoord point
	coords := make(map[point]bool)
	coords[curCoord] = true
	direction := north
	//Для каждого шага:
	for _, step := range steps {
		step = strings.TrimSpace(step)
		if len(step) == 0 {
			continue
		}
		//поворот
		rotate := step[0]
		if rotate == 'R' {
			direction++
		} else if rotate == 'L' {
			direction--
		}
		if direction > west {
			direction = north
		} else if direction < north {
			direction = west
		}
		//перемещение
		n, err := strconv.Atoi(step[1:])
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; i < n; i++ {
			switch direction {
			case north:
				curCoord.Y++
				break
			case east:
				curCoord.X++
				break
			case south:
				curCoord.Y--
				break
			case west:
				curCoord.X--
				break
			}
			//повторность
			_, ok := coords[curCoord]
			if ok {
				result := int(math.Abs(float64(curCoord.X)) + math.Abs(float64(curCoord.Y)))
				fmt.Println("Первое место, которое посещено дважды в ", result, " кварталах")
				return
			}
			coords[curCoord] = true
		}
	}

	fmt.Println("Повторов не было(")
}
