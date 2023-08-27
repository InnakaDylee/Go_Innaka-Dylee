package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	
	fmt.Print("\nResult: ")
	fmt.Println(diffDiagonal(matrix))
}

func diffDiagonal(matrix [][]int) int {
	result := 0
	for i := 0; i < len(matrix); i++ {
		result += matrix[i][i] - matrix[i][len(matrix)-1-i]
	}
	if result < 0 {
		result *= -1
	}
	return result
}