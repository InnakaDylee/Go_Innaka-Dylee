package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// your code here
	temp := ""
	for i := 0; i < len(b); i++ {
		temp = b[:len(b)-i]

		if strings.Contains(a, temp) {
			return temp
		}
	}
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}
