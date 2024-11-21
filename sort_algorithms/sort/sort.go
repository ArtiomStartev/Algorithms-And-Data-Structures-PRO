package sort

import "time"

// SortResult is a struct that represents the result of a sorting operation.
type SortResult struct {
	Slice         []int         // The sorted slice
	Comparisons   int           // Number of comparisons made
	Swaps         int           // Number of swaps performed
	ExecutionTime time.Duration // Execution time for sorting
}

// Sorter is an interface that defines the behavior of a sorting algorithm.
type Sorter interface {
	Sort(slice []int) SortResult
}
