package staircase

func CountWaysWithDP(lookupTable map[int]int, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	val, ok := lookupTable[n]
	if ok {
		return val
	}

	step1 := 0
	step2 := 0
	step3 := 0
	val, ok = lookupTable[n-1]
	if ok {
		step1 = val
	} else {
		step1 = CountWaysWithDP(lookupTable, n-1)
		lookupTable[n-1] = step1
	}

	val, ok = lookupTable[n-2]
	if ok {
		step2 = val
	} else {
		step2 = CountWaysWithDP(lookupTable, n-2)
		lookupTable[n-2] = step2
	}

	val, ok = lookupTable[n-3]
	if ok {
		step3 = val
	} else {
		step3 = CountWaysWithDP(lookupTable, n-3)
		lookupTable[n-3] = step3
	}
	result := step1 + step2 + step3
	lookupTable[n] = result
	return result
}

func CoutWays(n int) int {
	lookupTable := make(map[int]int)
	return CountWaysWithDP(lookupTable, n)
}
