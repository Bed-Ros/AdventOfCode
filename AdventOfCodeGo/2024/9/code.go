package day9

import (
	"fmt"
	"os"
	"strconv"
)

const (
	inputDataPath = "2024/9/Вводные данные.txt"
)

func readInput() []int {
	bytes, _ := os.ReadFile(inputDataPath)
	var result []int
	for _, symbol := range string(bytes) {
		x, _ := strconv.Atoi(string(symbol))
		result = append(result, x)
	}
	return result
}

func createDiskFromItsMap(diskMap []int) []int {
	isSpace := false
	var fileIndex int
	var disk []int
	for _, n := range diskMap {
		if isSpace {
			for i := 0; i < n; i++ {
				disk = append(disk, -1)
			}
		} else {
			for i := 0; i < n; i++ {
				disk = append(disk, fileIndex)
			}
			fileIndex++
		}
		isSpace = !isSpace
	}
	return disk
}

func calculateDisk(disk []int) int {
	var result int
	for i, n := range disk {
		if n != -1 {
			result += i * n
		}
	}
	return result
}

func Part1() {
	diskMap := readInput()
	//Создаем диск
	disk := createDiskFromItsMap(diskMap)
	//Компактно укладываем
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			for {
				if disk[len(disk)-1] == -1 {
					disk = disk[:len(disk)-1]
				} else {
					break
				}
			}
			disk[i] = disk[len(disk)-1]
			disk = disk[:len(disk)-1]
		}
	}
	//Считаем
	result := calculateDisk(disk)
	fmt.Printf("Результат: %d", result)
}

func Part2() {
	//diskMap := readInput()
	////Создаем диск
	//disk := createDiskFromItsMap(diskMap)
	////Перемещаем файлы целиком
	//lastIndex := len(disk) - 1
	//var curFile []int
	//for {
	//	cur := disk[lastIndex]
	//	if cur == -1 {
	//		continue
	//	}
	//
	//}
	////Считаем
	//result := calculateDisk(disk)
	//fmt.Printf("Результат: %d", result)
}
