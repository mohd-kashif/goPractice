package fileread

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFromFile(filePath string) string {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return "error, is file exist"
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		fmt.Println(inputString)
		if readError == io.EOF {
			return "file read ended"
		}
	}
}
