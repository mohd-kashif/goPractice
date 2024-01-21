package fibonacci

var fibonacciLookupTable map[int]int

func init() {
	fibonacciLookupTable = make(map[int]int)
	fibonacciLookupTable[0] = 0
	fibonacciLookupTable[1] = 1
}

func GetFibonacciWithDP(number int) int {
	val, ok := fibonacciLookupTable[number]
	if ok {
		return val
	}

	previousFibonacci := GetFibonacciWithDP(number - 1)
	previousToPreviousFibonacci := GetFibonacciWithDP(number - 2)

	fibonacciLookupTable[number-1] = previousFibonacci
	fibonacciLookupTable[number-2] = previousToPreviousFibonacci

	result := previousFibonacci + previousToPreviousFibonacci
	fibonacciLookupTable[number] = result

	return result
}
