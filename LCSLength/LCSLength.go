package lcslength

import "fmt"

func LengthOfLCS(m string, n string) int {
	rows := len(m)
	cols := len(n)
	lookUpTable := make([][]int, rows+1)
	for i := range lookUpTable {
		lookUpTable[i] = make([]int, cols+1)
	}
	result := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if m[i-1] == n[j-1] {
				lookUpTable[i][j] = lookUpTable[i-1][j-1] + 1
				if lookUpTable[i][j] > result {
					result = lookUpTable[i][j]
				}
			}
		}
	}
	fmt.Println(lookUpTable)
	return result
}
