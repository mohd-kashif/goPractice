package longestpalindrome

func LongestPalindromeWithoutDP(str string) int {
	return RecursiveLongestPalindromeWithoutDP(str, 0, len(str)-1)
}

func RecursiveLongestPalindromeWithoutDP(str string, starting int, ending int) int {
	if starting > ending {
		return 0
	}
	if starting == ending {
		return 1
	}
	if len(str) == 0 {
		return 0
	}
	if str[starting] == str[ending] {
		return 2 + RecursiveLongestPalindromeWithoutDP(str, starting+1, ending-1)
	} else {
		fromLeft := RecursiveLongestPalindromeWithoutDP(str, starting+1, ending)
		fromRight := RecursiveLongestPalindromeWithoutDP(str, starting, ending-1)
		if fromLeft > fromRight {
			return fromLeft
		} else {
			return fromRight
		}
	}
}
