package leetcode

import (
	"math"
)

func KEmptySlots(bulbs []int, k int) int {
	n := len(bulbs)
	result := math.MaxInt
	position := make([]int, n+1)
	for i := 0; i < n; i++ {
		position[bulbs[i]] = i + 1
	}
	start := 1
	end := k + 2
	for i := 1; end <= n; i++ {
		if position[i] > position[start] && position[i] > position[end] {
			continue
		}
		if i == end {
			result = getMin(result, getMax(position[start], position[end]))
		}
		start = i
		end = i + k + 1
	}
	if result == math.MaxInt {
		result = -1
	}
	return result
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
