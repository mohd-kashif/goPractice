package setpartition

func CanPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	if sum%2 != 0 {
		return false
	}
	return PartitionRecursive(nums, sum/2, 0)
}

func PartitionRecursive(nums []int, sum int, currentIndex int) bool {
	length := len(nums)
	if sum == 0 {
		return true
	}
	if currentIndex >= len(nums) || length == 0 {
		return false
	}

	if sum >= nums[currentIndex] {
		if PartitionRecursive(nums, sum-nums[currentIndex], currentIndex+1) {
			return true
		}
	}
	return PartitionRecursive(nums, sum, currentIndex+1)
}
