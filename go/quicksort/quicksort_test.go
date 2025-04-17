package quicksort

import (
	"sort"
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

func FuzzQuicksort(f *testing.F) {

	f.Fuzz(func(t *testing.T, data []byte) {
		fz := fuzz.NewConsumer(data)
        var targetSlice []int
        if err := fz.CreateSlice(&targetSlice); err != nil {
			t.Skip()
		}

		oldLen := len(targetSlice)
		t.Log(targetSlice, oldLen)
		Quicksort(targetSlice, 0, oldLen - 1)

		// check if sorted
		if !sort.IntsAreSorted(targetSlice) {
			t.Errorf("not sorted: %v", targetSlice)
		}

		// check for length consistency
		if len(targetSlice) != oldLen {
			t.Errorf("length changed: before %d, after %d", oldLen, len(targetSlice))
		}
	})
}

func FuzzRandomQuicksort(f *testing.F) {

	f.Fuzz(func(t *testing.T, data []byte) {
		fz := fuzz.NewConsumer(data)
        var targetSlice []int
        if err := fz.CreateSlice(&targetSlice); err != nil {
			t.Skip()
		}

		oldLen := len(targetSlice)
		t.Log(targetSlice, oldLen)
		RandomQuicksort(targetSlice, 0, oldLen - 1)

		// check if sorted
		if !sort.IntsAreSorted(targetSlice) {
			t.Errorf("not sorted: %v", targetSlice)
		}

		// check for length consistency
		if len(targetSlice) != oldLen {
			t.Errorf("length changed: before %d, after %d", oldLen, len(targetSlice))
		}
	})
}
