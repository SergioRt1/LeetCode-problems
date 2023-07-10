package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList takes a head of a liked lists and reorder it L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 ...
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	list := iterativeToArray(head, 0)
	i, j := 0, len(list)-1
	for ; i < j; i, j = i+1, j-1 {
		list[i].Next, list[j].Next = list[j], list[i].Next
	}

	list[i].Next = nil
}

// iterativeToArray return an array with the pointers of a liked list
func iterativeToArray(head *ListNode, depth int) []*ListNode {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	return list
}

// recursiveToArray return an array with the pointers of a liked list
func recursiveToArray(head *ListNode, depth int) []*ListNode {
	if head == nil {
		return make([]*ListNode, depth)
	}
	list := recursiveToArray(head.Next, depth+1)
	list[depth] = head
	return list
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
	head := makeList([]int{1, 2, 3, 4, 5}, 0)
	reorderList(head)
	fmt.Print(head)
}
