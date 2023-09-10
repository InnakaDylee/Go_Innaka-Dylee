package main

import (
	"fmt"
	"time"
)

func main() {
	var numberMultiple, multipleTime int
	fmt.Print("Input Number        : ")
	fmt.Scanln(&numberMultiple)
	fmt.Print("Input Multiple Time : ")
	fmt.Scanln(&multipleTime)

	massage := make(chan int)// channel

	multiple := func(number int, multiple int, massage chan int) {
		var multipleNumber int
		for index := 1; index <= multiple; index++ {
			multipleNumber += number
			massage <- multipleNumber
		}
	}

	go multiple(numberMultiple, multipleTime, massage)

	go func() {
		for i := 1; i <= multipleTime; i++ {
			multiple := <-massage
			fmt.Printf("Time %d : %d\n", i, multiple)
			time.Sleep(3 * time.Second)
		}
	}()

	time.Sleep(time.Duration(multipleTime) * 3 * time.Second)

}
