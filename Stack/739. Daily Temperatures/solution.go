package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	temps := make([][]int, 101)
	for idx, temp := range temperatures {
		temps[temp] = append(temps[temp], idx)
	}
	ans := make([]int, len(temperatures))
	for idx, temp := range temperatures {
		ans[idx] = findWarmerTemp(idx, temp, temps)
	}

	return ans
}

func findWarmerTemp(currentDay, temp int, temps [][]int) int {
	bestWarmer := 1 << 31
	found := false

	for warmerT := temp + 1; warmerT <= 100; warmerT++ {
		if len(temps[warmerT]) > 0 {
			bestDay := findWarmerDayBinary(currentDay, temps[warmerT])

			if bestDay != 0 && bestDay < bestWarmer {
				bestWarmer = bestDay
				found = true
			}
			if bestDay == 1 {
				break
			}
		}
	}
	if !found {
		bestWarmer = 0
	}

	return bestWarmer
}

func findWarmerDay(currentDay int, days []int) int {
	for _, day := range days {
		if currentDay < day {
			return day - currentDay
		}
	}

	return 0
}

func findWarmerDayBinary(currentDay int, days []int) int {
	lo := 0
	hi := len(days) - 1
	result := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if days[mid] > currentDay {
			result = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if result != -1 {
		return days[result] - currentDay
	}
	return 0
}

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
}
