package main

import "fmt"
import "github.com/jerebenitez/DSA/go/heap"

func main() {
	test := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	heap.MaxHeapify(test, 1)
	fmt.Printf("Result: %v\n", test)

	create := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap.BuildMaxHeap(create)
	fmt.Printf("Result: %v\n", create)

	heap.HeapSort(create)
	fmt.Printf("Sorted: %v\n", create)
}
