package main

import (
	"fmt"
)

// subsetsWithDup calculate the power set of the given list of nums in order, accept duplicate numbers
func subsetsWithDup(nums []int) [][]int {
	r, _ := helper(nums)
	return r
}

//TODO use sorting to avoid duplicates

// helper uses a single list and map through the recursion tree to build the answer
func helper(nums []int) ([][]int, map[int]struct{}) {
	//base case
	if len(nums) == 0 {
		return [][]int{{}}, make(map[int]struct{}, 0)
	}
	//recursive decision
	subPowerSet, valuesSeen := helper(nums[1:])
	originalLen := len(subPowerSet)
	for i := 0; i < originalLen; i++ {
		set := subPowerSet[i]
		withCurrent := insertionSort(set, nums[0])

		subPowerSet = addSafe(withCurrent, valuesSeen, subPowerSet)
	}
	return subPowerSet, valuesSeen
}

// insertionSort takes a sorted list and insert an element keeping the order
func insertionSort(list []int, elem int) []int {
	newList := make([]int, 0, len(list)+1)
	notInserted := true
	for i := 0; i < len(list); {
		if notInserted && elem < list[i] {
			newList = append(newList, elem)
			notInserted = false
		} else {
			newList = append(newList, list[i])
			i++
		}
	}
	if notInserted {
		newList = append(newList, elem)
	}
	return newList
}

func addSafe(list []int, valuesSeen map[int]struct{}, response [][]int) [][]int {
	hashElem := hashList(list)
	if _, ok := valuesSeen[hashElem]; !ok {
		valuesSeen[hashElem] = struct{}{}
		response = append(response, list)
	}
	return response
}

const (
	offset32 = 2166136261
	prime32  = 16777619
)

// hashList takes a list and return a 32-bit FNV-1 hash
func hashList(list []int) int {
	hash := offset32
	for _, c := range list {
		hash *= prime32
		hash ^= c
	}
	return hash
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Print(subsetsWithDup(nums))
}
