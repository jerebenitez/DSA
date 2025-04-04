package heap

import "cmp"

func MaxHeapify[T cmp.Ordered](A []T, idx int) {
	// Go is 0-indexed
	l := 2 * idx + 1
	r := 2 * idx + 2

	largest := idx

	// assuming whole slice is a heap
	if l <= len(A) - 1 && A[l] > A[idx] {
		largest = l
	}

	if r <= len(A) - 1 && A[r] > A[largest] {
		largest = r
	}

	if largest != idx {
		A[idx] , A[largest] = A[largest], A[idx]
		MaxHeapify(A, largest)
	}
}

func BuildMaxHeap[T cmp.Ordered](A []T) {
	for i := (len(A) / 2) - 1; i >= 0; i-- {
		MaxHeapify(A, i)
	}
}

func HeapSort[T cmp.Ordered](A []T) {
	BuildMaxHeap(A)

	for i := len(A) - 1; i >= 0; i-- {
		A[0], A[i] = A[i], A[0]
		A = A[0:i]
		MaxHeapify(A, 0)
	}
}
