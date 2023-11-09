package main

import "fmt"

func findMin(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	min := nums[hi]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < min {
			min = nums[mid]
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return min
}

func main() {
	nums := []int{1}
	fmt.Println(findMin(nums))
}
