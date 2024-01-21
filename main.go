package main

import (
	"fmt"
	lcslength "goPractice/LCSLength"
)

func main() {
	m := "www.educative.io/explore"
	n := "educative.io/edpresso"
	// m = "ababczb"
	// n = "bababz"
	fmt.Println(lcslength.LengthOfLCS(m, n))
}
