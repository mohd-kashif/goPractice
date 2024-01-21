package staircase

func CountWaysWithoutDP(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	result := CountWaysWithoutDP(n-1) + CountWaysWithoutDP(n-2) + CountWaysWithoutDP(n-3)
	return result
}
