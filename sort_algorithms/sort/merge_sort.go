package sort

import "time"

// MergeSorter is a type that implements the Sorter interface.
type MergeSorter struct{}

// Sort sorts the slice using the merge sort algorithm.
func (ms MergeSorter) Sort(slice []int) SortResult {
	// Record the start time of the sorting process
	startTime := time.Now()

	// Initialize the comparison and swap counters
	var comparisons, swaps int

	// Perform the merge sort algorithm
	sortedSlice := ms.mergeSort(slice, &comparisons, &swaps)

	// Return the sorted slice and the counters
	return SortResult{
		Slice:         sortedSlice,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// mergeSort is a helper function that performs the merge sort algorithm on a slice.
func (ms MergeSorter) mergeSort(slice []int, comparisons, swaps *int) []int {
	// If the slice has 0 or 1 elements, it is already sorted
	if len(slice) <= 1 {
		return slice
	}

	// Divide the slice into two halves
	mid := len(slice) / 2
	left := ms.mergeSort(slice[:mid], comparisons, swaps)  // Recursively sort the left half
	right := ms.mergeSort(slice[mid:], comparisons, swaps) // Recursively sort the right half

	// Merge the two sorted halves
	return ms.merge(left, right, comparisons, swaps)
}

// merge is a helper function that merges two sorted slices into a single sorted slice.
func (ms MergeSorter) merge(left, right []int, comparisons, swaps *int) []int {
	// Create a new slice to store the sorted elements
	sorted := make([]int, 0, len(left)+len(right))

	// Compare the first elements of the two slices and add the smaller one to the sorted slice
	for len(left) > 0 || len(right) > 0 {
		// If the left slice is empty, add the remaining elements of the right slice to the sorted slice
		if len(left) == 0 {
			*swaps += len(right) // Count appending as swaps
			return append(sorted, right...)
		}

		// If the right slice is empty, add the remaining elements of the left slice to the sorted slice
		if len(right) == 0 {
			*swaps += len(left) // Count appending as swaps
			return append(sorted, left...)
		}

		*comparisons++ // Increment the comparison counter

		// Compare the first elements of the two slices
		if left[0] <= right[0] {
			*swaps++ // Increment the swap counter
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			*swaps++ // Increment the swap counter
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	return sorted // Return the sorted slice
}
