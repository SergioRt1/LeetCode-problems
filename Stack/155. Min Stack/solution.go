package main

import "fmt"

type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		mins: make([]int, 0),
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	lastElem := len(s.mins) - 1
	min := val
	if lastElem >= 0 {
		if min > s.mins[lastElem] {
			min = s.mins[lastElem]
		}
	}

	s.mins = append(s.mins, min)
}

func (s *MinStack) Pop() {
	lastElem := len(s.mins) - 1
	s.data = s.data[:lastElem]
	s.mins = s.mins[:lastElem]
}

func (s *MinStack) Top() int {
	lastElem := len(s.data) - 1
	if lastElem < 0 {
		return 0
	}

	return s.data[lastElem]
}

func (s *MinStack) GetMin() int {
	lastElem := len(s.mins) - 1
	if lastElem < 0 {
		return 0
	}

	return s.mins[lastElem]
}

func main() {
	//Your MinStack object will be instantiated and called as such:
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin()) // return -3
	obj.Pop()
	fmt.Println(obj.Top())    // return 0
	fmt.Println(obj.GetMin()) // return -2
}
