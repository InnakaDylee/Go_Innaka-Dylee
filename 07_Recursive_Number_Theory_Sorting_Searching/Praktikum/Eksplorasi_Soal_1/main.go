package main


import "fmt"


func MaxSequence(arr []int) int {

// your code here
	var sum,max int

	for i:=0;i < len(arr)-1;i++{
		for j:=i; j < len(arr)-1;j++{
			sum += arr[j]
			if sum > max{
				max = sum
			}
		}
		sum = 0
	}
	return max
}


func main() {

fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}