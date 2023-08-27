package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	tempArr := make(map[int]int)
	var result []int 

	for _, value := range angka {
		tempValue, _ := strconv.Atoi(string(value))
		tempArr[tempValue]++
	}
	for key, count := range tempArr{
		if count == 1{
			result = append(result, key)
		}
	}
	return result

}

func main() {

	// Test cases

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}
