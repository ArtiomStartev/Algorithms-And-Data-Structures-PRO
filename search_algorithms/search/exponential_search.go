package search

type ExponentialSearcher struct{}

// Search searches for the target in the slice and returns true if the target is found, otherwise false
func (ExponentialSearcher) Search(data any, target int) bool {
	// Type assertion to convert the interface{} type to a slice of integers
	slice, ok := data.([]int)
	if !ok {
		return false
	}

	// Check the first element of the slice and return true if it matches the target
	if slice[0] == target {
		return true
	}

	// Find the range for binary search by repeated doubling
	i := 1
	for i < len(slice) && slice[i] <= target {
		i *= 2
	}

	// Perform binary search in the range [i/2, min(i, len(slice))]
	return recursiveBinarySearch(slice, target, i/2, min(i, len(slice)))
}

// recursiveBinarySearch is a helper function that performs binary search recursively
func recursiveBinarySearch(slice []int, target, start, end int) bool {
	// Find the middle index
	middle := (start + end) / 2

	// If target is found at a middle, return true
	if slice[middle] == target {
		return true
	}

	// If the target is less than the middle element, search in the left half
	if slice[middle] > target {
		return recursiveBinarySearch(slice, target, start, middle-1)
	}

	// If the target is greater than the middle element, search in the right half
	if slice[middle] < target {
		return recursiveBinarySearch(slice, target, middle+1, end)
	}

	// If the target is not found, return false
	return false
}
