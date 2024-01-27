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

type DropTestCase struct {
	breaks []bool
	want   int
}

var dropTestCases = []DropTestCase{
	{breaks: []bool{false, false, true}, want: 2},
	{breaks: []bool{false, false, false, false, false, false, false, false, false, false, true, true, true}, want: 10},
}

func TestCrystalBall(t *testing.T) {
	for _, test := range dropTestCases {
		got := CrystalBal(test.breaks)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}
