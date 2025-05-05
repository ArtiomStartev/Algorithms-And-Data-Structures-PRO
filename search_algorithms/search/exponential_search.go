package search

type ExponentialSearcher struct{}

// Search performs an exponential search on a sorted array of integers
func (ExponentialSearcher) Search(data any, x int) bool {
	arr, ok := data.([]int)
	if !ok {
		return false
	}

	if arr[0] == x {
		return true
	}

	i := 1
	for i < len(arr) && arr[i] <= x {
		if i >= len(arr)/2 { // Check for possible overflow
			i = len(arr)
			break
		}
		i *= 2
	}

	return recursiveBinarySearch(arr, x, i/2, i)
}

// recursiveBinarySearch performs a recursive binary search on a sorted array of integers
func recursiveBinarySearch(arr []int, x, low, high int) bool {
	mid := low + (high-low)/2

	if arr[mid] == x {
		return true
	}

	if arr[mid] > x {
		return recursiveBinarySearch(arr, x, low, mid-1)
	}

	if arr[mid] < x {
		return recursiveBinarySearch(arr, x, mid+1, high)
	}

	return false
}
