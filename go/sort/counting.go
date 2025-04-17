package sort

import "slices"

func CountingSort(A []int) (B []int) {
	n := len(A)

	if n == 0 {
		return
	}

	maxValue := slices.Max(A)
	minValue := slices.Min(A)

	offset := 0
	if minValue < 0 {
		offset = -minValue
	}

	k := maxValue
	if minValue < 0 {
		k += offset
	}


	B = make([]int, n)
	C := make([]int, k + 1) // temporary storage, zero initialized

	for _, val := range A {
		C[val + offset]++ // count number of items == A[j]
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1] // add number of items < A[j]
	}

	for j := n - 1; j >= 0; j-- {
		B[C[A[j] + offset] - 1] = A[j]
		C[A[j] + offset]--  // handle duplicates
	}

	return B
}
