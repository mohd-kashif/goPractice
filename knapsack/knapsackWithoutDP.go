package knapsack

func KnapsackWithoutDP(profits []int, profitsLength int, weights []int, weightsLength int, capacity int, currentIndex int) int {
	if capacity <= 0 || currentIndex >= profitsLength || currentIndex < 0 || weightsLength != profitsLength {
		return 0
	}
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + KnapsackWithoutDP(profits, profitsLength, weights, weightsLength, capacity-weights[currentIndex], currentIndex+1)
	}
	profit2 := KnapsackWithoutDP(profits, profitsLength, weights, weightsLength, capacity, currentIndex+1)

	if profit1 > profit2 {
		return profit1
	}
	return profit2
}
