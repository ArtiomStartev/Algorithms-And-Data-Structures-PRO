package utils

import (
	"math"
	"math/rand"
	"slices"
)

type OrderType string

const (
	Ascending  OrderType = "Ascending"
	Descending OrderType = "Descending"
	Random     OrderType = "Random"
)

// GenerateArr generates an array of random integers of a given size and order type
func GenerateArr(size int, orderType OrderType) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Intn(math.MaxInt16)
	}

	switch orderType {
	case Ascending:
		slices.Sort(arr)
	case Descending:
		slices.SortFunc(arr, func(i, j int) int { return j - i })
	case Random:
		// Already filled with random numbers
	}

	return arr
}

// Unique returns a new array containing only the unique elements of the input array
func Unique(arr []int) []int {
	var result []int
	seen := make(map[int]bool)

	for _, val := range arr {
		if ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}

	return result
}
