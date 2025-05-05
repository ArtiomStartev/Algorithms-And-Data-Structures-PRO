package sort

import "time"

type SelectionSorter struct{}

// Sort performs a selection sort on the given array of integers
func (SelectionSorter) Sort(arr []int) SortResult {
	startTime := time.Now()
	var comparisons, swaps int

	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			comparisons++

			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
			swaps++
		}
	}

	return SortResult{
		Arr:           arr,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}
