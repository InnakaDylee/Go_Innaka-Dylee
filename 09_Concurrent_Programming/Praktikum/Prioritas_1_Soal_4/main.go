package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	mu.Lock()
	fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	numbers := []int{5, 7, 10, 3}

	for _, n := range numbers {
		wg.Add(1)
		go factorial(n, &wg, &mu)
	}

	wg.Wait()
}
