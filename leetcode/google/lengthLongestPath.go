package leetcode

import "strings"

func LengthLongestPath(input string) int {
	stack := []int{}
	stack = append(stack, 0)
	maxLength := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		index := strings.LastIndex(line, "\t")
		level := index + 1
		for level < len(stack)-1 {
			stack = stack[:len(stack)-1]
		}
		len := stack[len(stack)-1] + (len(line) - level) + 1
		stack = append(stack, len)
		if strings.Contains(line, ".") {
			if len > maxLength {
				maxLength = len
			}
		}
	}
	if maxLength == 0 {
		return maxLength
	}
	return maxLength - 1
}

// func lengthLongestPath(input string) int {
// 	lines := strings.Split(input, "\n")
// 	maxLen := 0
// 	pathLengths := make(map[int]int)

// 	for _, line := range lines {
// 		level := strings.Count(line, "\t")
// 		name := strings.TrimSpace(strings.Replace(line, "\t", "", -1))
// 		if strings.Contains(name, ".") { // file
// 			path := pathLengths[level-1] + len(name)
// 			if path > maxLen {
// 				maxLen = path
// 			}
// 		} else { // directory
// 			pathLengths[level] = pathLengths[level-1] + len(name) + 1 // plus 1 for the '/'
// 		}
// 	}

// 	return maxLen
// }
