package main

import (
	"fmt"
)

// calculateSum calculates the sum of elements in a portion of the array
func calculateSum(arr []int, start, end int, ch chan<- int) {
	sum := 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}
	ch <- sum
}

func main() {
	// Define the size of the array and the number of goroutines to use
	arrSize := 1000000
	numGoroutines := 4

	// Create a large array
	arr := make([]int, arrSize)
	for i := 0; i < arrSize; i++ {
		arr[i] = i + 1
	}

	// Create a channel to receive partial sums
	ch := make(chan int, numGoroutines)

	// Calculate the sum of elements concurrently using goroutines
	partSize := arrSize / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if i == numGoroutines-1 {
			end = arrSize
		}
		go calculateSum(arr, start, end, ch)
	}

	// Receive and accumulate the partial sums
	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-ch
	}

	fmt.Println("Sum of elements in the array:", totalSum)
}
