package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputDataPath = "2022/12/Вводные данные.txt"
)

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var heightMap [][]int
	var result int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {

		}
		if brkFor {
			break
		}
	}
	fmt.Printf("%d шагов необходимо сделать, чтобы переместиться из вашего "+
		"текущего положения в место, где должен быть лучший сигнал", result)
}
