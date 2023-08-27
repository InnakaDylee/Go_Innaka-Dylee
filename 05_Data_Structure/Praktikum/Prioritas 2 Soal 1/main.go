package main


import "fmt"


func PairSum(arr []int, target int) []int {
	tempArr := make(map[int]int)
	var value, index, index2, num int

	var isFound bool
	for index, num = range arr {
		value = target - num
		index2, isFound = tempArr[value]
		if isFound {
			return []int{index2, index}
		}
		tempArr[num] = index
	}
	return nil
}


func main() {

// Test cases

fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}