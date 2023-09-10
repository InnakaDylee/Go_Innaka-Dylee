package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	vals := []int{1000, 500, 100, 50, 10, 5, 1}
	symbols := []string{"M", "D", "C", "L", "X", "V", "I"}

	romanNum := ""
	i := 0

	for num > 0 {
		if num >= vals[i] {
			romanNum += symbols[i]
			num -= vals[i]
		} else {
			i++
		}
	}

	romanNum = strings.ReplaceAll(romanNum, "DCCCC", "CM")
	romanNum = strings.ReplaceAll(romanNum, "CCCC", "CD")
	romanNum = strings.ReplaceAll(romanNum, "LXXXX", "XC")
	romanNum = strings.ReplaceAll(romanNum, "XXXX", "XL")
	romanNum = strings.ReplaceAll(romanNum, "VIIII", "IX")
	romanNum = strings.ReplaceAll(romanNum, "IIII", "IV")

	return romanNum
}

func main() {
	fmt.Println(intToRoman(4)) //IV
	fmt.Println(intToRoman(9)) //IX
	fmt.Println(intToRoman(23)) //XXIII
	fmt.Println(intToRoman(2021)) //MMXXI
	fmt.Println(intToRoman(1646)) //MDCXLVI
}
