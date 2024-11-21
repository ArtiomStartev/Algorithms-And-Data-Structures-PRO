package sort

import "time"

// QuickSorter is a type that implements the Sorter interface for the quick sort algorithm.
type QuickSorter struct{}

// Sort sorts the slice using the quick sort algorithm.
func (qs QuickSorter) Sort(slice []int) SortResult {
	// Record the start time of the sorting process
	startTime := time.Now()

	// Initialize the comparison and swap counters
	var comparisons, swaps int

	// Perform the quick sort algorithm
	sortedSlice := qs.quickSort(slice, &comparisons, &swaps)

	// Return the sorted slice and the counters
	return SortResult{
		Slice:         sortedSlice,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// quickSort is a helper function that performs the quick sort algorithm on a slice.
func (qs QuickSorter) quickSort(slice []int, comparisons, swaps *int) []int {
	// If the slice has 0 or 1 element, it is already sorted
	if len(slice) <= 1 {
		return slice
	}

	// Choose the pivot element as the middle element of the slice
	pivot := len(slice) / 2

	// Initialize the left and right sub-slices
	var left, right []int

	// Partition the slice into two sub-slices
	for i := 0; i < len(slice); i++ {
		*comparisons++ // Increment the comparison counter
		if i == pivot {
			continue // Skip the pivot element
		}

		// Append elements less than the pivot to the left sub-slice
		if slice[i] < slice[pivot] {
			left = append(left, slice[i])
		} else { // Append elements greater than or equal to the pivot to the right sub-slice
			right = append(right, slice[i])
		}
	}

	// Recursively sort the left and right sub-slices
	sortedLeft := qs.quickSort(left, comparisons, swaps)
	sortedRight := qs.quickSort(right, comparisons, swaps)

	*swaps += len(left) + len(right) // Count combining as swaps
	return append(append(sortedLeft, slice[pivot]), sortedRight...)
}
