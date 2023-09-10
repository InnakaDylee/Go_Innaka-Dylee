package main

import (
	"fmt"
)

func pascal(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	n := 5
	fmt.Println(pascal(n))
	// pascalTriangle := pascal(n)

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n-i-1; j++ {
	// 		fmt.Print("  ")
	// 	}
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Printf("%4d", pascalTriangle[i][j])
	// 	}
	// 	fmt.Println()
	// }
}
