package sort

import "time"

// BubbleSorter is a type that implements the Sorter interface.
type BubbleSorter struct{}

// Sort sorts the slice using the bubble sort algorithm.
func (BubbleSorter) Sort(slice []int) SortResult {
	// Record the start time of the sorting process
	startTime := time.Now()

	// Initialize the comparison and swap counters
	var comparisons, swaps int

	// Outer loop iterates through each element of the slice
	for i := 0; i < len(slice); i++ {
		// The last i elements are already sorted, so we don't need to compare them
		for j := 0; j < len(slice)-i-1; j++ {
			comparisons++ // Increment the comparison counter

			// Compare adjacent elements and swap them if they are in the wrong order
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j] // Swap the elements
				swaps++                                     // Increment the swap counter
			}
		}
	}

	// Return the sorted slice and the counters
	return SortResult{
		Slice:         slice,
		Comparisons:   comparisons,
		Swaps:         swaps,
		ExecutionTime: time.Since(startTime),
	}
}
