package sort

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
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
		testname := fmt.Sprintf("%v", tc.in)
		t.Run(testname, func(t *testing.T) {
			sorted := CountingSort(tc.in)

			if !slices.Equal(sorted, tc.want) {
				t.Errorf("Sort: %v, want %v", sorted, tc.want)
			}
		})
	}
}

func FuzzCountingSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		fz := fuzz.NewConsumer(data)

		var targetSlice []int
		if err := fz.CreateSlice(&targetSlice); err != nil {
			t.Skip()
		}
		
		sorted := CountingSort(targetSlice)

		// check if sorted
		if !sort.IntsAreSorted(sorted) {
			t.Errorf("not sorted: %v", sorted)
		}

		// check for length consistency
		if len(targetSlice) != len(sorted) {
			t.Errorf("length changed: before %d, after %d", len(sorted), len(targetSlice))
		}
	})
}
