package fibonacci

func GetFibonacciWithoutDP(number int) int {
	result := 0
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	result = GetFibonacciWithoutDP(number-1) + GetFibonacciWithoutDP(number-2)
	return result
}
