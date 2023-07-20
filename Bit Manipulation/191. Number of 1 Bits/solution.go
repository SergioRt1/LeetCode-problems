package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func main() {
	fmt.Print(hammingWeight(11))
}
