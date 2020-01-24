// 写文件
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return 
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello World!\n"

	for i := 0; i < 10; i ++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
