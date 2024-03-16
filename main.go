package main

import (
	"fmt"
	lcslength "goPractice/LCSLength"
)

func main() {
	m := "www.educative.io/exploreeee"
	n := "educative.io/edpressoooo"
	// m = "ababczb"
	// n = "bababz"
	fmt.Println(lcslength.LengthOfLCS(m, n))
}
