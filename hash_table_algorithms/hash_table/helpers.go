package hash_table

type HashFunc func(string) int

// HashString is a simple hash function that returns the sum of the ASCII values of the characters in the string
func HashString(s string) int {
	var hash int
	for _, c := range s {
		hash += int(c)
	}
	return hash
}
