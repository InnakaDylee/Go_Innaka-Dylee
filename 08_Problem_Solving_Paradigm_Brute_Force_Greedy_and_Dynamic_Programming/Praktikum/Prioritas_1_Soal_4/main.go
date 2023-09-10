package main

import "fmt"

func SimpleEquations(a, b, c int) {
	var result [][]int
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			for k := 1; k <= c; k++ {
				if i+j+k == a && i*j*k == b && i*i+j*j+k*k == c {
					result = append(result, []int{i, j, k})
				}
			}
		}
	}
	
	if result == nil {
		fmt.Println("No Solution")
	} else {
		fmt.Println(result[0])
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
