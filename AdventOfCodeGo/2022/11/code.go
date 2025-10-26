package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputDataPath = "2022/11/Вводные данные.txt"
	Quintillion   = 1_000_000_000_000_000_000
)

type bigNumber struct {
	NumberOfQuintillions *bigNumber //количество слагаемых числа, равные 10^18
	LeftOver             int64      //остаток < 10^18
}

func newBigNumber(number int) *bigNumber {
	return &bigNumber{
		NumberOfQuintillions: nil,
		LeftOver:             int64(number),
	}
}

func (bn *bigNumber) Add(anotherBn *bigNumber) *bigNumber {
	result := bn
	//Добавляем квинтиллионы
	if anotherBn.NumberOfQuintillions != nil {
		if result.NumberOfQuintillions == nil {
			result.NumberOfQuintillions = anotherBn.NumberOfQuintillions
		} else {
			result.NumberOfQuintillions = result.NumberOfQuintillions.Add(anotherBn.NumberOfQuintillions)
		}
	}
	//Складываем остатки
	newLeftOver := result.LeftOver + anotherBn.LeftOver
	q := newLeftOver / Quintillion
	//Если сумма остатков больше квинтиллиона, то добавляем результату соответствующее количество квинтиллионов
	if q > 0 {
		if result.NumberOfQuintillions == nil {
			result.NumberOfQuintillions = &bigNumber{NumberOfQuintillions: &bigNumber{LeftOver: q}}
		} else {
			result.NumberOfQuintillions = result.NumberOfQuintillions.Add(&bigNumber{NumberOfQuintillions: &bigNumber{LeftOver: q}})
		}
		result.LeftOver = newLeftOver % Quintillion
	} else {
		result.LeftOver = newLeftOver
	}
	return result
}

func (bn *bigNumber) Multiply(anotherBn *bigNumber) *bigNumber {
	result := newBigNumber(0)
	//Перемножаем квинтиллионы
	if bn.NumberOfQuintillions != nil && anotherBn.NumberOfQuintillions != nil {
		result.NumberOfQuintillions = bn.NumberOfQuintillions.Multiply(anotherBn.NumberOfQuintillions)
	}
	//Перемножаем остатки
	var smallLeftOver, bigLeftOver int64
	if bn.LeftOver > anotherBn.LeftOver {
		bigLeftOver = bn.LeftOver
		smallLeftOver = anotherBn.LeftOver
	} else {
		smallLeftOver = bn.LeftOver
		bigLeftOver = anotherBn.LeftOver
	}
	for i := 0; int64(i) < smallLeftOver; i++ {
		result = result.Add(&bigNumber{LeftOver: bigLeftOver})
	}
	//Перемножаем остаткок1 и квинтиллионы2
	if anotherBn.NumberOfQuintillions != nil {
		for i := 0; int64(i) < bn.LeftOver; i++ {
			result = result.Add(&bigNumber{NumberOfQuintillions: anotherBn.NumberOfQuintillions})
		}
	}
	//Перемножаем квинтиллионы1 и остаткок2
	if bn.NumberOfQuintillions != nil {
		for i := 0; int64(i) < anotherBn.LeftOver; i++ {
			result = result.Add(&bigNumber{NumberOfQuintillions: bn.NumberOfQuintillions})
		}
	}
	return result
}

//func (bn *bigNumber) Divide(number int) *bigNumber {
//
//}

func (bn *bigNumber) Divisible(number int) bool {
	var result = *bn
	q := Quintillion % int64(number)
	for {
		if result.NumberOfQuintillions == nil {
			result.LeftOver = result.LeftOver % int64(number)
			break
		} else {
			x := result.NumberOfQuintillions.Multiply(&bigNumber{LeftOver: q})
			result.NumberOfQuintillions = nil
			result = *result.Add(x)
		}
	}
	return result.LeftOver == 0
}

type monkey struct {
	Items      []*bigNumber
	Operation  func(*bigNumber) *bigNumber
	TestNumber int
	TestTrue   int
	TestFalse  int
}

func readMonkeysFromFile() []monkey {
	//Открываем файл
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewReader(inputFile)
	//Нужные переменные
	brkFor := false
	var monkeys []monkey
	var curMonkey monkey
	var i int
	//Для каждой строки:
	for {
		curStr, err := scanner.ReadString('\n')
		if err != nil {
			brkFor = true
		}
		parts := strings.Fields(curStr)
		if len(parts) > 0 {
			switch i {
			case 1: //предметы
				strNumbers := strings.Split(strings.Join(parts[2:], ""), ",")
				for _, number := range strNumbers {
					n, _ := strconv.Atoi(number)
					curMonkey.Items = append(curMonkey.Items, newBigNumber(n))
				}
				break
			case 2: //операция
				var n int
				if parts[5] != "old" {
					n, _ = strconv.Atoi(parts[5])
				}
				switch parts[4] {
				case "*":
					if parts[5] == "old" {
						curMonkey.Operation = func(old *bigNumber) *bigNumber {
							return old.Multiply(old)
						}
					} else {
						curMonkey.Operation = func(old *bigNumber) *bigNumber {
							return old.Multiply(newBigNumber(n))
						}
					}
					break
				case "+":
					if parts[5] == "old" {
						curMonkey.Operation = func(old *bigNumber) *bigNumber {
							return old.Add(old)
						}
					} else {
						curMonkey.Operation = func(old *bigNumber) *bigNumber {
							return old.Add(newBigNumber(n))
						}
					}
					break
				}
				break
			case 3: //тест
				n, _ := strconv.Atoi(parts[len(parts)-1])
				curMonkey.TestNumber = n
				break
			case 4: //тест да
				n, _ := strconv.Atoi(parts[len(parts)-1])
				curMonkey.TestTrue = n
				break
			case 5: //тест нет
				n, _ := strconv.Atoi(parts[len(parts)-1])
				curMonkey.TestFalse = n
				break
			}
			i++
		} else {
			monkeys = append(monkeys, curMonkey)
			curMonkey = monkey{}
			i = 0
		}
		if brkFor {
			monkeys = append(monkeys, curMonkey)
			break
		}
	}
	return monkeys
}

func main(roundsNum int, divideByThree bool) {
	monkeys := readMonkeysFromFile()
	result := make([]int, len(monkeys))
	//Перекидываем предметы между обезьянами
	for i := 0; i < roundsNum; i++ {
		for curMonkeyNum, curMonkey := range monkeys {
			for _, curItem := range curMonkey.Items {
				result[curMonkeyNum]++
				fmt.Println(result)
				curItem = curMonkey.Operation(curItem)
				//if divideByThree {
				//	curItem = curItem.Divide(3)
				//}
				if curItem.Divisible(curMonkey.TestNumber) {
					monkeys[curMonkey.TestTrue].Items = append(monkeys[curMonkey.TestTrue].Items, curItem)
				} else {
					monkeys[curMonkey.TestFalse].Items = append(monkeys[curMonkey.TestFalse].Items, curItem)
				}
			}
			monkeys[curMonkeyNum].Items = nil
		}
	}
	//Результат
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	fmt.Printf(
		"После %d раундов обезьяньих махинаций с перебрасыванием вещей уровень обезьяньего бизнеса будет равен %d\n",
		roundsNum,
		result[0]*result[1])
}

func Part1() {
	main(20, true)
}

func Part2() {
	main(10000, false)
}
