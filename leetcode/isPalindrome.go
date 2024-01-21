package leetcode

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = convertToLower(s)
	s = removeNonAlphanumeric(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func convertToLower(input string) string {
	return strings.ToLower(input)
}

func removeNonAlphanumeric(input string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	return regex.ReplaceAllString(input, "")
}
