package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMedian(nums2)
	}
	if len(nums2) == 0 {
		return findMedian(nums1)
	}
	if nums1[0] > nums2[0] { // nums1[0] is always the smallest value
		nums1, nums2 = nums2, nums1
	}
	last1 := len(nums1) - 1

	if nums1[last1] < nums2[0] { // separated case: [ ... ] + { ... }
		size := len(nums1) + len(nums2)
		mid := size / 2

		if size&1 == 0 {
			n1 := findMid(nums1, nums2, mid)
			n2 := findMid(nums1, nums2, mid-1)
			return float64(n1+n2) / 2
		} else {
			return float64(findMid(nums1, nums2, mid))
		}
	}

	// partial overlap: [ .. { ... ] .. } and
	// full overlap:    [ .. { ... } .. ]
	return findMedianBinary(nums1, nums2)
}

type binaryList struct {
	L    int
	H    int
	mid  int
	nums []int
}

func (b *binaryList) getMid() int {
	return b.nums[b.mid]
}

func findMedianBinary(nums1 []int, nums2 []int) float64 {
	l1 := &binaryList{
		L:    0,
		H:    len(nums1) - 1,
		nums: nums1,
	}
	l2 := &binaryList{
		L:    0,
		H:    len(nums2) - 1,
		nums: nums2,
	}
	size := len(nums1) + len(nums2)
	idealMid := size / 2

	var currentMid int

	for l1.L <= l1.H && l2.L <= l2.H {
		if l1.H-l1.L < l2.H-l2.L { // l1 is always bigger
			l1, l2 = l2, l1
		}
		l1.mid = l1.L + (l1.H-l1.L)/2
		l2.mid = l2.L + binarySearch(l2.nums[l2.L:l2.H+1], l1.getMid())
		if l2.getMid() < l1.getMid() {
			currentMid = l1.mid + l2.mid + 1
		} else {
			currentMid = l1.mid + l2.mid
		}
		if currentMid == idealMid {
			return solveParity(size, l1, l2)
		} else if currentMid < idealMid {
			l1.L = l1.mid + 1
			l2.L = l2.mid
			if l2.getMid() < l1.getMid() {
				l2.L += 1
			}
		} else { // idealMid < l1[mid]
			l1.H = l1.mid - 1
			l2.H = l2.mid
			if l2.getMid() > l1.getMid() {
				l2.H -= 1
			}
		}

	}
	// validate one l is done but the other not
	if l1.L <= l1.H {
		return solve(l1, l2, size)
	}
	if l2.L <= l2.H {
		return solve(l2, l1, size)
	}

	return 0
}

func solve(l1 *binaryList, l2 *binaryList, size int) float64 {
	if l1.H-l1.L > 0 {
		if l1.nums[l1.H] < l2.nums[0] { // l2 is in the right
			l1.mid = (len(l1.nums) + len(l2.nums)) / 2
		} else if l2.nums[len(l2.nums)-1] < l1.nums[l1.L] { // l2 is in the left
			l1.mid = (len(l1.nums)+len(l2.nums))/2 - len(l2.nums) - 1
		} else { // is in the middle
			l1.mid = l1.L + (l1.H-l1.L)/2
		}
	}
	return solveParity(size, l1, l2)
}

func solveParity(size int, l1 *binaryList, l2 *binaryList) float64 {
	if size&1 == 0 {
		before := getBefore(l1, l2)
		return float64(l1.getMid()+before) / 2
	} else {
		return float64(l1.getMid())
	}
}

// get the value before l1.mid
func getBefore(l1 *binaryList, l2 *binaryList) (before int) {
	candidates := make([]int, 0, 2)
	if l1.mid-1 >= 0 {
		candidates = append(candidates, l1.nums[l1.mid-1])
	}
	if l2.getMid() <= l1.getMid() {
		candidates = append(candidates, l2.getMid())
	} else if l2.mid-1 >= 0 && l2.nums[l2.mid-1] <= l1.getMid() {
		candidates = append(candidates, l2.nums[l2.mid-1])
	}
	before = -1 << 30
	for _, v := range candidates {
		if v > before {
			before = v
		}
	}

	return
}

func binarySearch(nums []int, t int) int {
	lo, hi := 0, len(nums)-1
	lowest := lo
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == t {
			return mid
		} else if nums[mid] < t {
			lowest = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lowest
}

func findMid(nums1 []int, nums2 []int, mid int) int {
	if mid < len(nums1) {
		return nums1[mid]
	} else {
		return nums2[mid-len(nums1)]
	}
}

func findMedian(nums []int) float64 {
	if len(nums) == 0 {
		return 0 // This should never happen
	}
	m := len(nums) / 2
	if len(nums)&1 == 0 {
		return float64(nums[m]+nums[m-1]) / 2
	} else {
		return float64(nums[m])
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{1, 2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
