package levenshteinDistance

type LevenshteinDistance struct {
}

func (r *LevenshteinDistance) GetLavenshteinDistance(str1 string, str2 string) int {
	lenStr1 := len(str1)
	lenStr2 := len(str2)
	shouldReturn, returnValue := r.checkBaseCondition(lenStr1, lenStr2)
	if shouldReturn {
		return returnValue
	}
	result := 0
	mapStr1 := make(map[rune]int)
	for _, v := range str1 {
		val, ok := mapStr1[v]
		if ok {
			mapStr1[v] = val + 1
		} else {
			mapStr1[v] = 1
		}
	}

	uniqueStr2Count := 0
	uniqueStr1Count := 0
	for _, v := range str2 {
		val, ok := mapStr1[v]
		if ok {
			if val == 1 {
				delete(mapStr1, v)
			} else {
				mapStr1[v] = val - 1
			}
		} else {
			uniqueStr2Count++
		}
	}
	for _, v := range mapStr1 {
		uniqueStr1Count = uniqueStr1Count + v
	}

	if uniqueStr1Count == 0 && uniqueStr2Count == 0 {
		return result
	}
	if uniqueStr1Count == 0 && uniqueStr2Count != 0 {
		return uniqueStr2Count
	}
	if uniqueStr2Count == 0 && uniqueStr1Count != 0 {
		return uniqueStr1Count
	}
	if uniqueStr1Count > uniqueStr2Count {
		return uniqueStr1Count
	}
	if uniqueStr2Count > uniqueStr1Count {
		return uniqueStr2Count
	}
	if uniqueStr1Count == uniqueStr2Count {
		return uniqueStr1Count
	}
	return result
}

func (*LevenshteinDistance) checkBaseCondition(lenStr1 int, lenStr2 int) (bool, int) {
	if lenStr1 == 0 && lenStr2 != 0 {
		return true, lenStr2
	}
	if lenStr2 == 0 && lenStr1 != 0 {
		return true, lenStr1
	}
	if lenStr1 == 0 && lenStr2 == 0 {
		return true, 0
	}
	return false, 0
}
