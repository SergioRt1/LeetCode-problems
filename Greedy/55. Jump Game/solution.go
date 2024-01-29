package main

import "fmt"

func canJump(nums []int) bool {
	reachPoint := 0
	for i, n := range nums {
		if i > reachPoint {
			return false
		}
		if i+n > reachPoint {
			reachPoint = i + n
		}
	}
	return true
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
