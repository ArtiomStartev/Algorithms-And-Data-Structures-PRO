package sort

import "time"

type SortResult struct {
	Arr           []int
	Comparisons   int
	Swaps         int
	ExecutionTime time.Duration
}

type Sorter interface {
	Sort(arr []int) SortResult
}
