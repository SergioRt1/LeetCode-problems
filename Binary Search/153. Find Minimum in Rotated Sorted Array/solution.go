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

func findMin2(nums []int) int {
	lo, hi := 0, len(nums)-1
	min := nums[hi]
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] { // rotation point on the right side
			lo = mid + 1
		} else {
			min = nums[mid]
			hi = mid
		}
	}

	return min
}

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(findMin2(nums))
}
