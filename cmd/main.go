package main

import (
	"fmt"

	"Heap"
)

func main() {
	heap := Heap.NewHeap(Heap.MAX)
	heap.Insert(10)
	heap.Insert(9)
	heap.Insert(8)
	heap.Insert(7)
	heap.Insert(15)
	heap.Insert(12)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(1)
	heap.Print()
	fmt.Println(heap.Sort())
}
