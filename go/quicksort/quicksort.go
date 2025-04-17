package quicksort

import (
	"cmp"
	"errors"
	"log"
	"math/rand"
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

func randomPartition[T cmp.Ordered](A []T, p int, r int) (int, error) {
	i :=  rand.Intn(r - p) + p
	A[i], A[r] = A[r], A[i]

	return partition(A, p, r)
}

func Quicksort[T cmp.Ordered](A []T, p int, r int) {
	if r >= len(A) {
		log.Fatalf("index out of range: p=%d, r=%d", p, r)
	}

	if len(A) < 2 {
		return
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


func RandomQuicksort[T cmp.Ordered](A []T, p int, r int) {
	if r >= len(A) {
		log.Fatal("index out of range")
	}

	if p < r {
		if q, err := randomPartition(A, p, r); err == nil {
			RandomQuicksort(A, p, q - 1)
			RandomQuicksort(A, q + 1, r)
		} else {
			log.Fatal(err.Error())
		}
	}
}
