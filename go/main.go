package main

import (
	"fmt"
	"os"

	"github.com/jerebenitez/DSA/go/heap"
	"github.com/jerebenitez/DSA/go/quicksort"
)

func main() {
	test := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	heap.MaxHeapify(test, 1)
	fmt.Printf("Result: %v\n", test)

	create := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap.BuildMaxHeap(create)
	fmt.Printf("Result: %v\n", create)

	heap.HeapSort(create)
	fmt.Printf("Sorted: %v\n", create)

	/***********************/
	// Build a priority queue for strings
	queue := make([]heap.Element[string], 0)

	// Create an element with a priority of 1
	// and insert it into queue
	el1 := heap.Element[string]{ 
		Key: 1,
		Object: "test",
	}
	var err error
	if queue, err = heap.Insert(queue, el1); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	fmt.Println(queue)

	// Create an element with a higher priority
	// and insert it into queue
	el2 := heap.Element[string]{
		Key: 10,
		Object: "higher",
	}
	if queue, err = heap.Insert(queue, el2); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	fmt.Println(queue)

	/**********************/
	A := []int{15, 24, 6, 8, 99, 106, 2, 392, 94}
	quicksort.Quicksort(A, 0, len(A) - 1)
	fmt.Printf("Sorted: %v\n", A)

	B := []int{15, 24, 6, 8, 99, 106, 2, 392, 94}
	quicksort.RandomQuicksort(B, 0, len(B) - 1)
	fmt.Printf("Sorted: %v\n", B)
}
