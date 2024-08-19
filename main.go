package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}
	res := maxPoints(points)
	fmt.Println(res)
}

func maxPoints(points [][]int) int64 {
	rows := len(points)
	cols := len(points[0])

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for j := 0; j < cols; j++ {
		dp[0][j] = points[0][j]
	}

	for i := 1; i < rows; i++ {
		left := make([]int, cols)
		right := make([]int, cols)

		left[0] = dp[i-1][0]
		for j := 1; j < cols; j++ {
			left[j] = max(left[j-1]-1, dp[i-1][j])
		}

		right[cols-1] = dp[i-1][cols-1]
		for j := cols - 2; j >= 0; j-- {
			right[j] = max(right[j+1]-1, dp[i-1][j])
		}

		for j := 0; j < cols; j++ {
			dp[i][j] = points[i][j] + max(left[j], right[j])
		}
	}

	result := math.MinInt64
	for j := 0; j < cols; j++ {
		if dp[rows-1][j] > result {
			result = dp[rows-1][j]
		}
	}

	return int64(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
