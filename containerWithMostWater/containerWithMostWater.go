package containerwithmostwater

func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	area := 0
	maxArea := area
	for i < j {
		min := getMin(height[i], height[j])
		area = min * (j - i)
		if area > maxArea {
			maxArea = area
		}
		if height[i] == height[j] {
			i++
			j--
		} else {
			if height[i] < height[j] {
				i++
			} else {
				j++
			}
		}
	}
	return maxArea
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
