package main

import "fmt"

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[lo] > nums[hi] { // rotation in current interval
			if (nums[mid] < nums[hi] && (nums[mid] <= target && target <= nums[hi])) || // is in right side !R
				(nums[mid] > nums[hi] && (nums[mid] <= target || target <= nums[hi])) { // is in right side R
				lo = mid + 1
			} else if (nums[lo] < nums[mid] && (nums[lo] <= target && target <= nums[mid])) || // is in left side !R
				(nums[lo] > nums[mid] && (nums[lo] <= target || target <= nums[mid])) { // is in left side R
				hi = mid - 1
			} else { // is not in the array
				return -1
			}
		} else { // do normal binary
			if nums[mid] < target { // take right side
				lo = mid + 1
			} else if target < nums[mid] { // take left side
				hi = mid - 1
			} else {
				return -1
			}
		}
	}
	return -1
}

func searchSimplify(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}

		// Check if the left half is sorted
		if nums[lo] <= nums[mid] {
			// If target is in the left half
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			// The right half must be sorted
			// If target is in the right half
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{5, 5, 7, 7, 1, 2, 3}
	target := 1

	fmt.Println(searchSimplify(nums, target))
}
