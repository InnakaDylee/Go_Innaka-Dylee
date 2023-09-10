package main

import (
	"fmt"
	"sync"
)

func countLetterFrequency(text string, wg *sync.WaitGroup, resultChan chan<- map[rune]int) {
	defer wg.Done()

	letterFrequency := make(map[rune]int)

	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			letterFrequency[char]++
		}
	}

	resultChan <- letterFrequency
}

func mergeResults(results []map[rune]int) map[rune]int {
	combinedResult := make(map[rune]int)

	for _, result := range results {
		for char, count := range result {
			combinedResult[char] += count
		}
	}

	return combinedResult
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

	numCPU := 4

	textChunks := make([]string, numCPU)
	chunkSize := len(text) / numCPU
	start := 0

	for i := 0; i < numCPU; i++ {
		end := start + chunkSize

		if i == numCPU-1 {
			end = len(text)
		}
		textChunks[i] = text[start:end]
		start = end
	}

	resultChan := make(chan map[rune]int, numCPU)
	var wg sync.WaitGroup

	for _, chunk := range textChunks {
		wg.Add(1)
		go countLetterFrequency(chunk, &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	var results []map[rune]int
	for result := range resultChan {
		results = append(results, result)
	}

	combinedResult := mergeResults(results)

	for char, count := range combinedResult {
		fmt.Printf("%c : %d\n", char, count)
	}
}
