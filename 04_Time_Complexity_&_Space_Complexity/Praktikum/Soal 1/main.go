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

func primeNumber(i int)bool{
	input := strconv.Itoa(i)
	var n string = string(input[len(input)-1])
	fN := string(input[0])
	firstNumber, _ := strconv.Atoi(fN)
	floatNumber,_ := strconv.ParseFloat(input, 64)
	
		if i < 2{
			return false
		}
		if i == 2 || i == 3 || i == 5{
			return true
		}
		if n == "0" || n == "2" || n == "4" || n == "5" || n == "6" || n == "8"{
			return false
		}
	
		for i := 1; i <= len(input)-1;i++{
			m := string(input[i])
			bilangan, _ := strconv.Atoi(m)
			firstNumber += bilangan
		}
		totalNumber := firstNumber
	
		if totalNumber%3 == 0 {
			return false
		}
		squareNumber := int(math.Sqrt(floatNumber))
		for j := 2; j < squareNumber;j+=1{
			if primeNumber(j) && i%j == 0 {
					return false	
			}
		} 
		return true
}