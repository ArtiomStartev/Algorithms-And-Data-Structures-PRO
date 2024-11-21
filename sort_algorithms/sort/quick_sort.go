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
	qs.quickSort(slice, &comparisons, &swaps)

	// Return the sorted slice and the counters
	return SortResult{
		Slice:         slice,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// quickSort sorts the slice using the quick sort algorithm.
func (qs QuickSorter) quickSort(slice []int, comparisons, swaps *int) {
	// If the slice has 0 or 1 element, it is already sorted
	if len(slice) <= 1 {
		return
	}

	// Partition the slice and get the pivot index
	pivotIndex := qs.partition(slice, comparisons, swaps)

	// Recursively sort elements before and after the pivot
	qs.quickSort(slice[:pivotIndex], comparisons, swaps)
	qs.quickSort(slice[pivotIndex+1:], comparisons, swaps)
}

// partition partitions the slice and returns the pivot index.
func (qs QuickSorter) partition(slice []int, comparisons, swaps *int) int {
	pivot := slice[len(slice)-1] // Choose the last element as the pivot
	i := -1                      // Index for the smaller element

	for j := 0; j < len(slice)-1; j++ {
		*comparisons++

		// If the current element is smaller than or equal to the pivot
		if slice[j] <= pivot {
			i++

			// Swap elements
			slice[i], slice[j] = slice[j], slice[i]
			*swaps++
		}
	}

	// Place the pivot in the correct position
	slice[i+1], slice[len(slice)-1] = slice[len(slice)-1], slice[i+1]
	*swaps++

	return i + 1 // Return the index of the pivot element
}
