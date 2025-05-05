package sort

import "time"

type BubbleSorter struct{}

// Sort performs a bubble sort on the given array of integers
func (BubbleSorter) Sort(arr []int) SortResult {
	startTime := time.Now()
	var comparisons, swaps int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			comparisons++

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
			}
		}
	}

	return SortResult{
		Arr:           arr,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}
