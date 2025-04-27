package neetcode150

// https://leetcode.com/problems/contains-duplicate/submissions/1619123759/
func ContainsDuplicate(nums []int) bool {
	values := make(map[int]bool)
	for _, val := range nums {
		exist := values[val]
		if !exist {
			values[val] = true
		} else {
			return true
		}
	}

	return false
}
