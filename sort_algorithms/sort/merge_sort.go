package sort

import "time"

type MergeSorter struct{}

// Sort performs a merge sort on the given array of integers
func (ms MergeSorter) Sort(arr []int) SortResult {
	startTime := time.Now()
	var comparisons, swaps int

	sortedArr := ms.mergeSort(arr, &comparisons, &swaps)

	return SortResult{
		Arr:           sortedArr,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// mergeSort recursively sorts the array using the merge sort algorithm
func (ms MergeSorter) mergeSort(arr []int, comparisons, swaps *int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := ms.mergeSort(arr[:mid], comparisons, swaps)
	right := ms.mergeSort(arr[mid:], comparisons, swaps)

	return ms.merge(left, right, comparisons, swaps)
}

// merge is a helper function that merges two sorted arrays into one sorted array
func (ms MergeSorter) merge(left, right []int, comparisons, swaps *int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	var i, j int

	for i < len(left) && j < len(right) {
		*comparisons++

		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}

		*swaps++
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
		*swaps++
	}

	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
		*swaps++
	}

	return sorted
}
