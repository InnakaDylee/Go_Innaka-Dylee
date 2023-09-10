package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	jumpCost := make([]int, n)

	jumpCost[0] = 0
	jumpCost[1] = int(math.Abs(float64(jumps[0] - jumps[1])))

	for i := 2; i < n; i++ {
		jumpCost[i] = int(math.Min(
			float64(jumpCost[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))),
			float64(jumpCost[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2])))),
		))
	}

	return jumpCost[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))   //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
