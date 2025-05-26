package MaxHeap

import "container/heap"

/**
本库的Heap的封装实现使用了container/heap中的heap接口
*/
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
