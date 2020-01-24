/*
将整个文件的内容读到一个字符串里
io/ioutil包里的ioutil.ReadFile()方法,
该方法第一个返回值得类型的[]byte,里面存放读取到的内容，第二个返回值是错误，
如果没有错误发生，第二个返回值为nil. 函数WriteFile()可以将[]byte的值写入
文件
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
