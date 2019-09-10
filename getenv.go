package main 

import (
	"os"
	"fmt"
)

var (
	//  func Environ() []string 返回表示环境变量的格式的"key=value"的字符串的切片拷贝
	ENV = os.Environ()
	// func Getenv(key string) string 
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	for _, value := range ENV {
		fmt.Println(value)
	}
	fmt.Println(HOME, USER, GOROOT)

	s := "hello"
fmt.Println(s[0])
}