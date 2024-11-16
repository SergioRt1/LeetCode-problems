package main

import (
	"fmt"
)

// Heap implementation

type CustomHeap []int

func Init(l []int) CustomHeap {
	for i := len(l)/2 - 1; i >= 0; i-- {
		shiftDown(l, i)
	}

	return l
}

func (c *CustomHeap) Add(val int) {
	*c = append(*c, val)
	last := len(*c) - 1
	shiftUp(*c, last)
}

func (c *CustomHeap) Pop() int {
	slice := *c
	current := slice[0]
	last := len(slice) - 1
	slice[0] = slice[last]
	slice = slice[:last]
	shiftDown(slice, 0)

	*c = slice

	return current
}

func getParent(i int) int {
	return (i - 1) / 2
}

func shiftUp(h CustomHeap, idx int) {
	for parent := getParent(idx); idx != parent && h[parent] >= h[idx]; idx, parent = parent, getParent(parent) { // parent
		h[parent], h[idx] = h[idx], h[parent]
	}
}

func shiftDown(h CustomHeap, root int) {
	left := root*2 + 1
	for left < len(h) {
		minChild := left
		if right := left + 1; right < len(h) && h[right] < h[left] {
			minChild = right
		}

		if h[root] > h[minChild] {
			h[root], h[minChild] = h[minChild], h[root]
		} else {
			break
		}
		root = minChild
		left = root*2 + 1
	}
}

// Problem solution

type KthLargest struct {
	heap CustomHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := Init(nums)

	for len(h) > k {
		h.Pop()
	}

	return KthLargest{
		heap: h,
		k:    k,
	}
}

func (k *KthLargest) Add(val int) int {
	k.heap.Add(val)
	if len(k.heap) > k.k {
		k.heap.Pop()
	}

	return k.heap[0]
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
