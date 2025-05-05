package sort

import "time"

type QuickSorter struct{}

// Sort performs a quick sort on the given array of integers
func (qs QuickSorter) Sort(arr []int) SortResult {
	startTime := time.Now()
	var comparisons, swaps int

	qs.quickSort(arr, &comparisons, &swaps)

	return SortResult{
		Arr:           arr,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// quickSort recursively sorts the array using the quick sort algorithm
func (qs QuickSorter) quickSort(arr []int, comparisons, swaps *int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := qs.partition(arr, comparisons, swaps)

	qs.quickSort(arr[:pivotIndex], comparisons, swaps)
	qs.quickSort(arr[pivotIndex+1:], comparisons, swaps)
}

// partition partitions the array around a pivot element
func (qs QuickSorter) partition(arr []int, comparisons, swaps *int) int {
	pivotIndex := len(arr) - 1
	i := -1

	for j := 0; j < pivotIndex; j++ {
		*comparisons++

		if arr[j] <= arr[pivotIndex] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			*swaps++
		}
	}

	arr[i+1], arr[pivotIndex] = arr[pivotIndex], arr[i+1]
	*swaps++

	return i + 1
}
