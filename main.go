package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE        = 100_000_000
	CHUNKS      = 8
	randomLimit = 1_000_000
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(randomLimit) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	result := data[0]
	for _, value := range data[1:] {
		if value > result {
			result = value
		}
	}
	return result
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunkCount := CHUNKS
	if len(data) < chunkCount {
		chunkCount = len(data)
	}
	maxima := make([]int, chunkCount)
	var group sync.WaitGroup
	for i := 0; i < chunkCount; i++ {
		start := i * len(data) / chunkCount
		end := (i + 1) * len(data) / chunkCount
		group.Add(1)
		go func(index int, chunk []int) {
			defer group.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	group.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	started := time.Now()
	result := maximum(data)
	elapsed := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", result, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	started = time.Now()
	parallelResult := maxChunks(data)
	elapsed = time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", parallelResult, elapsed)
}
