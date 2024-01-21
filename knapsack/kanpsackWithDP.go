package knapsack

func KnapsackWithDP(lookupTable [][]int, profits []int, profitsLength int, weights []int, weightsLength int, capacity int, currentIndex int) int {
	if capacity <= 0 || currentIndex >= profitsLength || currentIndex < 0 || weightsLength != profitsLength {
		return 0
	}

	if lookupTable[currentIndex][capacity] != 0 {
		return lookupTable[currentIndex][capacity]
	}
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + KnapsackWithDP(lookupTable, profits, profitsLength, weights, weightsLength, capacity-weights[currentIndex], currentIndex+1)
	}
	profit2 := KnapsackWithDP(lookupTable, profits, profitsLength, weights, weightsLength, capacity, currentIndex+1)

	if profit1 > profit2 {
		lookupTable[currentIndex][capacity] = profit1
		return profit1
	}
	lookupTable[currentIndex][capacity] = profit2
	return profit2
}

func Knapsack(profits []int, profitsLength int, weights []int, weightsLength int, capacity int) int {
	lookupTable := make([][]int, profitsLength)

	for i := 0; i < profitsLength; i++ {
		lookupTable[i] = make([]int, capacity+1)
		for j := 0; j < capacity+1; j++ {
			lookupTable[i][j] = 0
		}
	}
	return KnapsackWithDP(lookupTable, profits, profitsLength, weights, weightsLength, capacity, 0)
}
