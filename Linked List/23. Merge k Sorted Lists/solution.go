package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

const inf = 1 << 30

// mergeKLists Merge all the linked-lists into one sorted linked-list and return it.
func mergeKLists(lists []*ListNode) *ListNode {
	return recursiveMerge(lists)
}

// recursiveMerge do mergeKLists in O(n*log2(k)) using a mergeSort strategy
func recursiveMerge(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := recursiveMerge(lists[:mid])
	right := recursiveMerge(lists[mid:])

	head := new(ListNode)
	prev := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			prev.Next = left
			left = left.Next
		} else {
			prev.Next = right
			right = right.Next
		}
		prev = prev.Next
	}
	for left != nil {
		prev.Next = left
		left = left.Next
		prev = prev.Next
	}
	for right != nil {
		prev.Next = right
		right = right.Next
		prev = prev.Next
	}

	return head.Next
}

// simpleMerge do mergeKLists in O(n*k).
func simpleMerge(lists []*ListNode) *ListNode {
	head := new(ListNode)
	prev := head
	var bestIdx int
	var bestVal int
	notDone := true
	for notDone {
		notDone = false
		bestVal = inf
		var list *ListNode
		for idx := 0; idx < len(lists); idx++ {
			list = lists[idx]
			if list == nil {
				lists = fastRemove(lists, idx)
				idx--
			} else if bestVal > list.Val {
				bestVal = list.Val
				bestIdx = idx
				notDone = true
			}
		}
		if notDone {
			prev.Next = lists[bestIdx]
			lists[bestIdx] = lists[bestIdx].Next
			prev = prev.Next
		}
	}
	return head.Next
}

func fastRemove(l []*ListNode, idx int) []*ListNode {
	last := len(l) - 1
	l[idx], l = l[last], l[:last]
	return l
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

func main() {
	lists := []*ListNode{
		makeList([]int{1, 4, 5}, 0),
		makeList([]int{1, 3, 4}, 0),
		makeList([]int{2, 6}, 0),
	}
	fmt.Print(mergeKLists(lists))
}
