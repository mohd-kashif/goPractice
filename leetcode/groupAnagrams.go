package leetcode

var alphabetMap = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}
var DONE = "DONE"

// https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	strMap := [26]int{}
	for key, str := range strs {
		if str == DONE {
			continue
		}
		intermediateResult := []string{}
		intermediateResult = append(intermediateResult, str)
		strMap = [26]int{}
		for _, v := range str {
			index := alphabetMap[v]
			strMap[index] = strMap[index] + 1
		}
		for j := key + 1; j < len(strs); j++ {
			if len(str) != len(strs[j]) {
				continue
			}
			if IsAnagram(strMap, strs[j]) {
				intermediateResult = append(intermediateResult, strs[j])
				strs[j] = DONE
			}
		}
		result = append(result, intermediateResult)
	}
	return result
}

func IsAnagram(strMap [26]int, str string) bool {
	for _, v := range str {
		index := alphabetMap[v]
		val := strMap[index]
		if val == 0 {
			return false
		}
		strMap[index] = val - 1
	}
	for _, v := range strMap {
		if v > 0 {
			return false
		}
	}
	return true
}
