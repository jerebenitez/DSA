package sort

func CountingSort(A []int, k int) (B []int) {
	n := len(A)

	if n == 0 {
		return
	}

	B = make([]int, n)
	C := make([]int, k + 1) // temporary storage, zero initialized

	for _, val := range A {
		C[val]++ // count number of items == A[j]
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1] // add number of items < A[j]
	}

	for j := n - 1; j >= 0; j-- {
		B[C[A[j]] - 1] = A[j]
		C[A[j]]--  // handle duplicates
	}

	return B
}
