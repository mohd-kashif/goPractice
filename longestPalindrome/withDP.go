package longestpalindrome

func LongestPalindromeWithDP(str string) int {
	strLen := len(str)
	rows := strLen * strLen
	hashMap := make([][]int, rows)
	for i := range hashMap {
		hashMap[i] = make([]int, rows)
	}
	rs := RecursiveLongestPalindromeWithDP(str, 0, len(str)-1, hashMap)
	return rs
}

func RecursiveLongestPalindromeWithDP(str string, starting int, ending int, hashMap [][]int) int {
	if starting > ending {
		return 0
	}
	if starting == ending {
		return 1
	}
	if hashMap[starting][ending] == 0 {
		if str[starting] == str[ending] {
			hashMap[starting][ending] = 2 + RecursiveLongestPalindromeWithDP(str, starting+1, ending-1, hashMap)
		} else {
			fromLeft := RecursiveLongestPalindromeWithoutDP(str, starting+1, ending)
			fromRight := RecursiveLongestPalindromeWithoutDP(str, starting, ending-1)
			if fromLeft > fromRight {
				hashMap[starting][ending] = fromLeft
			} else {
				hashMap[starting][ending] = fromRight
			}
		}
	}
	return hashMap[starting][ending]
}
