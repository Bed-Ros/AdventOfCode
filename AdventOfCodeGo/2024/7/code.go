package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2024/7/Вводные данные.txt"
)

func readInput() []Equation {
	file, err := os.Open(inputDataPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result []Equation

	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, ":")
		r, _ := strconv.Atoi(parts[0])
		newEquation := Equation{Result: r}
		for _, s := range strings.Fields(parts[1]) {
			n, _ := strconv.Atoi(s)
			newEquation.Numbers = append(newEquation.Numbers, n)
		}
		result = append(result, newEquation)
	}
	return result
}

type Equation struct {
	Result  int
	Numbers []int
}

func (e Equation) FindSolution(operations []func(x1, x2 int) int) bool {
	if len(e.Numbers) == 0 {
		return false
	}
	if len(e.Numbers) == 1 {
		return e.Numbers[0] == e.Result
	}

	n1 := e.Numbers[0]
	n2 := e.Numbers[1]
	for _, operation := range operations {
		result := operation(n1, n2)
		if result > e.Result {
			continue
		}
		subEquation := Equation{Result: e.Result, Numbers: append([]int{result}, e.Numbers[2:]...)}
		if subEquation.FindSolution(operations) {
			return true
		}
	}
	return false
}

func MainFunc(operations []func(x1, x2 int) int) {
	equations := readInput()
	var result int
	for _, equation := range equations {
		if equation.FindSolution(operations) {
			result += equation.Result
		}
	}
	fmt.Printf("Результат: %d", result)
}

func Part1() {
	operations := []func(x1, x2 int) int{
		func(x1, x2 int) int { return x1 * x2 },
		func(x1, x2 int) int { return x1 + x2 },
	}
	MainFunc(operations)
}

func Part2() {
	operations := []func(x1, x2 int) int{
		func(x1, x2 int) int { return x1 * x2 },
		func(x1, x2 int) int { return x1 + x2 },
		func(x1, x2 int) int {
			r, _ := strconv.Atoi(strconv.Itoa(x1) + strconv.Itoa(x2))
			return r
		},
	}
	MainFunc(operations)
}
