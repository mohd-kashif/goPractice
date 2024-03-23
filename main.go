package main

import (
	"fmt"
	fileread "goPractice/fileRead"
)

func main() {
	output := fileread.ReadFromFile("/Users/mishipay/Documents/goPractice/fileRead/file.txt")
	fmt.Println(output)
}
