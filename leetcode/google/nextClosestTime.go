package leetcode

import (
	"sort"
	"strconv"
)

// https://leetcode.com/problems/next-closest-time/
func NextClosestTime(time string) string {
	result := ""
	values := []int{}
	for _, char := range time {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			values = append(values, digit)
		}
	}
	sortedValues := make([]int, len(values))
	copy(sortedValues, values)
	sort.Ints(sortedValues)
	indexes := make(map[int]int)
	for k, v := range sortedValues {
		val, ok := indexes[v]
		if !ok || k > val {
			indexes[v] = k
		}
	}
	for i := 3; i >= 0; i-- {
		val := values[i]
		index := indexes[val]
		higher, found := getHigher(index, sortedValues)
		if !found || isHighest(values[i], i, values[0]) || isInValid(higher, i, values[0]) {
			lower := sortedValues[0]
			values[i] = lower
		} else if found {
			values[i] = higher
			break
		}
	}
	for i := 0; i < 4; i++ {
		if i == 2 {
			result = result + ":"
		}
		result = result + strconv.Itoa(values[i])
	}
	return result
}

func isInValid(higher int, index int, val int) bool {
	if index == 3 && higher > 9 {
		return true
	} else if index == 2 && higher > 5 {
		return true
	} else if index == 1 {
		if val == 2 && higher > 3 {
			return true
		}
		if (val == 0 || val == 1) && higher > 9 {
			return true
		}
	} else if index == 0 && higher > 2 {
		return true
	}
	return false
}

func isHighest(higher int, index int, val int) bool {
	if index == 3 && higher == 9 {
		return true
	} else if index == 2 && higher == 5 {
		return true
	} else if index == 1 {
		if val == 2 && higher == 3 {
			return true
		}
		if (val == 0 || val == 1) && higher == 9 {
			return true
		}
	} else if index == 0 && higher == 2 {
		return true
	}
	return false
}

func getHigher(index int, values []int) (int, bool) {
	if index == 3 {
		return 0, false
	}
	return values[index+1], true
}
