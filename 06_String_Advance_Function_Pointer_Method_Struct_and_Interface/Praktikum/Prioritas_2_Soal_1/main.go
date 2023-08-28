package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	// your code here
	input =  strings.ToLower(input)
	var result string
	for _, value := range input{
		if value >= 'a' && value <= 'z'{
			result += string('a' + (value-'a'+int32(offset))%26)
		} else {
			result += string(value)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // def
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}