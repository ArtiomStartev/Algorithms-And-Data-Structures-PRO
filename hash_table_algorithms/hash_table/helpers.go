package hash_table

// HashString is a simple hash function that returns the sum of the ASCII values of the characters in the string
func HashString(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
	}

	return hash
}
