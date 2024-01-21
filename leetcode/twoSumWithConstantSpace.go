package leetcode

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSumWithConstantSpace(numbers []int, target int) []int {
	result := [2]int{}
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			result[0] = i + 1
			result[1] = j + 1
			return result[:]
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return result[:]
}
