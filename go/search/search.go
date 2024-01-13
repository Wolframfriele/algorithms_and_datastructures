package search

func LinearSearch(haystack []int, needle int) int {
	// If the needle is found, it will return the index
	// Else it will return -1
	// Runs in O(n)
	for i := range haystack {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}

func BinarySearch(haystack []int, needle int) int {
	l := 0
	h := len(haystack)
	for l < h {
		m := l + (h-l)/2
		switch {
		case haystack[m] == needle:
			return m
		case haystack[m] < needle:
			l = m + 1
		default:
			h = m
		}
	}
	return -1
}
