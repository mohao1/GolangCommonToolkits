package Heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	NewHeap := NewMaxHeap()
	heap.Push(NewHeap, 10)
	heap.Push(NewHeap, 20)
	heap.Push(NewHeap, 30)
	heap.Push(NewHeap, 6)
	heap.Push(NewHeap, 7)
	for i := 0; i < 5; i++ {
		fmt.Println(heap.Pop(NewHeap))
	}
}

func TestMinHeap(t *testing.T) {
	NewHeap := NewMinHeap()
	heap.Push(NewHeap, 10)
	heap.Push(NewHeap, 20)
	heap.Push(NewHeap, 30)
	heap.Push(NewHeap, 6)
	heap.Push(NewHeap, 7)
	for i := 0; i < 5; i++ {
		fmt.Println(heap.Pop(NewHeap))
	}
}
