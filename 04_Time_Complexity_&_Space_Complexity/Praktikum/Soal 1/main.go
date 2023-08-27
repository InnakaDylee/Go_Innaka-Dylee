package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(13)) // true

	fmt.Println(primeNumber(17)) // true

	fmt.Println(primeNumber(20)) // false

	fmt.Println(primeNumber(35)) // false

}

func primeNumber(i int) bool {
	input := strconv.Itoa(i)
	var n string = string(input[len(input)-1])
	fN := string(input[0])
	firstNumber, _ := strconv.Atoi(fN)
	floatNumber, _ := strconv.ParseFloat(input, 64)

	if input == "2" || input == "3" || input == "5" {
		return true
	}
	if n == "0" || n == "2" || n == "4" || n == "5" || n == "6" || n == "8" {
		return false
	}

	for i := 1; i <= len(input)-1; i++ {
		m := string(input[i])
		bilangan, _ := strconv.Atoi(m)
		firstNumber += bilangan
	}
	lastNumber := firstNumber

	if lastNumber%3 == 0 {
		return false
	}
	squareNumber := int(math.Sqrt(floatNumber))
	for j := 2; j < squareNumber; j += 1 {
		if primeNumber(j) == true {
			if i%j == 0 {
				return false
			}
		}
	}
	return true
}
