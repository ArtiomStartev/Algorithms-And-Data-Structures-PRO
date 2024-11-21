package search

type BinarySearcher struct{}

// Search searches for the target in the slice and returns true if the target is found, otherwise false
func (BinarySearcher) Search(data any, target int) bool {
	// Type assertion to convert the interface{} type to a slice of integers
	slice, ok := data.([]int)
	if !ok {
		return false
	}

	var start int
	end := len(slice) - 1

	for start <= end {
		// Find the middle index
		middle := start + (end-start)/2

		// If target is found at a middle, return true
		if slice[middle] == target {
			return true
		}

		// If the target is less than the middle element, search in the left half
		if slice[middle] > target {
			end = middle - 1
		}

		// If the target is greater than the middle element, search in the right half
		if slice[middle] < target {
			start = middle + 1
		}
	}

	// If the target is not found, return false
	return false
}
