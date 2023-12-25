package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	size := 0
	for node := head; node != nil; node = node.Next {
		size++
	}
	merged := &ListNode{}
	return mergeSort(head, size, merged)
}

func mergeSort(head *ListNode, size int, merged *ListNode) *ListNode {
	if size <= 1 {
		return head
	}
	mid := size / 2
	var left, right *ListNode

	left = head
	for i := 0; i < mid-1; i++ {
		head = head.Next
	}
	right, head.Next = head.Next, nil

	l := mergeSort(left, mid, merged)
	r := mergeSort(right, size-mid, merged)

	head = merged

	for l != nil && r != nil {
		if l.Val < r.Val {
			merged.Next = l
			l = l.Next
		} else {
			merged.Next = r
			r = r.Next
		}
		merged = merged.Next
	}
	for l != nil {
		merged.Next = l
		l = l.Next
		merged = merged.Next
	}
	for r != nil {
		merged.Next = r
		r = r.Next
		merged = merged.Next
	}

	return head.Next
}

func main() {
	head := makeList([]int{4, 2, 1, 3}, 0)
	sortList(head)
}

func makeList(l []int, i int) *ListNode {
	if i >= len(l) {
		return nil
	}
	return &ListNode{
		Val:  l[i],
		Next: makeList(l, i+1),
	}
}
