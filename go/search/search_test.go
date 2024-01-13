package search

import "testing"

type SearchTestCase struct {
	haystack []int
	needle   int
	want     int
}

var searchTests = []SearchTestCase{
	{haystack: []int{1, 2, 3}, needle: 2, want: 1},
	{haystack: []int{1, 2, 3}, needle: 4, want: -1},
	{haystack: []int{1, 2, 3, 5, 6, 8, 20, 21, 22, 25, 70}, needle: 25, want: 9},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		got := LinearSearch(test.haystack, test.needle)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		got := BinarySearch(test.haystack, test.needle)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}
