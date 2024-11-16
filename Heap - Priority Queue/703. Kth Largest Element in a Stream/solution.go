package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	obj := KthLargest{
		heap: &h,
		k:    k,
	}
	heap.Init(obj.heap)
	for obj.heap.Len() > k {
		heap.Pop(obj.heap)
	}

	return obj
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k.heap, val)
	if k.heap.Len() > k.k {
		heap.Pop(k.heap)
	}

	return (*k.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}
