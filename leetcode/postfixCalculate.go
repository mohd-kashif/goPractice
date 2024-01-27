package leetcode

import "strconv"

var Operands = map[string]bool{
	"+": true,
	"*": true,
	"-": true,
	"/": true,
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, val := range tokens {
		_, ok := Operands[val]
		if !ok {
			v, _ := strconv.Atoi(val)
			stack = append(stack, v)
		} else {
			lenStack := len(stack)
			val1 := stack[lenStack-1]
			val2 := stack[lenStack-2]
			stack = stack[:lenStack-2]
			res := 0
			if val == "+" {
				res = val2 + val1
			} else if val == "-" {
				res = val2 - val1
			} else if val == "*" {
				res = val2 * val1
			} else {
				res = val2 / val1
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
