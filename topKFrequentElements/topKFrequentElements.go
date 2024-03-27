package topkfrequentelements

import "sort"

// https://leetcode.com/problems/top-k-frequent-elements/
func TopKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		val, ok := hashMap[num]
		if ok {
			hashMap[num] = val + 1
		} else {
			hashMap[num] = 1
		}
	}
	sortedSlice := sortByValue(hashMap)
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, sortedSlice[i].Key)
	}
	return result
}

type KeyValue struct {
	Key   int
	Value int
}

func sortByValue(hashMap map[int]int) []KeyValue {
	keyValueSlice := []KeyValue{}
	for k, v := range hashMap {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}

	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[j].Key > keyValueSlice[i].Key
	})

	return keyValueSlice
}
