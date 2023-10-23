package main

import (
	"fmt"
	"strings"
)

func characterReplacement(s string, k int) int {
	return slidingWindow(s, k)
}

type SuperMap struct {
	data       map[uint8]int
	maxCount   int
	maxCountID uint8
}

func NewSuperMap(size int) *SuperMap {
	return &SuperMap{
		data: make(map[uint8]int, size),
	}
}

func (s *SuperMap) Increment(id uint8) {
	s.data[id]++
	if id == s.maxCountID {
		s.maxCount++
	} else if s.data[id] > s.maxCount {
		s.maxCount = s.data[id]
		s.maxCountID = id
	}
}

func (s *SuperMap) GetMax() int {
	return s.maxCount
}

func (s *SuperMap) Decrement(id uint8) {
	s.data[id]--
	if id == s.maxCountID {
		s.maxCount--
		for k, v := range s.data {
			if v > s.maxCount {
				s.maxCount = v
				s.maxCountID = k
			}
		}
	}
}

// O(n) solution
func slidingWindow(s string, k int) int {
	words := NewSuperMap(26)
	i := 0
	best := 0
	for j := i; j < len(s); j++ {
		words.Increment(s[j])
		longest := words.GetMax()
		for (j-i+1)-longest > k {
			words.Decrement(s[i])
			longest = words.GetMax()
			i++
		}
		if j-i+1 > best {
			best = j - i + 1
		}
	}

	return best

}

// Don't work O(2n^2)
func simpleSolution(s string, k int) int {
	diffLetters := make(map[uint8]struct{}, k+1)

	i := 0
	for ; i < len(s); i++ {
		diffLetters[s[i]] = struct{}{}
		if len(diffLetters) > k+1 {
			break
		}
	}
	normal := findBest(s, k)
	rev := findBest(revString(s), k)

	if normal > rev {
		return normal
	} else {
		return rev
	}
}

func revString(s string) string {
	newStr := strings.Builder{}
	newStr.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newStr.WriteByte(s[i])
	}
	return newStr.String()
}

func findBest(s string, k int) int {
	best := 0
	nextDiff := 0
	for i := 0; i < len(s); i = nextDiff {
		c := s[i]
		currentK := k
		j := i + 1
		setNextDiff := false
		if best >= len(s)-i {
			return best
		}

		for ; j < len(s); j++ {
			if s[j] != c {
				if !setNextDiff {
					setNextDiff = true
					nextDiff = j
				}
				if currentK > 0 {
					currentK--
				} else {
					break
				}
			}
		}

		if best < j-i {
			best = j - i
		}
	}

	return best
}

// Don't work O(26*n*k)
func dpSetup(s string, k int) int {
	var best int
	memo := make([][][]int, 26)
	for char := 0; char < 26; char++ {
		memo[char] = make([][]int, len(s))
		for i := 0; i < len(s); i++ {
			memo[char][i] = make([]int, k+1)
			for j := 0; j < k+1; j++ {
				memo[char][i][j] = -1
			}
		}
	}

	for idx := range s {
		current := dp(idx, k, s[idx]-'A', s, memo)
		if current > best {
			best = current
		}
	}
	return best
}

func dp(i, k int, c uint8, s string, memo [][][]int) int {
	if i >= len(s) {
		return 0
	}
	if memo[c][i][k] != -1 {
		return memo[c][i][k]
	}
	if k > 0 {
		if subCase := memo[c][i][k-1]; subCase != -1 {
			if subCase == 0 {
				memo[c][i][k] = 0
			} else {
				memo[c][i][k] = subCase + dp(i+subCase, 1, c, s, memo)
			}
			return memo[c][i][k]
		}
	}
	if newC := s[i] - 'A'; newC == c {
		memo[c][i][k] = dp(i+1, k, c, s, memo) + 1
	} else if k > 0 {
		memo[c][i][k] = dp(i+1, k-1, c, s, memo) + 1
	} else {
		memo[c][i][k] = 0
	}
	return memo[c][i][k]
}

func main() {
	s := "ACBBBXBBB"
	k := 2
	fmt.Print(characterReplacement(s, k))
}
