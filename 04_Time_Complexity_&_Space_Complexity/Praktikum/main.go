package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	var input int
	fmt.Scanln(&input)

	startTime := time.Now()
	fmt.Println(startTime)

	fmt.Println(prima(input))

	endTime := time.Now()
	fmt.Println(endTime)
	
}

func prima(i int)bool{
	result := 0

	input := strconv.Itoa(i)
	var n string = string(input[len(input)-1])
	fN := string(input[0])
	firstNumber, _ := strconv.Atoi(fN)
	floatNumber,_ := strconv.ParseFloat(input, 64)
	

		if input == "2" || input == "3" || input == "5"{
			result += 1
			return true
		}else{
			result += 1
		}
		if n == "0" || n == "2" || n == "4" || n == "5" || n == "6" || n == "8"{
			result += 1
			return false
		}else{
			result += 1
		}
	
		for i := 1; i <= len(input)-1;i++{
			m := string(input[i])
			result += 1
			bilangan, _ := strconv.Atoi(m)
			firstNumber += bilangan
		}
		lastNumber := firstNumber
	
		if lastNumber%3 == 0 {
			result += 1
			return false
		}
		squareNumber := int(math.Sqrt(floatNumber))
		for j := 2; j < squareNumber;j+=1{
			if prima(j) == true {
				result += 1
				if i%j == 0 {
					result += 1
					return false
				}
			}
		} 
		return true
}
