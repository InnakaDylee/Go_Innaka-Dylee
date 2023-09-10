package main

import "fmt"

func fibonacci(number int) int {

	// your code here
	if number <= 1 {
		return number
	}
	
	fibo := make([]int, number+1)

	fibo[0] = 0
	fibo[1] = 1
	for i := 2; i <= number; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}

	return fibo[number]
}

func main() {

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}
