package main

import (
	"fmt"
	containerwithmostwater "goPractice/containerWithMostWater"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(containerwithmostwater.MaxArea(height))

}
