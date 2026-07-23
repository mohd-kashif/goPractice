package main

import (
	"fmt"
	agodaprep "goPractice/agodaPrep"
)

func main() {
	str := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	result := agodaprep.OrangesRotting(str)
	fmt.Println(result)
}
