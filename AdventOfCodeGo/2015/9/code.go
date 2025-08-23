package day9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2015/9/Вводные данные.txt"
)

var (
	allCities     []city
	citiesIndexes map[city]int
	roadsLengths  map[road]int
)

type city string
type road string

func newRoad(cities ...city) (direct, reverse road) {
	for _, c := range cities {
		direct = road(string(direct) + string(rune(citiesIndexes[c])))
	}
	for i := len(cities) - 1; i >= 0; i-- {
		reverse = road(string(reverse) + string(rune(citiesIndexes[cities[i]])))
	}
	return
}

func parseInputFile() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}

	var i int

	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			parts := strings.Fields(curStr)
			l, err := strconv.Atoi(parts[4])
			if err != nil {
				log.Fatalln(err)
			}
			city1 := city(parts[0])
			city2 := city(parts[2])
			citiesIndexes[city1] = i
			i++
			citiesIndexes[city2] = i
			i++
			road1, road2 := newRoad(city1, city2)
			roadsLengths[road1] = l
			roadsLengths[road2] = l
		} else {
			break
		}
	}
	for s := range citiesIndexes {
		allCities = append(allCities, s)
	}
}

type roadBranch struct {
	CurCity  city
	Children []*roadBranch
}

func createRoadTree(parent city, child []city) *roadBranch {
	result := roadBranch{CurCity: parent}
	if len(child) > 0 {
		for i, c := range child {
			result.Children = append(result.Children, createRoadTree(c, append(child[:i], child[i+1:]...)))
		}
	}
	return &result
}

func (b *roadBranch) minLength() int {
	result := math.MaxInt
	for _, branch := range b.Children {
		l := branch.minLength()
		if l < result {
			result = l
		}
	}
	return result
}

func bruteForceWithMemory() int {
	fullTree := createRoadTree("", allCities)
	result := math.MaxInt
	for _, branch := range fullTree.Children {
		l := branch.minLength()
		if l < result {
			result = l
		}
	}
	return result
}

func Part1() {
	parseInputFile()
	fmt.Println("Минимальное расстояние: ", bruteForceWithMemory())
}

func Part2() {

}
