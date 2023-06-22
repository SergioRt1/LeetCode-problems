package main

import (
	"fmt"
	"sort"
)

type pair struct {
	i int
	j int
}

// Runtime error, too much memory used :(
func threeSum(nums []int) [][]int {
	sums := make(map[int][]pair, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	unic := make(map[pair]struct{})
	for i := range nums {
		for j := range nums {
			p := pair{nums[j], nums[i]}
			if i != j {
				if _, seen := unic[p]; !seen {
					sum := nums[i] + nums[j]
					sums[sum] = append(sums[sum], pair{j, i})
					unic[p] = struct{}{}
				}
			} else {
				break
			}
		}
	}
	for k := range unic {
		delete(unic, k)
	}

	var result [][]int
	for k := range nums {
		com := -nums[k]
		if v, ok := sums[com]; ok {
			for idx := 0; idx < len(v); idx++ {
				p := v[idx]
				if p.i != k && p.j != k {
					r := []int{nums[p.i], nums[p.j], nums[k]}
					sort.Slice(r, func(i, j int) bool {
						return r[i] < r[j]
					})
					newP := pair{r[0], r[1]}
					if _, seen := unic[newP]; !seen {
						result = append(result, r)
						unic[newP] = struct{}{}
						v = fastRemove(v, idx)
						idx--
					}
				}
			}
			sums[com] = v
		}
	}

	return result
}

func fastRemove(s []pair, i int) []pair {
	l := len(s) - 1
	s[i] = s[l]
	return s[:l]
}

func main() {
	nums := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	fmt.Print(threeSum(nums))
}
