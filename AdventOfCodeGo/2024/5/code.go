package day5

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2024/5/Вводные данные.txt"
)

func stringsToInts(list []string) []int {
	result := make([]int, len(list))
	for i, s := range list {
		n, _ := strconv.Atoi(s)
		result[i] = n
	}
	return result
}

func readInput() (RulesCollection, [][]int) {
	text, _ := os.ReadFile(inputDataPath)
	parts := strings.Split(string(text), "\r\n\r\n")
	rules := make(map[int]*Rule)
	for _, line := range strings.Split(parts[0], "\r\n") {
		lineNums := stringsToInts(strings.Split(line, "|"))
		curRule, ok := rules[lineNums[0]]
		if ok {
			curRule.After = append(curRule.After, lineNums[1])
		} else {
			rules[lineNums[0]] = &Rule{After: []int{lineNums[1]}}
		}
		curRule, ok = rules[lineNums[1]]
		if ok {
			curRule.Before = append(curRule.Before, lineNums[0])
		} else {
			rules[lineNums[1]] = &Rule{Before: []int{lineNums[0]}}
		}
	}
	var updates [][]int
	for _, line := range strings.Split(parts[1], "\r\n") {
		lineNums := stringsToInts(strings.Split(line, ","))
		updates = append(updates, lineNums)
	}
	return rules, updates
}

type RulesCollection map[int]*Rule

func (c RulesCollection) Check(update []int) bool {
	for i, n := range update {
		curRule, ok := c[n]
		if !ok {
			continue
		}
		if curRule.Check(update, i) != 0 {
			return false
		}
	}
	return true
}

type Rule struct {
	Before []int
	After  []int
}

func (r *Rule) Check(update []int, index int) int {
	for _, n := range update[index+1:] {
		if slices.Contains(r.Before, n) {
			return 1
		}
	}
	for _, n := range update[:index] {
		if slices.Contains(r.After, n) {
			return -1
		}
	}
	return 0
}

func (c RulesCollection) CreateValidUpdate(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	result := []int{input[0]}
	for _, n := range input[1:] {
		result = append(result, n)
		curRule, ok := c[n]
		if !ok {
			continue
		}
		index := len(result) - 1
		for {
			step := curRule.Check(result, index)
			if step == 0 {
				break
			}
			newIndex := index + step
			result[index], result[newIndex] = result[newIndex], result[index]
			index = newIndex
		}
	}
	return result
}

func Part1() {
	rules, updates := readInput()
	var result int
	for _, update := range updates {
		if rules.Check(update) {
			result += update[len(update)/2]
		}
	}
	log.Printf("Результат: %d", result)
}

func Part2() {
	rules, updates := readInput()
	var result int
	for _, update := range updates {
		if !rules.Check(update) {
			validUpdate := rules.CreateValidUpdate(update)
			result += validUpdate[len(validUpdate)/2]
		}
	}
	log.Printf("Результат: %d", result)
}
