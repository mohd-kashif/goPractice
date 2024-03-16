package main

import (
	"fmt"
	lcslength "goPractice/LCSLength"
)

func main() {
	m := "www.educative.io/exploreeeejasdg"
	n := "educative.io/edpressoooaasd,hao"
	// m = "ababczb"
	// n = "bababz"
	fmt.Println(lcslength.LengthOfLCS(m, n))
}
