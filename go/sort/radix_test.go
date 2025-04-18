package sort

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

func TestRadixSort(t *testing.T) {
	testcases := []struct {
		in, want []int
	}{
		{[]int{}, []int{}},
		{[]int{10, 21, 23}, []int{10, 21, 23}},
		{[]int{30, 21, 14}, []int{14, 21, 30}},
		{[]int{0, 0, -1}, []int{-1, 0, 0}},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("%v", tc.in)
		t.Run(testname, func(t *testing.T) {
			sorted := RadixSort(tc.in)

			if !slices.Equal(sorted, tc.want) {
				t.Errorf("Sort: %v, want %v", sorted, tc.want)
			}
		})
	}
}

func FuzzRadixSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		fz := fuzz.NewConsumer(data)

		var targetSlice []int
		if err := fz.CreateSlice(&targetSlice); err != nil {
			t.Skip()
		}
		
		sorted := RadixSort(targetSlice)

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
