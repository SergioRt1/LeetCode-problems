package main

import (
	"fmt"
	"strings"
)

func findLongestWordLength(s string) int {
	return solutionWithoutExtraMemory(s)
}

func solutionWithoutExtraMemory(s string) int {
	best, current := -1, 0
	for _, c := range s {
		if c == ' ' {
			if current > best {
				best = current
			}
			current = 0
		} else {
			current++
		}
	}
	if current > best {
		best = current
	}
	return best
}

func easySolution(s string) int {
	words := strings.Split(s, " ")
	best := -1
	for _, w := range words {
		if len(w) > best {
			best = len(w)
		}
	}

	return best
}

func main() {
	s := "What if we try a super-long word such as otorhinolaryngology"
	fmt.Println(findLongestWordLength(s))
}
