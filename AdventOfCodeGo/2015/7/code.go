package day7

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
	inputDataPath                = "2015/7/Вводные данные.txt"
	wordBetweenLeftAndRightParts = "->"
)

var (
	//Инструкции
	instructions = make(map[string]LeftPart)
	//Результаты инструкций
	results = make(map[string]uint16)
)

type LeftPart struct {
	Instruction    func(wires []uint16) uint16
	WiresOrNumbers []string
}

func NewLeftPart(instruction string) LeftPart {
	result := LeftPart{}

	parts := strings.Fields(instruction)
	if len(parts) == 1 {
		result.WiresOrNumbers = []string{parts[0]}
		result.Instruction = func(wires []uint16) uint16 {
			return wires[0]
		}
	} else if len(parts) == 2 {
		//NOT провод ->
		result.WiresOrNumbers = []string{parts[1]}
		result.Instruction = func(wires []uint16) uint16 {
			return math.MaxUint16 ^ wires[0]
		}
	} else if len(parts) == 3 {
		result.WiresOrNumbers = []string{parts[0], parts[2]}
		switch parts[1] {
		//провод AND провод ->
		case "AND":
			result.Instruction = func(wires []uint16) uint16 {
				return wires[0] & wires[1]
			}
			break
		//провод OR провод ->
		case "OR":
			result.Instruction = func(wires []uint16) uint16 {
				return wires[0] | wires[1]
			}
			break
		//провод LSHIFT провод ->
		case "LSHIFT":
			result.Instruction = func(wires []uint16) uint16 {
				return wires[0] << wires[1]
			}
			break
		//провод RSHIFT провод ->
		case "RSHIFT":
			result.Instruction = func(wires []uint16) uint16 {
				return wires[0] >> wires[1]
			}
			break
		}
	}

	return result
}

func Part1() {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}

	//Для каждой строки:
	scanner := bufio.NewReader(inputFile)
	for {
		curStr, _ := scanner.ReadString('\n')
		if len(curStr) > 0 {
			//Записываем инструкцию
			parts := strings.Split(curStr, wordBetweenLeftAndRightParts)
			instructions[strings.TrimSpace(parts[1])] = NewLeftPart(parts[0])
		} else {
			break
		}
	}

	//Решаем задачу
	stack := []string{"a"}
	for {
		if len(stack) == 0 {
			break
		}
		//Берем последний элемент стака
		rightPart := stack[len(stack)-1]
		leftPart := instructions[rightPart]
		//Берем все провода, которые нужны для результата данной инструкции
		//и проверяем на наличие их результата,
		//если у провода нет результата, то кладем его в стак
		allWiresHaveResult := true
		wiresResults := make([]uint16, len(leftPart.WiresOrNumbers))
		for i, wire := range leftPart.WiresOrNumbers {
			//пробуем превратить в число
			number, err := strconv.Atoi(wire)
			if err == nil {
				//это число
				wiresResults[i] = uint16(number)
			} else {
				//это не число
				//проверяем есть ли уже результат у этого провода
				num, ok := results[wire]
				if ok {
					//есть
					wiresResults[i] = num
				} else {
					//нет
					stack = append(stack, wire)
					allWiresHaveResult = false
				}
			}
		}
		//Если значения всех проводов известно,
		//то находим результат данной инструкции и убираем данную инструкцию из стака
		if allWiresHaveResult {
			results[rightPart] = leftPart.Instruction(wiresResults)
			stack = stack[:len(stack)-1]
		}
	}

	fmt.Println("Провод 'a' подаст сигнал: ", results["a"])
}

func Part2() {
	Part1()
	wireA := results["a"]
	results = make(map[string]uint16)
	results["b"] = wireA
	Part1()
}
