package sort

import "time"

// SelectionSorter is a type that implements the Sorter interface.
type SelectionSorter struct{}

// Sort sorts the slice using the selection sort algorithm.
func (SelectionSorter) Sort(slice []int) SortResult {
	// Record the start time of the sorting process
	startTime := time.Now()

	// Initialize the comparison and swap counters
	var comparisons, swaps int

	// Outer loop iterates through each element of the slice
	for i := 0; i < len(slice); i++ {
		minIndex := i // Initialize minIndex to the current index i

		// Inner loop iterates through the unsorted part of the slice
		for j := i + 1; j < len(slice); j++ {
			comparisons++ // Increment the comparison counter

			// Update minIndex if a smaller element is found
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}

		// If minIndex is different from i, swap the elements at indices i and minIndex
		if i != minIndex {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
			swaps++ // Increment the swap counter
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
