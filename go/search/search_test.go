package search

import "testing"

type LinearSearchTestCase struct {
	haystack []int
	needle   int
	want     bool
}

func TestLinearSearch(t *testing.T) {
	linearSearchTests := []LinearSearchTestCase{
		{haystack: []int{1, 2, 3}, needle: 2, want: true},
		{haystack: []int{1, 2, 3}, needle: 4, want: false},
	}
	for _, test := range linearSearchTests {
		got := LinearSearch(test.haystack, test.needle)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}
