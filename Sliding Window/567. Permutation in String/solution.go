package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	i, j := 0, 0
	set := make(map[byte]int)
	for w := 0; w < len(s1); w++ {
		set[s1[w]]++
	}

	for i < len(s2) && j < len(s2) {
		letter := s2[j]
		if count, isIn := set[letter]; isIn {
			if count == 1 {
				delete(set, letter)
				if len(set) == 0 {
					return true
				}
			} else {
				set[letter]--
			}
			j++
		} else {
			// undo i
			if i < j {
				letter = s2[i]
				set[letter]++
				i++
			} else {
				j++
				i++
			}
		}
	}

	return false
}

func main() {
	s1 := "adc"
	s2 := "dcma"
	fmt.Println(checkInclusion(s1, s2))
}
