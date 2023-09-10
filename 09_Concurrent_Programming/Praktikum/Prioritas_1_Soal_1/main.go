package main

import (
	"fmt"
	"time"
)

func multiple(multipleNumber, multipleTime int) {
	var multiple int
	for i := 1; i <= multipleTime; i++ {
		multiple += multipleNumber
		fmt.Printf("time %d : %d \n",i , multiple)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var multipleNumber, multipleTime int
	fmt.Print("Input Number        : ")
	fmt.Scanln(&multipleNumber)
	fmt.Print("Input Multiple Time : ")
	fmt.Scanln(&multipleTime)
	go multiple(multipleNumber, multipleTime)
	time.Sleep(time.Duration(multipleTime) * 3 * time.Second)
}
