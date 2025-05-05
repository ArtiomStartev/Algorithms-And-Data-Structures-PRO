package search

type LinearSearcher struct{}

// Search performs a linear search on an array of integers
func (LinearSearcher) Search(data any, x int) bool {
	arr, ok := data.([]int)
	if !ok {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}

	return false
}
