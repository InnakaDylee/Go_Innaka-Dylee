package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	// your code here
	if number == 0{
		return 0
	}

	isPrime := func(prime int)bool{
		if prime <= 3{
			return true
		}
		if prime == 5{
			return true
		}
		if prime%2 == 0 || prime%3 == 0 || prime%5 == 0{
			return false
		}
		squareRoot := int(math.Sqrt(float64(prime)))
		for i := 2; i < squareRoot; i++{
			if prime%i == 0{
				return false
			}
		} 
		return true
	}

	index := 0
	prime := 2
	for{
		if isPrime(prime){
			index++
		}
		if index == number{
			return prime
		}
		prime++
	}

}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}