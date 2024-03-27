package generateparenthesis

func GenerateParenthesis(n int) []string {
	result := []string{}
	generateStrings(n, 0, 0, "", &result)
	return result
}

func generateStrings(n int, openCount int, closeCount int, currentString string, result *[]string) {
	if openCount == n && closeCount == n {
		*result = append(*result, currentString)
	}
	if openCount < n {
		generateStrings(n, openCount+1, closeCount, currentString+"(", result)
	}
	if openCount > closeCount {
		generateStrings(n, openCount, closeCount+1, currentString+")", result)
	}
}
