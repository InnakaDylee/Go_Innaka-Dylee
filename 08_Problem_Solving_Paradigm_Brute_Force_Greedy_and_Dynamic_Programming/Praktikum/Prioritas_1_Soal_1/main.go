package main

import (
	"fmt"
)

func generateBinaryRepresentation(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		// fmt.Print(" | ",i," : ")
		binaryInt := 0
		base := 1
		for j := i; j > 0; j /= 2 {
			// fmt.Print(" < ",(j%2), " ")
			binaryInt += (j % 2) * base
			base *= 10
		}
		ans[i] = binaryInt
	}
	return ans
}

func main() {
	fmt.Println(generateBinaryRepresentation(5)) // Output: [0 1 10]
	fmt.Println(generateBinaryRepresentation(10)) // Output: [0 1 10 11]
}