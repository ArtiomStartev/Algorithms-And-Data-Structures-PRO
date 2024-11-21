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
	qs.quickSort(slice, 0, len(slice)-1, &comparisons, &swaps)

	// Return the sorted slice and the counters
	return SortResult{
		Slice:         slice,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}

// quickSort sorts the slice using the quick sort algorithm.
func (qs QuickSorter) quickSort(slice []int, low, high int, comparisons, swaps *int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := qs.partition(slice, low, high, comparisons, swaps)

		// Recursively sort elements before and after the pivot
		qs.quickSort(slice, low, pivotIndex-1, comparisons, swaps)
		qs.quickSort(slice, pivotIndex+1, high, comparisons, swaps)
	}
}

// partition partitions the slice and returns the pivot index.
func (qs QuickSorter) partition(slice []int, low, high int, comparisons, swaps *int) int {
	pivot := slice[high] // Choose the last element as the pivot
	i := low - 1         // Index for the smaller element

	for j := low; j < high; j++ {
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
	slice[i+1], slice[high] = slice[high], slice[i+1]
	*swaps++

	return i + 1 // Return the index of the pivot element
}
