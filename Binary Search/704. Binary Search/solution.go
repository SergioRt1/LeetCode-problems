package main

import "fmt"

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0)
}

func binarySearch(nums []int, target int, offset int) int {
	if len(nums) == 0 {
		return -1
	}
	half := len(nums) / 2
	if value := nums[half]; value == target {
		return half + offset
	} else if value < target {
		return binarySearch(nums[half+1:], target, half+1+offset)
	} else {
		return binarySearch(nums[:half], target, offset)
	}
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Print(search(nums, target))
}
