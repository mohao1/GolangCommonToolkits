package Heap

import "container/heap"

// 小顶堆MinHeap
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

// NewMinHeap 获得一个初始化的小顶堆
func NewMinHeap() *minHeap {
	NewHeap := new(minHeap)
	heap.Init(NewHeap)
	return NewHeap
}

// 大顶堆MaxHeap
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewMaxHeap 获得一个初始化的大顶堆
func NewMaxHeap() *maxHeap {
	NewHeap := new(maxHeap)
	heap.Init(NewHeap)
	return NewHeap
}
