package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// your code here
	var result []pair
	temp := map[string]int{}
	for _, value := range items {
		temp[value]++
	}
	for key, countMap := range temp {
		result = append(result, pair{name: key, count: countMap})
	}

	//bubble sorting
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].count > result[j].count {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
