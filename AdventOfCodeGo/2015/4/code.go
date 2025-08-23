package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"unicode"
)

const (
	inputData = "yzbqklnj"
)

func foo(zeroes int) {
	i := 1
	resultFound := false
	for {
		hash := md5.Sum([]byte(inputData + strconv.Itoa(i)))
		hexHash := hex.EncodeToString(hash[:])
		for n, b := range hexHash {
			if n < zeroes {
				if b == '0' {
					continue
				} else {
					break
				}
			} else {
				if unicode.IsDigit(b) {
					resultFound = true
				}
				break
			}
		}
		if resultFound {
			break
		}
		i++
	}
	fmt.Println("Минимальное число: ", i)
}

func Part1() {
	foo(5)
}

func Part2() {
	foo(6)
}
