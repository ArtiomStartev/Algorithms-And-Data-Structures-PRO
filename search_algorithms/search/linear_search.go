package search

type LinearSearcher struct{}

// Search searches for the target in the slice and returns true if the target is found, otherwise false
func (LinearSearcher) Search(data any, target int) bool {
	// Type assertion to convert the interface{} type to a slice of integers
	slice, ok := data.([]int)
	if !ok {
		return false
	}

	// Iterate through each element of the slice
	for i := 0; i < len(slice); i++ {
		// If the current element is equal to the target, return true
		if slice[i] == target {
			return true
		}
	}

	// If the target element is not found, return false
	return false
}
