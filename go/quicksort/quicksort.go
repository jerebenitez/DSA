package quicksort

import (
	"cmp"
	"errors"
	"log"
)

func partition[T cmp.Ordered](A []T, p int, r int) (int, error) {
	if r > len(A) || p > len(A) {
		return 0, errors.New("index out of bounds")
	}

	pivot := A[r]
	i := p - 1

	for j := p; j < r; j++ {
		if A[j] <= pivot {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}

	A[i+1], A[r] = A[r], A[i+1]

	return i + 1, nil
}

func Quicksort[T cmp.Ordered](A []T, p int, r int) {
	if p >= len(A) || r >= len(A) {
		log.Fatal("index out of range")
	}

	if p < r {
		if q, err := partition(A, p, r); err == nil {
			Quicksort(A, p, q - 1)
			Quicksort(A, q + 1, r)
		} else {
			log.Fatal(err.Error())
		}
	}
}
