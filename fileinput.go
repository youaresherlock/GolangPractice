// 文件使用指向os.File类型的指针来表示的,叫做文件句柄
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	// 在一个无限循环中使用ReadSTring('\n')将文件的内容逐行读取出来
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

/*
在之前的例子中，我们看到，Unix 和 Linux 的行结束符是 \n，而 Windows 的行结束符是 \r\n。
在使用 ReadString 和 ReadBytes 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 
就可以了。另外，我们也可以使用 ReadLine() 方法来实现相同的功能。

一旦读取到文件末尾，变量 readerError 的值将变成非空（事实上，常量 io.EOF 的值是 true）
，我们就会执行 return 语句从而退出循环。
*/
