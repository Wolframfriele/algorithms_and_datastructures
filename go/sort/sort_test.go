package sort

import (
	"testing"
)

type SearchTestCase struct {
	sortable []int
	want     []int
}

var testCases = []SearchTestCase{
	{sortable: []int{2, 4, 1, 7, 3}, want: []int{1, 2, 3, 4, 7}},
	{sortable: []int{1, 2, 3, 4, 7}, want: []int{1, 2, 3, 4, 7}},
	{sortable: []int{7, 5, 4, 3, 1}, want: []int{1, 3, 4, 5, 7}},
}

func TestBubbleSort(t *testing.T) {
	for _, test := range testCases {
		got := append([]int{}, test.sortable...)
		BubbleSort(got)
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("Got %v, want %v", got[i], test.want[i])
			}
		}
	}
}
