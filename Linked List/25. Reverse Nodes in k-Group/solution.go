package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return recursiveInit(head, k)
}

func recursiveInit(first *ListNode, k int) *ListNode {
	if k == 1 {
		return first
	}

	preGroup := &ListNode{
		Next: first,
	}
	head := preGroup

	for first != nil && first.Next != nil {
		nextGroup, ok := recursive(first.Next, preGroup, first, k-1)
		if !ok {
			preGroup.Next = first
			break
		}
		first.Next = nextGroup

		first, preGroup = nextGroup, first
	}

	return head.Next
}

func simpleReverse(node, prev *ListNode) *ListNode {
	if node == nil {
		return prev
	}
	last := simpleReverse(node.Next, node)
	node.Next = prev

	return last
}

func recursive(node, preGroup, prev *ListNode, k int) (nextGroup *ListNode, ok bool) {
	if node == nil && k >= 0 {
		return nil, false
	}
	if k == 1 {
		nextGroup = node.Next
		ok = true
		preGroup.Next = node
		node.Next = prev
		return
	}
	nextGroup, ok = recursive(node.Next, preGroup, node, k-1)
	if !ok {
		return
	}

	node.Next = prev

	return
}

func makeList(l []int, idx int) *ListNode {
	if idx >= len(l) {
		return nil
	}
	return &ListNode{
		Val:  l[idx],
		Next: makeList(l, idx+1),
	}
}

func printList(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(",")
		}

		head = head.Next
	}
	fmt.Println("]")
}

func main() {
	list := makeList([]int{1, 2, 3, 4, 5}, 0)
	k := 1
	printList(list)
	newList := reverseKGroup(list, k)
	printList(newList)
}
