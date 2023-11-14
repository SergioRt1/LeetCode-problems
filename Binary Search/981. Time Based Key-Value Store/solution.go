package main

import "fmt"

type TimeMap struct {
	data map[string][]dataPoint
}

type dataPoint struct {
	Timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]dataPoint),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	list := m.data[key]

	m.data[key] = append(list, dataPoint{
		Timestamp: timestamp,
		value:     value,
	})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	list := m.data[key]

	return binary(list, timestamp)
}

func binary(data []dataPoint, timestamp int) (ans string) {
	lo, hi := 0, len(data)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if midData := data[mid]; midData.Timestamp == timestamp {
			return midData.value
		} else if midData.Timestamp < timestamp { // take the right side
			ans = midData.value
			lo = mid + 1
		} else { // take the left side
			hi = mid - 1
		}
	}

	return
}

func main() {
	obj := Constructor()

	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))

	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))

	fmt.Println(obj.Get("foo", 1))
}
