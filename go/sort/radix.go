package sort

import "slices"

func expCountingSort(A []int, k int, getIndex func (int) int) (B []int) {
	n := len(A)
	B = A
	if n < 2 {
		return
	}

	B = make([]int, n)
	C := make([]int, k + 1) // temporary storage, zero initialized

	for _, val := range A {
		C[getIndex(val)]++ // count number of items == A[j]
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1] // add number of items < A[j]
	}

	for j := n - 1; j >= 0; j-- {
		B[C[getIndex(A[j])] - 1] = A[j]
		C[getIndex(A[j])]--  // handle duplicates
	}

	return B
}

func RadixSort(A []int) (B []int) {
	B = A
	if len(A) < 2 {
		return
	}

	// maxValue is just 9 since we're counting a single digit
	maxValue := 9
	minValue := slices.Min(A)

	offset := 0
	if minValue < 0 {
		offset = -minValue
	}

	k := maxValue
	if minValue < 0 {
		k += offset
	}

	m := slices.Max(A) + offset

	for exp := 1; m/exp > 0; exp *= 10 {
		B = expCountingSort(B, k, func (i int) int {
			i += offset
			return (i / exp) % 10
		})
	}

	return B
}
