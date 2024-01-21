package leetcode

// https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			l := len(stack)

			if l == 0 {
				return false
			} else {
				pop := stack[l-1]
				stack = stack[:l-1]

				if pop != char {
					return false
				}
			}
		}
	}

	return len(stack) == 0

}
