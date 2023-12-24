package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	//count length of list
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}
	groupSize := l / k
	remainder := l % k

	// generate response
	response := make([]*ListNode, k)
	idx := 0
	currentSize := 0
	var endList bool
	for node := head; node != nil; {
		if currentSize == 0 {
			response[idx] = node
		}
		currentSize++

		if remainder <= idx {
			endList = currentSize == groupSize
		} else {
			endList = currentSize == groupSize+1
		}

		if endList {
			currentSize = 0
			idx++
			node, node.Next = node.Next, nil
		} else {
			node = node.Next
		}
	}

	return response
}

func makeList(list []int, idx int) *ListNode {
	if idx >= len(list) {
		return nil
	}
	return &ListNode{
		Val:  list[idx],
		Next: makeList(list, idx+1),
	}
}

func main() {
	l := makeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0)
	k := 3

	fmt.Println(splitListToParts(l, k))
}
