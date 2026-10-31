package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	nums := make([]int, size)

	for i := range size {
		nums[i] = rand.Int()
	}

	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	max := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// если в data меньше 8 (CHUNKS) элементов, то нет смысла делить его на 8 частей
	// можно сразу найти максимум через maximum(data)
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	values := make([]int, CHUNKS)
	chunk := len(data) / CHUNKS

	for j := 0; j < CHUNKS; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			start := j * chunk
			end := (j + 1) * chunk

			// если размер слайса не кратен 8, то элементы в "остатке" не проверятся
			// чтобы избежать этого, для последнего среза end равен длине исходного слайса
			// тогда если размер слайса не кратен 8, "остаток" будет включен в последний срез
			if j == CHUNKS-1 {
				end = len(data)
			}

			// для поиска максимума в срезе используется функция maximum()
			values[j] = maximum(data[start:end])
		}(j)
	}
	wg.Wait()

	return maximum(values)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	t := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(t)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(t)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
