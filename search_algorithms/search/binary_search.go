package search

type BinarySearcher struct{}

// Search performs a binary search on a sorted array of integers
func (BinarySearcher) Search(data any, x int) bool {
	arr, ok := data.([]int)
	if !ok {
		return false
	}

	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return true
		}

		if arr[mid] > x {
			high = mid - 1
		}

		if arr[mid] < x {
			low = mid + 1
		}
	}

	return false
}
