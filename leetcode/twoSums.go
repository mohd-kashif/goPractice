package leetcode

// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	result := []int{}
	values := make(map[int]int)
	complement := make(map[int]int)
	for k, v := range nums {
		complement[v] = target - v
		values[v] = k
	}
	for k, v := range nums {
		comp := complement[v]
		val, ok := values[comp]
		if ok && k != val {
			result = append(result, k)
			result = append(result, val)
			return result
		}
	}
	return result
}
