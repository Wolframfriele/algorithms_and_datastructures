package search

import "math"

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
	// If the needle is found it will return the index
	// Else it will return -1
	// Runs in O(log(n))
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

func CrystalBal(heights []bool) int {
	// Given two crystal balls that break if dropped from high enough,
	// determine the exact spot in wich it will break in the most optimized way
	// So there are two chances to determine the breaking point.
	lastGood := 0
	for i := 0; i <= len(heights); i += int(math.Sqrt(float64(len(heights)))) {
		if heights[i] == true {
			break
		}
		lastGood = i
	}
	for j := lastGood; j <= len(heights); j++ {
		if heights[j] == true {
			return j
		}
	}
	return -1
}
