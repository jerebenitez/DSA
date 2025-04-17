package sort

import (
	"slices"
	"testing"
)

func TestCountingSort(t *testing.T) {
	testcases := []struct {
		in, want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{0, 0, -1}, []int{-1, 0, 0}},
	}

	for _, tc := range testcases {
		sorted := CountingSort(tc.in)

		if !slices.Equal(sorted, tc.want) {
			t.Errorf("Sort: %v, want %v", sorted, tc.want)
		}
	}
}
