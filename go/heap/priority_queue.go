package heap

import "errors"

type Element[T comparable] struct {
	Key int
	Object T
}

func maxHeapify[T comparable](A []Element[T], idx int) {
	// Go is 0-indexed
	l := 2 * idx + 1
	r := 2 * idx + 2

	largest := idx

	// assuming whole slice is a heap
	if l <= len(A) - 1 && A[l].Key > A[idx].Key {
		largest = l
	}

	if r <= len(A) - 1 && A[r].Key > A[largest].Key {
		largest = r
	}

	if largest != idx {
		A[idx] , A[largest] = A[largest], A[idx]
		maxHeapify(A, largest)
	}
}

func buildMaxHeap[T comparable](A []Element[T]) {
	for i := (len(A) / 2) - 1; i >= 0; i-- {
		maxHeapify(A, i)
	}
}

func heapSort[T comparable](A []Element[T]) {
	buildMaxHeap(A)

	for i := len(A) - 1; i >= 0; i-- {
		A[0], A[i] = A[i], A[0]
		A = A[0:i]
		maxHeapify(A, 0)
	}
}

func Insert[T comparable](A []Element[T], x Element[T]) ([]Element[T], error) {
	// No need to check if there's enough space in A since I'm working with slices

	k := x.Key
	// refer to https://stackoverflow.com/a/6878625
	x.Key = -(int(^uint(0) >> 1)) - 1

	A = append(A, x)
	if err := IncreaseKey(A, x, k); err != nil {
		return nil, err
	}

	return A, nil
}

func Maximum[T comparable](A []Element[T]) Element[T] {
	return A[0]
}

func ExtractMax[T comparable](A []Element[T]) Element[T] {
	maximum := Maximum(A)
	n := len(A) - 1
	A[0], A[n] = A[n], A[0]
	A = A[0:n]

	maxHeapify(A, 0)

	return maximum
}

func IncreaseKey[T comparable](A []Element[T], x Element[T], key int) error {
	if key < x.Key {
		return errors.New("new key is smaller than current key")
	}

	idx, err := findIndex(A, x)
	A[idx].Key = key

	if err != nil {
		return err
	}

	parent := (idx-1)/2

	for idx > 0 && A[parent].Key < A[idx].Key {
		A[idx], A[parent] = A[parent], A[idx]
		idx = parent
	}

	return nil	
}

func findIndex[T comparable](A []Element[T], x Element[T]) (int, error) {
	for idx, el := range A {
		// x has to point to the same memory location as el
		if el.Object == x.Object {
			return idx, nil
		}
	}

	return 0, errors.New("element not found in queue")
}
