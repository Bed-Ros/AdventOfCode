package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const (
	inputDataPath = "2022/7/Вводные данные.txt"
)

type dirTree struct {
	Parent    *dirTree
	Size      int
	ChildDirs []*dirTree
	Name      string
}

func (t dirTree) FullSizeOfSmallestDir(biggerThen int) int {
	curMin := syscall.INFINITE
	for _, child := range t.ChildDirs {
		if len(child.ChildDirs) > 0 {
			a := child.FullSize()
			if a < curMin && a >= biggerThen {
				curMin = a
			}
			a = child.FullSizeOfSmallestDir(biggerThen)
			if a < curMin && a >= biggerThen {
				curMin = a
			}
		}
	}
	return curMin
}

func (t dirTree) SizeOfAllSmallDirs(lessThen int) int {
	var result int
	for _, child := range t.ChildDirs {
		if len(child.ChildDirs) > 0 {
			result += child.SizeOfAllSmallDirs(lessThen)
			a := child.FullSize()
			if a <= lessThen {
				result += a
			}
		}
	}
	return result
}

func (t dirTree) FullSize() int {
	result := t.Size
	for _, child := range t.ChildDirs {
		result += child.FullSize()
	}
	return result
}

func createDirTree() dirTree {
	//Открывам файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	baseDir := dirTree{}
	curDir := &baseDir
	readingFilesList := false
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			//Создание дерева папок и файлов
			if parts[0] == "$" {
				readingFilesList = false
				switch parts[1] {
				case "cd":
					switch parts[2] {
					case "/":
						curDir = &baseDir
						break
					case "..":
						curDir = curDir.Parent
						break
					default:
						for _, child := range curDir.ChildDirs {
							if child.Name == parts[2] {
								curDir = child
							}
						}
					}
					break
				case "ls":
					readingFilesList = true
					break
				}
			} else if readingFilesList {
				number, _ := strconv.Atoi(parts[0])
				curDir.ChildDirs = append(curDir.ChildDirs, &dirTree{
					Parent:    curDir,
					Size:      number,
					ChildDirs: nil,
					Name:      parts[1],
				})
			}
		}
		if brkFor {
			break
		}
	}
	return baseDir
}

func Part1() {
	baseDir := createDirTree()
	fmt.Printf("Общая сумма размеров каталогов размером до 100000 равна %d", baseDir.SizeOfAllSmallDirs(100000))
}

func Part2() {
	baseDir := createDirTree()
	remainingEmptySpace := 70000000 - baseDir.FullSize()
	neededExtraSpaceForUpdate := 30000000 - remainingEmptySpace
	fmt.Printf("Размер самого маленького каталога, удаление которого освободит достаточно места в файловой системе "+
		"для запуска обновления равен %d", baseDir.FullSizeOfSmallestDir(neededExtraSpaceForUpdate))
}
